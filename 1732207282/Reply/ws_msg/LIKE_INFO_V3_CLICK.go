package part

type LIKE_INFO_V3_CLICK struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ShowArea   int    `json:"show_area"`
		MsgType    int    `json:"msg_type"`
		LikeIcon   string `json:"like_icon"`
		UID        int    `json:"uid"`
		LikeText   string `json:"like_text"`
		Uname      string `json:"uname"`
		UnameColor string `json:"uname_color"`
		Identities []int  `json:"identities"`
		FansMedal  struct {
			TargetID         int    `json:"target_id"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			GuardLevel       int    `json:"guard_level"`
			Special          string `json:"special"`
			IconID           int    `json:"icon_id"`
			AnchorRoomid     int    `json:"anchor_roomid"`
			Score            int    `json:"score"`
		} `json:"fans_medal"`
		ContributionInfo struct {
			Grade int `json:"grade"`
		} `json:"contribution_info"`
	} `json:"data"`
}
