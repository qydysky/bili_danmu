package send

import (
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"time"

	c "github.com/qydysky/bili_danmu/CV"

	file "github.com/qydysky/part/file"
	limit "github.com/qydysky/part/limit"
	reqf "github.com/qydysky/part/reqf"
)

// 每5s一个令牌，最多等20秒
var danmu_s_limit = limit.New(1, "5s", "20s")
var damnu_official = make(map[string]string)

// 初始化表情代码
func init() {
	f := file.New("config/config_danmu_official.json", 0, true)
	if !f.IsExist() {
		return
	}
	bb, err := f.ReadAll(1000, 1<<16)
	if !errors.Is(err, io.EOF) {
		return
	}
	var buf map[string]interface{}
	_ = json.Unmarshal(bb, &buf)
	for k, v := range buf {
		if k == v {
			continue
		}
		damnu_official[k] = v.(string)
	}
}

// 弹幕发送
func Danmu_s(msg string, roomid int) error {
	data := map[string]string{
		`msg`:    msg,
		`roomid`: strconv.Itoa(roomid),
	}

	if v, ok := c.C.K_v.LoadV(`弹幕_识别表情代码`).(bool); ok && v {
		if v, ok := damnu_official[msg]; ok {
			data[`msg`] = v
			data[`dm_type`] = `1`
		}
	}

	return Danmu_s2(data)
}

var (
	ErrLimit     = errors.New("ErrLimit")
	ErrMsgEmpty  = errors.New("ErrMsgEmpty")
	ErrRoomEmpty = errors.New("ErrRoomEmpty")
	ErrNoLogin   = errors.New("ErrNoLogin")
	ErrRes       = errors.New("ErrRes")
)

// 通用发送
func Danmu_s2(data map[string]string) error {
	//等待令牌时阻塞，超时返回true
	if danmu_s_limit.TO() {
		return ErrLimit
	}

	l := c.C.Log.Base("弹幕发送")

	if _, ok := data[`msg`]; !ok {
		l.L(`E: `, "必须输入参数msg")
		return ErrMsgEmpty
	}

	if _, ok := data[`roomid`]; !ok {
		l.L(`E: `, "必须输入参数roomid")
		return ErrRoomEmpty
	}

	csrf, _ := c.C.Cookie.LoadV(`bili_jct`).(string)
	if csrf == `` {
		l.L(`E: `, "Cookie错误,无bili_jct=")
		return ErrNoLogin
	}

	if _, ok := data[`bubble`]; !ok {
		data[`bubble`] = `0`
	}
	if _, ok := data[`color`]; !ok {
		data[`color`] = `5816798`
	}
	if _, ok := data[`mode`]; !ok {
		data[`mode`] = `1`
	}
	if _, ok := data[`fontsize`]; !ok {
		data[`fontsize`] = `25`
	}
	data[`rnd`] = strconv.Itoa(int(time.Now().Unix()))
	data[`csrf`] = csrf
	data[`csrf_token`] = csrf

	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	postStr, contentType := reqf.ToForm(data)
	l.L(`I: `, "发送", data[`msg`], "至", data[`roomid`])

	r := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(r)
	err := r.Reqf(reqf.Rval{
		Url:     "https://api.live.bilibili.com/msg/send",
		PostStr: postStr,
		Retry:   2,
		Timeout: 5 * 1000,
		Proxy:   c.C.Proxy,
		Header: map[string]string{
			`Host`:            `api.live.bilibili.com`,
			`User-Agent`:      c.UA,
			`Accept`:          `application/json, text/javascript, */*; q=0.01`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Content-Type`:    contentType,
			`Origin`:          `https://live.bilibili.com`,
			`Connection`:      `keep-alive`,
			`Pragma`:          `no-cache`,
			`Cache-Control`:   `no-cache`,
			`Referer`:         "https://live.bilibili.com/" + data[`roomid`],
			`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
		},
	})
	if err != nil {
		l.L(`E: `, err)
		return err
	}

	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	if e := json.Unmarshal(r.Respon, &res); e != nil {
		l.L(`E: `, e)
		return e
	}

	if res.Code != 0 {
		l.L(`E: `, `产生错误：`, res.Code, res.Message)
		return errors.Join(ErrRes, errors.New(res.Message))
	}

	return nil
}
