package part

type RANK_CHANGED struct {
	Cmd  string `json:"cmd"`
	Data struct {
		UID              int    `json:"uid"`
		Rank             int    `json:"rank"`
		Countdown        int    `json:"countdown"`
		Timestamp        int    `json:"timestamp"`
		OnRankNameByType string `json:"on_rank_name_by_type"`
		RankNameByType   string `json:"rank_name_by_type"`
		URLByType        string `json:"url_by_type"`
		RankByType       int    `json:"rank_by_type"`
		RankType         int    `json:"rank_type"`
	} `json:"data"`
}
