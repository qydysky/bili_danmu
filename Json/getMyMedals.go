package part

type GetMyMedals struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	TTL     int              `json:"ttl"`
	Data    GetMyMedals_Data `json:"data"`
}
type GetMyMedals_Items struct {
	CanDeleted       bool   `json:"can_deleted"`
	DayLimit         int    `json:"day_limit"`
	GuardLevel       int    `json:"guard_level"`
	GuardMedalTitle  string `json:"guard_medal_title"`
	Intimacy         int    `json:"intimacy"`
	IsLighted        int    `json:"is_lighted"`
	Level            int    `json:"level"`
	MedalName        string `json:"medal_name"`
	MedalColorBorder int    `json:"medal_color_border"`
	MedalColorEnd    int    `json:"medal_color_end"`
	MedalColorStart  int    `json:"medal_color_start"`
	MedalID          int    `json:"medal_id"`
	NextIntimacy     int    `json:"next_intimacy"`
	TodayFeed        int    `json:"today_feed"`
	Roomid           int    `json:"roomid"`
	Status           int    `json:"status"`
	TargetID         int    `json:"target_id"`
	TargetName       string `json:"target_name"`
	Uname            string `json:"uname"`
}
type GetMyMedals_PageInfo struct {
	CurPage   int `json:"cur_page"`
	TotalPage int `json:"total_page"`
}
type GetMyMedals_Data struct {
	Items    []GetMyMedals_Items  `json:"items"`
	PageInfo GetMyMedals_PageInfo `json:"page_info"`
	Count    int                  `json:"count"`
}
