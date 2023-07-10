package reply

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	brotli "github.com/andybalholm/brotli"
	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	ws_msg "github.com/qydysky/bili_danmu/Reply/ws_msg"
	send "github.com/qydysky/bili_danmu/Send"
	p "github.com/qydysky/part"
	mq "github.com/qydysky/part/msgq"
	pstrings "github.com/qydysky/part/strings"
)

var reply_log = c.C.Log.Base(`Reply`)

// 返回数据分派
// 传入接受到的ws数据
// 判断进行解压，并对每个json对象进行分派
func Reply(b []byte) {
	reply_log := reply_log.Base_add(`返回分派`)

	if len(b) <= c.WS_PACKAGE_HEADER_TOTAL_LENGTH {
		reply_log.L(`W: `, "包缺损")
		return
	}

	head := F.HeadChe(b[:c.WS_PACKAGE_HEADER_TOTAL_LENGTH])
	if int(head.PackL) > len(b) {
		reply_log.L(`E: `, "首包缺损")
		return
	}

	switch head.BodyV {
	case 1: // 心跳
	case c.WS_BODY_PROTOCOL_VERSION_NORMAL: // 无加密
	case c.WS_BODY_PROTOCOL_VERSION_DEFLATE: // DEFLATE
		readc, err := zlib.NewReader(bytes.NewReader(b[16:]))
		if err != nil {
			reply_log.L(`E: `, "解压错误")
			return
		}
		defer readc.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := buf.ReadFrom(readc); err != nil {
			reply_log.L(`E: `, "解压错误")
			return
		}
		b = buf.Bytes()
	case c.WS_BODY_PROTOCOL_VERSION_BROTLI: // BROTLI
		readc := brotli.NewReader(bytes.NewReader(b[16:]))

		buf := bytes.NewBuffer(nil)
		if _, err := buf.ReadFrom(readc); err != nil {
			reply_log.L(`E: `, "解压错误")
			return
		}
		b = buf.Bytes()
	default:
		reply_log.L(`E: `, "未知的编码方式", head.BodyV)
	}

	for len(b) != 0 {
		head := F.HeadChe(b[:c.WS_PACKAGE_HEADER_TOTAL_LENGTH])
		if int(head.PackL) > len(b) {
			reply_log.L(`E: `, "包缺损")
			return
		}

		contain := b[c.WS_PACKAGE_HEADER_TOTAL_LENGTH:head.PackL]
		switch head.OpeaT {
		case c.WS_OP_MESSAGE:
			Msg(contain)
			SaveToJson.Write(contain)
		case c.WS_OP_HEARTBEAT_REPLY: //心跳响应
			Heart(contain)
			return //忽略剩余内容
		default:
			reply_log.L(`W: `, "unknow reply", contain)
		}

		b = b[head.PackL:]
	}
}

// 所有的json对象处理子函数类
// 包含Msg和HeartBeat两大类
type replyF struct{}

// 默认未识别Msg
func (replyF) defaultMsg(s string) {
	msglog.Base_add("Unknow").L(`E: `, s)
}

// 登录提示
func (replyF) log_in_notice(s string) {
	msglog := msglog.Base_add("房")
	var j ws_msg.LOG_IN_NOTICE
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
		return
	}
	Gui_show(j.Data.NoticeMsg, `0room`)
	msglog.L(`I: `, j.Data.NoticeMsg)
}

// 超管切直播
func (replyF) cut_off(s string) {
	msglog := msglog.Base_add("房", "超管")
	var j ws_msg.CUT_OFF
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
		return
	}

	j.Msg = fmt.Sprint(j.Msg)

	//直播流服务弹幕
	SendStreamWs(Danmu_item{
		auth:   "超管",
		border: true,
		color:  "#FF0000",
		msg:    j.Msg,
	})
	Gui_show(j.Msg, `0room`)
	msglog.L(`I: `, "超管切断了直播，理由:"+j.Msg)
}

// 大乱斗pk开始
func (replyF) pk_lottery_start(s string) {
	msglog := msglog.Base_add("大乱斗")
	var j ws_msg.PK_LOTTERY_START
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
		return
	}
	Gui_show(j.Data.Title, `0room`)
	msglog.L(`I: `, j.Data.Title)
}

// 连麦人状态
func (replyF) voice_join_status(s string) {
	msglog := msglog.Base_add("连麦")
	var j ws_msg.VOICE_JOIN_STATUS
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
		return
	}
	if j.Data.UserName == `` {
		return
	}

	Gui_show(`连麦中:`+j.Data.UserName, `0room`)
	msglog.L(`I: `, `连麦中:`, j.Data.UserName)
}

