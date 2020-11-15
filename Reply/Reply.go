package reply

import (
	"fmt"
	"bytes"
	"strconv"
	"compress/zlib"

	p "github.com/qydysky/part"
	F "github.com/qydysky/bili_danmu/F"
	S "github.com/qydysky/bili_danmu/Send"
	c "github.com/qydysky/bili_danmu/CV"
)

var replylog = p.Logf().New().Open("danmu.log").Base(-1, "Reply.go")

//返回数据分派
//传入接受到的ws数据
//判断进行解压，并对每个json对象进行分派
func Reply(b []byte) {
	replylog.Base(-1, "返回分派")
	defer replylog.Base(0)

	head := F.HeadChe(b[:c.WS_PACKAGE_HEADER_TOTAL_LENGTH])
	if int(head.PackL) > len(b) {replylog.E("包缺损");return}

	if head.BodyV == c.WS_BODY_PROTOCOL_VERSION_DEFLATE {
		readc, err := zlib.NewReader(bytes.NewReader(b[16:]))
		if err != nil {replylog.E("解压错误");return}
		defer readc.Close()
		
		buf := bytes.NewBuffer(nil)
		if _, err := buf.ReadFrom(readc);err != nil {replylog.E("解压错误");return}
		b = buf.Bytes()
	}

	for len(b) != 0 {
		head := F.HeadChe(b[:c.WS_PACKAGE_HEADER_TOTAL_LENGTH])
		if int(head.PackL) > len(b) {replylog.E("包缺损");return}
		
		contain := b[c.WS_PACKAGE_HEADER_TOTAL_LENGTH:head.PackL]
		switch head.OpeaT {
		case c.WS_OP_MESSAGE:Msg(contain)
		case c.WS_OP_HEARTBEAT_REPLY:Heart(contain)
		default :replylog.W("unknow reply", contain)
		}

		b = b[head.PackL:]
	}
}

//所有的json对象处理子函数类
//包含Msg和HeartBeat两大类
type replyF struct {}

//默认未识别Msg
func (replyF) defaultMsg(s string){
	msglog.Base(1, "Unknow").E(s)
}

//Msg-天选之人开始
func (replyF) anchor_lot_start(s string){
	award_name := p.Json().GetValFromS(s, "data.award_name");
	var sh = []interface{}{">天选"}
	if award_name != nil {
		sh = append(sh, award_name, "开始")
	}

	{//额外 ass
		Assf(fmt.Sprintln("天选之人", award_name, "开始"))
	}
	fmt.Println(sh...)
	Gui_show(Itos(sh))

	msglog.Base(1, "房").Fileonly(true).I(sh...).Fileonly(false)
}

//Msg-天选之人结束
func (replyF) anchor_lot_award(s string){
	award_name := p.Json().GetValFromS(s, "data.award_name");
	award_users := p.Json().GetValFromS(s, "data.award_users");

	var sh = []interface{}{">天选"}

	if award_name != nil {
		sh = append(sh, award_name, "获奖[")
	}
	if award_users != nil {
		for _,v := range award_users.([]interface{}) {
			uname := p.Json().GetValFrom(v, "uname");
			uid := p.Json().GetValFrom(v, "uid");
			if uname != nil && uid != nil {
				sh = append(sh, uname, "(", uid, ")")
			}
		}
	}
	sh = append(sh, "]")
	{//额外 ass
		Assf(fmt.Sprintln("天选之人", award_name, "结束"))
	}
	fmt.Println(sh...)
	Gui_show(Itos(sh))

	msglog.Base(1, "房").Fileonly(true).I(sh...).Fileonly(false)
}

