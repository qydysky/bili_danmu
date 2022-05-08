package part

type COMMON_NOTICE_DANMAKU struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ContentSegments []struct {
			FontColor string `json:"font_color"`
			Text      string `json:"text"`
			Type      int    `json:"type"`
		} `json:"content_segments"`
		Dmscore   int   `json:"dmscore"`
		Terminals []int `json:"terminals"`
	} `json:"data"`
}

/*
{"cmd":"COMMON_NOTICE_DANMAKU","data":{"content_segments":[{"font_color":"#FB7299","text":"大肚罐罐子在元气赏中五连抽！送出了好多礼物！","type":1}],"dmscore":144,"terminals":[1,2,3,4,5]}}
*/
