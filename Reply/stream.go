package reply

import (
	"bytes"
	"encoding/base64"
	"errors"
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
	funcCtrl "github.com/qydysky/part/funcCtrl"
	idpool "github.com/qydysky/part/idpool"
	log "github.com/qydysky/part/log"
	msgq "github.com/qydysky/part/msgq"
	reqf "github.com/qydysky/part/reqf"
	signal "github.com/qydysky/part/signal"
	sync "github.com/qydysky/part/sync"
)

type M4SStream struct {
	Status               *signal.Signal     //IsLive()是否运行中
	exitSign             *signal.Signal     //IsLive()是否等待退出中
	log                  *log.Log_interface //日志
	config               M4SStream_Config   //配置
	stream_last_modified time.Time          //流地址更新时间
	// stream_expires       int64              //流到期时间
	stream_hosts      sync.Map   //使用的流服务器
	stream_type       string     //流类型
	Stream_msg        *msgq.Msgq //流数据消息 tag:data
	first_buf         []byte     //m4s起始块 or flv起始块
	boot_buf          [][]byte   //快速启动缓冲
	boot_buf_size     int        //快速启动缓冲长度
	boot_buf_locker   funcCtrl.BlockFunc
	last_m4s          *m4s_link_item   //最后一个切片
	common            c.Common         //通用配置副本
	Current_save_path string           //明确的直播流保存目录
	Callback_start    func(*M4SStream) //开始的回调
	Callback_stop     func(*M4SStream) //结束的回调
	reqPool           *idpool.Idpool   //请求池
}

type M4SStream_Config struct {
	save_path     string //直播流保存目录
	want_qn       int    //直播流清晰度
	want_type     string //直播流类型
	bufsize       int    //直播hls流缓冲
	banlance_host bool   //直播hls流均衡
}