//msg-通常是大航海购买续费
func (replyF) user_toast_msg(s string){
	username := p.Json().GetValFromS(s, "data.username");
	op_type := p.Json().GetValFromS(s, "data.op_type");
	role_name := p.Json().GetValFromS(s, "data.role_name");
	num := p.Json().GetValFromS(s, "data.num");
	unit := p.Json().GetValFromS(s, "data.unit");
	// target_guard_count := p.Json().GetValFromS(s, "data.target_guard_count");
	price := p.Json().GetValFromS(s, "data.price");

	var sh []interface{}

	if username != nil {
		sh = append(sh, username)
	}
	if op_type != nil {
		switch op_type.(float64) {
		case 1:
			sh = append(sh, "购买了")
		case 2:
			sh = append(sh, "续费了")
		case 3:
			sh = append(sh, "自动续费了")
		default:
			msglog.W(s)
			sh = append(sh, op_type)
		}
	}
	if num != nil {
		sh = append(sh, num, "x")
	}
	if unit != nil {
		sh = append(sh, unit)
	}
	if role_name != nil {
		sh = append(sh, role_name)
	}
	if price != nil {
		sh = append(sh, "￥", int(price.(float64)) / 1000)
		c.Rev += price.(float64) / 1000
	}
	{//额外 ass
		Assf(fmt.Sprintln(sh...))
	}
	fmt.Println("\n====")
	fmt.Println(sh...)
	fmt.Print("====\n\n")

	// Gui_show("\n====")
	Gui_show(Itos(sh), "0buyguide")
	// Gui_show("====\n")

	msglog.Base(1, "礼").Fileonly(true).I(sh...).Fileonly(false)
}

//HeartBeat-心跳用来传递人气值
func (replyF) heartbeat(s string){
	if s == "1" {return}//人气为1,不输出
	heartlog.I("当前人气", s)
}

//Msg-房间特殊活动
func (replyF) win_activity(s string){
	msglog.Fileonly(true)
	defer msglog.Fileonly(false)

	title := p.Json().GetValFromS(s, "data.title");

	fmt.Println("活动", title, "已开启")
	msglog.Base(1, "房").I("活动", title, "已开启")
}

//Msg-特殊礼物，当前仅观察到节奏风暴
func (replyF) special_gift(s string){
	msglog.Fileonly(true)
	defer msglog.Fileonly(false)

	content := p.Json().GetValFromS(s, "data.39.content");
	action := p.Json().GetValFromS(s, "data.39.action");

	var sh []interface{}

	if action != nil && action.(string) == "end" {
		return
	}
	if content != nil {
		sh = append(sh, "节奏风暴", content, "￥ 100")
		c.Rev += 100
	}
	{//额外
		Assf(fmt.Sprintln(sh...))
	}
	fmt.Println("\n====")
	fmt.Println(sh...)
	fmt.Print("====\n\n")

	// Gui_show("\n====")
	Gui_show(Itos(sh), "0jiezou")
	// Gui_show("====\n")

	msglog.Base(1, "礼").I(sh...)

}

//Msg-大航海购买，由于信息少，用user_toast_msg进行替代
func (replyF) guard_buy(s string){
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
	{//额外 ass
		Assf(fmt.Sprintln(sh...))
	}
	fmt.Println("\n====")
	fmt.Println(sh...)
	fmt.Print("====\n\n")

	msglog.Base(1, "礼").Fileonly(true).I(sh...).Fileonly(false)
}

//Msg-房间信息改变，标题等
func (replyF) room_change(s string){
	title := p.Json().GetValFromS(s, "data.title");
	area_name := p.Json().GetValFromS(s, "data.area_name");

	var sh  = []interface{}{"房间改变"}

	if title != nil {
		sh = append(sh, title)
		c.Title = title.(string)
	}
	if area_name != nil {
		sh = append(sh, area_name)
	}
	Gui_show(Itos(sh), "0room")

	msglog.Base(1, "房").I(sh...)
}

//Msg-大航海欢迎信息
func (replyF) welcome_guard(s string){

	username := p.Json().GetValFromS(s, "data.username");
	guard_level := p.Json().GetValFromS(s, "data.guard_level");
	img := "0default"

	var sh = []interface{}{"欢迎"}

	if guard_level != nil {
		switch guard_level.(float64) {
		case 1:sh = append(sh, "总督");img="0level1"
		case 2:sh = append(sh, "提督");img="0level2"
		case 3:sh = append(sh, "舰长");img="0level3"
		default:sh = append(sh, "等级", guard_level)
		}
	}
	if username != nil {
		sh = append(sh, username, "进入直播间")
	}

	fmt.Print(">>> ")
	fmt.Println(sh...)
	Gui_show(Itos(append([]interface{}{">>> "}, sh...)), img)

	msglog.Base(1, "房").Fileonly(true).I(sh...).Fileonly(false)
}

