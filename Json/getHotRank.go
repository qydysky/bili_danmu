package part

type GetHotRank struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		List []struct {
			UID            int    `json:"uid"`
			Uname          string `json:"uname"`
			Face           string `json:"face"`
			Rank           int    `json:"rank"`
			Score          int    `json:"score"`
			AreaID         int    `json:"area_id"`
			AreaName       string `json:"area_name"`
			AreaParentID   int    `json:"area_parent_id"`
			AreaParentName string `json:"area_parent_name"`
			RoomID         int    `json:"room_id"`
			LiveStatus     int    `json:"live_status"`
			Verify         int    `json:"verify"`
		} `json:"list"`
		Own struct {
			UID            int    `json:"uid"`
			Uname          string `json:"uname"`
			Face           string `json:"face"`
			Rank           int    `json:"rank"`
			Score          int    `json:"score"`
			NeedScore      int    `json:"need_score"`
			AreaID         int    `json:"area_id"`
			AreaName       string `json:"area_name"`
			AreaParentID   int    `json:"area_parent_id"`
			AreaParentName string `json:"area_parent_name"`
			RoomID         int    `json:"room_id"`
			Verify         int    `json:"verify"`
		} `json:"own"`
		Data struct {
			Countdown    int `json:"countdown"`
			Refresh      int `json:"refresh"`
			IntervalTime int `json:"interval_time"`
			Jumpfrom     int `json:"jumpfrom"`
			BeforeTime   struct {
				Hour   string `json:"hour"`
				Minute string `json:"minute"`
			} `json:"before_time"`
			ParentAreaID   int           `json:"parent_area_id"`
			ParentAreaName string        `json:"parent_area_name"`
			ParentAreaList []interface{} `json:"parent_area_list"`
			SiteAreaID     int           `json:"site_area_id"`
		} `json:"data"`
	} `json:"data"`
}
