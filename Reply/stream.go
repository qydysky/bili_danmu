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
	stream_expires       int64              //流到期时间
	last_m4s             *m4s_link_item     //最后一个切片
	stream_hosts         sync.Map           //使用的流服务器
	Newst_m4s            *msgq.Msgq         //m4s消息 tag:m4s
	first_m4s            []byte             //m4s起始块
	common               c.Common           //通用配置副本
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

func (t *M4SStream) getFirstM4S() []byte {
	return t.first_m4s
}

func (t *M4SStream) fetchCheckStream() bool {
	// 获取流地址
	t.common.Live_want_qn = t.config.want_qn
	if F.Get(&t.common).Get(`Live`); len(t.common.Live) == 0 {
		return false
	}

	// 保存流地址过期时间
	if m3u8_url, err := url.Parse(t.common.Live[0]); err != nil {
		t.log.L(`E: `, err.Error())
		return false
	} else {
		expires, _ := strconv.Atoi(m3u8_url.Query().Get("expires"))
		t.stream_expires = int64(expires)
	}

	// 检查是否可以获取
	CookieM := make(map[string]string)
	t.common.Cookie.Range(func(k, v interface{}) bool {
		CookieM[k.(string)] = v.(string)
		return true
	})

	var req = reqf.New()
	if e := req.Reqf(reqf.Rval{
		Url:       t.common.Live[0],
		Retry:     10,
		SleepTime: 1000,
		Proxy:     t.common.Proxy,
		Header: map[string]string{
			`Cookie`: reqf.Map_2_Cookies_String(CookieM),
		},
		Timeout:          5 * 1000,
		JustResponseCode: true,
	}); e != nil {
		t.log.L(`W: `, e)
	}

	if req.Response == nil {
		t.log.L(`W: `, `live响应错误`)
		return false
	} else if req.Response.StatusCode != 200 {
		t.log.L(`W: `, `live响应错误`, req.Response.Status, string(req.Respon))
		return false
	}
	return true
}

