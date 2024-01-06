package part

type LIKE_GUIDE_USER struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ShowArea   int    `json:"show_area"`
		LikeText   string `json:"like_text"`
		UID        int    `json:"uid"`
		Identities []int  `json:"identities"`
		MsgType    int    `json:"msg_type"`
		Dmscore    int    `json:"dmscore"`
	} `json:"data"`
}
