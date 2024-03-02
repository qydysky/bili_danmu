package reply

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/http/pprof"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	psql "github.com/qydysky/part/sql"
	"golang.org/x/text/encoding/simplifiedchinese"
	_ "modernc.org/sqlite"

	"github.com/dustin/go-humanize"
	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	_ "github.com/qydysky/bili_danmu/Reply/F"
	"github.com/qydysky/bili_danmu/Reply/F/danmuXml"
	videoInfo "github.com/qydysky/bili_danmu/Reply/F/videoInfo"
	send "github.com/qydysky/bili_danmu/Send"

	p "github.com/qydysky/part"
	pctx "github.com/qydysky/part/ctx"
	file "github.com/qydysky/part/file"
	pio "github.com/qydysky/part/io"
	limit "github.com/qydysky/part/limit"
	msgq "github.com/qydysky/part/msgq"
	slice "github.com/qydysky/part/slice"
	psync "github.com/qydysky/part/sync"
	pweb "github.com/qydysky/part/web"
	websocket "github.com/qydysky/part/websocket"

	encoder "golang.org/x/text/encoding"
)

/*
F额外功能区
*/
var flog = c.C.Log.Base(`功能`)

// 功能开关选取函数
func IsOn(s string) bool {
	return c.C.IsOn(s)
}

// 字符重复度检查
// a在buf中出现的字符占a的百分数
func cross(a string, buf []string) float32 {
	var s float32
	var matched bool
	for _, v1 := range a {
		for _, v2 := range buf {
			for _, v3 := range v2 {
				if v3 == v1 {
					matched = true
					break
				}
			}
			if matched {
				break
			}
		}
		if matched {
			s += 1
		}
		matched = false
	}
	return s / float32(len([]rune(a)))
}

// 在a中仅出现一次出现的字符占a的百分数
// func selfcross(a string) float32 {
// 	buf := make(map[rune]bool)
// 	for _, v := range a {
// 		if _, ok := buf[v]; !ok {
// 			buf[v] = true
// 		}
// 	}
// 	return 1 - float32(len(buf))/float32(len([]rune(a)))
// }

// 在a的每个字符串中
// 出现的字符次数最多的
// 占出现的字符总数的百分数
// *单字符串中的重复出现计为1次
func selfcross2(a []string) (float32, string) {
	buf := make(map[rune]float32)
	for _, v := range a {
		block := make(map[rune]bool)
		for _, v1 := range v {
			if _, ok := block[v1]; ok {
				continue
			}
			block[v1] = true
			buf[v1] += 1
		}
	}
	var (
		max  float32
		maxS string
		all  float32
	)
	for k, v := range buf {
		all += v
		if v > max {
			max = v
			maxS = string(k)
		}
	}
	return max / all, maxS
}

// 功能区

// 显示营收
func init() {
	if !IsOn("统计营收") {
		return
	}
	go func() {
		var ShowRev = make(map[int]float64)

		clog := c.C.Log.Base_add(`营收`)
		for {
			if _, ok := ShowRev[c.C.Roomid]; !ok && c.C.Roomid != 0 {
				ShowRev[c.C.Roomid] = 0
			}
			for room, rev := range ShowRev {
				if c.C.Roomid != room {
					clog.L(`I: `, fmt.Sprintf("%d ￥%.2f", room, c.C.Rev))
					delete(ShowRev, room)
				} else if c.C.Rev != rev {
					ShowRev[room] = c.C.Rev
					clog.L(`I: `, fmt.Sprintf("%d ￥%.2f", room, c.C.Rev))
				}
			}
			time.Sleep(time.Minute)
		}
	}()
}

// Ass 弹幕转字幕
type Ass struct {
	file   string           //弹幕ass文件名
	startT time.Time        //开始记录的基准时间
	header string           //ass开头
	wrap   encoder.Encoding //编码
}

var (
	Ass_height = 720  //字幕高度
	Ass_width  = 1280 //字幕宽度
	Ass_font   = 50   //字幕字体大小
	Ass_T      = 7    //单条字幕显示时间
	Ass_loc    = 7    //字幕位置 小键盘对应的位置
)

var ass = Ass{
	header: `[Script Info]
Title: Default Ass file
ScriptType: v4.00+
WrapStyle: 0
ScaledBorderAndShadow: yes
PlayResX: ` + strconv.Itoa(Ass_height) + `
PlayResY: ` + strconv.Itoa(Ass_width) + `

[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: Default,,` + strconv.Itoa(Ass_font) + `,&H40FFFFFF,&H000017FF,&H80000000,&H40000000,0,0,0,0,100,100,0,0,1,4,4,` + strconv.Itoa(Ass_loc) + `,20,20,50,1

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`,
	wrap: simplifiedchinese.GB18030,
}

func init() {
	accept := map[string]bool{
		`GB18030`: true,
		`utf-8`:   true,
	}
	if v, ok := c.C.K_v.LoadV("Ass编码").(string); ok {
		if v1, ok := accept[v]; ok && v1 {
			c.C.Log.Base(`Ass`).L(`T: `, "编码:", v)
			if v == `utf-8` {
				ass.wrap = nil
			}
		}
	}
}

// 设定字幕文件名，为""时停止输出
func Ass_f(ctx context.Context, save_path string, filePath string, st time.Time) {
	if !IsOn(`仅保存当前直播间流`) {
		return
	}
	ass.file = filePath
	if filePath == "" {
		return
	}
	fl := flog.Base_add(`Ass`)
	fl.L(`I: `, `开始`)
	f := &file.File{
		Config: file.Config{
			FilePath:  ass.file + ".ass",
			AutoClose: true,
			Coder:     ass.wrap,
		},
	}
	_, _ = f.Write([]byte(ass.header), true)
	ass.startT = st

	ctx, done := pctx.WaitCtx(ctx)
	defer done()
	<-ctx.Done()

	ass.file = ""
	fl.L(`I: `, "结束")
}

