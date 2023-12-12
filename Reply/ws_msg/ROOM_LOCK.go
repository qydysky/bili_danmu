package part

type ROOM_LOCK struct {
	Cmd    string `json:"cmd"`
	Expire string `json:"expire"`
	Roomid int    `json:"roomid"`
}
