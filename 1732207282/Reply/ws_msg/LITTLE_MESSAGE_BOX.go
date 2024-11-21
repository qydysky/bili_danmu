package part

type LITTLE_MESSAGE_BOX struct {
	Cmd  string `json:"cmd"`
	Data struct {
		From     string `json:"from"`
		Platform struct {
			Android bool `json:"android"`
			Ios     bool `json:"ios"`
			Web     bool `json:"web"`
		} `json:"platform"`
		Msg     string `json:"msg"`
		Room_id int    `json:"room_id"`
		Type    int    `json:"type"`
	} `json:"data"`
}
