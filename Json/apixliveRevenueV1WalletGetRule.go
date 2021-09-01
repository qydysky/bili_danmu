package part

type ApixliveRevenueV1WalletGetRule struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Coin2SilverRateNormal int     `json:"coin_2_silver_rate_normal"`
		Coin2SilverRateVip    int     `json:"coin_2_silver_rate_vip"`
		Coin2SilverRate       int     `json:"coin_2_silver_rate"`
		Coin2SilverFee        float64 `json:"coin_2_silver_fee"`
		Coin2SilverLimit      int     `json:"coin_2_silver_limit"`
		Coin2SilverLimitVip   int     `json:"coin_2_silver_limit_vip"`
		Silver2CoinPrice      int     `json:"silver_2_coin_price"`
		Silver2CoinLimit      int     `json:"silver_2_coin_limit"`
		Coin2SilverRealRate   int     `json:"coin_2_silver_real_rate"`
		Gold2SilverBonus      struct {
			Num10000  float64 `json:"10000"`
			Num100000 float64 `json:"100000"`
			Num500000 float64 `json:"500000"`
		} `json:"gold_2_silver_bonus"`
	} `json:"data"`
}