// 传入要显示的单条字幕
func Assf(s string) {
	if !IsOn("生成Ass弹幕") {
		return
	}
	if ass.file == "" {
		return
	}

	if s == "" {
		return
	}

	st := time.Since(ass.startT) + time.Duration(p.Rand().MixRandom(0, 2000))*time.Millisecond
	et := st + time.Duration(Ass_T)*time.Second

	var b string
	// b += "Comment: " + strconv.Itoa(loc) + " "+ Dtos(showedt) + "\n"
	b += `Dialogue: 0,`
	b += dtos(st) + `,` + dtos(et)
	b += `,Default,,0,0,0,,{\fad(200,500)\blur3}` + s + "\n"

	f := file.New(ass.file+".ass", -1, true)
	f.Config.Coder = ass.wrap
	if _, e := f.Write([]byte(b), true); e != nil {
		flog.Base(`Assf`).L(`E: `, e)
	}
}

// 时间转化为0:00:00.00规格字符串
func dtos(t time.Duration) string {
	M := int(math.Floor(t.Minutes())) % 60
	S := int(math.Floor(t.Seconds())) % 60
	Ns := t.Nanoseconds() / int64(time.Millisecond) % 1000 / 10

	return fmt.Sprintf("%d:%02d:%02d.%02d", int(math.Floor(t.Hours())), M, S, Ns)
}

// fmp4
// https://datatracker.ietf.org/doc/html/draft-pantos-http-live-streaming
var streamO = new(sync.Map)

// 获取实例的Common
func StreamOCommon(roomid int) (array []*c.Common) {
	if roomid != -1 { //返回特定房间
		if v, ok := streamO.Load(roomid); ok {
			return []*c.Common{v.(*M4SStream).Common()}
		}
	} else { //返回所有
		streamO.Range(func(_, v interface{}) bool {
			array = append(array, v.(*M4SStream).Common())
			return true
		})
	}
	return
}

// 获取实例的录制状态
func StreamOStatus(roomid int) (Islive bool) {
	v, ok := streamO.Load(roomid)
	return ok && (!pctx.Done(v.(*M4SStream).Status) || v.(*M4SStream).exitSign.Islive())
}

// 开始实例
func StreamOStart(roomid int) {
	if StreamOStatus(roomid) {
		flog.L(`W: `, `已录制 `+strconv.Itoa(roomid)+` 不能重复录制`)
		return
	}

	var tmp = new(M4SStream)

	if e := tmp.LoadConfig(c.C.Copy()); e != nil {
		flog.L(`E: `, e)
		return
	}
	tmp.common.Roomid = roomid
	//实例回调，避免重复录制
	tmp.Callback_start = func(ms *M4SStream) error {
		//流服务添加
		if _, ok := streamO.LoadOrStore(ms.common.Roomid, tmp); ok {
			return fmt.Errorf("已存在此直播间(%d)录制", ms.common.Roomid)
		}
		return nil
	}
	tmp.Callback_stop = func(ms *M4SStream) {
		streamO.Delete(ms.common.Roomid) //流服务去除
	}
	tmp.Start()
}

// 停止实例
//
// -2 其他房间
// -1 所有房间
// 针对某房间
func StreamOStop(roomid int) {
	switch roomid {
	case -2: // 其他房间
		streamO.Range(func(_roomid, v interface{}) bool {
			if c.C.Roomid == _roomid {
				return true
			}
			if !pctx.Done(v.(*M4SStream).Status) {
				v.(*M4SStream).Stop()
			}
			streamO.Delete(_roomid)
			return true
		})
	case -1: // 所有房间
		streamO.Range(func(k, v interface{}) bool {
			if !pctx.Done(v.(*M4SStream).Status) {
				v.(*M4SStream).Stop()
			}
			streamO.Delete(k)
			return true
		})
	default: // 针对某房间
		if v, ok := streamO.Load(roomid); ok {
			if !pctx.Done(v.(*M4SStream).Status) {
				v.(*M4SStream).Stop()
			}
			streamO.Delete(roomid)
		}
	}
}

// 实例切断
func StreamOCut(roomid int) (setTitle func(string)) {
	if v, ok := streamO.Load(roomid); ok {
		if !pctx.Done(v.(*M4SStream).Status) {
			v.(*M4SStream).Cut()
			flog.L(`I: `, `已切片 `+strconv.Itoa(roomid))
			return func(title string) {
				if title != "" {
					v.(*M4SStream).Common().Title = title
				}
			}
		}
	}
	return func(s string) {}
}

// type Obs struct {
// 	c    obsws.Client
// 	Prog string //程序路径
// }

// var obs = Obs{
// 	c:    obsws.Client{Host: "127.0.0.1", Port: 4444},
// 	Prog: "obs",
// }

// func Obsf(on bool) {
// 	if !IsOn("调用obs") {
// 		return
// 	}
// 	l := c.C.Log.Base(`obs`)

// 	if on {
// 		if sys.Sys().CheckProgram("obs")[0] != 0 {
// 			l.L(`W: `, "obs已经启动")
// 			return
// 		}
// 		if sys.Sys().CheckProgram("obs")[0] == 0 {
// 			if obs.Prog == "" {
// 				l.L(`E: `, "未知的obs程序位置")
// 				return
// 			}
// 			l.L(`I: `, "启动obs")
// 			p.Exec().Start(exec.Command(obs.Prog))
// 			sys.Sys().Timeoutf(3)
// 		}

