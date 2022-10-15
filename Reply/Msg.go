package reply

import (
	"encoding/json"
	"io/ioutil"

	c "github.com/qydysky/bili_danmu/CV"
)

/*
	数据为WS_OP_MESSAGE类型的数据分派
*/

var msglog = c.C.Log.Base(`Msg`)

// Msg类型数据处理方法map
var Msg_map = map[string]func(replyF, string){
	`VOICE_JOIN_ROOM_COUNT_INFO`:        replyF.voice_join_room_count_info, //连麦等待
	`VOICE_JOIN_LIST`:                   nil,
	`VOICE_JOIN_STATUS`:                 replyF.voice_join_status,     //连麦人状态
	`STOP_LIVE_ROOM_LIST`:               nil,                          //停止直播的直播间
	`PK_LOTTERY_START`:                  replyF.pk_lottery_start,      //大乱斗pk
	`PK_BATTLE_PRE_NEW`:                 nil,                          //pk准备
	`PK_BATTLE_START_NEW`:               nil,                          //pk开始
	`PK_BATTLE_PROCESS_NEW`:             replyF.pk_battle_process_new, //pk进行中
	`VTR_GIFT_LOTTERY`:                  replyF.vtr_gift_lottery,      //特别礼物
	`ENTRY_EFFECT_MUST_RECEIVE`:         nil,                          //高能榜前三进入
	`GIFT_BAG_DOT`:                      nil,
	`LITTLE_MESSAGE_BOX`:                replyF.little_message_box,           //小消息
	`MESSAGEBOX_USER_MEDAL_CHANGE`:      replyF.messagebox_user_medal_change, //粉丝牌切换
	`HOT_RANK_SETTLEMENT`:               nil,                                 //replyF.hot_rank_settlement, 热门榜获得
	`HOT_RANK_SETTLEMENT_V2`:            replyF.hot_rank_settlement_v2,       //热门榜获得v2
	`HOT_RANK_CHANGED`:                  nil,                                 //replyF.hot_rank_changed, 热门榜变动
	`HOT_RANK_CHANGED_V2`:               nil,                                 //replyF.hot_rank_changed_v2, 热门榜变动v2
	`CARD_MSG`:                          nil,                                 //提示关注
	`WIDGET_BANNER`:                     nil,                                 //每日任务
	`ROOM_ADMINS`:                       nil,                                 //房管列表
	`ONLINE_RANK_TOP3`:                  nil,
	`ONLINE_RANK_COUNT`:                 nil,
	`ONLINE_RANK_V2`:                    nil,
	"TRADING_SCORE":                     nil, //每日任务
	"MATCH_ROOM_CONF":                   nil, //赛事房间配置
	"HOT_ROOM_NOTIFY":                   nil, //热点房间
	"MATCH_TEAM_GIFT_RANK":              nil, //赛事人气比拼
	"ACTIVITY_MATCH_GIFT":               nil, //赛事礼物
	"PK_BATTLE_PRE":                     nil, //人气pk
	"PK_BATTLE_START":                   nil, //人气pk
	"PK_BATTLE_PROCESS":                 nil, //人气pk
	"PK_BATTLE_END":                     nil, //人气pk
	"PK_BATTLE_RANK_CHANGE":             nil, //人气pk
	"PK_BATTLE_SETTLE_USER":             nil, //人气pk
	"PK_BATTLE_SETTLE_V2":               nil, //人气pk
	"PK_BATTLE_SETTLE":                  nil, //人气pk
	"SYS_MSG":                           nil, //系统消息
	"ROOM_SKIN_MSG":                     nil,
	"GUARD_ACHIEVEMENT_ROOM":            nil,
	"ANCHOR_LOT_START":                  replyF.anchor_lot_start, //天选之人开始
	"ANCHOR_LOT_CHECKSTATUS":            nil,
	"ANCHOR_LOT_END":                    nil,                     //天选之人结束
	"ANCHOR_LOT_AWARD":                  replyF.anchor_lot_award, //天选之人获奖
	"COMBO_SEND":                        nil,
	"INTERACT_WORD":                     replyF.interact_word, //进入信息，包含直播间关注提示
	"ACTIVITY_BANNER_UPDATE_V2":         nil,
	"NOTICE_MSG":                        nil,
	"ROOM_BANNER":                       nil,
	"ONLINERANK":                        nil,
	"WELCOME":                           nil,
	"HOUR_RANK_AWARDS":                  nil,
	"ROOM_RANK":                         nil,
	"ROOM_SHIELD":                       nil,
	"USER_TOAST_MSG":                    replyF.user_toast_msg,     //大航海购买信息
	"WIN_ACTIVITY":                      replyF.win_activity,       //活动
	"SPECIAL_GIFT":                      replyF.special_gift,       //节奏风暴
	"GUARD_BUY":                         nil,                       //replyF.guard_buy,//大航海购买
	"WELCOME_GUARD":                     nil,                       //replyF.welcome_guard,//大航海进入 ？已废弃？
	"DANMU_MSG":                         replyF.danmu,              //弹幕
	"DANMU_MSG:4:0:2:2:2:0":             replyF.danmu,              //弹幕
	"ROOM_CHANGE":                       replyF.room_change,        //房间信息分区改变
	"ROOM_SILENT_OFF":                   replyF.roomsilent,         //禁言结束
	"ROOM_SILENT_ON":                    replyF.roomsilent,         //禁言开始
	"SEND_GIFT":                         replyF.send_gift,          //礼物
	"ROOM_BLOCK_MSG":                    replyF.room_block_msg,     //封禁
	"PREPARING":                         replyF.preparing,          //下播
	"LIVE":                              replyF.live,               //开播
	"SUPER_CHAT_ENTRANCE":               nil,                       //SC入口
	"SUPER_CHAT_MESSAGE_DELETE":         nil,                       //SC删除
	"SUPER_CHAT_MESSAGE":                nil,                       //replyF.super_chat_message,//SC
	"SUPER_CHAT_MESSAGE_JPN":            replyF.super_chat_message, //SC
	"PANEL":                             nil,                       //replyF.panel,//排行榜 被HOT_RANK_CHANGED替代
	"ENTRY_EFFECT":                      replyF.entry_effect,       //进入特效
	"ROOM_REAL_TIME_MESSAGE_UPDATE":     nil,                       //replyF.roominfo,//粉丝数
	"WATCHED_CHANGE":                    replyF.watched_change,     //Msg-观看人数
	"FULL_SCREEN_SPECIAL_EFFECT":        nil,
	"GIFT_BOARD_RED_DOT":                nil,
	"USER_PANEL_RED_ALARM":              nil,
	"POPULARITY_RED_POCKET_NEW":         replyF.popularity_red_pocket_new,   //老板打赏新礼物红包
	"POPULARITY_RED_POCKET_START":       replyF.popularity_red_pocket_start, //老板打赏礼物红包开始
	"POPULARITY_RED_POCKET_WINNER_LIST": nil,                                //老板打赏礼物红包的得奖名单
	"COMMON_NOTICE_DANMAKU":             replyF.common_notice_danmaku,       //元气赏连抽
	"ACTIVITY_BANNER_CHANGE":            nil,                                //活动标题改变
	"ACTIVITY_BANNER_CHANGE_V2":         replyF.activity_banner_change_v2,   //活动标题改变v2
	"VIDEO_CONNECTION_JOIN_START":       replyF.video_connection_join_start, //开始了与某人的视频连线
	"VIDEO_CONNECTION_JOIN_END":         replyF.video_connection_join_end,   //结束了与某人的视频连线
	"VIDEO_CONNECTION_MSG":              replyF.video_connection_msg,        //视频连线状态改变
	"WARNING":                           replyF.warning,                     //超管警告
	"DANMU_AGGREGATION":                 nil,                                //聚合弹幕
	"GUARD_HONOR_THOUSAND":              nil,
	"LIKE_INFO_V3_CLICK":                replyF.like_info_v3_click, //为主播点赞了
	"LIKE_INFO_V3_UPDATE":               nil,                       //为主播点赞了总个数
	"USER_TASK_PROGRESS":                nil,
	"LITTLE_TIPS":                       replyF.little_tips, //小提示窗口
	"LIKE_INFO_V3_NOTICE":               nil,
}

// 屏蔽不需要的消息
func init() {
	{ //加载不需要的消息
		bb, err := ioutil.ReadFile("config/config_disable_msg.json")
		if err != nil {
			return
		}
		var buf map[string]interface{}
		json.Unmarshal(bb, &buf)
		for k, v := range buf {
			if able, ok := v.(bool); ok { //设置为true时，使用默认显示
				if able {
					Msg_map[k] = replyF.defaultMsg
				} else {
					Msg_map[k] = nil
				}
			}
		}
	}
}

// Msg类型数据处理方法挑选
// 识别cmd字段类型，查找上述map中设置的方法，并将json转为字符串型传入
func Msg(b []byte) {

	msglog := msglog.Base_add(`select func`)
	var tmp struct {
		Cmd string `json:"cmd"`
	}
	if e := json.Unmarshal(b, &tmp); e != nil {
		msglog.L(`E: `, e)
		return
	}
	if F, ok := Msg_map[tmp.Cmd]; ok {
		if F != nil {
			F(replyF{}, string(b))
		}
	} else {
		(replyF{}).defaultMsg(string(b))
	}
}
