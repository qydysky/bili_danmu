package part

type Search struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Seid           string `json:"seid"`
		Page           int    `json:"page"`
		Pagesize       int    `json:"pagesize"`
		NumResults     int    `json:"numResults"`
		NumPages       int    `json:"numPages"`
		SuggestKeyword string `json:"suggest_keyword"`
		RqtType        string `json:"rqt_type"`
		CostTime       struct {
			ParamsCheck         string `json:"params_check"`
			IsRiskQuery         string `json:"is_risk_query"`
			IllegalHandler      string `json:"illegal_handler"`
			DeserializeResponse string `json:"deserialize_response"`
			AsResponseFormat    string `json:"as_response_format"`
			AsRequest           string `json:"as_request"`
			SaveCache           string `json:"save_cache"`
			AsDocRequest        string `json:"as_doc_request"`
			AsRequestFormat     string `json:"as_request_format"`
			Total               string `json:"total"`
			MainHandler         string `json:"main_handler"`
		} `json:"cost_time"`
		ExpList struct {
			Num5501 bool `json:"5501"`
			Num6609 bool `json:"6609"`
			Num7708 bool `json:"7708"`
		} `json:"exp_list"`
		EggHit int `json:"egg_hit"`
		Result []struct {
			RankOffset int      `json:"rank_offset"`
			UID        int      `json:"uid"`
			Tags       string   `json:"tags"`
			Type       string   `json:"type"`
			LiveTime   string   `json:"live_time"`
			HitColumns []string `json:"hit_columns"`
			CateName   string   `json:"cate_name"`
			LiveStatus int      `json:"live_status"`
			Area       int      `json:"area"`
			IsLive     bool     `json:"is_live"`
			Uname      string   `json:"uname"`
			AreaV2ID   int      `json:"area_v2_id"`
			Uface      string   `json:"uface"`
			RankIndex  int      `json:"rank_index"`
			RankScore  int      `json:"rank_score"`
			Roomid     int      `json:"roomid"`
			Attentions int      `json:"attentions"`
		} `json:"result"`
		ShowColumn int `json:"show_column"`
		InBlackKey int `json:"in_black_key"`
		InWhiteKey int `json:"in_white_key"`
	} `json:"data"`
}
