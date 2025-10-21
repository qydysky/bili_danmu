package reply

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	p "github.com/qydysky/part"
	part "github.com/qydysky/part/log"

	"github.com/dustin/go-humanize"
	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	replyFunc "github.com/qydysky/bili_danmu/Reply/F"
	"github.com/qydysky/bili_danmu/Reply/F/danmuXml"
	videoInfo "github.com/qydysky/bili_danmu/Reply/F/videoInfo"
	send "github.com/qydysky/bili_danmu/Send"

	compress "github.com/qydysky/part/compress"
	pctx "github.com/qydysky/part/ctx"
	perrors "github.com/qydysky/part/errors"
	file "github.com/qydysky/part/file"
	fctrl "github.com/qydysky/part/funcCtrl"
	pio "github.com/qydysky/part/io"
	pweb "github.com/qydysky/part/web"
	websocket "github.com/qydysky/part/websocket"
)

/*
F额外功能区
*/
var flog = c.C.Log.Base(`功能`)

// 功能开关选取函数
func IsOn(s string) bool {
	return c.C.IsOn(s)
}

// 功能区

// 获取实例的Common
func StreamOCommon(roomid int) (array []*c.Common) {
	if roomid != -1 { //返回特定房间
		if v, ok := c.StreamO.Load(roomid); ok {
			return []*c.Common{v.(*M4SStream).Common()}
		}
	} else { //返回所有
		c.StreamO.Range(func(_, v any) bool {
			array = append(array, v.(*M4SStream).Common())
			return true
		})
	}
	return
}

// 获取实例的录制状态
func StreamOStatus(roomid int) (Islive bool) {
	v, ok := c.StreamO.Load(roomid)
	return ok && (!pctx.Done(v.(*M4SStream).Status) || v.(*M4SStream).exitSign.Islive())
}

// 开始实例
func StreamOStart(roomid int) {
	if StreamOStatus(roomid) {
		flog.L(`W: `, `已录制 `+strconv.Itoa(roomid)+` 不能重复录制`)
		return
	}

	common, _ := c.CommonsLoadOrInit.LoadOrInitPThen(roomid)(func(actual *c.Common, loaded bool) (*c.Common, bool) {
		return actual, loaded
	})

	if tmp, e := NewM4SStream(common); e != nil {
		flog.L(`E: `, e)
	} else {
		//实例回调，避免重复录制
		tmp.Callback_start = func(ms *M4SStream) error {
			//流服务添加
			if _, ok := c.StreamO.LoadOrStore(roomid, tmp); ok {
				return fmt.Errorf("已存在此直播间(%d)录制", roomid)
			}
			return nil
		}
		tmp.Callback_stop = func(ms *M4SStream) {
			c.StreamO.Delete(roomid) //流服务去除
		}
		tmp.Start()
	}
}

// 停止实例
func StreamOStopAll() {
	c.StreamO.Range(func(k, v any) bool {
		if !pctx.Done(v.(*M4SStream).Status) {
			v.(*M4SStream).Stop()
		}
		c.StreamO.Delete(k)
		return true
	})
}

// 停止实例
func StreamOStopOther(roomid int) {
	c.StreamO.Range(func(_roomid, v any) bool {
		if roomid == _roomid {
			return true
		}
		if !pctx.Done(v.(*M4SStream).Status) {
			v.(*M4SStream).Stop()
		}
		c.StreamO.Delete(_roomid)
		return true
	})
}

// 停止实例
func StreamOStop(roomid int) {
	if v, ok := c.StreamO.Load(roomid); ok {
		if !pctx.Done(v.(*M4SStream).Status) {
			v.(*M4SStream).Stop()
		}
		c.StreamO.Delete(roomid)
	}
}

// 实例切断
func StreamOCut(roomid int) (setTitle func(title ...string)) {
	if v, ok := c.StreamO.Load(roomid); ok {
		if !pctx.Done(v.(*M4SStream).Status) {
			v.(*M4SStream).Cut()
			flog.L(`I: `, `已切片 `+strconv.Itoa(roomid))
			return func(title ...string) {
				if len(title) > 0 {
					v.(*M4SStream).Common().Title = title[0]
				}
			}
		}
	}
	return func(s ...string) {}
}

// 进入房间发送弹幕
func Entry_danmu(common *c.Common) {
	flog := flog.Base_add(`进房弹幕`)

	//检查与切换粉丝牌，只在cookie存在时启用
	F.Api.Get(common, `CheckSwitch_FansMedal`)

	if v, _ := common.K_v.LoadV(`进房弹幕_有粉丝牌时才发`).(bool); v && common.Wearing_FansMedal == 0 {
		flog.L(`T: `, `无粉丝牌`)
		return
	}
	if v, _ := common.K_v.LoadV(`进房弹幕_仅发首日弹幕`).(bool); v {
		res, e := F.Get_weared_medal(common.Uid, common.UpUid)
		if e != nil {
			return
		}
		if res.TodayIntimacy > 0 {
			flog.L(`T: `, `今日已发弹幕`)
			return
		}
	}
	if array, ok := common.K_v.LoadV(`进房弹幕_内容`).([]any); ok && len(array) != 0 {
		replyFunc.KeepMedalLight.Run2(func(kmli replyFunc.KeepMedalLightI) {
			kmli.Do(array[p.Rand().MixRandom(0, int64(len(array)-1))].(string))
		})
	}
}

