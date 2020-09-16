package bili_danmu

import (
	"fmt"
	"bytes"
	"compress/zlib"

	p "github.com/qydysky/part"
)
/*
	数据为WS_OP_MESSAGE类型的
*/

var msglog = p.Logf().New().Base(-1, "Msg.go").Open("danmu.log").Level(1)
var Msg_cookie string
var Msg_roomid int
var Msg_map = map[string]func(replayF, string) {
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
	"WIN_ACTIVITY":nil,
	"GUARD_BUY":replayF.guard_buy,//大航海购买
	"WELCOME_GUARD":nil,//replayF.welcome_guard,//大航海进入
	"DANMU_MSG":replayF.danmu,//弹幕
	"ROOM_CHANGE":replayF.room_change,//房间信息分区改变
	"ROOM_SILENT_OFF":replayF.roomsilent,//禁言结束
	"ROOM_SILENT_ON":replayF.roomsilent,//禁言开始
	"SEND_GIFT":replayF.send_gift,//礼物
	"ROOM_BLOCK_MSG":replayF.room_block_msg,//封禁
	"PREPARING":replayF.preparing,//下播
	"LIVE":replayF.live,//开播
	"SUPER_CHAT_MESSAGE":nil,//replayF.super_chat_message,//打赏
	"SUPER_CHAT_MESSAGE_JPN":replayF.super_chat_message,//打赏
	"PANEL":replayF.panel,//排行榜
	"ENTRY_EFFECT":replayF.entry_effect,//进入特效
	"ROOM_REAL_TIME_MESSAGE_UPDATE":nil,//replayF.roominfo,//粉丝数
}

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
			var f replayF

			if F, ok := Msg_map[cmd.(string)]; ok {
				if F != nil {F(f, s)}
			} else {
				f.defaultMsg(s)
			}
		}
	}

	return 
}

type replayF struct {}

func (replayF) defaultMsg(s string){
	msglog.Base(1, "Unknow cmd").E(s)
}

func (replayF) guard_buy(s string){
	msglog.Fileonly(true).Base(-1, "礼")
	defer msglog.Base(0).Fileonly(false)

	username := p.Json().GetValFromS(s, "data.username");
	gift_name := p.Json().GetValFromS(s, "data.gift_name");
	price := p.Json().GetValFromS(s, "data.price");

	var sh []interface{}

	if username != nil {
		sh = append(sh, username)
	}
	if gift_name != nil {
		sh = append(sh, "购买了", gift_name)
	}
	if price != nil {
		sh = append(sh, "￥", int(price.(float64)) / 1000)
	}

	fmt.Println("====")
	fmt.Println(sh...)
	fmt.Println("====")
	msglog.I(sh...)

}

func (replayF) room_change(s string){
	title := p.Json().GetValFromS(s, "data.title");
	area_name := p.Json().GetValFromS(s, "data.area_name");

	var sh  = []interface{}{"房间改变"}

	if title != nil {
		sh = append(sh, title)
	}
	if area_name != nil {
		sh = append(sh, area_name)
	}
	msglog.Base(1, "房").I(sh...)
}

func (replayF) welcome_guard(s string){

	username := p.Json().GetValFromS(s, "data.username");
	guard_level := p.Json().GetValFromS(s, "data.guard_level");

	var sh = []interface{}{"欢迎"}

	if username != nil {
		sh = append(sh, username, "进入直播间")
	}
	if guard_level != nil {
		sh = append(sh, "等级", guard_level)
	}

	msglog.Base(1, "房").I(sh...)
}

func (replayF) send_gift(s string){
	// coin_type := p.Json().GetValFromS(s, "data.coin_type");
	num := p.Json().GetValFromS(s, "data.num");
	uname := p.Json().GetValFromS(s, "data.uname");
	action := p.Json().GetValFromS(s, "data.action");
	giftName := p.Json().GetValFromS(s, "data.giftName");
	price := p.Json().GetValFromS(s, "data.price");

	var sh []interface{}
	var allprice int64

	if uname != nil {
		sh = append(sh, uname)
	}
	if action != nil {
		sh = append(sh, action)
	}
	if num != nil {
		sh = append(sh, num, "x")
	}
	if giftName != nil {
		sh = append(sh, giftName)
	}
	if price != nil {
		allprice = int64(num.(float64) * price.(float64) / 1000)
		sh = append(sh, "￥", allprice)
	}

	if len(sh) == 0 {return}

	msglog.Fileonly(true).Base(-1, "礼")
	defer msglog.Base(0).Fileonly(false)

	//小于3万金瓜子
	if allprice < 30000 {msglog.T(sh...);return}

	fmt.Println("====")
	fmt.Println(sh...)
	fmt.Println("====")
	msglog.I(sh...)
}

