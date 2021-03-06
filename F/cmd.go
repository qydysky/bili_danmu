package F

import (
	"os"
	"bufio"
	"strings"
	"strconv"

	p "github.com/qydysky/part"
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
			cmdlog.L(`I: `, "切换房间->输入数字回车")
			cmdlog.L(`I: `, "发送弹幕->输入' 字符串'回车")
			cmdlog.L(`I: `, "查看直播中主播->输入' live'回车")
			cmdlog.L(`I: `, "其他输出隔断不影响")
		} else if inputs[0] == 27 {//屏蔽功能键
			cmdlog.L(`W: `, "不支持功能键")
		} else if inputs[0] == 32 {// 开头
			if strings.Contains(inputs, ` live`) {//直播间切换
				if len(inputs) > 5 {
					if room,ok := liveList[inputs];ok{
						c.Roomid = room
						cmdlog.L(`I: `, "进入房间",room)
						c.Danmu_Main_mq.Push_tag(`change_room`,nil)
						continue
					}
					cmdlog.L(`W: `, "输入错误", inputs)
					continue
				}
				for k,v := range Feed_list() {
					liveList[` live`+strconv.Itoa(k)] = v.Roomid
					cmdlog.L(`I: `, k, v.Uname, v.Title)
				}
				cmdlog.L(`I: `, "回复' live(序号)'进入直播间")
				continue
			}
			{//弹幕发送
				if len(inputs) < 2 {
					cmdlog.L(`W: `, "输入长度过短", inputs)
					continue
				}
				Cookie := make(map[string]string)
				c.Cookie.Range(func(k,v interface{})(bool){
					Cookie[k.(string)] = v.(string)
					return true
				})
				send.Danmu_s(inputs[1:], p.Map_2_Cookies_String(Cookie),c.Roomid)
			}
		} else if room,err := strconv.Atoi(inputs);err == nil {//直接进入房间
			c.Roomid = room
			cmdlog.L(`I: `, "进入房间",room)
			c.Danmu_Main_mq.Push_tag(`change_room`,nil)
		} else {//其余字符串
			Cookie := make(map[string]string)
			c.Cookie.Range(func(k,v interface{})(bool){
				Cookie[k.(string)] = v.(string)
				return true
			})
			send.Danmu_s(inputs, p.Map_2_Cookies_String(Cookie),c.Roomid)
		}
	}
}