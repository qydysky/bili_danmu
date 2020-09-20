package bili_danmu

import (
	// "fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	p "github.com/qydysky/part"
) 

//功能开关
var AllF = map[string]bool{
	"Autoban":true,//自动封禁(仅提示，未完成)
	"Jiezou":true,//带节奏预警，提示弹幕礼仪
	"Danmuji":true,//反射型弹幕机，回应弹幕
	"Danmuji_auto":false,//自动型弹幕机，定时输出
	"Autoskip":true,//刷屏缩减，相同合并
	"Lessdanmu":true,//弹幕缩减，屏蔽与前n条弹幕重复的字数占比度高于阈值的弹幕
	"Moredanmu":false,//弹幕增量
	"Shortdanmu":true,//上下文相同文字缩减
}

func IsOn(s string) bool {
	if v, ok := AllF[s]; ok && v {
		return true
	}
	return false
}
//公共
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
func selfcross(a string) (float32) {
	buf := make(map[rune]bool)
	for _,v := range a {
		if _,ok := buf[v]; !ok {
			buf[v] = true
		}
	}
	return 1 - float32(len(buf)) / float32(len([]rune(a)))
}
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
type Autoban struct {
	Banbuf []string
	buf []string
	Inuse bool
}

var autoban = Autoban {
	Inuse:IsOn("Autoban"),
}

func Autobanf(s string) bool {
	if !autoban.Inuse {return false}

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

	pt := float32(len([]rune(s)))
	if pt <= 5 {return false}//字数过少去除
	res = append(res, pt)

	pt = selfcross(s);
	if pt > 0.5 {return false}//自身重复高去除
	res = append(res, pt)

	pt = cross(s, autoban.buf);
	if pt < 0.7 {return false}//历史重复低去除
	res = append(res, pt)

	pt = cross(s, autoban.Banbuf);
	if pt < 0.8 {return false}//ban字符重复低去除
	res = append(res, pt)

	l := p.Logf().New().Open("danmu.log").Base(1, "autoban")
	l.W(res)
	return true
}

type Danmuji struct {
	buf map[string]string
	Inuse bool
	Inuse_auto bool

	mute bool
}

var danmuji = Danmuji{
	Inuse:IsOn("Danmuji"),
	Inuse_auto:IsOn("Danmuji_auto"),
	buf:map[string]string{
		"弹幕机在么":"在",
	},
}

func Danmujif(s string) {
	if !danmuji.Inuse {return}
	if v, ok := danmuji.buf[s]; ok {
		Msg_senddanmu(v)
	}
}

func Danmuji_auto(sleep int) {
	if !danmuji.Inuse || !danmuji.Inuse_auto || danmuji.mute {return}
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
	Inuse bool
	num int
	buf sync.Map
	bufbreak chan bool
}

var autoskip = Autoskip{
	Inuse:IsOn("Autoskip"),
	bufbreak:make(chan bool, 10),
}

func Autoskipf(s string, maxNum,muteSecond int) int {
	if !autoskip.Inuse || s == "" || maxNum <= 0 || muteSecond <= 0 {return 0}
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
	Inuse bool
	buf []string
}

var lessdanmu = Lessdanmu{
	Inuse:IsOn("Lessdanmu"),
}

func Lessdanmuf(s string, bufsize int) float32 {
	if !lessdanmu.Inuse {return 0}
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
// 	Inuse bool
// 	buf []string
// }

// var moredanmu = Moredanmu{
// 	Inuse:IsOn("Moredanmu"),
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
	Inuse bool
	lastdanmu []rune
}

var shortdanmu = Shortdanmu{
	Inuse:IsOn("Shortdanmu"),
}

func Shortdanmuf(s string) string {
	if !shortdanmu.Inuse {return s}
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
	if new == "" {new = "...."}
	shortdanmu.lastdanmu = []rune(s)
	return new
}

type Jiezou struct {
	Inuse bool
	alertdanmu string
	skipS map[string]interface{}

	avg float32
	turn int
	sync.Mutex
}

var jiezou = Jiezou{
	Inuse:IsOn("Jiezou"),
	alertdanmu:"",
	skipS:map[string]interface{}{//常见语气词忽略
		"了":nil,
		"的":nil,
		"哈":nil,
		"是":nil,
	},
}

func Jiezouf(s []string) bool {
	if !jiezou.Inuse {return false}
	now,S := selfcross2(s)
	jiezou.avg = (9 * jiezou.avg + now)/10
	if jiezou.turn < len(s) {jiezou.turn += 1;return false}
	
	if _,ok := jiezou.skipS[S]; ok {return false}

	jiezou.Lock()
	if now > 1.3 * jiezou.avg {//触发
		l := p.Logf().New().Open("danmu.log").Base(1, "jiezou")
		l.W("节奏注意", now, jiezou.avg, S)
		jiezou.avg = now //沉默
		jiezou.Unlock()

		//发送弹幕
		if jiezou.alertdanmu != "" {Msg_senddanmu(jiezou.alertdanmu)}
		return true
	}
	jiezou.Unlock()
	return false
}