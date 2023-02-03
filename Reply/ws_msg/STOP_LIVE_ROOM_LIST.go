package part

type STOP_LIVE_ROOM_LIST struct {
	Cmd  string `json:"cmd"`
	Data struct {
		RoomIDList []int `json:"room_id_list"`
	} `json:"data"`
}

//{"cmd":"STOP_LIVE_ROOM_LIST","data":{"room_id_list":[22301508,14091554,21782859]}}
