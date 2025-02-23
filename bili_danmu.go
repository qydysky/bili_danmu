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
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	reply "github.com/qydysky/bili_danmu/Reply"
	replyFunc "github.com/qydysky/bili_danmu/Reply/F"
	"github.com/qydysky/bili_danmu/Reply/F/danmuReLiveTriger"
	"github.com/qydysky/bili_danmu/Reply/F/genCpuPprof"
	"github.com/qydysky/bili_danmu/Reply/F/recStartEnd"
	send "github.com/qydysky/bili_danmu/Send"
	Cmd "github.com/qydysky/bili_danmu/cmd"
	pctx "github.com/qydysky/part/ctx"
	part "github.com/qydysky/part/log"
	sys "github.com/qydysky/part/sys"

	msgq "github.com/qydysky/part/msgq"
	reqf "github.com/qydysky/part/reqf"
	ws "github.com/qydysky/part/websocket"
)

func Start() {
	danmulog := c.C.Log.Base(`bilidanmu`)
	danmulog.L(`I: `, `当前PID:`, c.C.PID)
	danmulog.L(`I: `, "version: ", c.C.Version)

	//检查配置
	if c.C.K_v.Len() == 0 {
		panic("未能加载配置")
	}

	// 保持唤醒
	var stop = sys.Sys().PreventSleep()
	defer stop.Done()

	mainCtx, mainDone := pctx.WithWait(context.Background(), 0, time.Minute)
	defer func() {
		danmulog.L(`I: `, fmt.Sprintf("等待%v协程结束", time.Minute))
		if e := mainDone(); errors.Is(e, pctx.ErrWaitTo) {
			danmulog.L(`W: `, `等待退出超时`)
		} else {
			danmulog.L(`I: `, "结束")
		}
	}()

	// 用户中断
	var cancelInterrupt, interrupt_chan = c.C.Danmu_Main_mq.Pull_tag_chan(`interrupt`, 2, mainCtx)
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

	{
		//如果连接中断，则等待
		F.KeepConnect()
		//获取cookie
		F.Get(c.C).Get(`Cookie`)
		//获取LIVE_BUVID
		//获取uid
		F.Get(c.C).Get(`LIVE_BUVID`)
		//命令行操作 切换房间 发送弹幕
		go Cmd.Cmd()
		// F.Get(c.C).Get(`Uid`)
		//兑换硬币
		F.Get(c.C).Silver_2_coin()
		//每日签到
		// F.Dosign()
		// 附加功能 savetojson
		reply.SaveToJson.Init()
		// 附加功能 保持牌子点亮
		reply.KeepMedalLight(mainCtx, c.C)
		//ass初始化
		replyFunc.Ass.Init(c.C.K_v.LoadV("Ass"))
		// 指定房间录制区间
		if _, err := recStartEnd.InitF.Run(mainCtx, c.C); err != nil {
			danmulog.Base("功能", "指定房间录制区间").L(`E: `, err)
		} else {
			_, _ = recStartEnd.LoopCheck.Run(mainCtx, recStartEnd.StreamCtl{
				C:       c.C,
				Commons: c.Commons,
				State:   reply.StreamOStatus,
				Start:   reply.StreamOStart,
				End:     reply.StreamOStop,
				Cut:     func(i int) { reply.StreamOCut(i) },
			})
		}
		// 指定弹幕重启录制
		if _, err := danmuReLiveTriger.Init.Run(mainCtx, danmuReLiveTriger.DanmuReLiveTriger{
			StreamCut: func(i int, title ...string) {
				reply.StreamOCut(i)(title...)
			},
			C: c.C,
		}); err != nil {
			danmulog.Base("功能", "指定弹幕重启录制").L(`E: `, err)
		}
		// pgo gen
		if file, ok := c.C.K_v.LoadV("生成pgo").(string); ok {
			if _, e := genCpuPprof.Start.Run(mainCtx, file); e != nil {
				danmulog.Base("功能", "生成pgo").L(`E: `, e)
			}
		}

		//使用带tag的消息队列在功能间传递消息
		{
			var cancelfunc = c.C.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
				// `change_room`: func(_ any) bool { //房间改变
				// 	c.C.Rev = 0.0     // 营收
				// 	c.C.Renqi = 1     // 人气置1
				// 	c.C.Watched = 0   // 观看人数
				// 	c.C.OnlineNum = 0 // 在线人数
				// 	c.C.GuardNum = 0  // 舰长数
				// 	c.C.Note = ``     // 分区排行
				// 	c.C.Uname = ``    // 主播id
				// 	c.C.Title = ``
				// 	c.C.Wearing_FansMedal = 0
				// 	return false
				// },
				`c.Rev_add`: func(data any) bool { //收入
					if rev, ok := data.(struct {
						Roomid int
						Rev    float64
					}); ok {
						common, ok := c.Commons.LoadV(c.C.Roomid).(*c.Common)
						if ok {
							common.Rev += rev.Rev
						}
					}
					return false
				},
				`c.Renqi`: func(data any) bool { //人气更新
					if tmp, ok := data.(struct {
						Roomid int
						Renqi  int
					}); ok {
						common, ok := c.Commons.LoadV(c.C.Roomid).(*c.Common)
						if ok {
							common.Renqi = tmp.Renqi
						}
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
				`new day`: func(_ any) bool { //日期更换
					go func() {
						//每日签到
						// F.Dosign()
						//每日兑换硬币
						F.Get(c.C).Silver_2_coin()
						//附加功能 每日发送弹幕
						reply.Entry_danmu(c.C)
						//附加功能 自动发送即将过期礼物
						reply.AutoSend_silver_gift(c.C)
					}()
					return false
				},
			})
			defer cancelfunc()
		}

		for exitSign := false; !exitSign; {
			if c.C.Roomid == 0 {
				fmt.Println("回车查看指令")
				ctx, cancel := context.WithCancel(mainCtx)
				cancel1, ch := c.C.Danmu_Main_mq.Pull_tag_chan(`change_room`, 1, ctx)
				select {
				case roomid := <-ch:
					c.C.Roomid = roomid.(int)
				case <-interrupt_chan:
					exitSign = true
				}
				cancel1()
				cancel()
			} else {
				danmulog.L(`T: `, "房间号: ", strconv.Itoa(c.C.Roomid))
			}

			if exitSign {
				break
			}

			danmulog.L(`T: `, "准备")

			//如果连接中断，则等待
			F.KeepConnect()
			//获取cookie，检查是否登陆失效
			F.Get(c.C).Get(`Cookie`)
			//获取LIVE_BUVID
			F.Get(c.C).Get(`LIVE_BUVID`)

			// 获取房间实际id
			c.C.Roomid = F.GetRoomRealId(c.C.Roomid)

			common, loaded := c.CommonsLoadOrStore.LoadOrStoreP(c.Commons, c.C.Roomid)

			if loaded {
				common.InIdle = false
				common.Rev = 0.0 // 营收
			}

			exitSign = entryRoom(mainCtx, danmulog.BaseAdd(common.Roomid), common)

			common.InIdle = true

			time.Sleep(time.Second)
		}

		{ //附加功能 直播流停止 ws信息保存
			reply.SaveToJson.Close()
			reply.StreamOStopAll()
		}
	}
}

