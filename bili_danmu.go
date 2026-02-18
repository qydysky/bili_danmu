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
	"syscall"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
	reply "github.com/qydysky/bili_danmu/Reply"
	replyFunc "github.com/qydysky/bili_danmu/Reply/F"
	"github.com/qydysky/bili_danmu/Reply/F/danmuReLiveTriger"
	"github.com/qydysky/bili_danmu/Reply/F/recStartEnd"
	send "github.com/qydysky/bili_danmu/Send"
	pctx "github.com/qydysky/part/ctx"
	fc "github.com/qydysky/part/funcCtrl"
	plog "github.com/qydysky/part/log/v2"
	sys "github.com/qydysky/part/sys"

	msgq "github.com/qydysky/part/msgq"
	ws "github.com/qydysky/part/websocket"
)

func Start(rootCtx context.Context) {
	danmulog := c.C.Log.Base(`bilidanmu`)
	danmulog.I(`当前PID:`, c.C.PID)
	danmulog.I("version: ", c.C.Version)

	//检查配置
	if c.C.K_v.Len() == 0 {
		panic("未能加载配置")
	}

	// 保持唤醒
	var stop = sys.Sys().PreventSleep()
	defer stop.Done()

	mainCtx, mainDone := pctx.WithWait(context.Background(), 0, time.Minute)
	defer func() {
		danmulog.I(fmt.Sprintf("等待%v协程结束", time.Minute))
		if e := mainDone(); errors.Is(e, pctx.ErrWaitTo) {
			danmulog.W(`等待退出超时`)
		} else {
			danmulog.I("结束")
		}
	}()

	// 用户中断
	var cancelInterrupt, interrupt_chan = c.C.Danmu_Main_mq.Pull_tag_chan(`interrupt`, 2, mainCtx)
	defer cancelInterrupt()

	//ctrl+c退出
	go func() {
		var interrupt = make(chan os.Signal, 2)
		//捕获ctrl+c、容器退出
		signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
		danmulog.I("3s内2次ctrl+c退出")
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
		// 校验必要组件
		_ = replyFunc.ParseM3u8.Err()
		//命令行操作 切换房间 发送弹幕
		Cmd.Run2(func(ci CmdI) {
			ci.Cmd()
		})
		// 附加功能 savetojson
		replyFunc.SaveToJson.Run2(func(i replyFunc.SaveToJsonI) {
			i.Init(c.C.K_v.LoadV(`save_to_json`))
		})
		//ass初始化
		replyFunc.Ass.Run2(func(ai replyFunc.AssI) {
			ai.Init(c.C.K_v.LoadV("Ass"))
		})
		//tts初始化
		replyFunc.TTS.Run2(func(t replyFunc.TTSI) {
			t.Init(mainCtx, danmulog, c.C.K_v.LoadV("TTS"))
		})
		if reply.IsOn(`相似弹幕忽略`) {
			if max_num, ok := c.C.K_v.LoadV(`每秒显示弹幕数`).(float64); ok && int(max_num) >= 1 {
				replyFunc.LessDanmu.Run2(func(ldi replyFunc.LessDanmuI) {
					ldi.Init(int(max_num))
				})
			}
		}
		//rev初始化
		if c.C.IsOn(`统计营收`) {
			replyFunc.Rev.Run2(func(ri replyFunc.RevI) {
				ri.Init(danmulog)
			})
		}
		// 指定房间录制区间
		if _, err := recStartEnd.InitF.Run(mainCtx, c.C); err != nil {
			danmulog.Base("功能", "指定房间录制区间").E(err)
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
			danmulog.Base("功能", "指定弹幕重启录制").E(err)
		}
		// pgo gen
		if file, ok := c.C.K_v.LoadV("生成pgo").(string); ok {
			replyFunc.GenCpuPprof.Run2(func(inter interface {
				Start(ctx context.Context, file string) (any, error)
			}) {
				if _, e := inter.Start(mainCtx, file); e != nil {
					danmulog.Base("功能", "生成pgo").E(e)
				}
			})
		}
		// login
		c.C.Danmu_Main_mq.Pull_tag_only(`login`, func(_ any) (disable bool) {
			F.Api.Get(c.C, `CookieNoBlock`)
			return false
		})

		var (
			exitSign                     = false
			rs       fc.RangeSource[any] = func(yield func(any) bool) {
				for !exitSign {
					if !yield(nil) {
						break
					}
				}
			}
		)

		for ctx := range rs.RangeCtx(mainCtx) {
			if c.C.Roomid == 0 {
				if Cmd.Err() == nil {
					fmt.Println("回车查看指令")
				}
				cancel1, ch := c.C.Danmu_Main_mq.Pull_tag_chan(`change_room`, 1, ctx)
				select {
				case roomid := <-ch:
					c.C.Roomid = roomid.(int)
				case <-interrupt_chan:
					exitSign = true
				}
				cancel1()
			} else {
				danmulog.T("房间号: ", strconv.Itoa(c.C.Roomid))
			}

			if exitSign {
				break
			}

			danmulog.T("准备")

			//如果连接中断，则等待
			if !F.IsConnected() {
				cancel1, ch := c.C.Danmu_Main_mq.Pull_tag_chan(`exit_room`, 1, rootCtx)
				select {
				case <-ch:
					reply.StreamOStop(c.C.Roomid)
					danmulog.I("退出房间", c.C.Roomid)
					c.C.Roomid = 0
				case <-interrupt_chan:
					exitSign = true
				case <-time.After(time.Duration(30) * time.Second):
				}
				cancel1()
				continue
			}

			//获取cookie
			F.Api.Get(c.C, `Cookie`)
			//获取LIVE_BUVID
			// F.Api.Get(c.C, `LIVE_BUVID`)
			//兑换硬币
			F.Api.Get(c.C, `Silver2Coin`)
			// 获取房间实际id
			F.Api.Get(c.C, `Roomid`)

			//使用带tag的消息队列在功能间传递消息
			var cancelfunc = c.C.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
				`c.Rev_add`: func(data any) bool { //收入
					if rev, ok := data.(struct {
						Roomid int
						Rev    float64
					}); ok {
						common, ok := c.Commons.LoadV(c.C.Roomid).(*c.Common)
						if ok {
							common.Rev += rev.Rev
							// 显示营收
							replyFunc.Rev.Run2(func(ri replyFunc.RevI) {
								ri.ShowRev(common.Roomid, common.Rev)
							})
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
							danmulog.BaseAdd(`私信`).E(e)
						}
					}
					return false
				},
				`new day`: func(_ any) bool { //日期更换
					go func() {
						//每日兑换硬币
						F.Api.Get(c.C, `Silver2Coin`)
						//附加功能 每日发送弹幕
						reply.EntryDanmu(c.C)
						//附加功能 自动发送即将过期礼物
						reply.AutoSend_silver_gift(c.C)
					}()
					return false
				},
			})

			common, _ := c.CommonsLoadOrInit.LoadOrInitPThen(c.C.Roomid)(func(actual *c.Common, loaded bool) (*c.Common, bool) {
				if loaded {
					actual.InIdle = false
					actual.Rev = 0.0 // 营收
				} else {
					actual.UpUid = c.C.UpUid
					actual.Uname = c.C.Uname
					actual.ParentAreaID = c.C.ParentAreaID
					actual.AreaID = c.C.AreaID
					actual.Title = c.C.Title
					actual.Live_Start_Time = c.C.Live_Start_Time
					actual.Liveing = c.C.Liveing
					actual.Roomid = c.C.Roomid
				}
				return actual, loaded
			})

			exitSign = entryRoom(rootCtx, ctx, danmulog.BaseAdd(common.Roomid), common)

			cancelfunc()
			common.InIdle = true
			time.Sleep(time.Second)
		}

		{ //附加功能 直播流停止 ws信息保存
			replyFunc.SaveToJson.Run2(func(i replyFunc.SaveToJsonI) {
				i.Close()
			})
			reply.StreamOStopAll()
		}
	}
}

