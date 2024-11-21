package part

type POPULAR_RANK_CHANGED struct {
	Cmd  string `json:"cmd"`
	Data struct {
		UID       int    `json:"uid"`
		Rank      int    `json:"rank"`
		Countdown int    `json:"countdown"`
		Timestamp int    `json:"timestamp"`
		CacheKey  string `json:"cache_key"`
	} `json:"data"`
}

// {"cmd":"POPULAR_RANK_CHANGED","data":{"uid":13046,"rank":77,"countdown":1421,"timestamp":1672662980,"cache_key":"rank_change:e759ceed19234dbd9517829adb9b0b6c"}}
