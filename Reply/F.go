package Reply

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
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	p "github.com/qydysky/part"

	"github.com/dustin/go-humanize"
	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	replyFunc "github.com/qydysky/bili_danmu/Reply/F"
	"github.com/qydysky/bili_danmu/Reply/F/danmuXml"
	videoInfo "github.com/qydysky/bili_danmu/Reply/F/videoInfo"
	"github.com/qydysky/bili_danmu/Reply/decoder"
	send "github.com/qydysky/bili_danmu/Send"

	pctx "github.com/qydysky/part/ctx"
	perrors "github.com/qydysky/part/errors"
	file "github.com/qydysky/part/file"
	fctrl "github.com/qydysky/part/funcCtrl"
	pio "github.com/qydysky/part/io"
	plog "github.com/qydysky/part/log/v2"
	pool "github.com/qydysky/part/pool"
	ps "github.com/qydysky/part/slice"
	unsafe "github.com/qydysky/part/unsafe"
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
		flog.W(`已录制 ` + strconv.Itoa(roomid) + ` 不能重复录制`)
		return
	}

	common, _ := c.CommonsLoadOrInit.LoadOrInitPThen(roomid)(func(actual *c.Common, loaded bool) (*c.Common, bool) {
		return actual, loaded
	})

	if tmp, e := NewM4SStream(common); e != nil {
		flog.E(e)
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
			flog.I(`已切片 ` + strconv.Itoa(roomid))
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
func EntryDanmu(common *c.Common) {
	flog := flog.BaseAdd(`进房弹幕`)

	//检查与切换粉丝牌，只在cookie存在时启用
	F.Api.Get(common, `CheckSwitch_FansMedal`)

	if v, _ := common.K_v.LoadV(`进房弹幕_有粉丝牌时才发`).(bool); v && common.Wearing_FansMedal == 0 {
		flog.T(`无粉丝牌`)
		return
	}
	if array, ok := common.K_v.LoadV(`进房弹幕_内容`).([]any); ok && len(array) != 0 {
		if v, _ := common.K_v.LoadV(`进房弹幕_仅发首日弹幕`).(bool); v {
			if func() (skip bool) {
				defer common.Lock()()
				skip = common.EntryDanmuT.IsZero() || time.Now().Day() == common.EntryDanmuT.Day()
				common.EntryDanmuT = time.Now()
				return
			}() {
				flog.T(`初次启动或今日已发弹幕`)
				return
			}
		}
		replyFunc.KeepMedalLight.Run2(func(kmli replyFunc.KeepMedalLightI) {
			kmli.Do(array[p.Rand().MixRandom(0, int64(len(array)-1))].(string))
		})
	} else {
		flog.T(`进房弹幕_内容为空，不发送`)
		return
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
		flog.BaseAdd(`自动送礼`).I(`已完成`)
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
		flog.BaseAdd(`流服务弹幕`).E(err)
		return
	}
	StreamWs.Interface().Push_tag(`send`, websocket.Uinterface{
		Id:   0,
		Data: data,
	})
}

// 节目单
type PlayItem struct {
	Uname         string         `json:"uname"`         // 主播名 // 自动从Live[0]取
	UpUid         int            `json:"upUid"`         // 主播uid // 自动从Live[0]取
	Roomid        int            `json:"roomid"`        // 房间号 // 自动从Live[0]取
	Qn            string         `json:"qn"`            // 画质 // 自动从Live[0]取
	Name          string         `json:"name"`          // 自定义标题
	StartT        string         `json:"startT"`        // 本段起始时间 // 自动从Live[0]取
	StartTS       int64          `json:"-"`             // 本段起始时间unix // 自动从Live取
	EndT          string         `json:"endT"`          // 本段停止时间 // 自动从Live[len(Live)-1]取
	EndTS         int64          `json:"-"`             // 本段停止时间unix // 自动从Live取
	Dur           time.Duration  `json:"-"`             // 本段时长 // 自动从Live取
	Path          string         `json:"path"`          // 自定义目录名
	Format        string         `json:"format"`        // 格式 // 自动从Live[0]取
	Codec         string         `json:"codec"`         // 格式 // 自动从Live[0]取
	StartLiveT    string         `json:"startLiveT"`    // 本场起始时间 // 自动从Live[0]取
	OnlinesPerMin []int          `json:"onlinesPerMin"` // 人数 // 自动从Live取
	Lives         []PlayItemlive `json:"lives,omitempty"`
	Cuts          []PlayCut      `json:"cuts,omitempty"`
}

var (
	ErrMultiDirMatched = perrors.Action("ErrMultiDirMatched")
)

func getLiveDir[T PlayItem | *PlayItem](pareDir string, playitems ...T) error {
	if len(playitems) == 0 {
		return nil
	}

	for dir := range file.Open(pareDir + "/").DirFilesRange(func(fi os.FileInfo) bool {
		return !fi.IsDir()
	}) {
		switch any(playitems[0]).(type) {
		case PlayItem:
			for _, playitem := range ps.Range(any(playitems).([]PlayItem)) {
				for _, live := range ps.Range(playitem.Lives) {
					if live.liveDirExp == nil || !live.liveDirExp.MatchString(dir.Name()) {
						continue
					} else if live.liveDirExpRes && live.LiveDir != dir.SelfName() {
						return ErrMultiDirMatched.New(live.LiveDir + "匹配结果不唯一")
					} else {
						live.liveDirExpRes = true
						live.LiveDir = dir.SelfName()
					}
				}
			}
		case *PlayItem:
			for _, playitem := range any(playitems).([]*PlayItem) {
				for _, live := range ps.Range(playitem.Lives) {
					if live.liveDirExp == nil || !live.liveDirExp.MatchString(dir.Name()) {
						continue
					} else if live.liveDirExpRes && live.LiveDir != dir.SelfName() {
						return ErrMultiDirMatched.New(live.LiveDir + "匹配结果不唯一")
					} else {
						live.liveDirExpRes = true
						live.LiveDir = dir.SelfName()
					}
				}
			}
		}
	}
	return nil
}

type PlayItemlive struct {
	LiveDir string `json:"liveDir"`
	StartT  string `json:"startT,omitempty"` // 只有第一个设置
	Dur     string `json:"dur,omitempty"`    // 只有最后一个设置

	liveDirExp    *regexp.Regexp `json:"-"`
	liveDirExpRes bool           `json:"-"`

	infoDur time.Duration `json:"-"` // 录播dur
}

func (t *PlayItemlive) UnmarshalJSON(b []byte) (e error) {
	var tmp = struct {
		LiveDir string `json:"liveDir"`
		StartT  string `json:"startT,omitempty"` // 只有第一个设置
		Dur     string `json:"dur,omitempty"`    // 只有最后一个设置
	}{}
	e = json.Unmarshal(b, &tmp)
	if e == nil {
		t.LiveDir = tmp.LiveDir
		t.StartT = tmp.StartT
		t.Dur = tmp.Dur
		t.liveDirExp, _ = regexp.Compile(t.LiveDir)
	}
	return
}

type PlayCut struct {
	Title   string        `json:"title"`
	LiveDir string        `json:"liveDir,omitempty"`
	St      string        `json:"st,omitempty"`
	stt     time.Duration `json:"-"`
	Et      string        `json:"et,omitempty"`
	ett     time.Duration `json:"-"`
	Dur     string        `json:"dur,omitempty"`
}

func (t *PlayCut) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Title string `json:"title"`
		St    string `json:"st,omitempty"`
		Et    string `json:"et,omitempty"`
		Dur   string `json:"dur,omitempty"`
	}{
		Title: t.Title,
		St:    fmt.Sprintf("%.1f", parseDuration(t.St).Minutes()),
		Et:    fmt.Sprintf("%.1f", parseDuration(t.Et).Minutes()),
		Dur:   fmt.Sprintf("%.1f", parseDuration(t.Dur).Minutes()),
	})
}

