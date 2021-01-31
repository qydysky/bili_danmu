package cv

import (
	"time"
	mq "github.com/qydysky/part/msgq"
	s "github.com/qydysky/part/buf"
	log "github.com/qydysky/part/log"
)


var (
	Uid = 0//client uid

	Live []string//直播链接
	Live_qn string
	Roomid int
	Cookie = make(map[string]string)
	Title string//直播标题
	Uname string//主播名
	Rev float64//营收
	Renqi int//人气
	GuardNum int//舰长数
	Note string//分区排行
	Live_Start_Time time.Time//直播开始时间
	Liveing bool//是否在直播
)

//消息队列
type Danmu_Main_mq_item struct {
	Class string
	Data interface{}
}
//200长度防止push击穿
var Danmu_Main_mq = mq.New(200)

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

//k-v
var K_v =make(map[string]interface{})

func init() {
	buf := s.New()
	buf.Load("config/config_K_v.json")
	for k,v := range buf.B {
		K_v[k] = v
	}
}

//from player-loader-2.0.11.min.js
/*
	customAuthParam
*/
var (
	VERSION = "2.0.11"
)

var Default_qn = map[string]string{
	"10000":"原画",
	"800":"4K",
	"401":"蓝光(杜比)",
	"400":"蓝光",
	"250":"超清",
	"150":"高清",
	"80":"流畅",
}