var fc_AutoSend_silver_gift fctrl.SkipFunc

// 自动发送即将过期的银瓜子礼物
func AutoSend_silver_gift(common *c.Common) {
	if fc_AutoSend_silver_gift.NeedSkip() {
		return
	} else {
		defer fc_AutoSend_silver_gift.UnSet()
	}

	day, _ := common.K_v.LoadV(`发送还有几天过期的礼物`).(float64)
	if day <= 0 {
		return
	}

	if common.UpUid == 0 {
		F.Api.Get(common, `UpUid`)
	}

	sended := false
	for _, v := range F.Gift_list() {
		if time.Now().Add(time.Hour*time.Duration(24*int(day))).Unix() > int64(v.Expire_at) {
			send.Send_gift(common, v.Gift_id, v.Bag_id, v.Gift_num)
			sended = true
		}
	}

	if sended {
		flog.Base_add(`自动送礼`).L(`I: `, `已完成`)
	}
}

// 直播Web服务口
var StreamWs = websocket.New_server()

func SendStreamWs(item *Danmu_item) {
	var msg string
	if item.auth != nil && !item.hideAuth {
		msg += fmt.Sprint(item.auth) + `: `
	}
	msg += item.msg
	msg = strings.ReplaceAll(msg, "\n", "")
	msg = strings.ReplaceAll(msg, "\\", "\\\\")

	type DataStyle struct {
		Color  string `json:"color"`
		Border bool   `json:"border"`
		Mode   int    `json:"mode"`
	}

	type Data struct {
		Text  string    `json:"text"`
		Style DataStyle `json:"style"`
		Time  float64   `json:"time"`
	}

	var data, err = json.Marshal(Data{
		Text: msg,
		Style: DataStyle{
			Color:  item.color,
			Border: item.border,
			Mode:   item.mode,
		},
	})

	if err != nil {
		flog.Base_add(`流服务弹幕`).L(`E: `, err)
		return
	}
	StreamWs.Interface().Push_tag(`send`, websocket.Uinterface{
		Id:   0,
		Data: data,
	})
}

// 节目单
type PlayItem struct {
	Uname      string         `json:"uname"`      // 主播名 // 自动从Live[0]取
	UpUid      int            `json:"upUid"`      // 主播uid // 自动从Live[0]取
	Roomid     int            `json:"roomid"`     // 房间号 // 自动从Live[0]取
	Qn         string         `json:"qn"`         // 画质 // 自动从Live[0]取
	Name       string         `json:"name"`       // 自定义标题
	StartT     string         `json:"startT"`     // 本段起始时间 // 自动从Live[0]取
	StartTS    int64          `json:"-"`          // 本段起始时间unix // 自动从Live[0]取
	EndT       string         `json:"endT"`       // 本段停止时间 // 自动从Live[len(Live)-1]取
	Path       string         `json:"path"`       // 自定义目录名
	Format     string         `json:"format"`     // 格式 // 自动从Live[0]取
	StartLiveT string         `json:"startLiveT"` // 本场起始时间 // 自动从Live[0]取
	Live       []PlayItemlive `json:"live"`
}

type PlayItemlive struct {
	LiveDir string `json:"liveDir"`
	StartT  string `json:"startT"` // 只有第一个设置
	Dur     string `json:"dur"`    // 只有最后一个设置
}

