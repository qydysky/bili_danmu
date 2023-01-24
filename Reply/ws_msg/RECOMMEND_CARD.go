package part

type RECOMMEND_CARD struct {
	Cmd  string `json:"cmd"`
	Data struct {
		TitleIcon     string `json:"title_icon"`
		RecommendList []struct {
			ShoppingCardDetail struct {
				GoodsID             string        `json:"goods_id"`
				GoodsName           string        `json:"goods_name"`
				GoodsPrice          string        `json:"goods_price"`
				GoodsMaxPrice       string        `json:"goods_max_price"`
				SaleStatus          int           `json:"sale_status"`
				CouponName          string        `json:"coupon_name"`
				GoodsIcon           string        `json:"goods_icon"`
				GoodsStatus         int           `json:"goods_status"`
				Source              int           `json:"source"`
				H5URL               string        `json:"h5_url"`
				JumpLink            string        `json:"jump_link"`
				SchemaURL           string        `json:"schema_url"`
				IsPreSale           int           `json:"is_pre_sale"`
				ActivityInfo        interface{}   `json:"activity_info"`
				PreSaleInfo         interface{}   `json:"pre_sale_info"`
				EarlyBirdInfo       interface{}   `json:"early_bird_info"`
				Timestamp           int           `json:"timestamp"`
				CouponDiscountPrice string        `json:"coupon_discount_price"`
				SellingPoint        string        `json:"selling_point"`
				HotBuyNum           int           `json:"hot_buy_num"`
				GiftBuyInfo         []interface{} `json:"gift_buy_info"`
				IsExclusive         bool          `json:"is_exclusive"`
				CouponID            string        `json:"coupon_id"`
				RewardInfo          interface{}   `json:"reward_info"`
				GoodsTagList        interface{}   `json:"goods_tag_list"`
				VirtualExtraInfo    interface{}   `json:"virtual_extra_info"`
				PriceInfo           interface{}   `json:"price_info"`
				BtnInfo             interface{}   `json:"btn_info"`
				GoodsSortID         int           `json:"goods_sort_id"`
			} `json:"shopping_card_detail"`
			RecommendCardExtra interface{} `json:"recommend_card_extra"`
		} `json:"recommend_list"`
		Timestamp int `json:"timestamp"`
	} `json:"data"`
}
