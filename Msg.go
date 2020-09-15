package bili_danmu

import (
	"bytes"
	"compress/zlib"

	p "github.com/qydysky/part"
)
/*
	数据为WS_OP_MESSAGE类型的
*/

var msglog = p.Logf().New().Base(-1, "Msg.go>").Open("danmu.log").Level(1)
var Msg_cookie string
var Msg_roomid int

func Msg(b []byte, compress bool) {
	if compress {
		readc, err := zlib.NewReader(bytes.NewReader(b[16:]))
		if err != nil {msglog.E("解压错误");return}
		defer readc.Close()
		
		buf := bytes.NewBuffer(nil)
		if _, err := buf.ReadFrom(readc);err != nil {msglog.E("解压错误");return}
		b = buf.Bytes()
	}

	for len(b) != 0 {
		
		var packL int32
		if ist, packl := headChe(b[:16], len(b), WS_BODY_PROTOCOL_VERSION_NORMAL, WS_OP_MESSAGE, 0, 0); !ist {
			msglog.E("头错误");return
		} else {
			packL = packl
		}

		s := string(b[16:packL])
		b = b[packL:]
		if cmd := p.Json().GetValFromS(s, "cmd");cmd == nil {
			msglog.E("cmd", s)
			return
		} else {
			switch cmd.(string) {
			case "ANCHOR_LOT_START"://天选之人开始
			case "ANCHOR_LOT_CHECKSTATUS":
			case "ANCHOR_LOT_END"://天选之人结束
			case "ANCHOR_LOT_AWARD"://天选之人获奖
			case "COMBO_SEND":
			case "INTERACT_WORD":
			case "ACTIVITY_BANNER_UPDATE_V2":
			case "NOTICE_MSG":
			case "ROOM_BANNER":
			case "ONLINERANK":
			case "WELCOME":
			case "HOUR_RANK_AWARDS":
			case "ROOM_RANK":
			case "ROOM_SHIELD":
			case "USER_TOAST_MSG":
			case "GUARD_BUY"://大航海购买
			case "WELCOME_GUARD"://welcome_guard(s)//大航海进入
			case "ROOM_SILENT_OFF", "ROOM_SILENT_ON":roomsilent(s);//禁言
			case "SEND_GIFT":send_gift(s)//礼物
			case "ROOM_BLOCK_MSG":room_block_msg(s)//封禁
			case "PREPARING":preparing(s)//下播
			case "LIVE":live(s)//开播
			case "SUPER_CHAT_MESSAGE", "SUPER_CHAT_MESSAGE_JPN":super_chat_message(s)//打赏
			case "PANEL":panel(s)//排行榜
			case "ENTRY_EFFECT":entry_effect(s)//进入特效
			case "ROOM_REAL_TIME_MESSAGE_UPDATE":roominfo(s)//粉丝数
			case "DANMU_MSG":danmu(s)//弹幕
			default:msglog.I("Unknow cmd", s)
			}
		}
	}

	return 
}

func welcome_guard(s string){
	msglog.Base(1, "房")

	username := p.Json().GetValFromS(s, "data.username");
	guard_level := p.Json().GetValFromS(s, "data.guard_level");

	var sh = []interface{}{"欢迎"}

	if username != nil {
		sh = append(sh, username.(string), "进入直播间")
	}
	if guard_level != nil {
		sh = append(sh, "等级", int64(guard_level.(float64)))
	}
	if len(sh) == 0 {return}

	msglog.I(sh...)
}

func send_gift(s string){
	msglog.Base(1, "礼")

	coin_type := p.Json().GetValFromS(s, "data.coin_type");
	num := p.Json().GetValFromS(s, "data.num");
	uname := p.Json().GetValFromS(s, "data.uname");
	action := p.Json().GetValFromS(s, "data.action");
	giftName := p.Json().GetValFromS(s, "data.giftName");
	price := p.Json().GetValFromS(s, "data.price");

	var sh []interface{}
	var allprice int64

	if num != nil {
		sh = append(sh, int64(num.(float64)), "x")
	}
	if price != nil {
		allprice = int64(num.(float64) * price.(float64))
		sh = append(sh, "(", allprice, "x 金瓜子 )")
	}
	if uname != nil {
		sh = append(sh, uname.(string))
	}
	if action != nil {
		sh = append(sh, action.(string))
	}
	if giftName != nil {
		sh = append(sh, giftName.(string))
	}
	
	if len(sh) == 0 {return}

	//小于1万金瓜子 银瓜子不显示
	if allprice < 10000 || coin_type.(string) == "silver" {msglog.T(sh...);return}
	msglog.I(sh...)
}