func init() {
	flog := flog.Base_add(`直播Web服务`)

	var liveRootDir string
	if s, ok := c.C.K_v.LoadV(`直播流保存位置`).(string); ok && s != "" {
		if strings.HasSuffix(s, "/") || strings.HasSuffix(s, "\\") {
			liveRootDir = s[:len(s)-1]
		} else {
			liveRootDir = s
		}
	}

	if spath, ok := c.C.K_v.LoadV(`直播Web服务路径`).(string); ok {
		if spath[0] != '/' {
			flog.L(`E: `, `直播Web服务路径错误`)
			return
		}

		// 直播流回放连接限制
		var climit pweb.Limits
		if limits, ok := c.C.K_v.LoadV(`直播流回放连接限制`).([]any); ok {
			for i := 0; i < len(limits); i++ {
				if vm, ok := limits[i].(map[string]any); ok {
					if cidr, ok := vm["cidr"].(string); !ok {
						continue
					} else if max, ok := vm["max"].(float64); !ok {
						continue
					} else {
						climit.AddLimitItem(pweb.NewLimitItem(int(max)).Cidr(cidr))
					}
				}
			}
		}

		// cache
		var cache pweb.Cache

		// 直播流主页
		c.C.SerF.Store(spath, func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet, http.MethodHead) {
				return
			}

			p := strings.TrimPrefix(r.URL.Path, spath)

			if len(p) == 0 || p[len(p)-1] == '/' {
				p += "index.html"
			}
			if strings.HasSuffix(p, ".js") {
				w.Header().Set("Content-Type", "application/javascript")
			} else if strings.HasSuffix(p, ".map") {
				w.Header().Set("Content-Type", "application/json")
			} else if strings.HasSuffix(p, ".css") {
				w.Header().Set("Content-Type", "text/css")
			} else if strings.HasSuffix(p, ".html") {
				w.Header().Set("Content-Type", "text/html")
			}

			f := file.New("html/streamList/"+p, 0, true).CheckRoot("html/streamList/")
			if !f.IsExist() || f.IsDir() {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			// mod
			if info, e := f.Stat(); e == nil && pweb.NotModified(r, w, info.ModTime()) {
				return
			}

			b, _ := f.ReadAll(humanize.KByte, 10*humanize.MByte)
			b, _ = compress.InGzip(b, 1)
			w.Header().Set("Content-Encoding", "gzip")
			_, _ = w.Write(b)
		})

		// 直播流文件弹幕统计api
		c.C.SerF.Store(spath+"danmuCountPerMin", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet) {
				return
			}
			qref := r.URL.Query().Get("ref")
			if qref == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if liveRootDir == "" {
				w.WriteHeader(http.StatusServiceUnavailable)
				flog.L(`W: `, `直播流保存位置无效`)
				return
			} else if qref == `now` {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte("[]"))
			} else {
				replyFunc.DanmuCountPerMin.Run2(func(dcpmi replyFunc.DanmuCountPerMinI) {
					dcpmi.CheckRoot(liveRootDir)
				})
				v := liveRootDir + "/" + qref + "/"
				if rawPath, e := url.PathUnescape(v); e != nil {
					w.WriteHeader(http.StatusServiceUnavailable)
					flog.L(`I: `, "路径解码失败", v)
					return
				} else {
					v = rawPath
				}
				replyFunc.DanmuCountPerMin.Run2(func(dcpmi replyFunc.DanmuCountPerMinI) {
					if e := dcpmi.GetRec(v, r, w); e != nil && !errors.Is(e, os.ErrNotExist) {
						flog.L(`W: `, "获取弹幕统计", e)
					}
				})
			}
		})

		// 直播流文件弹幕统计apis
		c.C.SerF.Store(spath+"danmuCountPerMins", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodPost) {
				return
			}

			if liveRootDir == "" {
				w.WriteHeader(http.StatusServiceUnavailable)
				flog.L(`W: `, `直播流保存位置无效`)
				return
			} else {
				replyFunc.DanmuCountPerMin.Run2(func(dcpmi replyFunc.DanmuCountPerMinI) {
					dcpmi.CheckRoot(liveRootDir)
				})
			}

			buf := make([]byte, humanize.KByte)
			if n, e := io.ReadFull(r.Body, buf); n == len(buf) && e != io.EOF {
				w.WriteHeader(http.StatusRequestEntityTooLarge)
				return
			} else {
				buf = buf[:n]
			}
			var refs []string
			if e := json.Unmarshal(buf, &refs); e != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")

			_, _ = w.Write([]byte("{"))
			for i := 0; i < len(refs); i++ {
				qref := refs[i]
				if qref == "" {
					continue
				}
				_, _ = w.Write([]byte(`"` + qref + `":`))
				if qref == `now` {
					_, _ = w.Write([]byte("[]"))
				} else {
					qref = liveRootDir + "/" + qref + "/"
					if rawPath, e := url.PathUnescape(qref); e != nil {
						flog.L(`I: `, "路径解码失败", qref)
						continue
					} else {
						qref = rawPath
					}
					if e := replyFunc.DanmuCountPerMin.Run(func(dcpmi replyFunc.DanmuCountPerMinI) error {
						if e := dcpmi.GetRec2(qref, w); e != nil {
							if !errors.Is(e, os.ErrNotExist) {
								flog.L(`W: `, "获取弹幕统计", e)
							}
							_, _ = w.Write([]byte("[]"))
						}
						return nil
					}); e != nil {
						_, _ = w.Write([]byte("[]"))
					}
				}
				if i < len(refs)-1 {
					_, _ = w.Write([]byte(","))
				}
			}
			_, _ = w.Write([]byte("}"))
		})

		// 实时回放模式api
		c.C.SerF.Store(spath+"streamMode", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet) {
				return
			}
			var ms = []string{}
			if modes, ok := c.C.K_v.LoadV(`实时回放预处理`).(map[string]any); ok && len(modes) != 0 {
				for k, v := range modes {
					if len(k) == 0 || k[0] == '_' {
						continue
					}
					if _, ok := v.(map[string]any); ok {
						ms = append(ms, k)
					}
				}
			}
			c.ResStruct{Code: 0, Message: "ok", Data: ms}.Write(w)
		})

		// 直播流文件列表api
		c.C.SerF.Store(spath+"filePath", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet) {
				return
			}

			if liveRootDir == "" {
				c.ResStruct{Code: -1, Message: "直播流保存位置无效", Data: nil}.Write(w)
				flog.L(`W: `, `直播流保存位置无效`)
				return
			}

			//cache
			if bp, ok := cache.IsCache(r.RequestURI); ok {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Cache-Control", "max-age=5")
				_, _ = w.Write(*bp)
				return
			}
			w = cache.Cache(r.RequestURI, time.Second*5, w)

			var filePaths = []*videoInfo.Paf{}

			// 获取当前房间的
			var currentStreamO *M4SStream
			c.StreamO.Range(func(key, value any) bool {
				if key != nil && c.C.Roomid == key.(int) {
					currentStreamO = value.(*M4SStream)
					return false
				}
				return true
			})

			// 支持ref
			v := liveRootDir
			if qref := r.URL.Query().Get("ref"); qref != "" {
				if rawPath, e := url.PathUnescape(qref); e != nil {
					w.WriteHeader(http.StatusServiceUnavailable)
					flog.L(`I: `, "路径解码失败", qref)
					return
				} else {
					if strings.HasSuffix(v, "/") || strings.HasSuffix(v, "\\") {
						v += rawPath
					} else {
						v += "/" + rawPath
					}
				}
			}

			dir := file.New(v, 0, true)
			defer func() {
				_ = dir.Close()
			}()
			if !dir.IsDir() {
				c.ResStruct{Code: -1, Message: "not dir", Data: nil}.Write(w)
				return
			}

			skip, _ := strconv.Atoi(r.URL.Query().Get("skip"))
			size, _ := strconv.Atoi(r.URL.Query().Get("size"))
			sortS := r.URL.Query().Get("sort")

			if playlist := file.Open(v + "/0.json"); playlist.IsExist() {
				if data, e := playlist.ReadAll(humanize.KByte, humanize.MByte); e != nil && !errors.Is(e, io.EOF) {
					w.WriteHeader(http.StatusServiceUnavailable)
					flog.L(`W: `, `节目单错误`, v, e)
				} else {
					var playlists []PlayItem
					if e := json.Unmarshal(data, &playlists); e != nil {
						w.WriteHeader(http.StatusServiceUnavailable)
						flog.L(`W: `, `解析节目单`, v, e)
					} else {
						// 从子live里获取信息
						for i := 0; i < len(playlists); i++ {
							if len(playlists[i].Live) == 0 {
								continue
							}
							if fi, e := videoInfo.Get.Run(context.Background(), v+"/"+playlists[i].Live[0].LiveDir); e != nil {
								flog.L(`W: `, `解析节目单`, v, e)
								continue
							} else {
								playlists[i].StartT = fi.StartT
								playlists[i].StartTS = fi.StartTS
								playlists[i].StartLiveT = fi.StartLiveT
								playlists[i].Format = fi.Format
								playlists[i].Roomid = fi.Roomid
								playlists[i].Uname = fi.Uname
								playlists[i].Qn = fi.Qn
								playlists[i].UpUid = fi.UpUid
							}
							if fi, e := videoInfo.Get.Run(context.Background(), v+"/"+playlists[i].Live[len(playlists[i].Live)-1].LiveDir); e != nil {
								flog.L(`W: `, `解析节目单`, v, e)
								continue
							} else {
								playlists[i].EndT = fi.EndT
							}
						}
						switch sortS {
						case `startTAsc`:
							slices.SortFunc(playlists, func(a, b PlayItem) int {
								return int(a.StartTS - b.StartTS)
							})
						case `startTDsc`:
							slices.SortFunc(playlists, func(a, b PlayItem) int {
								return int(b.StartTS - a.StartTS)
							})
						}
						if skip >= 0 {
							skip = min(skip, len(playlists))
							playlists = playlists[skip:]
						}
						if size >= 0 {
							playlists = playlists[:min(size, len(playlists))]
						}
						c.ResStruct{Code: 0, Message: "ok", Data: playlists}.Write(w)
					}
				}
				return
			}

			if fs, e := dir.DirFiles(func(fi os.FileInfo) bool { return !fi.IsDir() }); e != nil {
				c.ResStruct{Code: -1, Message: e.Error(), Data: nil}.Write(w)
				return
			} else {
				uname := r.URL.Query().Get("uname")
				startT := r.URL.Query().Get("startT")
				startLiveT := r.URL.Query().Get("startLiveT")
				for i, n := 0, len(fs); i < n; i++ {
					if filePath, e := videoInfo.Get.Run(context.Background(), fs[i]); e != nil {
						if !errors.Is(e, os.ErrNotExist) {
							flog.L(`W: `, fs[i], e)
						}
						continue
					} else {
						if uname != "" && uname != filePath.Uname {
							continue
						}
						if startT != "" && !strings.HasPrefix(filePath.StartT, startT) {
							continue
						}
						if startLiveT != "" && !strings.HasPrefix(filePath.StartLiveT, startLiveT) {
							continue
						}
						if currentStreamO != nil &&
							currentStreamO.Common().Liveing &&
							strings.Contains(currentStreamO.GetSavePath(), filePath.Path) {
							filePath.Name = "Now: " + filePath.Name
							filePath.Path = "now"
						}
						filePaths = append(filePaths, filePath)
					}
				}
				switch sortS {
				case `startTAsc`:
					slices.SortFunc(filePaths, func(a, b *videoInfo.Paf) int {
						return int(a.StartTS - b.StartTS)
					})
				case `startTDsc`:
					slices.SortFunc(filePaths, func(a, b *videoInfo.Paf) int {
						return int(b.StartTS - a.StartTS)
					})
				}
				if skip >= 0 {
					skip = min(skip, len(filePaths))
					filePaths = filePaths[skip:]
				}
				if size >= 0 {
					filePaths = filePaths[:min(size, len(filePaths))]
				}
			}
			c.ResStruct{Code: 0, Message: "ok", Data: filePaths}.Write(w)
		})

		// 表情
		c.C.SerF.Store(spath+"emots/", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet) {
				return
			}

			ref := ""
			if u, e := url.Parse(r.Header.Get("Referer")); e == nil {
				ref = u.Query().Get("ref")
				if ref != "now" {
					ref = liveRootDir + "/" + ref + "/"
				} else {
					ref = ""
				}
			}
			if replyFunc.DanmuEmotes.Run(func(dei replyFunc.DanmuEmotesI) error {
				emoteDir := dei.GetEmotesDir(ref)

				if f, e := emoteDir.Open(strings.TrimPrefix(r.URL.Path, spath+"emots/")); e != nil {
					if errors.Is(e, fs.ErrNotExist) {
						w.WriteHeader(http.StatusNotFound)
					} else {
						w.WriteHeader(http.StatusBadRequest)
					}
				} else if info, e := f.Stat(); e != nil {
					w.WriteHeader(http.StatusNotFound)
				} else if !pweb.NotModified(r, w, info.ModTime()) {
					_, _ = io.Copy(w, f)
				}
				return nil
			}) != nil {
				w.WriteHeader(http.StatusNotFound)
			}
		})

		// 直播流播放器
		c.C.SerF.Store(spath+"player/", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet) {
				return
			}

			p := strings.TrimPrefix(r.URL.Path, spath+"player/")

			if len(p) == 0 {
				p = "index.html"
			}

			var s string
			if strings.HasPrefix(p, "emots/") {
				s = "emots/"
				replyFunc.DanmuEmotes.Run2(func(dei replyFunc.DanmuEmotesI) {
					p = dei.Hashr(p)
				})
			} else {
				s = "html/artPlayer/"
				p = "html/artPlayer/" + p
			}

			if strings.HasSuffix(p, ".js") {
				w.Header().Set("content-type", "application/javascript")
			} else if strings.HasSuffix(p, ".css") {
				w.Header().Set("content-type", "text/css")
			} else if strings.HasSuffix(p, ".html") {
				w.Header().Set("content-type", "text/html")
			}

			f := file.New(p, 0, true).CheckRoot(s)
			if !f.IsExist() {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			// mod
			if info, e := f.Stat(); e == nil && pweb.NotModified(r, w, info.ModTime()) {
				return
			}

			b, _ := f.ReadAll(humanize.KByte, humanize.MByte)
			_, _ = w.Write(b)
		})

		// 对于经过代理层，有可能浏览器标签页已经关闭，但代理层不关闭连接，导致连接不能释放
		var expirer = pweb.NewExprier(0)
		if v, ok := c.C.K_v.LoadV(`直播流回放连接检查`).(float64); ok && v > 0 {
			expirer.SetMax(int(v))
		}

		c.C.SerF.Store(spath+"keepAlive", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet) {
				return
			}
			if key, e := expirer.Reg(time.Second*30, r.URL.Query().Get("key")); e != nil {
				w.WriteHeader(http.StatusForbidden)
			} else {
				_, _ = w.Write([]byte(key))
			}
		})

		var (
			parseDuration = func(s string) (t time.Duration) {
				t, _ = time.ParseDuration(s)
				return
			}
			readFile = func(w http.ResponseWriter, liveRoot, videoDirFromRoot string, startT, duration time.Duration, skipHeader, writeLastBuf bool, offsetByte *int) {
				if rawPath, e := url.PathUnescape(videoDirFromRoot); e != nil {
					w.WriteHeader(http.StatusServiceUnavailable)
					flog.L(`I: `, "路径解码失败", videoDirFromRoot)
					return
				} else {
					videoDirFromRoot = rawPath
				}

				videoDir := liveRoot + videoDirFromRoot
				videoType := ""

				if file.New(videoDir+"0.flv", 0, true).CheckRoot(liveRoot).IsExist() {
					videoType = "flv"
					w.Header().Set("Content-Type", "flv-application/octet-stream")
				} else if file.New(videoDir+"0.mp4", 0, true).CheckRoot(liveRoot).IsExist() {
					videoType = "mp4"
					w.Header().Set("Content-Type", "video/mp4")
				} else {
					w.Header().Set("Retry-After", "1")
					w.WriteHeader(http.StatusServiceUnavailable)
					flog.L(`I: `, "未找到流文件", videoDir)
					return
				}

				f := file.Open(videoDir + "0." + videoType).CheckRoot(liveRoot)
				defer f.CloseErr()

				if fsize, e := f.GetFileSize(); e == nil && fsize <= int64(*offsetByte) {
					*offsetByte -= int(fsize)
					return
				}

				// 直播流回放速率
				var speed, _ = humanize.ParseBytes("1 M")
				if rc, ok := c.C.K_v.LoadV(`直播流回放速率`).(string); ok {
					if s, e := humanize.ParseBytes(rc); e != nil {
						w.WriteHeader(http.StatusServiceUnavailable)
						flog.L(`W: `, `直播流回放速率不合法:`, e)
						return
					} else {
						speed = s
					}
				}

				if startT+duration != 0 || skipHeader {
					type decodeCuter interface {
						CutSeed(reader io.Reader, startT time.Duration, duration time.Duration, w io.Writer, seeker io.Seeker, getIndex func(seedTo time.Duration) (int64, error), skipHeader, writeLastBuf bool) (err error)
						Cut(reader io.Reader, startT time.Duration, duration time.Duration, w io.Writer, skipHeader, writeLastBuf bool) (err error)
					}

					var cuter decodeCuter

					switch videoType {
					case "flv":
						w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%d.flv\"", time.Now().Unix()))
						flvDecoder := NewFlvDecoder()
						if v, ok := c.C.K_v.LoadV(`flv音视频时间戳容差ms`).(float64); ok && v > 100 {
							flvDecoder.Diff = v
						}
						cuter = flvDecoder
					case "mp4":
						w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%d.mp4\"", time.Now().Unix()))
						fmp4Decoder := NewFmp4Decoder()
						if v, ok := c.C.K_v.LoadV(`debug模式`).(bool); ok && v {
							fmp4Decoder.Debug = true
						}
						if v, ok := c.C.K_v.LoadV(`fmp4音视频时间戳容差s`).(float64); ok && v > 0.1 {
							fmp4Decoder.AVTDiff = v
						}
						cuter = fmp4Decoder
					default:
						w.WriteHeader(http.StatusServiceUnavailable)
						flog.L(`W: `, `未配置的视频类型`, videoDir)
						return
					}

					res := pio.WriterWithConfig(w, pio.CopyConfig{BytePerSec: speed, SkipByte: *offsetByte})

					// fastSeed
					if file.IsExist(videoDir + "0." + videoType + ".fastSeed") {
						if e := replyFunc.VideoFastSeed.Run(func(vfsi replyFunc.VideoFastSeedI) error {
							if gf, e := vfsi.InitGet(videoDir + "0." + videoType + ".fastSeed"); e != nil {
								flog.L(`E: `, e)
								return e
							} else if e := cuter.CutSeed(f, startT, duration, res, f, gf, skipHeader, writeLastBuf); e != nil && !errors.Is(e, io.EOF) {
								flog.L(`E: `, e)
								return e
							}
							return nil
						}); e != nil {
							flog.L(`E: `, e)
							if e := cuter.Cut(f, startT, duration, res, skipHeader, writeLastBuf); e != nil && !errors.Is(e, io.EOF) {
								flog.L(`I: `, e)
							}
						}
					} else if e := cuter.Cut(f, startT, duration, res, skipHeader, writeLastBuf); e != nil && !errors.Is(e, io.EOF) {
						flog.L(`I: `, e)
					}
				} else if e := f.CopyToIoWriter(w, pio.CopyConfig{BytePerSec: speed, SkipByte: *offsetByte}); e != nil {
					flog.L(`I: `, perrors.ErrorFormat(e, perrors.ErrActionInLineFunc))
				}
			}
		)

		// 流地址
		c.C.SerF.Store(spath+"stream", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet, http.MethodPost) {
				return
			}

			if liveRootDir == "" {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusServiceUnavailable)
				flog.L(`W: `, `直播流保存位置无效`)
				return
			}

			// 直播流回放连接限制
			if climit.AddCount(r) {
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}

			// 检查key
			{
				var checkKey = true

				if v, ok := c.C.K_v.LoadV(`直播流回放连接检查忽略key`).([]any); ok && len(v) != 0 {
					for i := 0; i < len(v); i++ {
						if s, ok := v[i].(string); ok && s != "" && r.URL.Query().Get("key") == s {
							checkKey = false
							break
						}
					}
				}

				if checkKey {
					if e := expirer.LoopCheck(r.Context(), r.URL.Query().Get("key"), func(key string, e error) {
						_ = c.C.SerF.GetConn(r).Close()
					}); e != nil {
						w.WriteHeader(http.StatusTooManyRequests)
						return
					}
				}
			}

			//header
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Connection", "keep-alive")
			w.Header().Set("Content-Transfer-Encoding", "binary")

			switch r.URL.Query().Get("ref") {
			case ``:
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusBadRequest)
			case `now`:
				// 获取当前房间的
				var currentStreamO *M4SStream
				c.StreamO.Range(func(key, value any) bool {
					if key != nil && c.C.Roomid == key.(int) {
						currentStreamO = value.(*M4SStream)
						return false
					}
					return true
				})

				// 未准备好
				if currentStreamO == nil || pctx.Done(currentStreamO.Status) {
					w.Header().Set("Retry-After", "1")
					w.WriteHeader(http.StatusNotFound)
					return
				}

				// w.WriteHeader(http.StatusOK)

				// 推送数据
				{
					startFunc := func(_ *M4SStream) error {
						flog.L(`T: `, r.RemoteAddr, `接入直播`)
						return nil
					}
					stopFunc := func(_ *M4SStream) error {
						flog.L(`T: `, r.RemoteAddr, `断开直播`)
						return nil
					}

					conn, _ := r.Context().Value(c.C.SerF).(net.Conn)

					// 在客户端存在某种代理时，将有可能无法监测到客户端关闭，这有可能导致goroutine泄漏
					// if to, ok := c.C.K_v.LoadV(`直播流回放限时min`).(float64); ok && to > 0 {
					// 	if e := conn.SetDeadline(time.Now().Add(time.Duration(int(time.Minute) * int(to)))); e != nil {
					// 		flog.L(`W: `, `设置直播流回放限时min错误`, e)
					// 	}
					// }

					if e := currentStreamO.PusherToHttp(flog, conn, w, r, PusherEvent{startFunc, nil, stopFunc}); e != nil {
						flog.L(`W: `, e)
					}
				}
			default:
				// 读取区间
				var rangeHeaderNum int
				if rangeHeader := r.Header.Get(`range`); rangeHeader != "" {
					var e error
					if strings.Index(rangeHeader, "bytes=") != 0 {
						w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
						flog.L(`W: `, `请求的范围不合法:仅支持bytes`)
						return
					} else if strings.Contains(rangeHeader, ",") && strings.Index(rangeHeader, "-") != len(rangeHeader)-1 {
						w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
						flog.L(`W: `, `请求的范围不合法:仅支持向后范围`)
						return
					} else if rangeHeaderNum, e = strconv.Atoi(string(rangeHeader[6 : len(rangeHeader)-1])); e != nil {
						w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
						flog.L(`W: `, `请求的范围不合法:`, e)
						return
					}
				}

				ref, st, dur := r.URL.Query().Get("ref"), parseDuration(r.URL.Query().Get("st")), parseDuration(r.URL.Query().Get("dur"))
				if st < 0 || dur < 0 {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				if f := file.Open(liveRootDir + "/" + ref + "/0.json").CheckRoot(liveRootDir); f.IsExist() {
					flog.L(`T: `, r.RemoteAddr, `接入录播`)
					defer func(ts time.Time) {
						flog.L(`T: `, r.RemoteAddr, `断开录播`, time.Since(ts))
					}(time.Now())
					readFile(w, liveRootDir+"/", ref+"/", st, dur, false, true, &rangeHeaderNum)
				} else if f = file.Open(filepath.Dir(liveRootDir+"/"+ref) + "/0.json"); f.IsExist() {
					// 节目单
					var playlists []PlayItem
					if data, e := f.ReadAll(humanize.KByte, humanize.MByte); e != nil && !errors.Is(e, io.EOF) {
						w.WriteHeader(http.StatusServiceUnavailable)
						flog.L(`W: `, `读取节目单`, e)
						return
					} else if e := json.Unmarshal(data, &playlists); e != nil {
						w.WriteHeader(http.StatusServiceUnavailable)
						flog.L(`W: `, `解析节目单`, e)
						return
					}

					if playlistI := slices.IndexFunc(playlists, func(i PlayItem) bool {
						return i.Path == ref
					}); playlistI > -1 {
						v := playlists[playlistI]
						if len(v.Live) > 0 {
							flog.L(`T: `, r.RemoteAddr, `接入录播`)
							defer func(ts time.Time) {
								flog.L(`T: `, r.RemoteAddr, `断开录播`, time.Since(ts))
							}(time.Now())
							sst, sdur, skipHeader := parseDuration(v.Live[0].StartT)+st, parseDuration(v.Live[len(v.Live)-1].Dur), false
							nodur := dur == 0
							for i := 0; i < len(v.Live) && (nodur || dur > 0); i++ {
								if fi, e := videoInfo.Get.Run(context.Background(), liveRootDir+"/"+v.Live[i].LiveDir); e != nil {
									flog.L(`W: `, `读取元数据`, liveRootDir+"/"+v.Live[i].LiveDir, e)
									w.WriteHeader(http.StatusServiceUnavailable)
									return
								} else if sst > fi.Dur {
									sst -= fi.Dur
									continue
								} else {
									if i == len(v.Live)-1 && !nodur && sdur != 0 {
										dur = min(dur, sdur)
									}
									readFile(w, liveRootDir+"/", v.Live[i].LiveDir+"/", sst, dur, skipHeader, true, &rangeHeaderNum)
									skipHeader = true
									if !nodur {
										dur -= (fi.Dur - sst)
									}
									sst = 0
								}
							}
						} else {
							w.WriteHeader(http.StatusNotFound)
							return
						}
					} else {
						w.WriteHeader(http.StatusNotFound)
						return
					}
				} else {
					w.WriteHeader(http.StatusNotFound)
					return
				}
			}
		})

		// 弹幕回放
		c.C.SerF.Store(spath+"player/ws", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet) {
				return
			}

			// 直播流回放连接限制
			if climit.AddCount(r) {
				w.WriteHeader(http.StatusTooManyRequests)
				_, _ = w.Write([]byte(http.StatusText(http.StatusTooManyRequests)))
				return
			}

			var rpath string

			if qref := r.URL.Query().Get("ref"); rpath == "" && qref != "" {
				rpath = "/" + qref + "/"
			}

			if rpath == "" {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			if rpath != `/now/` {
				if liveRootDir == "" {
					w.Header().Set("Retry-After", "1")
					w.WriteHeader(http.StatusServiceUnavailable)
					flog.L(`W: `, `直播流保存位置无效`)
					return
				}
				var v string = liveRootDir + rpath

				if !file.New(v+"0.csv", 0, true).CheckRoot(liveRootDir).IsExist() {
					w.WriteHeader(http.StatusNotFound)
					return
				}

				if s, closeF := PlayRecDanmu(v + "0.csv"); s == nil {
					w.WriteHeader(http.StatusNotFound)
					return
				} else {
					defer closeF()
					defer s.Interface().Pull_tag(map[string]func(any) (disable bool){
						`error`: func(a any) (disable bool) {
							flog.L(`T: `, a)
							return false
						},
					})()
					//获取通道
					conn := s.WS(w, r)
					//由通道获取本次会话id，并测试 提示
					<-conn
					//等待会话结束，通道释放
					<-conn
				}
				return
			} else if IsOn("直播Web可以发送弹幕") {
				StreamWs.Interface().Pull_tag(map[string](func(any) bool){
					`recv`: func(i any) bool {
						if u, ok := i.(websocket.Uinterface); ok {
							if bytes.Equal(u.Data[:2], []byte("%S")) && len(u.Data) > 0 {
								flog.Base_add(`流服务弹幕`).L(`I: `, string(u.Data[2:]))
								Msg_senddanmu(string(u.Data[2:]))
							}
						}
						return false
					},
					`close`: func(i any) bool { return true },
				})
			}

			//获取通道
			conn := StreamWs.WS(w, r)
			//由通道获取本次会话id，并测试 提示
			<-conn
			//等待会话结束，通道释放
			<-conn
		})

		// 弹幕回放xml
		c.C.SerF.Store(spath+"player/xml", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet) {
				return
			}

			var rpath string

			if qref := r.URL.Query().Get("ref"); rpath == "" && qref != "" {
				rpath = "/" + qref + "/"
			}

			if rpath == "" {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}

			if liveRootDir == "" {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusServiceUnavailable)
				flog.L(`W: `, `直播流保存位置无效`)
			} else {
				var v string = liveRootDir + rpath

				if !file.New(v+"0.xml", 0, true).CheckRoot(liveRootDir).IsExist() {
					if !file.New(v+"0.csv", 0, true).CheckRoot(liveRootDir).IsExist() {
						w.WriteHeader(http.StatusNotFound)
						return
					}
					if _, e := danmuXml.DanmuXml.Run(context.Background(), &v); e != nil {
						msglog.L(`E: `, e)
					}
				}

				if e := file.New(v+"0.xml", 0, true).CheckRoot(liveRootDir).CopyToIoWriter(w, pio.CopyConfig{}); e != nil {
					flog.L(`W: `, e)
				}
			}
		})

		if s, ok := c.C.K_v.LoadV("直播Web服务路径").(string); ok && s != "" {
			flog.L(`I: `, `启动于 `+c.C.Stream_url.String()+s)
		}
	}
}

