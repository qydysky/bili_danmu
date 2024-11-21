package part

type POPULARITY_RED_POCKET_START struct {
	Cmd  string `json:"cmd"`
	Data struct {
		LotID           int    `json:"lot_id"`
		SenderUID       int    `json:"sender_uid"`
		SenderName      string `json:"sender_name"`
		SenderFace      string `json:"sender_face"`
		JoinRequirement int    `json:"join_requirement"`
		Danmu           string `json:"danmu"`
		CurrentTime     int    `json:"current_time"`
		StartTime       int    `json:"start_time"`
		EndTime         int    `json:"end_time"`
		LastTime        int    `json:"last_time"`
		RemoveTime      int    `json:"remove_time"`
		ReplaceTime     int    `json:"replace_time"`
		LotStatus       int    `json:"lot_status"`
		H5URL           string `json:"h5_url"`
		UserStatus      int    `json:"user_status"`
		Awards          []struct {
			GiftID   int    `json:"gift_id"`
			GiftName string `json:"gift_name"`
			GiftPic  string `json:"gift_pic"`
			Num      int    `json:"num"`
		} `json:"awards"`
		LotConfigID int `json:"lot_config_id"`
		TotalPrice  int `json:"total_price"`
		WaitNum     int `json:"wait_num"`
	} `json:"data"`
}

/*
{"cmd":"POPULARITY_RED_POCKET_START","data":{"lot_id":2495708,"sender_uid":2128651899,"sender_name":"一个不想起床的屑","sender_face":"http://i0.hdslb.com/bfs/face/e32a3b04c1b60f1070bc22536611525aef3af729.jpg","join_requirement":1,"danmu":"老板大气！点点红包抽礼物！","current_time":1652008432,"start_time":1652008432,"end_time":1652008612,"last_time":180,"remove_time":1652008627,"replace_time":1652008622,"lot_status":1,"h5_url":"https://live.bilibili.com/p/html/live-app-red-envelope/popularity.html?is_live_half_webview=1\u0026hybrid_half_ui=1,5,100p,100p,000000,0,50,0,0,1;2,5,100p,100p,000000,0,50,0,0,1;3,5,100p,100p,000000,0,50,0,0,1;4,5,100p,100p,000000,0,50,0,0,1;5,5,100p,100p,000000,0,50,0,0,1;6,5,100p,100p,000000,0,50,0,0,1;7,5,100p,100p,000000,0,50,0,0,1;8,5,100p,100p,000000,0,50,0,0,1\u0026hybrid_rotate_d=1\u0026hybrid_biz=popularityRedPacket\u0026lotteryId=2495708","user_status":2,"awards":[{"gift_id":31215,"gift_name":"花式夸夸","gift_pic":"https://s1.hdslb.com/bfs/live/ce0efeceae7054d1ee835864eace28f08a54a37d.png","num":1},{"gift_id":31212,"gift_name":"打call","gift_pic":"https://s1.hdslb.com/bfs/live/f75291a0e267425c41e1ce31b5ffd6bfedc6f0b6.png","num":12},{"gift_id":31214,"gift_name":"牛哇","gift_pic":"https://s1.hdslb.com/bfs/live/b8a38b4bd3be120becddfb92650786f00dffad48.png","num":26}],"lot_config_id":5,"total_price":41600,"wait_num":2}}
*/
