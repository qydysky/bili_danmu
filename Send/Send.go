package send

import (
	"encoding/json"
	"net/url"
	"strconv"

	c "github.com/qydysky/bili_danmu/CV"

	limit "github.com/qydysky/part/limit"
	reqf "github.com/qydysky/part/reqf"
	sys "github.com/qydysky/part/sys"
)

// 每5s一个令牌，最多等20秒
var danmu_s_limit = limit.New(1, 5000, 20000)

// 弹幕发送
func Danmu_s(msg string, roomid int) {
	//等待令牌时阻塞，超时返回true
	if danmu_s_limit.TO() {
		return
	}

	l := c.C.Log.Base("弹幕发送")

	if msg == "" || roomid == 0 {
		l.L(`E: `, "输入参数不足")
		return
	}

	csrf, _ := c.C.Cookie.LoadV(`bili_jct`).(string)
	if csrf == `` {
		l.L(`E: `, "Cookie错误,无bili_jct=")
		return
	}

	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	PostStr := `color=16777215&fontsize=25&mode=1&msg=` + msg + `&rnd=` + strconv.Itoa(int(sys.Sys().GetSTime())) + `&roomid=` + strconv.Itoa(roomid) + `&bubble=0&csrf_token=` + csrf + `&csrf=` + csrf
	l.L(`I: `, "发送", msg, "至", roomid)

	reqi := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(reqi)
	r := reqi.Item.(*reqf.Req)
	err := r.Reqf(reqf.Rval{
		Url:     "https://api.live.bilibili.com/msg/send",
		PostStr: url.PathEscape(PostStr),
		Retry:   2,
		Timeout: 5 * 1000,
		Proxy:   c.C.Proxy,
		Header: map[string]string{
			`Host`:            `api.live.bilibili.com`,
			`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
			`Accept`:          `application/json, text/javascript, */*; q=0.01`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Content-Type`:    `application/x-www-form-urlencoded; charset=UTF-8`,
			`Origin`:          `https://live.bilibili.com`,
			`Connection`:      `keep-alive`,
			`Pragma`:          `no-cache`,
			`Cache-Control`:   `no-cache`,
			`Referer`:         "https://live.bilibili.com/" + strconv.Itoa(roomid),
			`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
		},
	})
	if err != nil {
		l.L(`E: `, err)
		return
	}

	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	if e := json.Unmarshal(r.Respon, &res); e != nil {
		l.L(`E: `, e)
	}

	if res.Code != 0 {
		l.L(`E: `, `产生错误：`, res.Code, res.Message)
	}
}