// 		// Connect a client.
// 		if err := obs.c.Connect(); err != nil {
// 			l.L(`E: `, err)
// 			return
// 		}
// 	} else {
// 		if sys.Sys().CheckProgram("obs")[0] == 0 {
// 			l.L(`W: `, "obs未启动")
// 			return
// 		}
// 		obs.c.Disconnect()
// 	}
// }

// func Obs_R(on bool) {
// 	if !IsOn("调用obs") {
// 		return
// 	}

// 	l := c.C.Log.Base("obs_R")

// 	if sys.Sys().CheckProgram("obs")[0] == 0 {
// 		l.L(`W: `, "obs未启动")
// 		return
// 	} else {
// 		if err := obs.c.Connect(); err != nil {
// 			l.L(`E: `, err)
// 			return
// 		}
// 	}
// 	//录
// 	if on {
// 		req := obsws.NewStartRecordingRequest()
// 		if err := req.Send(obs.c); err != nil {
// 			l.L(`E: `, err)
// 			return
// 		}
// 		resp, err := req.Receive()
// 		if err != nil {
// 			l.L(`E: `, err)
// 			return
// 		}
// 		if resp.Status() == "ok" {
// 			l.L(`I: `, "开始录制")
// 		}
// 	} else {
// 		req := obsws.NewStopRecordingRequest()
// 		if err := req.Send(obs.c); err != nil {
// 			l.L(`E: `, err)
// 			return
// 		}
// 		resp, err := req.Receive()
// 		if err != nil {
// 			l.L(`E: `, err)
// 			return
// 		}
// 		if resp.Status() == "ok" {
// 			l.L(`I: `, "停止录制")
// 		}
// 		sys.Sys().Timeoutf(3)
// 	}
// }

// type Autoban struct {
// 	Banbuf []string
// 	buf    []string
// }

// var autoban = Autoban{}

// func Autobanf(s string) bool {
// 	if !IsOn("Autoban") {
// 		return false
// 	}

// 	l := c.C.Log.Base("autoban")

// 	if len(autoban.Banbuf) == 0 {
// 		f := file.New("Autoban.txt", -1, false)
// 		for {
// 			if data, e := f.ReadUntil('\n', 50, 5000); e != nil {
// 				if !errors.Is(e, io.EOF) {
// 					l.L(`E: `, e)
// 				}
// 				break
// 			} else {
// 				autoban.Banbuf = append(autoban.Banbuf, string(data))
// 			}
// 		}
// 	}

// 	if len(autoban.buf) < 10 {
// 		autoban.buf = append(autoban.buf, s)
// 		return false
// 	}
// 	defer func() {
// 		autoban.buf = append(autoban.buf[1:], s)
// 	}()

// 	var res []float32
// 	{
// 		pt := float32(len([]rune(s)))
// 		if pt <= 5 {
// 			return false
// 		} //字数过少去除
// 		res = append(res, pt)
// 	}
// 	{
// 		pt := selfcross(s)
// 		// if pt > 0.5 {return false}//自身重复高去除
// 		// res = append(res, pt)

// 		pt1 := cross(s, autoban.buf)
// 		if pt+pt1 > 0.3 {
// 			return false
// 		} //历史重复高去除
// 		res = append(res, pt, pt1)
// 	}
// 	{
// 		pt := cross(s, autoban.Banbuf)
// 		if pt < 0.8 {
// 			return false
// 		} //ban字符重复低去除
// 		res = append(res, pt)
// 	}
// 	l.L(`W: `, res)
// 	return true
// }

type Danmuji struct {
	Buf           map[string]string
	Inuse_auto    bool
	reflect_limit *limit.Limit

	mute bool
}

var danmuji = Danmuji{
	Inuse_auto: IsOn("自动弹幕机"),
	Buf: map[string]string{
		"弹幕机在么": "在",
	},
	reflect_limit: limit.New(1, "4s", "8s"),
}

func init() { //初始化反射型弹幕机
	f := file.New("config/config_auto_reply.json", 0, true)
	if !f.IsExist() {
		return
	}
	bb, err := f.ReadAll(100, 1<<16)
	if !errors.Is(err, io.EOF) {
		return
	}
	var buf map[string]interface{}
	_ = json.Unmarshal(bb, &buf)
	for k, v := range buf {
		if k == v {
			continue
		}
		danmuji.Buf[k] = v.(string)
	}
}

func Danmujif(s string) {
	if !IsOn("反射弹幕机") {
		return
	}

	if danmuji.reflect_limit.TO() {
		return
	}

	for k, v := range danmuji.Buf {
		if strings.Contains(s, k) {
			Msg_senddanmu(v)
			break
		}
	}
}

func Danmuji_auto() {
	if !IsOn("反射弹幕机") || !IsOn("自动弹幕机") || danmuji.mute {
		return
	}

	danmuji.mute = true

	var (
		list    []string
		timeout int
	)
	for _, v := range c.C.K_v.LoadV(`自动弹幕机_内容`).([]interface{}) {
		list = append(list, v.(string))
	}
	timeout = int(c.C.K_v.LoadV(`自动弹幕机_发送间隔s`).(float64))
	if timeout < 5 {
		timeout = 5
	}

	go func() {
		for i := 0; true; i++ {
			if i >= len(list) {
				i = 0
			}
			if msg := list[i]; msg != `` {
				Msg_senddanmu(msg)
			}
			time.Sleep(time.Duration(timeout) * time.Second)
		}
	}()
}