func (t *PlayCut) UnmarshalJSON(b []byte) (e error) {
	var tmp = struct {
		Title   string        `json:"title"`
		LiveDir string        `json:"liveDir,omitempty"`
		St      string        `json:"st,omitempty"`
		stt     time.Duration `json:"-"`
		Et      string        `json:"et,omitempty"`
		ett     time.Duration `json:"-"`
		Dur     string        `json:"dur,omitempty"`
	}{}
	e = json.Unmarshal(b, &tmp)
	if e == nil {
		t.Title = tmp.Title
		t.LiveDir = tmp.LiveDir
		t.St = tmp.St
		t.stt = parseDuration(tmp.St)
		t.Et = tmp.Et
		t.ett = parseDuration(tmp.Et)
		t.Dur = tmp.Dur
		if t.ett > 0 && tmp.Dur == "" {
			t.Dur = (t.ett - t.stt).String()
		}
	}
	return
}

func init() {
	flog := flog.BaseAdd(`直播Web服务`)

	var liveRootDir string
	if s, ok := c.C.K_v.LoadV(`直播流保存位置`).(string); ok && s != "" {
		if strings.HasSuffix(s, "/") || strings.HasSuffix(s, "\\") {
			liveRootDir = s[:len(s)-1]
		} else {
			liveRootDir = s
		}
	}
	if liveRootDir == "" {
		flog.W(`直播流保存位置无效`)
	} else {
		replyFunc.DanmuCountPerMin.Run2(func(dcpmi replyFunc.DanmuCountPerMinI) {
			dcpmi.CheckRoot(liveRootDir)
		})
	}

	if spath, ok := c.C.K_v.LoadV(`直播Web服务路径`).(string); ok {
		if spath[0] != '/' {
			flog.E(`直播Web服务路径错误`)
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

			wf, close := pweb.WithEncoding(w, r)
			_ = f.CopyToIoWriter(wf, pio.CopyConfig{})
			_ = close()
		})

		// 直播流文件弹幕统计apis
		//
		// req body ["1","2","3"]
		c.C.SerF.Store(spath+"danmuCountPerMins", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodPost) {
				return
			}

			if liveRootDir == "" {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
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
			for i, qref := range refs {
				if qref == "" {
					continue
				}

				if e, hasLivsJson, dir, playlists := LiveDirF(liveRootDir, qref); e != nil {
					flog.I("路径解码失败", qref, perrors.ErrorFormat(e, perrors.ErrActionInLineFunc))
					continue
				} else {
					_, _ = w.Write([]byte(`"` + qref + `":`))
					if qref != `now` {
						if e := replyFunc.DanmuCountPerMin.Run(func(dcpmi replyFunc.DanmuCountPerMinI) error {
							var points []int
							_, _ = w.Write([]byte("["))
							if hasLivsJson {
								var (
									totalDur time.Duration
									totalNum int
								)
								for _, live := range playlists[0].Lives {
									if e := dcpmi.GetRec4(dir+"/"+live.LiveDir, &points); e != nil {
										if !errors.Is(e, os.ErrNotExist) {
											flog.W("获取弹幕统计", e)
										}
										break
									} else {
										st, dur := func() (st, dur int) {
											st, dur = int(parseDuration(live.StartT).Minutes()), int(parseDuration(live.Dur).Minutes())
											if st < 0 {
												st = 0
											} else if st > len(points) {
												st = len(points)
											}
											if dur <= 0 || st+dur > len(points) {
												dur = len(points) - st
											}
											return
										}()
										totalDur += live.infoDur
										for i := 0; i < dur && float64(totalNum) < totalDur.Minutes(); i++ {
											if totalNum > 0 || i > 0 {
												_, _ = w.Write([]byte(","))
											}
											_, _ = w.Write([]byte(strconv.Itoa(points[st+i])))
											totalNum += 1
										}
									}
								}
							} else {
								if e := dcpmi.GetRec4(dir, &points); e != nil {
									if !errors.Is(e, os.ErrNotExist) {
										flog.W("获取弹幕统计", e)
									}
								} else {
									for i := 0; i < len(points); i++ {
										_, _ = w.Write([]byte(strconv.Itoa(points[i])))
										if i != len(points)-1 {
											_, _ = w.Write([]byte{','})
										}
									}
								}
							}
							_, _ = w.Write([]byte("]"))
							return nil
						}); e != nil {
							_, _ = w.Write([]byte("[]"))
						}
					} else {
						_, _ = w.Write([]byte("[]"))
					}
					if i < len(refs)-1 {
						_, _ = w.Write([]byte(","))
					}
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
				flog.W(`直播流保存位置无效`)
				return
			}

			if pweb.NotModifiedDur(r, w, time.Second*5) {
				return
			}

			// 获取当前房间的
			var currentStreamO *M4SStream
			c.StreamO.Range(func(key, value any) bool {
				if key != nil && c.C.Roomid == key.(int) {
					currentStreamO = value.(*M4SStream)
					return false
				}
				return true
			})

			qref := r.URL.Query().Get("ref")
			skip, _ := strconv.Atoi(r.URL.Query().Get("skip"))
			size, _ := strconv.Atoi(r.URL.Query().Get("size"))
			sortS := r.URL.Query().Get("sort")

			// 支持ref
			e, _, _, playlists := LiveDirF(liveRootDir, qref)
			if e != nil {
				w.WriteHeader(http.StatusBadRequest)
				flog.I("路径解码失败", qref, perrors.ErrorFormat(e, perrors.ErrActionInLineFunc))
				return
			}

			uname := r.URL.Query().Get("uname")
			startT := r.URL.Query().Get("startT")
			startLiveT := r.URL.Query().Get("startLiveT")

			ps.Del(&playlists, func(t *PlayItem) (del bool) {
				if currentStreamO != nil &&
					currentStreamO.Common().Liveing &&
					filepath.Base(currentStreamO.GetSavePath()) == t.Path {
					t.Name = "Now: " + t.Name
					t.Path = "now"
				}
				return (uname != "" && uname != t.Uname) || (startT != "" && !strings.HasPrefix(t.StartT, startT)) || (startLiveT != "" && !strings.HasPrefix(t.StartLiveT, startLiveT))
			})

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
		})

		// 表情
		c.C.SerF.Store(spath+"emots/", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpFunc(c.C, w, r, http.MethodGet) {
				return
			}

			if u, e := url.Parse(r.Header.Get("Referer")); e == nil {
				qref := u.Query().Get("ref")
				if qref == "now" {
					if replyFunc.DanmuEmotes.Run(func(dei replyFunc.DanmuEmotesI) error {
						emoteDir := dei.GetEmotesDir("")
						defer emoteDir.Close()
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
				} else if e, hasLivsJson, dir, playlists := LiveDirF(liveRootDir, qref); e != nil {
					w.WriteHeader(http.StatusBadRequest)
					flog.I("路径解码失败", qref, perrors.ErrorFormat(e, perrors.ErrActionInLineFunc))
					return
				} else if hasLivsJson {
					if _, playlist := ps.Search(playlists, func(t *PlayItem) bool {
						return t.Path == filepath.Base(qref)
					}); playlist == nil {
						w.WriteHeader(http.StatusNotFound)
					} else {
						switch e := replyFunc.DanmuEmotes.Run(func(dei replyFunc.DanmuEmotesI) error {
							if err := getLiveDir(dir, playlist); err != nil {
								return err
							}
							for _, live := range playlist.Lives {
								f := dei.GetEmotesDir(dir + "/" + live.LiveDir)
								defer f.Close()
								if f, e := f.Open(strings.TrimPrefix(r.URL.Path, spath+"emots/")); e != nil {
									if errors.Is(e, fs.ErrNotExist) {
										continue
									} else {
										return e
									}
								} else if info, e := f.Stat(); e != nil {
									if errors.Is(e, fs.ErrNotExist) {
										continue
									} else {
										return e
									}
								} else if !pweb.NotModified(r, w, info.ModTime()) {
									_, _ = io.Copy(w, f)
									return nil
								} else {
									return nil
								}
							}
							return os.ErrNotExist
						}); e {
						case nil:
						case os.ErrNotExist:
							w.WriteHeader(http.StatusNotFound)
						default:
							w.WriteHeader(http.StatusBadRequest)
						}
					}
				} else if replyFunc.DanmuEmotes.Run(func(dei replyFunc.DanmuEmotesI) error {
					emoteDir := dei.GetEmotesDir(dir)
					defer emoteDir.Close()
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
			} else {
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
			readFile = func(w http.ResponseWriter, liveRoot, videoDirFromRoot string, startT, duration time.Duration, skipHeader, writeLastBuf bool, offsetByte *int) {
				if rawPath, e := url.PathUnescape(videoDirFromRoot); e != nil {
					w.WriteHeader(http.StatusServiceUnavailable)
					flog.I("路径解码失败", videoDirFromRoot)
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
					flog.I("未找到流文件", videoDir)
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
						flog.W(`直播流回放速率不合法:`, e)
						return
					} else {
						speed = s
					}
				}

				if startT+duration != 0 || skipHeader {
					func() {
						type decodeCuter interface {
							CutSeed(reader io.Reader, startT time.Duration, duration time.Duration, w io.Writer, seeker io.Seeker, getIndex func(seedTo time.Duration) (int64, error), skipHeader, writeLastBuf bool) (err error)
							Cut(reader io.Reader, startT time.Duration, duration time.Duration, w io.Writer, skipHeader, writeLastBuf bool) (err error)
							GenFastSeed(reader io.Reader, save func(seedTo time.Duration, cuIndex int64) error) (err error)
						}

						var cuter decodeCuter

						switch videoType {
						case "flv":
							w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%d.flv\"", time.Now().Unix()))
							flvDecoder := decoder.NewFlvDecoder()
							if v, ok := c.C.K_v.LoadV(`flv音视频时间戳容差ms`).(float64); ok && v > 100 {
								flvDecoder.Diff = v
							}
							cuter = flvDecoder
						case "mp4":
							w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%d.mp4\"", time.Now().Unix()))
							fmp4Decoder := decoder.Fmp4DecoderPool.Get()
							defer decoder.Fmp4DecoderPool.Put(fmp4Decoder)
							if v, ok := c.C.K_v.LoadV(`debug模式`).(bool); ok && v {
								fmp4Decoder.Debug = true
							}
							if v, ok := c.C.K_v.LoadV(`fmp4音视频时间戳容差s`).(float64); ok && v > 0.1 {
								fmp4Decoder.AVTDiff = v
							}
							cuter = fmp4Decoder
						default:
							w.WriteHeader(http.StatusServiceUnavailable)
							flog.W(`未配置的视频类型`, videoDir)
							return
						}

						res := pio.WriterWithConfig(w, pio.CopyConfig{BytePerSec: speed, SkipByte: *offsetByte})

						// fastSeed
						if e := replyFunc.VideoFastSeed.Run(func(vfsi replyFunc.VideoFastSeedI) error {

							f := file.Open(videoDir + "0." + videoType)

							if !file.IsExist(videoDir + "0." + videoType + ".fastSeed") {
								flog := flog.BaseAdd("重生成索引")
								defer f.Close()
								if sf, delete, e := vfsi.InitSav(videoDir + "0." + videoType + ".fastSeed"); e != nil {
									flog.E(e)
								} else if e := cuter.GenFastSeed(f, sf); e != nil && !errors.Is(e, io.EOF) {
									flog.E(perrors.ErrorFormat(e, perrors.ErrActionInLineFunc))
									delete()
								}
								_, _ = f.Seek(0, int(file.AtOrigin))
							}

							if gf, e := vfsi.InitGet(videoDir + "0." + videoType + ".fastSeed"); e != nil {
								flog.E(e)
								return e
							} else if e := cuter.CutSeed(f, startT, duration, res, f, gf, skipHeader, writeLastBuf); e != nil && !errors.Is(e, io.EOF) {
								flog.E(perrors.ErrorFormat(e, perrors.ErrActionInLineFunc))
								return e
							}
							return nil
						}); e != nil {
							flog.E(e)
							if e := cuter.Cut(f, startT, duration, res, skipHeader, writeLastBuf); e != nil && !errors.Is(e, io.EOF) {
								flog.I(e)
							}
						}
					}()
				} else if e := f.CopyToIoWriter(w, pio.CopyConfig{BytePerSec: speed, SkipByte: *offsetByte}); e != nil {
					flog.I(perrors.ErrorFormat(e, perrors.ErrActionInLineFunc))
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
				flog.W(`直播流保存位置无效`)
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

			switch ref := r.URL.Query().Get("ref"); ref {
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
						flog.T(r.RemoteAddr, `接入直播`)
						return nil
					}
					stopFunc := func(_ *M4SStream) error {
						flog.T(r.RemoteAddr, `断开直播`)
						return nil
					}

					conn, _ := r.Context().Value(c.C.SerF).(net.Conn)

					// 在客户端存在某种代理时，将有可能无法监测到客户端关闭，这有可能导致goroutine泄漏
					// if to, ok := c.C.K_v.LoadV(`直播流回放限时min`).(float64); ok && to > 0 {
					// 	if e := conn.SetDeadline(time.Now().Add(time.Duration(int(time.Minute) * int(to)))); e != nil {
					// 		flog.W(`设置直播流回放限时min错误`, e)
					// 	}
					// }

					if e := currentStreamO.PusherToHttp(flog, conn, w, r, PusherEvent{startFunc, nil, stopFunc}); e != nil {
						flog.W(e)
					}
				}
			default:
				// 读取区间
				var rangeHeaderNum int
				if rangeHeader := r.Header.Get(`range`); rangeHeader != "" {
					var e error
					if strings.Index(rangeHeader, "bytes=") != 0 {
						w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
						flog.W(`请求的范围不合法:仅支持bytes`)
						return
					} else if strings.Contains(rangeHeader, ",") && strings.Index(rangeHeader, "-") != len(rangeHeader)-1 {
						w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
						flog.W(`请求的范围不合法:仅支持向后范围`)
						return
					} else if rangeHeaderNum, e = strconv.Atoi(string(rangeHeader[6 : len(rangeHeader)-1])); e != nil {
						w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
						flog.W(`请求的范围不合法:`, e)
						return
					}
				}

				st, dur := parseDuration(r.URL.Query().Get("st")), parseDuration(r.URL.Query().Get("dur"))
				if st < 0 || dur < 0 {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				if e, hasLivsJson, dir, playlists := LiveDirF(liveRootDir, ref); e != nil {
					flog.I("路径解码失败", ref, perrors.ErrorFormat(e, perrors.ErrActionInLineFunc))
					w.WriteHeader(http.StatusServiceUnavailable)
					return
				} else if !hasLivsJson {
					flog.T(r.RemoteAddr, `接入录播`)
					defer func(ts time.Time) {
						flog.T(r.RemoteAddr, `断开录播`, time.Since(ts))
					}(time.Now())
					readFile(w, dir+"/", "/", st, dur, false, true, &rangeHeaderNum)
				} else {
					if _, playlist := ps.Search(playlists, func(t *PlayItem) bool { return t.Path == filepath.Base(ref) }); playlist == nil || len(playlist.Lives) == 0 {
						w.WriteHeader(http.StatusNotFound)
						return
					} else {
						if err := getLiveDir(dir, playlist); err != nil {
							flog.W(`解析节目单失败`, dir, perrors.ErrorFormat(err, perrors.ErrActionInLineFunc))
							w.WriteHeader(http.StatusServiceUnavailable)
							return
						}
						flog.T(r.RemoteAddr, `接入录播`)
						defer func(ts time.Time) {
							flog.T(r.RemoteAddr, `断开录播`, time.Since(ts))
						}(time.Now())
						sst, sdur, skipHeader := parseDuration(playlist.Lives[0].StartT)+st, parseDuration(playlist.Lives[len(playlist.Lives)-1].Dur), false
						nodur := dur == 0
						for i := 0; i < len(playlist.Lives) && (nodur || dur > 0); i++ {
							if fi, e := videoInfo.Get.Run(context.Background(), dir+"/"+playlist.Lives[i].LiveDir); e != nil {
								flog.W(`读取节目单元数据失败`, dir+"/"+playlist.Lives[i].LiveDir, e)
								w.WriteHeader(http.StatusServiceUnavailable)
								return
							} else if sst > fi.Dur {
								sst -= fi.Dur
								continue
							} else {
								if i == len(playlist.Lives)-1 && !nodur && sdur != 0 {
									dur = min(dur, sdur)
								}
								readFile(w, dir+"/", playlist.Lives[i].LiveDir+"/", sst, dur, skipHeader, true, &rangeHeaderNum)
								skipHeader = true
								if !nodur {
									dur -= (fi.Dur - sst)
								}
								sst = 0
							}
						}
						return
					}
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

			switch r.URL.Query().Get("ref") {
			case ``:
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusBadRequest)
			case `now`:
				if IsOn("直播Web可以发送弹幕") {
					StreamWs.Interface().Pull_tag(map[string]func(websocket.Uinterface) (disable bool){
						`recv`: func(u websocket.Uinterface) bool {
							if bytes.Equal(u.Data[:2], []byte("%S")) && len(u.Data) > 0 {
								flog.BaseAdd(`流服务弹幕`).I(unsafe.B2S(u.Data[2:]))
								Msg_senddanmu(unsafe.B2S(u.Data[2:]))
							}
							return false
						},
						`close`: func(i websocket.Uinterface) bool { return true },
					})
				}
				//获取通道
				conn := StreamWs.WS(w, r)
				//由通道获取本次会话id，并测试 提示
				<-conn
				//等待会话结束，通道释放
				<-conn
			default:
				if liveRootDir == "" {
					w.Header().Set("Retry-After", "1")
					w.WriteHeader(http.StatusServiceUnavailable)
					flog.W(`直播流保存位置无效`)
					return
				}
				if !IsOn("弹幕回放") {
					w.WriteHeader(http.StatusNotFound)
					return
				}

				ref := r.URL.Query().Get("ref")
				if e, hasLivsJson, dir, playlists := LiveDirF(liveRootDir, ref); e != nil {
					flog.I("路径解码失败", ref, perrors.ErrorFormat(e, perrors.ErrActionInLineFunc))
					w.WriteHeader(http.StatusServiceUnavailable)
					return
				} else if !hasLivsJson {
					if file.Open(dir + "/0.csv").IsExist() {
						if s, closeF := websocket.Plays(func(reg func(filepath string, start, dur time.Duration) error) {
							if e := reg(dir+"/0.csv", 0, 0); e != nil {
								flog.W(`加载弹幕失败`, e)
							}
						}); s == nil {
							w.WriteHeader(http.StatusNotFound)
							return
						} else {
							defer closeF()
							//获取通道
							conn := s.WS(w, r)
							//由通道获取本次会话id，并测试 提示
							<-conn
							//等待会话结束，通道释放
							<-conn
						}
						return
					} else {
						w.WriteHeader(http.StatusNotFound)
						return
					}
				} else {
					if _, playlist := ps.Search(playlists, func(t *PlayItem) bool { return t.Path == filepath.Base(ref) }); playlist == nil || len(playlist.Lives) == 0 {
						w.WriteHeader(http.StatusNotFound)
						return
					} else {
						if s, closeF := websocket.Plays(func(reg func(filepath string, start, dur time.Duration) error) {
							if err := getLiveDir(dir, playlist); err != nil {
								flog.W(`解析节目单失败`, dir, perrors.ErrorFormat(err, perrors.ErrActionInLineFunc))
								w.WriteHeader(http.StatusServiceUnavailable)
								return
							}
							for i, live := range playlist.Lives {
								st, dur := parseDuration(live.StartT), parseDuration(live.Dur)
								// 列表中间的必须填入时长，如未填入，尝试从元数据中获取
								if dur == 0 && i < len(playlist.Lives)-1 {
									if fi, e := videoInfo.Get.Run(context.Background(), dir+"/"+live.LiveDir); e != nil {
										flog.W(`读取节目单元数据失败`, dir+"/"+live.LiveDir, e)
										return
									} else {
										dur = fi.Dur
									}
								}
								if e := reg(dir+"/"+live.LiveDir+"/0.csv", st, dur); e != nil {
									flog.W(`加载节目单弹幕失败`, e)
									return
								}
							}
						}); s == nil {
							w.WriteHeader(http.StatusNotFound)
							return
						} else {
							defer closeF()
							//获取通道
							conn := s.WS(w, r)
							//由通道获取本次会话id，并测试 提示
							<-conn
							//等待会话结束，通道释放
							<-conn
						}
						return
					}
				}
			}
		})

		if s, ok := c.C.K_v.LoadV("直播Web服务路径").(string); ok && s != "" {
			flog.I(`启动于 ` + c.C.Stream_url.String() + s)
		}
	}
}

