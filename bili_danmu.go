package bili_danmu

import (
	"fmt"
	"flag"
	"strconv"
	"os"
	"os/signal"

	p "github.com/qydysky/part"
	reply "github.com/qydysky/bili_danmu/Reply"
	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
)

var danmulog = p.Logf().New().Open("danmu.log").Base(-1, "bili_danmu.go").Level(c.LogLevel)

func Demo() {

	danmulog.Base(-1, "测试")
	defer danmulog.Base(0)
	
	//ctrl+c退出
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)


	{
		var groomid = flag.Int("r", 0, "roomid")
		flag.Parse()
	
		var room = *groomid
		fmt.Printf("输入房间号: ")
		if room == 0 {
			_, err := fmt.Scanln(&room)
			if err != nil {
				danmulog.E("输入错误", err)
				return
			}
		} else {
			fmt.Print(strconv.Itoa(room), "\n")
		}

		var break_sign bool
		for !break_sign {
			//获取房间相关信息
			api := F.New_api(room).Get_host_Token().Get_live()
			if len(api.Url) == 0 || api.Roomid == 0 || api.Token == "" || api.Uid == 0 {
				danmulog.E("some err")
				return
			}
			danmulog.I("连接到房间", room)

			//对每个弹幕服务器尝试
			for _, v := range api.Url {
				//ws启动
				ws := New_ws(v).Handle()
	
				//SendChan 传入发送[]byte
				//RecvChan 接收[]byte
				danmulog.I("连接", v)
				ws.SendChan <- F.HelloGen(api.Roomid, api.Token)
				if F.HelloChe(<- ws.RecvChan) {
					danmulog.I("已连接到房间", room)

					//30s获取一次人气
					go func(){
						danmulog.I("获取人气")
						p.Sys().MTimeoutf(500)//500ms
						heartbeatmsg, heartinterval := F.Heartbeat()
						ws.Heartbeat(1000 * heartinterval, heartbeatmsg)
						
						//打招呼
						if p.Checkfile().IsExist("cookie.txt") {
							f := p.File().FileWR(p.Filel{
								File:"cookie.txt",
								Write:false,
							})
							//传输变量，以便响应弹幕"弹幕机在么"
							c.Roomid = api.Roomid
							c.Live = api.Live
							c.Cookie = f
							{//附加功能 弹幕机 直播流保存
								reply.Danmuji_auto(1)
								go reply.Saveflvf()
							}
						}
					}()
				}

				var isclose bool
				for !isclose {
					select {
					case i := <- ws.RecvChan:
						if len(i) == 0 && ws.Isclose() {
							isclose = true
						} else {
							go reply.Reply(i)
						}
					case <- interrupt:
						ws.Close()
						danmulog.I("停止，等待服务器断开连接")
						break_sign = true
					}

				}

				if break_sign {break}
			}

			p.Sys().Timeoutf(1)
		}

		{//附加功能 直播流转码
			reply.Saveflv_wait()
		}
		danmulog.I("结束退出")
	}
}

