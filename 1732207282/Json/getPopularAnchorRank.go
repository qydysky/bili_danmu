package part

type GetPopularAnchorRank struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		List []struct {
			UID             int    `json:"uid"`
			Uname           string `json:"uname"`
			Face            string `json:"face"`
			Rank            int    `json:"rank"`
			Score           int    `json:"score"`
			RoomID          int    `json:"room_id"`
			LiveStatus      int    `json:"live_status"`
			Verify          int    `json:"verify"`
			UserNum         int    `json:"user_num"`
			LotStatus       int    `json:"lot_status"`
			RedPocketStatus int    `json:"red_pocket_status"`
			RoomLink        string `json:"room_link"`
		} `json:"list"`
		Anchor struct {
			UID               int    `json:"uid"`
			Uname             string `json:"uname"`
			Face              string `json:"face"`
			Rank              int    `json:"rank"`
			Score             int    `json:"score"`
			RankDistanceScore int    `json:"rank_distance_score"`
			RoomID            int    `json:"room_id"`
			Verify            int    `json:"verify"`
			UserNum           int    `json:"user_num"`
			FansClubStatus    int    `json:"fans_club_status"`
			InBlack           int    `json:"in_black"`
		} `json:"anchor"`
		UserMedal struct {
			MedalID          int    `json:"medal_id"`
			Level            int    `json:"level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLight          int    `json:"is_light"`
			GuardLevel       int    `json:"guard_level"`
			GuardIcon        string `json:"guard_icon"`
			HonorIcon        string `json:"honor_icon"`
		} `json:"user_medal"`
		Data struct {
			Countdown    int `json:"countdown"`
			Refresh      int `json:"refresh"`
			IntervalTime int `json:"interval_time"`
			Jumpfrom     int `json:"jumpfrom"`
		} `json:"data"`
	} `json:"data"`
}
