package part

type Nav struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		IsLogin bool `json:"isLogin"`
	} `json:"data"`
}
