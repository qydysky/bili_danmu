package part

type WEALTH_NOTIFY struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Flag int `json:"flag"`
		Info struct {
			EffectKey       int   `json:"effect_key"`
			HasItemsChanged int   `json:"has_items_changed"`
			Level           int   `json:"level"`
			SendTime        int64 `json:"send_time"`
			Status          int   `json:"status"`
		} `json:"info"`
	} `json:"data"`
	IsReport bool   `json:"is_report"`
	MsgID    string `json:"msg_id"`
	SendTime int64  `json:"send_time"`
}