// 弹幕回放
func StartRecDanmu(ctx context.Context, flog *plog.Log, filePath string) {
	if !IsOn(`仅保存当前直播间流`) || !IsOn("弹幕回放") {
		return
	}
	f := flog.BaseAdd("弹幕回放")
	var Recoder = websocket.Recorder{
		Server: StreamWs,
	}
	if e := Recoder.Start(filePath + "0.csv"); e == nil {
		f.I(`开始`)
	} else {
		f.E(e)
	}

	ctx, done := pctx.WaitCtx(ctx)
	defer done()
	<-ctx.Done()

	f.I(`结束`)

	// 弹幕录制结束
	if _, e := danmuXml.DanmuXml.Run(context.Background(), &filePath); e != nil {
		msglog.E(e)
	}

	// Ass
	replyFunc.Ass.Run2(func(ai replyFunc.AssI) {
		ai.ToAss(filePath)
	})

	// emots
	replyFunc.DanmuEmotes.Run2(func(dei replyFunc.DanmuEmotesI) {
		if e := dei.PackEmotes(filePath); e != nil {
			msglog.E(e)
		}
	})

	Recoder.Stop()
}

func parseDuration(s string) (t time.Duration) {
	t, _ = time.ParseDuration(s)
	return
}

var (
	ErrLiveDirF      = perrors.Action(`ErrLiveDirF`)
	ErrUrl           = ErrLiveDirF.Append(`.ErrUrl`)
	ErrNoDir         = ErrLiveDirF.Append(`.ErrNoDir`)
	ErrPlaylistRead  = ErrLiveDirF.Append(`.ErrPlaylistRead`)
	ErrPlaylistParse = ErrLiveDirF.Append(`.ErrPlaylistParse`)
	ErrDirFiles      = ErrLiveDirF.Append(`.ErrDirFiles`)
	ErrPlayInfoRead  = ErrLiveDirF.Append(`.ErrPlayInfoRead`)
	ErrNotExist      = ErrLiveDirF.Append(`.ErrNotExist`)
)