func entryRoom(rootCtx, mainCtx context.Context, danmulog *plog.Log, common *c.Common) (exitSign bool) {
	var (
		heartbeatmsg, heartinterval                        = F.Heartbeat()
		loopCtx, loopCancel                                = context.WithTimeout(mainCtx, time.Hour*3)
		rangeSource                 fc.RangeSource[string] = func(yield func(string) bool) {
			for !pctx.Done(loopCtx) {
				//如果连接中断，则等待
				if !F.IsConnected() {
					cancel1, ch := c.C.Danmu_Main_mq.Pull_tag_chan(`exit_room`, 1, mainCtx)
					select {
					case <-ch:
						reply.StreamOStop(c.C.Roomid)
						danmulog.I("退出房间", c.C.Roomid)
						c.C.Roomid = 0
					case <-mainCtx.Done():
					case <-time.After(time.Duration(30) * time.Second):
					}
					cancel1()
					continue
				}
				//获取cookie，检查是否登录失效
				F.Api.Get(common, `Cookie`)
				//获取LIVE_BUVID
				// F.Api.Get(common, `LIVE_BUVID`)
				//附加功能 自动发送即将过期礼物
				reply.AutoSend_silver_gift(common)
				//获取热门榜
				F.Api.Get(common, `Note`)
				// 检查与切换粉丝牌，只在cookie存在时启用
				F.Api.Get(common, `CheckSwitch_FansMedal`)
				// 附加功能 保持牌子点亮
				if reply.IsOn(`保持牌子亮着`) && common.Wearing_FansMedal != 0 {
					replyFunc.KeepMedalLight.Run2(func(kmli replyFunc.KeepMedalLightI) {
						kmli.Init(danmulog.Base("保持牌子点亮"), common.Roomid, send.Danmu_s, []any{}) // 不传入默认发送弹幕，使用点赞
					})
				} else {
					replyFunc.KeepMedalLight.Run2(func(kmli replyFunc.KeepMedalLightI) {
						kmli.Clear()
					})
				}
				if reply.IsOn(`相似弹幕忽略`) {
					replyFunc.LessDanmu.Run2(func(ldi replyFunc.LessDanmuI) {
						ldi.InitRoom(common.Roomid)
					})
				} else {
					replyFunc.LessDanmu.Run2(func(ldi replyFunc.LessDanmuI) {
						ldi.Unset()
					})
				}
				//tts
				defer func() {
					replyFunc.TTS.Run2(func(t replyFunc.TTSI) {
						t.Clear()
					})
				}()
				danmulog.I("连接到房间", common.Roomid)
				// 获取弹幕服务器
				F.Api.Get(common, `WSURL`)

				for {
					unlock := common.Lock()
					if len(common.WSURL) == 0 {
						unlock()
						break
					}
					wsUrl := common.WSURL[0]
					common.WSURL = common.WSURL[1:]
					unlock()
					if !yield(wsUrl) {
						return
					}
				}
			}
		}
	)

	for ctx, v := range rangeSource.RangeCtxCancel(loopCtx, loopCancel) {
		//ws启动
		danmulog.T("连接 " + v)
		u, _ := url.Parse(v)
		ws_c, err := ws.New_client(&ws.Client{
			// BufSize:           150,
			Url:               v,
			RTOMs:             (heartinterval + 5) * 1000,
			WTOMs:             (heartinterval + 5) * 1000,
			Proxy:             common.Proxy,
			Func_abort_close:  func() { danmulog.I(`服务器连接中断`) },
			Func_normal_close: func() { danmulog.I(`服务器连接关闭`) },
			Header: map[string]string{
				`Cookie`:          common.GenReqCookie(),
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
			danmulog.E("初始化连接错误", err)
			continue
		}
		wsmsg, err := ws_c.Handle()
		if err != nil {
			danmulog.E("尝试连接错误", err)
			continue
		}
		if ws_c.Isclose() {
			if err := ws_c.Error(); err != nil {
				danmulog.E("连接时错误", err)
			}
			continue
		}

		// auth
		{
			wsmsg.PushLock_tag(`send`, &ws.WsMsg{
				Msg: func(f func([]byte) error) error {
					return f(F.HelloGen(common.Roomid, common.Token))
				},
			})
			waitCheckAuth, cancel := context.WithTimeout(ctx, 5*time.Second)
			doneAuth := wsmsg.Pull_tag_only(`recv`, func(wm *ws.WsMsg) (disable bool) {
				_ = wm.Msg(func(b []byte) error {
					if F.HelloChe(b) {
						cancel()
					}
					return nil
				})
				return true
			})
			<-waitCheckAuth.Done()
			doneAuth()
			if err := waitCheckAuth.Err(); errors.Is(err, context.DeadlineExceeded) {
				danmulog.E("连接验证失败")
				continue
			}
		}

		danmulog.I("已连接到房间", common.Uname, `(`, common.Roomid, `)`)
		reply.Gui_show(`进入直播间: `+common.Uname+` (`+strconv.Itoa(common.Roomid)+`)`, `0room`)
		if common.Title != `` {
			danmulog.I(`房间标题: ` + common.Title)
			reply.Gui_show(`房间标题: `+common.Title, `0room`)
		}

		// 直播状态
		if F.Api.Get(common, `Liveing`); common.Liveing {
			danmulog.I("直播中")
		} else {
			danmulog.I("未直播")
		}

		// 处理ws消息
		var cancelDeal = wsmsg.Pull_tag(map[string]func(*ws.WsMsg) (disable bool){
			`recv`: func(wm *ws.WsMsg) (disable bool) {
				go func() {
					_ = wm.Msg(func(b []byte) error {
						reply.Reply(common, b)
						return nil
					})
				}()
				return false
			},
			`close`: func(_ *ws.WsMsg) (disable bool) {
				return true
			},
		})

		//30s获取一次心跳人气
		go func() {
			danmulog.T("获取心跳人气")

			replayTO := -1.0
			if tmp, _ := c.C.K_v.LoadV("消息响应超时s").(float64); tmp >= 300 || tmp == -1 {
				replayTO = tmp
			}

			for !ws_c.Isclose() {
				heartBeatSendT := time.Now()
				wsmsg.Push_tag(`send`, &ws.WsMsg{
					Msg: func(f func([]byte) error) error {
						return f(heartbeatmsg)
					},
				})
				time.Sleep(time.Millisecond * time.Duration(heartinterval*1000))
				if lastReplyT := func() time.Time {
					defer common.Lock()()
					return common.RepleyT
				}(); lastReplyT.IsZero() {
					danmulog.WF("心跳无响应")
				} else if to := lastReplyT.Sub(heartBeatSendT).Seconds(); replayTO > 0 && to > replayTO {
					danmulog.WF("心跳响应超时(%v)，重新进入房间", to)
					common.Danmu_Main_mq.Push_tag(`flash_room`, nil)
				}
			}
		}()

		// 刷新舰长数
		F.Api.Get(common, `GuardNum`)
		// 在线人数
		F.Api.Get(common, `getOnlineGoldRank`)
		//附加功能 弹幕机 无cookie无法发送弹幕
		if common.IsLogin() && reply.IsOn("自动弹幕机") {
			replyFunc.Danmuji.Run2(func(di replyFunc.DanmujiI) {
				di.Danmuji_auto(ctx, c.C.K_v.LoadV(`自动弹幕机_内容`).([]any), c.C.K_v.LoadV(`自动弹幕机_发送间隔s`).(float64), reply.Msg_senddanmu)
			})
		}
		{ //附加功能 进房间发送弹幕 直播流保存 每日签到
			F.RoomEntryAction(common.Roomid)
			reply.EntryDanmu(common)
			// go F.Dosign()
			if _, e := recStartEnd.RecStartCheck.Run(ctx, common); e == nil {
				reply.StreamOStart(common.Roomid)
			} else {
				danmulog.Base("功能", "指定房间录制区间").I(common.Roomid, e)
			}
			//弹幕合并
			if reply.IsOn("弹幕合并") {
				replyFunc.DanmuMerge.Run2(func(dmi replyFunc.DanmuMergeI) {
					dmi.Init(ctx, common.Roomid)
				})
			} else {
				replyFunc.DanmuMerge.Run2(func(dmi replyFunc.DanmuMergeI) {
					dmi.Unset()
				})
			}
		}

		//当前ws
		{
			var login = common.Login
			// 处理各种指令
			var cancelfunc = common.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
				`interrupt`: func(_ any) (disable bool) {
					reply.StreamOStopAll() //停止录制
					exitSign = true
					danmulog.I("停止，等待服务器断开连接")
					loopCancel()
					ws_c.Close()
					return true
				},
				`exit_room`: func(_ any) bool { //退出当前房间
					reply.StreamOStop(common.Roomid)
					c.C.Roomid = 0
					danmulog.I("退出房间", common.Roomid)
					loopCancel()
					ws_c.Close()
					return true
				},
				`change_room`: func(roomid any) bool { //换房时退出当前房间
					c.C.Roomid = roomid.(int)
					if v, ok := common.K_v.LoadV(`仅保存当前直播间流`).(bool); ok && v {
						reply.StreamOStopOther(c.C.Roomid) //停止其他房间录制
					}
					loopCancel()
					ws_c.Close()
					return true
				},
				`flash_room`: func(_ any) bool { //重进房时退出当前房间
					F.Api.Get(common, `WSURL`)
					ws_c.Close()
					return true
				},
				`guard_update`: func(_ any) bool { //舰长更新
					go F.Api.Get(common, `GuardNum`)
					return false
				},
				`every100s`: func(_ any) bool { //每100s
					if login != c.C.IsLogin() {
						common.Danmu_Main_mq.Push_tag(`flash_room`, nil)
						return false
					}
					if v, ok := common.K_v.LoadV("保持牌子亮着-开播时也发送").(bool); !common.Liveing || (ok && v) {
						replyFunc.KeepMedalLight.Run2(func(kmli replyFunc.KeepMedalLightI) {
							kmli.Do()
						})
					}
					if v, ok := common.K_v.LoadV("下播后不记录人气观看人数").(bool); ok && v && !common.Liveing {
						return false
					}
					// 在线人数
					go F.Api.Get(common, `getOnlineGoldRank`)
					return false
				},
			})

			danmulog.T("启动完成", common.Uname, `(`, common.Roomid, `)`)

			{
				cancel, c := wsmsg.Pull_tag_chan(`exit`, 1, ctx)
				select {
				case <-ctx.Done():
					common.Danmu_Main_mq.Push_tag(`flash_room`, nil)
				case <-c:
				case <-rootCtx.Done():
					common.Danmu_Main_mq.Push_tag(`interrupt`, nil)
					<-c
				}
				cancel()
			}

			if err := ws_c.Error(); err != nil {
				danmulog.E("结束连接错误", err)
			}

			cancelfunc()
			time.Sleep(time.Second)
		}
		cancelDeal()
	}

	return
}
