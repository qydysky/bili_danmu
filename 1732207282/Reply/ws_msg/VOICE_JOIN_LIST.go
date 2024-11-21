package part

type VOICE_JOIN_LIST struct {
	Cmd  string `json:"cmd"`
	Data struct {
		RoomID     int `json:"room_id"`
		Category   int `json:"category"`
		ApplyCount int `json:"apply_count"`
		RedPoint   int `json:"red_point"`
		Refresh    int `json:"refresh"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

//{"cmd":"VOICE_JOIN_LIST","data":{"room_id":1017,"category":1,"apply_count":62,"red_point":1,"refresh":0},"roomid":1017}
