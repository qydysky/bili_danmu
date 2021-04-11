package reply

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"math"
	"time"
	"os/exec"
    "path/filepath"

	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	send "github.com/qydysky/bili_danmu/Send"
	
	p "github.com/qydysky/part"
	funcCtrl "github.com/qydysky/part/funcCtrl"
	msgq "github.com/qydysky/part/msgq"
	reqf "github.com/qydysky/part/reqf"
	b "github.com/qydysky/part/buf"
	s "github.com/qydysky/part/signal"

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
	c.Log.Base(`Ass`).L(`I: `,"保存至", ass.file + ".ass")

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
	wait *s.Signal
	cancel *s.Signal
	skipFunc funcCtrl.SkipFunc
}

var saveflv = Saveflv {
}

func init(){
	//使用带tag的消息队列在功能间传递消息
	c.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
		`saveflv`:func(data interface{})(bool){
			if saveflv.cancel.Islive() {
				Saveflv_wait()
			} else {
				go Saveflvf()
			}

			return false
		},
	})
}

//已go func形式调用，将会获取直播流
func Saveflvf(){
	l := c.Log.Base(`saveflv`)

	//避免多次开播导致的多次触发
	{
		if saveflv.skipFunc.NeedSkip() {
			l.L(`T: `,`已存在实例`)
			return
		}
		defer saveflv.skipFunc.UnSet()
	}

	qn, ok := c.K_v.LoadV("flv直播流清晰度").(float64)
	if !ok || qn < 0 {return}

	{
		AcceptQn := []int{}
		for k,_ := range c.AcceptQn {
			if k <= int(qn) {AcceptQn = append(AcceptQn, k)}
		}
		MaxQn := 0
		for i:=0; len(AcceptQn)>i; i+=1{
			if AcceptQn[i] > MaxQn {
				MaxQn = AcceptQn[i]
			}
		}
		if MaxQn == 0 {
			l.L(`W: `,"使用默认清晰度")
		}
		c.Live_qn = MaxQn
	}

	if saveflv.cancel.Islive() {return}

	cuLinkIndex := 0
	for {
		F.Get(`Liveing`)
		if !c.Liveing {break}

		F.Get(`Live`)
		if len(c.Live)==0 {break}

		if path,ok := c.K_v.LoadV("直播流保存位置").(string);ok{
			if path,err := filepath.Abs(path);err == nil{
				saveflv.path = path+"/"
			}
		}
		

		saveflv.path += strconv.Itoa(c.Roomid) + "_" + time.Now().Format("2006_01_02_15-04-05-000")

		saveflv.wait = s.Init()
		saveflv.cancel = s.Init()
		
		rr := reqf.Req()
		go func(){
			saveflv.cancel.Wait()
			rr.Close()
		}()


		CookieM := make(map[string]string)
		c.Cookie.Range(func(k,v interface{})(bool){
			CookieM[k.(string)] = v.(string)
			return true
		})

		{//重试
			l.L(`I: `,"尝试连接live")
			if e := rr.Reqf(reqf.Rval{
				Url:c.Live[cuLinkIndex],
				Retry:10,
				SleepTime:5,
				Header:map[string]string{
					`Cookie`:reqf.Map_2_Cookies_String(CookieM),
				},
				Timeout:5,
				JustResponseCode:true,
			}); e != nil{l.L(`W: `,e)}

			if rr.Response == nil ||
			rr.Response.StatusCode != 200 {
				saveflv.wait.Done()
				saveflv.cancel.Done()
				cuLinkIndex += 1
				if cuLinkIndex >= len(c.Live) {cuLinkIndex = 0}
				time.Sleep(time.Second*5)
				continue
			}
		}

		Ass_f(saveflv.path, time.Now())
		l.L(`I: `,"保存到", saveflv.path + ".flv")

		if e := rr.Reqf(reqf.Rval{
			Url:c.Live[cuLinkIndex],
			Retry:10,
			SleepTime:5,
			Header:map[string]string{
				`Cookie`:reqf.Map_2_Cookies_String(CookieM),
			},
			SaveToPath:saveflv.path + ".flv",
			Timeout:-1,
		}); e != nil{l.L(`W: `,e)}

		l.L(`I: `,"结束")
		Ass_f("", time.Now())//ass
		p.FileMove(saveflv.path+".flv.dtmp", saveflv.path+".flv")
		if !saveflv.cancel.Islive() {break}//cancel
		/*
			Saveflv需要外部组件
			ffmpeg http://ffmpeg.org/download.html
		*/
		// if p.Checkfile().IsExist(saveflv.path+".flv"){
		// 	l.L(`I: `,"转码中")
		// 	p.Exec().Run(false, "ffmpeg", "-i", saveflv.path+".flv", "-c", "copy", saveflv.path+".mkv")
		// 	if p.Checkfile().IsExist(saveflv.path+".mkv"){os.Remove(saveflv.path+".flv")}
		// }

		// l.L(`I: `,"转码结束")
		saveflv.wait.Done()
		saveflv.cancel.Done()
	}
	saveflv.wait.Done()
	saveflv.cancel.Done()
}

//已func形式调用，将会停止保存直播流
func Saveflv_wait(){
	qn, ok := c.K_v.LoadV("flv直播流清晰度").(float64)
	if !ok || qn < 0 {return}

	saveflv.cancel.Done()
	c.Log.Base(`saveflv`).L(`T: `,"等待")
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

	mute bool
}

var danmuji = Danmuji{
	Inuse_auto:IsOn("自动弹幕机"),
	Buf:map[string]string{
		"弹幕机在么":"在",
	},
}

func init(){//初始化反射型弹幕机
	buf := b.New()
	buf.Load("config/config_auto_reply.json")
	for k,v := range buf.B {
		danmuji.Buf[k] = v.(string)
	}
}

func Danmujif(s string) {
	if !IsOn("反射弹幕机") {return}
	if v, ok := danmuji.Buf[s]; ok {
		Msg_senddanmu(v)
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
}

var lessdanmu = Lessdanmu{
}

func Lessdanmuf(s string, bufsize int) float32 {
	if !IsOn("相似弹幕忽略") {return 0}
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