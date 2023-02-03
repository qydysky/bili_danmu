package part

type VTR_GIFT_LOTTERY struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ActName          string `json:"act_name"`
		AwardUsername    string `json:"award_username"`
		InteractMsg      string `json:"interact_msg"`
		ToastMsg         string `json:"toast_msg"`
		RoomID           int    `json:"room_id"`
		UID              int    `json:"uid"`
		HighlightCol     string `json:"highlight_col"`
		DarkHighlightCol string `json:"dark_highlight_col"`
		LotteryID        string `json:"lottery_id"`
	} `json:"data"`
}

/*
{
    "cmd": "VTR_GIFT_LOTTERY",
    "data": {
        "dark_highlight_col": "#000000",
        "highlight_col": "#000000",
        "interact_msg": "恭喜和铜肆年赠送银色闪耀魔盒触发银色闪耀魔盒！+活动积分100分！",
        "room_id": 5050,
        "toast_msg": ""
    }
}
*/
