package part

type VIDEO_CONNECTION_JOIN_START struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Status       int    `json:"status"`
		InvitedUID   int    `json:"invited_uid"`
		ChannelID    string `json:"channel_id"`
		InvitedUname string `json:"invited_uname"`
		InvitedFace  string `json:"invited_face"`
		StartAt      int    `json:"start_at"`
		CurrentTime  int    `json:"current_time"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

/*
{"cmd":"VIDEO_CONNECTION_JOIN_START","data":{"status":1,"invited_uid":66391032,"channel_id":"interact_connection_bbf4a128-c627-49da-b794-a8101635a6dc_22","invited_uname":"帅soserious","invited_face":"http://i0.hdslb.com/bfs/face/40a663bb18e9064a97901b96aaf7d84d8056e98b.jpg","start_at":1651980743,"current_time":1651980743},"roomid":2823677}
*/
