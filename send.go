package bili_danmu

import (
	"strings"
	"strconv"

	p "github.com/qydysky/part"
)

//每秒一个令牌，最多等5秒
var danmu_s_limit = p.Limit(1, 1500, 5000)

//弹幕发送
func Danmu_s(msg,Cookie string, roomid int) {
	//等待令牌时阻塞，超时返回false
	if danmu_s_limit.TO() {return}

	l := p.Logf().New().Base(-1, "弹幕发送").Level(1)
	defer l.Block()

	if msg == "" || Cookie == "" || roomid == 0{
		l.E("输入参数不足")
		return
	}
	if i := strings.Index(Cookie, "{"); i != -1 {
		l.E("Cookie格式错误,需为 key=val; key=val 式")
		return
	}

	if i := strings.Index(Cookie, "PVID="); i == -1 {
		l.E("Cookie错误,无PVID=")
		return
	} else {
		if d := strings.Index(Cookie[i:], ";"); d == -1 {
			Cookie = Cookie[:i]
		} else {
			Cookie = Cookie[:i] + Cookie[i + d + 1:]
		}
	}

	var csrf string
	if i := strings.Index(Cookie, "bili_jct="); i == -1 {
		l.E("Cookie错误,无bili_jct=")
		return
	} else {
		if d := strings.Index(Cookie[i + 9:], ";"); d == -1 {
			csrf = Cookie[i + 9:]
		} else {
			csrf = Cookie[i + 9:][:d]
		}
	}

	PostStr := `color=16777215&fontsize=25&mode=1&msg=` + msg + `&rnd=` + strconv.Itoa(int(p.Sys().GetMTime())) + `&roomid=` + strconv.Itoa(roomid) + `&bubble=0&csrf_token=` + csrf + `&csrf=` + csrf

	l.I("发送", msg, "至", roomid)
	r := p.Req()
	err := r.Reqf(p.Rval{
		Url:"https://api.live.bilibili.com/msg/send",
		PostStr:PostStr,
		Timeout:5,
		Referer:"https://live.bilibili.com/" + strconv.Itoa(roomid),
		Cookie:Cookie,
	})
	if err != nil {
		l.E(err)
		return
	}
	
	if code := p.Json().GetValFromS(string(r.Respon), "code");code == nil || code.(float64) != 0 {
		if message := p.Json().GetValFromS(string(r.Respon), "message");message != nil {
			l.E(message)
		} else {
			l.E(string(r.Respon))
		}
		return
	}

}