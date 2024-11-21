package part

type LITTLE_TIPS struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Msg string `json:"msg"`
	} `json:"data"`
}
