package reply

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"

	p "github.com/qydysky/part"
	file "github.com/qydysky/part/file"
	funcCtrl "github.com/qydysky/part/funcCtrl"
	idpool "github.com/qydysky/part/idpool"
	log "github.com/qydysky/part/log"
	msgq "github.com/qydysky/part/msgq"
	pool "github.com/qydysky/part/pool"
	reqf "github.com/qydysky/part/reqf"
	signal "github.com/qydysky/part/signal"
	slice "github.com/qydysky/part/slice"
	pstring "github.com/qydysky/part/strings"
)

type M4SStream struct {
	Status               *signal.Signal     //IsLive()是否运行中
	exitSign             *signal.Signal     //IsLive()是否等待退出中
	log                  *log.Log_interface //日志
	config               M4SStream_Config   //配置
	stream_last_modified time.Time          //流地址更新时间
	// stream_expires       int64              //流到期时间
	// stream_hosts      psync.Map  //使用的流服务器
	stream_type       string     //流类型
	Stream_msg        *msgq.Msgq //流数据消息 tag:data
	first_buf         []byte     //m4s起始块 or flv起始块
	boot_buf          [][]byte   //快速启动缓冲
	boot_buf_size     int        //快速启动缓冲长度
	boot_buf_locker   funcCtrl.BlockFunc
	last_m4s          *m4s_link_item           //最后一个切片
	m4s_pool          *pool.Buf[m4s_link_item] //切片pool
	common            c.Common                 //通用配置副本
	Current_save_path string                   //明确的直播流保存目录
	Callback_start    func(*M4SStream) error   //实例开始的回调
	Callback_startRec func(*M4SStream) error   //录制开始的回调
	Callback_stopRec  func(*M4SStream)         //录制结束的回调
	Callback_stop     func(*M4SStream)         //实例结束的回调
	reqPool           *idpool.Idpool           //请求池
}

type M4SStream_Config struct {
	save_path     string //直播流保存目录
	want_qn       int    //直播流清晰度
	want_type     string //直播流类型
	save_as_mp4   bool   //直播hls流保存为MP4
	banlance_host bool   //直播hls流故障转移
}

type m4s_link_item struct {
	Url          string    // m4s链接
	Base         string    // m4s文件名
	status       int       // 下载状态 0:未下载 1:正在下载 2:下载完成 3:下载失败
	tryDownCount int       // 下载次数 当=3时，不再下载，忽略此块
	err          error     // 下载中出现的错误
	data         []byte    // 下载的数据
	createdTime  time.Time // 创建时间
	pooledTime   time.Time // 到pool时间
}

func (t *m4s_link_item) copyTo(to *m4s_link_item) {
	// fmt.Println("copy to ", t.Base)
	to.Url = t.Url
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
	t.Base = ""
	t.status = 0
	t.tryDownCount = 0
	t.data = t.data[:0]
	t.createdTime = time.Now()
	return t
}

func (t *m4s_link_item) isInit() bool {
	return strings.Contains(t.Base, "h")
}

func (t *m4s_link_item) getNo() (int, error) {
	var base = t.Base
	if t.isInit() {
		base = base[1:]
	}
	return strconv.Atoi(base[:len(base)-4])
}

func (t *M4SStream) getM4s() (p *m4s_link_item) {
	if t.m4s_pool == nil {
		t.m4s_pool = pool.New(
			func() *m4s_link_item {
				return &m4s_link_item{}
			},
			func(t *m4s_link_item) bool {
				return !t.pooledTime.IsZero() || t.createdTime.After(t.pooledTime) || time.Now().Before(t.pooledTime.Add(time.Second*10))
			},
			func(t *m4s_link_item) *m4s_link_item {
				return t.reset()
			},
			func(t *m4s_link_item) *m4s_link_item {
				t.pooledTime = time.Now()
				return t
			},
			50,
		)
	}
	return t.m4s_pool.Get()
}

func (t *M4SStream) putM4s(ms ...*m4s_link_item) {
	t.m4s_pool.Put(ms...)
}

func (t *M4SStream) Common() c.Common {
	return t.common
}

