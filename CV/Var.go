package cv

import (
	"time"
)

var (
	Live []string//直播链接
	Live_qn string
	Roomid int
	Cookie string
	Title string
	Rev float64//营收
	Live_Start_Time time.Time//直播开始时间
	Liveing bool//是否在直播
)

//from player-loader-2.0.11.min.js
/*
	customAuthParam
*/
var (
	VERSION = "2.0.11"
)

var Default_qn = map[string]string{
	"10000":"原画",
	"400":"蓝光",
	"250":"超清",
	"150":"高清",
	"80":"流畅",
}