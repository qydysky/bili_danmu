package part

type Following struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	TTL     int           `json:"ttl"`
	Data    FollowingData `json:"data"`
}
type FollowingDataList struct {
	Roomid           int    `json:"roomid"`
	UID              int    `json:"uid"`
	Uname            string `json:"uname"`
	Title            string `json:"title"`
	Face             string `json:"face"`
	LiveStatus       int    `json:"live_status"`
	RecordNum        int    `json:"record_num"`
	RecentRecordID   string `json:"recent_record_id"`
	IsAttention      int    `json:"is_attention"`
	Clipnum          int    `json:"clipnum"`
	FansNum          int    `json:"fans_num"`
	AreaName         string `json:"area_name"`
	AreaValue        string `json:"area_value"`
	Tags             string `json:"tags"`
	RecentRecordIDV2 string `json:"recent_record_id_v2"`
	RecordNumV2      int    `json:"record_num_v2"`
}
type FollowingData struct {
	Title     string              `json:"title"`
	PageSize  int                 `json:"pageSize"`
	TotalPage int                 `json:"totalPage"`
	List      []FollowingDataList `json:"list"`
	Count     int                 `json:"count"`
}