func entryRoom(mainCtx context.Context, danmulog *part.Log_interface, common *c.Common) (exitSign bool) {
	//附加功能 自动发送即将过期礼物
	go reply.AutoSend_silver_gift(common)
	//获取热门榜
	F.Get(common).Get(`Note`)

	danmulog.L(`I: `, "连接到房间", common.Roomid)

	Cookie := make(map[string]string)
	common.Cookie.Range(func(k, v any) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	// 检查与切换粉丝牌，只在cookie存在时启用
	F.Get(common).Get(`CheckSwitch_FansMedal`)

	// 对每个弹幕服务器尝试
	F.Get(common).Get(`WSURL`)
	aliveT := time.Now().Add(3 * time.Hour)
	heartbeatmsg, heartinterval := F.Heartbeat()
	for i, exitloop := 0, false; !exitloop && i < len(common.WSURL) && time.Now().Before(aliveT); {
		v := common.WSURL[i]
		//ws启动
		danmulog.L(`T: `, "连接 "+v)
		u, _ := url.Parse(v)
		ws_c, err := ws.New_client(&ws.Client{
			Url:               v,
			TO:                (heartinterval + 5) * 1000,
			Proxy:             common.Proxy,
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
			if err := ws_c.Error(); err != nil {
				danmulog.L(`E: `, "连接错误", err)
			}
			i += 1
			continue
		}

		// auth
		{
			wsmsg.PushLock_tag(`send`, &ws.WsMsg{
				Msg: F.HelloGen(common.Roomid, common.Token),
			})
			waitCheckAuth, cancel := context.WithTimeout(mainCtx, 5*time.Second)
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

		danmulog.L(`I: `, "已连接到房间", common.Uname, `(`, common.Roomid, `)`)
		reply.Gui_show(`进入直播间: `+common.Uname+` (`+strconv.Itoa(common.Roomid)+`)`, `0room`)
		if common.Title != `` {
			danmulog.L(`I: `, `房间标题: `+common.Title)
			reply.Gui_show(`房间标题: `+common.Title, `0room`)
		}

		// 直播状态
		if F.Get(common).Get(`Liveing`); common.Liveing {
			danmulog.L(`I: `, "直播中")
		} else {
			danmulog.L(`I: `, "未直播")
		}

		// 处理ws消息
		var cancelDeal = wsmsg.Pull_tag(map[string]func(*ws.WsMsg) (disable bool){
			`rec`: func(wm *ws.WsMsg) (disable bool) {
				go reply.Reply(common, wm.Msg)
				return false
			},
			`close`: func(_ *ws.WsMsg) (disable bool) {
				return true
			},
		})

		//30s获取一次人气
		go func() {
			danmulog.L(`T: `, "获取人气")
			for !ws_c.Isclose() {
				wsmsg.Push_tag(`send`, &ws.WsMsg{
					Msg: heartbeatmsg,
				})
				time.Sleep(time.Millisecond * time.Duration(heartinterval*1000))
			}
		}()

		// 刷新舰长数
		F.Get(common).Get(`GuardNum`)
		// 在线人数
		F.Get(common).Get(`getOnlineGoldRank`)
		//验证cookie
		if missKey := F.CookieCheck([]string{
			`bili_jct`,
			`DedeUserID`,
			`LIVE_BUVID`,
		}); len(missKey) == 0 && reply.IsOn("自动弹幕机") {
			//附加功能 弹幕机 无cookie无法发送弹幕
			replyFunc.Danmuji.Danmuji_auto(mainCtx, c.C.K_v.LoadV(`自动弹幕机_内容`).([]any), c.C.K_v.LoadV(`自动弹幕机_发送间隔s`).(float64), reply.Msg_senddanmu)
		}
		{ //附加功能 进房间发送弹幕 直播流保存 每日签到
			// go F.Dosign()
			go reply.Entry_danmu(common)
			if _, e := recStartEnd.RecStartCheck.Run(mainCtx, common); e == nil {
				go reply.StreamOStart(common, common.Roomid)
			} else {
				danmulog.Base("功能", "指定房间录制区间").L(`I: `, common.Roomid, e)
			}
			go F.RoomEntryAction(common.Roomid)
		}

		//当前ws
		{
			// 处理各种指令
			var cancelfunc = common.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
				`interrupt`: func(_ any) (disable bool) {
					exitloop = true
					exitSign = true
					ws_c.Close()
					danmulog.L(`I: `, "停止，等待服务器断开连接")
					reply.StreamOStopAll() //停止录制
					return true
				},
				`exit_room`: func(_ any) bool { //退出当前房间
					exitloop = true
					reply.StreamOStop(common.Roomid)
					danmulog.L(`I: `, "退出房间", common.Roomid)
					c.C.Roomid = 0
					ws_c.Close()
					return true
				},
				`change_room`: func(roomid any) bool { //换房时退出当前房间
					c.C.Roomid = roomid.(int)
					exitloop = true
					ws_c.Close()
					if v, ok := common.K_v.LoadV(`仅保存当前直播间流`).(bool); ok && v {
						reply.StreamOStopOther(c.C.Roomid) //停止其他房间录制
					}
					return true
				},
				`flash_room`: func(_ any) bool { //重进房时退出当前房间
					go F.Get(common).Get(`WSURL`)
					ws_c.Close()
					return true
				},
				`guard_update`: func(_ any) bool { //舰长更新
					go F.Get(common).Get(`GuardNum`)
					return false
				},
				`every100s`: func(_ any) bool { //每100s
					if time.Now().After(aliveT) {
						common.Danmu_Main_mq.Push_tag(`flash_room`, nil)
						return false
					}
					if v, ok := common.K_v.LoadV("下播后不记录人气观看人数").(bool); ok && v && !common.Liveing {
						return false
					}
					// 在线人数
					go F.Get(common).Get(`getOnlineGoldRank`)
					return false
				},
			})

			{
				cancel, c := wsmsg.Pull_tag_chan(`exit`, 1, mainCtx)
				<-c
				cancel()
			}

			if err := ws_c.Error(); err != nil {
				danmulog.L(`E: `, "连接错误", err)
			}

			cancelfunc()
			time.Sleep(time.Second)
		}
		cancelDeal()
	}

	return
}