var (
	parseBuf = pool.NewPoolBlocks[byte]()
)

// qref 以/结尾或为空，指向 目录
//
// 1. 当有节目单json时
//
// hasLivsJson = true, dir = qref目录， refs = 节目单nfos
//
// 2. 当无节目单json时
//
// hasLivsJson = false, dir = qref目录， refs = 目录下的录播infos
//
// qref 不以/结尾，指向 录播目录
//
// 1. 当是节目单json的录播路径时
//
// hasLivsJson = true, dir = qref目录， refs[0] = 节目单info
//
// 2. 当不是节目单json的录播路径时
//
// hasLivsJson = false, dir = 录播路径， refs[0] = 录播目录info
func LiveDirF(liveRootDir, qref string) (e error, hasLivsJson bool, dir string, refs []PlayItem) {
	qref, e = url.PathUnescape(qref)
	if e != nil {
		e = ErrUrl.NewErr(e)
		return
	}
	playlist := file.Open(filepath.Dir(liveRootDir+"/"+qref) + "/1.json").CheckRoot(liveRootDir)
	hasLivsJson = playlist.IsExist()
	// 录播父级目录
	if qref == "" || strings.HasSuffix(qref, "/") {
		if hasLivsJson {
			tmpBuf := parseBuf.Get()
			defer parseBuf.Put(tmpBuf)

			dir = filepath.Dir(liveRootDir + "/" + qref)
			if err := playlist.ReadToBuf(tmpBuf, humanize.KByte, humanize.MByte); err != nil && !errors.Is(err, io.EOF) {
				e = ErrPlaylistRead.NewErr(err)
				return
			} else if err = json.Unmarshal(*tmpBuf, &refs); err != nil {
				e = ErrPlaylistParse.NewErr(err)
				return
			} else {
				// 从子live里获取信息
				if err := getLiveDir(dir, refs...); err != nil {
					e = ErrPlaylistParse.NewErr(err)
					return
				}
				for _, info := range ps.Range(refs) {
					for j, live := range ps.Range(info.Lives) {
						fi, e := videoInfo.Get.Run(context.Background(), dir+"/"+live.LiveDir)
						if e != nil {
							flog.W(`读取节目单元数据失败`, dir+"/"+live.LiveDir, e)
							break
						}

						st, dur := parseDuration(live.StartT), parseDuration(live.Dur)
						live.infoDur = fi.Dur - st
						if dur > 0 {
							live.infoDur = min(live.infoDur, dur)
						}
						info.Dur += live.infoDur

						// 根据cut 的 LiveDir 重新计算开始时刻
						for _, cut := range ps.Range(info.Cuts) {
							if cut.LiveDir == "" {
								continue
							} else if !strings.Contains(live.LiveDir, cut.LiveDir) {
								cut.stt += live.infoDur
							} else {
								cut.St = fmt.Sprint(cut.stt)
								cut.LiveDir = ""
							}
						}
						if info.EndTS > 0 {
							if diff := info.EndTS - fi.StartTS; diff < -1 {
								flog.W(`拼接节目单元数据`, dir+"/"+live.LiveDir, `早于上个视频结束时间s`, diff)
							} else if diff > 1 {
								flog.W(`拼接节目单元数据`, dir+"/"+live.LiveDir, `晚于上个视频结束时间s`, diff)
							}
						}
						if j == 0 {
							sts, _ := time.Parse(time.DateTime, fi.StartT)
							sts = sts.Add(st)
							info.StartT = sts.Format(time.DateTime)
							info.StartTS = sts.Unix()
							info.StartLiveT = fi.StartLiveT
							info.Format = fi.Format
							info.Codec = fi.Codec
							info.Roomid = fi.Roomid
							info.Uname = fi.Uname
							info.Qn = fi.Qn
							info.UpUid = fi.UpUid
						}
						// if j == len(info.Lives)-1 {
						end := func() time.Time {
							sts, _ := time.Parse(time.DateTime, info.StartT)
							return sts.Add(info.Dur)
						}()
						info.EndT = end.Format(time.DateTime)
						info.EndTS = end.Unix()
						// }
						sst, sdur := int(st.Minutes()), int(dur.Minutes())
						if sst < 0 {
							sst = 0
						} else if sst > len(fi.OnlinesPerMin) {
							sst = len(fi.OnlinesPerMin)
						}
						if sdur <= 0 {
							sdur = len(fi.OnlinesPerMin)
						} else if sdur > len(fi.OnlinesPerMin) {
							sdur = len(fi.OnlinesPerMin)
						}
						info.OnlinesPerMin = append(info.OnlinesPerMin, fi.OnlinesPerMin[sst:sdur]...)
					}
				}
			}
			return
		} else {
			pdir := file.Open(liveRootDir + "/" + qref).CheckRoot(liveRootDir)
			if !pdir.IsDir() {
				e = ErrNoDir
				return
			} else if fs, err := pdir.DirFiles(func(fi os.FileInfo) bool { return !fi.IsDir() }); err != nil {
				e = ErrDirFiles.NewErr(err)
				return
			} else {
				dir = filepath.Dir(liveRootDir + "/" + qref)
				for i, n := 0, len(fs); i < n; i++ {
					if info, err := videoInfo.Get.Run(context.Background(), dir+"/"+fs[i].SelfName()); err != nil {
						if !errors.Is(err, os.ErrNotExist) {
							e = ErrPlayInfoRead.NewErr(err)
						}
					} else {
						refs = append(refs, PlayItem{
							Uname:         info.Uname,
							UpUid:         info.UpUid,
							Roomid:        info.Roomid,
							Qn:            info.Qn,
							Name:          info.Name,
							StartT:        info.StartT,
							StartTS:       info.StartTS,
							EndT:          info.EndT,
							EndTS:         info.EndTS,
							Dur:           info.Dur,
							Path:          info.Path,
							Format:        info.Format,
							Codec:         info.Codec,
							StartLiveT:    info.StartLiveT,
							OnlinesPerMin: info.OnlinesPerMin,
						})
					}
				}
				return
			}
		}
	} else if qref != "" {
		// 录播目录
		if hasLivsJson {
			// 为节目单的虚拟录播目录
			tmpBuf := parseBuf.Get()
			defer parseBuf.Put(tmpBuf)

			dir = filepath.Dir(liveRootDir + "/" + qref)
			if err := playlist.ReadToBuf(tmpBuf, humanize.KByte, humanize.MByte); err != nil && !errors.Is(err, io.EOF) {
				e = ErrPlaylistRead.NewErr(err)
			} else if err = json.Unmarshal(*tmpBuf, &refs); err != nil {
				e = ErrPlaylistParse.NewErr(err)
			}
			if i, _ := ps.Search(refs, func(t *PlayItem) bool {
				return t.Path == filepath.Base(qref)
			}); i != -1 {
				refs = refs[i : i+1]
			} else {
				e = ErrNoDir
				return
			}
			// 从子live里获取信息
			if err := getLiveDir(dir, refs[0]); err != nil {
				e = ErrPlaylistParse.NewErr(err)
				return
			}
			for j, live := range ps.Range(refs[0].Lives) {
				fi, err := videoInfo.Get.Run(context.Background(), dir+"/"+live.LiveDir)
				if err != nil {
					flog.W(`读取节目单元数据失败`, dir+"/"+live.LiveDir, err)
					return
				}
				st, dur := parseDuration(live.StartT), parseDuration(live.Dur)
				live.infoDur = fi.Dur - st
				if dur > 0 {
					live.infoDur = min(live.infoDur, dur)
				}
				refs[0].Dur += live.infoDur
				// 根据cut 的 LiveDir 重新计算开始时刻
				for _, cut := range ps.Range(refs[0].Cuts) {
					if cut.LiveDir == "" {
						continue
					} else if !strings.Contains(live.LiveDir, cut.LiveDir) {
						cut.stt += live.infoDur
					} else {
						cut.St = fmt.Sprint(cut.stt)
						cut.LiveDir = ""
					}
				}
				if refs[0].EndTS > 0 {
					if diff := refs[0].EndTS - fi.StartTS; diff < -1 {
						flog.W(`拼接节目单元数据`, dir+"/"+live.LiveDir, `早于上个视频结束时间s`, diff)
					} else if diff > 1 {
						flog.W(`拼接节目单元数据`, dir+"/"+live.LiveDir, `晚于上个视频结束时间s`, diff)
					}
				}
				if j == 0 {
					sts, _ := time.Parse(time.DateTime, fi.StartT)
					sts = sts.Add(st)
					refs[0].StartT = sts.Format(time.DateTime)
					refs[0].StartTS = sts.Unix()
					refs[0].StartLiveT = fi.StartLiveT
					refs[0].Format = fi.Format
					refs[0].Codec = fi.Codec
					refs[0].Roomid = fi.Roomid
					refs[0].Uname = fi.Uname
					refs[0].Qn = fi.Qn
					refs[0].UpUid = fi.UpUid
				}
				end := func() time.Time {
					sts, _ := time.Parse(time.DateTime, refs[0].StartT)
					return sts.Add(refs[0].Dur)
				}()
				refs[0].EndT = end.Format(time.DateTime)
				refs[0].EndTS = end.Unix()
				sst, sdur := int(st.Minutes()), int(dur.Minutes())
				if sst < 0 {
					sst = 0
				} else if sst > len(fi.OnlinesPerMin) {
					sst = len(fi.OnlinesPerMin)
				}
				if sdur <= 0 {
					sdur = len(fi.OnlinesPerMin)
				} else if sdur > len(fi.OnlinesPerMin) {
					sdur = len(fi.OnlinesPerMin)
				}
				refs[0].OnlinesPerMin = append(refs[0].OnlinesPerMin, fi.OnlinesPerMin[sst:sdur]...)
			}
			return
		} else {
			videoDir := file.Open(liveRootDir + "/" + qref + "/").CheckRoot(liveRootDir)
			if !videoDir.IsDir() {
				e = ErrNoDir
				return
			}
			dir = filepath.Dir(liveRootDir + "/" + qref + "/")
			if info, err := videoInfo.Get.Run(context.Background(), liveRootDir+"/"+qref+"/"); err != nil {
				e = ErrPlayInfoRead.NewErr(err)
				return
			} else {
				refs = append(refs, PlayItem{
					Uname:         info.Uname,
					UpUid:         info.UpUid,
					Roomid:        info.Roomid,
					Qn:            info.Qn,
					Name:          info.Name,
					StartT:        info.StartT,
					StartTS:       info.StartTS,
					EndT:          info.EndT,
					EndTS:         info.EndTS,
					Dur:           info.Dur,
					Path:          info.Path,
					Format:        info.Format,
					Codec:         info.Codec,
					StartLiveT:    info.StartLiveT,
					OnlinesPerMin: info.OnlinesPerMin,
				})
				return
			}
		}
	}
	e = ErrNotExist.New(qref)
	return
}
