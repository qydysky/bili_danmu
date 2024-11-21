package part

type WARNING struct {
	Cmd    string `json:"cmd"`
	Msg    string `json:"msg"`
	Roomid int    `json:"roomid"`
}

/*
{"cmd":"WARNING","msg":"\u865a\u62df\u5f62\u8c61\u4e0d\u9002\u5b9c\uff0c\u8bf7\u7acb\u5373\u8c03\u6574","roomid":22259479}
*/
