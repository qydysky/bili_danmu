package part

type FULL_SCREEN_MASK_OPEN struct {
	Cmd  string `json:"cmd"`
	Data struct {
		HasMask       bool   `json:"has_mask"`
		Icon          string `json:"icon"`
		Lines         string `json:"lines"`
		OverlaySecond int    `json:"overlay_second"`
		RoomID        int    `json:"room_id"`
		Title         string `json:"title"`
	} `json:"data"`
}
