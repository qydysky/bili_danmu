package part

type GetOnlineGoldRank struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		OnlineNum      int `json:"onlineNum"`
		OnlineRankItem []struct {
			UserRank  int    `json:"userRank"`
			UID       int    `json:"uid"`
			Name      string `json:"name"`
			Face      string `json:"face"`
			Score     int    `json:"score"`
			MedalInfo struct {
				GuardLevel       int    `json:"guardLevel"`
				MedalColorStart  int    `json:"medalColorStart"`
				MedalColorEnd    int    `json:"medalColorEnd"`
				MedalColorBorder int    `json:"medalColorBorder"`
				MedalName        string `json:"medalName"`
				Level            int    `json:"level"`
				TargetID         int    `json:"targetId"`
				IsLight          int    `json:"isLight"`
			} `json:"medalInfo"`
			GuardLevel int `json:"guard_level"`
		} `json:"OnlineRankItem"`
		OwnInfo struct {
			UID        int    `json:"uid"`
			Name       string `json:"name"`
			Face       string `json:"face"`
			Rank       int    `json:"rank"`
			NeedScore  int    `json:"needScore"`
			Score      int    `json:"score"`
			GuardLevel int    `json:"guard_level"`
		} `json:"ownInfo"`
		TipsText  string `json:"tips_text"`
		ValueText string `json:"value_text"`
	} `json:"data"`
}