func (t *M4SStream) fetchParseM3U8() (m4s_links []*m4s_link_item, m3u8_addon []byte) {
	// 请求解析m3u8内容
	for _, v := range t.common.Live {
		m3u8_url, err := url.Parse(v)
		if err != nil {
			t.log.L(`E: `, err.Error())
			return
		}

		// 设置请求参数
		rval := reqf.Rval{
			Url:            m3u8_url.String(),
			ConnectTimeout: 2000,
			ReadTimeout:    1000,
			Timeout:        2000,
			Proxy:          c.C.Proxy,
			Header: map[string]string{
				`Host`:            m3u8_url.Host,
				`User-Agent`:      `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
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
		var r = reqf.New()
		if e := r.Reqf(rval); e != nil {
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
				t.log.L(`W: `, err, string(m3u8_respon))
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
			u, e := url.Parse("./" + m4s_link + "?trid=" + m3u8_url.Query().Get("trid"))
			if e != nil {
				t.log.L(`E: `, e)
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

		if t.last_m4s == nil {
			return
		}

		// 只返回新增加的
		for k, m4s_link := range m4s_links {
			if m4s_link.Base == t.last_m4s.Base {
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
			}
		}

		// 来到此处说明出现了丢失 尝试补充
		var guess_end_no, _ = m4s_links[0].getNo()
		for no, _ := t.last_m4s.getNo(); no < guess_end_no; no += 1 {
			// 补充m3u8
			m3u8_addon = append([]byte(`#EXTINF:1.00\n`+strconv.Itoa(no)+`.m4s\n`), m3u8_addon...)

			//获取切片地址
			u, e := url.Parse("./" + strconv.Itoa(no) + `.m4s`)
			if e != nil {
				t.log.L(`E: `, e)
				return
			}

			//将切片添加到返回切片数组前
			m4s_links = append([]*m4s_link_item{
				{
					Url:  m3u8_url.ResolveReference(u).String(),
					Base: strconv.Itoa(no) + `.m4s`,
				},
			}, m4s_links...)
		}

		// 请求解析成功，退出获取循环
		break
	}

	return
}

func (t *M4SStream) saveStream() {
	// 设置保存路径
	var save_path = t.config.save_path + strconv.Itoa(t.common.Roomid) + "_" + time.Now().Format("2006_01_02_15-04-05-000") + `/`

	// 显示保存位置
	if rel, err := filepath.Rel(t.config.save_path, save_path); err == nil {
		t.log.L(`I: `, "保存到", rel+`/0.m3u8`)
	} else {
		t.log.L(`W: `, err)
	}

	// 获取流
	if strings.Contains(t.common.Live[0], `m3u8`) {
		t.stream_expires = time.Now().Add(time.Minute * 2).Unix() // 流链接过期时间

		// 同时下载数限制
		var download_limit = funcCtrl.BlockFuncN{
			Max: 3,
		}

		// 下载循环
		for download_seq := []*m4s_link_item{}; ; {
			// 下载切片
			for _, v := range download_seq {
				v.status = 1 // 设置切片状态为正在下载

				// 均衡负载
				if link_url, e := url.Parse(v.Url); e == nil {
					if t.stream_hosts.Len() != 1 {
						t.stream_hosts.Range(func(key, value interface{}) bool {
							// 故障转移
							if v.status == 3 && link_url.Host == key.(string) {
								return true
							}
							// 随机
							link_url.Host = key.(string)
							return false
						})
					}
					v.Url = link_url.String()
				}

				download_limit.Block()
				go func(link *m4s_link_item, path string) {
					defer download_limit.UnBlock()

					r := reqf.New()
					if e := r.Reqf(reqf.Rval{
						Url:            link.Url,
						SaveToPath:     path + link.Base,
						ConnectTimeout: 2000,
						ReadTimeout:    1000,
						Timeout:        2000,
						Proxy:          t.common.Proxy,
					}); e != nil && !errors.Is(e, io.EOF) {
						t.log.L(`E: `, `hls切片下载失败:`, e)
						link.status = 3 // 设置切片状态为下载失败
					} else {
						if usedt := r.UsedTime.Seconds(); usedt > 700 {
							t.log.L(`I: `, `hls切片下载慢`, usedt, `ms`)
						}
						link.data = r.Respon
						link.status = 2 // 设置切片状态为下载完成
					}
				}(v, save_path)
			}

			// 等待队列下载完成
			download_limit.None()
			download_limit.UnNone()

			//添加失败切片 传递切片
			{
				var tmp_seq []*m4s_link_item
				for _, v := range download_seq {
					if strings.Contains(v.Base, `h`) {
						t.first_m4s = v.data
					}

					if v.status == 3 {
						tmp_seq = append(tmp_seq, v)
					} else {
						t.Newst_m4s.Push_tag(`m4s`, v.data)
					}
				}
				download_seq = tmp_seq
			}

			// 停止录制
			if !t.Status.Islive() {
				if len(download_seq) != 0 {
					t.log.L(`I: `, `下载最后切片:`, len(download_seq))
					continue
				}
				break
			}

			// 刷新流地址
			if time.Now().Unix()+60 > t.stream_expires {
				t.fetchCheckStream()
			}

			// 获取解析m3u8
			var m4s_links, m3u8_addon = t.fetchParseM3U8()
			if len(m4s_links) == 0 {
				time.Sleep(time.Second)
				continue
			}

			// 添加新切片到下载队列
			download_seq = append(download_seq, m4s_links...)

			// 添加m3u8字节
			p.File().FileWR(p.Filel{
				File:    save_path + "0.m3u8.dtmp",
				Loc:     -1,
				Context: []interface{}{m3u8_addon},
			})
		}

		// 结束
		if p.Checkfile().IsExist(save_path + "0.m3u8.dtmp") {
			f := p.File()
			f.FileWR(p.Filel{
				File:    save_path + "0.m3u8.dtmp",
				Loc:     -1,
				Context: []interface{}{"#EXT-X-ENDLIST"},
			})
			p.FileMove(save_path+"0.m3u8.dtmp", save_path+"0.m3u8")
		}

	}
}

func (t *M4SStream) Start() {
	// 清晰度-1 or 路径存在问题 不保存
	if t.config.want_qn == -1 || t.config.save_path == "" {
		return
	}

	// 状态检测与设置
	if t.Status.Islive() {
		t.log.L(`T: `, `已存在实例`)
		return
	}
	t.Status = signal.Init()
	defer t.Status.Done()

	// 初始化切片消息
	t.Newst_m4s = msgq.New(10)

	// 主循环
	for t.Status.Islive() {
		// 是否在直播
		t.log.L(`I: `, t.common.Roomid)
		F.Get(&t.common).Get(`Liveing`)
		if !t.common.Liveing {
			t.log.L(`T: `, `未直播`)
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

	t.log.L(`T: `, `结束`+strconv.Itoa(t.common.Roomid))
	t.exitSign.Done()
}

func (t *M4SStream) Stop() {
	t.exitSign = signal.Init()
	t.Status.Done()
	t.log.L(`I: `, `正在等待切片下载...`)
	t.exitSign.Wait()
}