func (replayF) room_block_msg(s string) {
	msglog.Fileonly(true).Base(-1, "封")
	defer msglog.Base(0).Fileonly(false)

	if uname := p.Json().GetValFromS(s, "uname");uname == nil {
		msglog.E("uname", uname)
		return
	} else {
	fmt.Println("用户", uname, "已被封禁")
	msglog.I("用户", uname, "已被封禁")
	}
}

func (replayF) preparing(s string) {
	msglog.Base(1, "房")

	if roomid := p.Json().GetValFromS(s, "roomid");roomid == nil {
		msglog.E("roomid", roomid)
		return
	} else {
		msglog.I("房间", roomid, "下播了")
	}
}

func (replayF) live(s string) {
	msglog.Base(1, "房")

	if roomid := p.Json().GetValFromS(s, "roomid");roomid == nil {
		msglog.E("roomid", roomid)
		return
	} else {
		msglog.I("房间", roomid, "开播了")
	}
}

func (replayF) super_chat_message(s string){
	uname := p.Json().GetValFromS(s, "data.user_info.uname");
	price := p.Json().GetValFromS(s, "data.price");
	message := p.Json().GetValFromS(s, "data.message");
	message_jpn := p.Json().GetValFromS(s, "data.message_jpn");

	var sh = []interface{}{"打赏: "}

	if uname != nil {
		sh = append(sh, uname)
	}
	if price != nil {
		sh = append(sh, "￥", price)
	}
	if message != nil {
		sh = append(sh, message)
	}
	if message_jpn != nil && message != message_jpn {
		sh = append(sh, message_jpn)
	}
	msglog.Fileonly(true)
	defer msglog.Fileonly(false)

	fmt.Println("====")
	fmt.Println(sh...)
	fmt.Println("====")
	msglog.Base(1, "礼").I(sh...)
}

func (replayF) panel(s string){
	msglog.Fileonly(true).Base(1, "房")
	defer msglog.Fileonly(false)

	if note := p.Json().GetValFromS(s, "data.note");note == nil {
		msglog.E("note", note)
		return
	} else {
		fmt.Println("排行", note)
		msglog.I("排行", note)
	}
}

func (replayF) entry_effect(s string){
	msglog.Fileonly(true).Base(-1, "房")
	defer msglog.Base(0).Fileonly(false)

	if copy_writing := p.Json().GetValFromS(s, "data.copy_writing");copy_writing == nil {
		msglog.E("copy_writing", copy_writing)
		return
	} else {
		msglog.I(copy_writing)
		fmt.Println(copy_writing)
	}

}

func (replayF) roomsilent(s string){
	msglog.Base(1, "房")

	if level := p.Json().GetValFromS(s, "data.level");level == nil {
		msglog.E("level", level)
		return
	} else {
		if level.(float64) == 0 {msglog.I("主播关闭了禁言")}
		msglog.I("主播开启了等级禁言:", level)
	}
}

func (replayF) roominfo(s string){
	fans := p.Json().GetValFromS(s, "data.fans");
	fans_club := p.Json().GetValFromS(s, "data.fans_club");

	var sh []interface{}

	if fans != nil {
		sh = append(sh, "粉丝总人数:", fans)
	}
	if fans_club != nil {
		sh = append(sh, "粉丝团人数:", fans_club)
	}

	if len(sh) != 0 {msglog.Base(1, "粉").I(sh...)}
}

func (replayF) danmu(s string) {
	if info := p.Json().GetValFromS(s, "info");info == nil {
		msglog.E("info", info)
		return
	} else {
		infob := info.([]interface{})
		msg := infob[1].(string)
		auth := infob[2].([]interface{})[1]

		msglog.Fileonly(true)
		defer msglog.Fileonly(false)

		//F附加方法
		Danmujif(msg, Msg_cookie, Msg_roomid)
		if Autobanf(msg) > 0.5 {
			msglog.Base(1, "风险").I(msg)
			return
		}
		if i := Autoskipf(msg, 50, 15); i > 0 {
			msglog.I(auth, ":", msg)
			return
		}
		if Lessdanmuf(msg, 200) {
			msglog.I(auth, ":", msg)
			return
		}

		fmt.Println(msg)
		msglog.I(auth, ":", msg)
	}
}
