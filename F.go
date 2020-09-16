package bili_danmu

import (
	"fmt"
	"bytes"
	"sync"
	"time"

	p "github.com/qydysky/part"
) 

//功能开关
var AllF = map[string]bool{
	"Autoban":false,//自动封禁(仅提示，未完成)
	"Danmuji":true,//反射型弹幕机，回应弹幕
	"Danmuji_auto":false,//自动型弹幕机，定时输出
	"Autoskip":true,//刷屏缩减，相同合并
	"Lessdanmu":true,//弹幕缩减，显示差异大的
}

func IsOn(s string) bool {
	if v, ok := AllF[s]; ok && v {
		return true
	}
	return false
}

//功能区
type Autoban struct {
	buf []byte
	Inuse bool
}

var autoban = Autoban {
	Inuse:IsOn("Autoban"),
}

func Autobanf(s string) float32 {
	if !autoban.Inuse {return 0}

	if len(autoban.buf) == 0 {
		f := p.File().FileWR(p.Filel{
			File:"Autoban.txt",
			Write:false,
		})
		autoban.buf = []byte(f)
	}

	var scop int
	for _, v := range []byte(s) {
		if bytes.Contains(autoban.buf, []byte{v}) {scop += 1}
	}
	return float32(scop) / float32(len(s))
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

func Danmujif(s,cookie string, roomid int) {
	if !danmuji.Inuse {return}
	if cookie == "" || roomid == 0 {return}
	if v, ok := danmuji.buf[s]; ok {
		Danmu_s(v, cookie, roomid)
	}
}

func Danmuji_auto(cookie string, sleep,roomid int) {
	if !danmuji.Inuse || !danmuji.Inuse_auto || danmuji.mute {return}
	if cookie == "" || roomid == 0 || sleep == 0 {return}

	danmuji.mute = true
	var list = []string{
		"当前正在直播",
		"12345",
	}
	go func(){
		for i := 0; true; i++{
			if i >= len(list) {i = 0}
			Danmu_s(list[i], cookie, roomid)
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
		if ok && i.(int) > 0 {fmt.Println(s, "+", i)}
	}()
	return 0
}

type Lessdanmu struct {
	Inuse bool
	buf []string

	avg float32
}

var lessdanmu = Lessdanmu{
	Inuse:IsOn("Lessdanmu"),
}

func Lessdanmuf(s string, bufsize int) bool {
	if !lessdanmu.Inuse {return false}
	if len(lessdanmu.buf) > bufsize {
		lessdanmu.buf = append(lessdanmu.buf[1:], s)
	} else {
		lessdanmu.buf = append(lessdanmu.buf, s)
	}

	o := cross(s, lessdanmu.buf)
	lessdanmu.avg = (0.8 * lessdanmu.avg + 0.2 * o)
	return o > lessdanmu.avg 
}

func cross(a string,buf []string) float32 {
	var (
		s float32
		all float32
	)
	for _,v1 := range buf {
		for _,v2 := range v1 {
			for _,v3 := range a {
				if v3 == v2 {s += 1}
				all += 1
			}
		}

	}
	return s / all
}