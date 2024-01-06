package part

type GetRoomBaseInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		ByRoomIds map[string]GetRoomBaseInfoD `json:"by_room_ids"`
	} `json:"data"`
}

type GetRoomBaseInfoD struct {
	RoomID         int    `json:"room_id"`
	UID            int    `json:"uid"`
	AreaID         int    `json:"area_id"`
	LiveStatus     int    `json:"live_status"`
	// LiveURL        string `json:"live_url"`
	ParentAreaID   int    `json:"parent_area_id"`
	Title          string `json:"title"`
	// ParentAreaName string `json:"parent_area_name"`
	// AreaName       string `json:"area_name"`
	LiveTime       string `json:"live_time"`
	// Description    string `json:"description"`
	// Tags           string `json:"tags"`
	Attention      int    `json:"attention"`
	Online         int    `json:"online"`
	ShortID        int    `json:"short_id"`
	Uname          string `json:"uname"`
	Cover          string `json:"cover"`
	Background     string `json:"background"`
	JoinSlide      int    `json:"join_slide"`
	LiveID         int64  `json:"live_id"`
	LiveIDStr      string `json:"live_id_str"`
}
