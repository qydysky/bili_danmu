package part

type ROOM_CHANGE struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Title          string `json:"title"`
		AreaID         int    `json:"area_id"`
		ParentAreaID   int    `json:"parent_area_id"`
		AreaName       string `json:"area_name"`
		ParentAreaName string `json:"parent_area_name"`
		LiveKey        string `json:"live_key"`
		SubSessionKey  string `json:"sub_session_key"`
	} `json:"data"`
}
