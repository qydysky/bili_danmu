package send

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	J "github.com/qydysky/bili_danmu/Json"

	limit "github.com/qydysky/part/limit"
	reqf "github.com/qydysky/part/reqf"
)

// 每2s一个令牌，最多等10秒
var gift_limit = limit.New(1, "2s", "10s")

func Send_gift(gift_id, bag_id, gift_num int) {
	log := c.C.Log.Base_add(`发送礼物`)

	if gift_limit.TO() {
		log.L(`W: `, "超时")
		return
	}

	if c.C.UpUid == 0 {
		log.L(`W: `, "还未获取到Up主uid")
		return
	}

	{ //发送请求（银瓜子礼物）
		csrf, _ := c.C.Cookie.LoadV(`bili_jct`).(string)
		if csrf == `` {
			log.L(`E: `, "Cookie错误,无bili_jct=")
			return
		}

		var sendStr = `uid=` + strconv.Itoa(c.C.Uid) + `&` +
			`gift_id=` + strconv.Itoa(gift_id) + `&` +
			`ruid=` + strconv.Itoa(c.C.UpUid) + `&` +
			`send_ruid=0&` +
			`gift_num=` + strconv.Itoa(gift_num) + `&` +
			`bag_id=` + strconv.Itoa(bag_id) + `&` +
			`platform=pc&` +
			`biz_code=live&` +
			`biz_id=` + strconv.Itoa(c.C.Roomid) + `&` +
			`rnd=` + strconv.Itoa(int(time.Now().Unix())) + `&` +
			`storm_beat_id=0&` +
			`metadata=&` +
			`price=0&` +
			`csrf_token=` + csrf + `&` +
			`csrf=` + csrf + `&` +
			`visit_id=`

		Cookie := make(map[string]string)
		c.C.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(req)
		if e := req.Reqf(reqf.Rval{
			Url:     `https://api.live.bilibili.com/xlive/revenue/v2/gift/sendBag`,
			PostStr: url.PathEscape(sendStr),
			Timeout: 10 * 1000,
			Proxy:   c.C.Proxy,
			Header: map[string]string{
				`Host`:            `api.vc.bilibili.com`,
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
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
			return
		}

		var res J.SendBag

		if e := json.Unmarshal(req.Respon, &res); e != nil {
			log.L(`E: `, e)
			return
		}

		if res.Code != 0 {
			log.L(`W: `, res.Message)
			return
		}
		for i := 0; i < len(res.Data.GiftList); i++ {
			log.L(`I: `, `给`, c.C.Roomid, `赠送了`, res.Data.GiftList[i].GiftNum, `个`, res.Data.GiftList[i].GiftName)
		}
	}
}
