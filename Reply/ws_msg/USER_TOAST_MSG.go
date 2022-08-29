package part

type USER_TOAST_MSG struct {
	Cmd  string `json:"cmd"`
	Data struct {
		AnchorShow       bool   `json:"anchor_show"`
		Color            string `json:"color"`
		Dmscore          int    `json:"dmscore"`
		EndTime          int    `json:"end_time"`
		GuardLevel       int    `json:"guard_level"`
		IsShow           int    `json:"is_show"`
		Num              int    `json:"num"`
		OpType           int    `json:"op_type"`
		PayflowID        string `json:"payflow_id"`
		Price            int    `json:"price"`
		RoleName         string `json:"role_name"`
		StartTime        int    `json:"start_time"`
		SvgaBlock        int    `json:"svga_block"`
		TargetGuardCount int    `json:"target_guard_count"`
		ToastMsg         string `json:"toast_msg"`
		UID              int    `json:"uid"`
		Unit             string `json:"unit"`
		UserShow         bool   `json:"user_show"`
		Username         string `json:"username"`
	} `json:"data"`
}
/*
{"cmd":"USER_TOAST_MSG","data":{"anchor_show":true,"color":"#00D1F1","dmscore":90,"end_time":1623612866,"guard_level":3,"is_show":0,"num":1,
"op_type":3,"payflow_id":"2106140334131392145517282","price":138000,"role_name":"舰长","start_time":1623612866,"svga_block":0,"target_guard_
count":820,"toast_msg":"\u003c%大桶麦丽素%\u003e 自动续费了舰长","uid":40144551,"unit":"月","user_show":true,"username":"大桶麦丽素"}}
*/