package Cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	reply "github.com/qydysky/bili_danmu/Reply"
	send "github.com/qydysky/bili_danmu/Send"
)

// 直播间缓存
var liveList = make(map[string]int)

func Cmd() {

	cmdlog := c.C.Log.Base_add(`命令行操作`).L(`T: `, `回车查看帮助`)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if inputs := scanner.Text(); inputs == `` { //帮助
			fmt.Print("\n")
			fmt.Println("切换房间->输入数字回车")
			if c.C.Roomid == 0 {
				if _, ok := c.C.Cookie.LoadV(`bili_jct`).(string); ok {
					fmt.Println("查看直播中主播->输入' liv'回车")
					fmt.Println("查看历史观看主播->输入' his'回车")
				} else {
					fmt.Println("登陆->输入' login'回车")
				}
				fmt.Println("搜索主播->输入' sea关键词'回车")
				fmt.Println("其他输出隔断不影响")
				fmt.Print("\n")
				continue
			}
			if _, ok := c.C.Cookie.LoadV(`bili_jct`).(string); ok {
				fmt.Println("发送弹幕->输入' 字符串'回车")
				fmt.Println("查看直播中主播->输入' liv'回车")
			} else {
				fmt.Println("登陆->输入' login'回车")
			}
			fmt.Println("重载弹幕->输入' reload'回车")
			fmt.Println("搜索主播->输入' sea关键词'回车")
			fmt.Println("房间信息->输入' room'回车")
			fmt.Println("开始结束录制->输入' rec'回车")
			fmt.Println("其他输出隔断不影响")
			fmt.Print("\n")
		} else if inputs[0] == 27 { //屏蔽功能键
			cmdlog.L(`W: `, "不支持功能键")
		} else if inputs[0] == 32 { // 开头
			cmdlog.L(`T: `, "指令("+inputs+")")
			//录制切换
			if strings.Contains(inputs, ` rec`) {
				if len(inputs) > 4 {
					if v, ok := c.C.K_v.LoadV(`仅保存当前直播间流`).(bool); ok && v {
						cmdlog.L(`W: `, "输入错误", inputs)
						continue
					}
					if room, err := strconv.Atoi(inputs[4:]); err == nil {
						c.C.Danmu_Main_mq.Push_tag(`savestream`, reply.SavestreamO{
							Roomid: room,
							IsRec:  !reply.StreamOStatus(room),
						})
						continue
					}
					cmdlog.L(`W: `, "输入错误", inputs)
				} else {
					c.C.Danmu_Main_mq.Push_tag(`savestream`, reply.SavestreamO{
						Roomid: c.C.Roomid,
						IsRec:  !reply.StreamOStatus(c.C.Roomid),
					})
				}
				continue
			}
			//进入房间
			if strings.Contains(inputs, ` to`) {
				if len(inputs) == 3 {
					cmdlog.L(`W: `, "未输入进入序号")
					continue
				}

				fmt.Print("\n")
				if room, ok := liveList[inputs]; ok {
					c.C.Roomid = room
					c.C.Danmu_Main_mq.Push_tag(`change_room`, nil)
					continue
				} else {
					cmdlog.L(`W: `, "输入错误", inputs)
				}
				continue
			}
			//直播间切换
			if strings.Contains(inputs, ` liv`) {
				if _, ok := c.C.Cookie.LoadV(`bili_jct`).(string); !ok {
					cmdlog.L(`W: `, "尚未登陆，未能获取关注主播")
					continue
				}
				fmt.Print("\n")
				for k, v := range F.Feed_list() {
					liveList[` to`+strconv.Itoa(k)] = v.Roomid
					fmt.Printf("%d\t%s(%d)\n\t\t\t%s\n", k, v.Uname, v.Roomid, v.Title)
				}
				fmt.Println("回复' to(序号)'进入直播间")
				fmt.Print("\n")
				continue
			}
			//直播间历史
			if strings.Contains(inputs, ` his`) {
				if _, ok := c.C.Cookie.LoadV(`bili_jct`).(string); !ok {
					cmdlog.L(`W: `, "尚未登陆，未能获取关注主播")
					continue
				}
				fmt.Print("\n")
				for k, v := range F.GetHisStream() {
					liveList[` to`+strconv.Itoa(k)] = v.Roomid
					if v.LiveStatus == 1 {
						fmt.Printf("%d\t%s\t%s(%d)\n\t\t\t%s\n", k, `☁`, v.Uname, v.Roomid, v.Title)
					} else {
						fmt.Printf("%d\t%s\t%s(%d)\n\t\t\t%s\n", k, ` `, v.Uname, v.Roomid, v.Title)
					}
				}
				fmt.Println("回复' to(序号)'进入直播间")
				fmt.Print("\n")
				continue
			}
			//登陆
			if strings.Contains(inputs, ` login`) {
				//获取cookie
				F.Get(&c.C).Get(`Cookie`)
				continue
			}
			//搜索主播
			if strings.Contains(inputs, ` sea`) {
				if len(inputs) == 4 {
					cmdlog.L(`W: `, "未输入搜索内容")
					continue
				}

				fmt.Print("\n")
				for k, v := range F.SearchUP(inputs[4:]) {
					liveList[` to`+strconv.Itoa(k)] = v.Roomid
					if v.Is_live {
						fmt.Printf("%d\t%s\t%s(%d)\n", k, `☁`, v.Uname, v.Roomid)
					} else {
						fmt.Printf("%d\t%s\t%s(%d)\n", k, ` `, v.Uname, v.Roomid)
					}
				}
				fmt.Println("回复' to(序号)'进入直播间")
				fmt.Print("\n")

				continue
			}
			//重载弹幕
			if strings.Contains(inputs, ` reload`) && c.C.Roomid != 0 {
				c.C.Danmu_Main_mq.Push_tag(`flash_room`, nil)
				continue
			}
			//当前直播间信息
			if strings.Contains(inputs, ` room`) && c.C.Roomid != 0 {
				fmt.Print("\n")
				fmt.Println("当前直播间(" + strconv.Itoa(c.C.Roomid) + ")信息")
				{
					living := `未在直播`
					if c.C.Liveing {
						living = `直播中`
					}
					fmt.Println(c.C.Uname, c.C.Title, living)
				}
				{
					if c.C.Liveing {
						d := time.Since(c.C.Live_Start_Time).Round(time.Second)
						h := d / time.Hour
						d -= h * time.Hour
						m := d / time.Minute
						d -= m * time.Minute
						s := d / time.Second
						fmt.Println(`已直播时长:`, fmt.Sprintf("%02d:%02d:%02d", h, m, s))
					}
				}
				{
					fmt.Println(`营收:`, fmt.Sprintf("￥%.2f", c.C.Rev))
				}
				fmt.Println(`舰长数:`, c.C.GuardNum)
				fmt.Println(`分区排行:`, c.C.Note, `人气：`, c.C.Renqi, `观看人数：`, c.C.Watched, `在线人数：`, c.C.OnlineNum)
				for _, v := range c.C.Stream_url {
					fmt.Println(`直播Web服务:`, v)
				}
				var array = reply.StreamOCommon(-1)
				fmt.Println(`正在录制的房间：`)
				for _, v := range array {
					fmt.Println("\t" + v.Uname + "(" + strconv.Itoa(v.Roomid) + ") " + v.Title)
				}
				fmt.Print("输入` rec` 来启停当前房间录制")

				if v, ok := c.C.K_v.LoadV(`仅保存当前直播间流`).(bool); !ok || !v {
					fmt.Print(" 输入` rec房间号` 来启停其他录制")
				}

				fmt.Print("\n\n")

				continue
			}
			//直接进入房间
			if room, err := strconv.Atoi(inputs[1:]); err == nil {
				c.C.Roomid = room
				cmdlog.L(`I: `, "进入房间", room)
				c.C.Danmu_Main_mq.Push_tag(`change_room`, nil)
			}
			cmdlog.L(`W: `, "无效指令("+inputs+")")
		} else { //其余字符串
			if c.C.Roomid == 0 {
				continue
			}
			send.Danmu_s(inputs, c.C.Roomid)
		}
	}
}
