package part

type GetHistory struct {
	Code int `json:"code"`
	Data struct {
		Admin []interface{} `json:"admin"`
		Room  []struct {
			Text        string        `json:"text"`
			UID         int           `json:"uid"`
			Nickname    string        `json:"nickname"`
			UnameColor  string        `json:"uname_color"`
			Timeline    string        `json:"timeline"`
			Isadmin     int           `json:"isadmin"`
			Vip         int           `json:"vip"`
			Svip        int           `json:"svip"`
			Medal       []interface{} `json:"medal"`
			Title       []string      `json:"title"`
			UserLevel   []interface{} `json:"user_level"`
			Rank        int           `json:"rank"`
			Teamid      int           `json:"teamid"`
			Rnd         string        `json:"rnd"`
			UserTitle   string        `json:"user_title"`
			GuardLevel  int           `json:"guard_level"`
			Bubble      int           `json:"bubble"`
			BubbleColor string        `json:"bubble_color"`
			CheckInfo   struct {
				Ts int    `json:"ts"`
				Ct string `json:"ct"`
			} `json:"check_info"`
			Lpl int `json:"lpl"`
		} `json:"room"`
	} `json:"data"`
	Message string `json:"message"`
	Msg     string `json:"msg"`
}
