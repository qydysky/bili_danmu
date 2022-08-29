package part

type VIDEO_CONNECTION_MSG struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ChannelID   string `json:"channel_id"`
		CurrentTime int    `json:"current_time"`
		Dmscore     int    `json:"dmscore"`
		Toast       string `json:"toast"`
	} `json:"data"`
}

/*
{"cmd":"VIDEO_CONNECTION_MSG","data":{"channel_id":"interact_connection_bbf4a128-c627-49da-b794-a8101635a6dc_22","current_time":1651982364,"dmscore":4,"toast":"主播结束了视频连线"}}
*/