// 连麦等待
func (replyF) voice_join_room_count_info(s string) {
	msglog := msglog.Base_add("连麦")
	var j ws_msg.VOICE_JOIN_ROOM_COUNT_INFO
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
		return
	}
	Gui_show(`连麦等待:`+strconv.Itoa(j.Data.ApplyCount), `0room`)
	msglog.L(`I: `, `连麦等待人数`, j.Data.ApplyCount)
}

// 大乱斗pk状态
func (replyF) pk_battle_process_new(s string) {
	msglog := msglog.Base_add("大乱斗")
	var j ws_msg.PK_BATTLE_PROCESS_NEW
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
		return
	}
	if diff := j.Data.InitInfo.Votes - j.Data.MatchInfo.Votes; diff < 0 {
		Gui_show(`大乱斗落后`, strconv.Itoa(diff), `0room`)
		msglog.L(`I: `, `大乱斗落后`, diff)
	}
}

// msg-特别礼物
func (replyF) vtr_gift_lottery(s string) {
	msglog := msglog.Base_add("特别礼物")
	var j ws_msg.VTR_GIFT_LOTTERY
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
		return
	}
	{ //语言tts
		c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{
			uid: `0room`,
			m: map[string]string{
				`{msg}`: j.Data.InteractMsg,
			},
		})
	}
	Gui_show(j.Data.InteractMsg, `0room`)
	msglog.L(`I: `, j.Data.InteractMsg)
}

// msg-直播间进入信息，此处用来提示关注
func (replyF) interact_word(s string) {
	msg_type := p.Json().GetValFromS(s, "data.msg_type")
	if v, ok := msg_type.(float64); !ok || v < 2 {
		return
	} //关注时为2,进入时为1
	uname := p.Json().GetValFromS(s, "data.uname")
	if v, ok := uname.(string); ok {
		{ //语言tts
			c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{
				uid: `0follow`,
				msg: fmt.Sprint(v + `关注了直播间`),
			})
		}
		Gui_show(v+`关注了直播间`, `0follow`)
		msglog.Base_add("房").Log_show_control(false).L(`I`, v+`关注了直播间`)
	}
}

// Msg-天选之人开始
func (replyF) anchor_lot_start(s string) {
	award_name := p.Json().GetValFromS(s, "data.award_name")
	var sh = []interface{}{">天选"}
	if award_name != nil {
		sh = append(sh, award_name, "开始")
	}

	{ //额外 ass
		Assf(fmt.Sprintln("天选之人", award_name, "开始"))
	}
	fmt.Println(sh...)
	Gui_show(Itos(sh), `0tianxuan`)

	msglog.Base_add("房").Log_show_control(false).L(`I`, sh...)
}

// Msg-天选之人结束
func (replyF) anchor_lot_award(s string) {
	award_name := p.Json().GetValFromS(s, "data.award_name")
	award_users := p.Json().GetValFromS(s, "data.award_users")

	var sh = []interface{}{">天选"}

	if award_name != nil {
		sh = append(sh, award_name, "获奖[")
	}
	if award_users != nil {
		for _, v := range award_users.([]interface{}) {
			uname := p.Json().GetValFrom(v, "uname")
			uid := p.Json().GetValFrom(v, "uid")
			if uname != nil && uid != nil {
				if v, ok := uid.(float64); ok { //uid可能为float型
					sh = append(sh, uname, "(", strconv.Itoa(int(v)), ")")
				} else {
					sh = append(sh, uname, "(", uid, ")")
				}
			}
		}
	}
	sh = append(sh, "]")
	{ //额外 ass
		Assf(fmt.Sprintln("天选之人", award_name, "结束"))
	}
	fmt.Println(sh...)
	Gui_show(Itos(sh), `0tianxuan`)

	msglog.Base_add("房").Log_show_control(false).L(`I`, sh...)
}

