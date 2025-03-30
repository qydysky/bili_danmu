package reply

import (
	"context"
	"crypto/md5"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"math"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dustin/go-humanize"
	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"

	replyFunc "github.com/qydysky/bili_danmu/Reply/F"
	videoInfo "github.com/qydysky/bili_danmu/Reply/F/videoInfo"
	pctx "github.com/qydysky/part/ctx"
	perrors "github.com/qydysky/part/errors"
	file "github.com/qydysky/part/file"
	funcCtrl "github.com/qydysky/part/funcCtrl"
	pio "github.com/qydysky/part/io"
	log "github.com/qydysky/part/log"
	msgq "github.com/qydysky/part/msgq"
	pool "github.com/qydysky/part/pool"
	reqf "github.com/qydysky/part/reqf"
	signal "github.com/qydysky/part/signal"
	slice "github.com/qydysky/part/slice"
	pstring "github.com/qydysky/part/strings"
	pu "github.com/qydysky/part/util"
	pweb "github.com/qydysky/part/web"
)

const (
	defaultStartCount uint = 3 //直播流接收n帧才保存,默认值
)

type M4SStream struct {
	Status               context.Context       //IsLive()是否运行中
	exitSign             *signal.Signal        //IsLive()是否等待退出中
	log                  *log.Log_interface    //日志
	config               M4SStream_Config      //配置
	stream_last_modified time.Time             //流地址更新时间
	stream_type          string                //流类型
	Stream_msg           *msgq.MsgType[[]byte] //流数据消息 tag:data
	first_buf            []byte                //m4s起始块 or flv起始块
	frameCount           uint                  //关键帧数量
	boot_buf             []byte                //快速启动缓冲
	boot_buf_locker      sync.RWMutex
	m4s_pool             *pool.Buf[m4s_link_item] //切片pool
	common               *c.Common                //通用配置副本
	Current_save_path    string                   //明确的直播流保存目录
	// 事件周期 start: 开始实例 startRec：开始录制 load：接收到视频头
	// keyFrame: 接收到关键帧 cut：切 stopRec：结束录制 stop：结束实例
	msg               *msgq.MsgType[*M4SStream] //实例的各种事件回调
	Callback_start    func(*M4SStream) error    //实例开始的回调
	Callback_startRec func(*M4SStream) error    //录制开始的回调
	Callback_stopRec  func(*M4SStream)          //录制结束的回调
	Callback_stop     func(*M4SStream)          //实例结束的回调
	reqPool           *pool.Buf[reqf.Req]       //请求池
}

type M4SStream_Config struct {
	save_path    string //直播流保存目录
	want_qn      int    //直播流清晰度
	want_type    string //直播流类型
	save_to_file bool   //保存到文件
}

type m4s_link_item struct {
	Url          string           // m4s链接
	Base         string           // m4s文件名
	isHeader     bool             // m4sHeader
	status       int              // 下载状态 0:未下载 1:正在下载 2:下载完成 3:下载失败
	err          error            // 下载状态 3:下载失败 时的错误
	tryDownCount int              // 下载次数 当=3时，不再下载，忽略此块
	data         *slice.Buf[byte] // 下载的数据
	createdTime  time.Time        // 创建时间
	pooledTime   time.Time        // 到pool时间
	SerUuid      string           // 使用的流服务器uuid
}

// 更换服务器
func (t *m4s_link_item) replaceSer(v *c.LiveQn) {
	t.SerUuid = v.Uuid
	t.Url = F.ResolveReferenceLast(v.Url, t.Base+"?trid="+F.ParseQuery(v.Url, "trid="))
	t.status = 0
}

func (t *m4s_link_item) copyTo(to *m4s_link_item) {
	// fmt.Println("copy to ", t.Base)
	to.Url = t.Url
	to.isHeader = t.isHeader
	to.Base = t.Base
	to.status = t.status
	to.tryDownCount = t.tryDownCount
	to.createdTime = t.createdTime
}

func (t *m4s_link_item) reset() *m4s_link_item {
	if t == nil {
		return t
	}
	t.Url = ""
	t.isHeader = false
	t.Base = ""
	t.status = 0
	t.tryDownCount = 0
	t.data.Reset()
	t.createdTime = time.Now()
	return t
}

func (t *m4s_link_item) isInit() bool {
	return t.isHeader
}

func (t *m4s_link_item) getNo() (int, error) {
	var base = t.Base
	if t.Base == "" {
		return 0, nil
	}
	if t.isInit() {
		base = base[1:]
	}
	return strconv.Atoi(base[:len(base)-4])
}

var (
	ActionErrFmp4DownloadCareTO perrors.Action = `ActionErrFmp4DownloadCareTO`
)

func (link *m4s_link_item) download(reqPool *pool.Buf[reqf.Req], reqConfig reqf.Rval) (err error) {
	link.status = 1 // 设置切片状态为正在下载
	link.err = nil
	link.tryDownCount += 1
	link.data.Reset()

	r := reqPool.Get()
	defer reqPool.Put(r)
	reqConfig.Url = link.Url

	if e := r.Reqf(reqConfig); e != nil && !errors.Is(e, io.EOF) {
		link.status = 3 // 设置切片状态为下载失败
		link.err = e
		return e
	} else if e = link.data.Append(r.Respon); e != nil {
		link.status = 3 // 设置切片状态为下载失败
		link.err = e
		return e
	} else {
		if int64(reqConfig.Timeout) < r.UsedTime.Milliseconds()+3000 {
			err = ActionErrFmp4DownloadCareTO.New(fmt.Sprintf("fmp4切片下载超时s(%d)或许应该大于%d", reqConfig.Timeout/1000, (r.UsedTime.Milliseconds()+4000)/1000))
		}
		link.status = 2 // 设置切片状态为下载完成
		return
	}
}

func (t *M4SStream) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(t.common, "", "    ")
}

func (t *M4SStream) getM4s() (p *m4s_link_item) {
	return t.m4s_pool.Get()
}

func (t *M4SStream) putM4s(ms ...*m4s_link_item) {
	t.m4s_pool.Put(ms...)
}

func (t *M4SStream) Common() *c.Common {
	return t.common
}

func NewM4SStream(c *c.Common) (*M4SStream, error) {
	var t = &M4SStream{
		common: c,
		log:    c.Log.Base(`直播流保存`),
	}

	//读取配置
	if path, ok := c.K_v.LoadV("直播流保存位置").(string); ok {
		if path, err := filepath.Abs(path); err == nil {
			if fs, err := os.Stat(path); err != nil {
				if errors.Is(err, os.ErrNotExist) {
					if err := os.Mkdir(path, os.ModePerm); err != nil {
						return nil, errors.New(`直播流保存位置错误` + err.Error())
					}
				} else {
					return nil, errors.New(`直播流保存位置错误` + err.Error())
				}
			} else if !fs.IsDir() {
				return nil, errors.New(`直播流保存位置不是目录`)
			}
			t.config.save_path = path
		} else {
			return nil, errors.New(`直播流保存位置错误` + err.Error())
		}
	} else {
		return nil, errors.New(`未配置直播流保存位置`)
	}
	if v, ok := c.K_v.LoadV(`直播流保存到文件`).(bool); ok {
		t.config.save_to_file = v
	}
	if v, ok := c.K_v.LoadV(`直播流清晰度`).(float64); ok {
		t.config.want_qn = int(v)
	}
	if v, ok := c.K_v.LoadV(`直播流类型`).(string); ok {
		t.config.want_type = v
	}
	return t, nil
}

