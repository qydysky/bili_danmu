package reply

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"

	videoInfo "github.com/qydysky/bili_danmu/Reply/F/videoInfo"
	pctx "github.com/qydysky/part/ctx"
	pe "github.com/qydysky/part/errors"
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
	pweb "github.com/qydysky/part/web"
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

func (link *m4s_link_item) download(reqPool *pool.Buf[reqf.Req], reqConfig reqf.Rval) error {
	link.status = 1 // 设置切片状态为正在下载
	link.err = nil
	link.tryDownCount += 1
	link.data.Reset()

	r := reqPool.Get()
	defer reqPool.Put(r)
	reqConfig.Url = link.Url

	// fmt.Println(`T: `, `下载`, link.Base)
	// defer t.log.L(`T: `, `下载完成`, link.Base, link.status, link.err)

	if e := r.Reqf(reqConfig); e != nil && !errors.Is(e, io.EOF) {
		// t.log.L(`T: `, `下载错误`, link.Base, e)
		// if !reqf.IsTimeout(e) {
		// 	// 发生非超时错误
		// 	link.err = e
		// 	link.tryDownCount = 3 // 设置切片状态为下载失败
		// }
		link.status = 3 // 设置切片状态为下载失败
		link.err = e
		return e
	} else if e = link.data.Append(r.Respon); e != nil {
		link.status = 3 // 设置切片状态为下载失败
		link.err = e
		return e
	} else {
		link.status = 2 // 设置切片状态为下载完成
		return nil
	}
}

