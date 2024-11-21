package part

type LIVE_INTERACTIVE_GAME struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Type           int         `json:"type"`
		UID            int         `json:"uid"`
		Uname          string      `json:"uname"`
		Uface          string      `json:"uface"`
		GiftID         int         `json:"gift_id"`
		GiftName       string      `json:"gift_name"`
		GiftNum        int         `json:"gift_num"`
		Price          int         `json:"price"`
		Paid           bool        `json:"paid"`
		Msg            string      `json:"msg"`
		FansMedalLevel int         `json:"fans_medal_level"`
		GuardLevel     int         `json:"guard_level"`
		Timestamp      int         `json:"timestamp"`
		AnchorLottery  interface{} `json:"anchor_lottery"`
		PkInfo         interface{} `json:"pk_info"`
		AnchorInfo     interface{} `json:"anchor_info"`
	} `json:"data"`
}