// Deprecated: use NewM4SStream
func (t *M4SStream) LoadConfig(common *c.Common) (e error) {
	t.common = common
	t.log = common.Log.Base(`直播流保存`)

	//读取配置
	if path, ok := common.K_v.LoadV("直播流保存位置").(string); ok {
		if path, err := filepath.Abs(path); err == nil {
			if fs, err := os.Stat(path); err != nil {
				if errors.Is(err, os.ErrNotExist) {
					if err := os.Mkdir(path, os.ModePerm); err != nil {
						return errors.New(`直播流保存位置错误` + err.Error())
					}
				} else {
					return errors.New(`直播流保存位置错误` + err.Error())
				}
			} else if !fs.IsDir() {
				return errors.New(`直播流保存位置不是目录`)
			}
			t.config.save_path = path
		} else {
			return errors.New(`直播流保存位置错误` + err.Error())
		}
	} else {
		return errors.New(`未配置直播流保存位置`)
	}
	if v, ok := common.K_v.LoadV(`直播流保存到文件`).(bool); ok {
		t.config.save_to_file = v
	}
	if v, ok := common.K_v.LoadV(`直播流清晰度`).(float64); ok {
		t.config.want_qn = int(v)
	}
	if v, ok := common.K_v.LoadV(`直播流类型`).(string); ok {
		t.config.want_type = v
	}
	return
}

func (t *M4SStream) getFirstBuf() []byte {
	if t == nil {
		return []byte{}
	}
	return t.first_buf
}

