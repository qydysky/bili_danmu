package bili_danmu

import (
	"fmt"
	"bytes"
	"strconv"
	"os"
	"os/signal"

	p "github.com/qydysky/part"
)

const LogLevel = 1
var danmulog = p.Logf().New().Open("danmu.log").Base(-1, "bili_danmu.go").Level(LogLevel)

func Demo() {
	danmulog.Base(-1, "测试")
	defer danmulog.Base(0)
	
	//ctrl+c退出
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	{
		var room int
		fmt.Printf("输入房间号: ")
		_, err := fmt.Scanln(&room)
		if err != nil {
			danmulog.E("输入错误", err)
			return
		}
		
		var break_sign bool
		for !break_sign {
			//获取房间相关信息
			api := New_api(room).Get_host_Token()
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
				ws.SendChan <- hello_send(api.Roomid, api.Token)
				if hello_ok(<- ws.RecvChan) {
					danmulog.I("已连接到房间", room)

					//开始心跳
					go func(){
						danmulog.I("开始心跳")
						p.Sys().MTimeoutf(500)//500ms
						heartbeatmsg, heartinterval := heartbeat()
						ws.Heartbeat(1000 * heartinterval, heartbeatmsg)
						
						//打招呼
						if p.Checkfile().IsExist("cookie.txt") {
							f := p.File().FileWR(p.Filel{
								File:"cookie.txt",
								Write:false,
							})
							//传输变量至Msg，以便响应弹幕"弹幕机在么"
							Msg_roomid = api.Roomid
							Msg_cookie = f
							Danmuji_auto(Msg_cookie, 5, Msg_roomid)
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
							go Reply(i)
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

		danmulog.I("结束退出")
	}
}

//from player-loader-2.0.4.min.js
const (
	WS_PACKAGE_HEADER_TOTAL_LENGTH = 16
	WS_HEADER_DEFAULT_VERSION = 1
	WS_BODY_PROTOCOL_VERSION_NORMAL = 0
	WS_BODY_PROTOCOL_VERSION_DEFLATE = 2
	WS_OP_USER_AUTHENTICATION = 7
	WS_OP_HEARTBEAT = 2
	WS_HEADER_DEFAULT_SEQUENCE = 1
	WS_OP_CONNECT_SUCCESS = 8
	WS_OP_MESSAGE = 5
	WS_OP_HEARTBEAT_REPLY = 3
)

//返回数据分派
func Reply(b []byte) {

	if ist, _ := headChe(b[:16], len(b), WS_BODY_PROTOCOL_VERSION_DEFLATE, WS_OP_MESSAGE, 0, 4); ist {
		Msg(b, true);return
	}
	if ist, _ := headChe(b[:16], len(b), WS_BODY_PROTOCOL_VERSION_NORMAL, WS_OP_MESSAGE, 0, 4); ist {
		Msg(b, false);return
	}

	danmulog.Base(1, "返回分派")

	if ist, _ := headChe(b[:16], len(b), WS_HEADER_DEFAULT_VERSION, WS_OP_HEARTBEAT_REPLY, WS_HEADER_DEFAULT_SEQUENCE, 4); ist {
		danmulog.T("heartbeat replay!");
		return
	}

	danmulog.T("unknow reply", b)
}

//头部生成与检查
func headGen(datalenght,Opeation,Sequence int) []byte {
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.Write(Itob32(int32(datalenght + WS_PACKAGE_HEADER_TOTAL_LENGTH)))
	buffer.Write(Itob16(WS_PACKAGE_HEADER_TOTAL_LENGTH))
	buffer.Write(Itob16(WS_HEADER_DEFAULT_VERSION))
	buffer.Write(Itob32(int32(Opeation)))
	buffer.Write(Itob32(int32(Sequence)))

	return buffer.Bytes()
}

func headChe(head []byte, datalenght,Bodyv,Opeation,Sequence,show int) (bool,int32) {
	if len(head) != WS_PACKAGE_HEADER_TOTAL_LENGTH {return false, 0}
	
	danmulog.Base(-1, "头部检查").Level(show)
	defer danmulog.Base(0).Level(LogLevel)
	

	packL := Btoi32(head[:4])
	headL := Btoi16(head[4:6])
	BodyV := Btoi16(head[6:8])
	OpeaT := Btoi32(head[8:12])
	Seque := Btoi32(head[12:16])

	if packL > int32(datalenght) {danmulog.E("包缺损", packL, datalenght);return false, packL}
	if headL != WS_PACKAGE_HEADER_TOTAL_LENGTH {danmulog.E("头错误", headL);return false, packL}
	if OpeaT != int32(Opeation) {danmulog.E("类型错误");return false, packL}
	if Seque != int32(Sequence) {danmulog.E("Seq错误");return false, packL}
	if BodyV != int16(Bodyv) {danmulog.E("压缩算法错误");return false, packL}

	return true, packL
}

//认证生成与检查
func hello_send(roomid int, key string) []byte {

	if roomid == 0 || key == "" {
		danmulog.Base(1, "认证生成").E("roomid == 0 || key == \"\"")
		return []byte("")
	}
	
	//from player-loader-2.0.4.min.js
	/*
		customAuthParam
	*/
	const (
		_uid = 0
		_protover = 2
		_platform = "web"
		VERSION = "2.0.4"
		_type = 2
	)

	var obj = `{"uid":` + strconv.Itoa(_uid) + 
	`,"roomid":` + strconv.Itoa(roomid) + 
	`,"protover":` + strconv.Itoa(_protover) + 
	`,"platform":"`+ _platform + 
	`","clientver":"` + VERSION + 
	`","type":` + strconv.Itoa(_type) + 
	`,"key":"` + key + `"}`

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	
	buffer.Write(headGen(len(obj), WS_OP_USER_AUTHENTICATION, WS_HEADER_DEFAULT_SEQUENCE))

	buffer.Write([]byte(obj))

	return buffer.Bytes()
}

func hello_ok(r []byte) bool {
	if len(r) == 0 {return false}

	var obj = `{"code":0}`

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.Write(headGen(len(obj), WS_OP_CONNECT_SUCCESS, WS_HEADER_DEFAULT_SEQUENCE))

	buffer.Write([]byte(obj))

	h := buffer.Bytes()

	if len(h) != len(r) {return false}

	for k, v := range r {
		if v != h[k] {return false}
	}
	return true
}

//心跳生成
func heartbeat() ([]byte, int) {
	//from player-loader-2.0.4.min.js
	const heartBeatInterval = 30

	var obj = `[object Object]`

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	
	buffer.Write(headGen(len(obj), WS_OP_HEARTBEAT, WS_HEADER_DEFAULT_SEQUENCE))

	buffer.Write([]byte(obj))

	return buffer.Bytes(), heartBeatInterval

}