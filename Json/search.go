package part

type Search struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Seid           string `json:"seid"`
		Page           int    `json:"page"`
		Pagesize       int    `json:"pagesize"`
		Numresults     int    `json:"numResults"`
		Numpages       int    `json:"numPages"`
		SuggestKeyword string `json:"suggest_keyword"`
		RqtType        string `json:"rqt_type"`
		CostTime       struct {
			ParamsCheck         string `json:"params_check"`
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
		ExpList interface{} `json:"exp_list"`
		EggHit  int         `json:"egg_hit"`
		Result  []struct {
			RankOffset int      `json:"rank_offset"`
			UID        int      `json:"uid"`
			Tags       string   `json:"tags"`
			Type       string   `json:"type"`
			LiveTime   string   `json:"live_time"`
			HitColumns []string `json:"hit_columns"`
			LiveStatus int      `json:"live_status"`
			Area       int      `json:"area"`
			IsLive     bool     `json:"is_live"`
			Uname      string   `json:"uname"`
			Uface      string   `json:"uface"`
			RankIndex  int      `json:"rank_index"`
			RankScore  int      `json:"rank_score"`
			Roomid     int      `json:"roomid"`
			Attentions int      `json:"attentions"`
		} `json:"result"`
		ShowColumn int `json:"show_column"`
	} `json:"data"`
}