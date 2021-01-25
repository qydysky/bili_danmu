package part

type VTR_GIFT_LOTTERY struct {
	Cmd string `json:"cmd"`
	Data struct {
		Dark_highlight_col int `json:"dark_highlight_col"`
        Highlight_col int `json:"highlight_col"`
        Interact_msg int `json:"interact_msg"`
        Room_id string `json:"room_id"`
    	Toast_msg int `json:"toast_msg"`
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