// msg-通常是大航海购买续费
func (replyF) user_toast_msg(s string) {
	msglog := msglog.Base_add("礼", "大航海")

	var j ws_msg.USER_TOAST_MSG
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
		return
	}

	username := j.Data.Username
	op_type := j.Data.OpType
	uid := j.Data.UID
	role_name := j.Data.RoleName
	num := j.Data.Num
	unit := j.Data.Unit
	// target_guard_count := j.Data.target_guard_count
	price := j.Data.Price

	var sh []interface{}
	var sh_log []interface{}

	if username != "" {
		sh = append(sh, username)
	}
	op_name := ""
	switch op_type {
	case 1:
		op_name = `购买`
		sh = append(sh, "购买了")
	case 2:
		op_name = `续费`
		sh = append(sh, "续费了")
	case 3:
		op_name = `自动续费`
		sh = append(sh, "自动续费了")
	default:
		msglog.L(`W`, s)
		sh = append(sh, op_type)
	}
	if num != 0 {
		sh = append(sh, num, "个")
	}
	if unit != "" {
		sh = append(sh, unit)
	}
	if role_name != "" {
		sh = append(sh, role_name)
	}
	if price != 0 {
		sh_log = append(sh, "￥", price/1000) //不在界面显示价格
		c.C.Danmu_Main_mq.Push_tag(`c.Rev_add`, float64(price)/1000)
	}
	{ //语言tts
		c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{ //传入消息队列
			uid: `0buyguide`,
			m: map[string]string{
				`{username}`:  username,
				`{op_name}`:   op_name,
				`{role_name}`: role_name,
				`{num}`:       strconv.Itoa(num),
				`{unit}`:      unit,
			},
		})
	}
	{ //额外 ass 私信
		Assf(fmt.Sprintln(sh...))
		c.C.Danmu_Main_mq.Push_tag(`guard_update`, nil) //使用连续付费的新舰长无法区分，刷新舰长数
		if msg := c.C.K_v.LoadV(`上舰私信`).(string); uid != 0 && msg != "" {
			c.C.Danmu_Main_mq.Push_tag(`pm`, send.Pm_item{
				Uid: uid,
				Msg: msg,
			}) //上舰私信
		}
		if msg, uid := c.C.K_v.LoadV(`上舰私信(额外)`).(string), c.C.K_v.LoadV(`额外私信对象`).(float64); uid != 0 && msg != "" {
			c.C.Danmu_Main_mq.Push_tag(`pm`, send.Pm_item{
				Uid: int(uid),
				Msg: msg,
			}) //上舰私信-对额外
		}
	}
	fmt.Println("\n====")
	fmt.Println(sh...)
	fmt.Print("====\n\n")

	// Gui_show("\n====")
	Gui_show(Itos(sh), "0buyguide")
	// Gui_show("====\n")

	msglog.Log_show_control(false).L(`I: `, sh_log...)
}

// HeartBeat-心跳用来传递人气值
var (
	cuRoom        int
	renqi_last    int
	renqi_old     int
	watched_old   float64
	onlinenum_old int
	renqi_l       float64
	watched_l     float64
	onlinenum_l   float64
)

func (replyF) heartbeat(s int) {
	if v, ok := c.C.K_v.LoadV("下播后不记录人气观看人数").(bool); ok && v && !c.C.Liveing {
		return
	}
	c.C.Danmu_Main_mq.Push_tag(`c.Renqi`, s) //使用连续付费的新舰长无法区分，刷新舰长数
	if s == 1 {
		return
	} //人气为1,不输出
	if c.C.Roomid != cuRoom {
		cuRoom = c.C.Roomid
		renqi_last = 0
		renqi_old = 0
		watched_old = 0
		onlinenum_old = 0
		renqi_l = 0
		watched_l = 0
		onlinenum_l = 0
	}
	if renqi_last != s {
		var (
			tmp         string
			watchPerMin float64
			tmp1        string
			tmp2        string
		)
		if time.Since(c.C.Live_Start_Time) > time.Minute {
			watchPerMin = float64(c.C.Watched) / float64(time.Since(c.C.Live_Start_Time)/time.Minute)
		}
		if renqi_old != 0 {
			renqi_l = (renqi_l + 100*float64(s-renqi_old)/float64(renqi_old)) / 2
			if renqi_l > 0 {
				tmp = `+`
			} else if renqi_l == 0 {
				tmp = `=`
			}
			tmp += fmt.Sprintf("%.1f%%", renqi_l)
			tmp = `(` + tmp + `)`
		} else {
			tmp = "(=0.0%)"
		}
		if watched_old != 0 {
			watched_l = (watched_l + 100*float64(watchPerMin-watched_old)/float64(watched_old)) / 2
			if watched_l > 0 {
				tmp1 = `+`
			} else if watched_l == 0 {
				tmp1 = `=`
			}
			tmp1 += fmt.Sprintf("%.1f%%", watched_l)
			tmp1 = `(` + tmp1 + `)`
		} else {
			tmp1 = "(=0.0%)"
		}
		if onlinenum_old != 0 {
			onlinenum_l = (onlinenum_l + 100*float64(c.C.OnlineNum-onlinenum_old)/float64(onlinenum_old)) / 2
			if onlinenum_l > 0 {
				tmp2 = `+`
			} else if onlinenum_l == 0 {
				tmp2 = `=`
			}
			tmp2 += fmt.Sprintf("%.1f%%", onlinenum_l)
			tmp2 = `(` + tmp2 + `)`
		} else {
			tmp2 = "(=0.0%)"
		}
		fmt.Printf("+----\n|当前人气:%s%d\n|平均观看:%s%d\n|在线人数:%s%d\n+----\n", tmp, s, tmp1, int(watchPerMin), tmp2, c.C.OnlineNum)
		renqi_old = s
		watched_old = watchPerMin
		onlinenum_old = c.C.OnlineNum
	}
	renqi_last = s
	reply_log.Base_add(`人气`).Log_show_control(false).L(`I: `, "当前人气", s)
}

