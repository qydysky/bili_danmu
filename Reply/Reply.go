package reply

import (
	"fmt"
	"time"
	"bytes"
	"strings"
	"strconv"
	"compress/zlib"
	"encoding/json"

	p "github.com/qydysky/part"
	mq "github.com/qydysky/part/msgq"
	F "github.com/qydysky/bili_danmu/F"
	ws_msg "github.com/qydysky/bili_danmu/Reply/ws_msg"
	send "github.com/qydysky/bili_danmu/Send"
	c "github.com/qydysky/bili_danmu/CV"
)

var reply_log = c.Log.Base(`Reply`)

//返回数据分派
//传入接受到的ws数据
//判断进行解压，并对每个json对象进行分派
func Reply(b []byte) {
	reply_log := reply_log.Base_add(`返回分派`)

	if len(b) <= c.WS_PACKAGE_HEADER_TOTAL_LENGTH {reply_log.L(`W: `,"包缺损");return}

	head := F.HeadChe(b[:c.WS_PACKAGE_HEADER_TOTAL_LENGTH])
	if int(head.PackL) > len(b) {reply_log.L(`E: `,"包缺损");return}

	if head.BodyV == c.WS_BODY_PROTOCOL_VERSION_DEFLATE {
		readc, err := zlib.NewReader(bytes.NewReader(b[16:]))
		if err != nil {reply_log.L(`E: `,"解压错误");return}
		defer readc.Close()
		
		buf := bytes.NewBuffer(nil)
		if _, err := buf.ReadFrom(readc);err != nil {reply_log.L(`E: `,"解压错误");return}
		b = buf.Bytes()
	}

	for len(b) != 0 {
		head := F.HeadChe(b[:c.WS_PACKAGE_HEADER_TOTAL_LENGTH])
		if int(head.PackL) > len(b) {reply_log.L(`E: `,"包缺损");return}
		
		contain := b[c.WS_PACKAGE_HEADER_TOTAL_LENGTH:head.PackL]
		switch head.OpeaT {
		case c.WS_OP_MESSAGE:
			Msg(contain)
			Save_to_json(-1, []interface{}{contain,`,`})
		case c.WS_OP_HEARTBEAT_REPLY:Heart(contain)
		default :reply_log.L(`W: `,"unknow reply", contain)
		}

		b = b[head.PackL:]
	}
}

//所有的json对象处理子函数类
//包含Msg和HeartBeat两大类
type replyF struct {}

//默认未识别Msg
func (replyF) defaultMsg(s string){
	msglog.Base_add("Unknow").L(`E: `, s)
}

//msg-特别礼物
func (replyF) vtr_gift_lottery(s string){
	msglog := msglog.Base_add("特别礼物")
	var j ws_msg.VTR_GIFT_LOTTERY
	if e := json.Unmarshal([]byte(s), &j);e != nil{
		msglog.L(`E: `, e)
		return
	}
	{//语言tts
		c.Danmu_Main_mq.Push_tag(`tts`,Danmu_mq_t{
			uid:`0room`,
			msg:fmt.Sprint(j.Data.InteractMsg),
		})
	}
	Gui_show(j.Data.InteractMsg,`0room`)
	msglog.L(`I`, j.Data.InteractMsg)
}

//msg-直播间进入信息，此处用来提示关注
func (replyF) interact_word(s string){
	msg_type := p.Json().GetValFromS(s, "data.msg_type");
	if v,ok := msg_type.(float64);!ok || v < 2 {return}//关注时为2,进入时为1
	uname := p.Json().GetValFromS(s, "data.uname");
	if v,ok := uname.(string);ok {
		{//语言tts
			c.Danmu_Main_mq.Push_tag(`tts`,Danmu_mq_t{
				uid:`0follow`,
				msg:fmt.Sprint(v + `关注了直播间`),
			})
		}
		Gui_show(v + `关注了直播间`,`0follow`)
		msglog.Base_add("房").Log_show_control(false).L(`I`, v + `关注了直播间`)
	}
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
	Gui_show(Itos(sh),`0tianxuan`)

	msglog.Base_add("房").Log_show_control(false).L(`I`, sh...)
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
				if v,ok := uid.(float64);ok {//uid可能为float型
					sh = append(sh, uname, "(", strconv.Itoa(int(v)), ")")
				} else {
					sh = append(sh, uname, "(", uid, ")")
				}
			}
		}
	}
	sh = append(sh, "]")
	{//额外 ass
		Assf(fmt.Sprintln("天选之人", award_name, "结束"))
	}
	fmt.Println(sh...)
	Gui_show(Itos(sh),`0tianxuan`)

	msglog.Base_add("房").Log_show_control(false).L(`I`, sh...)
}

