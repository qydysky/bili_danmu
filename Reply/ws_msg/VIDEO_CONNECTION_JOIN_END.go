package part

type VIDEO_CONNECTION_JOIN_END struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ChannelID   string `json:"channel_id"`
		StartAt     int    `json:"start_at"`
		Toast       string `json:"toast"`
		CurrentTime int    `json:"current_time"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

/*
{"cmd":"VIDEO_CONNECTION_JOIN_END","data":{"channel_id":"interact_connection_bbf4a128-c627-49da-b794-a8101635a6dc_22","start_at":1651982364,"toast":"主播结束了与帅soserious的连线.","current_time":1651982364},"roomid":1017}
*/
