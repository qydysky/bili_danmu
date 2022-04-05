package part

type HOT_RANK_SETTLEMENT_V2 struct {
	Cmd  string                      `json:"cmd"`
	Data HOT_RANK_SETTLEMENT_V2_Data `json:"data"`
}
type HOT_RANK_SETTLEMENT_V2_Data struct {
	Rank      int    `json:"rank"`
	Uname     string `json:"uname"`
	Face      string `json:"face"`
	Timestamp int    `json:"timestamp"`
	Icon      string `json:"icon"`
	AreaName  string `json:"area_name"`
	URL       string `json:"url"`
	CacheKey  string `json:"cache_key"`
	DmMsg     string `json:"dm_msg"`
}
