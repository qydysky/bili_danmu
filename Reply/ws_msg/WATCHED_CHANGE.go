package part

type WATCHED_CHANGE struct {
	Cmd  string              `json:"cmd"`
	Data WATCHED_CHANGE_Data `json:"data"`
}
type WATCHED_CHANGE_Data struct {
	Num       int    `json:"num"`
	TextSmall string `json:"text_small"`
	TextLarge string `json:"text_large"`
}
