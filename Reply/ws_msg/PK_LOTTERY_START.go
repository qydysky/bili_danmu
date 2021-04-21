package part

type PK_LOTTERY_START struct {
	Cmd  string `json:"cmd"`
	Data struct {
		AssetAnimationPic string `json:"asset_animation_pic"`
		AssetIcon         string `json:"asset_icon"`
		FromUser          struct {
			Face  string `json:"face"`
			UID   int    `json:"uid"`
			Uname string `json:"uname"`
		} `json:"from_user"`
		ID        int    `json:"id"`
		MaxTime   int    `json:"max_time"`
		PkID      int    `json:"pk_id"`
		RoomID    int    `json:"room_id"`
		ThankText string `json:"thank_text"`
		Time      int    `json:"time"`
		TimeWait  int    `json:"time_wait"`
		Title     string `json:"title"`
		Weight    int    `json:"weight"`
	} `json:"data"`
}
