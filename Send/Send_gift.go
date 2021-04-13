package send

import (
	"time"
	"net/url"
	"encoding/json"
	"strconv"

	c "github.com/qydysky/bili_danmu/CV"

	reqf "github.com/qydysky/part/reqf"
	limit "github.com/qydysky/part/limit"
)

//每2s一个令牌，最多等10秒
var gift_limit = limit.New(1, 2000, 10000)

func Send_gift(gift_id,bag_id,gift_num int) {
	log := c.Log.Base_add(`发送礼物`)

	if gift_limit.TO() {log.L(`W: `,"超时");return}

	if c.UpUid == 0 {log.L(`W: `,"还未获取到Up主uid");return}

	{//发送请求（银瓜子礼物）
		csrf,_ := c.Cookie.LoadV(`bili_jct`).(string)
		if csrf == `` {log.L(`E: `,"Cookie错误,无bili_jct=");return}

		var sendStr = 
		`uid=`+strconv.Itoa(c.Uid)+`&`+
		`gift_id=`+strconv.Itoa(gift_id)+`&`+
		`ruid=`+strconv.Itoa(c.UpUid)+`&`+
		`send_ruid=0&`+
		`gift_num=`+strconv.Itoa(gift_num)+`&`+
		`bag_id=`+strconv.Itoa(bag_id)+`&`+
		`platform=pc&`+
		`biz_code=live&`+
		`biz_id=`+strconv.Itoa(c.Roomid)+`&`+
		`rnd=`+strconv.Itoa(int(time.Now().Unix()))+`&`+
		`storm_beat_id=0&`+
		`metadata=&`+
		`price=0&`+
		`csrf_token=`+csrf+`&`+
		`csrf=`+csrf+`&`+
		`visit_id=`

		Cookie := make(map[string]string)
		c.Cookie.Range(func(k,v interface{})(bool){
			Cookie[k.(string)] = v.(string)
			return true
		})
		
		req := reqf.Req()
		if e:= req.Reqf(reqf.Rval{
			Url:`https://api.live.bilibili.com/gift/v2/live/bag_send`,
			PostStr:url.PathEscape(sendStr),
			Timeout:10,
			Header:map[string]string{
				`Host`: `api.vc.bilibili.com`,
				`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
				`Accept`: `application/json, text/javascript, */*; q=0.01`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Content-Type`: `application/x-www-form-urlencoded; charset=UTF-8`,
				`Origin`: `https://message.bilibili.com`,
				`Connection`: `keep-alive`,
				`Pragma`: `no-cache`,
				`Cache-Control`: `no-cache`,
				`Referer`:"https://message.bilibili.com",
				`Cookie`:reqf.Map_2_Cookies_String(Cookie),
			},
		});e != nil {
			log.L(`E: `,e)
			return
		}
	
		var res struct{
			Code int `json:"code"`
			Message string `json:"message"`
			Data struct{
				Gift_name string `json:"gift_name"`
				Gift_num int `json:"gift_num"`
			} `json:"data"`
		}

		if e := json.Unmarshal(req.Respon, &res);e != nil {
			log.L(`E: `,e)
			return
		}

		if res.Code != 0 {
			log.L(`W: `,res.Message)
			return
		}
		log.L(`I: `,`给`,c.Roomid,`赠送了`,res.Data.Gift_num,`个`,res.Data.Gift_name)
	}
}
