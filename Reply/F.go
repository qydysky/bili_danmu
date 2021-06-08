package reply

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"math"
	"time"
	"os/exec"
    "path/filepath"
	"path"
    "net/http"
	"context"
	"net/url"
	"errors"
	"bytes"
	"encoding/base64"
	// "runtime"

	"golang.org/x/text/encoding/simplifiedchinese"

	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	send "github.com/qydysky/bili_danmu/Send"

	p "github.com/qydysky/part"
	funcCtrl "github.com/qydysky/part/funcCtrl"
	idpool "github.com/qydysky/part/idpool"
	msgq "github.com/qydysky/part/msgq"
	reqf "github.com/qydysky/part/reqf"
	web "github.com/qydysky/part/web"
	b "github.com/qydysky/part/buf"
	s "github.com/qydysky/part/signal"
	limit "github.com/qydysky/part/limit"

	"github.com/christopher-dG/go-obs-websocket"
)

/*
	F额外功能区
*/
var flog = c.Log.Base(`功能`)

//功能开关选取函数
func IsOn(s string) bool {
	v, ok := c.K_v.LoadV(s).(bool)
	return ok && v
}

//字符重复度检查
//a在buf中出现的字符占a的百分数
func cross(a string,buf []string) (float32) {
	var s float32
	var matched bool
	for _,v1 := range a {
		for _,v2 := range buf {
			for _,v3 := range v2 {
				if v3 == v1 {matched = true;break}
			}
			if matched {break}
		}
		if matched {s += 1}
		matched = false
	}
	return s / float32(len([]rune(a)))
}

//在a中仅出现一次出现的字符占a的百分数
func selfcross(a string) (float32) {
	buf := make(map[rune]bool)
	for _,v := range a {
		if _,ok := buf[v]; !ok {
			buf[v] = true
		}
	}
	return 1 - float32(len(buf)) / float32(len([]rune(a)))
}

//在a的每个字符串中
//出现的字符次数最多的
//占出现的字符总数的百分数
//*单字符串中的重复出现计为1次
func selfcross2(a []string) (float32, string) {
	buf := make(map[rune]float32)
	for _,v := range a {
		block := make(map[rune]bool)
		for _,v1 := range v {
			if _,ok := block[v1]; ok {continue}
			block[v1] = true
			buf[v1] += 1
		}
	}
	var (
		max float32
		maxS string
		all float32
	)
	for k,v := range buf {
		all += v
		if v > max {max = v;maxS = string(k)}
	}
	return max / all, maxS
}

//功能区
//ShowRev 显示h营收
var (
	ShowRev_old float64
	ShowRev_start bool
)

func ShowRevf(){
	if!IsOn("统计营收") {return}
	if ShowRev_start {
		c.Log.Base(`功能`).L(`I: `, fmt.Sprintf("营收 ￥%.2f",c.Rev))
		return
	}
	ShowRev_start = true
	for {
		c.Log.Base(`功能`).L(`I: `, fmt.Sprintf("营收 ￥%.2f",c.Rev))
		for c.Rev == ShowRev_old {p.Sys().Timeoutf(60)}
		ShowRev_old = c.Rev
	}
}

//Ass 弹幕转字幕
type Ass struct {
	file string//弹幕ass文件名
	startT time.Time//开始记录的基准时间
	header string//ass开头
	encoderS func(string)(string,error)//编码
}

var (
	Ass_height = 720//字幕高度
	Ass_width = 1280//字幕宽度
	Ass_font = 50//字幕字体大小
	Ass_T = 7//单条字幕显示时间
	Ass_loc = 7//字幕位置 小键盘对应的位置
)

var ass = Ass {
header:`[Script Info]
Title: Default Ass file
ScriptType: v4.00+
WrapStyle: 0
ScaledBorderAndShadow: yes
PlayResX: `+strconv.Itoa(Ass_height)+`
PlayResY: `+strconv.Itoa(Ass_width)+`

[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: Default,,`+strconv.Itoa(Ass_font)+`,&H40FFFFFF,&H000017FF,&H80000000,&H40000000,0,0,0,0,100,100,0,0,1,4,4,`+strconv.Itoa(Ass_loc)+`,20,20,50,1

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`,
encoderS:simplifiedchinese.GB18030.NewEncoder().String,
}

func init(){
	accept := map[string]bool{
		`GB18030`:true,
		`utf-8`:true,
	}
	if v,ok := c.K_v.LoadV("Ass编码").(string);ok{
		if v1,ok := accept[v];ok && v1 {
			c.Log.Base(`Ass`).L(`T: `,"编码:", v)
			if v == `utf-8` {
				ass.encoderS = func(b string)(string,error){
					return b,nil
				}
			}
		}
	}
}

//设定字幕文件名，为""时停止输出
func Ass_f(file string, st time.Time){
	ass.file = file
	if file == "" {return}

	if rel, err := filepath.Rel(savestream.base_path, ass.file);err == nil {
		c.Log.Base(`Ass`).L(`I: `,"保存到", rel + ".ass")
	} else {
		c.Log.Base(`Ass`).L(`I: `, "保存到", ass.file + ".ass")
		c.Log.Base(`Ass`).L(`W: `, err)
	}

	if tmp,err := ass.encoderS(ass.header);err != nil {
		c.Log.Base(`Ass`).L(`W: `, err)
	} else {
		p.File().FileWR(p.Filel{
			File:ass.file + ".ass",
			Write:true,
			Loc:0,
			Context:[]interface{}{tmp},
		})
		ass.startT = st
	}
}

//传入要显示的单条字幕
func Assf(s string){
	if !IsOn("生成Ass弹幕") {return}
	if ass.file == "" {return}

	if s == "" {return}

	st := time.Since(ass.startT) + time.Duration(p.Rand().MixRandom(0, 2000)) * time.Millisecond
	et := st + time.Duration(Ass_T) * time.Second

	var b string
	// b += "Comment: " + strconv.Itoa(loc) + " "+ Dtos(showedt) + "\n"
	b += `Dialogue: 0,`
	b += dtos(st) + `,` + dtos(et)
	b += `,Default,,0,0,0,,{\fad(200,500)\blur3}` + s + "\n"

	if tmp,err := ass.encoderS(b);err != nil {
		c.Log.Base(`Ass`).L(`W: `, err)
	} else {
		p.File().FileWR(p.Filel{
			File:ass.file + ".ass",
			Write:true,
			Loc:-1,
			Context:[]interface{}{tmp},
		})
	}
}

//时间转化为0:00:00.00规格字符串
func dtos(t time.Duration) string {
	M := int(math.Floor(t.Minutes())) % 60
	S := int(math.Floor(t.Seconds())) % 60
	Ns := t.Nanoseconds() / int64(time.Millisecond) % 1000 / 10

	return fmt.Sprintf("%d:%02d:%02d.%02d", int(math.Floor(t.Hours())), M, S, Ns)
}

//hls
//https://datatracker.ietf.org/doc/html/draft-pantos-http-live-streaming

//直播流保存
type Savestream struct {
	base_path string
	path string
	hls_stream struct {
		b []byte//发送给客户的m3u8字节
		t time.Time
	}
	flv_front []byte//flv头及首tag
	flv_stream *msgq.Msgq//发送给客户的flv流关键帧间隔片

	m4s_hls int//hls list 中的m4s数量
	hlsbuffersize int//hls list缓冲m4s数量
	hls_banlance_host bool//使用均衡hls服务器

	wait *s.Signal
	cancel *s.Signal
	skipFunc funcCtrl.SkipFunc
}

type hls_generate struct {
	hls_file_header []byte//发送给客户的m3u8不变头
	m4s_list []*m4s_link_item//m4s列表 缓冲
}

type m4s_link_item struct {//使用指针以设置是否已下载
	Url string// m4s链接
	Base string//m4s文件名
	Offset_line int//m3u8中的行下标
	status int//该m4s下载状态 s_noload:未下载 s_loading正在下载 s_fin下载完成 s_fail下载失败
	isshow bool
}
//m4s状态
const (
	s_noload = iota
	s_loading
	s_fin
	s_fail
)

