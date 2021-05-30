package F

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"time"
	"fmt"

	send "github.com/qydysky/bili_danmu/Send"
	c "github.com/qydysky/bili_danmu/CV"
)

//直播间缓存
var liveList =make(map[string]int)

func Cmd() {
	
	cmdlog := c.Log.Base_add(`命令行操作`).L(`T: `,`回车查看帮助`)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if inputs := scanner.Text();inputs == `` {//帮助
			fmt.Print("\n")
			fmt.Println("切换房间->输入数字回车")
			if _,ok := c.Cookie.LoadV(`bili_jct`).(string);ok {
				fmt.Println("发送弹幕->输入' 字符串'回车")
				fmt.Println("查看直播中主播->输入' live'回车")
				fmt.Println("获取小心心->输入' getheart'回车")
			} else {
				fmt.Println("登陆->输入' login'回车")
			}
			fmt.Println("搜索主播->输入' search关键词'回车")
			fmt.Println("房间信息->输入' room'回车")
			fmt.Println("开始结束录制->输入' rec'回车")
			fmt.Println("其他输出隔断不影响")
			fmt.Print("\n")
		} else if inputs[0] == 27 {//屏蔽功能键
			cmdlog.L(`W: `, "不支持功能键")
		} else if inputs[0] == 32 {// 开头
			//录制切换
			if strings.Contains(inputs, ` rec`) {
				if !c.Liveing {
					cmdlog.L(`W: `, "不能切换录制状态，未在直播")
					continue
				}
				c.Danmu_Main_mq.Push_tag(`savestream`, nil)
				continue
			}
			//直播间切换
			if strings.Contains(inputs, ` live`) {
				if _,ok := c.Cookie.LoadV(`bili_jct`).(string);!ok {
					cmdlog.L(`W: `, "尚未登陆，未能获取关注主播")
					continue
				}
				fmt.Print("\n")
				if len(inputs) > 5 {
					if room,ok := liveList[inputs];ok{
						c.Roomid = room
						c.Danmu_Main_mq.Push_tag(`change_room`,nil)
						continue
					}
					cmdlog.L(`W: `, "输入错误", inputs)
					continue
				}
				for k,v := range Feed_list() {
					liveList[` live`+strconv.Itoa(k)] = v.Roomid
					fmt.Printf("%d\t%s\n\t\t\t%s\n", k, v.Uname, v.Title)
				}
				fmt.Println("回复' live(序号)'进入直播间")
				fmt.Print("\n")
				continue
			}
			//登陆
			if strings.Contains(inputs, ` login`) {
				if _,ok := c.Cookie.LoadV(`bili_jct`).(string);ok {
					cmdlog.L(`W: `, "已登陆")
					continue
				}
				//获取cookie
				Get(`Cookie`)

				continue
			}
			//获取小心心
			if strings.Contains(inputs, ` getheart`) {
				if _,ok := c.Cookie.LoadV(`bili_jct`).(string);!ok {
					cmdlog.L(`W: `, "尚未登陆，不能获取小心心")
					continue
				}
				//获取小心心
				go F_x25Kn()

				continue
			}
			//搜索主播
			if strings.Contains(inputs, ` search`) {
				if len(inputs) == 7 {
					cmdlog.L(`W: `, "未输入搜索内容")
					continue
				}

				fmt.Print("\n")
				for k,v := range SearchUP(inputs[7:]) {
					liveList[` live`+strconv.Itoa(k)] = v.Roomid
					if v.Is_live {
						fmt.Printf("%d\t%s\t%s\n", k, `☁`, v.Uname)
					} else {
						fmt.Printf("%d\t%s\t%s\n", k, ` `, v.Uname)
					}
				}
				fmt.Println("回复' live(序号)'进入直播间")
				fmt.Print("\n")

				continue
			}
			//当前直播间信息
			if strings.Contains(inputs, ` room`) {
				fmt.Print("\n")
				fmt.Println("当前直播间信息")
				{
					living := `未在直播`
					if c.Liveing {living = `直播中`}
					fmt.Println(c.Uname, c.Title, living)
				}
				{
					if c.Liveing {
						d := time.Since(c.Live_Start_Time).Round(time.Second)
						h := d / time.Hour
						d -= h * time.Hour
						m := d / time.Minute
						d -= m * time.Minute
						s := d / time.Second
						fmt.Println(`已直播时长:`, fmt.Sprintf("%02d:%02d:%02d", h, m, s))
					}
				}
				{
					fmt.Println(`营收:`, fmt.Sprintf("￥%.2f",c.Rev))
				}
				fmt.Println(`舰长数:`, c.GuardNum)
				fmt.Println(`分区排行:`, c.Note, `人气：`, c.Renqi)
				if c.Stream_url != ""{fmt.Println(`直播Web服务:`, c.Stream_url)}
				fmt.Print("\n")

				continue
			}
			{//弹幕发送
				if _,ok := c.Cookie.LoadV(`bili_jct`).(string);!ok {
					cmdlog.L(`W: `, "尚未登陆，不能发送弹幕")
					continue
				}
				if len(inputs) < 2 {
					cmdlog.L(`W: `, "输入长度过短", inputs)
					continue
				}
				send.Danmu_s(inputs[1:], c.Roomid)
			}
		} else if room,err := strconv.Atoi(inputs);err == nil {//直接进入房间
			c.Roomid = room
			cmdlog.L(`I: `, "进入房间",room)
			c.Danmu_Main_mq.Push_tag(`change_room`,nil)
		} else {//其余字符串
			send.Danmu_s(inputs, c.Roomid)
		}
	}
}