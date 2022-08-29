package part

type POPULARITY_RED_POCKET_NEW struct {
	Cmd  string `json:"cmd"`
	Data struct {
		LotID       int    `json:"lot_id"`
		StartTime   int    `json:"start_time"`
		CurrentTime int    `json:"current_time"`
		WaitNum     int    `json:"wait_num"`
		Uname       string `json:"uname"`
		UID         int    `json:"uid"`
		Action      string `json:"action"`
		Num         int    `json:"num"`
		GiftName    string `json:"gift_name"`
		GiftID      int    `json:"gift_id"`
		Price       int    `json:"price"`
		NameColor   string `json:"name_color"`
		MedalInfo   struct {
			TargetID         int    `json:"target_id"`
			Special          string `json:"special"`
			IconID           int    `json:"icon_id"`
			AnchorUname      string `json:"anchor_uname"`
			AnchorRoomid     int    `json:"anchor_roomid"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			GuardLevel       int    `json:"guard_level"`
		} `json:"medal_info"`
	} `json:"data"`
}

/*
{"cmd":"POPULARITY_RED_POCKET_NEW","data":{"lot_id":2495925,"start_time":1652009382,"current_time":1652008665,"wait_num":5,"uname":"一个不想起床的屑","uid":2128651899,"action":"送出","num":1,"gift_name":"红包","gift_id":13000,"price":520,"name_color":"","medal_info":{"target_id":168598,"special":"","icon_id":0,"anchor_uname":"","anchor_roomid":0,"medal_level":4,"medal_name":"刺儿","medal_color":6067854,"medal_color_start":6067854,"medal_color_end":6067854,"medal_color_border":6067854,"is_lighted":1,"guard_level":0}}}
*/