//msg-通常是大航海购买续费
func (replyF) user_toast_msg(s string){
	username := p.Json().GetValFromS(s, "data.username");
	op_type := p.Json().GetValFromS(s, "data.op_type");
	uid := p.Json().GetValFromS(s, "data.uid");
	role_name := p.Json().GetValFromS(s, "data.role_name");
	num := p.Json().GetValFromS(s, "data.num");
	unit := p.Json().GetValFromS(s, "data.unit");
	// target_guard_count := p.Json().GetValFromS(s, "data.target_guard_count");
	price := p.Json().GetValFromS(s, "data.price");

	var sh []interface{}
	var sh_log []interface{}

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
			msglog.L(`W`, s)
			sh = append(sh, op_type)
		}
	}
	if num != nil {
		sh = append(sh, num, "个")
	}
	if unit != nil {
		sh = append(sh, unit)
	}
	if role_name != nil {
		sh = append(sh, role_name)
	}
	if price != nil {
		sh_log = append(sh, "￥", int(price.(float64)) / 1000)//不在界面显示价格
		c.Danmu_Main_mq.Push_tag(`c.Rev_add`,price.(float64) / 1000)
	}
	{//语言tts
		c.Danmu_Main_mq.Push_tag(`tts`,Danmu_mq_t{//传入消息队列
			uid:`0buyguide`,
			msg:fmt.Sprint(sh...),
		})
	}
	{//额外 ass 私信
		Assf(fmt.Sprintln(sh...))
		c.Danmu_Main_mq.Push_tag(`guard_update`,nil)//使用连续付费的新舰长无法区分，刷新舰长数
		if uid != 0 {
			c.Danmu_Main_mq.Push_tag(`pm`,send.Pm_item{
				Uid:int(uid.(float64)),
				Msg:c.K_v.LoadV(`上舰私信`).(string),
			})//上舰私信
		}
		if c.K_v.LoadV(`额外私信对象`).(float64) != 0 {
			c.Danmu_Main_mq.Push_tag(`pm`,send.Pm_item{
				Uid:int(c.K_v.LoadV(`额外私信对象`).(float64)),
				Msg:c.K_v.LoadV(`上舰私信(额外)`).(string),
			})//上舰私信-对额外
		}
	}
	fmt.Println("\n====")
	fmt.Println(sh...)
	fmt.Print("====\n\n")

	// Gui_show("\n====")
	Gui_show(Itos(sh), "0buyguide")
	// Gui_show("====\n")

	msglog.Base_add("礼").Log_show_control(false).L(`I: `,sh_log...)
}

//HeartBeat-心跳用来传递人气值
func (replyF) heartbeat(s int){
	c.Danmu_Main_mq.Push_tag(`c.Renqi`,s)//使用连续付费的新舰长无法区分，刷新舰长数
	if s == 1 {return}//人气为1,不输出
	reply_log.Base_add(`人气`).L(`I: `,"当前人气", s)
}

//Msg-房间特殊活动
func (replyF) win_activity(s string){
	title := p.Json().GetValFromS(s, "data.title");

	fmt.Println("活动", title, "已开启")
	msglog.Base_add("房").Log_show_control(false).L(`I: `,"活动", title, "已开启")
}

