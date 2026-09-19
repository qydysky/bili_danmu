package part

type PANEL_INTERACTIVE_NOTIFY_CHANGE struct {
	Cmd  string `json:"cmd"`
	Data struct {
		BizID    int    `json:"biz_id"`
		EndTime  int    `json:"end_time"`
		Icon     string `json:"icon"`
		LastTime int    `json:"last_time"`
		Level    int    `json:"level"`
		Text     string `json:"text"`
	} `json:"data"`
}

/*
预言
{"cmd":"PANEL_INTERACTIVE_NOTIFY_CHANGE","data":{"biz_id":4,"end_time":180,"icon":"https://i0.hdslb.com/bfs/live/164a37487431ce065981d76afe6c2fb2083facee.png","last_time":5,"level":1,"text":"主播开启预言"}}

{"cmd":"PANEL_INTERACTIVE_NOTIFY_CHANGE","data":{"biz_id":4,"end_time":0,"icon":"https://i0.hdslb.com/bfs/live/164a37487431ce065981d76afe6c2fb2083facee.png","last_time":0,"level":1,"text":"预言状态变更"}}
*/
