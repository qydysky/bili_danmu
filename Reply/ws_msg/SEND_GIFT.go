package part

type SEND_GIFT struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Action         string `json:"action"`
		BatchComboID   string `json:"batch_combo_id"`
		BatchComboSend struct {
			Action        string      `json:"action"`
			BatchComboID  string      `json:"batch_combo_id"`
			BatchComboNum int         `json:"batch_combo_num"`
			BlindGift     interface{} `json:"blind_gift"`
			GiftID        int         `json:"gift_id"`
			GiftName      string      `json:"gift_name"`
			GiftNum       int         `json:"gift_num"`
			SendMaster    interface{} `json:"send_master"`
			UID           int         `json:"uid"`
			Uname         string      `json:"uname"`
		} `json:"batch_combo_send"`
		Beatid           string      `json:"beatId"`
		BizSource        string      `json:"biz_source"`
		BlindGift        interface{} `json:"blind_gift"`
		BroadcastID      int         `json:"broadcast_id"`
		CoinType         string      `json:"coin_type"`
		ComboResourcesID int         `json:"combo_resources_id"`
		ComboSend        struct {
			Action     string      `json:"action"`
			ComboID    string      `json:"combo_id"`
			ComboNum   int         `json:"combo_num"`
			GiftID     int         `json:"gift_id"`
			GiftName   string      `json:"gift_name"`
			GiftNum    int         `json:"gift_num"`
			SendMaster interface{} `json:"send_master"`
			UID        int         `json:"uid"`
			Uname      string      `json:"uname"`
		} `json:"combo_send"`
		ComboStayTime  int     `json:"combo_stay_time"`
		ComboTotalCoin int     `json:"combo_total_coin"`
		CritProb       int     `json:"crit_prob"`
		Demarcation    int     `json:"demarcation"`
		Dmscore        int     `json:"dmscore"`
		Draw           int     `json:"draw"`
		Effect         int     `json:"effect"`
		EffectBlock    int     `json:"effect_block"`
		Face           string  `json:"face"`
		Giftid         int     `json:"giftId"`
		Giftname       string  `json:"giftName"`
		Gifttype       int     `json:"giftType"`
		Gold           int     `json:"gold"`
		GuardLevel     int     `json:"guard_level"`
		IsFirst        bool    `json:"is_first"`
		IsSpecialBatch int     `json:"is_special_batch"`
		Magnification  float64 `json:"magnification"`
		MedalInfo      struct {
			AnchorRoomid     int    `json:"anchor_roomid"`
			AnchorUname      string `json:"anchor_uname"`
			GuardLevel       int    `json:"guard_level"`
			IconID           int    `json:"icon_id"`
			IsLighted        int    `json:"is_lighted"`
			MedalColor       int    `json:"medal_color"`
			MedalColorBorder int    `json:"medal_color_border"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			Special          string `json:"special"`
			TargetID         int    `json:"target_id"`
		} `json:"medal_info"`
		NameColor         string      `json:"name_color"`
		Num               int         `json:"num"`
		OriginalGiftName  string      `json:"original_gift_name"`
		Price             int         `json:"price"`
		Rcost             int64       `json:"rcost"`
		Remain            int         `json:"remain"`
		Rnd               string      `json:"rnd"`
		SendMaster        interface{} `json:"send_master"`
		Silver            int         `json:"silver"`
		Super             int         `json:"super"`
		SuperBatchGiftNum int         `json:"super_batch_gift_num"`
		SuperGiftNum      int         `json:"super_gift_num"`
		SvgaBlock         int         `json:"svga_block"`
		TagImage          string      `json:"tag_image"`
		Tid               string      `json:"tid"`
		Timestamp         int         `json:"timestamp"`
		TopList           interface{} `json:"top_list"`
		TotalCoin         int         `json:"total_coin"`
		UID               int         `json:"uid"`
		Uname             string      `json:"uname"`
	} `json:"data"`
}

/*
{"cmd":"SEND_GIFT","data":{"action":"投喂","batch_combo_id":"batch:gift:combo_id:333010355:585267:20004:1623602183.2621","batch_combo_send":
{"action":"投喂","batch_combo_id":"batch:gift:combo_id:333010355:585267:20004:1623602183.2621","batch_combo_num":1,"blind_gift":null,"gift_i
d":20004,"gift_name":"吃瓜","gift_num":1,"send_master":null,"uid":333010355,"uname":"小尕子_"},"beatId":"","biz_source":"Live","blind_gift":
null,"broadcast_id":0,"coin_type":"gold","combo_resources_id":1,"combo_send":{"action":"投喂","combo_id":"gift:combo_id:333010355:585267:200
04:1623602183.2615","combo_num":1,"gift_id":20004,"gift_name":"吃瓜","gift_num":1,"send_master":null,"uid":333010355,"uname":"小尕子_"},"com
bo_stay_time":3,"combo_total_coin":100,"crit_prob":0,"demarcation":1,"dmscore":16,"draw":0,"effect":0,"effect_block":0,"face":"http://i2.hds
lb.com/bfs/face/dfb44e63d5ac42c28a37c295f6005561be2b089f.jpg","giftId":20004,"giftName":"吃瓜","giftType":1,"gold":0,"guard_level":0,"is_fir
st":true,"is_special_batch":0,"magnification":1,"medal_info":{"anchor_roomid":0,"anchor_uname":"","guard_level":0,"icon_id":0,"is_lighted":0
,"medal_color":6067854,"medal_color_border":12632256,"medal_color_end":12632256,"medal_color_start":12632256,"medal_level":1,"medal_name":"
少废话","special":"","target_id":14110780},"name_color":"","num":1,"original_gift_name":"","price":100,"rcost":3892073142,"remain":0,"rnd":"
1332407012","send_master":null,"silver":0,"super":0,"super_batch_gift_num":1,"super_gift_num":1,"svga_block":0,"tag_image":"","tid":"1623602
183121500002","timestamp":1623602183,"top_list":null,"total_coin":100,"uid":333010355,"uname":"小尕子_"}}
*/
