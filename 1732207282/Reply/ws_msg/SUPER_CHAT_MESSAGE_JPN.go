package part

type SUPER_CHAT_MESSAGE_JPN struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ID                    string      `json:"id"`
		UID                   string      `json:"uid"`
		Price                 int         `json:"price"`
		Rate                  int         `json:"rate"`
		Message               string      `json:"message"`
		MessageJpn            string      `json:"message_jpn"`
		IsRanked              int         `json:"is_ranked"`
		BackgroundImage       string      `json:"background_image"`
		BackgroundColor       string      `json:"background_color"`
		BackgroundIcon        string      `json:"background_icon"`
		BackgroundPriceColor  string      `json:"background_price_color"`
		BackgroundBottomColor string      `json:"background_bottom_color"`
		Ts                    int         `json:"ts"`
		Token                 string      `json:"token"`
		MedalInfo             interface{} `json:"medal_info"`
		UserInfo              struct {
			Uname      string `json:"uname"`
			Face       string `json:"face"`
			FaceFrame  string `json:"face_frame"`
			GuardLevel int    `json:"guard_level"`
			UserLevel  int    `json:"user_level"`
			LevelColor string `json:"level_color"`
			IsVip      int    `json:"is_vip"`
			IsSvip     int    `json:"is_svip"`
			IsMainVip  int    `json:"is_main_vip"`
			Title      string `json:"title"`
			Manager    int    `json:"manager"`
		} `json:"user_info"`
		Time      int `json:"time"`
		StartTime int `json:"start_time"`
		EndTime   int `json:"end_time"`
		Gift      struct {
			Num      int    `json:"num"`
			GiftID   int    `json:"gift_id"`
			GiftName string `json:"gift_name"`
		} `json:"gift"`
	} `json:"data"`
	Roomid string `json:"roomid"`
}

/*
{"cmd":"SUPER_CHAT_MESSAGE_JPN","data":{"id":"1852575","uid":"696837750","price":2000,"rate":1000,"message":"\u8fd9\u80fd\u770b\u5230\u51e0\
u70b9","message_jpn":"","is_ranked":1,"background_image":"https:\/\/i0.hdslb.com\/bfs\/live\/a712efa5c6ebc67bafbe8352d3e74b820a00c13e.png","
background_color":"#FFD8D8","background_icon":"https:\/\/i0.hdslb.com\/bfs\/live\/0d9cbbdbad7d3371266cd5b568065415415316ae.png","background_
price_color":"#C86A7A","background_bottom_color":"#AB1A32","ts":1623612710,"token":"5491A5EB","medal_info":null,"user_info":{"uname":"Red\u4
e00","face":"http:\/\/i0.hdslb.com\/bfs\/face\/member\/noface.jpg","face_frame":"http:\/\/i0.hdslb.com\/bfs\/live\/ceb8e7fc5e207b8e6219c9917
e3ef1b22f3df61a.png","guard_level":1,"user_level":43,"level_color":"#ff86b2","is_vip":0,"is_svip":0,"is_main_vip":0,"title":"0","manager":0}
,"time":7200,"start_time":1623612710,"end_time":1623619910,"gift":{"num":1,"gift_id":12000,"gift_name":"\u9192\u76ee\u7559\u8a00"}},"roomid"
:"47867"}
*/
