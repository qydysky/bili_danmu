package reply

import (
	p "github.com/qydysky/part"
)
/*
	数据为WS_OP_MESSAGE类型的
*/

var msglog = p.Logf().New().Base(-1, "Msg.go").Open("danmu.log").Level(1)
var Msg_map = map[string]func(replyF, string) {
	"ROOM_SKIN_MSG":nil,
	"GUARD_ACHIEVEMENT_ROOM":nil,
	"ANCHOR_LOT_START":nil,//天选之人开始
	"ANCHOR_LOT_CHECKSTATUS":nil,
	"ANCHOR_LOT_END":nil,//天选之人结束
	"ANCHOR_LOT_AWARD":nil,//天选之人获奖
	"COMBO_SEND":nil,
	"INTERACT_WORD":nil,
	"ACTIVITY_BANNER_UPDATE_V2":nil,
	"NOTICE_MSG":nil,
	"ROOM_BANNER":nil,
	"ONLINERANK":nil,
	"WELCOME":nil,
	"HOUR_RANK_AWARDS":nil,
	"ROOM_RANK":nil,
	"ROOM_SHIELD":nil,
	"USER_TOAST_MSG":nil,
	"WIN_ACTIVITY":replyF.win_activity,//活动
	"SPECIAL_GIFT":replyF.special_gift,//节奏风暴
	"GUARD_BUY":replyF.guard_buy,//大航海购买
	"WELCOME_GUARD":replyF.welcome_guard,//大航海进入
	"DANMU_MSG":replyF.danmu,//弹幕
	"ROOM_CHANGE":replyF.room_change,//房间信息分区改变
	"ROOM_SILENT_OFF":replyF.roomsilent,//禁言结束
	"ROOM_SILENT_ON":replyF.roomsilent,//禁言开始
	"SEND_GIFT":replyF.send_gift,//礼物
	"ROOM_BLOCK_MSG":replyF.room_block_msg,//封禁
	"PREPARING":replyF.preparing,//下播
	"LIVE":replyF.live,//开播
	"SUPER_CHAT_MESSAGE":nil,//replyF.super_chat_message,//SC
	"SUPER_CHAT_MESSAGE_JPN":replyF.super_chat_message,//SC
	"PANEL":replyF.panel,//排行榜
	"ENTRY_EFFECT":nil,//replyF.entry_effect,//进入特效
	"ROOM_REAL_TIME_MESSAGE_UPDATE":nil,//replyF.roominfo,//粉丝数
}

func Msg(b []byte) {
	s := string(b)
	if cmd := p.Json().GetValFromS(s, "cmd");cmd == nil {
		msglog.E("cmd", s)
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
