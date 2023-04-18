package bili_danmu

import (
	_ "embed"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	reply "github.com/qydysky/bili_danmu/Reply"
	send "github.com/qydysky/bili_danmu/Send"
	Cmd "github.com/qydysky/bili_danmu/cmd"
	sys "github.com/qydysky/part/sys"

	msgq "github.com/qydysky/part/msgq"
	reqf "github.com/qydysky/part/reqf"
	ws "github.com/qydysky/part/websocket"
)

//go:embed VERSION
var version string

func init() {
	go func() { //日期变化
		var old = time.Now().Hour()
		for {
			if now := time.Now().Hour(); now == 0 && old != now {
				c.C.Danmu_Main_mq.Push_tag(`new day`, nil)
				old = now
			}
			c.C.Danmu_Main_mq.Push_tag(`every100s`, nil)
			time.Sleep(time.Second * time.Duration(100))
		}
	}()
}

func Start() {
	var danmulog = c.C.Log.Base(`bilidanmu`)
	defer danmulog.Block(1000)

	var stop = sys.Sys().PreventSleep()
	defer stop.Done()

	danmulog.L(`I: `, "version: ", strings.TrimSpace(version))

	//ctrl+c退出
	interrupt := make(chan os.Signal, 2)
	go func() {
		danmulog.L(`T: `, "两次ctrl+c强制退出")
		for len(interrupt) < 2 {
			time.Sleep(time.Second * 3)
		}
		danmulog.L(`I: `, "强制退出!").Block(1000)
		os.Exit(1)
	}()

	danmulog.L(`I: `, `当前PID:`, c.C.PID)

	//启动时显示ip
	{
		if v, ok := c.C.K_v.LoadV("启动时显示ip").(bool); ok && v {
			for _, v := range sys.GetIntranetIp(``) {
				danmulog.L(`I: `, `当前ip：http://`+v)
			}
		}
	}

	//检查配置
	{
		if c.C.K_v.Len() == 0 {
			panic("未能加载配置")
		}
	}

	{
		var (
			change_room_chan = make(chan struct{})
			flash_room_chan  = make(chan struct{})
		)

		//如果连接中断，则等待
		F.KeepConnect()
		//获取cookie
		F.Get(c.C).Get(`Cookie`)
		//获取LIVE_BUVID
		F.Get(c.C).Get(`LIVE_BUVID`)

		// 房间初始化
		if c.C.Roomid == 0 {
			c.C.Log.Block(1000) //等待所有日志输出完毕
			fmt.Println("回车查看指令")
		} else {
			fmt.Print("房间号: ", strconv.Itoa(c.C.Roomid), "\n")
			go func() { change_room_chan <- struct{}{} }()
		}

		//命令行操作 切换房间 发送弹幕
		go Cmd.Cmd()

		//使用带tag的消息队列在功能间传递消息
		c.C.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
			`flash_room`: func(_ any) bool { //房间重进
				F.Get(c.C).Get(`WSURL`)
				select {
				case flash_room_chan <- struct{}{}:
				default:
				}
				return false
			},
			`change_room`: func(_ any) bool { //房间改变
				c.C.Rev = 0.0     // 营收
				c.C.Renqi = 1     // 人气置1
				c.C.Watched = 0   // 观看人数
				c.C.OnlineNum = 0 // 在线人数
				c.C.GuardNum = 0  // 舰长数
				c.C.Note = ``     // 分区排行
				c.C.Uname = ``    // 主播id
				c.C.Title = ``
				c.C.Wearing_FansMedal = 0
				for len(change_room_chan) != 0 {
					<-change_room_chan
				}
				change_room_chan <- struct{}{}
				return false
			},
			`c.Rev_add`: func(data any) bool { //收入
				c.C.Rev += data.(float64)
				return false
			},
			`c.Renqi`: func(data any) bool { //人气更新
				if tmp, ok := data.(int); ok {
					c.C.Renqi = tmp
				}
				return false
			},
			`gtk_close`: func(_ any) bool { //gtk关闭信号
				interrupt <- os.Interrupt
				return false
			},
			`pm`: func(data any) bool { //私信
				if tmp, ok := data.(send.Pm_item); ok {
					if e := send.Send_pm(tmp.Uid, tmp.Msg); e != nil {
						danmulog.Base_add(`私信`).L(`E: `, e)
					}
				}
				return false
			},
		})

		<-change_room_chan

		//捕获ctrl+c退出
		signal.Notify(interrupt, os.Interrupt)
		//获取uid
		F.Get(c.C).Get(`Uid`)
		//兑换硬币
		F.Get(c.C).Get(`Silver_2_coin`)
		//每日签到
		F.Dosign()
		// //客户版本 不再需要
		// F.Get(`VERSION`)
		// 附加功能 savetojson
		reply.SaveToJson.Init()
		//附加功能 保持牌子点亮
		go reply.Keep_medal_light()
		//附加功能 自动发送即将过期礼物
		go reply.AutoSend_silver_gift()

		for exit_sign := true; exit_sign; {

			danmulog.L(`T: `, "准备")
			//如果连接中断，则等待
			F.KeepConnect()
			//获取热门榜
			F.Get(c.C).Get(`Note`)

			danmulog.L(`I: `, "连接到房间", c.C.Roomid)

			Cookie := make(map[string]string)
			c.C.Cookie.Range(func(k, v any) bool {
				Cookie[k.(string)] = v.(string)
				return true
			})

			F.Get(c.C).Get(`Liveing`)
			//检查与切换粉丝牌，只在cookie存在时启用
			F.Get(c.C).Get(`CheckSwitch_FansMedal`)

			//直播状态
			if c.C.Liveing {
				danmulog.L(`I: `, "直播中")
			} else {
				danmulog.L(`I: `, "未直播")
			}

			//对每个弹幕服务器尝试
			F.Get(c.C).Get(`WSURL`)
			for i := 0; i < len(c.C.WSURL); i += 1 {
				v := c.C.WSURL[i]
				//ws启动
				danmulog.L(`T: `, "连接 "+v)
				u, _ := url.Parse(v)
				ws_c := ws.New_client(ws.Client{
					Url:               v,
					TO:                35 * 1000,
					Proxy:             c.C.Proxy,
					Func_abort_close:  func() { danmulog.L(`I: `, `服务器连接中断`) },
					Func_normal_close: func() { danmulog.L(`I: `, `服务器连接关闭`) },
					Header: map[string]string{
						`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
						`Host`:            u.Hostname(),
						`User-Agent`:      `Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0`,
						`Accept`:          `*/*`,
						`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
						`Origin`:          `https://live.bilibili.com`,
						`Pragma`:          `no-cache`,
						`Cache-Control`:   `no-cache`,
					},
				}).Handle()
				if ws_c.Isclose() {
					danmulog.L(`E: `, "连接错误", ws_c.Error())
					continue
				}

				//SendChan 传入发送[]byte
				//RecvChan 接收[]byte
				ws_c.SendChan <- F.HelloGen(c.C.Roomid, c.C.Token)
				if F.HelloChe(<-ws_c.RecvChan) {
					danmulog.L(`I: `, "已连接到房间", c.C.Uname, `(`, c.C.Roomid, `)`)
					reply.Gui_show(`进入直播间: `+c.C.Uname+` (`+strconv.Itoa(c.C.Roomid)+`)`, `0room`)
					if c.C.Title != `` {
						danmulog.L(`I: `, c.C.Title)
						reply.Gui_show(`房间标题: `+c.C.Title, `0room`)
					}
					//30s获取一次人气
					go func() {
						sys.Sys().MTimeoutf(500) //500ms
						danmulog.L(`T: `, "获取人气")
						go func() {
							heartbeatmsg, heartinterval := F.Heartbeat()
							for !ws_c.Isclose() {
								ws_c.SendChan <- heartbeatmsg
								time.Sleep(time.Millisecond * time.Duration(heartinterval*1000))
							}
						}()

						//订阅消息，以便刷新舰长数
						F.Get(c.C).Get(`GuardNum`)
						// 在线人数
						F.Get(c.C).Get(`getOnlineGoldRank`)
						//当前ws的消息队列
						c.C.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
							`exit_cu_room`: func(_ any) bool { //退出
								return true
							},
							`guard_update`: func(_ any) bool { //舰长更新
								go F.Get(c.C).Get(`GuardNum`)
								return false
							},
							`flash_room`: func(_ any) bool { //重进房时退出当前房间
								return true
							},
							`change_room`: func(_ any) bool { //换房时退出当前房间
								return true
							},
							`every100s`: func(_ any) bool { //每100s
								if v, ok := c.C.K_v.LoadV("下播后不记录人气观看人数").(bool); ok && v && !c.C.Liveing {
									return false
								}
								// 在线人数
								F.Get(c.C).Get(`getOnlineGoldRank`)
								return false
							},
							`new day`: func(_ any) bool { //日期更换
								//每日签到
								F.Dosign()
								// //获取用户版本  不再需要
								// go F.Get(`VERSION`)
								//每日兑换硬币
								go F.Get(c.C).Silver_2_coin()
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
						}); len(missKey) == 0 {
							//附加功能 弹幕机 无cookie无法发送弹幕
							reply.Danmuji_auto()
						}

						{ //附加功能 进房间发送弹幕 直播流保存 营收 每日签到
							go F.Dosign()
							go reply.Entry_danmu()
							go reply.StreamOStart(c.C.Roomid)
							go reply.ShowRevf()
							go F.RoomEntryAction(c.C.Roomid)
						}
					}()
				}

				var isclose bool
				var break_sign bool
				for !isclose {
					select {
					case i := <-ws_c.RecvChan:
						if len(i) == 0 && ws_c.Isclose() {
							isclose = true
						} else {
							go reply.Reply(i)
						}
					case <-interrupt:
						ws_c.Close()
						danmulog.L(`I: `, "停止，等待服务器断开连接")
						break_sign = true
						exit_sign = false
					case <-flash_room_chan:
						ws_c.Close()
						danmulog.L(`I: `, "停止，等待服务器断开连接")
						//刷新WSURL
						F.Get(c.C).Get(`WSURL`)
						i = 0
					case <-change_room_chan:
						ws_c.Close()
						danmulog.L(`I: `, "停止，等待服务器断开连接")
						break_sign = true
					}
				}

				// 当前ws链接断开时，取消当前ws的消息队列
				c.C.Danmu_Main_mq.Push_tag(`exit_cu_room`, nil)

				if break_sign {
					break
				}
			}
			{ //附加功能
				if v, ok := c.C.K_v.LoadV(`仅保存当前直播间流`).(bool); ok && v {
					reply.StreamOStop(-2) //停止其他房间录制
				}
			}
			sys.Sys().Timeoutf(1)
		}

		{ //附加功能 直播流停止 ws信息保存
			reply.SaveToJson.Close()
			reply.StreamOStop(-1)
		}
		// close(interrupt)
		danmulog.L(`I: `, "结束退出")
	}
}
