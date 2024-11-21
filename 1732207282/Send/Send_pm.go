package send

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"

	c "github.com/qydysky/bili_danmu/CV"

	limit "github.com/qydysky/part/limit"
	reqf "github.com/qydysky/part/reqf"
	sys "github.com/qydysky/part/sys"

	"github.com/google/uuid"
)

type Pm_item struct {
	Uid int
	Msg string
}

// 每5s一个令牌，最多等10秒
var pm_limit = limit.New(1, "5s", "10s")

func Send_pm(uid int, msg string) error {
	if msg == `` || uid == 0 {
		return errors.New(`msg == "" || uid == 0`)
	}

	log := c.C.Log.Base_add(`私信`)

	if c.C.Uid == 0 {
		log.L(`E: `, `client uid == 0`)
		return errors.New(`client uid == 0`)
	} else if c.C.Uid == uid {
		log.L(`W: `, `不能发送给自己`)
		return errors.New(`不能发送给自己`)
	}

	csrf, _ := c.C.Cookie.LoadV(`bili_jct`).(string)
	if csrf == `` {
		return errors.New("Cookie错误,无bili_jct=")
	}

	if pm_limit.TO() {
		return errors.New("TO")
	}

	var send_str = `msg[sender_uid]=` + strconv.Itoa(c.C.Uid) + `&msg[receiver_id]=` + strconv.Itoa(uid) + `&msg[receiver_type]=1&msg[msg_type]=1&msg[msg_status]=0&msg[content]={"content":"` + msg + `"}&msg[timestamp]=` + strconv.Itoa(int(sys.Sys().GetSTime())) + `&msg[new_face_version]=0&msg[dev_id]=` + strings.ToUpper(uuid.New().String()) + `&from_firework=0&build=0&mobi_app=web&csrf_token=` + csrf + `&csrf=` + csrf

	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	req := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(req)
	if e := req.Reqf(reqf.Rval{
		Url:     `https://api.vc.bilibili.com/web_im/v1/web_im/send_msg`,
		PostStr: url.PathEscape(send_str),
		Timeout: 10 * 1000,
		Proxy:   c.C.Proxy,
		Header: map[string]string{
			`Host`:            `api.vc.bilibili.com`,
			`User-Agent`:      c.UA,
			`Accept`:          `application/json, text/javascript, */*; q=0.01`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Content-Type`:    `application/x-www-form-urlencoded; charset=UTF-8`,
			`Origin`:          `https://message.bilibili.com`,
			`Connection`:      `keep-alive`,
			`Pragma`:          `no-cache`,
			`Cache-Control`:   `no-cache`,
			`Referer`:         "https://message.bilibili.com",
			`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
		},
	}); e != nil {
		log.L(`E: `, e)
		return e
	}

	J := struct {
		Code int `json:"code"`
	}{}

	if e := json.Unmarshal(req.Respon, &J); e != nil {
		log.L(`E: `, string(req.Respon))
		return errors.New(string(req.Respon))
	}

	log.L(`I: `, `发送私信给`, uid, `:`, msg)
	return nil
}