//Msg-礼物处理，对于小于30人民币的礼物不显示
func (replyF) send_gift(s string){
	coin_type := p.Json().GetValFromS(s, "data.coin_type");
	if coin_type != nil && coin_type == "silver" {return}

	num := p.Json().GetValFromS(s, "data.num");
	uname := p.Json().GetValFromS(s, "data.uname");
	action := p.Json().GetValFromS(s, "data.action");
	giftName := p.Json().GetValFromS(s, "data.giftName");
	total_coin := p.Json().GetValFromS(s, "data.total_coin");

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
	if total_coin != nil {
		allprice = int64(total_coin.(float64) / 1000)
		sh = append(sh, "￥", allprice)
		c.Rev += total_coin.(float64) / 1000
	}

	if len(sh) == 0 {return}

	//小于3万金瓜子
	if allprice < 30 {msglog.T(sh...);return}
	{//额外
		Assf(fmt.Sprintln(sh...))
	}
	fmt.Println("\n====")
	fmt.Println(sh...)
	fmt.Print("====\n\n")

	// Gui_show("\n====")
	Gui_show(Itos(sh), "0gift")
	// Gui_show("====\n")
	
	msglog.Base(1, "礼").Fileonly(true).I(sh...).Fileonly(false)
}

//Msg-房间封禁信息
func (replyF) room_block_msg(s string) {
	msglog.Fileonly(true)
	defer msglog.Fileonly(false)

	if uname := p.Json().GetValFromS(s, "uname");uname == nil {
		msglog.E("uname", uname)
		return
	} else {
		Gui_show(Itos([]interface{}{"用户", uname, "已被封禁"}), "0room")
		fmt.Println("用户", uname, "已被封禁")
		msglog.Base(1, "封").I("用户", uname, "已被封禁")
	}
}

//Msg-房间准备信息，通常出现在下播而不出现在开播
func (replyF) preparing(s string) {
	msglog.Base(1, "房")

	if roomid := p.Json().GetValFromS(s, "roomid");roomid == nil {
		msglog.E("roomid", roomid)
		return
	} else {
		{//附加功能 obs结束 saveflv结束
			Obs_R(false)
			Obsf(false)
			Saveflv_wait()
			c.Rev = 0
			go ShowRevf()
		}
		if p.Sys().Type(roomid) == "float64" {
			Gui_show(Itos([]interface{}{"房间", roomid, "下播了"}), "0room")
			msglog.I("房间", int(roomid.(float64)), "下播了")
			return
		}
		Gui_show(Itos([]interface{}{"房间", roomid, "下播了"}), "0room")
		msglog.I("房间", roomid, "下播了")
	}
}

//Msg-房间开播信息
func (replyF) live(s string) {
	msglog.Base(1, "房")

	if roomid := p.Json().GetValFromS(s, "roomid");roomid == nil {
		msglog.E("roomid", roomid)
		return
	} else {
		{//附加功能 obs录播
			Obsf(true)
			Obs_R(true)
			go Saveflvf()
		}
		if p.Sys().Type(roomid) == "float64" {
			Gui_show(Itos([]interface{}{"房间", roomid, "开播了"}), "0room")
			msglog.I("房间", int(roomid.(float64)), "开播了")
			return
		}
		Gui_show(Itos([]interface{}{"房间", roomid, "开播了"}), "0room")
		msglog.I("房间", roomid, "开播了")
	}
}

//Msg-超级留言处理
var sc_buf = make(map[string]bool)
func (replyF) super_chat_message(s string){
	id := p.Json().GetValFromS(s, "data.id");
	if id != nil {
		if _,ok := sc_buf[id.(string)];ok{return}
		if len(sc_buf) >= 10 {
			for k,_ := range sc_buf {delete(sc_buf, k);break}
		}
		sc_buf[id.(string)] = true
	}
	uname := p.Json().GetValFromS(s, "data.user_info.uname");
	price := p.Json().GetValFromS(s, "data.price");
	message := p.Json().GetValFromS(s, "data.message");
	message_jpn := p.Json().GetValFromS(s, "data.message_jpn");

	var sh = []interface{}{"SC: "}

	if uname != nil {
		sh = append(sh, uname)
	}
	if price != nil {
		sh = append(sh, "￥", price, "\n")
		c.Rev += price.(float64)
	}
	fmt.Println("====")
	fmt.Println(sh...)
	// Gui_show("\n====")
	if message != nil && message.(string) != ""{
		fmt.Println(message)
		// Gui_show(message.(string))
		sh = append(sh, message)
	}
	if message_jpn != nil && message.(string) != message_jpn.(string) && message_jpn.(string) != "" {
		fmt.Println(message_jpn)
		// Gui_show(message_jpn.(string))
		sh = append(sh, message_jpn)
	}
	fmt.Print("====\n\n")
	
	{//额外
		Assf(fmt.Sprintln(sh...))
		// Gui_show("====\n")
		Gui_show(Itos(sh), "0superchat")
	}
	msglog.Base(1, "礼").Fileonly(true).I(sh...).Fileonly(false)
}

