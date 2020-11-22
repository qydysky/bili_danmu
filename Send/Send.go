package send

import (
	"strings"
	"strconv"

	p "github.com/qydysky/part"
)

//每5s一个令牌，最多等10秒
var danmu_s_limit = p.Limit(1, 5000, 10000)

//弹幕发送
func Danmu_s(msg,Cookie string, roomid int) {
	//等待令牌时阻塞，超时返回true
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

	PostStr := `color=16777215&fontsize=25&mode=1&msg=` + msg + `&rnd=` + strconv.Itoa(int(p.Sys().GetSTime())) + `&roomid=` + strconv.Itoa(roomid) + `&bubble=0&csrf_token=` + csrf + `&csrf=` + csrf
	l.I("发送", msg, "至", roomid)
	r := p.Req()
	err := r.Reqf(p.Rval{
		Url:"https://api.live.bilibili.com/msg/send",
		PostStr:PostStr,
		Timeout:5,
		Header:map[string]string{
			`Host`: `api.live.bilibili.com`,
			`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
			`Accept`: `application/json, text/javascript, */*; q=0.01`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Content-Type`: `application/x-www-form-urlencoded; charset=UTF-8`,
			`Origin`: `https://live.bilibili.com`,
			`Connection`: `keep-alive`,
			`Pragma`: `no-cache`,
			`Cache-Control`: `no-cache`,
			`Referer`:"https://live.bilibili.com/" + strconv.Itoa(roomid),
			`Cookie`:Cookie,
		},
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