package part

type LOG_IN_NOTICE struct {
	Cmd  string `json:"cmd"`
	Data struct {
		NoticeMsg string `json:"notice_msg"`
		ImageWeb  string `json:"image_web"`
		ImageApp  string `json:"image_app"`
	} `json:"data"`
}

// {"cmd":"LOG_IN_NOTICE","data":{"notice_msg":"为保护用户隐私，未注册登陆用户将无法查看他人昵称","image_web":"http://i0.hdslb.com/bfs/dm/75e7c16b99208df259fe0a93354fd3440cbab412.png","image_app":""}}