//Msg-特殊礼物，当前仅观察到节奏风暴
func (replyF) special_gift(s string){
	
	content := p.Json().GetValFromS(s, "data.39.content");
	action := p.Json().GetValFromS(s, "data.39.action");

	var sh []interface{}

	if action != nil && action.(string) == "end" {
		return
	}
	if content != nil {
		sh = append(sh, "节奏风暴", content)
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

	msglog.Base_add("礼").Log_show_control(false).L(`I: `, sh...)

}

//Msg-大航海购买，由于信息少，用user_toast_msg进行替代
func (replyF) guard_buy(s string){
	username := p.Json().GetValFromS(s, "data.username");
	gift_name := p.Json().GetValFromS(s, "data.gift_name");
	price := p.Json().GetValFromS(s, "data.price");

	var sh []interface{}
	var sh_log []interface{}

	if username != nil {
		sh = append(sh, username)
	}
	if gift_name != nil {
		sh = append(sh, "购买了", gift_name)
	}
	if price != nil {
		sh_log = append(sh, "￥", int(price.(float64)) / 1000)//不在界面显示价格
	}
	{//额外 ass
		Assf(fmt.Sprintln(sh...))
	}
	fmt.Println("\n====")
	fmt.Println(sh...)
	fmt.Print("====\n\n")
	msglog.Base_add("礼").Log_show_control(false).L(`I: `, sh_log...)

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

	msglog.Base_add("房").L(`I: `, sh...)
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
	{//语言tts
		c.Danmu_Main_mq.Push_tag(`tts`,Danmu_mq_t{//传入消息队列
			uid:img,
			msg:fmt.Sprintln(sh...),
		})
	}
	fmt.Print(">>> ")
	fmt.Println(sh...)
	Gui_show(Itos(append([]interface{}{">>> "}, sh...)), img)

	msglog.Base_add("房").Log_show_control(false).L(`I: `, sh...)
}

//Msg-礼物处理
func (replyF) send_gift(s string){
	coin_type := p.Json().GetValFromS(s, "data.coin_type");
	if coin_type != nil && coin_type == "silver" {return}

	num := p.Json().GetValFromS(s, "data.num");
	uname := p.Json().GetValFromS(s, "data.uname");
	action := p.Json().GetValFromS(s, "data.action");
	giftName := p.Json().GetValFromS(s, "data.giftName");
	total_coin := p.Json().GetValFromS(s, "data.total_coin");

	var sh []interface{}
	var sh_log []interface{}
	var allprice float64

	if uname != nil {
		sh = append(sh, uname)
	}
	if action != nil {
		sh = append(sh, action)
	}
	if num != nil {
		sh = append(sh, num, "个")
	}
	if giftName != nil {
		sh = append(sh, giftName)
	}
	if total_coin != nil {
		allprice = total_coin.(float64) / 1000
		sh_log = append(sh, fmt.Sprintf("￥%.1f",allprice))//不在界面显示价格
		c.Danmu_Main_mq.Push_tag(`c.Rev_add`,allprice)
	}

	if len(sh) == 0 {return}
	msglog := msglog.Base_add("礼").Log_show_control(false)

	//小于设定
	{
		var tmp = 20.0
		if v,ok := c.K_v.Load(`弹幕_礼物金额显示阈值`);ok {
			tmp = v.(float64)
		}
		if allprice < tmp {msglog.L(`T: `, sh_log...);return}
		msglog.L(`I: `, sh_log...);
	}

	{//语言tts
		c.Danmu_Main_mq.Push_tag(`tts`,Danmu_mq_t{//传入消息队列
			uid:`0gift`,
			msg:fmt.Sprintln(sh...),
		})
	}
	{//额外
		Assf(fmt.Sprintln(sh...))
	}
	fmt.Println("\n====")
	fmt.Println(sh...)
	fmt.Print("====\n\n")

	// Gui_show("\n====")
	Gui_show(Itos(sh), "0gift")
	// Gui_show("====\n")
}

//Msg-房间封禁信息
func (replyF) room_block_msg(s string) {
	if uname := p.Json().GetValFromS(s, "uname");uname == nil {
		msglog.L(`E: `, "uname", uname)
		return
	} else {
		Gui_show(Itos([]interface{}{"用户", uname, "已被封禁"}), "0room")
		fmt.Println("用户", uname, "已被封禁")
		msglog.Base_add("封").Log_show_control(false).L(`I: `, "用户", uname, "已被封禁")
	}
}

//Msg-房间准备信息，通常出现在下播而不出现在开播
func (replyF) preparing(s string) {
	msglog := msglog.Base_add("房")

	if roomid := p.Json().GetValFromS(s, "roomid");roomid == nil {
		msglog.L(`E: `, "roomid", roomid)
		return
	} else {
		{//附加功能 obs结束 `savestream`结束
			Obs_R(false)
			Obsf(false)
			Savestream_wait()
			go ShowRevf()
			c.Liveing = false
		}
		if p.Sys().Type(roomid) == "float64" {
			Gui_show(Itos([]interface{}{"房间", roomid, "下播了"}), "0room")
			msglog.L(`I: `, "房间", int(roomid.(float64)), "下播了")
			return
		}
		Gui_show(Itos([]interface{}{"房间", roomid, "下播了"}), "0room")
		msglog.L(`I: `, "房间", roomid, "下播了")
	}
}

//Msg-房间开播信息
func (replyF) live(s string) {
	msglog := msglog.Base_add("房")

	if roomid := p.Json().GetValFromS(s, "roomid");roomid == nil {
		msglog.L(`E: `, "roomid", roomid)
		return
	} else {
		{//附加功能 obs录播
			Obsf(true)
			Obs_R(true)
			go Savestreamf()
		}
		{
			c.Rev = 0.0 //营收
			c.Liveing = true //直播i标志
			c.Live_Start_Time = time.Now() //开播h时间
		}
		if p.Sys().Type(roomid) == "float64" {
			Gui_show(Itos([]interface{}{"房间", roomid, "开播了"}), "0room")
			msglog.L(`I: `, "房间", int(roomid.(float64)), "开播了")
			return
		}
		Gui_show(Itos([]interface{}{"房间", roomid, "开播了"}), "0room")
		msglog.L(`I: `, "房间", roomid, "开播了")
	}
}

//Msg-超级留言处理
var sc_buf = make(map[string]struct{})
func (replyF) super_chat_message(s string){
	id := p.Json().GetValFromS(s, "data.id");
	if id != nil {
		if _,ok := sc_buf[id.(string)];ok{return}
		if len(sc_buf) >= 10 {
			for k,_ := range sc_buf {delete(sc_buf, k);break}
			{//copy map
				tmp := make(map[string]struct{})
				for k,v := range sc_buf {tmp[k] = v}
				sc_buf = tmp
			}
		}
		sc_buf[id.(string)] = struct{}{}
	}
	uname := p.Json().GetValFromS(s, "data.user_info.uname");
	price := p.Json().GetValFromS(s, "data.price");
	message := p.Json().GetValFromS(s, "data.message");
	message_jpn := p.Json().GetValFromS(s, "data.message_jpn");

	var sh = []interface{}{"SC: "}

	if uname != nil {
		sh = append(sh, uname)
	}
	logg := sh
	if price != nil {
		sh = append(sh, "\n")//界面不显示价格
		logg = append(logg, "￥", price)
		c.Danmu_Main_mq.Push_tag(`c.Rev_add`,price.(float64))
	}
	fmt.Println("====")
	fmt.Println(sh...)
	// Gui_show("\n====")
	if message != nil && message.(string) != ""{
		fmt.Println(message)
		// Gui_show(message.(string))
		sh = append(sh, message)
		logg = append(logg, message)
	}
	{//语言tts
		c.Danmu_Main_mq.Push_tag(`tts`,Danmu_mq_t{//传入消息队列
			uid:`0superchat`,
			msg:fmt.Sprintln(sh...),
		})
	}
	if message_jpn != nil && message.(string) != message_jpn.(string) && message_jpn.(string) != "" {
		fmt.Println(message_jpn)
		// Gui_show(message_jpn.(string))
		sh = append(sh, message_jpn)
		logg = append(logg, message_jpn)
	}
	fmt.Print("====\n\n")
	
	{//额外
		Assf(fmt.Sprintln(sh...))
		// Gui_show("====\n")
		Gui_show(Itos(sh), "0superchat")
	}
	msglog.Base_add("礼").Log_show_control(false).L(`I: `, logg...)
}

//Msg-分区排行 使用热门榜替代
func (replyF) panel(s string){
	msglog := msglog.Base_add("房").Log_show_control(false)

	if note := p.Json().GetValFromS(s, "data.note");note == nil {
		msglog.L(`E: `, "note", note)
		return
	} else {
		if v,ok := note.(string);ok{c.Note = v}
		fmt.Println("排行", note)
		msglog.L(`I: `, "排行", note)
	}
}

//Msg-热门榜变动
func (replyF) hot_rank_changed(s string){
	msglog := msglog.Base_add("房")

	var type_item ws_msg.HOT_RANK_CHANGED
	if e := json.Unmarshal([]byte(s), &type_item);e != nil {
		msglog.L(`E: `, e)
	}
	if type_item.Data.Area_name != `` {
		c.Note = type_item.Data.Area_name + " "
		if type_item.Data.Rank == 0 {
			c.Note += "50+"
		} else {
			c.Note += strconv.Itoa(type_item.Data.Rank)
		}
		msglog.L(`I: `, "热门榜", c.Note)
	}
}

//Msg-热门榜获得
func (replyF) hot_rank_settlement(s string){
	msglog := msglog.Base_add("房")

	var type_item ws_msg.HOT_RANK_SETTLEMENT
	if e := json.Unmarshal([]byte(s), &type_item);e != nil {
		msglog.L(`E: `, e)
	}
	var tmp = `获得:`
	if type_item.Data.Area_name != `` {
		tmp += type_item.Data.Area_name + " 第"
	}
	if type_item.Data.Rank != 0 {
		tmp += strconv.Itoa(type_item.Data.Rank)
	}
	Gui_show(tmp, "0rank")
	c.Danmu_Main_mq.Push_tag(`tts`,Danmu_mq_t{//传入消息队列
		uid:"0rank",
		msg:tmp,
	})
	msglog.L(`I: `, "热门榜", tmp)
}

//Msg-小消息
func (replyF) little_message_box(s string){
	msglog := msglog.Base_add("系统")

	var type_item ws_msg.LITTLE_MESSAGE_BOX
	if e := json.Unmarshal([]byte(s), &type_item);e != nil {
		msglog.L(`E: `, e)
	}
	if type_item.Data.Msg != `` {
		msglog.L(`I: `, type_item.Data.Msg)
		if strings.Contains(type_item.Data.Msg,`小心心`) && strings.Contains(type_item.Data.Msg,`上限`) {
				F.F_x25Kn_cancel()
		}
	}
}

//Msg-粉丝牌切换
func (replyF) messagebox_user_medal_change(s string){
	msglog := msglog.Base_add("房")

	var type_item ws_msg.MESSAGEBOX_USER_MEDAL_CHANGE
	if e := json.Unmarshal([]byte(s), &type_item);e != nil {
		msglog.L(`E: `, e)
	}
	if type_item.Data.Medal_name != `` {
		msglog.L(`I: `, "粉丝牌切换至", type_item.Data.Medal_name, type_item.Data.Medal_level)
	}
}

//Msg-进入特效
func (replyF) entry_effect(s string){

	var res struct{
		Data struct{
			Copy_writing string `json:"copy_writing"`
		} `json:"data"`
	}
	if e := json.Unmarshal([]byte(s), &res);e != nil {
		msglog.L(`E: `, e)
	}
	//处理特殊字符
	copy_writing := strings.ReplaceAll(res.Data.Copy_writing, `<%`, ``)
	copy_writing = strings.ReplaceAll(copy_writing, `%>`, ``)

	img := "0default"
	if strings.Contains(copy_writing, `总督`) {
		img = "0level1"
	} else if strings.Contains(copy_writing, `提督`) {
		img = "0level2"
	} else if strings.Contains(copy_writing, `舰长`) {
		img = "0level3"
	}

	{//语言tts
		c.Danmu_Main_mq.Push_tag(`tts`,Danmu_mq_t{//传入消息队列
			uid:img,
			msg:fmt.Sprintln(copy_writing),
		})
	}
	fmt.Print(">>> ")
	fmt.Println(copy_writing)
	Gui_show(copy_writing, img)

	msglog.Base_add("房").Log_show_control(false).L(`I: `, copy_writing)
}

//Msg-房间禁言
func (replyF) roomsilent(s string){
	msglog := msglog.Base_add("房")

	if level := p.Json().GetValFromS(s, "data.level");level == nil {
		msglog.L(`E: `, "level", level)
		return
	} else {
		if level.(float64) == 0 {msglog.L(`I: `, "主播关闭了禁言")}
		msglog.L(`I: `, "主播开启了等级禁言:", level)
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

	if len(sh) != 0 {msglog.Base_add("粉").L(`I: `, sh...)}
}

//Msg-弹幕处理
func (replyF) danmu(s string) {
	if info := p.Json().GetValFromS(s, "info");info == nil {
		msglog.L(`E: `, "info", info)
		return
	} else {
		infob := info.([]interface{})
		msg := infob[1].(string)
		auth := infob[2].([]interface{})[1]
		uid := strconv.Itoa(int(infob[2].([]interface{})[0].(float64)))

		msglog := msglog.Log_show_control(false)

		{//附加功能 弹幕机 封禁 弹幕合并
			Danmujif(msg)
			if Autobanf(msg) {
				Gui_show(Itos([]interface{}{"风险", auth, ":", msg}))
				fmt.Println("风险", auth, ":", msg)
				msglog.Base_add("风险").L(`I: `, auth, ":", msg)
				return
			}
			if i := Autoskipf(msg); i > 0 {
				msglog.L(`I: `, auth, ":", msg)
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
	if missKey := F.CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	});len(missKey) != 0 || c.Roomid == 0 {
		msglog.L(`E: `,`c.Roomid == 0 || Cookie无Key:`,missKey)
		return
	}
	send.Danmu_s(msg, c.Roomid)
}

//弹幕显示
//由于额外功能有些需要显示，为了统一管理，使用此方法进行处理
func Msg_showdanmu(auth interface{}, m ...string) {
	msg := m[0]
	msglog := msglog.Log_show_control(false)
	{//附加功能 更少弹幕
		if !Lessdanmuf(msg) {
			if auth != nil {msglog.L(`I: `, auth, ":", msg)}
			return
		}
		if _msg := Shortdanmuf(msg); _msg == "" {
			if auth != nil {msglog.L(`I: `, auth, ":", msg)}
			return
		} else {msg = _msg}
		Assf(msg)//ass
		if auth != nil {
			Gui_show(fmt.Sprint(auth) +`: `+ msg, m[1])
		} else {
			Gui_show(m...)
		}	
	}
	{//语言tts 私信
		if len(m) > 1 {
			c.Danmu_Main_mq.Push_tag(`tts`,Danmu_mq_t{//传入消息队列
				uid:m[1],
				msg:msg,
			})
			if i,e := strconv.Atoi(m[1]);e == nil {
				c.Danmu_Main_mq.Push_tag(`pm`,send.Pm_item{
					Uid:i,
					Msg:c.K_v.LoadV(`弹幕私信`).(string),
				})//上舰私信
			}
			if c.K_v.LoadV(`额外私信对象`).(float64) != 0 {
				c.Danmu_Main_mq.Push_tag(`pm`,send.Pm_item{
					Uid:int(c.K_v.LoadV(`额外私信对象`).(float64)),
					Msg:c.K_v.LoadV(`弹幕私信(额外)`).(string),
				})//上舰私信-对额外
			}
		}
	}
	fmt.Println(msg)
	if auth != nil {msglog.L(`I: `, auth, ":", msg)}
}

type Danmu_mq_t struct {
	uid string
	msg string
}
var Danmu_mq = mq.New(10)

func Gui_show(m ...string){
	//m[0]:msg m[1]:uid
	uid := ""
	if len(m) > 1 {uid = m[1]}

	Danmu_mq.Push_tag(`danmu`,Danmu_mq_t{
		uid:uid,
		msg:m[0],
	})
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