func (t *M4SStream) MarshalJSON() ([]byte, error) {
	t.common.Rev = c.C.Rev
	t.common.Watched = c.C.Watched
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
	// 获取流地址
	t.common.Live_want_qn = t.config.want_qn
	if F.Get(t.common).Get(`Live`); len(t.common.Live) == 0 {
		return false
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

	var nomcdn bool
	if v, ok := t.common.K_v.LoadV("直播流不使用mcdn").(bool); ok && v {
		nomcdn = true
	}

	r := t.reqPool.Get()
	defer t.reqPool.Put(r)

	for _, v := range t.common.Live {
		if nomcdn && strings.Contains(v.Url, ".mcdn.") {
			v.Disable(time.Now().Add(time.Hour * 100))
			t.log.L(`I: `, `停用流服务器`, F.ParseHost(v.Url))
			continue
		}

		if e := r.Reqf(reqf.Rval{
			Url:   v.Url,
			Proxy: t.common.Proxy,
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
			t.log.L(`W: `, F.ParseHost(v.Url), e)
			v.DisableAuto()
			continue
		}

		if r.Response == nil {
			t.log.L(`W: `, `live响应错误`, F.ParseHost(v.Url))
			v.DisableAuto()
			continue
		} else if r.Response.StatusCode&200 != 200 {
			t.log.L(`W: `, `live响应错误`, F.ParseHost(v.Url), r.Response.Status)
			v.DisableAuto()
			continue
		}

		// 显示使用流服务器
		t.log.L(`I: `, `使用流服务器`, F.ParseHost(v.Url))
	}

	return t.common.ValidLive() != nil
}

func (t *M4SStream) fetchParseM3U8(lastM4s *m4s_link_item, fmp4ListUpdateTo float64) (m4s_links []*m4s_link_item, e error) {
	{
		n := t.common.ValidNum()
		if d, ok := t.common.K_v.LoadV("fmp4获取更多服务器").(bool); ok && d && n <= 1 {
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
	for _, v := range t.common.Live {
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
			// 1min后重新启用
			v.DisableAuto()
			t.log.L("W: ", fmt.Sprintf("服务器 %s 发生故障 %s", F.ParseHost(v.Url), pe.ErrorFormat(err, pe.ErrSimplifyFunc)))
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

		// m3u8字节流
		var m3u8_respon = r.Respon

		// base64解码
		if len(m3u8_respon) != 0 && !bytes.Contains(m3u8_respon, []byte("#")) {
			var err error
			m3u8_respon, err = base64.StdEncoding.DecodeString(string(m3u8_respon))
			if err != nil {
				e = err
				return
			}
		}

		// 解析m3u8
		// var tmp []*m4s_link_item
		var lastNo int
		if lastM4s != nil {
			lastNo, _ = lastM4s.getNo()
		}

		m3u := bytes.Split(m3u8_respon, []byte("\n"))
		for i := 0; i < len(m3u); i++ {
			line := m3u[i]
			if len(line) == 0 {
				continue
			}

			var (
				m4s_link string //切片文件名
				isHeader bool
			)

			if line[0] == '#' {
				if strings.HasPrefix(string(line), "#EXT-X-MAP") {
					if lastM4s != nil {
						continue
					}
					e := bytes.Index(line[16:], []byte(`"`)) + 16
					m4s_link = string(line[16:e])
					isHeader = true
				} else if strings.HasPrefix(string(line), "#EXT-X-STREAM-INF") {
					// m3u8 指向新连接
					i += 1
					line = m3u[i]
					newUrl := strings.TrimSpace(string(line))
					t.log.L(`I: `, `指向新连接`, v.Host(), "=>", F.ParseHost(newUrl))
					v.SetUrl(newUrl)
					continue
				} else {
					continue
				}
			} else {
				m4s_link = string(line)
			}

			if !isHeader {
				// 只增加新的切片
				if no, _ := strconv.Atoi(m4s_link[:len(m4s_link)-4]); lastNo >= no {
					continue
				}
			}

			//将切片添加到返回切片数组
			p := t.getM4s()
			p.Url = F.ResolveReferenceLast(v.Url, m4s_link+"?trid="+F.ParseQuery(v.Url, "trid="))
			p.Base = m4s_link
			p.isHeader = isHeader
			p.createdTime = time.Now()
			m4s_links = append(m4s_links, p)
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
			t.log.L(`I: `, `发现`, linksFirstNo-lastNo-1, `个切片遗漏，重新下载`)
			for guess_no := linksFirstNo - 1; guess_no > lastNo; guess_no-- {
				//将切片添加到返回切片数组前
				p := t.getM4s()
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
		f := file.New(t.config.save_path+"tmp.create", 0, true)
		f.Create()
		_ = f.Delete()
	} else {
		t.log.L(`W: `, err)
	}
}

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
		var startCount uint = 3
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
				readTO        int64 = 3
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
					ticker   = time.NewTicker(time.Second)
					buff     = slice.New[byte]()
					keyframe = slice.New[byte]()
					buf      = make([]byte, 1<<16)
				)
				defer ticker.Stop()

				for {
					n, e := pipe.Read(buf)
					_ = buff.Append(buf[:n])
					if e != nil {
						pctx.PutVal(cancelC, &errCtx, e)
						break
					}

					select {
					case <-ticker.C:
					default:
						continue
					}

					if !buff.IsEmpty() {
						keyframe.Reset()
						buf, unlock := buff.GetPureBufRLock()
						front_buf, last_available_offset, e := Search_stream_tag(buf, keyframe)
						unlock()
						if e != nil {
							if strings.Contains(e.Error(), `no found available tag`) {
								continue
							}
							pctx.PutVal(cancelC, &errCtx, errors.New("[decoder]"+e.Error()))
							//丢弃所有数据
							buff.Reset()
						}
						// 存在有效数据
						if len(front_buf) != 0 || keyframe.Size() != 0 {
							leastReadUnix.Store(time.Now().Unix())
						}
						if len(front_buf) != 0 && len(t.first_buf) == 0 {
							t.first_buf = make([]byte, len(front_buf))
							copy(t.first_buf, front_buf)
							// fmt.Println("write front_buf")
							// t.Stream_msg.PushLock_tag(`data`, t.first_buf)
							t.msg.Push_tag(`load`, t)
						}
						if keyframe.Size() != 0 {
							if len(t.first_buf) == 0 {
								t.log.L(`W: `, `flv未接收到起始段`)
								pctx.PutVal(cancelC, &errCtx, errors.New(`flv未接收到起始段`))
								break
							}
							buf, unlock := keyframe.GetPureBufRLock()
							t.bootBufPush(buf)
							t.Stream_msg.PushLock_tag(`data`, buf)
							unlock()
							keyframe.Reset()
							t.frameCount += 1
							t.msg.Push_tag(`keyFrame`, t)
						}
						if last_available_offset > 1 {
							// fmt.Println("write Sync")
							_ = buff.RemoveFront(last_available_offset - 1)
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

func (t *M4SStream) saveStreamM4s() (e error) {

	if v, ok := t.common.K_v.LoadV(`debug模式`).(bool); ok && v {
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

	var (
		// 同时下载数限制
		downloadLimit    = funcCtrl.NewBlockFuncN(3)
		buf              = slice.New[byte]()
		fmp4Decoder      = NewFmp4Decoder()
		keyframe         = slice.New[byte]()
		lastM4s          *m4s_link_item
		to               = 3
		fmp4ListUpdateTo = 5.0
		fmp4Count        = 0
		startT           = time.Now()
	)

	if v, ok := t.common.K_v.LoadV(`fmp4音视频时间戳容差s`).(float64); ok && v > 0.1 {
		fmp4Decoder.AVTDiff = v
	}
	if v, ok := t.common.K_v.LoadV(`fmp4切片下载超时s`).(float64); ok && to < int(v) {
		to = int(v)
	}
	if v, ok := t.common.K_v.LoadV(`fmp4列表更新超时s`).(float64); ok && fmp4ListUpdateTo < v {
		fmp4ListUpdateTo = v
	}

	// 下载循环
	for download_seq := []*m4s_link_item{}; ; {
		// 获取解析m3u8
		{
			// 防止过快的下载
			dru := time.Since(startT).Seconds()
			if wait := float64(fmp4Count) - dru - 1; wait > 5 {
				time.Sleep(time.Second * 5)
			} else if wait > 2 {
				time.Sleep(time.Duration(wait) * time.Second)
			} else {
				time.Sleep(time.Second * 2)
			}

			var m4s_links, err = t.fetchParseM3U8(lastM4s, fmp4ListUpdateTo)
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

			// n秒未产出切片
			fmp4Count = len(m4s_links)
			if dru > fmp4ListUpdateTo && fmp4Count == 0 {
				e = fmt.Errorf("%.2f 秒未产出切片", dru)
				t.log.L("E: ", "获取解析m3u8发生错误", e)
				break
			}
			if fmp4Count != 0 {
				startT = time.Now()
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
			downOk := true
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
						oldHost := linkUrl.Host
						// 将此切片服务器设置停用
						hadDisable := t.common.DisableLiveAuto(oldHost)
						// 从其他服务器获取此切片
						if vl := t.common.ValidLive(); vl == nil {
							return errors.New(`全部流服务器故障`)
						} else {
							linkUrl.Host = vl.Host()
							if !hadDisable {
								t.log.L(`W: `, `切片下载失败，故障转移`, oldHost, ` -> `, linkUrl.Host)
							}
						}
						download_seq[i].Url = linkUrl.String()
					}
				}

				done := downloadLimit.Block()
				go func(link *m4s_link_item) {
					defer done()

					if e := link.download(t.reqPool, reqf.Rval{
						Timeout:     to * 1000,
						WriteLoopTO: (to + 2) * 1000,
						Proxy:       t.common.Proxy,
						Header: map[string]string{
							`Connection`: `close`,
						},
					}); e != nil {
						downOk = false
						t.log.L(`W: `, `切片下载失败`, link.Base, e)
					}
				}(download_seq[i])
			}

			// 等待队列下载完成
			downloadLimit.BlockAll()()

			if downOk {
				break
			}
		}

		// 传递已下载切片
		for k := 0; k < len(download_seq) && download_seq[k].status == 2; k++ {

			if download_seq[k].isInit() {
				{
					buf, unlock := download_seq[k].data.GetPureBufRLock()
					front_buf, e := fmp4Decoder.Init_fmp4(buf)
					unlock()
					if e != nil {
						t.log.L(`E: `, e, `重试!`)
						download_seq[k].status = 3
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
				t.putM4s(download_seq[k])
				download_seq = append(download_seq[:k], download_seq[k+1:]...)
				k -= 1
				continue
			} else if t.first_buf == nil {
				t.putM4s(download_seq[k])
				download_seq = append(download_seq[:k], download_seq[k+1:]...)
				k -= 1
				continue
			}

			if e := download_seq[k].data.AppendTo(buf); e != nil {
				t.log.L(`E: `, e)
			}
			t.putM4s(download_seq[k])
			download_seq = append(download_seq[:k], download_seq[k+1:]...)
			k -= 1

			buff, unlock := buf.GetPureBufRLock()
			last_available_offset, err := fmp4Decoder.Search_stream_fmp4(buff, keyframe)
			unlock()

			if err != nil && !errors.Is(err, io.EOF) {
				t.log.L(`E: `, err)
				//丢弃所有数据
				buf.Reset()
				e = err
				return
			}

			// no, _ := download_seq[k].getNo()
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
					// 当cut时，取消上次录制
					contextC, cancel := context.WithCancel(mainCtx)
					fc.FlashWithCallback(cancel)

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
					if e, _ := videoInfo.Save.Run(contextC, ms); e != nil {
						l.L(`E: `, e)
					}

					go StartRecDanmu(contextC, ms.GetSavePath())                           //保存弹幕
					go Ass_f(contextC, ms.GetSavePath(), ms.GetSavePath()+"0", time.Now()) //开始ass

					startT := time.Now()
					if e := ms.PusherToFile(contextC, ms.GetSavePath()+`0.`+ms.GetStreamType(), startf, stopf); e != nil {
						l.L(`E: `, e)
					}
					duration := time.Since(startT)

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

	contextC, done := pctx.WaitCtx(contextC)
	defer done()

	_, _ = f.Write(t.getFirstBuf(), true)
	cancelRec := t.Stream_msg.Pull_tag(map[string]func([]byte) bool{
		`data`: func(b []byte) bool {
			select {
			case <-contextC.Done():
				return true
			default:
			}
			if len(b) == 0 {
				return true
			}
			_, _ = f.Write(b, true)
			return false
		},
		`close`: func(_ []byte) bool {
			return true
		},
	})
	<-contextC.Done()
	cancelRec()

	if e := stopFunc(t); e != nil {
		return e
	}

	return nil
}

// 流服务推送方法
//
// 在客户端存在某种代理时，将有可能无法监测到客户端关闭，这有可能导致goroutine泄漏
func (t *M4SStream) PusherToHttp(conn net.Conn, w http.ResponseWriter, r *http.Request, startFunc func(*M4SStream) error, stopFunc func(*M4SStream) error) error {
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

	//写入头
	{
		retry := 5
		for retry > 0 {
			select {
			case <-r.Context().Done():
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

	var cancelRec = t.Stream_msg.Pull_tag(map[string]func([]byte) bool{
		`data`: func(b []byte) bool {
			select {
			case <-r.Context().Done():
				return true
			default:
			}
			if len(b) == 0 {
				return true
			}
			_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 30))
			if n, err := w.Write(b); err != nil || n == 0 {
				return true
			}
			return false
		},
		`close`: func(_ []byte) bool {
			return true
		},
	})

	<-r.Context().Done()
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
