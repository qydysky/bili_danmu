package part

type HOT_RANK_CHANGED struct {
	Cmd string `json:"cmd"`
	Data struct {
		Rank int `json:"rank"`
		Trend int `json:"trend"`
		Countdown string `json:"countdown"`
		Timestamp string `json:"timestamp"`
		Web_url string `json:"web_url"`
		Live_url string `json:"live_url"`
		Blink_url string `json:"blink_url"`
		Live_link_url string `json:"live_link_url"`
		Pc_link_url string `json:"pc_link_url"`
		Icon string `json:"icon"`
		Area_name string `json:"area_name"`
	} `json:"data"`
}