type Autoskip struct {
	roomid int
	buf    map[string]Autoskip_item
	sync.Mutex
	now    uint
	ticker *time.Ticker
}

type Autoskip_item struct {
	Exprie uint
	Num    uint
}

var autoskip = Autoskip{
	buf:    make(map[string]Autoskip_item),
	ticker: time.NewTicker(time.Duration(2) * time.Second),
}

func init() {
	go func() {
		for {
			<-autoskip.ticker.C
			autoskip.Lock()
			if len(autoskip.buf) == 0 {
				autoskip.Unlock()
				continue
			}
			autoskip.now += 1
			if autoskip.roomid != c.C.Roomid {
				autoskip.buf = make(map[string]Autoskip_item)
				autoskip.roomid = c.C.Roomid
				flog.Base_add(`弹幕合并`).L(`T: `, `房间更新:`, autoskip.roomid)
				autoskip.Unlock()
				continue
			}
			for k, v := range autoskip.buf {
				if v.Exprie <= autoskip.now {
					delete(autoskip.buf, k)
					{ //超时显示
						if v.Num > 3 {
							c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{ //传入消息队列
								uid: `0multi`,
								m: map[string]string{
									`{num}`: strconv.Itoa(int(v.Num)),
									`{msg}`: k,
								},
							})
							Msg_showdanmu(Danmu_item{
								msg:    strconv.Itoa(int(v.Num)) + " x " + k,
								uid:    `0multi`,
								roomid: autoskip.roomid,
							})
						} else if v.Num > 1 {
							Msg_showdanmu(Danmu_item{
								msg:    strconv.Itoa(int(v.Num)) + " x " + k,
								uid:    `0default`,
								roomid: autoskip.roomid,
							})
						}
					}
				}
			}
			{ //copy map
				tmp := make(map[string]Autoskip_item)
				for k, v := range autoskip.buf {
					tmp[k] = v
				}
				autoskip.buf = tmp
			}
			autoskip.Unlock()
		}
	}()
}

func Autoskipf(s string) uint {
	if !IsOn("弹幕合并") || s == "" {
		return 0
	}
	autoskip.Lock()
	defer autoskip.Unlock()
	if autoskip.roomid != c.C.Roomid {
		autoskip.buf = make(map[string]Autoskip_item)
		autoskip.roomid = c.C.Roomid
		flog.Base_add(`弹幕合并`).L(`T: `, `房间更新:`, autoskip.roomid)
		return 0
	}
	{ //验证是否已经存在
		if v, ok := autoskip.buf[s]; ok && autoskip.now < v.Exprie {
			autoskip.buf[s] = Autoskip_item{
				Exprie: v.Exprie,
				Num:    v.Num + 1,
			}
			return v.Num
		}
	}
	{ //设置
		autoskip.buf[s] = Autoskip_item{
			Exprie: autoskip.now + 8,
			Num:    1,
		}
	}
	return 0
}

type Lessdanmu struct {
	roomid    int
	buf       []string
	limit     *limit.Limit
	max_num   int
	threshold float32
}

var lessdanmu = Lessdanmu{
	threshold: 0.7,
}

func init() {
	if max_num, ok := c.C.K_v.LoadV(`每秒显示弹幕数`).(float64); ok && int(max_num) >= 1 {
		flog.Base_add(`更少弹幕`).L(`T: `, `每秒弹幕数:`, int(max_num))
		lessdanmu.max_num = int(max_num)
		lessdanmu.limit = limit.New(int(max_num), "1s", "0s") //timeout right now
	}
}

