package reply

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"math"
	"time"
	"os/exec"

	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	"github.com/christopher-dG/go-obs-websocket"
	p "github.com/qydysky/part"
	s "github.com/qydysky/part/buf"
)

/*
	F额外功能区
*/

//功能开关
var AllF = map[string]bool{
	`ShowRev`:true,//显示本次营收
	"Gtk":false,//Gtk弹幕窗口
	"Saveflv":true,//保存直播流(默认高清，有cookie默认蓝光)
	"Ass":true,//Ass弹幕生成，由于时间对应关系,仅开启流保存时生效
	"Obs":false,//obs组件(仅录播)
	/*
		Obs需要外部组件:
		obs https://obsproject.com/download
		obs-websocket https://github.com/Palakis/obs-websocket/releases
	*/
	"Autoban":false,//自动封禁(仅提示，未完成)
	"Jiezou":true,//带节奏预警，提示弹幕礼仪
	"Danmuji":true,//反射型弹幕机，回应弹幕
	"Danmuji_auto":false,//自动型弹幕机，定时输出
	"Autoskip":true,//刷屏缩减，相同合并
	"Lessdanmu":true,//弹幕缩减，屏蔽与前n条弹幕重复的字数占比度高于阈值的弹幕
	"Moredanmu":false,//弹幕增量
	"Shortdanmu":true,//上下文相同文字缩减
}

//从config.json初始化
func init(){
	buf := s.New()
	buf.Load("config_F.json")
	for k,v := range buf.B {
		AllF[k] = v.(bool)
	}
}

//功能开关选取函数
func IsOn(s string) bool {
	if v, ok := AllF[s]; ok && v {
		return true
	}
	return false
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
	if!IsOn("ShowRev") {return}
	if ShowRev_start {
		p.Logf().New().Open("danmu.log").Base(1, "Rev").I("营收 ￥", ShowRev_old)
		return
	}
	ShowRev_start = true
	for {
		p.Logf().New().Open("danmu.log").Base(1, "Rev").I("营收 ￥", ShowRev_old)
		for c.Rev == ShowRev_old {p.Sys().Timeoutf(60)}
		ShowRev_old = c.Rev
	}
}

//Gtk 弹幕Gtk窗口
func Gtkf(){
	if!IsOn("Gtk") {return}
	Gtk_danmu()
}

//Ass 弹幕转字幕
type Ass struct {
	
	file string//弹幕ass文件名
	startT time.Time//开始记录的基准时间
	header string//ass开头
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
}

//设定字幕文件名，为""时停止输出
func Ass_f(file string, st time.Time){
	ass.file = file
	if file == "" {return}
	p.Logf().New().Open("danmu.log").Base(1, "Ass").I("保存至", ass.file + ".ass")

	p.File().FileWR(p.Filel{
		File:ass.file + ".ass",
		Write:true,
		Loc:0,
		Context:[]interface{}{ass.header},
	})
	ass.startT = st
}

//传入要显示的单条字幕
func Assf(s string){
	if !IsOn("Ass") {return}
	if ass.file == "" {return}

	if s == "" {return}

	st := time.Since(ass.startT) + time.Duration(p.Rand().MixRandom(0, 2000)) * time.Millisecond
	et := st + time.Duration(Ass_T) * time.Second

	var b string
	// b += "Comment: " + strconv.Itoa(loc) + " "+ Dtos(showedt) + "\n"
	b += `Dialogue: 0,`
	b += dtos(st) + `,` + dtos(et)
	b += `,Default,,0,0,0,,{\fad(200,500)\blur3}` + s + "\n"

	p.File().FileWR(p.Filel{
		File:ass.file + ".ass",
		Write:true,
		Loc:-1,
		Context:[]interface{}{b},
	})
}

//时间转化为0:00:00.00规格字符串
func dtos(t time.Duration) string {
	M := int(math.Floor(t.Minutes())) % 60
	S := int(math.Floor(t.Seconds())) % 60
	Ns := t.Nanoseconds() / int64(time.Millisecond) % 1000 / 10

	return fmt.Sprintf("%d:%02d:%02d.%02d", int(math.Floor(t.Hours())), M, S, Ns)
}