//Msg-分区排行
func (replyF) panel(s string){
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

//Msg-进入特效，大多为大航海进入，信息少，使用welcome_guard替代
func (replyF) entry_effect(s string){
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

//Msg-房间禁言
func (replyF) roomsilent(s string){
	msglog.Base(1, "房")

	if level := p.Json().GetValFromS(s, "data.level");level == nil {
		msglog.E("level", level)
		return
	} else {
		if level.(float64) == 0 {msglog.I("主播关闭了禁言")}
		msglog.I("主播开启了等级禁言:", level)
	}
}

//Msg-粉丝信息，常刷屏，不显示
func (replyF) roominfo(s string){
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

//Msg-弹幕处理
func (replyF) danmu(s string) {
	if info := p.Json().GetValFromS(s, "info");info == nil {
		msglog.E("info", info)
		return
	} else {
		infob := info.([]interface{})
		msg := infob[1].(string)
		auth := infob[2].([]interface{})[1]
		uid := strconv.Itoa(int(infob[2].([]interface{})[0].(float64)))

		msglog.Fileonly(true)
		defer msglog.Fileonly(false)

		{//附加功能 弹幕机 封禁 弹幕合并
			Danmujif(msg)
			if Autobanf(msg) {
				Gui_show(Itos([]interface{}{"风险", auth, ":", msg}))
				fmt.Println("风险", auth, ":", msg)
				msglog.Base(1, "风险").I(auth, ":", msg)
				return
			}
			if i := Autoskipf(msg, 50, 15); i > 0 {
				msglog.I(auth, ":", msg)
				return
			}
		}
		Msg_showdanmu(auth, msg, uid)
	}
}

//弹幕发送
//传入字符串即可发送
//需要cookie
func Msg_senddanmu(msg string){
	if c.Cookie == "" || c.Roomid == 0 {return}
	S.Danmu_s(msg, c.Cookie, c.Roomid)
}

//弹幕显示
//由于额外功能有些需要显示，为了统一管理，使用此方法进行处理
func Msg_showdanmu(auth interface{}, m ...string) {
	msg := m[0]
	{//附加功能 更少弹幕
		if Lessdanmuf(msg, 20) > 0.7 {//与前20条弹幕重复的字数占比度>0.7的屏蔽
			if auth != nil {msglog.I(auth, ":", msg)}
			return
		}
		if _msg := Shortdanmuf(msg); _msg == "" {
			if auth != nil {msglog.I(auth, ":", msg)}
			return
		} else {msg = _msg}
		Assf(msg)//ass
		if auth != nil {
			Gui_show(fmt.Sprint(auth) +`: `+ msg, m[1])
		} else {
			Gui_show(msg)
		}	
	}
	
	fmt.Println(msg)
	if auth != nil {msglog.I(auth, ":", msg)}
}

func Gui_show(m ...string){
	//m[0]:msg m[1]:uid
	if Gtk_on {
		if len(m) > 1 {
			Gtk_danmuChan_uid <- m[1]
		} else {Gtk_danmuChan_uid <- ""}
		Gtk_danmuChan <- m[0]
	}
}

func Itos(i []interface{}) string {
	r := ""
	for _,v := range i {
		switch v.(type) {
		case string:r += v.(string)
		case int:r += strconv.Itoa(v.(int))
		case int64: r += strconv.Itoa(int(v.(int64)))
		case float64: r+= strconv.Itoa(int(v.(float64)))
		default:fmt.Println("unkonw type", v)
		}
		r += " "
	}
	return r
}