func (t *M4SStream) fetchCheckStream() bool {
	_log := t.log.BaseAdd("获取流")
	// 获取流地址
	t.common.Live_want_qn = t.config.want_qn
	if F.Get(t.common).Get(`Live`); len(t.common.Live) == 0 {
		return false
	}

	// 直播流仅清晰度
	if v, ok := t.common.K_v.LoadV("直播流仅清晰度").(bool); ok && v {
		if _, ok := t.common.Qn[t.config.want_qn]; ok {
			if t.config.want_qn != t.common.Live_qn {
				_log.L(`W: `, `仅清晰度true,当前清晰度`, t.common.Qn[t.common.Live_qn])
				return false
			}
		}
	}

	// 直播流仅类型
	if v, ok := t.common.K_v.LoadV("直播流仅类型").(bool); ok && v {
		// 期望类型
		if vt, ok := t.common.K_v.LoadV(`直播流类型`).(string); ok {
			var (
				pass   bool
				cuType string
				cuCode string
			)

			if strings.Contains(t.common.Live[0].Codec, `hevc`) {
				cuCode = `hevc`
			} else if strings.Contains(t.common.Live[0].Codec, `avc`) {
				cuCode = `avc`
			} else {
				cuCode = `unknow`
			}

			if strings.Contains(t.common.Live[0].Url, `m3u8`) {
				cuType = `m3u8`
			} else if strings.Contains(t.common.Live[0].Url, `flv`) {
				cuType = `flv`
			} else {
				cuType = `unknow`
			}

			switch vt {
			case `fmp4H`:
				pass = cuType == `m3u8` && cuCode == `hevc`
			case `fmp4`:
				pass = cuType == `m3u8` && cuCode == `avc`
			case `flvH`:
				pass = cuType == `flv` && cuCode == `hevc`
			case `flv`:
				pass = cuType == `flv` && cuCode == `avc`
			}

			if !pass {
				_log.L(`W: `, `仅类型true,当前类型`, cuType, cuCode)
				return false
			}
		}
	}

	// 保存流类型
	if strings.Contains(t.common.Live[0].Url, `m3u8`) {
		t.stream_type = "mp4"
	} else if strings.Contains(t.common.Live[0].Url, `flv`) {
		t.stream_type = "flv"
	}

	// 检查是否可以获取
	CookieM := make(map[string]string)
	t.common.Cookie.Range(func(k, v interface{}) bool {
		CookieM[k.(string)] = v.(string)
		return true
	})

	var (
		noSer  []*regexp.Regexp
		noSerF = func(url string) (ban bool) {
			for i := 0; i < len(noSer) && !ban; i++ {
				ban = noSer[i].MatchString(url)
			}
			return
		}
	)
	if v, ok := t.common.K_v.LoadV("直播流不使用mcdn").(bool); ok && v {
		if reg, err := regexp.Compile(`\.mcdn\.`); err != nil {
			_log.L(`W: `, `停用流服务器`, `正则错误`, err)
		} else {
			noSer = append(noSer, reg)
		}
	}
	if v, ok := t.common.K_v.LoadV("直播流停用服务器").([]any); ok {
		for i := 0; i < len(v); i++ {
			if s, ok := v[i].(string); ok {
				if reg, err := regexp.Compile(s); err != nil {
					_log.L(`W: `, `停用流服务器`, `正则错误`, err)
				} else {
					noSer = append(noSer, reg)
				}
			}
		}
	}

	r := t.reqPool.Get()
	defer t.reqPool.Put(r)

	for _, v := range t.common.Live {
		if noSerF(v.Url) {
			v.Disable(time.Now().Add(time.Hour * 100))
			_log.L(`I: `, `停用流服务器`, F.ParseHost(v.Url))
			continue
		}

		if e := r.Reqf(reqf.Rval{
			Method: http.MethodGet,
			Url:    v.Url,
			Proxy:  t.common.Proxy,
			Header: map[string]string{
				`User-Agent`:      c.UA,
				`Accept`:          `*/*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`:          `https://live.bilibili.com`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Referer`:         "https://live.bilibili.com/",
				`Cookie`:          reqf.Map_2_Cookies_String(CookieM),
				`Connection`:      `close`,
			},
			Timeout:          5 * 1000,
			JustResponseCode: true,
		}); e != nil {
			_log.L(`W: `, F.ParseHost(v.Url), e)
			v.DisableAuto()
			continue
		}

		if r.Response == nil {
			_log.L(`W: `, `live响应错误`, F.ParseHost(v.Url))
			v.DisableAuto()
			continue
		} else if r.Response.StatusCode&200 != 200 {
			_log.L(`W: `, `live响应错误`, F.ParseHost(v.Url), r.Response.Status)
			v.DisableAuto()
			continue
		}

		// 显示使用流服务器
		_log.L(`I: `, `使用流服务器`, F.ParseHost(v.Url))
	}

	return t.common.ValidLive() != nil
}

func (t *M4SStream) fetchParseM3U8(lastM4s *m4s_link_item, fmp4ListUpdateTo float64) (m4s_links []*m4s_link_item, guessCount int, e error) {
	{
		n := t.common.ValidNum()
		if d, ok := t.common.K_v.LoadV("fmp4获取更多服务器").(bool); ok && d && n <= 1 && len(t.common.Live) <= 5 {
			t.log.L("I: ", "获取更多服务器...")
			if !t.fetchCheckStream() {
				e = errors.New("全部流服务器发生故障")
				return
			}
		} else if n == 0 {
			e = errors.New("全部流服务器发生故障")
			return
		}
	}

	// 开始请求
	r := t.reqPool.Get()
	defer t.reqPool.Put(r)

	// 请求解析m3u8内容
	for i := 0; i < len(t.common.Live); i++ {
		v := t.common.Live[i]

		// 跳过尚未启用的live地址
		if !v.Valid() {
			continue
		}

		// 设置请求参数
		rval := reqf.Rval{
			Url:     v.Url,
			Timeout: 2000,
			Proxy:   c.C.Proxy,
			Header: map[string]string{
				`Host`:            F.ParseHost(v.Url),
				`User-Agent`:      c.UA,
				`Accept`:          `*/*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`:          `https://live.bilibili.com`,
				`Connection`:      `keep-alive`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Referer`:         "https://live.bilibili.com/",
			},
		}
		if !t.stream_last_modified.IsZero() {
			rval.Header[`If-Modified-Since`] = t.stream_last_modified.Add(time.Second).Format(time.RFC1123)
		}

		if err := r.Reqf(rval); err != nil {
			v.DisableAuto()
			t.log.L("W: ", fmt.Sprintf("服务器 %s 发生故障 %s", F.ParseHost(v.Url), perrors.ErrorFormat(err, perrors.ErrActionInLineFunc)))
			if t.common.ValidLive() == nil {
				e = errors.New("全部流服务器发生故障")
				break
			}
			continue
		}

		if r.Response.StatusCode == http.StatusNotModified {
			t.log.L(`T: `, `hls未更改`)
			return
		}

		// 保存最后m3u8修改时间
		if last_mod, ok := r.Response.Header[`Last-Modified`]; ok && len(last_mod) > 0 {
			if lm, e := time.Parse(time.RFC1123, last_mod[0]); e == nil {
				t.stream_last_modified = lm
			}
		}

		// 解析m3u8
		// var tmp []*m4s_link_item
		var lastNo int
		if lastM4s != nil {
			lastNo, _ = lastM4s.getNo()
		}

		if rg, redirectUrl, err := replyFunc.ParseM3u8.Parse(r.Respon, lastNo); err != nil {
			if replyFunc.ParseM3u8.IsErrRedirect(err) {
				// 指向新连接
				t.log.L(`I: `, `指向新连接`, v.Host(), "=>", F.ParseHost(redirectUrl))
				v.SetUrl(redirectUrl)
				i -= 1
			} else {
				// 1min后重新启用
				t.log.L("W: ", fmt.Sprintf("服务器 %s 发生故障 %v", F.ParseHost(v.Url), err))
				v.DisableAuto()
				if t.common.ValidLive() == nil {
					e = errors.New("全部切片服务器发生故障")
					break
				}
			}
			continue
		} else {
			for m4sLinkI := range rg {
				//将切片添加到返回切片数组
				p := t.getM4s()
				p.SerUuid = v.Uuid
				p.Url = F.ResolveReferenceLast(v.Url, m4sLinkI.M4sLink()+"?trid="+F.ParseQuery(v.Url, "trid="))
				p.Base = m4sLinkI.M4sLink()
				p.isHeader = m4sLinkI.IsHeader()
				p.createdTime = time.Now()
				m4s_links = append(m4s_links, p)
			}
		}

		if len(m4s_links) == 0 {
			if lastM4s != nil &&
				!lastM4s.createdTime.IsZero() &&
				time.Since(lastM4s.createdTime).Seconds() > fmp4ListUpdateTo {
				// 1min后重新启用
				v.DisableAuto()
				t.log.L("W: ", fmt.Sprintf("服务器 %s 发生故障 %.2f 秒未产出切片", F.ParseHost(v.Url), time.Since(lastM4s.createdTime).Seconds()))
				if t.common.ValidLive() == nil {
					e = errors.New("全部切片服务器发生故障")
					break
				}
				continue
			}
			// fmt.Println("->", "empty", lastNo)
			return
		}

		// 检查是否服务器发生故障,产出切片错误
		if lastM4s != nil {
			timed := m4s_links[len(m4s_links)-1].createdTime.Sub(lastM4s.createdTime).Seconds()
			nos, _ := m4s_links[len(m4s_links)-1].getNo()
			noe := lastNo
			if (timed > 5 && nos-noe == 0) || (nos-noe > 50) {
				// 1min后重新启用
				v.DisableAuto()
				t.log.L("W: ", fmt.Sprintf("服务器 %s 发生故障 %d 秒产出了 %d 切片", F.ParseHost(v.Url), int(timed), nos-noe))
				if t.common.ValidLive() == nil {
					e = errors.New("全部切片服务器发生故障")
					break
				}
				continue
			}
		}

		// 首次下载
		if lastM4s == nil {
			return
		}

		// 刚刚好承接之前的结尾
		if linksFirstNo, _ := m4s_links[0].getNo(); linksFirstNo == lastNo+1 {
			return
		} else {
			// 来到此处说明出现了丢失 尝试补充
			guessCount = linksFirstNo - lastNo - 1
			for guess_no := linksFirstNo - 1; guess_no > lastNo; guess_no-- {
				//将切片添加到返回切片数组前
				p := t.getM4s()
				p.SerUuid = v.Uuid
				p.Base = strconv.Itoa(guess_no) + `.m4s`
				p.isHeader = false
				p.Url = F.ResolveReferenceLast(v.Url, p.Base)
				p.createdTime = time.Now()
				slice.AddFront(&m4s_links, p)
			}
		}

		// 请求解析成功，退出获取循环
		return
	}

	if e != nil {
		e = errors.New(e.Error() + " 未能找到可用流服务器")
	}
	return
}

// 移除历史流
func (t *M4SStream) removeStream() (e error) {
	if d, ok := t.common.K_v.LoadV("直播流保存天数").(float64); ok && d >= 1 {
		if v, ok := t.common.K_v.LoadV(`直播流保存位置`).(string); ok && v != "" {
			type dirEntryDirs []fs.DirEntry
			var list dirEntryDirs
			f, err := http.Dir(v).Open("/")
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err = f.Stat(); err != nil {
				return err
			}
			if d, ok := f.(fs.ReadDirFile); ok {
				list, err = d.ReadDir(-1)
			}

			if err != nil {
				return err
			}

			var (
				oldest   float64
				oldIndex []int
			)
			for i, n := 0, len(list); i < n; i++ {
				if list[i].IsDir() && len(list[i].Name()) > 20 {
					if file.New(v+"/"+list[i].Name()+"/.keep", 0, true).IsExist() {
						continue
					}
					if tt, err := time.Parse("2006_01_02-15_04_05", list[i].Name()[:19]); err == nil {
						if ts := time.Since(tt).Seconds(); ts > d*24*60*60 && ts > oldest {
							oldest = ts
							oldIndex = append(oldIndex, i)
						}
					}
				}
			}

			for n, i := 2, len(oldIndex)-1; n > 0 && i >= 0; n, i = n-1, i-1 {
				t.log.L(`I: `, "移除历史流", v+"/"+list[oldIndex[i]].Name())
				os.RemoveAll(v + "/" + list[oldIndex[i]].Name())
			}
		}
	}
	return nil
}

// 设置保存路径
func (t *M4SStream) getSavepath() {
	w := md5.New()
	_, _ = io.WriteString(w, t.common.Title)

	t.Current_save_path = fmt.Sprintf("%s/%s-%d-%d-%x-%s/",
		t.config.save_path,
		time.Now().Format("2006_01_02-15_04_05"),
		t.common.Roomid,
		t.common.Live_qn,
		w.Sum(nil)[:3],
		pstring.Rand(2, 3))

	// 显示保存位置
	if rel, err := filepath.Rel(t.config.save_path, t.Current_save_path); err == nil {
		t.log.L(`I: `, "保存到", rel+`/0.`+t.stream_type)
		f := file.New(t.Current_save_path+"tmp.create", 0, true)
		f.Create()
		_ = f.Delete()
	} else {
		t.log.L(`W: `, err)
	}
}

var ErrDecode = perrors.Action("ErrDecode")

func (t *M4SStream) saveStream() (e error) {
	// 清除初始值
	t.first_buf = nil
	t.frameCount = 0

	if s, ok := t.common.K_v.LoadV("直播Web服务路径").(string); ok && s != "" {
		t.log.L(`I: `, "Web服务地址", t.common.Stream_url.String()+s)
	}

	// 录制回调
	t.msg.Push_tag(`startRec`, t)
	if t.Callback_startRec != nil {
		if err := t.Callback_startRec(t); err != nil {
			t.log.L(`W: `, `开始录制回调错误`, err.Error())
			return err
		}
	}

	// 保存到文件
	if t.config.save_to_file {
		var startCount uint = defaultStartCount
		if s, ok := t.common.K_v.LoadV("直播流接收n帧才保存").(float64); ok && s > 0 && uint(s) > startCount {
			startCount = uint(s)
		}
		// 确保能接收到第n个帧才开始录制
		var cancelkeyFrame = t.msg.Pull_tag_only(`keyFrame`, func(ms *M4SStream) (disable bool) {
			if startCount <= t.frameCount {
				ms.msg.Push_tag(`cut`, ms)
				return true
			}
			t.log.L(`T: `, fmt.Sprintf("%d帧后开始录制", startCount-t.frameCount))
			return false
		})
		defer cancelkeyFrame()
	}

	// 获取流
	switch t.stream_type {
	case `mp4`:
		e = t.saveStreamM4s()
	case `flv`:
		e = t.saveStreamFlv()
	default:
		e = errors.New("undefind stream type")
		t.log.L(`E: `, e)
	}

	// 退出当前方法时，结束录制
	t.msg.Push_tag(`stopRec`, t)
	return
}

func (t *M4SStream) saveStreamFlv() (e error) {
	// 开始获取
	r := t.reqPool.Get()
	defer t.reqPool.Put(r)

	CookieM := make(map[string]string)
	t.common.Cookie.Range(func(k, v interface{}) bool {
		CookieM[k.(string)] = v.(string)
		return true
	})

	var surl *url.URL

	// 找到可用流服务器
	for {
		// 移除失效源
		t.removeSer()

		v := t.common.ValidLive()
		if v == nil {
			return errors.New("未能找到可用流服务器")
		}

		//reset e
		e = nil

		{
			var err error
			surl, err = url.Parse(v.Url)
			if err != nil {
				t.log.L(`E: `, err)
				e = err
				v.DisableAuto()
				continue
			}
		}

		//结束退出
		if pctx.Done(t.Status) {
			return
		}

		//检查
		if e := r.Reqf(reqf.Rval{
			Url:              surl.String(),
			NoResponse:       true,
			JustResponseCode: true,
			Proxy:            t.common.Proxy,
			Timeout:          5000,
			Header: map[string]string{
				`Host`:            surl.Host,
				`User-Agent`:      c.UA,
				`Accept`:          `*/*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Origin`:          `https://live.bilibili.com`,
				`Connection`:      `keep-alive`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Referer`:         "https://live.bilibili.com/",
				`Cookie`:          reqf.Map_2_Cookies_String(CookieM),
			},
		}); e != nil && reqf.IsTimeout(e) {
			v.DisableAuto()
			continue
		}

		// flv获取
		cancelC, cancel := context.WithCancel(t.Status)
		errCtx := pctx.Value[error]{}
		cancelC = errCtx.LinkCtx(cancelC)
		{
			pipe := pio.NewPipe()
			var (
				leastReadUnix atomic.Int64
				readTO        int64 = 10
			)
			leastReadUnix.Store(time.Now().Unix())
			if v, ok := t.common.K_v.LoadV(`flv断流超时s`).(float64); ok && int64(v) > readTO {
				readTO = int64(v)
			}

			// read timeout
			go func() {
				defer cancel()

				timer := time.NewTicker(time.Duration(readTO * int64(time.Second)))
				defer timer.Stop()

				for {
					select {
					case <-cancelC.Done():
						return
					case curT := <-timer.C:
						if curT.Unix()-leastReadUnix.Load() > readTO {
							t.log.L(`W: `, fmt.Sprintf("%vs未接收到有效数据", readTO))
							pctx.PutVal(cancelC, &errCtx, fmt.Errorf("%vs未接收到有效数据", readTO))
							// 时间段内未接收到任何数据
							return
						}
						if v, ok := c.C.K_v.LoadV(`直播流清晰度`).(float64); ok {
							if t.config.want_qn != int(v) {
								t.log.L(`I: `, "直播流清晰度改变:", t.common.Qn[t.config.want_qn], "=>", t.common.Qn[int(v)])
								t.config.want_qn = int(v)
								return
							}
						}
					}
				}
			}()

			// read
			go func() {
				defer cancel()

				var (
					buff       = slice.New[byte]()
					keyframe   = slice.New[byte]()
					buf        = make([]byte, humanize.KByte)
					flvDecoder = NewFlvDecoder()
					bufSize    = humanize.KByte * 1100
				)

				if v, ok := c.C.K_v.LoadV(`flv音视频时间戳容差ms`).(float64); ok && v > 100 {
					flvDecoder.Diff = v
				}

				for {
					if buff.Size() < bufSize {
						if n, e := pipe.Read(buf); e != nil {
							pctx.PutVal(cancelC, &errCtx, e)
							break
						} else if e = buff.Append(buf[:n]); e != nil {
							pctx.PutVal(cancelC, &errCtx, e)
							break
						}
						continue
					}

					if !buff.IsEmpty() {
						// front_buf
						buf, unlock := buff.GetPureBufRLock()
						frontBuf, dropOffset, e := flvDecoder.Parse(buf, keyframe)
						unlock()
						if e != nil {
							t.log.L(`E: `, e)
							pctx.PutVal(cancelC, &errCtx, errors.New("[decoder]"+e.Error()))
							break
						}

						if len(frontBuf) != 0 {
							t.first_buf = frontBuf
							t.msg.Push_tag(`load`, t)
						}

						if keyframe.Size() != 0 {
							if l := leastReadUnix.Load(); l > 0 && time.Now().Unix()-l > readTO-5 {
								t.log.L(`W: `, fmt.Sprintf("flv断流超时s(%d)或许应该大于%d", readTO, (time.Now().Unix()-l+5)))
							}
							// 存在有效数据
							leastReadUnix.Store(time.Now().Unix())

							buf, unlock := keyframe.GetPureBufRLock()
							t.bootBufPush(buf)
							t.Stream_msg.PushLock_tag(`data`, buf)
							unlock()

							keyframe.Reset()
							t.frameCount += 1
							t.msg.Push_tag(`keyFrame`, t)
						} else {
							bufSize += humanize.KByte * 50
							if bufSize > humanize.MByte*10 {
								t.log.L(`E: `, `缓冲池过大`)
								pctx.PutVal(cancelC, &errCtx, errors.New("缓冲池过大"))
								break
							}
						}

						if dropOffset > 0 {
							_ = buff.RemoveFront(dropOffset)
						}
					}
				}
				buf = nil
				buff.Reset()
			}()

			t.log.L(`I: `, `flv下载开始`, F.ParseHost(surl.String()))

			_ = r.Reqf(reqf.Rval{
				Ctx:         cancelC,
				Url:         surl.String(),
				SaveToPipe:  pipe,
				NoResponse:  true,
				Async:       true,
				Proxy:       t.common.Proxy,
				WriteLoopTO: int(readTO)*1000*2 + 1,
				Header: map[string]string{
					`Host`:            surl.Host,
					`User-Agent`:      c.UA,
					`Accept`:          `*/*`,
					`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
					`Origin`:          `https://live.bilibili.com`,
					`Connection`:      `keep-alive`,
					`Pragma`:          `no-cache`,
					`Cache-Control`:   `no-cache`,
					`Referer`:         "https://live.bilibili.com/",
					`Cookie`:          reqf.Map_2_Cookies_String(CookieM),
				},
			})
			if err := r.Wait(); err != nil && !errors.Is(err, io.EOF) {
				if reqf.IsCancel(err) {
					t.log.L(`I: `, `flv下载停止`, F.ParseHost(surl.String()))
					return
				} else if !reqf.IsTimeout(err) {
					e = err
					t.log.L(`E: `, `flv下载失败:`, F.ParseHost(surl.String()), err)
				} else {
					t.log.L(`E: `, `flv下载超时`, F.ParseHost(surl.String()))
				}
			} else if err := errCtx.Get(); err != nil && strings.HasPrefix(err.Error(), "[decoder]") {
				e = err
			}
		}

		cancel()

		if v1, ok := t.common.K_v.LoadV(`flv断流续接`).(bool); ok && !v1 {
			break
		}
		v.DisableAuto()
	}
	return
}

// 移除失效源
func (t *M4SStream) removeSer() {
	slice.Del(&t.common.Live, func(v **c.LiveQn) (del bool) {
		isDel := time.Now().Add(time.Minute * 2).Before((*v).ReUpTime)
		if isDel {
			t.log.L(`I: `, `移除流服务器`, (*v).Host())
		}
		return isDel
	})
}

func (t *M4SStream) saveStreamM4s() (e error) {

	var (
		// 同时下载数限制
		downloadLimit    = funcCtrl.NewBlockFuncN(3)
		buf              = slice.New[byte]()
		fmp4Decoder      = NewFmp4Decoder()
		keyframe         = slice.New[byte]()
		lastM4s          *m4s_link_item
		to               = 5
		fmp4ListUpdateTo = 5.0
		planSecPeriod    = 5.0
		lastNewT         = time.Now()
		skipErrFrame     = false
	)

	if v, ok := t.common.K_v.LoadV(`debug模式`).(bool); ok && v {
		fmp4Decoder.Debug = true
		cancle := make(chan struct{})
		defer close(cancle)
		go func() {
			for {
				select {
				case <-cancle:
					return
				case <-time.After(time.Minute):
				}
				reqState := t.m4s_pool.State()
				t.log.L(`T: `, fmt.Sprintf("m4sPoolState pooled/no(%d/%d), inuse/no(%d/%d), sum(%d), qts(%.2f)",
					reqState.Pooled, reqState.Nopooled, reqState.Inuse, reqState.Nouse, reqState.Sum, reqState.GetPerSec))
			}
		}()
	}
	if v, ok := t.common.K_v.LoadV(`fmp4音视频时间戳容差s`).(float64); ok && v > 0.1 {
		fmp4Decoder.AVTDiff = v
	}
	if v, ok := t.common.K_v.LoadV(`fmp4切片下载超时s`).(float64); ok && to < int(v) {
		to = int(v)
	}
	if v, ok := t.common.K_v.LoadV(`fmp4列表更新超时s`).(float64); ok && fmp4ListUpdateTo < v {
		fmp4ListUpdateTo = v
	}
	if v, ok := t.common.K_v.LoadV(`fmp4跳过解码出错的帧`).(bool); ok {
		skipErrFrame = v
	}

	// 下载循环
	for download_seq := []*m4s_link_item{}; ; {
		// 移除失效源
		t.removeSer()

		// 获取解析m3u8
		{
			// 防止过快的下载
			if needWaitSec := planSecPeriod - time.Since(lastNewT).Seconds(); needWaitSec > 0 {
				time.Sleep(time.Duration(needWaitSec) * time.Second)
			}

			var m4s_links, guessCount, err = t.fetchParseM3U8(lastM4s, fmp4ListUpdateTo)
			if err != nil {
				t.log.L(`E: `, `获取解析m3u8发生错误`, err)
				// if len(download_seq) != 0 {
				// 	continue
				// }
				if !reqf.IsTimeout(err) {
					e = err
					break
				}
			}

			countInLastPeriod := len(m4s_links)
			secInLastPeriod := time.Since(lastNewT).Seconds()
			if countInLastPeriod != 0 {
				lastNewT = time.Now()
			} else if secInLastPeriod > fmp4ListUpdateTo {
				// fmp4ListUpdateTo秒未产出切片
				e = fmt.Errorf("%.2f 秒未产出切片", secInLastPeriod)
				t.log.L("E: ", "获取解析m3u8发生错误", e)
				break
			}

			// 评估一轮正常下载的周期时长
			if countInLastPeriod-guessCount > 0 && countInLastPeriod > 0 && secInLastPeriod > 0 {
				planSecPeriod = float64(countInLastPeriod-guessCount) / float64(countInLastPeriod) * secInLastPeriod
				if guessCount == 0 {
					planSecPeriod += 0.2
				}
				if guessCount >= 3 {
					t.log.L(`I: `, `发现`, guessCount, `个切片遗漏，优先下载`)
				}
			}
			if planSecPeriod < 3 {
				planSecPeriod = 3
			}

			// 设置最后的切片
			for i := len(m4s_links) - 1; i >= 0; i-- {
				// fmt.Println("set last m4s", m4s_links[i].Base)
				if !m4s_links[i].isInit() && len(m4s_links[i].Base) > 0 {
					if lastM4s == nil {
						lastM4s = &m4s_link_item{}
					}
					m4s_links[i].copyTo(lastM4s)
					break
				}
			}

			// 添加新切片到下载队列
			download_seq = append(download_seq, m4s_links...)
		}

		// 下载切片
		for {
			var downErr atomic.Bool
			dCount := 0
			for i := 0; i < len(download_seq); i++ {
				// 已下载但还未移除的切片
				if download_seq[i].status == 2 {
					continue
				}

				// 每次最多只下载10个切片
				if dCount >= 10 {
					t.log.L(`T: `, `延迟切片下载 数量(`, len(download_seq)-i, `)`)
					break
				}
				dCount += 1

				// 故障转移
				if download_seq[i].status == 3 {
					if linkUrl, e := url.Parse(download_seq[i].Url); e == nil {
						// 将此切片服务器设置停用
						// hadDisable := t.common.DisableLiveAuto(oldHost)
						hadDisable := t.common.DisableLiveAutoByUuid(download_seq[i].SerUuid)
						// 从其他服务器获取此切片
						if vl := t.common.ValidLive(); vl == nil {
							return errors.New(`全部流服务器故障`)
						} else {
							download_seq[i].replaceSer(vl)
							if !hadDisable {
								t.log.L(`W: `, `切片下载失败，故障转移`, linkUrl.Host, ` -> `, vl.Host())
							}
						}
						// download_seq[i].Url = linkUrl.String()
					} else {
						return errors.New(`切片url错误`)
					}
				}

				done := downloadLimit.Block()
				go func(link *m4s_link_item) {
					defer done()

					e := link.download(t.reqPool, reqf.Rval{
						Timeout:     to * 1000,
						WriteLoopTO: (to + 2) * 1000,
						Proxy:       t.common.Proxy,
						Header: map[string]string{
							`Connection`: `close`,
						},
					})
					if ActionErrFmp4DownloadCareTO.Catch(e) {
						t.log.L(`W: `, e.Error())
					} else if e != nil {
						downErr.Store(true)
						if reqf.IsTimeout(e) {
							t.log.L(`W: `, fmt.Sprintf("fmp4切片下载超时s或许应该大于%d", to))
						}
						t.log.L(`W: `, `切片下载失败`, link.Base, perrors.ErrorFormat(e, perrors.ErrActionInLineFunc))
					}
				}(download_seq[i])
				// 间隔100ms发起
				time.Sleep(time.Millisecond * 100)
			}

			// 等待队列下载完成
			downloadLimit.BlockAll()()

			if !downErr.Load() {
				break
			}
		}

		// 传递已下载切片
		for k := 0; k < len(download_seq) && download_seq[k].status == 2; k++ {

			cu := download_seq[k]

			if cu.isInit() {
				{
					buf, unlock := cu.data.GetPureBufRLock()
					front_buf, e := fmp4Decoder.Init_fmp4(buf)
					unlock()
					if e != nil {
						t.log.L(`E: `, e, `重试!`)
						cu.status = 3
						break
					} else {
						for _, trak := range fmp4Decoder.traks {
							// fmt.Println(`T: `, "找到trak:", string(trak.handlerType), trak.trackID, trak.timescale)
							t.log.L(`T: `, "找到trak:", string(trak.handlerType), trak.trackID, trak.timescale)
						}
						t.first_buf = make([]byte, len(front_buf))
						copy(t.first_buf, front_buf)
						t.msg.Push_tag(`load`, t)
					}
				}
				t.putM4s(cu)
				download_seq = append(download_seq[:k], download_seq[k+1:]...)
				k -= 1
				continue
			} else if t.first_buf == nil {
				t.putM4s(cu)
				download_seq = append(download_seq[:k], download_seq[k+1:]...)
				k -= 1
				continue
			}

			if e := cu.data.AppendTo(buf); e != nil {
				t.log.L(`E: `, e)
			}
			t.putM4s(cu)
			download_seq = append(download_seq[:k], download_seq[k+1:]...)
			k -= 1

			buff, unlock := buf.GetPureBufRLock()
			last_available_offset, err := fmp4Decoder.Search_stream_fmp4(buff, keyframe)
			unlock()

			if err != nil && !errors.Is(err, io.EOF) {
				if ErrDecode.Catch(err) && skipErrFrame {
					t.log.L(`W: `, err)
					// 将此切片服务器设置停用
					// if u, e := url.Parse(cu.Url); e == nil {
					t.common.DisableLiveAutoByUuid(cu.SerUuid)
					// t.common.DisableLiveAuto(u.Host)
					// }
				} else {
					e = err
					download_seq = download_seq[:0]
					_ = pctx.CallCancel(t.Status)
					break
				}
			}

			// no, _ := cu.getNo()
			// fmt.Println(no, "fmp4KeyFrames", keyframe.Size(), last_available_offset, err)

			// 传递关键帧
			if !keyframe.IsEmpty() {
				keyframeBuf, unlock := keyframe.GetPureBufRLock()
				t.bootBufPush(keyframeBuf)
				t.Stream_msg.PushLock_tag(`data`, keyframeBuf)
				unlock()
				keyframe.Reset()
				t.frameCount += 1
				t.msg.Push_tag(`keyFrame`, t)
			}

			_ = buf.RemoveFront(last_available_offset)
		}

		// 停止录制
		if pctx.Done(t.Status) {
			if len(download_seq) != 0 {
				if time.Now().Unix() > t.stream_last_modified.Unix()+300 {
					e = errors.New("切片下载超时")
					t.log.L(`E: `, e)
				} else {
					t.log.L(`I: `, `下载最后切片:`, len(download_seq))
					continue
				}
			}
			break
		}

		if v, ok := c.C.K_v.LoadV(`直播流清晰度`).(float64); ok {
			if t.config.want_qn != int(v) {
				t.log.L(`I: `, "直播流清晰度改变:", t.common.Qn[t.config.want_qn], "=>", t.common.Qn[int(v)])
				t.config.want_qn = int(v)
				return
			}
		}
	}

	return
}

func (t *M4SStream) GetStreamType() string {
	return t.stream_type
}

func (t *M4SStream) GetSavePath() string {
	return t.Current_save_path
}

func (t *M4SStream) Start() bool {
	// 清晰度-1 or 路径存在问题 不保存
	if t.config.want_qn == -1 || t.config.save_path == "" {
		return false
	}

	// 状态检测与设置
	if !pctx.Done(t.Status) {
		t.log.L(`T: `, `已存在实例`)
		return false
	}

	// 是否在直播
	F.Get(t.common).Get(`Liveing`)
	if !t.common.Liveing {
		t.log.L(`I: `, `未直播`)
		return false
	}

	// 实例回调
	t.msg = msgq.NewType[*M4SStream](time.Second * 5)
	t.msg.Push_tag(`start`, t)
	if t.Callback_start != nil {
		if e := t.Callback_start(t); e != nil {
			t.log.L(`W: `, `开始回调错误`, e.Error())
			return false
		}
	}

	go func() {
		t.log.L(`I: `, `初始化录制(`+strconv.Itoa(t.common.Roomid)+`)`)

		defer t.log.L(`I: `, `结束录制(`+strconv.Itoa(t.common.Roomid)+`)`)
		defer func() {
			// use anonymous func avoid data race and unexpect sign wait
			t.exitSign.Done()
		}()

		// 初始化请求池
		t.reqPool = t.common.ReqPool

		// 初始化池
		t.m4s_pool = pool.New(
			pool.PoolFunc[m4s_link_item]{
				New: func() *m4s_link_item {
					return &m4s_link_item{
						data: slice.New[byte](),
					}
				},
				InUse: func(t *m4s_link_item) bool {
					return t.createdTime.After(t.pooledTime) || time.Now().Before(t.pooledTime.Add(time.Second*10))
				},
				Reuse: func(t *m4s_link_item) *m4s_link_item {
					return t.reset()
				},
				Pool: func(t *m4s_link_item) *m4s_link_item {
					t.pooledTime = time.Now()
					return t
				},
			},
			50,
		)

		// 初始化切片消息
		t.Stream_msg = msgq.NewType[[]byte]()

		// 设置事件
		// 当录制停止时，取消全部录制
		t.Status = pctx.CarryCancel(context.WithCancel(context.Background()))
		mainCtx, done := pctx.WithWait(t.Status, 0, time.Minute)
		defer func() {
			switch done() {
			case pctx.ErrWaitTo:
				t.log.L(`E: `, `结束超时`)
			case pctx.ErrNothingWait:
				fallthrough
			default:
			}
			_ = pctx.CallCancel(t.Status)
		}()

		if t.Callback_stopRec != nil {
			cancel := t.msg.Pull_tag_only(`stopRec`, func(ms *M4SStream) (disable bool) {
				ms.Callback_stopRec(ms)
				return false
			})
			defer cancel()
		}
		cancel := t.msg.Pull_tag_only("stop", func(ms *M4SStream) (disable bool) {
			if ms.Callback_stop != nil {
				ms.Callback_stop(ms)
			}
			_ = pctx.CallCancel(t.Status)
			t.msg.ClearAll()
			return true
		})
		defer cancel()

		if t.config.save_to_file {
			var fc funcCtrl.FlashFunc
			cancel := t.msg.Pull_tag_async(map[string]func(*M4SStream) (disable bool){
				`cut`: func(ms *M4SStream) (disable bool) {
					// 有时尚未初始化接收到新的cut信号，导致保存失败。可能在开播信号重复发出出现
					if ms.frameCount < defaultStartCount {
						ml := ms.log.Base_add(`分段`)
						ml.L(`I: `, "尚未接收到帧、跳过")
						return false
					}

					// 当cut时，取消上次录制
					ctx1, done := pctx.WithWait(mainCtx, 3, time.Second*30)
					fc.FlashWithCallback(func() { _ = done() })

					// 分段时长min
					if l, ok := ms.common.K_v.LoadV("分段时长min").(float64); ok && l > 0 {
						cutT := time.Duration(int64(time.Minute) * int64(l))
						ml := ms.log.Base_add(`分段`)
						ml.L(`I: `, fmt.Sprintf("分段启动 %v", cutT))
						defer time.AfterFunc(cutT, func() {
							ml.L(`I: `, ms.common.Roomid, "ok")
							ms.msg.Push_tag(`cut`, ms)
						}).Stop()
					}

					// 当stopRec时，取消录制
					cancelMsg := ms.msg.Pull_tag_only(`stopRec`, func(_ *M4SStream) (disable bool) {
						cancel()
						return true
					})
					defer cancelMsg()

					ms.getSavepath()

					l := ms.log.Base_add(`文件保存`)
					startf := func(_ *M4SStream) error {
						l.L(`T: `, `开始`)
						//弹幕分值统计
						replyFunc.DanmuCountPerMin.Rec(ctx1, ms.common.Roomid, ms.GetSavePath())(ms.common.K_v.LoadV("弹幕分值"))
						return nil
					}
					stopf := func(_ *M4SStream) error {
						l.L(`T: `, `结束`)
						return nil
					}

					// 移除历史流
					if err := ms.removeStream(); err != nil {
						l.L(`W: `, err)
					}

					// savestate
					if e, _ := videoInfo.Save.Run(ctx1, ms); e != nil {
						l.L(`E: `, e)
					}

					//保存弹幕
					go StartRecDanmu(ctx1, ms.GetSavePath())

					//指定房间录制回调
					if v, ok := ms.common.K_v.LoadV("指定房间录制回调").([]any); ok && len(v) > 0 {
						l := l.Base(`录制回调`)
						for i := 0; i < len(v); i++ {
							if vm, ok := v[i].(map[string]any); ok {
								if roomid, ok := vm["roomid"].(float64); ok && int(roomid) == ms.common.Roomid {
									var (
										durationS, _ = vm["durationS"].(float64)
										start, _     = vm["start"].([]any)
									)
									if len(start) >= 2 && durationS >= 0 {
										go func() {
											ctx2, done2 := pctx.WaitCtx(ctx1)
											defer done2()
											select {
											case <-ctx2.Done():
											case <-time.After(time.Second * time.Duration(durationS)):
												var cmds []string
												for i := 0; i < len(start); i++ {
													if cmd, ok := start[i].(string); ok && cmd != "" {
														cmds = append(cmds, strings.ReplaceAll(cmd, "{type}", ms.GetStreamType()))
													}
												}

												cmd := exec.Command(cmds[0], cmds[1:]...)
												cmd.Dir = ms.GetSavePath()
												l.L(`I: `, "启动", cmd.Args)
												if e := cmd.Run(); e != nil {
													l.L(`E: `, e)
												}
												l.L(`I: `, "结束")
											}
										}()
									}
								}
							}
						}
					}

					path := ms.GetSavePath() + `0.` + ms.GetStreamType()
					startT := time.Now()
					if e := ms.PusherToFile(ctx1, path, startf, stopf); e != nil {
						l.L(`E: `, e)
					}
					duration := time.Since(startT)

					// wait all goroutine exit
					if e := done(); e != nil && !errors.Is(e, pctx.ErrDoneCalled) {
						l.L(`E: `, e)
					}

					//PusherToFile fin genFastSeed
					if disableFastSeed, ok := ms.common.K_v.LoadV("禁用快速索引生成").(bool); !ok || !disableFastSeed {
						type deal interface {
							GenFastSeed(reader io.Reader, save func(seedTo time.Duration, cuIndex int64) error) (err error)
						}
						var dealer deal

						switch ms.GetStreamType() {
						case `mp4`:
							fmp4Decoder := NewFmp4Decoder()
							if v, ok := ms.common.K_v.LoadV(`fmp4音视频时间戳容差s`).(float64); ok && v > 0.1 {
								fmp4Decoder.AVTDiff = v
							}
							dealer = fmp4Decoder
						case `flv`:
							flvDecoder := NewFlvDecoder()
							if v, ok := ms.common.K_v.LoadV(`flv音视频时间戳容差ms`).(float64); ok && v > 100 {
								flvDecoder.Diff = v
							}
							dealer = flvDecoder
						default:
						}

						if dealer != nil {
							f := file.New(path, 0, false)
							if sf, e := replyFunc.VideoFastSeed.InitSav(path + ".fastSeed"); e != nil {
								l.L(`E: `, e)
							} else if e := dealer.GenFastSeed(f, sf); e != nil && !errors.Is(e, io.EOF) {
								l.L(`E: `, e)
							}
							f.Close()
						}
					}

					// 结束，不发送空值停止直播回放
					// t.Stream_msg.PushLock_tag(`data`, []byte{})

					//指定房间录制回调
					if v, ok := ms.common.K_v.LoadV("指定房间录制回调").([]any); ok && len(v) > 0 {
						l := l.Base(`录制回调`)
						for i := 0; i < len(v); i++ {
							if vm, ok := v[i].(map[string]any); ok {
								if roomid, ok := vm["roomid"].(float64); ok && int(roomid) == ms.common.Roomid {
									var (
										durationS, _ = vm["durationS"].(float64)
										after, _     = vm["after"].([]any)
									)
									if len(after) >= 2 && durationS >= 0 && duration.Seconds() > durationS {
										var cmds []string
										for i := 0; i < len(after); i++ {
											if cmd, ok := after[i].(string); ok && cmd != "" {
												cmds = append(cmds, strings.ReplaceAll(cmd, "{type}", ms.GetStreamType()))
											}
										}

										cmd := exec.Command(cmds[0], cmds[1:]...)
										cmd.Dir = ms.GetSavePath()
										l.L(`I: `, "启动", cmd.Args)
										if e := cmd.Run(); e != nil {
											l.L(`E: `, e)
										}
										l.L(`I: `, "结束")
									}
								}
							}
						}
					}
					return false
				},
			})
			defer cancel()
		}

		// 主循环
		for !pctx.Done(t.Status) {
			// 是否在直播
			F.Get(t.common).Get(`Liveing`)
			if !t.common.Liveing {
				t.log.L(`I: `, `未直播`)
				break
			}

			// 新循环，取消所有流
			t.common.Live = t.common.Live[:0]

			// 获取 and 检查流地址状态
			if !t.fetchCheckStream() {
				time.Sleep(time.Second * 5)
				continue
			}

			// 保存流
			err := t.saveStream()
			if err != nil {
				t.log.L(`E: `, "saveStream:", err)
			}
		}

		// 退出
		t.msg.Push_tag(`stop`, t)
	}()
	return true
}

func (t *M4SStream) Stop() {
	if pctx.Done(t.Status) {
		t.log.L(`I: `, `正在等待下载完成...`)
		t.exitSign.Wait()
		return
	}
	t.exitSign = signal.Init()
	_ = pctx.CallCancel(t.Status)
	t.log.L(`I: `, `正在等待下载完成...`)
	t.exitSign.Wait()
	t.log.L(`I: `, `结束`)
}

func (t *M4SStream) Cut() {
	t.msg.Push_tag(`cut`, t)
}

// 保存到文件
func (t *M4SStream) PusherToFile(contextC context.Context, filepath string, startFunc func(*M4SStream) error, stopFunc func(*M4SStream) error) error {
	f := file.New(filepath, 0, false)
	defer f.Close()
	_ = f.Delete()

	if e := startFunc(t); e != nil {
		return e
	}

	ctx1, done := pctx.WaitCtx(contextC)
	defer done()

	to := 2.0
	if tmp, ok := t.common.K_v.LoadV("直播流保存写入超时").(float64); ok && tmp > 2 {
		to = tmp
	}

	_, _ = f.Write(t.getFirstBuf(), true)
	cancelRec := t.Stream_msg.Pull_tag(map[string]func([]byte) bool{
		`data`: func(b []byte) bool {
			defer pu.Callback(func(startT time.Time, args ...any) {
				if dru := time.Since(startT).Seconds(); dru > to {
					t.log.L("W: ", "磁盘写入超时", dru)
					done()
				}
			})()

			select {
			case <-ctx1.Done():
				return true
			default:
			}
			if len(b) == 0 {
				return true
			}
			if n, err := f.Write(b, true); err != nil || n == 0 {
				done()
			}
			return false
		},
		`close`: func(_ []byte) bool {
			return true
		},
	})
	<-ctx1.Done()
	cancelRec()

	if e := stopFunc(t); e != nil {
		return e
	}

	return nil
}

// 流服务推送方法
//
// 在客户端存在某种代理时，将有可能无法监测到客户端关闭，这有可能导致goroutine泄漏
func (t *M4SStream) PusherToHttp(plog *log.Log_interface, conn net.Conn, w http.ResponseWriter, r *http.Request, startFunc func(*M4SStream) error, stopFunc func(*M4SStream) error) error {
	switch t.stream_type {
	case `m3u8`:
		fallthrough
	case `mp4`:
		w.Header().Set("Content-Type", "video/mp4")
	case `flv`:
		w.Header().Set("Content-Type", "video/x-flv")
	default:
		w.WriteHeader(http.StatusNotFound)
		return errors.New("pusher no support stream_type")
	}

	if e := startFunc(t); e != nil {
		return e
	}

	w = pweb.WithFlush(w)

	ctx, cancel := context.WithCancelCause(r.Context())
	defer cancel(nil)
	//写入头
	{
		retry := 5
		for retry > 0 {
			select {
			case <-ctx.Done():
				break
			default:
			}

			if len(t.getFirstBuf()) != 0 {
				if _, err := w.Write(t.getFirstBuf()); err != nil {
					return err
				}
				break
			}

			time.Sleep(time.Second)
			retry -= 1
		}
		if retry < 0 {
			w.WriteHeader(http.StatusNotFound)
			return errors.New("no stream")
		}
	}

	//写入快速启动缓冲
	if err := t.bootBufRead(func(data []byte) error {
		if len(data) != 0 {
			if _, err := w.Write(data); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	size := uint32(2)
	if tmp, ok := t.common.K_v.LoadV("直播流实时回放缓存").(float64); ok && int(tmp) > int(size) && tmp < math.MaxUint32 {
		size = uint32(tmp)
	}

	w = pweb.WithCache(w, size)

	var cancelRec = t.Stream_msg.Pull_tag(map[string]func([]byte) bool{
		`data`: func(b []byte) bool {
			select {
			case <-ctx.Done():
				return true
			default:
			}
			if len(b) == 0 {
				return true
			}

			_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 30))
			if n, err := w.Write(b); err != nil || n == 0 {
				if errors.Is(err, pio.ErrCacheWriterBusy) {
					plog.L(`I: `, r.RemoteAddr, "回放缓存跳过，或许应该增加`直播流实时回放缓存`")
				} else {
					plog.L(`W: `, r.RemoteAddr, err)
					cancel(err)
					return true
				}
			}
			return false
		},
		`close`: func(_ []byte) bool {
			return true
		},
	})

	<-ctx.Done()
	cancelRec()

	if e := stopFunc(t); e != nil {
		return e
	}

	return nil
}

func (t *M4SStream) bootBufPush(buf []byte) {
	t.boot_buf_locker.Lock()
	defer t.boot_buf_locker.Unlock()
	t.boot_buf = t.boot_buf[:0]
	t.boot_buf = append(t.boot_buf, buf...)
}

func (t *M4SStream) bootBufRead(r func(data []byte) error) error {
	t.boot_buf_locker.RLock()
	defer t.boot_buf_locker.RUnlock()
	return r(t.boot_buf)
}