//直播流保存
type Saveflv struct {
	path string
	wait p.Signal
	cancel p.Signal
}

var saveflv = Saveflv {
}

//已go func形式调用，将会获取直播流
func Saveflvf(){
	if !IsOn("Saveflv") {return}
	if saveflv.cancel.Islive() {return}

	l := p.Logf().New().Open("danmu.log").Base(-1, "saveflv")

	cuLinkIndex := 0
	api := F.New_api(c.Roomid)
	for api.Get_live(c.Live_qn).Live_status == 1 {
		c.Live = api.Live

		saveflv.path = strconv.Itoa(c.Roomid) + "_" + time.Now().Format("2006_01_02_15:04:05.000")

		saveflv.wait.Init()
		saveflv.cancel.Init()
		
		rr := p.Req()
		go func(){
			saveflv.cancel.Wait()
			rr.Close()
			os.Rename(saveflv.path+".flv.dtmp", saveflv.path+".flv")
		}()

		Cookie := c.Cookie
		if i := strings.Index(Cookie, "PVID="); i != -1 {
			if d := strings.Index(Cookie[i:], ";"); d == -1 {
				Cookie = Cookie[:i]
			} else {
				Cookie = Cookie[:i] + Cookie[i + d + 1:]
			}
		}

		{//重试
			l.I("尝试连接live")
			if e := rr.Reqf(p.Rval{
				Url:c.Live[cuLinkIndex],
				Retry:10,
				SleepTime:5,
				Header:map[string]string{
					`Cookie`:Cookie,
				},
				Timeout:5,
				JustResponseCode:true,
			}); e != nil{l.W(e)}

			if rr.Response.StatusCode != 200 {
				saveflv.wait.Done()
				saveflv.cancel.Done()
				cuLinkIndex += 1
				if cuLinkIndex >= len(c.Live) {cuLinkIndex = 0}
				continue
			}
		}

		Ass_f(saveflv.path, time.Now())
		l.I("保存到", saveflv.path + ".flv")

		if e := rr.Reqf(p.Rval{
			Url:c.Live[cuLinkIndex],
			Retry:10,
			SleepTime:5,
			Header:map[string]string{
				`Cookie`:Cookie,
			},
			SaveToPath:saveflv.path + ".flv",
			Timeout:-1,
		}); e != nil{l.W(e)}

		l.I("结束")
		Ass_f("", time.Now())//ass
		if !saveflv.cancel.Islive() {break}//cancel
		/*
			Saveflv需要外部组件
			ffmpeg http://ffmpeg.org/download.html
		*/
		// if p.Checkfile().IsExist(saveflv.path+".flv"){
		// 	l.I("转码中")
		// 	p.Exec().Run(false, "ffmpeg", "-i", saveflv.path+".flv", "-c", "copy", saveflv.path+".mkv")
		// 	if p.Checkfile().IsExist(saveflv.path+".mkv"){os.Remove(saveflv.path+".flv")}
		// }

		// l.I("转码结束")
		saveflv.wait.Done()
		saveflv.cancel.Done()
	}
	saveflv.wait.Done()
	saveflv.cancel.Done()
}

//已func形式调用，将会停止保存直播流
func Saveflv_wait(){
	if !IsOn("Saveflv") {return}
	saveflv.cancel.Done()
	p.Logf().New().Open("danmu.log").Base(-1, "saveflv").I("等待").Block()
	saveflv.wait.Wait()
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
	if !IsOn("Obs") {return}
	l := p.Logf().New().Open("danmu.log").Base(1, "obs")
	defer l.BC()

	if on {
		if p.Sys().CheckProgram("obs")[0] != 0 {l.W("obs已经启动");return}
		if p.Sys().CheckProgram("obs")[0] == 0 {
			if obs.Prog == "" {
				l.E("未知的obs程序位置")
				return
			}
			l.I("启动obs")
			p.Exec().Startf(exec.Command(obs.Prog))
			p.Sys().Timeoutf(3)
		}
		
		// Connect a client.
		if err := obs.c.Connect(); err != nil {
			l.E(err)
			return
		}
	} else {
		if p.Sys().CheckProgram("obs")[0] == 0 {l.W("obs未启动");return}
		obs.c.Disconnect()
	}
}