// Msg-房间特殊活动
func (replyF) win_activity(s string) {
	title := p.Json().GetValFromS(s, "data.title")

	fmt.Println("活动", title, "已开启")
	msglog.Base_add("房").Log_show_control(false).L(`I: `, "活动", title, "已开启")
}

// Msg-观看人数
func (replyF) watched_change(s string) {
	if v, ok := c.C.K_v.LoadV("下播后不记录人气观看人数").(bool); ok && v && !c.C.Liveing {
		return
	}
	var data ws_msg.WATCHED_CHANGE
	_ = json.Unmarshal([]byte(s), &data)
	// fmt.Printf("\t观看人数:%d\n", watched)
	if data.Data.Num == c.C.Watched {
		return
	}
	c.C.Watched = data.Data.Num
	var pperm = float64(c.C.Watched) / float64(time.Since(c.C.Live_Start_Time)/time.Minute)
	msglog.Base_add("房").Log_show_control(false).L(`I: `, "观看人数", data.Data.Num, fmt.Sprintf(" avg:%.1f人/分", pperm))
}

// Msg-特殊礼物，当前仅观察到节奏风暴
func (replyF) special_gift(s string) {

	content := p.Json().GetValFromS(s, "data.39.content")
	action := p.Json().GetValFromS(s, "data.39.action")

	var sh []interface{}

	if action != nil && action.(string) == "end" {
		return
	}
	if content != nil {
		sh = append(sh, "节奏风暴", content)
	}
	{ //额外
		Assf(fmt.Sprintln(sh...))
	}
	fmt.Println("\n====")
	fmt.Println(sh...)
	fmt.Print("====\n\n")

	// Gui_show("\n====")
	Gui_show(Itos(sh), "0jiezou")
	// Gui_show("====\n")

	msglog.Base_add("礼", "节奏风暴").Log_show_control(false).L(`I: `, sh...)

}

// Msg-房间信息改变，标题等
func (replyF) room_change(s string) {
	title := p.Json().GetValFromS(s, "data.title")
	area_name := p.Json().GetValFromS(s, "data.area_name")

	var sh = []interface{}{"房间改变"}

	if c.C.Title != title.(string) {
		StreamOCut(c.C.Roomid, title.(string))
	}

	if title != nil {
		sh = append(sh, title)
		c.C.Title = title.(string)
	}
	if area_name != nil {
		sh = append(sh, area_name)
	}
	Gui_show(Itos(sh), "0room")

	msglog.Base_add("房").L(`I: `, sh...)
}

// Msg-超管警告
func (replyF) warning(s string) {
	var type_item ws_msg.WARNING

	if e := json.Unmarshal([]byte(s), &type_item); e != nil {
		msglog.L(`E: `, e)
	}

	s, _ = pstrings.UnescapeUnicode(type_item.Msg)

	Gui_show(s, "0room")

	msglog.Base_add("房", "超管").L(`I: `, s)
}

// Msg-为主播点赞了
func (replyF) like_info_v3_click(s string) {
	var type_item ws_msg.LIKE_INFO_V3_CLICK

	if e := json.Unmarshal([]byte(s), &type_item); e != nil {
		msglog.L(`E: `, e)
	}
	s = type_item.Data.Uname + type_item.Data.LikeText

	Gui_show(s, "0room")

	msglog.Base_add("房").L(`I: `, s)
}

// Msg-小提示窗口
func (replyF) little_tips(s string) {
	var type_item ws_msg.LITTLE_TIPS

	if e := json.Unmarshal([]byte(s), &type_item); e != nil {
		msglog.L(`E: `, e)
	}
	s = type_item.Data.Msg

	Gui_show(s, "0room")

	msglog.Base_add("房").L(`I: `, s)
}

// Msg-人气排名
// func (replyF) popular_rank_changed(s string) {
// 	var type_item ws_msg.POPULAR_RANK_CHANGED

// 	if e := json.Unmarshal([]byte(s), &type_item); e != nil {
// 		msglog.L(`E: `, e)
// 	}
// 	s = fmt.Sprintf("人气排行 %d", type_item.Data.Rank)

// 	Gui_show(s, "0room")

// 	msglog.Base_add("房").L(`I: `, s)
// }

// Msg-开始了视频连线
func (replyF) video_connection_join_start(s string) {
	msglog := msglog.Base_add("房").Log_show_control(false)

	var j ws_msg.VIDEO_CONNECTION_JOIN_START
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
	}

	var tmp = `开始了与` + j.Data.InvitedUname + `的视频连线`

	Gui_show(tmp, "0room")

	msglog.Base_add("房").L(`I: `, tmp)
}