var savestream = Savestream {
	flv_stream:msgq.New(10),//队列最多保留10个关键帧间隔片
	m4s_hls:8,
}

func init(){
	//使用带tag的消息队列在功能间传递消息
	c.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
		`savestream`:func(data interface{})(bool){
			if savestream.cancel.Islive() {
				Savestream_wait()
			} else {
				go Savestreamf()
			}

			return false
		},
	})
	//base_path
	if path,ok := c.K_v.LoadV("直播流保存位置").(string);ok{
		if path,err := filepath.Abs(path);err == nil{
			savestream.base_path = path+"/"
		}
	}
	if v, ok := c.K_v.LoadV(`直播hls流缓冲`).(float64);ok && v > 0 {
		savestream.hlsbuffersize = int(v)
	}
	if v, ok := c.K_v.LoadV(`直播hls流均衡`).(bool);ok {
		savestream.hls_banlance_host = v
	}
}

//已go func形式调用，将会获取直播流
func Savestreamf(){
	l := c.Log.Base(`savestream`)

	//避免多次开播导致的多次触发
	{
		if savestream.skipFunc.NeedSkip() {
			l.L(`T: `,`已存在实例`)
			return
		}
		defer savestream.skipFunc.UnSet()
	}

	want_qn, ok := c.K_v.LoadV("直播流清晰度").(float64)
	if !ok || want_qn < 0 {return}
	c.Live_want_qn = int(want_qn)

	F.Get(`Live`)

	if savestream.cancel.Islive() {return}

	var (
		no_found_link = errors.New("no_found_link")
		no_Modified = errors.New("no_Modified")
		last_hls_Modified time.Time
		hls_get_link = func(m3u8_url string,last_download *m4s_link_item) (need_download []*m4s_link_item,m3u8_file_addition []byte,expires int,err error) {
			url_struct,e := url.Parse(m3u8_url)
			if e != nil {
				err = e
				return
			}

			query := url_struct.Query()

			var (
				r = reqf.New()
				rval = reqf.Rval{
					Url:m3u8_url,
					Retry:0,
					ConnectTimeout:3000,
					ReadTimeout:1000,
					Timeout:5000,
					Proxy:c.Proxy,
					Header:map[string]string{
						`Host`: url_struct.Host,
						`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
						`Accept`: `*/*`,
						`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
						`Accept-Encoding`: `gzip, deflate, br`,
						`Origin`: `https://live.bilibili.com`,
						`Connection`: `keep-alive`,
						`Pragma`: `no-cache`,
						`Cache-Control`: `no-cache`,
						`Referer`:"https://live.bilibili.com/",
					},
				}
			)
			if !last_hls_Modified.IsZero() {
				rval.Header[`If-Modified-Since`] = last_hls_Modified.Add(time.Second).Format("Mon, 02 Jan 2006 15:04:05 CST")
			}
			if e := r.Reqf(rval);e != nil {
				err = e
				return
			}
			if usedt := r.UsedTime.Seconds();usedt > 3000 {
				l.L(`I: `, `hls列表下载慢`, usedt, `ms`)
			}
			if r.Response.StatusCode == http.StatusNotModified {
				l.L(`T: `, `hls未更改`)
				err = no_Modified
				return
			}
			//last_hls_Modified
			if t,ok := r.Response.Header[`Last-Modified`];ok && len(t) > 0 {
				if lm,e := time.Parse("Mon, 02 Jan 2006 15:04:05 CST", t[0]);e == nil {
					last_hls_Modified = lm
				} else {
					l.L(`T: `, e)
				}
			}

			trid := query.Get("trid")
			expires,_ = strconv.Atoi(query.Get("expires"))
			buf := r.Respon

			//base-64
			if len(buf) != 0 && !bytes.Contains(buf, []byte("#")) {
				buf,err = base64.StdEncoding.DecodeString(string(buf))
				if err != nil {
					return
				}
				// fmt.Println(`base64`)
			}

			var m4s_links []*m4s_link_item
			lines := bytes.Split(buf, []byte("\n"))
			for i:=0;i<len(lines);i+=1 {
				line := lines[i]
				m4s_link := ""

				if bytes.Contains(line, []byte("EXT-X-MAP")) {
					o := bytes.Index(line,[]byte(`EXT-X-MAP:URI="`)) + 15
					e := bytes.Index(line[o:],[]byte(`"`)) + o
					m4s_link = string(line[o:e])
				} else if bytes.Contains(line, []byte(".m4s")) {
					m4s_link = string(line)
				}

				if m4s_link == "" {continue}

				u, e := url.Parse("./"+m4s_link+"?trid="+trid)
				if e != nil {
					err = e
					return
				}
				m4s_links = append(m4s_links, &m4s_link_item{
					Url:url_struct.ResolveReference(u).String(),
					Base:m4s_link,
					Offset_line:i,
				})
			}
			if len(m4s_links) == 0 {
				err = no_found_link
				return
			}

			if last_download == nil {
				m3u8_file_addition = buf
				need_download = m4s_links
				return
			}

			var found bool
			for i:=0;i<len(m4s_links);i+=1 {
				if found {
					offset := m4s_links[i].Offset_line-1
					for i:=offset;i<len(lines);i+=1 {
						m3u8_file_addition = append(m3u8_file_addition, lines[i]...)
						m3u8_file_addition = append(m3u8_file_addition,[]byte("\n")...)
					}
					m3u8_file_addition = m3u8_file_addition[:len(m3u8_file_addition)-1]

					need_download = append(need_download, m4s_links[i:]...)
					break
				}
				found = (*last_download).Base == m4s_links[i].Base
			}
			if !found {
				offset := m4s_links[1].Offset_line-1
				for i:=offset;i<len(lines);i+=1 {
					m3u8_file_addition = append(m3u8_file_addition, lines[i]...)
					m3u8_file_addition = append(m3u8_file_addition,[]byte("\n")...)
				}
				m3u8_file_addition = m3u8_file_addition[:len(m3u8_file_addition)-1]

				need_download = append(need_download, m4s_links[1:]...)
			}

			return
		}
		flv_get_link = func(link string) (need_download string,expires int,err error) {
			need_download = link

			url_struct,e := url.Parse(link)
			if e != nil {
				err = e
				return
			}
			query := url_struct.Query()
			expires,_ = strconv.Atoi(query.Get("expires"))

			return
		}
	)

	for {
		F.Get(`Liveing`)
		if !c.Liveing {break}

		F.Get(`Live`)
		if len(c.Live)==0 {break}

		savestream.path = savestream.base_path

		savestream.path += strconv.Itoa(c.Roomid) + "_" + time.Now().Format("2006_01_02_15-04-05-000")

		savestream.wait = s.Init()
		savestream.cancel = s.Init()

		CookieM := make(map[string]string)
		c.Cookie.Range(func(k,v interface{})(bool){
			CookieM[k.(string)] = v.(string)
			return true
		})

		{//重试
			r := reqf.New()
			go func(){
				savestream.cancel.Wait()
				r.Close()
			}()
			l.L(`I: `,"尝试连接live")
			if e := r.Reqf(reqf.Rval{
				Url:c.Live[0],
				Retry:10,
				SleepTime:1000,
				Proxy:c.Proxy,
				Header:map[string]string{
					`Cookie`:reqf.Map_2_Cookies_String(CookieM),
				},
				Timeout:5*1000,
				JustResponseCode:true,
			}); e != nil{l.L(`W: `,e)}

			if r.Response == nil {
				l.L(`W: `,`live响应错误`)
				savestream.wait.Done()
				savestream.cancel.Done()
				time.Sleep(time.Second*5)
				continue
			} else if r.Response.StatusCode != 200 {
				l.L(`W: `,`live响应错误`,r.Response.Status,string(r.Respon))
				savestream.wait.Done()
				savestream.cancel.Done()
				time.Sleep(time.Second*5)
				continue
			}
		}

		if strings.Contains(c.Live[0],"flv") {
			if rel, err := filepath.Rel(savestream.base_path, savestream.path);err == nil {
				l.L(`I: `,"保存到", rel + ".flv")
			} else {
				l.L(`I: `,"保存到", savestream.path + ".flv")
				l.L(`W: `, err)
			}
			Ass_f(savestream.path, time.Now())

			// no expect qn
			exit_chan := s.Init()
			go func(){
				savestream.cancel.Wait()
				exit_chan.Done()
			}()

			type link_stream struct {
				id *idpool.Id
				front []byte
				keyframe [][]byte
				// sync_buf []byte
				close func()
			}

			//chans
			var (
				reqs = msgq.New(10)
				id_pool = idpool.New()
			)

			//文件
			out, err := os.Create(savestream.path + ".flv" + ".dtmp")
			if err != nil {
				l.L(`E: `,err)
				return
			}

			//数据整合
			{
				type id_close struct {
					id uintptr
					close func()
				}

				var (
					reqs_used_id []id_close
					reqs_remove_id []id_close

					reqs_keyframe [][][]byte

					reqs_func_block funcCtrl.BlockFunc
					last_keyframe_timestamp int
				)
				reqs.Pull_tag(map[string]func(interface{})(bool){
					`req`:func(data interface{})(bool){
						req,ok := data.(link_stream)

						if !ok {return false}

						if len(req.keyframe) == 0 {
							// fmt.Println(`没有keyframe，退出`)
							req.close()
							return false
						}
						// fmt.Println(`处理req_id`,req.id.Id,`keyframe_len`,len(req.keyframe))

						if offset,_ := out.Seek(0,1);offset == 0 {
							// fmt.Println(`添加头`,len(req.front))
							//stream
							savestream.flv_front = req.front
							out.Write(req.front)
						}

						reqs_func_block.Block()
						defer reqs_func_block.UnBlock()

						for i:=0;i<len(reqs_remove_id);i+=1 {
							if reqs_remove_id[i].id == req.id.Id {
								req.close()
								return false
							}
						}

						var reqs_keyframe_index int = len(reqs_used_id)
						{
							var isnew bool = true
							for i:=0;i<len(reqs_used_id);i+=1 {
								if reqs_used_id[i].id == req.id.Id {
									reqs_keyframe_index = i
									isnew = false
									break
								}
							}
							if isnew {
								// fmt.Println(`新req`,req.id.Id,reqs_keyframe_index)
								reqs_used_id = append(reqs_used_id, id_close{
									id:req.id.Id,
									close:req.close,
								})
							}
						}

						if len(reqs_used_id) == 1 {
							// l.L(`T: `,"单req写入",len(req.keyframe))
							last_keyframe_timestamp,_ = Keyframe_timebase(req.keyframe,last_keyframe_timestamp)

							for i:=0;i<len(req.keyframe);i+=1 {
								//stream
								savestream.flv_stream.Push_tag("stream",req.keyframe[i])
								out.Write(req.keyframe[i])
							}
							return false
						}

						for reqs_keyframe_index >= len(reqs_keyframe) {
							reqs_keyframe = append(reqs_keyframe, [][]byte{})
						}
						reqs_keyframe[reqs_keyframe_index] = append(reqs_keyframe[reqs_keyframe_index], req.keyframe...)

						// fmt.Println(`merge,添加reqs_keyframe数据`,reqs_keyframe_index,len(reqs_keyframe[reqs_keyframe_index]))

						for _,v := range reqs_keyframe {
							if len(v) == 0 {
								// fmt.Println(`merge,req无数据`,k)
								return false
							}
						}

						if success_last_keyframe_timestamp,b,merged := Merge_stream(reqs_keyframe,last_keyframe_timestamp);merged == 0 {
							// fmt.Println(`merge失败，reqs_keyframe[1]`,reqs_keyframe[1][0][:11],reqs_keyframe[1][len(reqs_keyframe[1])-1][:11])
							size := 0
							for i:=1;i<len(reqs_keyframe);i+=1 {size += len(reqs_keyframe[i])}

							if reqs_keyframe_index == 0 {
								// l.L(`T: `,"flv拼合失败，reqs_keyframe[0]写入")
								// fmt.Println(`merge失败，reqs_keyframe[0]写入`,len(req.keyframe))

								last_keyframe_timestamp,_ = Keyframe_timebase(req.keyframe,last_keyframe_timestamp)

								for i:=0;i<len(req.keyframe);i+=1 {
									//stream
									savestream.flv_stream.Push_tag("stream",req.keyframe[i])
									out.Write(req.keyframe[i])
								}
								// reqs_keyframe[0] = [][]byte{reqs_keyframe[0][len(reqs_keyframe[0])-1]}
							} else if size > 4 {
								if reqs_keyframe_index == len(reqs_used_id)-1 {
									l.L(`T: `,"flv强行拼合")

									for i:=0;i<reqs_keyframe_index;i+=1 {
										reqs_remove_id = append(reqs_remove_id, reqs_used_id[i])
										reqs_used_id[i].close()
									}
									reqs_used_id = reqs_used_id[reqs_keyframe_index:]

									last_keyframe_timestamp,_ = Keyframe_timebase(req.keyframe,last_keyframe_timestamp)

									for i:=0;i<len(req.keyframe);i+=1 {
										//stream
										savestream.flv_stream.Push_tag("stream",req.keyframe[i])
										out.Write(req.keyframe[i])
									}

									reqs_keyframe = [][][]byte{}
								} else {
									req.close()
									return false
								}
							}
						} else {
							// fmt.Println(`merge成功`,len(b))
							l.L(`T: `,"flv拼合成功")

							last_keyframe_timestamp = success_last_keyframe_timestamp

							for i:=0;i<merged;i+=1 {
								reqs_remove_id = append(reqs_remove_id, reqs_used_id[i])
								reqs_used_id[i].close()
							}
							reqs_keyframe = [][][]byte{}

							reqs_used_id = reqs_used_id[merged:]

							//stream
							savestream.flv_stream.Push_tag("stream",b)
							out.Write(b)
						}

						return false
					},
					// 11区	1
					`close`:func(data interface{})(bool){
						// defer l.L(`I: `,"处理退出")
						for i:=0;i<len(reqs_used_id);i+=1 {
							reqs_used_id[i].close()
						}
						reqs_used_id = []id_close{}
						// reqs_remove_id = []id_close{}
						reqs_keyframe = [][][]byte{}
						last_keyframe_timestamp = 0
						return true
					},
				})
			}

			//连接保持
			for {
				//随机选取服务器，获取超时时间

				live_index := 0
				if len(c.Live) > 0 {
					live_index = int(p.Rand().MixRandom(0,int64(len(c.Live)-1)))
				}
				link,exp,e := flv_get_link(c.Live[live_index])
				if e != nil {
					l.L(`W: `,`流链接获取错误`,e)
					break
				}

				// 新建chan
				var (
					bc = make(chan []byte,1<<17)
					req = reqf.New()
					req_exit = s.Init()
				)

				l.L(`I: `,`新建请求`,req.Id())

				//新建请求
				go func(r *reqf.Req,rval reqf.Rval){
					go func(){
						select {
						case <-exit_chan.WaitC():;
						case <-req_exit.WaitC():;
						}
						r.Close()
					}()
					defer req_exit.Done()
					e := r.Reqf(rval)
					if r.Response == nil {
						l.L(`W: `,`请求退出`,r.Id(),e)
					} else if r.Response.StatusCode != 200 {
						l.L(`W: `,`请求退出`,r.Id(),e,r.Response.Status,string(r.Respon))
					} else {
						l.L(`W: `,`请求退出`,r.Id(),e,r.Response.Status)
					}
				}(req,reqf.Rval{
					Url:link,
					Proxy:c.Proxy,
					Header:map[string]string{
						`Cookie`:reqf.Map_2_Cookies_String(CookieM),
					},
					//SaveToPath:savestream.path + ".flv",
					SaveToChan:bc,
					Timeout:int(int64(exp) - p.Sys().GetSTime())*1000,
					ReadTimeout:5*1000,
					ConnectTimeout:10*1000,
				})

				//返回通道
				var item = link_stream{
						close:req.Close,
						id:id_pool.Get(),
					}
				l.L(`I: `,`新建连接`,item.id.Id)

				//解析
				go func(bc chan[]byte,item *link_stream,exit_chan *s.Signal){
					var (
						buf []byte
						skip_buf_size int
					)
					defer req_exit.Done()
					defer l.L(`W: `,`连接退出`,item.id.Id)
					for exit_chan.Islive() && req_exit.Islive() {
						select {
						case <-exit_chan.WaitC():return;
						case <-req_exit.WaitC():return;
						case b :=<- bc:
							if len(b) == 0 {
								// fmt.Println(`req退出`,item.id.Id)
								id_pool.Put(item.id)
								// reqs.Push_tag(`closereq`,*item)
								return
							}

							buf = append(buf, b...)

							if len(buf) < skip_buf_size {break}

							front,list,_ := Seach_stream_tag(buf)

							if len(front) != 0 && len(item.front) == 0 {
								// fmt.Println(item.id.Id,`获取到header`,len(front))
								item.front = make([]byte,len(front))
								copy(item.front, front)
							}

							if len(list) == 0 || len(item.front) == 0 {
								// fmt.Println(`再次查询bufsize`,skip_buf_size)
								skip_buf_size = 2*len(buf)
								break
							}

							item.keyframe = list

							{
								last_keyframe := list[len(list)-1]
								cut_offset := bytes.LastIndex(buf, last_keyframe)+len(last_keyframe)
								// fmt.Printf("buf截断 当前%d=>%d 下一header %b\n",len(buf),len(buf)-cut_offset,buf[:11])
								buf = buf[cut_offset:]
							}

							skip_buf_size = len(buf)+len(list[0])
							reqs.Push_tag(`req`,*item)
						}
					}
				}(bc,&item,exit_chan)

				expires := int64(exp) - p.Sys().GetSTime()-120
				// no expect qn
				if c.Live_want_qn < c.Live_qn {
					expires = time.Now().Add(time.Minute*2).Unix()
				}

				//等待过期/退出
				{
					var exit_sign bool
					select {
					case <- req_exit.Chan:;//本次连接错误，退出重试
					case <- exit_chan.Chan://要求退出
						exit_sign = true//
					case <- time.After(time.Second*time.Duration(int(expires))):;
					}
					if exit_sign {
						//退出
						// l.L(`T: `,"chan退出")
						break
					}
				}

				l.L(`I: `,"flv关闭，开始新连接")

				//即将过期，刷新c.Live
				F.Get(`Liveing`)
				if !c.Liveing {break}
				F.Get(`Live`)
				if len(c.Live)==0 {break}
			}

			exit_chan.Done()
			reqs.Push_tag(`close`,nil)
			out.Close()

			l.L(`I: `,"结束")
			Ass_f("", time.Now())//ass
			savestream.flv_front = []byte{}//flv头及首tag置空
			p.FileMove(savestream.path+".flv.dtmp", savestream.path+".flv")
		} else {
			savestream.path += "/"
			if rel, err := filepath.Rel(savestream.base_path, savestream.path);err == nil {
				l.L(`I: `,"保存到", rel+`/0.m3u8`)
			} else {
				l.L(`I: `,"保存到", savestream.path)
				l.L(`W: `, err)
			}
			Ass_f(savestream.path+"0", time.Now())

			var (
				hls_msg = msgq.New(10)
				hls_gen hls_generate
				DISCONTINUITY int
				SEQUENCE int
			)

			//hls stream gen 用户m3u8生成
			go func(){
				per_second := time.Tick(time.Second)
				for {
					select {
					case <- savestream.cancel.WaitC():return;//exit
					case now :=<- per_second:hls_msg.Push_tag(`clock`, now);
					}
				}
			}()

			//hls stream gen 用户m3u8生成
			hls_msg.Pull_tag(map[string]func(interface{})(bool){
				`header`:func(d interface{})(bool){
					if b,ok := d.([]byte);ok {
						hls_gen.hls_file_header = b
					}
					return false
				},
				`body`:func(d interface{})(bool){
					links,ok := d.([]*m4s_link_item)
					if !ok {return false}
					//remove hls first m4s
					if len(links) > 0 &&
						len((*links[0]).Base) > 0 &&
						(*links[0]).Base[0] == 104 {links = links[1:]}

					hls_gen.m4s_list = append(hls_gen.m4s_list, links...)
					
					return false
				},
				`clock`:func(now interface{})(bool){
					//buffer
					if len(hls_gen.m4s_list) - savestream.hlsbuffersize < 0 {
						return false
					}

					var (
						res []byte
						add = int(savestream.hls_stream.t.Unix() % 3)
					)

					//add block
					{
						var m4s_num int

						//m4s list
						m4s_list_b := []byte{}
						for k,v := range hls_gen.m4s_list {
							if v.status != s_fin {
								//#EXT-X-DISCONTINUITY-SEQUENCE
								//reset hls lists
								if k == m4s_num && m4s_num < 3 {
									m4s_list := append(hls_gen.m4s_list[:k], &m4s_link_item{
										Base:"DICONTINUITY",
										status:s_fin,
										isshow:true,
									})
									hls_gen.m4s_list = append(m4s_list, hls_gen.m4s_list[k:]...)
									m4s_list_b = append(m4s_list_b, []byte("#EXT-X-DICONTINUITY\n")...)
								}
								break
							}

							v.isshow = true

							if v.Base == "DICONTINUITY" {
								m4s_list_b = append(m4s_list_b, []byte("#EXT-X-DICONTINUITY\n")...)
								continue
							}

							if m4s_num >= savestream.m4s_hls+add {break}

							m4s_num += 1
							// if m4s_num == 1 {SEQUENCE = strings.ReplaceAll(v.Base, ".m4s", "")}
							m4s_list_b = append(m4s_list_b, []byte("#EXTINF:1,"+v.Base+"\n")...)
							m4s_list_b = append(m4s_list_b, v.Base...)
							m4s_list_b = append(m4s_list_b, []byte("\n")...)
						}

						//have useable m4s
						if m4s_num != 0 {
							//add header
							res = hls_gen.hls_file_header
							//add #EXT-X-DISCONTINUITY-SEQUENCE
							res = append(res, []byte("#EXT-X-DISCONTINUITY-SEQUENCE:"+strconv.Itoa(DISCONTINUITY)+"\n")...)
							//add #EXT-X-MEDIA-SEQUENCE
							res = append(res, []byte("#EXT-X-MEDIA-SEQUENCE:"+strconv.Itoa(SEQUENCE)+"\n")...)
							//add #INFO
							res = append(res, []byte(fmt.Sprintf("#INFO-BUFFER:%d/%d\n",m4s_num,len(hls_gen.m4s_list)))...)
							//add m4s
							res = append(res, m4s_list_b...)
						}

						//去除最后一个换行
						if len(res) > 0 {res = res[:len(res)-1]}

						//设置到全局变量，方便流服务器获取
						savestream.hls_stream.b = res
					}

					//设置到全局变量，方便流服务器获取
					savestream.hls_stream.t,_ = now.(time.Time)

					//del
					if add != 2 {return false}
					for del_num:=3;del_num > 0;hls_gen.m4s_list = hls_gen.m4s_list[1:] {
						//#EXT-X-DICONTINUITY
						if hls_gen.m4s_list[0].Base == "DICONTINUITY" {
							DISCONTINUITY += 1
							continue
						}
						del_num -= 1
						//#EXTINF
						if hls_gen.m4s_list[0].isshow {SEQUENCE += 1}
					}

					return false
				},
				`close`:func(d interface{})(bool){
					savestream.hls_stream.b = []byte{}//退出置空
					savestream.hls_stream.t = time.Now()
					return true
				},
			})

			type miss_download_T struct{
				List []*m4s_link_item
				sync.RWMutex
			}
			var (
				last_download *m4s_link_item
				miss_download miss_download_T
				download_limit = funcCtrl.BlockFuncN{
					Max:2,
				}//limit
			)
			expires := time.Now().Add(time.Minute*2).Unix()

			var (
				path_front string
				path_behind string
			)

			for {
				//退出，等待下载完成
				if !savestream.cancel.Islive() {
					l.L(`I: `,"退出，等待片段下载")
					download_limit.None()
					download_limit.UnNone()

					links := []*m4s_link_item{}
					//下载出错的
					miss_download.RLock()
					if len(miss_download.List) != 0 {
						miss_download.RUnlock()
						miss_download.Lock()
						links = append(miss_download.List, links...)
						miss_download.List = []*m4s_link_item{}
						miss_download.Unlock()
					} else {
						miss_download.RUnlock()
					}

					for k,v :=range links {
						l.L(`I: `,"正在下载最后片段:",k,"/",len(links))
						v.status = s_loading
						r := reqf.New()
						if e := r.Reqf(reqf.Rval{
							Url:v.Url,
							SaveToPath:savestream.path+v.Base,
							ConnectTimeout:5000,
							ReadTimeout:1000,
							Proxy:c.Proxy,
						}); e != nil{
							l.L(`I: `,e)
							if !reqf.IsTimeout(e) {
								v.status = s_fail
							}
						} else {
							if usedt := r.UsedTime.Seconds();usedt > 700 {
								l.L(`I: `, `hls切片下载慢`, usedt, `ms`)
							}
							v.status = s_fin
						}
					}
					//退出
					break
				}

				links,file_add,exp,e := hls_get_link(c.Live[0],last_download)
				if e != nil {
					if e == no_Modified {
						time.Sleep(time.Duration(2)*time.Second)
						continue
					} else if reqf.IsTimeout(e) || strings.Contains(e.Error(), "x509") {
						l.L(`I: `,e)
						continue
					} else {
						l.L(`W: `,e)
						break
					}
				}

				//first block 获取并设置不变头
				if last_download == nil {
					var res []byte
					{
						for row,i:=bytes.SplitAfter(file_add, []byte("\n")),0;i<len(row);i+=1 {
							if bytes.Contains(row[i], []byte("#EXT")) {
								if bytes.Contains(row[i], []byte("#EXT-X-MEDIA-SEQUENCE")) || bytes.Contains(row[i], []byte("#EXTINF")){continue}
								res = append(res, row[i]...)
							}
						}
					}
					hls_msg.Push_tag(`header`, res)
				}

				if len(links) == 0 {
					time.Sleep(time.Duration(2)*time.Second)
					continue
				}

				//random host load balance
				if savestream.hls_banlance_host {
					Host_list := []string{}
					for _,v := range c.Live { 
						url_struct,e := url.Parse(v)
						if e != nil {continue}
						Host_list = append(Host_list, url_struct.Hostname())
					}
					random := int(time.Now().Unix())
					for i:=0;i<len(links);i+=1 {
						url_struct,e := url.Parse(links[i].Url)
						if e != nil {continue}
						url_struct.Host = Host_list[(random+i) % len(Host_list)]
						links[i].Url = url_struct.String()
					}
				}

				//qn in expect , set expires
				if c.Live_want_qn >= c.Live_qn {
					expires = int64(exp)
				}

				//use guess
				if last_download != nil {
					previou,_ := strconv.Atoi((*last_download).Base[:len((*last_download).Base)-4])
					now,_ := strconv.Atoi(links[0].Base[:len(links[0].Base)-4])
					if previou < now-1 {
						if diff := now - previou;diff > 100 {
							l.L(`W: `,`diff too large `,diff)
							break
						} else {
							l.L(`I: `,`猜测hls`,previou,`-`,now,`(`,diff,`)`)
						}

						{//file_add
							for i:=now-1;i>previou;i-=1 {
								file_add = append([]byte(strconv.Itoa(i)+".m4s"),file_add...)
							}
						}
						{//links
							if path_front == "" || path_behind == "" {
								u, e := url.Parse(links[0].Url)
								if e != nil {
									l.L(`E: `,`fault to enable guess`,e)
									return
								}
								path_front = u.Scheme+"://"+path.Dir(u.Host+u.Path)+"/"
								path_behind = "?"+u.RawQuery
							}

							//下载出错的
							miss_download.RLock()
							if len(miss_download.List) != 0 {
								miss_download.RUnlock()
								miss_download.Lock()
								links = append(miss_download.List, links...)
								miss_download.List = []*m4s_link_item{}
								miss_download.Unlock()
							} else {
								miss_download.RUnlock()
							}

							//出错期间没能获取到的
							for i:=now-1;i>previou;i-=1 {
								base := strconv.Itoa(i)+".m4s"
								links = append([]*m4s_link_item{
									&m4s_link_item{
										Url:path_front+base+path_behind,
										Base:base,
									},
								}, links...)
							}
						}
					}
				}

				if len(links) > 10 {
					l.L(`T: `,`等待下载切片：`,len(links))
				} else if len(links) > 100 {
					l.L(`W: `,`重试,等待下载切片：`,len(links))

					if F.Get(`Liveing`);!c.Liveing {break}
					if F.Get(`Live`);len(c.Live) == 0 {break}

					// set expect
					expires = time.Now().Add(time.Minute*2).Unix()
					continue
				}

				//将links传送给hls生成器
				hls_msg.Push_tag(`body`, links)

				f := p.File()
				f.FileWR(p.Filel{
					File:savestream.path+"0.m3u8.dtmp",
					Write:true,
					Loc:-1,
					Context:[]interface{}{file_add},
				})

				for i:=0;i<len(links);i+=1 {
					//fmp4切片下载
					go func(link *m4s_link_item,path string){
						download_limit.Block()
						defer download_limit.UnBlock()

						link.status = s_loading
						r := reqf.New()
						if e := r.Reqf(reqf.Rval{
							Url: link.Url,
							SaveToPath: path + link.Base,
							ConnectTimeout: 3000,
							ReadTimeout: 1000,
							Timeout: 3000,
							Proxy: c.Proxy,
						}); e != nil{
							if reqf.IsTimeout(e) || strings.Contains(e.Error(), "x509") {
								l.L(`I: `, link.Base, `将重试！`)
								//避免影响后续猜测
								link.Offset_line = 0
								miss_download.Lock()
								miss_download.List = append(miss_download.List, link)
								miss_download.Unlock()
							} else {
								l.L(`W: `, e)
								link.status = s_fail
							}
						} else {
							if usedt := r.UsedTime.Seconds();usedt > 700 {
								l.L(`I: `, `hls切片下载慢`, usedt, `ms`)
							}
							link.status = s_fin
							//存入cache
							if _,ok := m4s_cache.Load(path + link.Base);!ok{
								m4s_cache.Store(path + link.Base, r.Respon)
								go func(){//移除
									time.Sleep(time.Second*time.Duration(savestream.hlsbuffersize+savestream.m4s_hls+1))
									m4s_cache.Delete(path + link.Base)
								}()
							}
						}
					}(links[i],savestream.path)

					//只记录最新
					if links[i].Offset_line > 0 {
						last_download = links[i]
					}
				}

				//m3u8_url 将过期
				if p.Sys().GetSTime()+60 > expires {
					if F.Get(`Liveing`);!c.Liveing {break}
					if F.Get(`Live`);len(c.Live) == 0 {break}
					// set expect
					expires = time.Now().Add(time.Minute*2).Unix()
				} else {
					time.Sleep(time.Second)
				}
			}

			if p.Checkfile().IsExist(savestream.path+"0.m3u8.dtmp") {
				f := p.File()
				f.FileWR(p.Filel{
					File:savestream.path+"0.m3u8.dtmp",
					Write:true,
					Loc:-1,
					Context:[]interface{}{"#EXT-X-ENDLIST"},
				})
				p.FileMove(savestream.path+"0.m3u8.dtmp", savestream.path+"0.m3u8")
			}

			hls_msg.Push_tag(`close`, nil)
			l.L(`I: `,"结束")
			Ass_f("", time.Now())//ass
		}
		//set ro ``
		savestream.path = ``

		if !savestream.cancel.Islive() {
			// l.L(`I: `,"退出")
			break
		}//cancel
		/*
			Savestream需要外部组件
			ffmpeg http://ffmpeg.org/download.html
		*/
		// if p.Checkfile().IsExist(savestream.path+".flv"){
		// 	l.L(`I: `,"转码中")
		// 	p.Exec().Run(false, "ffmpeg", "-i", savestream.path+".flv", "-c", "copy", savestream.path+".mkv")
		// 	if p.Checkfile().IsExist(savestream.path+".mkv"){os.Remove(savestream.path+".flv")}
		// }

		// l.L(`I: `,"转码结束")
		savestream.wait.Done()
		savestream.cancel.Done()
	}
	savestream.wait.Done()
	savestream.cancel.Done()
}

//已func形式调用，将会停止保存直播流
func Savestream_wait(){
	if !savestream.cancel.Islive() {return}

	savestream.cancel.Done()
	c.Log.Base(`savestream`).L(`I: `,"等待停止")
	savestream.wait.Wait()
}

type Obs struct {
	c obsws.Client
	Prog string//程序路径
}

var obs = Obs {
	c:obsws.Client{Host: "127.0.0.1", Port: 4444},
	Prog:"obs",
}

func Obsf(on bool){
	if !IsOn("调用obs") {return}
	l := c.Log.Base(`obs`)

	if on {
		if p.Sys().CheckProgram("obs")[0] != 0 {l.L(`W: `,"obs已经启动");return}
		if p.Sys().CheckProgram("obs")[0] == 0 {
			if obs.Prog == "" {
				l.L(`E: `,"未知的obs程序位置")
				return
			}
			l.L(`I: `,"启动obs")
			p.Exec().Start(exec.Command(obs.Prog))
			p.Sys().Timeoutf(3)
		}

		// Connect a client.
		if err := obs.c.Connect(); err != nil {
			l.L(`E: `,err)
			return
		}
	} else {
		if p.Sys().CheckProgram("obs")[0] == 0 {l.L(`W: `,"obs未启动");return}
		obs.c.Disconnect()
	}
}

func Obs_R(on bool){
	if !IsOn("调用obs") {return}

	l := c.Log.Base("obs_R")

	if p.Sys().CheckProgram("obs")[0] == 0 {
		l.L(`W: `,"obs未启动")
		return
	} else {
		if err := obs.c.Connect(); err != nil {
			l.L(`E: `,err)
			return
		}
	}
	//录
	if on {
		req := obsws.NewStartRecordingRequest()
		if err := req.Send(obs.c); err != nil {
			l.L(`E: `,err)
			return
		}
		resp, err := req.Receive()
		if err != nil {
			l.L(`E: `,err)
			return
		}
		if resp.Status() == "ok" {
			l.L(`I: `,"开始录制")
		}
	} else {
		req := obsws.NewStopRecordingRequest()
		if err := req.Send(obs.c); err != nil {
			l.L(`E: `,err)
			return
		}
		resp, err := req.Receive()
		if err != nil {
			l.L(`E: `,err)
			return
		}
		if resp.Status() == "ok" {
			l.L(`I: `,"停止录制")
		}
		p.Sys().Timeoutf(3)
	}
}

type Autoban struct {
	Banbuf []string
	buf []string
}

var autoban = Autoban {
}

func Autobanf(s string) bool {
	if !IsOn("Autoban") {return false}

	if len(autoban.Banbuf) == 0 {
		f := p.File().FileWR(p.Filel{
			File:"Autoban.txt",
		})

		for _,v := range strings.Split(f, "\n") {
			autoban.Banbuf = append(autoban.Banbuf, v)
		}
	}

	if len(autoban.buf) < 10 {
		autoban.buf = append(autoban.buf, s)
		return false
	}
	defer func(){
		autoban.buf = append(autoban.buf[1:], s)
	}()

	var res []float32
	{
		pt := float32(len([]rune(s)))
		if pt <= 5 {return false}//字数过少去除
		res = append(res, pt)
	}
	{
		pt := selfcross(s);
		// if pt > 0.5 {return false}//自身重复高去除
		// res = append(res, pt)

		pt1 := cross(s, autoban.buf);
		if pt + pt1 > 0.3 {return false}//历史重复高去除
		res = append(res, pt, pt1)
	}
	{
		pt := cross(s, autoban.Banbuf);
		if pt < 0.8 {return false}//ban字符重复低去除
		res = append(res, pt)
	}
	l := c.Log.Base("autoban")
	l.L(`W: `,res)
	return true
}

type Danmuji struct {
	Buf map[string]string
	Inuse_auto bool
	reflect_limit *limit.Limit

	mute bool
}

var danmuji = Danmuji{
	Inuse_auto:IsOn("自动弹幕机"),
	Buf:map[string]string{
		"弹幕机在么":"在",
	},
	reflect_limit:limit.New(1,4000,8000),
}

func init(){//初始化反射型弹幕机
	buf := b.New()
	buf.Load("config/config_auto_reply.json")
	for k,v := range buf.B {
		if k == v {continue}
		danmuji.Buf[k] = v.(string)
	}
}

func Danmujif(s string) {
	if !IsOn("反射弹幕机") {return}

	if danmuji.reflect_limit.TO() {return}

	for k,v := range danmuji.Buf {
		if strings.Contains(s, k) {
			Msg_senddanmu(v)
			break
		}
	}
}

func Danmuji_auto() {
	if !IsOn("反射弹幕机") || !IsOn("自动弹幕机") || danmuji.mute {return}

	danmuji.mute = true

	var (
		list []string
		timeout int
	)
	for _,v := range c.K_v.LoadV(`自动弹幕机_内容`).([]interface{}){
		list = append(list, v.(string))
	}
	timeout = int(c.K_v.LoadV(`自动弹幕机_发送间隔s`).(float64))
	if timeout < 5 {timeout = 5}

	go func(){
		for i := 0; true; i++{
			if i >= len(list) {i = 0}
			if msg := list[i];msg != ``{Msg_senddanmu(msg)}
			p.Sys().Timeoutf(timeout)
		}
	}()
}

type Autoskip struct {
	buf map[string]Autoskip_item
	sync.Mutex
	now uint
	ticker *time.Ticker
}

type Autoskip_item struct {
	Exprie uint
	Num uint
}

var autoskip = Autoskip{
	buf:make(map[string]Autoskip_item),
	ticker:time.NewTicker(time.Duration(2)*time.Second),
}

func init(){
	go func(){
		for {
			<-autoskip.ticker.C
			if len(autoskip.buf) == 0 {continue}
			autoskip.now += 1
			autoskip.Lock()
			for k,v := range autoskip.buf{
				if v.Exprie <= autoskip.now {
					delete(autoskip.buf,k)
					{//超时显示
						if v.Num > 3 {
							Msg_showdanmu(nil, strconv.Itoa(int(v.Num)) + " x " + k,`0multi`)
						} else if v.Num > 1 {
							Msg_showdanmu(nil, strconv.Itoa(int(v.Num)) + " x " + k,`0default`)
						}
					}
				}
			}
			{//copy map
				tmp := make(map[string]Autoskip_item)
				for k,v := range autoskip.buf {tmp[k] = v}
				autoskip.buf = tmp
			}
			autoskip.Unlock()
		}
	}()
}

func Autoskipf(s string) uint {
	if !IsOn("弹幕合并") || s == ""{return 0}
	autoskip.Lock()
	defer autoskip.Unlock()
	{//验证是否已经存在
		if v,ok := autoskip.buf[s];ok && autoskip.now < v.Exprie{
			autoskip.buf[s] = Autoskip_item{
				Exprie:v.Exprie,
				Num:v.Num+1,
			}
			return v.Num
		}
	}
	{//设置
		autoskip.buf[s] = Autoskip_item{
			Exprie:autoskip.now + 8,
			Num:1,
		}
	}
	return 0
}

type Lessdanmu struct {
	buf []string
	limit *limit.Limit
	max_num int
	threshold float32

	sync.RWMutex
}

var lessdanmu = Lessdanmu{
	threshold:0.7,
}

func init() {
	if max_num,ok := c.K_v.LoadV(`每10秒显示弹幕数`).(float64);ok && int(max_num) >= 1 {
		flog.Base_add(`更少弹幕`).L(`T: `,`每10秒弹幕数:`,int(max_num))
		lessdanmu.max_num = int(max_num)
		lessdanmu.limit = limit.New(int(max_num),10000,0)
		go func(){
			//等待启动
			for lessdanmu.limit.PTK() == lessdanmu.max_num {
				time.Sleep(time.Second*3)
			}

			for {
				time.Sleep(time.Second*10)

				lessdanmu.Lock()
				if ptk := lessdanmu.limit.PTK();ptk == lessdanmu.max_num {
					if lessdanmu.threshold > 0.03 {
						lessdanmu.threshold -= 0.03
					}
				} else if ptk == 0 {
					if lessdanmu.threshold < 0.97 {
						lessdanmu.threshold += 0.03
					}
				}
				lessdanmu.Unlock()
			}
		}()
	}
}

func Lessdanmuf(s string) (show bool) {
	if !IsOn("相似弹幕忽略") {return true}
	if len(lessdanmu.buf) < 20 {
		lessdanmu.buf = append(lessdanmu.buf, s)
		return true
	}

	o := cross(s, lessdanmu.buf)
	if o == 1 {return false}//完全无用

	Jiezouf(lessdanmu.buf)
	lessdanmu.buf = append(lessdanmu.buf[1:], s)

	lessdanmu.RLock()
	show = o < lessdanmu.threshold
	lessdanmu.RUnlock()

	if show && lessdanmu.max_num > 0 {
		lessdanmu.limit.TO()
	}
	return
}

/*
	Moredanmu
	目标：弹幕机自动发送弹幕
	原理：留存弹幕，称为buf。将当前若干弹幕在buf中的位置找出，根据位置聚集情况及该位置出现语句的频率，选择发送的弹幕
*/
// type Moredanmu struct {
// 	buf []string
// }

// var moredanmu = Moredanmu{
// }
// func moredanmuf(s string) {
// 	if !moredanmu.Inuse {return}
// 	// if len(moredanmu.buf) < bufsize {
// 		moredanmu.buf = append(moredanmu.buf, s)
// 	// }

// 	// b := p.Buf("danmu.buf").Load()
// 	// if b.Get() != nil {
// 	// 	moredanmu.buf = *b.Get()
// 	// }
// }

// func moredanmu_get(tb []string) {
// 	if !moredanmu.Inuse {return}

// 	var tmp string
// 	for _,v := range tb {
// 		tmp += v
// 	}
// 	// for _,v := range tb {
// 	// 	tmp += len([]rune(v[:len(v)-1]))
// 	// }

// 	var max float32
// 	var loc int
// 	for i := 0; len(moredanmu.buf) >= i + len(tb); i++ {
// 		if m := cross(tmp, moredanmu.buf[i:i + len(tb)]);m > max {
// 			max = m
// 			loc = i
// 		}
// 	}
// 	if loc != 0 {
// 		p := moredanmu.buf[loc:loc + len(tb)]
// 		for i,v := range p{
// 			if m := cross(v, p);m > max {
// 				max = m
// 				loc = i
// 			}
// 		}
// 		fmt.Println(len(moredanmu.buf),"=>",p[loc])
// 	}
// }

type Shortdanmu struct {
	lastdanmu []rune
}

var shortdanmu = Shortdanmu{
}

func Shortdanmuf(s string) string {
	if !IsOn("精简弹幕") {return s}
	if len(shortdanmu.lastdanmu) == 0 {shortdanmu.lastdanmu = []rune(s);return s}

	var new string

	for k,v := range []rune(s) {
		if k >= len(shortdanmu.lastdanmu) {
			new += string([]rune(s)[k:])
			break
		}
		if v != shortdanmu.lastdanmu[k] {
			switch k {
			case 0, 1, 2:new = s
			default:new = "..." + string([]rune(s)[k-1:])
			}
			break
		}
	}
	// if new == "" {new = "...."}
	shortdanmu.lastdanmu = []rune(s)
	return new
}

type Jiezou struct {
	alertdanmu string
	skipS map[string]interface{}

	avg float32
	turn int
	sync.Mutex
}

var jiezou = Jiezou{
	alertdanmu:"",
	skipS:map[string]interface{}{//常见语气词忽略
		"了":nil,
		"的":nil,
		"哈":nil,
		"是":nil,
		"，":nil,
		"这":nil,
	},
}

func Jiezouf(s []string) bool {
	if !IsOn("Jiezou") {return false}
	now,S := selfcross2(s)
	jiezou.avg = (8 * jiezou.avg + 2 * now)/10
	if jiezou.turn < len(s) {jiezou.turn += 1;return false}

	if _,ok := jiezou.skipS[S]; ok {return false}

	jiezou.Lock()
	if now > 1.3 * jiezou.avg {//触发
		c.Log.Base("jiezou").L(`W: `,"节奏注意", now, jiezou.avg, S)
		jiezou.avg = now //沉默
		jiezou.Unlock()

		//发送弹幕
		if jiezou.alertdanmu != "" {Msg_senddanmu(jiezou.alertdanmu)}
		return true
	}
	jiezou.Unlock()
	return false
}

//保存所有消息到json
func init(){
	Save_to_json(0, []interface{}{`[`})
	c.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
		`change_room`:func(data interface{})(bool){//房间改变
			Save_to_json(0, []interface{}{`[`})
			return false
		},
		`flash_room`:func(data interface{})(bool){//房间改变
			Save_to_json(0, []interface{}{`[`})
			return false
		},
	})
}

func Save_to_json(Loc int,Context []interface{}) {
	if path,ok := c.K_v.LoadV(`save_to_json`).(string);ok && path != ``{
		p.File().FileWR(p.Filel{
			File:path,
			Write:true,
			Loc:int64(Loc),
			Context:Context,
		})
	}
}

//进入房间发送弹幕
func Entry_danmu(){
	flog := flog.Base_add(`进房弹幕`)

	//检查与切换粉丝牌，只在cookie存在时启用
	F.Get(`CheckSwitch_FansMedal`)

	if v,_ := c.K_v.LoadV(`进房弹幕_有粉丝牌时才发`).(bool);v && c.Wearing_FansMedal == 0{
		flog.L(`T: `,`无粉丝牌`)
		return
	}
	if v,_ := c.K_v.LoadV(`进房弹幕_仅发首日弹幕`).(bool);v {
		res := F.Get_weared_medal()
		if res.Today_intimacy > 0 {
			flog.L(`T: `,`今日已发弹幕`)
			return
		}
	}
	if array,ok := c.K_v.LoadV(`进房弹幕_内容`).([]interface{});ok && len(array) != 0{
		rand := p.Rand().MixRandom(0,int64(len(array)-1))
		send.Danmu_s(array[rand].(string), c.Roomid)
	}
}

//保持所有牌子点亮
func Keep_medal_light() {
	if v,_ := c.K_v.LoadV(`保持牌子亮着`).(bool);!v {
		return
	}
	flog := flog.Base_add(`保持亮牌`)

	array,ok := c.K_v.LoadV(`进房弹幕_内容`).([]interface{})
	if !ok || len(array) == 0{
		flog.L(`I: `,`进房弹幕_内容 为 空，退出`)
		return
	}

	flog.L(`T: `,`开始`)

	var hasKeep bool
	for _,v := range F.Get_list_in_room() {
		if t := int64(v.Last_wear_time) - time.Now().Unix();t > 60*60*24*2 || t < 0{continue}//到期时间在2天以上或已过期

		hasKeep = true

		info := F.Info(v.Target_id)
		//两天内到期，发弹幕续期
		rand := p.Rand().MixRandom(0,int64(len(array)-1))
		send.Danmu_s(array[rand].(string), info.Data.LiveRoom.Roomid)
		time.Sleep(time.Second)
	}

	//重试，使用历史弹幕
	for _,v := range F.Get_list_in_room() {
		if t := int64(v.Last_wear_time) - time.Now().Unix();t > 60*60*24*2 || t < 0{continue}//到期时间在2天以上或已过期

		info := F.Info(v.Target_id)
		//两天内到期，发弹幕续期
		var Str string
		for _,v := range F.GetHistory(info.Data.LiveRoom.Roomid).Data.Room{
			if v.Text != "" {
				Str = v.Text
				break
			}
		}
		if Str == "" {
			rand := p.Rand().MixRandom(0,int64(len(array)-1))
			Str = array[rand].(string)
		}
		send.Danmu_s(Str,info.Data.LiveRoom.Roomid)
		time.Sleep(time.Second)
	}

	if hasKeep {
		flog.L(`I: `,`完成`)
	} else {
		flog.L(`T: `,`完成`)
	}
}

//自动发送即将过期的银瓜子礼物
func AutoSend_silver_gift() {
	day,_ := c.K_v.LoadV(`发送还有几天过期的礼物`).(float64)
	if day <= 0 {
		return
	}

	flog := flog.Base_add(`自动送礼`).L(`T: `,`开始`)

	if c.UpUid == 0 {F.Get(`UpUid`)}

	var hasSend bool

	for _,v := range F.Gift_list() {
		if time.Now().Add(time.Hour * time.Duration(24 * int(day))).Unix() > int64(v.Expire_at) {
			hasSend = true
			send.Send_gift(v.Gift_id, v.Bag_id, v.Gift_num)
		}
	}

	if hasSend {
		flog.L(`I: `,`完成`)
	} else {
		flog.L(`T: `,`完成`)
	}
}

var m4s_cache sync.Map//使用内存cache避免频繁io

//直播Web服务口
func init() {
	flog := flog.Base_add(`直播Web服务`)
	if port_f,ok := c.K_v.LoadV(`直播Web服务口`).(float64);ok && port_f >= 0 {
		port := int(port_f)

		base_dir := savestream.base_path

		addr := "0.0.0.0:"
		if port == 0 {
			addr += strconv.Itoa(p.Sys().GetFreePort())
		} else {
			addr += strconv.Itoa(port)
		}

		s := web.New(&http.Server{
			Addr: addr,
		})
		var (
			root = func(w http.ResponseWriter,r *http.Request){
				//header
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Headers", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Connection", "keep-alive")
				w.Header().Set("Content-Transfer-Encoding", "binary")
				start := time.Now()
				
				var path string = r.URL.Path[1:]

				if !p.Checkfile().IsExist(base_dir+path) {
					w.WriteHeader(http.StatusNotFound)
					return
				}

				if savestream.path != "" && strings.Contains(path, filepath.Base(savestream.path)) {
					w.Header().Set("Server", "live")
					if filepath.Ext(path) == `.dtmp` {
						if strings.Contains(path,".flv") {
							// path = base_dir+path
							w.Header().Set("Content-Type", "video/x-flv")
							w.WriteHeader(http.StatusOK)

							flusher, flushSupport := w.(http.Flusher)
							if flushSupport {flusher.Flush()}

							//写入flv头，首tag
							if _,err := w.Write(savestream.flv_front);err != nil {
								return
							} else if flushSupport {
								flusher.Flush()
							}

							cancel := make(chan struct{})

							//flv流关键帧间隔切片
							savestream.flv_stream.Pull_tag(map[string]func(interface{})(bool){
								`stream`:func(data interface{})(bool){
									if b,ok := data.([]byte);ok{
										if _,err := w.Write(b);err != nil {
											close(cancel)
											return true
										} else if flushSupport {
											flusher.Flush()
										}
									}
									return false
								},
								`close`:func(data interface{})(bool){
									close(cancel)
									return true
								},
							})

							<- cancel
						} else if strings.Contains(path,".m3u8") {
							w.Header().Set("Cache-Control", "max-age=1")
							w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
							w.Header().Set("Last-Modified", savestream.hls_stream.t.Format(http.TimeFormat))

							// //经常m4s下载速度赶不上，使用阻塞避免频繁获取列表带来的卡顿
							// if time.Now().Sub(savestream.hls_stream.t).Seconds() > 1 {
							// 	time.Sleep(time.Duration(3)*time.Second)
							// }

							res := savestream.hls_stream.b

							if len(res) == 0 {
								w.Header().Set("Retry-After", "1")
								w.WriteHeader(http.StatusServiceUnavailable)
								return
							}

							//Server-Timing
							w.Header().Set("Server-Timing", fmt.Sprintf("dur=%d", time.Since(start).Microseconds()))

							if _,err := w.Write(res);err != nil {
								flog.L(`E: `,err)
								return
							}
						}
					} else if filepath.Ext(path) == `.m4s` {
						w.Header().Set("Server", "live")
						w.Header().Set("Cache-Control", "Cache-Control:public, max-age=3600")

						path = base_dir+path

						var (
							buf []byte
							cached bool
						)

						if b,ok := m4s_cache.Load(path);!ok{
							f,err := os.OpenFile(path,os.O_RDONLY,0644)
							if err != nil {
								flog.L(`E: `,err);
								return
							}
							defer f.Close()

							b := make([]byte,1<<20)
							if n,e := f.Read(b);e != nil {
								flog.L(`E: `,e)
								w.Header().Set("Retry-After", "1")
								w.WriteHeader(http.StatusServiceUnavailable)
								return
							} else if n == 1<<20 {
								flog.L(`W: `,`buf limit`)
								w.Header().Set("Retry-After", "1")
								w.WriteHeader(http.StatusServiceUnavailable)
								return
							} else {
								buf = b[:n]
								m4s_cache.Store(path,buf)
								go func(){//移除
									time.Sleep(time.Second*time.Duration(savestream.m4s_hls+1))
									m4s_cache.Delete(path)
								}()
							}
						} else {
							cached = true
							buf,_ = b.([]byte)
						}

						if len(buf) == 0 {
							flog.L(`W: `,`buf size 0`)
							w.Header().Set("Retry-After", "1")
							w.WriteHeader(http.StatusServiceUnavailable)
							return
						}

						//Server-Timing
						w.Header().Add("Server-Timing", fmt.Sprintf("cache=%v;dur=%d", cached, time.Since(start).Microseconds()))
						w.WriteHeader(http.StatusOK)
						if _,err := w.Write(buf);err != nil {
							flog.L(`E: `,err)
							return
						}
					}
				} else {
					w.Header().Set("Server", "file")
					http.FileServer(http.Dir(base_dir)).ServeHTTP(w,r)
				}
			}
			now = func(w http.ResponseWriter,r *http.Request){
				//header
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Headers", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Cache-Control", "max-age=3")

				//最新直播流
				if savestream.path == `` {
					flog.L(`T: `,`还没有下载直播流-直播流为空`)
					w.WriteHeader(http.StatusNotFound)
					return
				}

				path := filepath.Base(savestream.path)
				if strings.Contains(c.Live[0],"flv") {
					path += ".flv.dtmp"
				} else {
					path += "/0.m3u8.dtmp"
				}

				if !p.Checkfile().IsExist(base_dir+path) {
					flog.L(`T: `,`还没有下载直播流-文件未能找到`)
					w.WriteHeader(http.StatusNotFound)
				} else {
					u, e := url.Parse("../"+path)
					if e != nil {
						flog.L(`E: `,e)
						w.Header().Set("Retry-After", "1")
						w.WriteHeader(http.StatusServiceUnavailable)
						return
					}
					// r.URL =
					// root(w, r)
					w.Header().Set("Location", r.URL.ResolveReference(u).String())
					w.WriteHeader(http.StatusTemporaryRedirect)
				}
				return
			}
		)
		s.Handle(map[string]func(http.ResponseWriter,*http.Request){
			`/`:root,
			`/now`:now,
			`/exit`:func(w http.ResponseWriter,r *http.Request){
				s.Server.Shutdown(context.Background())
			},
		})
		host := p.Sys().GetIntranetIp()
		c.Stream_url = strings.Replace(`http://`+s.Server.Addr,`0.0.0.0`,host,-1)
		flog.L(`I: `,`启动于`,c.Stream_url)
	}
}