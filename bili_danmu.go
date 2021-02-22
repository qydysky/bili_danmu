package bili_danmu

import (
	"fmt"
	"flag"
	"time"
	"net/url"
	"strconv"
	"os"
	"os/signal"

	p "github.com/qydysky/part"
	ws "github.com/qydysky/part/websocket"
	msgq "github.com/qydysky/part/msgq"
	reply "github.com/qydysky/bili_danmu/Reply"
	send "github.com/qydysky/bili_danmu/Send"
	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
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
	
	//ctrl+c退出
	interrupt := make(chan os.Signal, 1)
	
	{
		var groomid = flag.Int("r", 0, "roomid")
		var live_qn = flag.String("q", "0", "qn")
		flag.Parse()
	
		if _,ok := c.Default_qn[*live_qn]; ok{c.Live_qn = *live_qn}

		var exit_sign bool
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
				reply.Saveflv_wait()//停止保存直播流
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

		//ctrl+c退出
		signal.Notify(interrupt, os.Interrupt)

		for !exit_sign {
			//获取cookies
			{
				var q = p.Filel{
					Write:false,
				}
				var get_cookie = func(){
					danmulog.L(`I: `, "未检测到cookie.txt，如果需要登录请在本机打开以下网址扫码登录，不需要请忽略")
					//获取cookie
					F.Get_cookie()
					//验证cookie
					if missKey := F.CookieCheck([]string{
						`bili_jct`,
						`DedeUserID`,
					});len(missKey) == 0 {
						danmulog.L(`I: `,"你已登录，刷新房间！")
						//刷新
						c.Danmu_Main_mq.Push_tag(`change_room`,nil)
					}
				}
				
				if !p.Checkfile().IsExist("cookie.txt") {//读取cookie文件
					go get_cookie()
					p.Sys().Timeoutf(3)
				} else {
					q.File = "cookie.txt"
					cookieString := p.File().FileWR(q)

					if cookieString == `` {//cookie.txt为空
						danmulog.L(`T: `, `cookie.txt为空`)
						go get_cookie()
						p.Sys().Timeoutf(3)
					} else {
						for k,v := range p.Cookies_String_2_Map(cookieString){//cookie存入全局变量syncmap
							c.Cookie.Store(k, v)
						}
						if uid,ok := c.Cookie.LoadV(`DedeUserID`).(string);!ok{//cookie中无DedeUserID
							danmulog.L(`T: `, `读取cookie错误,无DedeUserID`)
							go get_cookie()
							p.Sys().Timeoutf(3)
						} else if uid,e := strconv.Atoi(uid);e != nil{
							danmulog.L(`E: `, e)
							go get_cookie()
							p.Sys().Timeoutf(3)
						} else {
							c.Uid = uid
						}
					}
				}
			}
			
			//获取房间相关信息
			api := F.New_api(c.Roomid).Get_host_Token().Get_live()
			c.Roomid = api.Roomid
			//每日签到
			F.Dosign()
			//每日兑换硬币
			F.Silver_2_coin()
			//获取用户版本
			api.Get_Version()
			//获取热门榜
			api.Get_HotRank()
			//检查与切换粉丝牌，只在cookie存在时启用
			api.CheckSwitch_FansMedal()
			//小心心
			go api.F_x25Kn()
			if len(api.Url) == 0 || api.Roomid == 0 || api.Token == "" || api.Uid == 0 || api.Locked {
				danmulog.L(`E: `,"some err")
				return
			}
			danmulog.L(`I: `,"连接到房间", c.Roomid)

			Cookie := make(map[string]string)
			c.Cookie.Range(func(k,v interface{})(bool){
				Cookie[k.(string)] = v.(string)
				return true
			})

			//对每个弹幕服务器尝试
			for _, v := range api.Url {
				//ws启动
				u, _ := url.Parse(v)
				ws_c := ws.New_client(ws.Client{
					Url:v,
					TO:35 * 1000,
					Func_abort_close:func(){danmulog.L(`I: `,`服务器连接中断`)},
					Func_normal_close:func(){danmulog.L(`I: `,`服务器连接关闭`)},
					Header: map[string]string{
						`Cookie`:p.Map_2_Cookies_String(Cookie),
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
				danmulog.L(`I: `,"连接", v)
				ws_c.SendChan <- F.HelloGen(c.Roomid, api.Token)
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
						danmulog.L(`I: `,"获取人气")
						go func(){
							heartbeatmsg, heartinterval := F.Heartbeat()
							for !ws_c.Isclose() {
								ws_c.SendChan <- heartbeatmsg
								time.Sleep(time.Millisecond*time.Duration(heartinterval*1000))
							}
						}()

						//传输变量，以便响应弹幕"弹幕机在么"
						c.Live = api.Live
						//获取过往营收 舰长数量
						// go api.Get_OnlineGoldRank()//高能榜显示的是在线观众的打赏

						//订阅消息，以便刷新舰长数
						api.Get_guardNum()
						//使用带tag的消息队列在功能间传递消息
						c.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
							`guard_update`:func(data interface{})(bool){//舰长更新
								go api.Get_guardNum()
								return false
							},
							`change_room`:func(data interface{})(bool){//换房时退出当前房间
								return true
							},
							`new day`:func(data interface{})(bool){//日期更换
								//每日签到
								F.Dosign()
								//每日兑换硬币
								go F.Silver_2_coin()
								//小心心
								go api.F_x25Kn()
								//附加功能 每日发送弹幕
								go reply.Entry_danmu()
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
							go reply.Saveflvf()
							go reply.ShowRevf()
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
						exit_sign = true
					case <- change_room_chan:
						ws_c.Close()
						danmulog.L(`I: `,"停止，等待服务器断开连接")
						break_sign = true
					}

				}

				if break_sign {break}
			}

			p.Sys().Timeoutf(1)
		}

		close(interrupt)
		{//附加功能 直播流
			reply.Saveflv_wait()
			reply.Save_to_json(-1, []interface{}{`{}]`})
		}
		danmulog.L(`I: `,"结束退出")
	}
}