func Lessdanmuf(s string) (show bool) {
	if !IsOn("相似弹幕忽略") {
		return true
	}
	if lessdanmu.roomid != c.C.Roomid {
		lessdanmu.buf = nil
		lessdanmu.roomid = c.C.Roomid
		lessdanmu.threshold = 0.7
		flog.Base_add(`更少弹幕`).L(`T: `, `房间更新:`, lessdanmu.roomid)
		return true
	}
	if len(lessdanmu.buf) < 20 {
		lessdanmu.buf = append(lessdanmu.buf, s)
		return true
	}

	o := cross(s, lessdanmu.buf)
	if o == 1 {
		return false
	} //完全无用

	Jiezouf(lessdanmu.buf)
	slice.DelFront(&lessdanmu.buf, 1)

	show = o < lessdanmu.threshold

	if show && lessdanmu.max_num > 0 {
		return !lessdanmu.limit.TO()
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
	l         sync.Mutex
}

var shortdanmu = Shortdanmu{}

func Shortdanmuf(s string) string {
	if !IsOn("精简弹幕") {
		return s
	}

	shortdanmu.l.Lock()
	defer shortdanmu.l.Unlock()

	if len(shortdanmu.lastdanmu) == 0 {
		shortdanmu.lastdanmu = []rune(s)
		return s
	}

	var new string

	for k, v := range []rune(s) {
		if k >= len(shortdanmu.lastdanmu) {
			new += string([]rune(s)[k:])
			break
		}
		if v != shortdanmu.lastdanmu[k] {
			switch k {
			case 0, 1, 2:
				new = s
			default:
				new = "..." + string([]rune(s)[k-1:])
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
	skipS      map[string]interface{}

	avg  float32
	turn int
	sync.Mutex
}

var jiezou = Jiezou{
	alertdanmu: "",
	skipS: map[string]interface{}{ //常见语气词忽略
		"了": nil,
		"的": nil,
		"哈": nil,
		"是": nil,
		"，": nil,
		"这": nil,
	},
}

func Jiezouf(s []string) bool {
	if !IsOn("Jiezou") {
		return false
	}
	now, S := selfcross2(s)
	jiezou.avg = (8*jiezou.avg + 2*now) / 10
	if jiezou.turn < len(s) {
		jiezou.turn += 1
		return false
	}

	if _, ok := jiezou.skipS[S]; ok {
		return false
	}

	jiezou.Lock()
	if now > 1.3*jiezou.avg { //触发
		c.C.Log.Base("jiezou").L(`W: `, "节奏注意", now, jiezou.avg, S)
		jiezou.avg = now //沉默
		jiezou.Unlock()

		//发送弹幕
		if jiezou.alertdanmu != "" {
			Msg_senddanmu(jiezou.alertdanmu)
		}
		return true
	}
	jiezou.Unlock()
	return false
}

// 保存所有消息到json
type saveToJson struct {
	msg  *msgq.MsgType[[]byte]
	once sync.Once
}

func (t *saveToJson) Init() {
	t.once.Do(func() {
		if path, ok := c.C.K_v.LoadV(`save_to_json`).(string); ok && path != `` {
			f := file.New(path, 0, false)
			_ = f.Delete()
			_, _ = f.Write([]byte("["), true)
			f.Close()

			t.msg = msgq.NewType[[]byte]()
			t.msg.Pull_tag(map[string]func([]byte) (disable bool){
				`data`: func(b []byte) (disable bool) {
					f := file.New(path, -1, false)
					_, _ = f.Write(b, true)
					_, _ = f.Write([]byte(","), true)
					f.Close()
					return false
				},
				`stop`: func(_ []byte) (disable bool) {
					f := file.New(path, -1, false)
					_ = f.SeekIndex(-1, file.AtEnd)
					_, _ = f.Write([]byte("]"), true)
					f.Close()
					return true
				},
			})
		}
	})
}

func (t *saveToJson) Write(data []byte) {
	if t.msg != nil {
		t.msg.PushLock_tag(`data`, data)
	}
}

func (t *saveToJson) Close() {
	if t.msg != nil {
		t.msg.PushLock_tag(`stop`, nil)
	}
}

var SaveToJson saveToJson

// 进入房间发送弹幕
func Entry_danmu() {
	flog := flog.Base_add(`进房弹幕`)

	//检查与切换粉丝牌，只在cookie存在时启用
	F.Get(c.C).Get(`CheckSwitch_FansMedal`)

	if v, _ := c.C.K_v.LoadV(`进房弹幕_有粉丝牌时才发`).(bool); v && c.C.Wearing_FansMedal == 0 {
		flog.L(`T: `, `无粉丝牌`)
		return
	}
	if v, _ := c.C.K_v.LoadV(`进房弹幕_仅发首日弹幕`).(bool); v {
		res := F.Get_weared_medal()
		if res.TodayIntimacy > 0 {
			flog.L(`T: `, `今日已发弹幕`)
			return
		}
	}
	if array, ok := c.C.K_v.LoadV(`进房弹幕_内容`).([]interface{}); ok && len(array) != 0 {
		rand := p.Rand().MixRandom(0, int64(len(array)-1))
		send.Danmu_s(array[rand].(string), c.C.Roomid)
	}
}

// 保持所有牌子点亮
func Keep_medal_light() {
	if v, _ := c.C.K_v.LoadV(`保持牌子亮着`).(bool); !v {
		return
	}
	flog := flog.Base_add(`保持亮牌`)

	array, ok := c.C.K_v.LoadV(`进房弹幕_内容`).([]interface{})
	if !ok || len(array) == 0 {
		flog.L(`I: `, `进房弹幕_内容 为 空，退出`)
		return
	}

	flog.L(`T: `, `开始`)
	defer flog.L(`I: `, `完成`)

	medals := F.Get_list_in_room()
	if len(medals) == 0 {
		return
	}
	for _, v := range medals {
		if v.IsLighted == 1 || v.RoomID == 0 {
			continue
		} //点亮状态

		//两天内到期，发弹幕续期
		rand := p.Rand().MixRandom(0, int64(len(array)-1))
		send.Danmu_s(array[rand].(string), v.RoomID)
		time.Sleep(time.Second * 5)
	}

	//重试，使用点赞
	medals = F.Get_list_in_room()
	if len(medals) == 0 {
		return
	}
	for _, v := range medals {
		if v.IsLighted == 1 || v.RoomID == 0 {
			continue
		}

		//两天内到期，发弹幕续期
		send.Danmu_s2(map[string]string{
			`msg`:     `official_147`,
			`dm_type`: `1`,
			`roomid`:  strconv.Itoa(v.RoomID),
		})
		time.Sleep(time.Second * 5)
	}

	//重试，使用历史弹幕
	medals = F.Get_list_in_room()
	if len(medals) == 0 {
		return
	}
	for _, v := range medals {
		if v.IsLighted == 1 || v.RoomID == 0 {
			continue
		}

		//两天内到期，发弹幕续期
		var Str string
		for _, v := range F.GetHistory(v.RoomID).Data.Room {
			if v.Text != "" {
				Str = v.Text
				break
			}
		}
		if Str == "" {
			rand := p.Rand().MixRandom(0, int64(len(array)-1))
			Str = array[rand].(string)
		}
		send.Danmu_s(Str, v.RoomID)
		time.Sleep(time.Second * 5)
	}
}

// 自动发送即将过期的银瓜子礼物
func AutoSend_silver_gift() {
	day, _ := c.C.K_v.LoadV(`发送还有几天过期的礼物`).(float64)
	if day <= 0 {
		return
	}

	if c.C.UpUid == 0 {
		F.Get(c.C).Get(`UpUid`)
	}

	for _, v := range F.Gift_list() {
		if time.Now().Add(time.Hour*time.Duration(24*int(day))).Unix() > int64(v.Expire_at) {
			send.Send_gift(v.Gift_id, v.Bag_id, v.Gift_num)
		}
	}

	flog.Base_add(`自动送礼`).L(`I: `, `已完成`)
}

// 直播Web服务口
var StreamWs = websocket.New_server()

func SendStreamWs(item Danmu_item) {
	var msg string
	if item.auth != nil {
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

func init() {
	flog := flog.Base_add(`直播Web服务`)
	if path, ok := c.C.K_v.LoadV(`直播Web服务路径`).(string); ok {
		if path[0] != '/' {
			flog.L(`E: `, `直播Web服务路径错误`)
			return
		}

		// debug模式
		if de, ok := c.C.K_v.LoadV(`debug模式`).(bool); ok && de {
			c.C.SerF.Store("/debug/pprof/", func(w http.ResponseWriter, r *http.Request) {
				if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
					return
				}
				pprof.Index(w, r)
			})
			c.C.SerF.Store("/debug/pprof/cmdline", func(w http.ResponseWriter, r *http.Request) {
				if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
					return
				}
				pprof.Cmdline(w, r)
			})
			c.C.SerF.Store("/debug/pprof/profile", func(w http.ResponseWriter, r *http.Request) {
				if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
					return
				}
				pprof.Profile(w, r)
			})
			c.C.SerF.Store("/debug/pprof/symbol", func(w http.ResponseWriter, r *http.Request) {
				if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
					return
				}
				pprof.Symbol(w, r)
			})
			c.C.SerF.Store("/debug/pprof/trace", func(w http.ResponseWriter, r *http.Request) {
				if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
					return
				}
				pprof.Trace(w, r)
			})
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
		c.C.SerF.Store(path, func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpCheck(c.C, w, r, http.MethodGet, http.MethodHead) {
				return
			}

			p := strings.TrimPrefix(r.URL.Path, path)

			if len(p) == 0 || p[len(p)-1] == '/' {
				p += "index.html"
			}
			if strings.HasSuffix(p, ".js") {
				w.Header().Set("content-type", "application/javascript")
			} else if strings.HasSuffix(p, ".css") {
				w.Header().Set("content-type", "text/css")
			} else if strings.HasSuffix(p, ".html") {
				w.Header().Set("content-type", "text/html")
			}

			//cache
			if bp, ok := cache.IsCache("html/streamList/" + p); ok {
				w.Header().Set("Cache-Control", "max-age=60")
				_, _ = w.Write(*bp)
				return
			}
			w = cache.Cache("html/streamList/"+p, time.Minute, w)

			f := file.New("html/streamList/"+p, 0, true)
			if !f.IsExist() || f.IsDir() {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			b, _ := f.ReadAll(humanize.KByte, humanize.MByte)
			_, _ = w.Write(b)
		})

		// 直播流文件列表api
		c.C.SerF.Store(path+"filePath", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
				return
			}

			//cache
			if bp, ok := cache.IsCache(path + "filePath"); ok {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Cache-Control", "max-age=5")
				_, _ = w.Write(*bp)
				return
			}
			w = cache.Cache(path+"filePath", time.Second*5, w)

			var filePaths []*videoInfo.Paf

			// 获取当前房间的
			var currentStreamO *M4SStream
			streamO.Range(func(key, value interface{}) bool {
				if key != nil && c.C.Roomid == key.(int) {
					currentStreamO = value.(*M4SStream)
					return false
				}
				return true
			})
			// if currentStreamO != nil && currentStreamO.Common().Liveing {
			// 	filePaths = append(filePaths, struct{
			// 		Name:       "Now: " + currentStreamO.Common().Title,
			// 		Path:       "now",
			// 		Qn:         c.C.Qn[currentStreamO.Common().Live_qn],
			// 		Uname:      currentStreamO.Common().Uname,
			// 		Format:     currentStreamO.stream_type,
			// 		StartT:     time.Now().Format(time.DateTime),
			// 		StartLiveT: currentStreamO.Common().Live_Start_Time.Format(time.DateTime),
			// 	})
			// }

			if v, ok := c.C.K_v.LoadV(`直播流保存位置`).(string); ok && v != "" {
				dir := file.New(v, 0, true)
				defer dir.Close()
				if !dir.IsDir() {
					c.ResStruct{Code: -1, Message: "not dir", Data: nil}.Write(w)
					return
				}

				if fs, e := dir.DirFiles(); e != nil {
					c.ResStruct{Code: -1, Message: e.Error(), Data: nil}.Write(w)
					return
				} else {
					for i, n := 0, len(fs); i < n; i++ {
						if filePath, e := videoInfo.Get.Run(context.Background(), fs[i]); e != nil {
							flog.L(`W: `, fs[i], e)
							continue
						} else {
							if t, e := time.Parse("2006_01_02-15_04_05", filePath.StartT); e == nil {
								filePath.StartT = t.Format(time.DateTime)
							}
							if currentStreamO != nil &&
								currentStreamO.Common().Liveing &&
								strings.Contains(currentStreamO.GetSavePath(), filePath.Path) {
								filePath.Name = "Now:" + filePath.Name
								filePath.Path = "now"
							}
							filePaths = append(filePaths, filePath)
						}
					}
				}
				sort.Slice(filePaths, func(i, j int) bool {
					return filePaths[i].StartT > filePaths[j].StartT
				})
			} else if len(filePaths) == 0 {
				c.ResStruct{Code: -1, Message: "直播流保存位置无效", Data: nil}.Write(w)
				flog.L(`W: `, `直播流保存位置无效`)
			}
			c.ResStruct{Code: 0, Message: "ok", Data: filePaths}.Write(w)
		})

		// 直播流播放器
		c.C.SerF.Store(path+"player/", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
				return
			}

			p := strings.TrimPrefix(r.URL.Path, path+"player/")
			if len(p) == 0 || p[len(p)-1] == '/' {
				p += "index.html"
			}

			if strings.HasSuffix(p, ".js") {
				w.Header().Set("content-type", "application/javascript")
			} else if strings.HasSuffix(p, ".css") {
				w.Header().Set("content-type", "text/css")
			} else if strings.HasSuffix(p, ".html") {
				w.Header().Set("content-type", "text/html")
			}

			//cache
			if bp, ok := cache.IsCache("html/artPlayer/" + p); ok {
				w.Header().Set("Cache-Control", "max-age=60")
				_, _ = w.Write(*bp)
				return
			}
			w = cache.Cache("html/artPlayer/"+p, time.Minute, w)

			f := file.New("html/artPlayer/"+p, 0, true)
			if !f.IsExist() {
				w.WriteHeader(http.StatusNotFound)
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

		c.C.SerF.Store(path+"keepAlive", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
				return
			}
			if key, e := expirer.Reg(time.Second*30, r.URL.Query().Get("key")); e != nil {
				w.WriteHeader(http.StatusForbidden)
			} else {
				_, _ = w.Write([]byte(key))
			}
		})

		// 流地址
		c.C.SerF.Store(path+"stream", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
				return
			}

			// 直播流回放连接限制
			if climit.AddCount(r) {
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}

			if e := expirer.LoopCheck(r.Context(), r.URL.Query().Get("key"), func(key string, e error) {
				_ = c.C.SerF.GetConn(r).Close()
			}); e != nil {
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}

			//header
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Connection", "keep-alive")
			w.Header().Set("Content-Transfer-Encoding", "binary")

			var rpath string

			if qref := r.URL.Query().Get("ref"); rpath == "" && qref != "" {
				rpath = "/" + qref + "/"
			}

			if rpath == "" {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if rpath != `/now/` {
				if v, ok := c.C.K_v.LoadV(`直播流保存位置`).(string); !ok || v == "" {
					w.Header().Set("Retry-After", "1")
					w.WriteHeader(http.StatusServiceUnavailable)
					flog.L(`W: `, `直播流保存位置无效`)
					return
				} else {
					if strings.HasSuffix(v, "/") || strings.HasSuffix(v, "\\") {
						v += rpath[1:]
					} else {
						v += rpath
					}
					if rawPath, e := url.PathUnescape(v); e != nil {
						w.WriteHeader(http.StatusServiceUnavailable)
						flog.L(`I: `, "路径解码失败", v)
						return
					} else {
						v = rawPath
					}
					if file.New(v+"0.flv", 0, true).IsExist() {
						v += "0.flv"
						w.Header().Set("Content-Type", "flv-application/octet-stream")
					} else if file.New(v+"0.mp4", 0, true).IsExist() {
						v += "0.mp4"
						w.Header().Set("Content-Type", "video/mp4")
					} else {
						w.Header().Set("Retry-After", "1")
						w.WriteHeader(http.StatusServiceUnavailable)
						flog.L(`I: `, "未找到流文件", v)
						return
					}

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

					f := file.New(v, int64(rangeHeaderNum), false)
					defer f.Close()

					// 设置当前返回区间，并拷贝
					// if fi, e := f.Stat(); e != nil {
					// 	w.WriteHeader(http.StatusServiceUnavailable)
					// 	flog.L(`W: `, e)
					// 	return
					// } else {
					// 	w.Header().Add(`Content-Range`, fmt.Sprintf("bytes %d-%d/%d", rangeHeaderNum, fi.Size()-1, fi.Size()))
					// 	w.WriteHeader(http.StatusPartialContent)

					flog.L(`T: `, r.RemoteAddr, `接入录播`)
					ts := time.Now()
					defer func() { flog.L(`T: `, r.RemoteAddr, `断开录播`, time.Since(ts)) }()

					if e := f.CopyToIoWriter(w, pio.CopyConfig{BytePerSec: speed}); e != nil {
						flog.L(`E: `, e)
					}
					// }
				}
				return
			}

			// 获取当前房间的
			var currentStreamO *M4SStream
			streamO.Range(func(key, value interface{}) bool {
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

				if e := currentStreamO.PusherToHttp(conn, w, r, startFunc, stopFunc); e != nil {
					flog.L(`W: `, e)
				}
			}
		})

		// 弹幕回放
		c.C.SerF.Store(path+"player/ws", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
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
				if v, ok := c.C.K_v.LoadV(`直播流保存位置`).(string); ok && v != "" {
					if strings.HasSuffix(v, "/") || strings.HasSuffix(v, "\\") {
						v += rpath[1:]
					} else {
						v += rpath
					}

					if !file.New(v+"0.csv", 0, true).IsExist() {
						w.WriteHeader(http.StatusNotFound)
						return
					} else if !file.New(v+"0.xml", 0, true).IsExist() {
						if _, e := danmuXml.DanmuXml.Run(context.Background(), &v); e != nil {
							msglog.L(`E: `, e)
						}
					}

					if s, closeF := PlayRecDanmu(v + "0.csv"); s == nil {
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
				} else {
					w.Header().Set("Retry-After", "1")
					w.WriteHeader(http.StatusServiceUnavailable)
					flog.L(`W: `, `直播流保存位置无效`)
				}
				return
			} else if IsOn("直播Web可以发送弹幕") {
				StreamWs.Interface().Pull_tag(map[string](func(interface{}) bool){
					`recv`: func(i interface{}) bool {
						if u, ok := i.(websocket.Uinterface); ok {
							if bytes.Equal(u.Data[:2], []byte("%S")) && len(u.Data) > 0 {
								flog.Base_add(`流服务弹幕`).L(`I: `, string(u.Data[2:]))
								Msg_senddanmu(string(u.Data[2:]))
							}
						}
						return false
					},
					`close`: func(i interface{}) bool { return true },
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
		c.C.SerF.Store(path+"player/xml", func(w http.ResponseWriter, r *http.Request) {
			if c.DefaultHttpCheck(c.C, w, r, http.MethodGet) {
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

			if v, ok := c.C.K_v.LoadV(`直播流保存位置`).(string); !ok || v == "" {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusServiceUnavailable)
				flog.L(`W: `, `直播流保存位置无效`)
			} else {
				if strings.HasSuffix(v, "/") || strings.HasSuffix(v, "\\") {
					v += rpath[1:]
				} else {
					v += rpath
				}

				if !file.New(v+"0.xml", 0, true).IsExist() {
					if !file.New(v+"0.csv", 0, true).IsExist() {
						w.WriteHeader(http.StatusNotFound)
						return
					}
					if _, e := danmuXml.DanmuXml.Run(context.Background(), &v); e != nil {
						msglog.L(`E: `, e)
					}
				}

				if e := file.New(v+"0.xml", 0, true).CopyToIoWriter(w, pio.CopyConfig{}); e != nil {
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
func StartRecDanmu(ctx context.Context, filePath string) {
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

	Recoder.Stop()
}

func PlayRecDanmu(filePath string) (*websocket.Server, func()) {
	if !IsOn(`仅保存当前直播间流`) || !IsOn("弹幕回放") {
		return nil, nil
	}
	return websocket.Play(filePath)
}

// 此次直播的交互人数
var communicate Communicate

type Communicate struct {
	Buf *psync.Map
}

func init() {
	communicate.Buf = new(psync.Map)
	c.C.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
		`change_room`: func(_ interface{}) bool { //房间改变
			communicate.Reset()
			return false
		},
		`flash_room`: func(_ interface{}) bool { //房间改变
			communicate.Reset()
			return false
		},
	})
}

func (t *Communicate) Reset() {
	t.Buf.Range(func(key, _ interface{}) bool {
		t.Buf.Delete(key)
		return true
	})
}

func (t *Communicate) Count() int {
	return t.Buf.Len()
}

func (t *Communicate) Store(k interface{}) {
	t.Buf.Store(k, nil)
}

// 保存弹幕至db
var saveDanmuToDB SaveDanmuToDB

type SaveDanmuToDB struct {
	dbname string
	db     *sql.DB
	insert string
	sync.Once
}

func (t *SaveDanmuToDB) init(c *c.Common) {
	t.Do(func() {
		if v, ok := c.K_v.LoadV(`保存弹幕至db`).(map[string]any); ok && len(v) != 0 {
			var (
				dbname, url, create                 string
				dbnameok, urlok, createok, insertok bool
			)

			dbname, dbnameok = v["dbname"].(string)
			url, urlok = v["url"].(string)
			create, createok = v["create"].(string)
			t.insert, insertok = v["insert"].(string)

			if dbname == "" || url == "" || t.insert == "" || !dbnameok || !urlok || !insertok {
				return
			}

			t.dbname = dbname

			if db, e := sql.Open(dbname, url); e != nil {
				c.Log.Base_add("保存弹幕至db").L(`E: `, e)
			} else {
				db.SetConnMaxLifetime(time.Minute * 3)
				db.SetMaxOpenConns(10)
				db.SetMaxIdleConns(10)
				t.db = db
				if createok {
					tx := psql.BeginTx[any](db, context.Background())
					tx.Do(psql.SqlFunc[any]{Query: create, SkipSqlErr: true})
					if _, e := tx.Fin(); e != nil {
						c.Log.Base_add("保存弹幕至db").L(`E: `, e)
						return
					}
				}
				c.Log.Base_add("保存弹幕至db").L(`I: `, dbname)
			}
		}
	})
}

func (t *SaveDanmuToDB) danmu(item Danmu_item) {
	if t.db == nil {
		return
	}
	if e := t.db.Ping(); e == nil {
		type DanmuI struct {
			Date   string
			Unix   int64
			Msg    string
			Color  string
			Auth   any
			Uid    string
			Roomid int64
		}

		var replaceF []func(index int, holder string) (replaceTo string)
		if t.dbname == "postgres" {
			replaceF = append(replaceF, func(index int, holder string) (replaceTo string) {
				return fmt.Sprintf("$%d", index+1)
			})
		}

		tx := psql.BeginTx[any](t.db, context.Background())
		tx.DoPlaceHolder(psql.SqlFunc[any]{Query: t.insert}, &DanmuI{
			Date:   time.Now().Format(time.DateTime),
			Unix:   time.Now().Unix(),
			Msg:    item.msg,
			Color:  item.color,
			Auth:   item.auth,
			Uid:    item.uid,
			Roomid: int64(item.roomid),
		}, replaceF...)
		tx.AfterEF(func(_ *any, result sql.Result, e *error) {
			if v, err := result.RowsAffected(); err != nil {
				*e = err
				return
			} else if v != 1 {
				*e = errors.New("插入数量错误")
				return
			}
		})
		if _, e := tx.Fin(); e != nil {
			c.C.Log.Base_add("保存弹幕至db").L(`E: `, e)
		}
	}
}
