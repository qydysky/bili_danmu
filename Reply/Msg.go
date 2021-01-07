package reply

import (
	p "github.com/qydysky/part"
	c "github.com/qydysky/bili_danmu/CV"
	s "github.com/qydysky/part/buf"
)

/*
	数据为WS_OP_MESSAGE类型的数据分派
*/

var msglog = c.Log.Base(`Msg`)

//Msg类型数据处理方法map
var Msg_map = map[string]func(replyF, string) {
	`HOT_RANK_SETTLEMENT`:nil,
	`HOT_RANK_CHANGED`:replyF.hot_rank_changed,//热门榜变动
	`CARD_MSG`:nil,//提示关注
	`LIVE_INTERACTIVE_GAME`:nil,
	`WIDGET_BANNER`:nil,//每日任务
	`ROOM_ADMINS`:nil,//房管列表
	`room_admin_entrance`:nil,
	`ONLINE_RANK_TOP3`:nil,
	`ONLINE_RANK_COUNT`:nil,
	`ONLINE_RANK_V2`:nil,
	"TRADING_SCORE":nil,//每日任务
	"MATCH_ROOM_CONF":nil,//赛事房间配置
	"HOT_ROOM_NOTIFY":nil,//热点房间
	"MATCH_TEAM_GIFT_RANK":nil,//赛事人气比拼
	"ACTIVITY_MATCH_GIFT":nil,//赛事礼物
	"PK_BATTLE_PRE":nil,//人气pk
	"PK_BATTLE_START":nil,//人气pk
	"PK_BATTLE_PROCESS":nil,//人气pk
	"PK_BATTLE_END":nil,//人气pk
	"PK_BATTLE_RANK_CHANGE":nil,//人气pk
	"PK_BATTLE_SETTLE_USER":nil,//人气pk
	"PK_BATTLE_SETTLE_V2":nil,//人气pk
	"PK_BATTLE_SETTLE":nil,//人气pk
	"SYS_MSG":nil,//系统消息
	"ROOM_SKIN_MSG":nil,
	"GUARD_ACHIEVEMENT_ROOM":nil,
	"ANCHOR_LOT_START":replyF.anchor_lot_start,//天选之人开始
	"ANCHOR_LOT_CHECKSTATUS":nil,
	"ANCHOR_LOT_END":nil,//天选之人结束
	"ANCHOR_LOT_AWARD":replyF.anchor_lot_award,//天选之人获奖
	"COMBO_SEND":nil,
	"INTERACT_WORD":replyF.interact_word,//进入信息，包含直播间关注提示
	"ACTIVITY_BANNER_UPDATE_V2":nil,
	"NOTICE_MSG":nil,
	"ROOM_BANNER":nil,
	"ONLINERANK":nil,
	"WELCOME":nil,
	"HOUR_RANK_AWARDS":nil,
	"ROOM_RANK":nil,
	"ROOM_SHIELD":nil,
	"USER_TOAST_MSG":replyF.user_toast_msg,//大航海购买信息
	"WIN_ACTIVITY":replyF.win_activity,//活动
	"SPECIAL_GIFT":replyF.special_gift,//节奏风暴
	"GUARD_BUY":nil,//replyF.guard_buy,//大航海购买
	"WELCOME_GUARD":replyF.welcome_guard,//大航海进入
	"DANMU_MSG":replyF.danmu,//弹幕
	"ROOM_CHANGE":replyF.room_change,//房间信息分区改变
	"ROOM_SILENT_OFF":replyF.roomsilent,//禁言结束
	"ROOM_SILENT_ON":replyF.roomsilent,//禁言开始
	"SEND_GIFT":replyF.send_gift,//礼物
	"ROOM_BLOCK_MSG":replyF.room_block_msg,//封禁
	"PREPARING":replyF.preparing,//下播
	"LIVE":replyF.live,//开播
	"SUPER_CHAT_ENTRANCE":nil,//SC入口
	"SUPER_CHAT_MESSAGE_DELETE":nil,//SC删除
	"SUPER_CHAT_MESSAGE":nil,//replyF.super_chat_message,//SC
	"SUPER_CHAT_MESSAGE_JPN":replyF.super_chat_message,//SC
	"PANEL":nil,//replyF.panel,//排行榜 被HOT_RANK_CHANGED替代
	"ENTRY_EFFECT":nil,//replyF.entry_effect,//进入特效
	"ROOM_REAL_TIME_MESSAGE_UPDATE":nil,//replyF.roominfo,//粉丝数
}

//屏蔽不需要的消息
func init(){
	{//加载不需要的消息
		buf := s.New()
		buf.Load("config/config_disable_msg.json")
		for k,v := range buf.B {
			if able,ok := v.(bool);ok {//设置为true时，使用默认显示
				if able {
					Msg_map[k] = replyF.defaultMsg
				} else {
					Msg_map[k] = nil
				}
			}
		}
	}
}

//Msg类型数据处理方法挑选
//识别cmd字段类型，查找上述map中设置的方法，并将json转为字符串型传入
func Msg(b []byte) {
	s := string(b)
	if cmd := p.Json().GetValFromS(s, "cmd");cmd == nil {
		msglog.L(`E: `,"cmd", s)
		return
	} else {
		var f replyF

		if F, ok := Msg_map[cmd.(string)]; ok {
			if F != nil {F(f, s)}
		} else {
			f.defaultMsg(s)
		}
	}

	return 
}
