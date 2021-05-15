package cv

import (
	"time"
	syncmap "github.com/qydysky/part/map"
	mq "github.com/qydysky/part/msgq"
	s "github.com/qydysky/part/buf"
	log "github.com/qydysky/part/log"
)


var (
	Uid = 0//client uid

	Live []string//直播流链接
	Live_qn int//当前直播流质量
	Live_want_qn int//期望直播流质量
	Roomid int
	Cookie syncmap.Map
	Title string//直播标题
	Uname string//主播名
	UpUid int//主播uid
	Rev float64//营收
	Renqi int//人气
	GuardNum int//舰长数
	ParentAreaID int//父分区
	AreaID int//子分区
	Locked bool//直播间封禁
	Note string//分区排行
	Live_Start_Time time.Time//直播开始时间
	Liveing bool//是否在直播
	Wearing_FansMedal int//当前佩戴的粉丝牌
	Token string//弹幕钥
	WSURL []string//弹幕链接
	LIVE_BUVID bool//cookies含LIVE_BUVID
)

var (
	Stream_url string//直播Web服务
)

//消息队列
type Danmu_Main_mq_item struct {
	Class string
	Data interface{}
}
//200长度防止push击穿
var Danmu_Main_mq = mq.New(200)

//k-v
var K_v syncmap.Map

func init() {
	buf := s.New()
	buf.Load("config/config_K_v.json")
	for k,v := range buf.B {
		K_v.Store(k, v)
	}
}

//constKv
var (
	Proxy string//全局代理
)
func init() {
	Proxy,_ = K_v.LoadV("http代理地址").(string)
}

//日志
var Log = log.New(log.Config{
	File:`danmu.log`,
	Stdout:true,
	Prefix_string:map[string]struct{}{
		`T: `:log.On,
		`I: `:log.On,
		`N: `:log.On,
		`W: `:log.On,
		`E: `:log.On,
	},
})

func init() {
	logmap := make(map[string]struct{})
	if array,ok := K_v.Load(`日志显示`);ok{
		for _,v := range array.([]interface{}){
			logmap[v.(string)] = log.On
		}
	}
	Log = Log.Level(logmap)
	return
}

//from player-loader-2.0.11.min.js
/*
	customAuthParam
*/
// var (
// 	VERSION = "2.0.11"
// ) // 不再需要

//允许的清晰度

var (
	AcceptQn = map[int]string{
		10000:"原画",
		800:"4K",
		401:"蓝光(杜比)",
		400:"蓝光",
		250:"超清",
		150:"高清",
		80:"流畅",
	}
	Qn = map[int]string{// no change
		10000:"原画",
		800:"4K",
		401:"蓝光(杜比)",
		400:"蓝光",
		250:"超清",
		150:"高清",
		80:"流畅",
	}
)