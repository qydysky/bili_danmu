package send

func Send_gift(){
	{//参数来源1
		//https://api.live.bilibili.com/xlive/web-room/v1/gift/bag_list?t=1614667412767&room_id=22347054

		/*
		{
			"code": 0,
			"message": "0",
			"ttl": 1,
			"data": {
				"list": [
					{
						"bag_id": 235312419,
						"gift_id": 1,
						"gift_name": "辣条",
						"gift_num": 1,
						"gift_type": 5,
						"expire_at": 1614700800,
						"corner_mark": "1天",
						"corner_color": "",
						"count_map": [
							{
								"num": 1,
								"text": ""
							}
						],
						"bind_roomid": 0,
						"bind_room_text": "",
						"type": 1,
						"card_image": "",
						"card_gif": "",
						"card_id": 0,
						"card_record_id": 0,
						"is_show_send": false
					},
					{
						"bag_id": 235157184,
						"gift_id": 30607,
						"gift_name": "小心心",
						"gift_num": 16,
						"gift_type": 5,
						"expire_at": 1615046400,
						"corner_mark": "5天",
						"corner_color": "",
						"count_map": [
							{
								"num": 1,
								"text": ""
							},
							{
								"num": 16,
								"text": "全部"
							}
						],
						"bind_roomid": 0,
						"bind_room_text": "",
						"type": 1,
						"card_image": "",
						"card_gif": "",
						"card_id": 0,
						"card_record_id": 0,
						"is_show_send": false
					},
					{
						"bag_id": 235162609,
						"gift_id": 30607,
						"gift_name": "小心心",
						"gift_num": 24,
						"gift_type": 5,
						"expire_at": 1615132800,
						"corner_mark": "6天",
						"corner_color": "",
						"count_map": [
							{
								"num": 1,
								"text": ""
							},
							{
								"num": 24,
								"text": "全部"
							}
						],
						"bind_roomid": 0,
						"bind_room_text": "",
						"type": 1,
						"card_image": "",
						"card_gif": "",
						"card_id": 0,
						"card_record_id": 0,
						"is_show_send": false
					},
					{
						"bag_id": 235077495,
						"gift_id": 1,
						"gift_name": "辣条",
						"gift_num": 20,
						"gift_type": 5,
						"expire_at": 1615132800,
						"corner_mark": "6天",
						"corner_color": "",
						"count_map": [
							{
								"num": 1,
								"text": ""
							},
							{
								"num": 20,
								"text": "全部"
							}
						],
						"bind_roomid": 0,
						"bind_room_text": "",
						"type": 1,
						"card_image": "",
						"card_gif": "",
						"card_id": 0,
						"card_record_id": 0,
						"is_show_send": false
					},
					{
						"bag_id": 235282499,
						"gift_id": 30607,
						"gift_name": "小心心",
						"gift_num": 3,
						"gift_type": 5,
						"expire_at": 1615219200,
						"corner_mark": "7天",
						"corner_color": "",
						"count_map": [
							{
								"num": 1,
								"text": ""
							},
							{
								"num": 3,
								"text": "全部"
							}
						],
						"bind_roomid": 0,
						"bind_room_text": "",
						"type": 1,
						"card_image": "",
						"card_gif": "",
						"card_id": 0,
						"card_record_id": 0,
						"is_show_send": false
					}
				],
				"time": 0
			}
		}
		*/
	}
	{//发送请求（银瓜子礼物）
		//https://api.live.bilibili.com/gift/v2/live/bag_send

		// req
		// {
		// 	"uid": "29183321",//客户uid
		// 	"gift_id": "1",//礼物id from 1
		// 	"ruid": "623441612",//主播uid
		// 	"send_ruid": "0",//固定值
		// 	"gift_num": "1",//发送数量 from 1
		// 	"bag_id": "235312419",//礼物包 from 1
		// 	"platform": "pc",//平台固定值
		// 	"biz_code": "live",//固定值
		// 	"biz_id": "22347054",//房间id
		// 	"rnd": "1614667054",//时间戳
		// 	"storm_beat_id": "0",//固定值
		// 	"metadata": "",//固定值
		// 	"price": "0",//价值
		// 	"csrf_token": "××",//from cookie
		// 	"csrf": "××",//from cookie
		// 	"visit_id": ""//固定值
		// }

		// res
		// {
		// 	"code": 0,
		// 	"msg": "success",
		// 	"message": "success",
		// 	"data": {
		// 		"tid": "1614667076110100001",
		// 		"uid": 29183321,
		// 		"uname": "qydysky丶",
		// 		"face": "https://i2.hdslb.com/bfs/face/8c6d1ce3f96dc86d0fc7876a2824910c92ae6802.jpg",
		// 		"guard_level": 0,
		// 		"ruid": 623441612,
		// 		"rcost": 18194089,
		// 		"gift_id": 1,
		// 		"gift_type": 5,
		// 		"gift_name": "辣条",
		// 		"gift_num": 1,
		// 		"gift_action": "投喂",
		// 		"gift_price": 100,
		// 		"coin_type": "silver",
		// 		"total_coin": 100,
		// 		"pay_coin": 100,
		// 		"metadata": "",
		// 		"fulltext": "",
		// 		"rnd": "1614667054",
		// 		"tag_image": "",
		// 		"effect_block": 1,
		// 		"extra": {
		// 			"wallet": null,
		// 			"gift_bag": {
		// 				"bag_id": 235312419,
		// 				"gift_num": 1
		// 			},
		// 			"top_list": [],
		// 			"follow": null,
		// 			"medal": null,
		// 			"title": null,
		// 			"pk": {
		// 				"pk_gift_tips": "",
		// 				"crit_prob": -1
		// 			},
		// 			"fulltext": "",
		// 			"event": {
		// 				"event_score": 0,
		// 				"event_redbag_num": 0
		// 			},
		// 			"capsule": null,
		// 			"lottery_id": ""
		// 		},
		// 		"blow_switch": 0,
		// 		"send_tips": "赠送成功",
		// 		"discount_id": 0,
		// 		"gift_effect": {
		// 			"super": 0,
		// 			"combo_timeout": 0,
		// 			"super_gift_num": 0,
		// 			"super_batch_gift_num": 0,
		// 			"batch_combo_id": "",
		// 			"broadcast_msg_list": [],
		// 			"small_tv_list": [],
		// 			"beat_storm": null,
		// 			"combo_id": "",
		// 			"smallTVCountFlag": true
		// 		},
		// 		"send_master": null,
		// 		"crit_prob": -1,
		// 		"combo_stay_time": 3,
		// 		"combo_total_coin": 0,
		// 		"demarcation": 1,
		// 		"magnification": 1,
		// 		"combo_resources_id": 1,
		// 		"is_special_batch": 0,
		// 		"send_gift_countdown": 6,
		// 		"bp_cent_balance": 0,
		// 		"price": 0,
		// 		"left_num": 0,
		// 		"need_num": 0,
		// 		"available_num": 0
		// 	}
		// }
	}
}
