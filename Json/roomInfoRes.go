package part

type Roominfores struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		RoomInfo struct {
			UID            int    `json:"uid"`
			RoomID         int    `json:"room_id"`
			ShortID        int    `json:"short_id"`
			Title          string `json:"title"`
			Cover          string `json:"cover"`
			Tags           string `json:"tags"`
			Background     string `json:"background"`
			Description    string `json:"description"`
			LiveStatus     int    `json:"live_status"`
			LiveStartTime  int    `json:"live_start_time"`
			LiveScreenType int    `json:"live_screen_type"`
			LockStatus     int    `json:"lock_status"`
			LockTime       int    `json:"lock_time"`
			HiddenStatus   int    `json:"hidden_status"`
			HiddenTime     int    `json:"hidden_time"`
			AreaID         int    `json:"area_id"`
			AreaName       string `json:"area_name"`
			ParentAreaID   int    `json:"parent_area_id"`
			ParentAreaName string `json:"parent_area_name"`
			Keyframe       string `json:"keyframe"`
			SpecialType    int    `json:"special_type"`
			UpSession      string `json:"up_session"`
			PkStatus       int    `json:"pk_status"`
			IsStudio       bool   `json:"is_studio"`
			Pendants       struct {
				Frame struct {
					Name  string `json:"name"`
					Value string `json:"value"`
					Desc  string `json:"desc"`
				} `json:"frame"`
			} `json:"pendants"`
			OnVoiceJoin int `json:"on_voice_join"`
			Online      int `json:"online"`
			RoomType    struct {
				Three13 int `json:"3-13"`
				Four1   int `json:"4-1"`
			} `json:"room_type"`
		} `json:"room_info"`
		AnchorInfo struct {
			BaseInfo struct {
				Uname        string `json:"uname"`
				Face         string `json:"face"`
				Gender       string `json:"gender"`
				OfficialInfo struct {
					Role  int    `json:"role"`
					Title string `json:"title"`
					Desc  string `json:"desc"`
				} `json:"official_info"`
			} `json:"base_info"`
			LiveInfo struct {
				Level        int    `json:"level"`
				LevelColor   int    `json:"level_color"`
				Score        int    `json:"score"`
				UpgradeScore int    `json:"upgrade_score"`
				Current      []int  `json:"current"`
				Next         []int  `json:"next"`
				Rank         string `json:"rank"`
			} `json:"live_info"`
			RelationInfo struct {
				Attention int `json:"attention"`
			} `json:"relation_info"`
			MedalInfo struct {
				MedalName string `json:"medal_name"`
				MedalID   int    `json:"medal_id"`
				Fansclub  int    `json:"fansclub"`
			} `json:"medal_info"`
		} `json:"anchor_info"`
		NewsInfo struct {
			UID     int    `json:"uid"`
			Ctime   string `json:"ctime"`
			Content string `json:"content"`
		} `json:"news_info"`
		RankdbInfo struct {
			Roomid    int    `json:"roomid"`
			RankDesc  string `json:"rank_desc"`
			Color     string `json:"color"`
			H5URL     string `json:"h5_url"`
			WebURL    string `json:"web_url"`
			Timestamp int    `json:"timestamp"`
		} `json:"rankdb_info"`
		AreaRankInfo struct {
			Arearank struct {
				Index int    `json:"index"`
				Rank  string `json:"rank"`
			} `json:"areaRank"`
			Liverank struct {
				Rank string `json:"rank"`
			} `json:"liveRank"`
		} `json:"area_rank_info"`
		BattleRankEntryInfo struct {
			FirstRankImgURL string `json:"first_rank_img_url"`
			RankName        string `json:"rank_name"`
			ShowStatus      int    `json:"show_status"`
		} `json:"battle_rank_entry_info"`
		TabInfo struct {
			List []struct {
				Type      string `json:"type"`
				Desc      string `json:"desc"`
				Isfirst   int    `json:"isFirst"`
				Isevent   int    `json:"isEvent"`
				Eventtype string `json:"eventType"`
				Listtype  string `json:"listType"`
				Apiprefix string `json:"apiPrefix"`
				RankName  string `json:"rank_name"`
			} `json:"list"`
		} `json:"tab_info"`
		ActivityInitInfo struct {
			Eventlist []interface{} `json:"eventList"`
			Weekinfo  struct {
				Bannerinfo interface{} `json:"bannerInfo"`
				Giftname   interface{} `json:"giftName"`
			} `json:"weekInfo"`
			Giftname interface{} `json:"giftName"`
			Lego     struct {
				Timestamp int    `json:"timestamp"`
				Config    string `json:"config"`
			} `json:"lego"`
		} `json:"activity_init_info"`
		VoiceJoinInfo struct {
			Status struct {
				Open        int    `json:"open"`
				AnchorOpen  int    `json:"anchor_open"`
				Status      int    `json:"status"`
				UID         int    `json:"uid"`
				UserName    string `json:"user_name"`
				HeadPic     string `json:"head_pic"`
				Guard       int    `json:"guard"`
				StartAt     int    `json:"start_at"`
				CurrentTime int    `json:"current_time"`
			} `json:"status"`
			Icons struct {
				IconClose    string `json:"icon_close"`
				IconOpen     string `json:"icon_open"`
				IconWait     string `json:"icon_wait"`
				IconStarting string `json:"icon_starting"`
			} `json:"icons"`
			WebShareLink string `json:"web_share_link"`
		} `json:"voice_join_info"`
		AdBannerInfo struct {
			Data []struct {
				ID       int    `json:"id"`
				Title    string `json:"title"`
				Location string `json:"location"`
				Position int    `json:"position"`
				Pic      string `json:"pic"`
				Link     string `json:"link"`
				Weight   int    `json:"weight"`
			} `json:"data"`
		} `json:"ad_banner_info"`
		SkinInfo struct {
			ID          int    `json:"id"`
			SkinName    string `json:"skin_name"`
			SkinConfig  string `json:"skin_config"`
			ShowText    string `json:"show_text"`
			SkinURL     string `json:"skin_url"`
			StartTime   int    `json:"start_time"`
			EndTime     int    `json:"end_time"`
			CurrentTime int    `json:"current_time"`
		} `json:"skin_info"`
		WebBannerInfo struct {
			ID               int    `json:"id"`
			Title            string `json:"title"`
			Left             string `json:"left"`
			Right            string `json:"right"`
			JumpURL          string `json:"jump_url"`
			BgColor          string `json:"bg_color"`
			HoverColor       string `json:"hover_color"`
			TextBgColor      string `json:"text_bg_color"`
			TextHoverColor   string `json:"text_hover_color"`
			LinkText         string `json:"link_text"`
			LinkColor        string `json:"link_color"`
			InputColor       string `json:"input_color"`
			InputTextColor   string `json:"input_text_color"`
			InputHoverColor  string `json:"input_hover_color"`
			InputBorderColor string `json:"input_border_color"`
			InputSearchColor string `json:"input_search_color"`
		} `json:"web_banner_info"`
		LolInfo struct {
			LolActivity struct {
				Status     int    `json:"status"`
				GuessCover string `json:"guess_cover"`
				VoteCover  string `json:"vote_cover"`
				VoteH5URL  string `json:"vote_h5_url"`
				VoteUseH5  bool   `json:"vote_use_h5"`
			} `json:"lol_activity"`
		} `json:"lol_info"`
		WishListInfo struct {
			List   []interface{} `json:"list"`
			Status int           `json:"status"`
		} `json:"wish_list_info"`
		ScoreCardInfo  interface{} `json:"score_card_info"`
		PkInfo         interface{} `json:"pk_info"`
		BattleInfo     interface{} `json:"battle_info"`
		SilentRoomInfo struct {
			Type       string `json:"type"`
			Level      int    `json:"level"`
			Second     int    `json:"second"`
			ExpireTime int    `json:"expire_time"`
		} `json:"silent_room_info"`
		SwitchInfo struct {
			CloseGuard   bool `json:"close_guard"`
			CloseGift    bool `json:"close_gift"`
			CloseOnline  bool `json:"close_online"`
			CloseDanmaku bool `json:"close_danmaku"`
		} `json:"switch_info"`
		RecordSwitchInfo struct {
			RecordTab bool `json:"record_tab"`
		} `json:"record_switch_info"`
		RoomConfigInfo struct {
			DmText string `json:"dm_text"`
		} `json:"room_config_info"`
		GiftMemoryInfo struct {
			List interface{} `json:"list"`
		} `json:"gift_memory_info"`
		NewSwitchInfo struct {
			RoomSocket           int `json:"room-socket"`
			RoomPropSend         int `json:"room-prop-send"`
			RoomSailing          int `json:"room-sailing"`
			RoomInfoPopularity   int `json:"room-info-popularity"`
			RoomDanmakuEditor    int `json:"room-danmaku-editor"`
			RoomEffect           int `json:"room-effect"`
			RoomFansMedal        int `json:"room-fans_medal"`
			RoomReport           int `json:"room-report"`
			RoomFeedback         int `json:"room-feedback"`
			RoomPlayerWatermark  int `json:"room-player-watermark"`
			RoomRecommendLiveOff int `json:"room-recommend-live_off"`
			RoomActivity         int `json:"room-activity"`
			RoomWebBanner        int `json:"room-web_banner"`
			RoomSilverSeedsBox   int `json:"room-silver_seeds-box"`
			RoomWishingBottle    int `json:"room-wishing_bottle"`
			RoomBoard            int `json:"room-board"`
			RoomSupplication     int `json:"room-supplication"`
			RoomHourRank         int `json:"room-hour_rank"`
			RoomWeekRank         int `json:"room-week_rank"`
			RoomAnchorRank       int `json:"room-anchor_rank"`
			RoomInfoIntegral     int `json:"room-info-integral"`
			RoomSuperChat        int `json:"room-super-chat"`
			RoomTab              int `json:"room-tab"`
			RoomHotRank          int `json:"room-hot-rank"`
		} `json:"new_switch_info"`
		SuperChatInfo struct {
			Status      int    `json:"status"`
			JumpURL     string `json:"jump_url"`
			Icon        string `json:"icon"`
			RankedMark  int    `json:"ranked_mark"`
			MessageList []struct {
				ID                    int    `json:"id"`
				UID                   int    `json:"uid"`
				Price                 int    `json:"price"`
				Rate                  int    `json:"rate"`
				BackgroundImage       string `json:"background_image"`
				BackgroundColor       string `json:"background_color"`
				BackgroundIcon        string `json:"background_icon"`
				BackgroundPriceColor  string `json:"background_price_color"`
				BackgroundBottomColor string `json:"background_bottom_color"`
				FontColor             string `json:"font_color"`
				Time                  int    `json:"time"`
				StartTime             int    `json:"start_time"`
				EndTime               int    `json:"end_time"`
				Message               string `json:"message"`
				TransMark             int    `json:"trans_mark"`
				MessageTrans          string `json:"message_trans"`
				Token                 string `json:"token"`
				Ts                    int    `json:"ts"`
				UserInfo              struct {
					Face       string `json:"face"`
					FaceFrame  string `json:"face_frame"`
					Uname      string `json:"uname"`
					UserLevel  int    `json:"user_level"`
					GuardLevel int    `json:"guard_level"`
					IsVip      int    `json:"is_vip"`
					IsSvip     int    `json:"is_svip"`
					IsMainVip  int    `json:"is_main_vip"`
				} `json:"user_info"`
			} `json:"message_list"`
		} `json:"super_chat_info"`
		OnlineGoldRankInfoV2 struct {
			List []struct {
				UID        int    `json:"uid"`
				Face       string `json:"face"`
				Uname      string `json:"uname"`
				Score      string `json:"score"`
				Rank       int    `json:"rank"`
				GuardLevel int    `json:"guard_level"`
			} `json:"list"`
		} `json:"online_gold_rank_info_v2"`
		VideoConnectionInfo interface{} `json:"video_connection_info"`
		PlayerThrottleInfo  struct {
			Status              int `json:"status"`
			NormalSleepTime     int `json:"normal_sleep_time"`
			FullscreenSleepTime int `json:"fullscreen_sleep_time"`
			TabSleepTime        int `json:"tab_sleep_time"`
			PromptTime          int `json:"prompt_time"`
		} `json:"player_throttle_info"`
		GuardInfo struct {
			Count                   int `json:"count"`
			AnchorGuardAchieveLevel int `json:"anchor_guard_achieve_level"`
		} `json:"guard_info"`
		HotRankInfo struct {
			Rank      int    `json:"rank"`
			Trend     int    `json:"trend"`
			Countdown int    `json:"countdown"`
			Timestamp int    `json:"timestamp"`
			URL       string `json:"url"`
			Icon      string `json:"icon"`
			AreaName  string `json:"area_name"`
		} `json:"hot_rank_info"`
	} `json:"data"`
}