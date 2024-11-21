package part

type GIFT_STAR_PROCESS struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Status int    `json:"status"`
		Tip    string `json:"tip"`
	} `json:"data"`
}
