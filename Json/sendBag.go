package part

type SendBag struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		UID               int    `json:"uid"`
		Uname             string `json:"uname"`
		Face              string `json:"face"`
		GuardLevel        int    `json:"guard_level"`
		Ruid              int    `json:"ruid"`
		RoomID            int    `json:"room_id"`
		Rcost             int    `json:"rcost"`
		TotalCoin         int    `json:"total_coin"`
		PayCoin           int    `json:"pay_coin"`
		BlowSwitch        int    `json:"blow_switch"`
		SendTips          string `json:"send_tips"`
		DiscountID        int    `json:"discount_id"`
		SendMaster        any    `json:"send_master"`
		ButtonComboType   int    `json:"button_combo_type"`
		SendGiftCountdown int    `json:"send_gift_countdown"`
		BlindGift         any    `json:"blind_gift"`
		Fulltext          string `json:"fulltext"`
		CritProb          int    `json:"crit_prob"`
		Price             int    `json:"price"`
		LeftNum           int    `json:"left_num"`
		NeedNum           int    `json:"need_num"`
		AvailableNum      int    `json:"available_num"`
		BpCentBalance     int    `json:"bp_cent_balance"`
		GiftList          []struct {
			Tid         string `json:"tid"`
			GiftID      int    `json:"gift_id"`
			GiftType    int    `json:"gift_type"`
			GiftName    string `json:"gift_name"`
			GiftNum     int    `json:"gift_num"`
			GiftAction  string `json:"gift_action"`
			GiftPrice   int    `json:"gift_price"`
			CoinType    string `json:"coin_type"`
			TagImage    string `json:"tag_image"`
			EffectBlock int    `json:"effect_block"`
			Extra       struct {
				Wallet  any `json:"wallet"`
				GiftBag struct {
					BagID   int `json:"bag_id"`
					GiftNum int `json:"gift_num"`
				} `json:"gift_bag"`
				Pk struct {
					PkGiftTips string `json:"pk_gift_tips"`
				} `json:"pk"`
				LotteryID string `json:"lottery_id"`
				Medal     struct {
					New       int    `json:"new"`
					MedalID   int    `json:"medal_id"`
					MedalName string `json:"medal_name"`
					Level     int    `json:"level"`
				} `json:"medal"`
			} `json:"extra"`
			GiftEffect struct {
				ComboTimeout      int    `json:"combo_timeout"`
				SuperGiftNum      int    `json:"super_gift_num"`
				SuperBatchGiftNum int    `json:"super_batch_gift_num"`
				BatchComboID      string `json:"batch_combo_id"`
				ComboID           string `json:"combo_id"`
			} `json:"gift_effect"`
			IsSpecialBatch    int  `json:"is_special_batch"`
			ComboStayTime     int  `json:"combo_stay_time"`
			ComboTotalCoin    int  `json:"combo_total_coin"`
			Demarcation       int  `json:"demarcation"`
			Magnification     int  `json:"magnification"`
			ComboResourcesID  int  `json:"combo_resources_id"`
			FloatScResourceID int  `json:"float_sc_resource_id"`
			IsNaming          bool `json:"is_naming"`
			ReceiveUserInfo   struct {
				Uname string `json:"uname"`
				UID   int    `json:"uid"`
			} `json:"receive_user_info"`
			IsJoinReceiver bool `json:"is_join_receiver"`
		} `json:"gift_list"`
		SendID string `json:"send_id"`
	} `json:"data"`
}