func (t *M4SStream) LoadConfig(common c.Common, l *log.Log_interface) {
	t.common = common
	t.log = l.Base(`直播流保存`)

	//读取配置
	if path, ok := common.K_v.LoadV("直播流保存位置").(string); ok {
		if path, err := filepath.Abs(path); err == nil {
			if fs, err := os.Stat(path); err != nil {
				if errors.Is(err, os.ErrNotExist) {
					if err := p.File().NewPath(path); err != nil {
						t.log.L(`E: `, `直播流保存位置错误`, err)
						return
					}
				} else {
					t.log.L(`E: `, `直播流保存位置错误`, err)
					return
				}
			} else if !fs.IsDir() {
				t.log.L(`E: `, `直播流保存位置不是目录`)
				return
			}
			t.config.save_path = path + "/"
		} else {
			t.log.L(`E: `, `直播流保存位置错误`, err)
			return
		}
	}
	if v, ok := common.K_v.LoadV(`直播hls流保存为MP4`).(bool); ok {
		t.config.save_as_mp4 = v
	}
	if v, ok := common.K_v.LoadV(`直播hls流故障转移`).(bool); ok {
		t.config.banlance_host = v
	}
	if v, ok := common.K_v.LoadV(`直播流清晰度`).(float64); ok {
		t.config.want_qn = int(v)
	}
	if v, ok := common.K_v.LoadV(`直播流类型`).(string); ok {
		t.config.want_type = v
	}
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
	if F.Get(&t.common).Get(`Live`); len(t.common.Live) == 0 {
		return false
	}

	// 保存流类型
	if strings.Contains(t.common.Live[0].Url, `m3u8`) {
		if t.config.save_as_mp4 {
			t.stream_type = "mp4"
		} else {
			t.stream_type = "m3u8"
		}
	} else if strings.Contains(t.common.Live[0].Url, `flv`) {
		t.stream_type = "flv"
	}

	// 检查是否可以获取
	CookieM := make(map[string]string)
	t.common.Cookie.Range(func(k, v interface{}) bool {
		CookieM[k.(string)] = v.(string)
		return true
	})

	for _, v := range t.common.Live {
		req := t.reqPool.Get()
		r := req.Item.(*reqf.Req)
		if e := r.Reqf(reqf.Rval{
			Url:       v.Url,
			Retry:     10,
			SleepTime: 1000,
			Proxy:     t.common.Proxy,
			Header: map[string]string{
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
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
			t.log.L(`W: `, e)
		}

		if r.Response == nil {
			t.log.L(`W: `, `live响应错误`)
			t.common.Live = t.common.Live[1:]
		} else if r.Response.StatusCode&200 != 200 {
			t.log.L(`W: `, `live响应错误`, r.Response.Status, string(r.Respon))
			t.common.Live = t.common.Live[1:]
		}
		t.reqPool.Put(req)
	}

	return len(t.common.Live) != 0
}

func (t *M4SStream) fetchParseM3U8() (m4s_links []*m4s_link_item, m3u8_addon []byte, e error) {
	if t.common.ValidLive() == nil {
		e = errors.New("全部流服务器发生故障")
		return
	}

	// 开始请求
	req := t.reqPool.Get()
	defer t.reqPool.Put(req)
	r := req.Item.(*reqf.Req)

	// 请求解析m3u8内容
	for k, v := range t.common.Live {
		// 跳过尚未启用的live地址
		if !v.Valid() {
			continue
		}

		m3u8_url, err := url.Parse(v.Url)
		if err != nil {
			e = err
			return
		}

		// 设置请求参数
		rval := reqf.Rval{
			Url:     m3u8_url.String(),
			Retry:   2,
			Timeout: 2000,
			Proxy:   c.C.Proxy,
			Header: map[string]string{
				`Host`:            m3u8_url.Host,
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
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
			rval.Header[`If-Modified-Since`] = t.stream_last_modified.Add(time.Second).Format("Mon, 02 Jan 2006 15:04:05 CST")
		}

		if err := r.Reqf(rval); err != nil {
			// 1min后重新启用
			t.common.Live[k].DisableAuto()
			t.log.L("W: ", fmt.Sprintf("服务器 %s 发生故障 %s", m3u8_url.Host, err.Error()))
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
			if lm, e := time.Parse("Mon, 02 Jan 2006 15:04:05 CST", last_mod[0]); e == nil {
				t.stream_last_modified = lm
			}
		}

		// m3u8字节流
		var m3u8_respon = r.Respon

		// base64解码
		if len(m3u8_respon) != 0 && !bytes.Contains(m3u8_respon, []byte("#")) {
			m3u8_respon, err = base64.StdEncoding.DecodeString(string(m3u8_respon))
			if err != nil {
				e = err
				return
			}
		}

		// 解析m3u8
		var tmp []*m4s_link_item
		var lastNo int
		if t.last_m4s != nil {
			lastNo, _ = t.last_m4s.getNo()
		}
		for _, line := range bytes.Split(m3u8_respon, []byte("\n")) {
			if len(line) == 0 {
				continue
			}

			var m4s_link string //切片文件名

			//获取附加的m3u8字节 忽略bili定制拓展
			if !bytes.Contains(line, []byte(`#EXT-X-BILI`)) {
				if t.last_m4s == nil {
					m3u8_addon = append(m3u8_addon, line...)
					m3u8_addon = append(m3u8_addon, []byte("\n")...)
				} else {
					if bytes.Contains(line, []byte(`#EXTINF`)) ||
						!bytes.Contains(line, []byte(`#`)) {
						m3u8_addon = append(m3u8_addon, line...)
						m3u8_addon = append(m3u8_addon, []byte("\n")...)
					}
				}
			}

			//获取切片文件名
			if bytes.Contains(line, []byte("EXT-X-MAP")) {
				o := bytes.Index(line, []byte(`EXT-X-MAP:URI="`)) + 15
				e := bytes.Index(line[o:], []byte(`"`)) + o
				m4s_link = string(line[o:e])
			} else if bytes.Contains(line, []byte("#EXT-X")) { //忽略扩展标签
				continue
			} else if bytes.Contains(line, []byte(".m4s")) {
				m4s_link = string(line)
			} else {
				continue
			}

			//获取切片地址
			u, err := url.Parse("./" + m4s_link + "?trid=" + m3u8_url.Query().Get("trid"))
			if err != nil {
				e = err
				return
			}

			{
				tmpBase := m4s_link
				// fmt.Println(tmpBase, t.last_m4s != nil)
				if tmpBase[0] == 'h' {
					if t.last_m4s != nil {
						continue
					} else {
						tmpBase = tmpBase[1:]
					}
				}
				if no, _ := strconv.Atoi(tmpBase[:len(tmpBase)-4]); lastNo >= no {
					// fmt.Println("skip", no)
					continue
				}
			}

			// fmt.Println("->", m4s_link)
			//将切片添加到返回切片数组
			p := t.getM4s()
			p.Url = m3u8_url.ResolveReference(u).String()
			p.Base = m4s_link
			p.createdTime = time.Now()
			tmp = append(tmp, p)
		}

		if len(tmp) == 0 {
			// fmt.Println("->", "empty", lastNo)
			return
		}

		// 检查是否服务器发生故障,产出多个切片
		if t.last_m4s != nil {
			timed := tmp[len(tmp)-1].createdTime.Sub(t.last_m4s.createdTime).Seconds()
			nos, _ := tmp[len(tmp)-1].getNo()
			noe, _ := t.last_m4s.getNo()
			if timed > 5 && nos-noe == 0 {
				// 1min后重新启用
				t.common.Live[k].DisableAuto()
				t.log.L("W: ", fmt.Sprintf("服务器 %s 发生故障 %d 秒产出了 %d 切片", m3u8_url.Host, int(timed), nos-noe))
				if t.common.ValidLive() == nil {
					e = errors.New("全部切片服务器发生故障")
					break
				}
				continue
			}
		}

		m4s_links = append(m4s_links, tmp...)

		// 首次下载
		if t.last_m4s == nil {
			return
		}

		// 去掉初始段 及 last之前的切片
		{
			last_no, _ := t.last_m4s.getNo()
			for k := 0; k < len(m4s_links); k++ {
				m4s_link := m4s_links[k]
				// 剔除初始段
				if m4s_link.isInit() {
					m4s_links = append(m4s_links[:k], m4s_links[k+1:]...)
					k--
					continue
				}
				no, _ := m4s_link.getNo()
				if no < last_no {
					continue
				} else if no == last_no {
					// 只返回新增加的切片,去掉无用切片
					t.putM4s(m4s_links[:k+1]...)
					m4s_links = m4s_links[k+1:]
					// 只返回新增加的m3u8_addon字节
					if index := bytes.Index(m3u8_addon, []byte(m4s_link.Base)); index != -1 {
						index += len([]byte(m4s_link.Base))
						if index == len(m3u8_addon) {
							m3u8_addon = []byte{}
						} else {
							m3u8_addon = m3u8_addon[index+1:]
						}
					}
					return
				} else if no == last_no+1 {
					// 刚刚好承接之前的结尾
					return
				} else {
					break
				}
			}
		}

		// 来到此处说明出现了丢失 尝试补充
		guess_end_no, _ := m4s_links[0].getNo()
		current_no, _ := t.last_m4s.getNo()

		if guess_end_no < current_no {
			return
		}

		t.log.L(`I: `, `发现`, guess_end_no-current_no-1, `个切片遗漏，重新下载`)
		for guess_no := guess_end_no - 1; guess_no > current_no; guess_no-- {
			// 补充m3u8
			m3u8_addon = append([]byte("#EXTINF:1.00\n"+strconv.Itoa(guess_no)+".m4s\n"), m3u8_addon...)

			//获取切片地址
			u, err := url.Parse("./" + strconv.Itoa(guess_no) + `.m4s`)
			if err != nil {
				e = err
				return
			}

			//将切片添加到返回切片数组前
			p := t.getM4s()
			p.Url = m3u8_url.ResolveReference(u).String()
			p.Base = strconv.Itoa(guess_no) + `.m4s`
			p.createdTime = time.Now()
			m4s_links = append([]*m4s_link_item{p}, m4s_links...)
		}

		// 请求解析成功，退出获取循环
		return
	}

	if e != nil {
		e = errors.New(e.Error() + " 未能找到可用流服务器")
	}
	return
}

func (t *M4SStream) saveStream() (e error) {
	// 设置保存路径
	t.Current_save_path = t.config.save_path + "/" +
		time.Now().Format("2006_01_02-15_04_05") + "-" +
		strconv.Itoa(t.common.Roomid) + "-" +
		strings.NewReplacer("\\", "", "\\/", "", ":", "", "*", "", "?", "", "\"", "", "<", "", ">", "", "|", "").Replace(t.common.Title) + "-" +
		pstring.Rand(2, 3) +
		`/`

	// 清除初始值
	t.last_m4s = nil

	// 显示保存位置
	if rel, err := filepath.Rel(t.config.save_path, t.Current_save_path); err == nil {
		t.log.L(`I: `, "保存到", rel+`/0.`+t.stream_type)
	} else {
		t.log.L(`W: `, err)
	}
	for _, v := range t.common.Stream_url {
		t.log.L(`I: `, "流地址:", v)
	}

	// 录制回调
	if t.Callback_startRec != nil {
		if err := t.Callback_startRec(t); err != nil {
			t.log.L(`W: `, `开始录制回调错误`, err.Error())
			return err
		}
	}
	if t.Callback_stopRec != nil {
		defer t.Callback_stopRec(t)
	}

	// 获取流
	switch t.stream_type {
	case `m3u8`:
		fallthrough
	case `mp4`:
		e = t.saveStreamM4s()
	case `flv`:
		e = t.saveStreamFlv()
	default:
		e = errors.New("undefind stream type")
		t.log.L(`E: `, e)
	}

	return
}

func (t *M4SStream) saveStreamFlv() (e error) {
	//对每个直播流进行尝试
	for _, v := range t.common.Live {
		surl, err := url.Parse(v.Url)
		if err != nil {
			t.log.L(`E: `, err)
			e = err
			continue
		}

		//结束退出
		if !t.Status.Islive() {
			return
		}

		// 如果被主动关闭，则退出saveStreamFlv，否则继续尝试其他live
		s := signal.Init()

		//开始获取
		req := t.reqPool.Get()
		{
			r := req.Item.(*reqf.Req)

			go func() {
				select {
				//停止录制
				case <-t.Status.WaitC():
					r.Cancel()
				//当前连接终止
				case <-s.WaitC():
				}
			}()

			out := file.New(t.Current_save_path+`0.flv`, -1, true).File()

			rc, rw := io.Pipe()
			var leastReadUnix = time.Now().Unix()
			// read timeout
			go func() {
				timer := time.NewTicker(5 * time.Second)
				defer timer.Stop()
				for {
					select {
					case <-s.WaitC():
						return
					case curT := <-timer.C:
						if curT.Unix()-leastReadUnix > 5 {
							t.log.L(`W: `, "5s未接收到任何数据")
							// 5s未接收到任何数据
							r.Cancel()
							return
						}
						if v, ok := c.C.K_v.LoadV(`直播流清晰度`).(float64); ok {
							if t.config.want_qn != int(v) {
								t.log.L(`I: `, "直播流清晰度改变:", t.common.Qn[t.config.want_qn], "=>", t.common.Qn[int(v)])
								t.config.want_qn = int(v)
								s.Done()
								r.Cancel()
								return
							}
						}
					}
				}
			}()

			// read
			go func() {
				var (
					ticker = time.NewTicker(time.Second)
					buff   = slice.New[byte]()
					buf    = make([]byte, 1<<16)
				)
				for {
					n, e := rc.Read(buf)
					buff.Append(buf[:n])
					if e != nil {
						out.Close()
						t.Stream_msg.Push_tag(`close`, nil)
						break
					}
					leastReadUnix = time.Now().Unix()

					skip := true
					select {
					case <-ticker.C:
						skip = false
					default:
					}
					if skip {
						continue
					}

					if !buff.IsEmpty() {
						front_buf, keyframe, last_available_offset, e := Search_stream_tag(buff.GetCopyBuf())
						if e != nil {
							if strings.Contains(e.Error(), `no found available tag`) {
								continue
							}
						}
						if len(front_buf)+len(keyframe) != 0 {
							if len(front_buf) != 0 {
								t.first_buf = front_buf
								// fmt.Println("write front_buf")
								out.Write(front_buf)
								t.Stream_msg.Push_tag(`data`, front_buf)
							}
							for _, frame := range keyframe {
								// fmt.Println("write frame")
								out.Write(frame)
								t.bootBufPush(frame)
								t.Stream_msg.Push_tag(`data`, frame)
							}
						}
						if last_available_offset > 1 {
							// fmt.Println("write Sync")
							buff.RemoveFront(last_available_offset - 1)
							out.Sync()
						}
					}
				}

				buf = nil
				buff.Reset()
			}()

			CookieM := make(map[string]string)
			t.common.Cookie.Range(func(k, v interface{}) bool {
				CookieM[k.(string)] = v.(string)
				return true
			})

			t.log.L(`I: `, `flv下载开始`)

			r.Reqf(reqf.Rval{
				Url:              surl.String(),
				SaveToPipeWriter: rw,
				NoResponse:       true,
				Async:            true,
				Proxy:            t.common.Proxy,
				Header: map[string]string{
					`Host`:            surl.Host,
					`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
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
					t.log.L(`I: `, `flv下载停止`)
				} else if err != nil && !reqf.IsTimeout(err) {
					e = err
					t.log.L(`E: `, `flv下载失败:`, err)
				}
			}
		}
		t.reqPool.Put(req)

		if s.Islive() {
			s.Done()
		} else {
			return
		}
	}

	e = errors.New("未能找到可用流服务器")
	return
}

func (t *M4SStream) saveStreamM4s() (e error) {
	// 同时下载数限制
	var download_limit = &funcCtrl.BlockFuncN{
		Max: 3,
	}

	var out *file.File
	if t.config.save_as_mp4 {
		out = file.New(t.Current_save_path+`0.mp4`, 0, false)
		defer out.Close()
	}

	//
	var (
		buf              = slice.New[byte]()
		fmp4KeyFrames    = slice.New[byte]()
		fmp4KeyFramesBuf []byte
		fmp4Decoder      = &Fmp4Decoder{}
		// flashingSer      bool
	)

	// 下载循环
	for download_seq := []*m4s_link_item{}; ; {

		// 刷新流地址
		// if !flashingSer && int64(t.common.Live[0].Expires)-time.Now().Unix() < 60 {
		// 	flashingSer = true
		// 	t.log.L(`T: `, `刷新流地址...`)
		// 	go func() {
		// 		t.fetchCheckStream()
		// 		flashingSer = false
		// 	}()
		// }

		// 存在待下载切片
		if len(download_seq) != 0 {
			var downingCount = 0 //本轮下载数量
			// 下载切片
			for _, v := range download_seq {

				// 已下载但还未移除的切片
				if v.status == 2 {
					continue
				}

				// 每次最多只下载10个切片
				if downingCount >= 10 {
					t.log.L(`T: `, `延迟切片下载 数量(`, len(download_seq)-downingCount, `)`)
					break
				}
				downingCount += 1

				download_limit.Block(func() {
					time.Sleep(time.Millisecond * 10)
				})

				// 故障转移
				if v.status == 3 {
					if linkUrl, e := url.Parse(v.Url); e == nil {
						oldHost := linkUrl.Host
						// 将此切片服务器设置停用
						t.common.DisableLiveAuto(oldHost)
						// 从其他服务器获取此切片
						if vl := t.common.ValidLive(); vl == nil {
							return errors.New(`全部流服务器故障`)
						} else {
							linkUrl.Host = vl.Host()
							t.log.L(`W: `, `切片下载失败，故障转移`, oldHost, ` -> `, linkUrl.Host)
						}
						v.Url = linkUrl.String()
					}
				}

				go func(link *m4s_link_item, path string) {
					defer download_limit.UnBlock()

					link.status = 1 // 设置切片状态为正在下载
					link.tryDownCount += 1

					req := t.reqPool.Get()
					defer t.reqPool.Put(req)

					r := req.Item.(*reqf.Req)
					reqConfig := reqf.Rval{
						Url:     link.Url,
						Timeout: 3000,
						Proxy:   t.common.Proxy,
						Header: map[string]string{
							`Connection`: `close`,
						},
					}
					if !t.config.save_as_mp4 {
						reqConfig.SaveToPath = path + link.Base
					}

					// t.log.L(`T: `, `下载`, link.Base)
					// defer t.log.L(`T: `, `下载完成`, link.Base, link.status, link.err)

					if e := r.Reqf(reqConfig); e != nil && !errors.Is(e, io.EOF) {
						// t.log.L(`T: `, `下载错误`, link.Base, e)
						// if !reqf.IsTimeout(e) {
						// 	// 发生非超时错误
						// 	link.err = e
						// 	link.tryDownCount = 3 // 设置切片状态为下载失败
						// }
						link.status = 3 // 设置切片状态为下载失败
					} else {
						// if usedt := r.UsedTime.Seconds(); usedt > 700 {
						// 	t.log.L(`I: `, `hls切片下载慢`, usedt, `ms`)
						// }
						if len(link.data) > len(r.Respon) {
							link.data = link.data[:copy(link.data, r.Respon)]
						} else {
							// if cap(link.data) < len(r.Respon) {
							// 	fmt.Println(cap(link.data), len(r.Respon))
							// }
							link.data = append(link.data[:0], r.Respon...)
						}
						link.status = 2 // 设置切片状态为下载完成
					}
				}(v, t.Current_save_path)
			}

			// 等待队列下载完成
			download_limit.BlockAll(func() {
				time.Sleep(time.Millisecond * 10)
			})
			download_limit.UnBlockAll()
		}

		// 传递已下载切片
		for k := 0; k < len(download_seq); k++ {
			// v := download_seq[k]

			if download_seq[k].status != 2 {
				if err := download_seq[k].err; err != nil {
					t.log.L(`E: `, `切片下载发生错误:`, err)
					e = err
					return
				}
				if download_seq[k].tryDownCount >= 3 {
					//下载了2次，任未下载成功，忽略此块
					t.putM4s(download_seq[k])
					download_seq = append(download_seq[:k], download_seq[k+1:]...)
					k -= 1
					continue
				} else {
					break
				}
			}

			// no, _ := download_seq[k].getNo()
			// fmt.Println("download_seq ", no, download_seq[k].status, t.first_buf == nil)

			if strings.Contains(download_seq[k].Base, `h`) {
				if header, e := fmp4Decoder.Init_fmp4(download_seq[k].data); e != nil {
					t.log.L(`E: `, e, `重试!`)
					download_seq[k].status = 3
					break
				} else {
					for _, trak := range fmp4Decoder.traks {
						// fmt.Println(`T: `, "找到trak:", string(trak.handlerType), trak.trackID, trak.timescale)
						t.log.L(`T: `, "找到trak:", string(trak.handlerType), trak.trackID, trak.timescale)
					}
					t.first_buf = header
					if out != nil {
						out.Write(t.first_buf, true)
						out.Sync()
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

			buf.Append(download_seq[k].data)
			t.putM4s(download_seq[k])
			download_seq = append(download_seq[:k], download_seq[k+1:]...)
			k -= 1

			last_available_offset, err := fmp4Decoder.Search_stream_fmp4(buf.GetPureBuf(), fmp4KeyFrames)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					t.log.L(`E: `, err)

					// no, _ := v.getNo()
					// file.New("error/"+strconv.Itoa(no)+".m4s", 0, true).Write(buf.getCopyBuf(), true)
					// file.New("error/"+strconv.Itoa(no)+"S.m4s", 0, true).Write(v.data, true)

					if err.Error() == "未初始化traks" {
						e = err
						return
					}
					//丢弃所有数据
					buf.Reset()
				} else {
					fmp4KeyFrames.Reset()
					last_available_offset = 0
				}
			}

			// no, _ := v.getNo()
			// fmt.Println(no, "fmp4KeyFrames", fmp4KeyFrames.size(), last_available_offset, err)

			if !fmp4KeyFrames.IsEmpty() {
				fmp4KeyFramesBuf = fmp4KeyFrames.GetCopyBuf()
				fmp4KeyFrames.Reset()
				t.bootBufPush(fmp4KeyFramesBuf)
				t.Stream_msg.Push_tag(`data`, fmp4KeyFramesBuf)
				if out != nil {
					out.Write(fmp4KeyFramesBuf, true)
					out.Sync()
				}
			}

			buf.RemoveFront(last_available_offset)
		}

		// 停止录制
		if !t.Status.Islive() {
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

		// 获取解析m3u8
		var m4s_links, m3u8_addon, err = t.fetchParseM3U8()
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

		// {
		// 	if t.last_m4s != nil {
		// 		l, _ := t.last_m4s.getNo()
		// 		fmt.Println("last", l)
		// 	}
		// 	for i := 0; i < len(m4s_links); i++ {
		// 		fmt.Println(m4s_links[i].getNo())
		// 	}
		// }

		if len(m4s_links) == 0 {
			time.Sleep(time.Second)
			continue
		} else {
			// 设置最后的切片
			if t.last_m4s == nil {
				t.last_m4s = &m4s_link_item{}
			}
			for i := len(m4s_links) - 1; i >= 0; i-- {
				// fmt.Println("set last m4s", m4s_links[i].Base)
				if !m4s_links[i].isInit() && len(m4s_links[i].Base) > 0 {
					m4s_links[i].copyTo(t.last_m4s)
					break
				}
			}
		}

		// 添加新切片到下载队列
		download_seq = append(download_seq, m4s_links...)

		if !t.config.save_as_mp4 {
			// 添加m3u8字节
			file.New(t.Current_save_path+"0.m3u8.dtmp", -1, true).Write(m3u8_addon, true)
		}
	}

	// 发送空字节会导致流服务终止
	t.Stream_msg.Push_tag(`data`, []byte{})

	if !t.config.save_as_mp4 {
		// 结束
		if p.Checkfile().IsExist(t.Current_save_path + "0.m3u8.dtmp") {
			file.New(t.Current_save_path+"0.m3u8.dtmp", -1, true).Write([]byte("#EXT-X-ENDLIST"), true)
			p.FileMove(t.Current_save_path+"0.m3u8.dtmp", t.Current_save_path+"0.m3u8")
		}
	}

	return
}

func (t *M4SStream) Start() bool {
	// 清晰度-1 or 路径存在问题 不保存
	if t.config.want_qn == -1 || t.config.save_path == "" {
		return false
	}

	// 状态检测与设置
	if t.Status.Islive() {
		t.log.L(`T: `, `已存在实例`)
		return false
	}

	// 是否在直播
	F.Get(&t.common).Get(`Liveing`)
	if !t.common.Liveing {
		t.log.L(`W: `, `未直播`)
		return false
	}

	// 实例回调
	if t.Callback_start != nil {
		if e := t.Callback_start(t); e != nil {
			t.log.L(`W: `, `开始回调错误`, e.Error())
			return false
		}
	}

	t.Status = signal.Init()
	go func() {
		defer t.Status.Done()

		if t.Callback_stop != nil {
			defer t.Callback_stop(t)
		}

		t.log.L(`I: `, `初始化录制(`+strconv.Itoa(t.common.Roomid)+`)`)

		// 初始化请求池
		t.reqPool = t.common.ReqPool

		// 初始化切片消息
		t.Stream_msg = msgq.New()

		// 初始化快速启动缓冲
		if v, ok := t.common.K_v.LoadV(`直播Web缓冲长度`).(float64); ok && v != 0 {
			t.boot_buf_size = int(v)
			t.boot_buf = make([][]byte, t.boot_buf_size)
			defer func() {
				t.boot_buf = nil
			}()
		}

		// 主循环
		for t.Status.Islive() {
			// 是否在直播
			F.Get(&t.common).Get(`Liveing`)
			if !t.common.Liveing {
				t.log.L(`W: `, `未直播`)
				break
			}

			// 获取 and 检查流地址状态
			if !t.fetchCheckStream() {
				time.Sleep(time.Second * 5)
				continue
			}

			// // 设置全部服务
			// for _, v := range t.common.Live {
			// 	if url_struct, e := url.Parse(v.Url); e == nil {
			// 		t.stream_hosts.Store(url_struct.Hostname(), v.)
			// 	}
			// }

			// 保存流
			err := t.saveStream()
			if err != nil {
				t.log.L(`E: `, "saveStream:", err)
			}

			// Deprecated: 默认总是获取到可用流
			// 直播流类型故障切换
			// if v, ok := t.common.K_v.LoadV(`直播流类型故障切换`).(bool); v && ok {
			// 	if err != nil && err.Error() == "未能找到可用流服务器" {
			// 		if v, ok := t.common.K_v.LoadV(`直播流类型`).(string); ok {
			// 			switch v {
			// 			case "fmp4":
			// 				t.common.K_v.Store(`直播流类型`, `flv`)
			// 			case "flv":
			// 				t.common.K_v.Store(`直播流类型`, `hls`)
			// 			default:
			// 				t.log.L(`E: `, `未知的流类型:`+v)
			// 			}
			// 		}
			// 	}
			// }

		}

		t.log.L(`I: `, `结束录制(`+strconv.Itoa(t.common.Roomid)+`)`)
		t.exitSign.Done()
	}()
	return true
}

func (t *M4SStream) Stop() {
	if !t.Status.Islive() {
		return
	}
	t.exitSign = signal.Init()
	t.Status.Done()
	t.log.L(`I: `, `正在等待下载完成...`)
	t.exitSign.Wait()
}

// 流服务推送方法
func (t *M4SStream) Pusher(w http.ResponseWriter, r *http.Request) {
	switch t.stream_type {
	case `m3u8`:
		t.pusherM4s(w, r)
	case `mp4`:
		t.pusherM4s(w, r)
	case `flv`:
		t.pusherFlv(w, r)
	default:
		t.log.L(`W: `, `Pusher no support stream_type`)
	}
}

func (t *M4SStream) pusherM4s(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "video/mp4")

	flusher, flushSupport := w.(http.Flusher)
	if flushSupport {
		flusher.Flush()
	}

	//写入hls头
	if _, err := w.Write(t.getFirstBuf()); err != nil {
		return
	} else if flushSupport {
		flusher.Flush()
	}

	//写入快速启动缓冲
	for i := 0; i < len(t.boot_buf); i++ {
		if _, err := w.Write(t.boot_buf[i]); err != nil {
			return
		}
	}
	if len(t.boot_buf) != 0 && flushSupport {
		flusher.Flush()
	}

	cancel := make(chan struct{})

	//hls切片
	t.Stream_msg.Pull_tag(map[string]func(interface{}) bool{
		`data`: func(data interface{}) bool {
			if b, ok := data.([]byte); ok {
				if len(b) == 0 {
					close(cancel)
					return true
				}
				if _, err := w.Write(b); err != nil {
					close(cancel)
					return true
				} else if flushSupport {
					flusher.Flush()
				}
			}
			return false
		},
		`close`: func(_ interface{}) bool {
			close(cancel)
			return true
		},
	})

	<-cancel
}

func (t *M4SStream) pusherFlv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "video/x-flv")

	flusher, flushSupport := w.(http.Flusher)
	if flushSupport {
		flusher.Flush()
	}

	//写入flv头
	if _, err := w.Write(t.getFirstBuf()); err != nil {
		return
	} else if flushSupport {
		flusher.Flush()
	}

	//写入快速启动缓冲
	for i := 0; i < len(t.boot_buf); i++ {
		if _, err := w.Write(t.boot_buf[i]); err != nil {
			return
		}
	}
	if len(t.boot_buf) != 0 && flushSupport {
		flusher.Flush()
	}

	cancel := make(chan struct{})

	//flv
	t.Stream_msg.Pull_tag(map[string]func(interface{}) bool{
		`data`: func(data interface{}) bool {
			if b, ok := data.([]byte); ok {
				if len(b) == 0 {
					close(cancel)
					return true
				}
				if _, err := w.Write(b); err != nil {
					close(cancel)
					return true
				} else if flushSupport {
					flusher.Flush()
				}
			}
			return false
		},
		`close`: func(_ interface{}) bool {
			close(cancel)
			return true
		},
	})

	<-cancel
}

func (t *M4SStream) bootBufPush(buf []byte) {
	if t.boot_buf != nil {
		t.boot_buf_locker.Block()
		defer t.boot_buf_locker.UnBlock()

		if len(t.boot_buf) == t.boot_buf_size {
			for i := 0; i < len(t.boot_buf)-1; i++ {
				t.boot_buf[i] = t.boot_buf[i+1]
			}
			t.boot_buf = t.boot_buf[:len(t.boot_buf)-1]
		}
		t.boot_buf = append(t.boot_buf, buf)
	}
}
