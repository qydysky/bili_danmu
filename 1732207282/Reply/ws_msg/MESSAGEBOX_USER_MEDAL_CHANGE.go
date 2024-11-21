package part

type MESSAGEBOX_USER_MEDAL_CHANGE struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Type                int    `json:"type"`
		Uid                 int    `json:"uid"`
		Up_uid              int    `json:"up_uid"`
		Medal_name          string `json:"medal_name"`
		Medal_level         int    `json:"medal_level"`
		Medal_color_start   int    `json:"medal_color_start"`
		Medal_color_end     int    `json:"medal_color_end"`
		Medal_color_border  int    `json:"medal_color_border"`
		Guard_level         int    `json:"guard_level"`
		Is_lighted          int    `json:"is_lighted"`
		Unlock              int    `json:"unlock"`
		Unlock_level        int    `json:"unlock_level"`
		Multi_unlock_level  string `json:"multi_unlock_level"`
		Upper_bound_content string `json:"upper_bound_content"`
	} `json:"data"`
}