// 弹幕回放
func StartRecDanmu(ctx context.Context, flog *part.Log_interface, filePath string) {
	if !IsOn(`仅保存当前直播间流`) || !IsOn("弹幕回放") {
		return
	}
	f := flog.Base_add("弹幕回放")
	var Recoder = websocket.Recorder{
		Server: StreamWs,
	}
	if e := Recoder.Start(filePath + "0.csv"); e == nil {
		f.L(`I: `, `开始`)
	} else {
		f.L(`E: `, e)
	}

	ctx, done := pctx.WaitCtx(ctx)
	defer done()
	<-ctx.Done()

	f.L(`I: `, `结束`)

	// 弹幕录制结束
	if _, e := danmuXml.DanmuXml.Run(context.Background(), &filePath); e != nil {
		msglog.L(`E: `, e)
	}

	// Ass
	replyFunc.Ass.Run2(func(ai replyFunc.AssI) {
		ai.ToAss(filePath)
	})

	// emots
	replyFunc.DanmuEmotes.Run2(func(dei replyFunc.DanmuEmotesI) {
		if e := dei.PackEmotes(filePath); e != nil {
			msglog.L(`E: `, e)
		}
	})

	Recoder.Stop()
}

func PlayRecDanmu(filePath string) (*websocket.Server, func()) {
	if !IsOn(`仅保存当前直播间流`) || !IsOn("弹幕回放") {
		return nil, nil
	}
	return websocket.Play(filePath)
}
