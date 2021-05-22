package part

type VOICE_JOIN_STATUS struct {
	Cmd  string `json:"cmd"`
	Data struct {
		RoomID       int    `json:"room_id"`
		Status       int    `json:"status"`
		Channel      string `json:"channel"`
		ChannelType  string `json:"channel_type"`
		UID          int    `json:"uid"`
		UserName     string `json:"user_name"`
		HeadPic      string `json:"head_pic"`
		Guard        int    `json:"guard"`
		StartAt      int    `json:"start_at"`
		CurrentTime  int    `json:"current_time"`
		WebShareLink string `json:"web_share_link"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}
//{"cmd":"VOICE_JOIN_STATUS","data":{"room_id":1017,"status":0,"channel":"","channel_type":"voice","uid":0,"user_name":"","head_pic":"","guard":0,"start_at":0,"current_time":1621702198,"web_share_link":"https:\/\/live.bilibili.com\/h5\/1017"},"roomid":1017}