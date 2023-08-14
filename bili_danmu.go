package bili_danmu

import (
	"context"
	_ "embed"
	"errors"
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

func Start() {
	danmulog := c.C.Log.Base(`bilidanmu`)
	danmulog.L(`I: `, `当前PID:`, c.C.PID)
	danmulog.L(`I: `, "version: ", strings.TrimSpace(version))

	//检查配置
	if c.C.K_v.Len() == 0 {
		panic("未能加载配置")
	}

	// 保持唤醒
	var stop = sys.Sys().PreventSleep()
	defer stop.Done()

	// 用户中断
	var cancelInterrupt, interrupt_chan = c.C.Danmu_Main_mq.Pull_tag_chan(`interrupt`, 2, context.Background())
	defer cancelInterrupt()

	//ctrl+c退出
	go func() {
		var interrupt = make(chan os.Signal, 2)
		//捕获ctrl+c退出
		signal.Notify(interrupt, os.Interrupt)
		danmulog.L(`I: `, "3s内2次ctrl+c退出")
		for {
			<-interrupt
			c.C.Danmu_Main_mq.Push_tag(`interrupt`, nil)
			select {
			case <-interrupt:
				c.C.Danmu_Main_mq.Push_tag(`interrupt`, nil)
				os.Exit(1)
			case <-time.After(time.Second * 3):
			}
		}
	}()

	// 启动时显示ip
	if v, ok := c.C.K_v.LoadV("启动时显示ip").(bool); ok && v {
		for _, v := range sys.GetIntranetIp(``) {
			danmulog.L(`I: `, `当前ip：http://`+v)
		}
	}

	{
		//如果连接中断，则等待
		F.KeepConnect()
		//获取cookie
		F.Get(c.C).Get(`Cookie`)
		//获取LIVE_BUVID
		F.Get(c.C).Get(`LIVE_BUVID`)
		//命令行操作 切换房间 发送弹幕
		go Cmd.Cmd()
		//获取uid
		F.Get(c.C).Get(`Uid`)
		//兑换硬币
		F.Get(c.C).Silver_2_coin()
		//每日签到
		F.Dosign()
		// 附加功能 savetojson
		reply.SaveToJson.Init()

		//使用带tag的消息队列在功能间传递消息
		{
			var cancelfunc = c.C.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
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
					c.C.Danmu_Main_mq.PushLock_tag(`interrupt`, nil)
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
			defer cancelfunc()
		}

		for exit_sign := true; exit_sign; {
			if c.C.Roomid == 0 {
				fmt.Println("回车查看指令")
				ctx, cancel := context.WithCancel(context.Background())
				cancel1, c := c.C.Danmu_Main_mq.Pull_tag_chan(`change_room`, 1, ctx)
				select {
				case <-c:
				case <-interrupt_chan:
					exit_sign = false
				}
				cancel1()
				cancel()
			} else {
				fmt.Print("房间号: ", strconv.Itoa(c.C.Roomid), "\n")
			}

			if !exit_sign {
				break
			}

			danmulog.L(`T: `, "准备")
			//如果连接中断，则等待
			F.KeepConnect()
			//附加功能 保持牌子点亮
			go reply.Keep_medal_light()
			//附加功能 自动发送即将过期礼物
			go reply.AutoSend_silver_gift()
			//获取热门榜
			F.Get(c.C).Get(`Note`)

			danmulog.L(`I: `, "连接到房间", c.C.Roomid)

			Cookie := make(map[string]string)
			c.C.Cookie.Range(func(k, v any) bool {
				Cookie[k.(string)] = v.(string)
				return true
			})

			F.Get(c.C).Get(`Liveing`)
			// 检查与切换粉丝牌，只在cookie存在时启用
			F.Get(c.C).Get(`CheckSwitch_FansMedal`)

			// 直播状态
			if c.C.Liveing {
				danmulog.L(`I: `, "直播中")
			} else {
				danmulog.L(`I: `, "未直播")
			}

			// 对每个弹幕服务器尝试
			F.Get(c.C).Get(`WSURL`)
			aliveT := time.Now().Add(3 * time.Hour)
			for i, exitloop := 0, false; !exitloop && i < len(c.C.WSURL) && time.Now().Before(aliveT); {
				v := c.C.WSURL[i]
				//ws启动
				danmulog.L(`T: `, "连接 "+v)
				u, _ := url.Parse(v)
				ws_c, err := ws.New_client(&ws.Client{
					Url:               v,
					TO:                35 * 1000,
					Proxy:             c.C.Proxy,
					Func_abort_close:  func() { danmulog.L(`I: `, `服务器连接中断`) },
					Func_normal_close: func() { danmulog.L(`I: `, `服务器连接关闭`) },
					Header: map[string]string{
						`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
						`Host`:            u.Hostname(),
						`User-Agent`:      c.UA,
						`Accept`:          `*/*`,
						`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
						`Origin`:          `https://live.bilibili.com`,
						`Pragma`:          `no-cache`,
						`Cache-Control`:   `no-cache`,
					},
				})
				if err != nil {
					danmulog.L(`E: `, "连接错误", err)
					i += 1
					continue
				}
				wsmsg, err := ws_c.Handle()
				if err != nil {
					danmulog.L(`E: `, "连接错误", err)
					i += 1
					continue
				}
				if ws_c.Isclose() {
					danmulog.L(`E: `, "连接错误", ws_c.Error())
					i += 1
					continue
				}

				// auth
				{
					wsmsg.PushLock_tag(`send`, &ws.WsMsg{
						Msg: F.HelloGen(c.C.Roomid, c.C.Token),
					})
					waitCheckAuth, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					doneAuth := wsmsg.Pull_tag_only(`rec`, func(wm *ws.WsMsg) (disable bool) {
						if F.HelloChe(wm.Msg) {
							cancel()
						}
						return true
					})
					<-waitCheckAuth.Done()
					doneAuth()
					if err := waitCheckAuth.Err(); errors.Is(err, context.DeadlineExceeded) {
						danmulog.L(`E: `, "连接验证失败")
						i += 1
						continue
					}
				}

				danmulog.L(`I: `, "已连接到房间", c.C.Uname, `(`, c.C.Roomid, `)`)
				reply.Gui_show(`进入直播间: `+c.C.Uname+` (`+strconv.Itoa(c.C.Roomid)+`)`, `0room`)
				if c.C.Title != `` {
					danmulog.L(`I: `, `房间标题: `+c.C.Title)
					reply.Gui_show(`房间标题: `+c.C.Title, `0room`)
				}

				// 处理ws消息
				var cancelDeal = wsmsg.Pull_tag(map[string]func(*ws.WsMsg) (disable bool){
					`rec`: func(wm *ws.WsMsg) (disable bool) {
						go reply.Reply(wm.Msg)
						return false
					},
					`close`: func(_ *ws.WsMsg) (disable bool) {
						return true
					},
				})

				//30s获取一次人气
				go func() {
					danmulog.L(`T: `, "获取人气")
					heartbeatmsg, heartinterval := F.Heartbeat()
					for !ws_c.Isclose() {
						wsmsg.Push_tag(`send`, &ws.WsMsg{
							Msg: heartbeatmsg,
						})
						time.Sleep(time.Millisecond * time.Duration(heartinterval*1000))
					}
				}()

				// 刷新舰长数
				F.Get(c.C).Get(`GuardNum`)
				// 在线人数
				F.Get(c.C).Get(`getOnlineGoldRank`)
				//验证cookie
				if missKey := F.CookieCheck([]string{
					`bili_jct`,
					`DedeUserID`,
					`LIVE_BUVID`,
				}); len(missKey) == 0 {
					//附加功能 弹幕机 无cookie无法发送弹幕
					reply.Danmuji_auto()
				}
				{ //附加功能 进房间发送弹幕 直播流保存 每日签到
					go F.Dosign()
					go reply.Entry_danmu()
					go reply.StreamOStart(c.C.Roomid)
					go F.RoomEntryAction(c.C.Roomid)
				}

				//当前ws
				{
					// 处理各种指令
					var cancelfunc = c.C.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
						`interrupt`: func(_ any) (disable bool) {
							exitloop = true
							exit_sign = false
							ws_c.Close()
							danmulog.L(`I: `, "停止，等待服务器断开连接")
							reply.StreamOStop(-1) //停止录制
							return true
						},
						`exit_room`: func(_ any) bool { //退出当前房间
							exitloop = true
							reply.StreamOStop(c.C.Roomid)
							danmulog.L(`I: `, "退出房间", c.C.Roomid)
							c.C.Roomid = 0
							ws_c.Close()
							return true
						},
						`change_room`: func(_ any) bool { //换房时退出当前房间
							exitloop = true
							ws_c.Close()
							if v, ok := c.C.K_v.LoadV(`仅保存当前直播间流`).(bool); ok && v {
								reply.StreamOStop(-2) //停止其他房间录制
							}
							return true
						},
						`flash_room`: func(_ any) bool { //重进房时退出当前房间
							go F.Get(c.C).Get(`WSURL`)
							ws_c.Close()
							return true
						},
						`guard_update`: func(_ any) bool { //舰长更新
							go F.Get(c.C).Get(`GuardNum`)
							return false
						},
						`every100s`: func(_ any) bool { //每100s
							if time.Now().After(aliveT) {
								c.C.Danmu_Main_mq.Push_tag(`flash_room`, nil)
								return false
							}
							if v, ok := c.C.K_v.LoadV("下播后不记录人气观看人数").(bool); ok && v && !c.C.Liveing {
								return false
							}
							// 在线人数
							go F.Get(c.C).Get(`getOnlineGoldRank`)
							return false
						},
						`new day`: func(_ any) bool { //日期更换
							go func() {
								//每日签到
								F.Dosign()
								//每日兑换硬币
								F.Get(c.C).Silver_2_coin()
								//附加功能 每日发送弹幕
								reply.Entry_danmu()
								//附加功能 保持牌子点亮
								reply.Keep_medal_light()
								//附加功能 自动发送即将过期礼物
								reply.AutoSend_silver_gift()
							}()
							return false
						},
					})

					{
						cancel, c := wsmsg.Pull_tag_chan(`exit`, 1, context.Background())
						<-c
						cancel()
					}

					cancelfunc()
					time.Sleep(time.Second)
				}
				cancelDeal()
			}
			time.Sleep(time.Second)
		}

		{ //附加功能 直播流停止 ws信息保存
			reply.SaveToJson.Close()
			reply.StreamOStop(-1)
		}
		danmulog.L(`I: `, "结束退出")
	}
}