// Msg-结束了视频连线
func (replyF) video_connection_join_end(s string) {
	msglog := msglog.Base_add("房").Log_show_control(false)

	var j ws_msg.VIDEO_CONNECTION_JOIN_END
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
	}

	var tmp = j.Data.Toast

	Gui_show(tmp, "0room")

	msglog.Base_add("房").L(`I: `, tmp)
}

// Msg-视频连线状态改变
func (replyF) video_connection_msg(s string) {
	msglog := msglog.Base_add("房").Log_show_control(false)

	var j ws_msg.VIDEO_CONNECTION_MSG
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
	}

	var tmp = j.Data.Toast

	Gui_show(tmp, "0room")

	msglog.Base_add("房").L(`I: `, tmp)
}

// Msg-活动标题改变v2
func (replyF) activity_banner_change_v2(s string) {
	msglog := msglog.Base_add("房").Log_show_control(false)

	var j ws_msg.ACTIVITY_BANNER_CHANGE_V2
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
	}

	if len(j.Data.List) == 0 {
		return
	}

	var tmp = `修改了活动标题 ` + j.Data.List[0].ActivityTitle

	Gui_show(tmp, "0room")

	msglog.Base_add("房").L(`I: `, tmp)
}

// Msg-礼物处理
func (replyF) send_gift(s string) {
	msglog := msglog.Base_add("礼").Log_show_control(false)

	var j ws_msg.SEND_GIFT
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
	}

	//忽略银瓜子
	if j.Data.CoinType == "silver" {
		return
	}

	num := j.Data.Num
	uname := j.Data.Uname
	action := j.Data.Action
	giftName := j.Data.Giftname
	total_coin := j.Data.TotalCoin

	var sh []interface{}
	var sh_log []interface{}
	var allprice float64

	sh = append(sh, uname)
	sh = append(sh, action)
	sh = append(sh, num, "个")
	sh = append(sh, giftName)

	if total_coin != 0 {
		allprice = float64(total_coin) / 1000
		sh_log = append(sh, fmt.Sprintf("￥%.1f", allprice)) //不在界面显示价格
		c.C.Danmu_Main_mq.Push_tag(`c.Rev_add`, allprice)
	}

	if len(sh) == 0 {
		return
	}

	//小于设定
	{
		var tmp = 20.0
		if v, ok := c.C.K_v.Load(`弹幕_礼物金额显示阈值`); ok {
			tmp = v.(float64)
		}
		if allprice < tmp {
			msglog.L(`T: `, sh_log...)
			return
		}
		msglog.L(`I: `, sh_log...)
	}

	{ //语言tts
		c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{ //传入消息队列
			uid: `0gift`,
			m: map[string]string{
				`{num}`:      strconv.Itoa(num),
				`{uname}`:    uname,
				`{action}`:   action,
				`{giftName}`: giftName,
			},
		})
	}
	{ //额外
		Assf(fmt.Sprintln(sh...))
	}
	fmt.Println("\n====")
	fmt.Println(sh...)
	fmt.Print("====\n\n")

	// Gui_show("\n====")
	Gui_show(Itos(sh), "0gift")
	// Gui_show("====\n")
}

// Msg-房间封禁信息
func (replyF) room_block_msg(s string) {
	if uname := p.Json().GetValFromS(s, "uname"); uname == nil {
		msglog.L(`E: `, "uname", uname)
		return
	} else {
		Gui_show(Itos([]interface{}{"用户", uname, "已被封禁"}), "0room")
		fmt.Println("用户", uname, "已被封禁")
		msglog.Base_add("封").Log_show_control(false).L(`I: `, "用户", uname, "已被封禁")
	}
}

// Msg-房间准备信息，通常出现在下播而不出现在开播
func (replyF) preparing(s string) {
	msglog := msglog.Base_add("房", "下播")

	var type_item ws_msg.PREPARING
	if err := json.Unmarshal([]byte(s), &type_item); err != nil {
		msglog.L(`E: `, err)
		return
	} else {
		{ //附加功能 savestream结束
			c.C.Liveing = false
			// 停止此房间录制
			var roomId, _ = strconv.Atoi(type_item.Roomid)
			StreamOStop(roomId)
		}
		Gui_show("房间", type_item.Roomid, "下播了", "0room")
		msglog.L(`I: `, "房间", type_item.Roomid, "下播了")
	}
}

