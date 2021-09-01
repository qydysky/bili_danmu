package part

type ApiXliveRevenueV1WalletGetStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Silver          int `json:"silver"`
		Gold            int `json:"gold"`
		Coin            int `json:"coin"`
		Bp              int `json:"bp"`
		Coin2SilverLeft int `json:"coin_2_silver_left"`
		Silver2CoinLeft int `json:"silver_2_coin_left"`
		Num             int `json:"num"`
		Status          int `json:"status"`
		Vip             int `json:"vip"`
	} `json:"data"`
}