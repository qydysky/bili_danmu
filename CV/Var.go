package cv
 
var (
	Live []string//直播链接
	Live_qn string
	Roomid int
	Cookie string
	Title string
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