// Msg-房间开播信息
func (replyF) live(s string) {
	msglog := msglog.Base_add("房", "开播")

	var type_item ws_msg.LIVE
	if err := json.Unmarshal([]byte(s), &type_item); err != nil {
		msglog.L(`E: `, err)
		return
	} else {
		{ //附加功能 obs录播
			// Obsf(true)
			// Obs_R(true)
		}
		{
			c.C.Rev = 0.0                    //营收
			c.C.Liveing = true               //直播i标志
			c.C.Live_Start_Time = time.Now() //开播h时间
		}
		//开始录制
		go func() {
			if v, ok := c.C.K_v.LoadV(`仅保存当前直播间流`).(bool); ok && v {
				StreamOStop(-2) //停止其他房间录制
			}
			StreamOStart(c.C.Roomid)
			//有时不返回弹幕 开播刷新弹幕
			c.C.Danmu_Main_mq.Push_tag(`flash_room`, nil)
		}()

		Gui_show(Itos([]interface{}{"房间", type_item.Roomid, "开播了"}), "0room")
		msglog.L(`I: `, "房间", type_item.Roomid, "开播了")
	}
}

// Msg-超级留言处理
var sc_buf = make(map[int]struct{})

func (replyF) super_chat_message(s string) {
	msglog := msglog.Base_add("礼", "SC")

	var j ws_msg.SUPER_CHAT_MESSAGE
	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
	}

	id := j.Data.ID
	if id != 0 {
		if _, ok := sc_buf[id]; ok {
			return
		}
		if len(sc_buf) >= 10 {
			for k := range sc_buf {
				delete(sc_buf, k)
				break
			}
			{ //copy map
				tmp := make(map[int]struct{})
				for k, v := range sc_buf {
					tmp[k] = v
				}
				sc_buf = tmp
			}
		}
		sc_buf[id] = struct{}{}
	}
	uname := j.Data.UserInfo.Uname
	price := j.Data.Price
	message := j.Data.Message

	var sh = []interface{}{"SC: "}

	sh = append(sh, uname)
	logg := sh
	if price != 0 {
		sh = append(sh, "\n") //界面不显示价格
		logg = append(logg, "￥", price)
		c.C.Danmu_Main_mq.Push_tag(`c.Rev_add`, float64(price))
	}
	fmt.Println("====")
	fmt.Println(sh...)
	if message != "" {
		fmt.Println(message)
		sh = append(sh, message)
		logg = append(logg, message)
	}
	{ //语言tts
		c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{ //传入消息队列
			uid: `0superchat`,
			m: map[string]string{
				`{uname}`:   uname,
				`{price}`:   strconv.Itoa(price),
				`{message}`: message,
			},
		})
	}
	fmt.Print("====\n")

	{ //额外
		Assf(fmt.Sprintln(sh...))
		Gui_show(Itos(sh), "0superchat")
		//直播流服务弹幕
		SendStreamWs(Danmu_item{
			auth:   uname,
			border: true,
			color:  "#FF0000",
			msg:    "SC: " + message,
		})
	}
	msglog.Log_show_control(false).L(`I: `, logg...)
}

// Msg-热门榜获得v2
func (replyF) hot_rank_settlement_v2(s string) {
	msglog := msglog.Base_add("房")

	var type_item ws_msg.HOT_RANK_SETTLEMENT_V2
	if e := json.Unmarshal([]byte(s), &type_item); e != nil {
		msglog.L(`E: `, e)
	}
	var tmp = `获得:`
	if type_item.Data.AreaName != `` {
		tmp += type_item.Data.AreaName + " 第"
	}
	if type_item.Data.Rank != 0 {
		tmp += strconv.Itoa(type_item.Data.Rank)
	}
	Gui_show(tmp, "0rank")
	c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{ //传入消息队列
		uid: "0rank",
		m: map[string]string{
			`{Area_name}`: type_item.Data.AreaName,
			`{Rank}`:      strconv.Itoa(type_item.Data.Rank),
		},
	})
	msglog.L(`I: `, "热门榜", tmp)
}

// Msg-老板打赏新礼物红包
func (replyF) popularity_red_pocket_new(s string) {
	msglog := msglog.Base_add("礼", "礼物红包")

	var type_item ws_msg.POPULARITY_RED_POCKET_NEW
	if e := json.Unmarshal([]byte(s), &type_item); e != nil {
		msglog.L(`E: `, e)
	}
	var tmp = type_item.Data.Uname + type_item.Data.Action + strconv.Itoa(type_item.Data.Num) + `个` + type_item.Data.GiftName
	Gui_show(tmp, "0gift")
	c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{ //传入消息队列
		uid: "0gift",
		m: map[string]string{
			`{num}`:      strconv.Itoa(type_item.Data.Num),
			`{uname}`:    type_item.Data.Uname,
			`{action}`:   type_item.Data.Action,
			`{giftName}`: type_item.Data.GiftName,
		},
	})
	msglog.L(`I: `, tmp)
}

// Msg-老板打赏礼物红包
func (replyF) popularity_red_pocket_start(s string) {
	msglog := msglog.Base_add("礼", "礼物红包")

	var type_item ws_msg.POPULARITY_RED_POCKET_START
	if e := json.Unmarshal([]byte(s), &type_item); e != nil {
		msglog.L(`E: `, e)
	}
	var tmp = type_item.Data.SenderName + `送出了礼物红包`
	Gui_show(tmp, "0room")
	c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{ //传入消息队列
		uid: "0room",
		m: map[string]string{
			`{msg}`: tmp,
		},
	})
	msglog.L(`I: `, tmp)
}

