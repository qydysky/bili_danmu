package part

type AREA_RANK_CHANGED struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ConfID      int    `json:"conf_id"`
		RankName    string `json:"rank_name"`
		UID         int    `json:"uid"`
		Rank        int    `json:"rank"`
		IconURLBlue string `json:"icon_url_blue"`
		IconURLPink string `json:"icon_url_pink"`
		IconURLGrey string `json:"icon_url_grey"`
		ActionType  int    `json:"action_type"`
		Timestamp   int    `json:"timestamp"`
		MsgID       string `json:"msg_id"`
		JumpURLLink string `json:"jump_url_link"`
		JumpURLPc   string `json:"jump_url_pc"`
		JumpURLPink string `json:"jump_url_pink"`
		JumpURLWeb  string `json:"jump_url_web"`
	} `json:"data"`
}

/*
{
    "cmd": "AREA_RANK_CHANGED",
    "data": {
        "conf_id": 22,
        "rank_name": "网游航海",
        "uid": 19738891,
        "rank": 0,
        "icon_url_blue": "https://i0.hdslb.com/bfs/live/18e2990a546d33368200f9058f3d9dbc4038eb5c.png",
        "icon_url_pink": "https://i0.hdslb.com/bfs/live/a6c490c36e88c7b191a04883a5ec15aed187a8f7.png",
        "icon_url_grey": "https://i0.hdslb.com/bfs/live/cb7444b1faf1d785df6265bfdc1fcfc993419b76.png",
        "action_type": 1,
        "timestamp": 1673762370,
        "msg_id": "4477c0ae-c259-4a22-82aa-ff76fa806246",
        "jump_url_link": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=3\u0026ruid=19738891\u0026conf_id=22\u0026is_live_half_webview=1\u0026hybrid_rotate_d=1\u0026is_cling_player=1\u0026hybrid_half_ui=1,3,100p,70p,f4eefa,0,30,100,0,0;2,2,375,100p,f4eefa,0,30,100,0,0;3,3,100p,70p,f4eefa,0,30,100,0,0;4,2,375,100p,f4eefa,0,30,100,0,0;5,3,100p,70p,f4eefa,0,30,100,0,0;6,3,100p,70p,f4eefa,0,30,100,0,0;7,3,100p,70p,f4eefa,0,30,100,0,0;8,3,100p,70p,f4eefa,0,30,100,0,0#/area-rank",
        "jump_url_pc": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=4\u0026ruid=19738891\u0026conf_id=22\u0026pc_ui=338,465,f4eefa,0#/area-rank",
        "jump_url_pink": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=1\u0026ruid=19738891\u0026conf_id=22\u0026is_live_half_webview=1\u0026hybrid_rotate_d=1\u0026is_cling_player=1\u0026hybrid_half_ui=1,3,100p,70p,f4eefa,0,30,100,0,0;2,2,375,100p,f4eefa,0,30,100,0,0;3,3,100p,70p,f4eefa,0,30,100,0,0;4,2,375,100p,f4eefa,0,30,100,0,0;5,3,100p,70p,f4eefa,0,30,100,0,0;6,3,100p,70p,f4eefa,0,30,100,0,0;7,3,100p,70p,f4eefa,0,30,100,0,0;8,3,100p,70p,f4eefa,0,30,100,0,0#/area-rank",
        "jump_url_web": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=2\u0026ruid=19738891\u0026conf_id=22#/area-rank"
    }
}
*/
