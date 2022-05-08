package part

type ACTIVITY_BANNER_CHANGE_V2 struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Timestamp int `json:"timestamp"`
		List      []struct {
			ID            int    `json:"id"`
			Position      string `json:"position"`
			Type          int    `json:"type"`
			ActivityTitle string `json:"activity_title"`
			Cover         string `json:"cover"`
			JumpURL       string `json:"jump_url"`
			IsClose       int    `json:"is_close"`
			Action        string `json:"action"`
			PlatformInfo  []struct {
				Platform  string `json:"platform"`
				Condition int    `json:"condition"`
				Build     int    `json:"build"`
			} `json:"platform_info"`
			ExtData string `json:"ext_data"`
		} `json:"list"`
	} `json:"data"`
}

/*
{"cmd":"ACTIVITY_BANNER_CHANGE_V2","data":{"timestamp":1651982404,"list":[{"id":1739,"position":"bottom","type":0,"activity_title":"宅家快乐燃脂","cover":"https://i0.hdslb.com/bfs/live/5fe7003e31f6a2adbdbdbc87d263cdfa3b34c858.jpg","jump_url":"https://www.bilibili.com/blackboard/activity-RmDTqgMcOa.html?-Abrowser=live\u0026is_live_half_webview=1\u0026hybrid_rotate_d=1\u0026is_cling_player=1\u0026hybrid_half_ui=1,3,100p,70p,0,1,30,100;2,2,375,100p,0,1,30,100;3,3,100p,70p,0,1,30,100;4,2,375,100p,0,1,30,100;5,3,100p,70p,0,1,30,100;6,3,100p,70p,0,1,30,100;7,3,100p,70p,0,1,30,100;8,3,100p,70p,0,1,30,100","is_close":1,"action":"delete","platform_info":[{"platform":"android","condition":0,"build":0},{"platform":"ios","condition":0,"build":0}],"ext_data":""}]}}
*/