type m4s_link_item struct {
	Url    string // m4s链接
	Base   string // m4s文件名
	status int    // 下载状态 0:未下载 1:正在下载 2:下载完成 3:下载失败
	data   []byte // 下载的数据
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

func (t *M4SStream) Common() c.Common {
	return t.common
}

func (t *M4SStream) LoadConfig(common c.Common, l *log.Log_interface) {
	t.common = common
	t.log = l.Base(`直播流保存`)

	//读取配置
	if path, ok := common.K_v.LoadV("直播流保存位置").(string); ok {
		if path, err := filepath.Abs(path); err == nil {
			if _, err := os.Stat(path); err != nil {
				if errors.Is(err, os.ErrNotExist) {
					if err := p.File().NewPath(path); err != nil {
						t.log.L(`E: `, `直播流保存位置错误`, err)
						return
					}
				} else {
					t.log.L(`E: `, `直播流保存位置错误`, err)
					return
				}
			}
			t.config.save_path = path + "/"
		} else {
			t.log.L(`E: `, `直播流保存位置错误`, err)
			return
		}
	}
	if v, ok := common.K_v.LoadV(`直播hls流缓冲`).(float64); ok && v > 0 {
		t.config.bufsize = int(v)
	}
	if v, ok := common.K_v.LoadV(`直播hls流均衡`).(bool); ok {
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
	if strings.Contains(t.common.Live[0], `m3u8`) {
		t.stream_type = "m3u8"
	} else if strings.Contains(t.common.Live[0], `flv`) {
		t.stream_type = "flv"
	}

	// // 保存流地址过期时间
	// if m3u8_url, err := url.Parse(t.common.Live[0]); err != nil {
	// 	t.log.L(`E: `, err.Error())
	// 	return false
	// } else {
	// 	expires, _ := strconv.Atoi(m3u8_url.Query().Get("expires"))
	// 	t.stream_expires = int64(expires)
	// }

	// 检查是否可以获取
	CookieM := make(map[string]string)
	t.common.Cookie.Range(func(k, v interface{}) bool {
		CookieM[k.(string)] = v.(string)
		return true
	})

	req := t.reqPool.Get()
	defer t.reqPool.Put(req)
	r := req.Item.(*reqf.Req)
	if e := r.Reqf(reqf.Rval{
		Url:       t.common.Live[0],
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
		return false
	} else if r.Response.StatusCode != 200 {
		t.log.L(`W: `, `live响应错误`, r.Response.Status, string(r.Respon))
		return false
	}
	return true
}

func (t *M4SStream) fetchParseM3U8() (m4s_links []*m4s_link_item, m3u8_addon []byte, e error) {
	// 请求解析m3u8内容
	for _, v := range t.common.Live {
		m3u8_url, err := url.Parse(v)
		if err != nil {
			e = err
			return
		}

		// 设置请求参数
		rval := reqf.Rval{
			Url:            m3u8_url.String(),
			Retry:          2,
			ConnectTimeout: 2000,
			ReadTimeout:    1000,
			Timeout:        2000,
			Proxy:          c.C.Proxy,
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

		// 开始请求
		req := t.reqPool.Get()
		defer t.reqPool.Put(req)
		r := req.Item.(*reqf.Req)
		if err := r.Reqf(rval); err != nil {
			e = err
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
		for _, line := range bytes.Split(m3u8_respon, []byte("\n")) {
			if len(line) == 0 {
				continue
			}

			var m4s_link = "" //切片文件名

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

			//将切片添加到返回切片数组
			m4s_links = append(m4s_links, &m4s_link_item{
				Url:  m3u8_url.ResolveReference(u).String(),
				Base: m4s_link,
			})
		}

		// 设置最后的切片
		defer func(last_m4s *m4s_link_item) {
			t.last_m4s = last_m4s
		}(m4s_links[len(m4s_links)-1])

		// 首次下载
		if t.last_m4s == nil {
			return
		}

		// 只返回新增加的
		{
			last_no, _ := t.last_m4s.getNo()
			for k, m4s_link := range m4s_links {
				// 剔除初始段
				if m4s_link.isInit() {
					m4s_links = append(m4s_links[:k], m4s_links[k+1:]...)
				}
				no, _ := m4s_link.getNo()
				if no == last_no {
					// 只返回新增加的切片
					m4s_links = m4s_links[k+1:]
					// 只返回新增加的m3u8字节
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
				}
			}
		}

		// 来到此处说明出现了丢失 尝试补充
		var (
			guess_end_no int
			current_no   int
		)
		// 找出不是初始段的第一个切片
		for _, v := range m4s_links {
			if v.isInit() {
				continue
			}
			guess_end_no, _ = v.getNo()
			break
		}
		current_no, _ = t.last_m4s.getNo()

		if guess_end_no < current_no {
			return
		}

		t.log.L(`I: `, `发现`, guess_end_no-current_no-1, `个切片遗漏，重新下载`)
		for guess_no := guess_end_no - 1; guess_no > current_no; guess_no -= 1 {
			// 补充m3u8
			m3u8_addon = append([]byte("#EXTINF:1.00\n"+strconv.Itoa(guess_no)+".m4s\n"), m3u8_addon...)

			//获取切片地址
			u, err := url.Parse("./" + strconv.Itoa(guess_no) + `.m4s`)
			if err != nil {
				e = err
				return
			}

			//将切片添加到返回切片数组前
			m4s_links = append([]*m4s_link_item{
				{
					Url:  m3u8_url.ResolveReference(u).String(),
					Base: strconv.Itoa(guess_no) + `.m4s`,
				},
			}, m4s_links...)
		}

		// 请求解析成功，退出获取循环
		break
	}

	e = nil
	return
}

func (t *M4SStream) saveStream() {
	// 设置保存路径
	t.Current_save_path = t.config.save_path + "/" + strconv.Itoa(t.common.Roomid) + "_" + time.Now().Format("2006_01_02_15-04-05-000") + `/`

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

	//开始,结束回调
	t.Callback_start(t)
	defer t.Callback_stop(t)

	// 获取流
	switch t.stream_type {
	case `m3u8`:
		t.saveStreamM4s()
	case `flv`:
		t.saveStreamFlv()
	default:
		t.log.L(`E: `, `undefind stream type`)
	}
}

func (t *M4SStream) saveStreamFlv() {
	//对每个直播流进行尝试
	for _, v := range t.common.Live {
		//结束退出
		if !t.Status.Islive() {
			break
		}

		surl, err := url.Parse(v)
		if err != nil {
			t.log.L(`E: `, err)
			break
		}

		//开始获取
		req := t.reqPool.Get()
		{
			s := signal.Init()
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

			out, err := os.Create(t.Current_save_path + `0.flv`)
			if err != nil {
				out.Close()
			}
			rc, rw := io.Pipe()
			go func() {
				var buff []byte
				var buf = make([]byte, 1<<16)
				for {
					n, e := rc.Read(buf)
					buff = append(buff, buf[:n]...)
					if n > 0 {
						front_buf, keyframe, last_avilable_offset, e := Seach_stream_tag(buff)
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
							if last_avilable_offset != 0 {
								// fmt.Println("write Sync")
								buff = buff[last_avilable_offset-1:]
								out.Sync()
							}
						}
					}
					if e != nil {
						out.Close()
						t.Stream_msg.Push_tag(`close`, nil)
						break
					}
				}

				buf = nil
				buff = nil
			}()

			CookieM := make(map[string]string)
			t.common.Cookie.Range(func(k, v interface{}) bool {
				CookieM[k.(string)] = v.(string)
				return true
			})

			if e := r.Reqf(reqf.Rval{
				Url:              surl.String(),
				SaveToPipeWriter: rw,
				NoResponse:       true,
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
			}); e != nil && !errors.Is(e, io.EOF) {
				if reqf.IsCancel(e) {
					t.log.L(`I: `, `flv下载停止`)
				} else if !reqf.IsTimeout(e) {
					t.log.L(`E: `, `flv下载失败:`, e)
				}
			}
			s.Done()
		}
		t.reqPool.Put(req)
	}
}

func (t *M4SStream) saveStreamM4s() {
	// 同时下载数限制
	var download_limit = &funcCtrl.BlockFuncN{
		Max: 3,
	}

	// 下载循环
	for download_seq := []*m4s_link_item{}; ; {

		// 存在待下载切片
		if len(download_seq) != 0 {
			// 设置限制计划
			download_limit.Plan(int64(len(download_seq)))

			// 下载切片
			for _, v := range download_seq {
				go func(link *m4s_link_item, path string) {
					defer download_limit.UnBlock()
					download_limit.Block()

					// 已下载但还未移除的切片
					if link.status == 2 {
						return
					}

					link.status = 1 // 设置切片状态为正在下载

					// 均衡负载
					if link_url, e := url.Parse(link.Url); e == nil {
						if t.stream_hosts.Len() != 1 {
							t.stream_hosts.Range(func(key, value interface{}) bool {
								// 故障转移
								if link.status == 3 && link_url.Host == key.(string) {
									return true
								}
								// 随机
								link_url.Host = key.(string)
								return false
							})
						}
						link.Url = link_url.String()
					}

					req := t.reqPool.Get()
					defer t.reqPool.Put(req)
					r := req.Item.(*reqf.Req)
					if e := r.Reqf(reqf.Rval{
						Url:            link.Url,
						SaveToPath:     path + link.Base,
						ConnectTimeout: 2000,
						ReadTimeout:    1000,
						Timeout:        2000,
						Proxy:          t.common.Proxy,
						Header: map[string]string{
							`Connection`: `close`,
						},
					}); e != nil && !errors.Is(e, io.EOF) {
						if !reqf.IsTimeout(e) {
							t.log.L(`E: `, `hls切片下载失败:`, e)
						}
						link.status = 3 // 设置切片状态为下载失败
					} else {
						if usedt := r.UsedTime.Seconds(); usedt > 700 {
							t.log.L(`I: `, `hls切片下载慢`, usedt, `ms`)
						}
						link.data = r.Respon
						link.status = 2 // 设置切片状态为下载完成
					}
				}(v, t.Current_save_path)
			}

			// 等待队列下载完成
			download_limit.PlanDone()
		}

		// 传递已下载切片
		{
			for _, v := range download_seq {
				if strings.Contains(v.Base, `h`) {
					t.first_buf = v.data
				}

				if v.status == 2 {
					download_seq = download_seq[1:]
					t.bootBufPush(v.data)
					t.Stream_msg.Push_tag(`data`, v.data)
				} else {
					break
				}
			}
		}

		// 停止录制
		if !t.Status.Islive() {
			if len(download_seq) != 0 {
				if time.Now().Unix() > t.stream_last_modified.Unix()+300 {
					t.log.L(`E: `, `切片下载超时`)
				} else {
					t.log.L(`I: `, `下载最后切片:`, len(download_seq))
					continue
				}
			}
			break
		}

		// 刷新流地址
		// 偶尔刷新后的切片编号与原来不连续，故不再提前检查，直到流获取失败再刷新
		// if time.Now().Unix()+60 > t.stream_expires {
		// 	t.stream_expires = time.Now().Add(time.Minute * 2).Unix() // 临时的流链接过期时间
		// 	go func() {
		// 		if t.fetchCheckStream() {
		// 			t.last_m4s = nil
		// 		}
		// 	}()
		// }

		// 获取解析m3u8
		var m4s_links, m3u8_addon, err = t.fetchParseM3U8()
		if err != nil {
			t.log.L(`E: `, `获取解析m3u8发生错误`, err)
			if len(download_seq) != 0 {
				continue
			}
			if !reqf.IsTimeout(err) {
				break
			}
		}
		if len(m4s_links) == 0 {
			time.Sleep(time.Second)
			continue
		}

		// 添加新切片到下载队列
		download_seq = append(download_seq, m4s_links...)

		// 添加m3u8字节
		p.File().FileWR(p.Filel{
			File:    t.Current_save_path + "0.m3u8.dtmp",
			Loc:     -1,
			Context: []interface{}{m3u8_addon},
		})
	}

	// 发送空字节会导致流服务终止
	t.Stream_msg.Push_tag(`data`, []byte{})

	// 结束
	if p.Checkfile().IsExist(t.Current_save_path + "0.m3u8.dtmp") {
		f := p.File()
		f.FileWR(p.Filel{
			File:    t.Current_save_path + "0.m3u8.dtmp",
			Loc:     -1,
			Context: []interface{}{"#EXT-X-ENDLIST"},
		})
		p.FileMove(t.Current_save_path+"0.m3u8.dtmp", t.Current_save_path+"0.m3u8")
	}
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

	t.Status = signal.Init()
	go func() {
		defer t.Status.Done()

		t.log.L(`I: `, `初始化录制(`+strconv.Itoa(t.common.Roomid)+`)`)

		// 初始化请求池
		t.reqPool = t.common.ReqPool

		// 初始化切片消息
		t.Stream_msg = msgq.New(15)

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

			// 设置均衡负载
			for _, v := range t.common.Live {
				if url_struct, e := url.Parse(v); e == nil {
					t.stream_hosts.Store(url_struct.Hostname(), nil)
				}
				if !t.config.banlance_host {
					break
				}
			}

			// 保存流
			t.saveStream()
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
	case `flv`:
		t.pusherFlv(w, r)
	default:
		t.log.L(`E: `, `no support stream_type`)
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
	if t.boot_buf != nil && len(t.boot_buf) > 0 {
		if _, err := w.Write(t.getBootBuf()); err != nil {
			return
		} else if flushSupport {
			flusher.Flush()
		}
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
		`close`: func(data interface{}) bool {
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
	if t.boot_buf != nil && len(t.boot_buf) > 0 {
		if _, err := w.Write(t.getBootBuf()); err != nil {
			return
		} else if flushSupport {
			flusher.Flush()
		}
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
		`close`: func(data interface{}) bool {
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
			t.boot_buf = t.boot_buf[1:]
		}
		t.boot_buf = append(t.boot_buf, buf)
	}
}

func (t *M4SStream) getBootBuf() (buf []byte) {
	for i := 0; i < len(t.boot_buf); i++ {
		buf = append(buf, t.boot_buf[i]...)
	}
	return buf
}
