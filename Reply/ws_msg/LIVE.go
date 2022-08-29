package part

type LIVE struct {
	Cmd             string `json:"cmd"`
	LiveKey         string `json:"live_key"`
	VoiceBackground string `json:"voice_background"`
	SubSessionKey   string `json:"sub_session_key"`
	LivePlatform    string `json:"live_platform"`
	LiveModel       int    `json:"live_model"`
	LiveTime        int    `json:"live_time"`
	Roomid          int    `json:"roomid"`
}

/*
{"cmd":"LIVE","live_key":"243098417424107244","voice_background":"","sub_session_key":"243098417424107244sub_time:1652355679","live_platform":"pc_link","live_model":0,"live_time":1652355679,"roomid":394988}
*/