// Msg-元气赏连抽
func (replyF) common_notice_danmaku(s string) {
	msglog := msglog.Base_add("房")

	var type_item ws_msg.COMMON_NOTICE_DANMAKU
	if e := json.Unmarshal([]byte(s), &type_item); e != nil {
		msglog.L(`E: `, e)
	}
	var tmp = type_item.Data.ContentSegments
	if len(tmp) == 0 {
		return
	}

	Gui_show(tmp[0].Text, "0room")
	c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{ //传入消息队列
		uid: "0room",
		m: map[string]string{
			`{msg}`: tmp[0].Text,
		},
	})
	msglog.L(`I: `, "元气赏连抽", tmp)
}

// Msg-小消息
func (replyF) little_message_box(s string) {
	msglog := msglog.Base_add("系统")

	var type_item ws_msg.LITTLE_MESSAGE_BOX
	if e := json.Unmarshal([]byte(s), &type_item); e != nil {
		msglog.L(`E: `, e)
	}
	if type_item.Data.Msg != `` {
		msglog.L(`I: `, type_item.Data.Msg)
	}
}

// Msg-粉丝牌切换
func (replyF) messagebox_user_medal_change(s string) {
	msglog := msglog.Base_add("房")

	var type_item ws_msg.MESSAGEBOX_USER_MEDAL_CHANGE
	if e := json.Unmarshal([]byte(s), &type_item); e != nil {
		msglog.L(`E: `, e)
	}
	if type_item.Data.Medal_name != `` {
		msglog.L(`I: `, "粉丝牌切换至", type_item.Data.Medal_name, type_item.Data.Medal_level)
	}
}

// Msg-进入特效
func (replyF) entry_effect(s string) {

	var res struct {
		Data struct {
			Copy_writing string `json:"copy_writing"`
		} `json:"data"`
	}
	if e := json.Unmarshal([]byte(s), &res); e != nil {
		msglog.L(`E: `, e)
	}

	var username string
	op := strings.Index(res.Data.Copy_writing, ` <%`)
	ed := strings.Index(res.Data.Copy_writing, `%> `)
	if op != -1 && ed != -1 {
		username = res.Data.Copy_writing[op+3 : ed]
	}
	//处理特殊字符
	copy_writing := strings.ReplaceAll(res.Data.Copy_writing, `<%`, ``)
	copy_writing = strings.ReplaceAll(copy_writing, `%>`, ``)

	guard_name := ""
	img := "0default"
	if strings.Contains(copy_writing, `总督`) {
		guard_name = `总督`
		img = "0level1"
	} else if strings.Contains(copy_writing, `提督`) {
		guard_name = `提督`
		img = "0level2"
	} else if strings.Contains(copy_writing, `舰长`) {
		guard_name = `舰长`
		img = "0level3"
	}

	{ //语言tts
		c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{ //传入消息队列
			uid: img,
			m: map[string]string{
				`{guard_name}`: guard_name,
				`{username}`:   username,
				`{msg}`:        copy_writing,
			},
		})
	}
	fmt.Print(">>> ")
	fmt.Println(copy_writing)
	Gui_show(copy_writing, img)

	msglog.Base_add("房").Log_show_control(false).L(`I: `, copy_writing)
}

// Msg-房间禁言
func (replyF) roomsilent(s string) {
	msglog := msglog.Base_add("房")

	if level := p.Json().GetValFromS(s, "data.level"); level == nil {
		msglog.L(`E: `, "level", level)
		return
	} else {
		if level.(float64) == 0 {
			msglog.L(`I: `, "主播关闭了禁言")
		}
		msglog.L(`I: `, "主播开启了等级禁言:", level)
	}
}

// Msg-弹幕处理
type Danmu_item struct {
	msg    string
	color  string
	border bool
	mode   int
	auth   interface{}
	uid    string
	roomid int //to avoid danmu show when room has changed
}

