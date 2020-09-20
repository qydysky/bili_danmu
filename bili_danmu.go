package bili_danmu

import (
	"fmt"
	"bytes"
	"strconv"
	"os"
	"os/signal"
	"compress/zlib"

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
							Danmuji_auto(5)
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
		os.Exit(0)
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
	danmulog.Base(-1, "返回分派").Level(4)
	defer danmulog.Base(0).Level(LogLevel)

	head := headChe(b[:16])
	if int(head.packL) > len(b) {danmulog.E("包缺损");return}

	if head.BodyV == WS_BODY_PROTOCOL_VERSION_DEFLATE {
		readc, err := zlib.NewReader(bytes.NewReader(b[16:]))
		if err != nil {danmulog.E("解压错误");return}
		defer readc.Close()
		
		buf := bytes.NewBuffer(nil)
		if _, err := buf.ReadFrom(readc);err != nil {danmulog.E("解压错误");return}
		b = buf.Bytes()
	}

	for len(b) != 0 {
		head := headChe(b[:16])
		if int(head.packL) > len(b) {danmulog.E("包缺损");return}
		
		contain := b[16:head.packL]
		switch head.OpeaT {
		case WS_OP_MESSAGE:Msg(contain)
		case WS_OP_HEARTBEAT_REPLY:danmulog.T("heartbeat replay!")
		default :danmulog.T("unknow reply", contain)
		}

		b = b[head.packL:]
	}
}

type header struct {
	packL int32
	headL int16
	BodyV int16
	OpeaT int32
	Seque int32
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

func headChe(head []byte) (header) {

	if len(head) != WS_PACKAGE_HEADER_TOTAL_LENGTH {danmulog.Base(1, "头部检查").E("输入头长度错误");return header{}}
	
	packL := Btoi32(head[:4])
	headL := Btoi16(head[4:6])
	BodyV := Btoi16(head[6:8])
	OpeaT := Btoi32(head[8:12])
	Seque := Btoi32(head[12:16])

	return header{
		packL :packL,
		headL :headL,
		BodyV :BodyV,
		OpeaT :OpeaT,
		Seque :Seque,
	}
}

//认证生成与检查
func hello_send(roomid int, key string) []byte {
	danmulog.Base(-1, "认证生成")
	defer danmulog.Base(0)

	if roomid == 0 || key == "" {
		danmulog.E("roomid == 0 || key == \"\"")
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