package part

type VOICE_JOIN_ROOM_COUNT_INFO struct {
	Cmd  string `json:"cmd"`
	Data struct {
		RoomID      int `json:"room_id"`
		RootStatus  int `json:"root_status"`
		RoomStatus  int `json:"room_status"`
		ApplyCount  int `json:"apply_count"`
		NotifyCount int `json:"notify_count"`
		RedPoint    int `json:"red_point"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

//{"cmd":"VOICE_JOIN_ROOM_COUNT_INFO","data":{"room_id":1017,"root_status":1,"room_status":1,"apply_count":62,"notify_count":0,"red_point":0},"roomid":1017}
