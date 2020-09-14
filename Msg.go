package bili_danmu

import (
	"bytes"
	"compress/zlib"

	p "github.com/qydysky/part"
)
/*
	数据为WS_OP_MESSAGE类型的
*/

var msglog = p.Logf().New().Open("danmu.log").Level(0)

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
			msglog.E("->", "cmd", s)
			return
		} else {
			switch cmd.(string) {
			case "COMBO_SEND":;
			case "INTERACT_WORD":;
			case "ACTIVITY_BANNER_UPDATE_V2":;
			case "SEND_GIFT":;//礼物
			case "NOTICE_MSG":;//礼物公告
			case "ROOM_BANNER":;//未知
			case "ONLINERANK":;//未知
			case "WELCOME":;//进入提示
			case "ROOM_SILENT_OFF", "ROOM_SILENT_ON":;
			case "HOUR_RANK_AWARDS":;
			case "ROOM_RANK":;
			case "WELCOME_GUARD":;
			case "GUARD_BUY":;
			case "ROOM_SHIELD":;
			case "USER_TOAST_MSG":;
			case "ROOM_BLOCK_MSG":room_block_msg(s)
			case "PREPARING":preparing(s)
			case "LIVE":live(s)
			case "SUPER_CHAT_MESSAGE", "SUPER_CHAT_MESSAGE_JPN":super_chat_message(s)
			case "PANEL":panel(s)
			case "ENTRY_EFFECT":entry_effect(s)
			case "ROOM_REAL_TIME_MESSAGE_UPDATE":roominfo(s)
			case "DANMU_MSG":danmu(s)
			default:msglog.I("Unknow cmd", s)
			}
		}
	}

	return 
}

func room_block_msg(s string) {
	if uname := p.Json().GetValFromS(s, "uname");uname == nil {
		msglog.E("->", "uname", uname)
		return
	} else {
		msglog.I("用户", uname.(string), "已被封禁")
	}
}

func preparing(s string) {
	if roomid := p.Json().GetValFromS(s, "roomid");roomid == nil {
		msglog.E("->", "roomid", roomid)
		return
	} else {
		msglog.I("房间", roomid.(string), "下播了")
	}
}

func live(s string) {
	if roomid := p.Json().GetValFromS(s, "roomid");roomid == nil {
		msglog.E("->", "roomid", roomid)
		return
	} else {
		msglog.I("房间", roomid.(string), "开播了")
	}
}

func super_chat_message(s string){
	uname := p.Json().GetValFromS(s, "data.user_info.uname");
	price := p.Json().GetValFromS(s, "data.price");
	message := p.Json().GetValFromS(s, "data.message");
	message_jpn := p.Json().GetValFromS(s, "data.message_jpn");

	var sh []interface{}

	if uname != nil {
		sh = append(sh, []interface{}{uname.(string)})
	}
	if price != nil {
		sh = append(sh, []interface{}{"￥", int64(price.(float64))})
	}
	if message != nil {
		sh = append(sh, []interface{}{message.(string)})
	}
	if message_jpn != nil {
		sh = append(sh, []interface{}{message_jpn.(string)})
	}

	if len(sh) != 0 {msglog.I("打赏: ", sh)}
}

func panel(s string){
	if note := p.Json().GetValFromS(s, "data.note");note == nil {
		msglog.E("->", "note", note)
		return
	} else {
		msglog.I(note.(string))
	}

}

func entry_effect(s string){
	if copy_writing := p.Json().GetValFromS(s, "data.copy_writing");copy_writing == nil {
		msglog.E("->", "copy_writing", copy_writing)
		return
	} else {
		msglog.I(copy_writing.(string))
	}

}

func roomsilent(s string){
	if level := p.Json().GetValFromS(s, "data.level");level == nil {
		msglog.E("->", "level", level)
		return
	} else {
		if level.(float64) == 0 {msglog.I("主播关闭了禁言")}
		msglog.I("主播开启了等级禁言:", int64(level.(float64)))
	}

}

func roominfo(s string){
	fans := p.Json().GetValFromS(s, "data.fans");
	fans_club := p.Json().GetValFromS(s, "data.fans_club");

	var sh []interface{}

	if fans != nil {
		sh = append(sh, []interface{}{"粉丝总人数:", int64(fans.(float64))})
	}
	if fans_club != nil {
		sh = append(sh, []interface{}{"粉丝团人数:", int64(fans_club.(float64))})
	}

	if len(sh) != 0 {msglog.I(sh)}
}

func danmu(s string) {
	if info := p.Json().GetValFromS(s, "info");info == nil {
		msglog.E("->", "info", info)
		return
	} else {
		infob := info.([]interface{})
		msg := infob[1].(string)
		auth := infob[2].([]interface{})[1].(string)
		msglog.I(auth, ":", msg)
	}
}
