package bili_danmu

import (
	"fmt"
	"flag"
	"time"
	"net/url"
	"strconv"
	"os"
	"os/signal"
		
	reply "github.com/qydysky/bili_danmu/Reply"
	send "github.com/qydysky/bili_danmu/Send"
	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"

	p "github.com/qydysky/part"
	ws "github.com/qydysky/part/websocket"
	msgq "github.com/qydysky/part/msgq"
	reqf "github.com/qydysky/part/reqf"
)

func init() {
	go func(){//日期变化
		var old = time.Now().Hour()
		for {
			if now := time.Now().Hour();now == 0 && old != now {
				c.Danmu_Main_mq.Push_tag(`new day`,nil)
				old = now
			}
			time.Sleep(time.Second*time.Duration(100))
		}
	}()
}

func Demo(roomid ...int) {
	var danmulog = c.Log.Base(`bilidanmu Demo`)
	defer danmulog.Block(1000)

	//ctrl+c退出
	interrupt := make(chan os.Signal,2)
	go func(){
		danmulog.L(`T: `, "两次ctrl+c强制退出")
		for len(interrupt) < 2 {
			time.Sleep(time.Second*3)
		}
		danmulog.L(`I: `, "强制退出!").Block(1000)
		os.Exit(1)
	}()

	{
		var groomid = flag.Int("r", 0, "roomid")
		flag.Parse()

		var change_room_chan = make(chan struct{})

		go func(){
			var room = *groomid
			if room == 0 && len(roomid) != 0 {
				room = roomid[0]
			}
			if room == 0 {
				c.Log.Block(1000)//等待所有日志输出完毕
				fmt.Printf("输入房间号: ")
				_, err := fmt.Scanln(&room)
				if err != nil {
					danmulog.L(`E: `, "输入错误", err)
					return
				}
			} else {
				fmt.Print("房间号: ", strconv.Itoa(room), "\n")
			}
			if c.Roomid == 0 {
				c.Roomid = room
				change_room_chan <- struct{}{}
			}
		}()
		
		//使用带tag的消息队列在功能间传递消息
		c.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
			`change_room`:func(data interface{})(bool){//房间改变
				c.Rev = 0.0 //营收
				c.Renqi = 1//人气置1
				c.GuardNum = 0//舰长数
				c.Note = ``//分区排行
				c.Uname = ``//主播id
				c.Title = ``
				c.Wearing_FansMedal = 0
				for len(change_room_chan) != 0 {<-change_room_chan}
				change_room_chan <- struct{}{}
				return false
			},
			`c.Rev_add`:func(data interface{})(bool){//收入
				c.Rev += data.(float64)
				return false
			},
			`c.Renqi`:func(data interface{})(bool){//人气更新
				if tmp,ok := data.(int);ok{
					c.Renqi = tmp
				}
				return false
			},
			`gtk_close`:func(data interface{})(bool){//gtk关闭信号
				interrupt <- os.Interrupt
				return false
			},
		})
		//单独，避免队列执行耗时block从而无法接收更多消息
		c.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
			`pm`:func(data interface{})(bool){//私信
				if tmp,ok := data.(send.Pm_item);ok{
					send.Send_pm(tmp.Uid,tmp.Msg)
				}
				return false
			},
		})

		<-change_room_chan

		//捕获ctrl+c退出
		signal.Notify(interrupt, os.Interrupt)
		//如果连接中断，则等待
		F.KeepConnect()
		//获取cookie
		F.Get(`Cookie`)
		//获取uid
		F.Get(`Uid`)
		//命令行操作 切换房间 发送弹幕
		go F.Cmd()
		//兑换硬币
		F.Get(`Silver_2_coin`)
		//每日签到
		F.Dosign()
		// //客户版本 不再需要
		// F.Get(`VERSION`)
		//附加功能 保持牌子点亮
		go reply.Keep_medal_light()
		//附加功能 自动发送即将过期礼物
		go reply.AutoSend_silver_gift()

		for exit_sign:=true;exit_sign; {

			danmulog.L(`T: `,"准备")
			//如果连接中断，则等待
			F.KeepConnect()
			//获取热门榜
			F.Get(`Note`)

			danmulog.L(`I: `,"连接到房间", c.Roomid)

			Cookie := make(map[string]string)
			c.Cookie.Range(func(k,v interface{})(bool){
				Cookie[k.(string)] = v.(string)
				return true
			})

			F.Get(`Liveing`)
			//检查与切换粉丝牌，只在cookie存在时启用
			F.Get(`CheckSwitch_FansMedal`)

			//直播状态
			if c.Liveing {
				danmulog.L(`I: `,"直播中")
			} else {
				danmulog.L(`I: `,"未直播")
			}

			//对每个弹幕服务器尝试
			F.Get(`WSURL`)
			for _, v := range c.WSURL {
				//ws启动
				u, _ := url.Parse(v)
				ws_c := ws.New_client(ws.Client{
					Url:v,
					TO:35 * 1000,
					Proxy:c.Proxy,
					Func_abort_close:func(){danmulog.L(`I: `,`服务器连接中断`)},
					Func_normal_close:func(){danmulog.L(`I: `,`服务器连接关闭`)},
					Header: map[string]string{
						`Cookie`:reqf.Map_2_Cookies_String(Cookie),
						`Host`: u.Hostname(),
						`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0`,
						`Accept`: `*/*`,
						`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
						`Origin`: `https://live.bilibili.com`,
						`Pragma`: `no-cache`,
						`Cache-Control`: `no-cache`,
					},
				}).Handle()
				if ws_c.Isclose() {
					danmulog.L(`E: `,"连接错误", ws_c.Error())
					continue
				}

				//SendChan 传入发送[]byte
				//RecvChan 接收[]byte
				danmulog.L(`T: `,"连接", v)
				ws_c.SendChan <- F.HelloGen(c.Roomid, c.Token)
				if F.HelloChe(<- ws_c.RecvChan) {
					danmulog.L(`I: `,"已连接到房间", c.Uname, `(`, c.Roomid, `)`)
					reply.Gui_show(`进入直播间: `+c.Uname+` (`+strconv.Itoa(c.Roomid)+`)`, `0room`)
					if c.Title != `` {
						danmulog.L(`I: `,c.Title)
						reply.Gui_show(`房间标题: `+c.Title, `0room`)
					}
					//30s获取一次人气
					go func(){
						p.Sys().MTimeoutf(500)//500ms
						danmulog.L(`T: `,"获取人气")
						go func(){
							heartbeatmsg, heartinterval := F.Heartbeat()
							for !ws_c.Isclose() {
								ws_c.SendChan <- heartbeatmsg
								time.Sleep(time.Millisecond*time.Duration(heartinterval*1000))
							}
						}()

						//订阅消息，以便刷新舰长数
						F.Get(`GuardNum`)
						//使用带tag的消息队列在功能间传递消息
						c.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
							`guard_update`:func(data interface{})(bool){//舰长更新
								go F.Get(`GuardNum`)
								return false
							},
							`change_room`:func(data interface{})(bool){//换房时退出当前房间
								return true
							},
							`new day`:func(data interface{})(bool){//日期更换
								//每日签到
								F.Dosign()
								// //获取用户版本  不再需要
								// go F.Get(`VERSION`)
								//每日兑换硬币
								go F.Silver_2_coin()
								//小心心
								go F.F_x25Kn()
								//附加功能 每日发送弹幕
								go reply.Entry_danmu()
								//附加功能 保持牌子点亮
								go reply.Keep_medal_light()
								//附加功能 自动发送即将过期礼物
								go reply.AutoSend_silver_gift()
								return false
							},
						})

						//验证cookie
						if missKey := F.CookieCheck([]string{
							`bili_jct`,
							`DedeUserID`,
							`LIVE_BUVID`,
						});len(missKey) == 0 {
							//附加功能 弹幕机 无cookie无法发送弹幕
							reply.Danmuji_auto()
						}

						{//附加功能 进房间发送弹幕 直播流保存 营收
							go reply.Entry_danmu()
							go reply.Savestreamf()
							go reply.ShowRevf()
							//小心心
							go F.F_x25Kn()
						}
					}()
				}

				var isclose bool
				var break_sign bool
				for !isclose {
					select {
					case i := <- ws_c.RecvChan:
						if len(i) == 0 && ws_c.Isclose() {
							isclose = true
						} else {
							go reply.Reply(i)
						}
					case <- interrupt:
						ws_c.Close()
						danmulog.L(`I: `,"停止，等待服务器断开连接")
						break_sign = true
						exit_sign = false
					case <- change_room_chan:
						ws_c.Close()
						danmulog.L(`I: `,"停止，等待服务器断开连接")
						break_sign = true
					}
				}

				if break_sign {break}
			}
			{//附加功能 直播流停止
				reply.Savestream_wait()
				reply.Save_to_json(-1, []interface{}{`{}]`})
			}
			p.Sys().Timeoutf(1)
		}

		close(interrupt)
		danmulog.L(`I: `,"结束退出")
	}
}