func (replyF) danmu(s string) {
	var j struct {
		Cmd  string        `json:"cmd"`
		Info []interface{} `json:"info"`
	}

	if e := json.Unmarshal([]byte(s), &j); e != nil {
		msglog.L(`E: `, e)
	}

	infob := j.Info
	item := Danmu_item{}
	{
		//解析
		if i, ok := infob[0].([]interface{}); ok {
			item.color = "#" + fmt.Sprintf("%x", F.Itob32(int32(i[3].(float64)))[1:])
		}
		if len(infob) > 0 {
			item.msg, _ = infob[1].(string)
			item.msg = strings.TrimSpace(item.msg)
		}
		if len(infob) > 1 {
			i, _ := infob[2].([]interface{})
			if len(i) > 0 {
				item.uid = strconv.Itoa(int(i[0].(float64)))
			}
			if len(i) > 1 {
				item.auth = i[1]
			}
		}
		item.roomid = c.C.Roomid
	}

	danmulog := msglog.Base("弹").LShow(false)

	if v, ok := c.C.K_v.LoadV(`弹幕输出到日志`).(bool); !ok || !v {
		danmulog.LFile("")
		danmulog.LDB(nil, "")
	}

	{ // 附加功能 弹幕机 封禁 弹幕合并
		// 保存弹幕至db
		saveDanmuToDB.init(c.C)
		saveDanmuToDB.danmu(item)
		// 对指定弹幕重新录制
		danmuReLiveTriger.Init(c.C)
		danmuReLiveTriger.Check(item.uid, item.msg)
		// 语言tts 私信
		{
			if item.uid != "" {
				if item.auth != nil {
					if s, ok := TTS_setting_string[item.uid]; ok && s != "" {
						c.C.Danmu_Main_mq.Push_tag(`tts`, Danmu_mq_t{ //传入消息队列
							uid: item.uid,
							m: map[string]string{
								`{auth}`: fmt.Sprint(item.auth),
								`{msg}`:  item.msg,
							},
						})
					}
				}
				if i, e := strconv.Atoi(item.uid); e == nil {
					if msg := c.C.K_v.LoadV(`弹幕私信`).(string); msg != "" {
						c.C.Danmu_Main_mq.Push_tag(`pm`, send.Pm_item{
							Uid: i,
							Msg: msg,
						}) //弹幕私信
					}
				}
				if c.C.K_v.LoadV(`额外私信对象`).(float64) != 0 {
					if msg, uid := c.C.K_v.LoadV(`弹幕私信(额外)`).(string), c.C.K_v.LoadV(`额外私信对象`).(float64); uid != 0 && msg != "" {
						c.C.Danmu_Main_mq.Push_tag(`pm`, send.Pm_item{
							Uid: int(uid),
							Msg: msg,
						}) //弹幕私信-对额外
					}
				}
			}
		}
		// 反射弹幕机
		if IsOn("反射弹幕机") {
			go Danmujif(item.msg)
		}
		if i := Autoskipf(item.msg); i > 0 {
			danmulog.L(`I: `, item.auth, ":", item.msg)
			return
		}
		//附加功能 更少弹幕
		if !Lessdanmuf(item.msg) {
			danmulog.L(`I: `, item.auth, ":", item.msg)
			return
		}
		if _msg := Shortdanmuf(item.msg); _msg == "" {
			danmulog.L(`I: `, item.auth, ":", item.msg)
			return
		} else {
			item.msg = _msg
		}
	}
	if item.auth != nil {
		danmulog.L(`I: `, item.auth, ":", item.msg)
	}
	Msg_showdanmu(item)
}

// 弹幕发送
// 传入字符串即可发送
// 需要cookie
func Msg_senddanmu(msg string) {
	if missKey := F.CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	}); len(missKey) != 0 || c.C.Roomid == 0 {
		msglog.L(`E: `, `c.Roomid == 0 || Cookie无Key:`, missKey)
		return
	}
	send.Danmu_s(msg, c.C.Roomid)
}

// 弹幕显示
// 由于额外功能有些需要显示，为了统一管理，使用此方法进行处理
func Msg_showdanmu(item Danmu_item) {
	//room change
	if item.roomid != 0 && item.roomid != c.C.Roomid {
		return
	}
	//展示
	{
		//ass
		Assf(item.msg)
		//直播流服务弹幕
		SendStreamWs(item)

		if item.auth != nil {
			Gui_show(fmt.Sprint(item.auth)+`: `+item.msg, item.uid)
		} else {
			Gui_show(item.msg, item.uid)
		}
	}
	fmt.Println(item.msg)
}

type Danmu_mq_t struct {
	uid string
	msg string
	m   map[string]string //tts参数替换列表
}

var Danmu_mq = mq.New()

// 消息显示
func Gui_show(m ...string) {
	//m[0]:msg m[1]:uid
	uid := ""
	if len(m) > 1 {
		uid = m[1]
	}
	Danmu_mq.Push_tag(`danmu`, Danmu_mq_t{
		uid: uid,
		msg: m[0],
	})
}

func Itos(i []interface{}) string {
	r := ""
	for _, v := range i {
		switch v := v.(type) {
		case string:
			r += v
		case int:
			r += strconv.Itoa(v)
		case int64:
			r += strconv.Itoa(int(v))
		case float64:
			r += strconv.Itoa(int(v))
		default:
			fmt.Println("unkonw type", v)
		}
		r += " "
	}
	return r
}