func room_block_msg(s string) {
	msglog.Base(1, "封")

	if uname := p.Json().GetValFromS(s, "uname");uname == nil {
		msglog.E("uname", uname)
		return
	} else {
		msglog.I("用户", uname, "已被封禁")
	}
}

func preparing(s string) {
	msglog.Base(1, "房")

	if roomid := p.Json().GetValFromS(s, "roomid");roomid == nil {
		msglog.E("roomid", roomid)
		return
	} else {
		msglog.I("房间", roomid.(string), "下播了")
	}
}

func live(s string) {
	msglog.Base(1, "房")

	if roomid := p.Json().GetValFromS(s, "roomid");roomid == nil {
		msglog.E("roomid", roomid)
		return
	} else {
		msglog.I("房间", roomid.(string), "开播了")
	}
}

func super_chat_message(s string){
	msglog.Base(1, "礼")

	uname := p.Json().GetValFromS(s, "data.user_info.uname");
	price := p.Json().GetValFromS(s, "data.price");
	message := p.Json().GetValFromS(s, "data.message");
	message_jpn := p.Json().GetValFromS(s, "data.message_jpn");

	var sh = []interface{}{"打赏: "}

	if uname != nil {
		sh = append(sh, uname.(string))
	}
	if price != nil {
		sh = append(sh, "￥", int64(price.(float64)))
	}
	if message != nil {
		sh = append(sh, message.(string))
	}
	if message_jpn != nil {
		sh = append(sh, message_jpn.(string))
	}

	if len(sh) != 0 {msglog.I(sh...)}
}

func panel(s string){
	msglog.Base(1, "房")

	if note := p.Json().GetValFromS(s, "data.note");note == nil {
		msglog.E("note", note)
		return
	} else {
		msglog.I("排行", note.(string))
	}
}

func entry_effect(s string){
	msglog.Base(1, "房")

	if copy_writing := p.Json().GetValFromS(s, "data.copy_writing");copy_writing == nil {
		msglog.E("copy_writing", copy_writing)
		return
	} else {
		msglog.I(copy_writing.(string))
	}

}

func roomsilent(s string){
	msglog.Base(1, "房")

	if level := p.Json().GetValFromS(s, "data.level");level == nil {
		msglog.E("level", level)
		return
	} else {
		if level.(float64) == 0 {msglog.I("主播关闭了禁言")}
		msglog.I("主播开启了等级禁言:", int64(level.(float64)))
	}
}

func roominfo(s string){
	msglog.Base(1, "粉")

	fans := p.Json().GetValFromS(s, "data.fans");
	fans_club := p.Json().GetValFromS(s, "data.fans_club");

	var sh []interface{}

	if fans != nil {
		sh = append(sh, "粉丝总人数:", int64(fans.(float64)))
	}
	if fans_club != nil {
		sh = append(sh, "粉丝团人数:", int64(fans_club.(float64)))
	}

	if len(sh) != 0 {msglog.I(sh...)}
}

func danmu(s string) {
	if info := p.Json().GetValFromS(s, "info");info == nil {
		msglog.E("info", info)
		return
	} else {
		infob := info.([]interface{})
		msg := infob[1].(string)
		auth := infob[2].([]interface{})[1].(string)

		if Autobanf(msg) > 0.5 {msglog.Base(1, "风险").I(msg)}
		if Msg_roomid != 0 && Msg_cookie != "" && msg == "弹幕机在么" {Danmu_s("在", Msg_cookie, Msg_roomid)}

		msglog.I(auth, ":", msg)
	}
}