func Obs_R(on bool){
	if !IsOn("Obs") {return}

	l := p.Logf().New().Open("danmu.log").Base(1, "obs_R")
	defer l.BC()

	if p.Sys().CheckProgram("obs")[0] == 0 {
		l.W("obs未启动")
		return
	} else {
		if err := obs.c.Connect(); err != nil {
			l.E(err)
			return
		}
	}
	//录
	if on {
		req := obsws.NewStartRecordingRequest()
		if err := req.Send(obs.c); err != nil {
			l.E(err)
			return
		}
		resp, err := req.Receive()
		if err != nil {
			l.E(err)
			return
		}
		if resp.Status() == "ok" {
			l.I("开始录制")
		}
	} else {
		req := obsws.NewStopRecordingRequest()
		if err := req.Send(obs.c); err != nil {
			l.E(err)
			return
		}
		resp, err := req.Receive()
		if err != nil {
			l.E(err)
			return
		}
		if resp.Status() == "ok" {
			l.I("停止录制")
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
	l := p.Logf().New().Open("danmu.log").Base(1, "autoban")
	l.W(res).Block()
	return true
}

type Danmuji struct {
	buf map[string]string
	Inuse_auto bool

	mute bool
}

var danmuji = Danmuji{
	Inuse_auto:IsOn("Danmuji_auto"),
	buf:map[string]string{
		"弹幕机在么":"在",
	},
}

func Danmujif(s string) {
	if !IsOn("Danmuji") {return}
	if v, ok := danmuji.buf[s]; ok {
		Msg_senddanmu(v)
	}
}

func Danmuji_auto(sleep int) {
	if !IsOn("Danmuji") || !IsOn("Danmuji_auto") || danmuji.mute {return}
	if sleep == 0 {return}

	danmuji.mute = true
	var list = []string{
		"当前正在直播",
		"12345",
	}
	go func(){
		for i := 0; true; i++{
			if i >= len(list) {i = 0}
			Msg_senddanmu(list[i])
			p.Sys().Timeoutf(sleep)
		}
	}()
}

type Autoskip struct {
	num int
	buf sync.Map
	bufbreak chan bool
}

var autoskip = Autoskip{
	bufbreak:make(chan bool, 100),
}

func Autoskipf(s string, maxNum,muteSecond int) int {
	if !IsOn("Autoskip") || s == "" || maxNum <= 0 || muteSecond <= 0 {return 0}
	if v, ok := autoskip.buf.LoadOrStore(s, 0); ok {
		autoskip.buf.Store(s, v.(int) + 1)
		return v.(int) + 1
	}
	
	autoskip.num += 1
	if autoskip.num > maxNum {autoskip.bufbreak <- true}
	
	go func(){
		select {
		case <- autoskip.bufbreak:
		case <- time.After(time.Duration(muteSecond)*time.Second):
		}
		autoskip.num -= 1
		i, ok := autoskip.buf.LoadAndDelete(s);
		if ok && i.(int) > 1 {Msg_showdanmu(nil, strconv.Itoa(i.(int)) + " x " + s)}
	}()
	return 0
}

type Lessdanmu struct {
	buf []string
}

var lessdanmu = Lessdanmu{
}

func Lessdanmuf(s string, bufsize int) float32 {
	if !IsOn("Lessdanmu") {return 0}
	if len(lessdanmu.buf) < bufsize {
		lessdanmu.buf = append(lessdanmu.buf, s)
		return 0
	}

	o := cross(s, lessdanmu.buf)
	if o == 1 {return 1}//完全无用
	Jiezouf(lessdanmu.buf)
	lessdanmu.buf = append(lessdanmu.buf[1:], s)

	return o
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
	if !IsOn("Shortdanmu") {return s}
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
		l := p.Logf().New().Open("danmu.log").Base(1, "jiezou")
		l.W("节奏注意", now, jiezou.avg, S).Block()
		jiezou.avg = now //沉默
		jiezou.Unlock()

		//发送弹幕
		if jiezou.alertdanmu != "" {Msg_senddanmu(jiezou.alertdanmu)}
		return true
	}
	jiezou.Unlock()
	return false
}