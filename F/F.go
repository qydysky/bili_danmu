package F

import (
	"bytes"
	"strconv"

	p "github.com/qydysky/part"
	. "github.com/qydysky/bili_danmu/Const"
)

var flog = p.Logf().New().Open("danmu.log").Base(-1, "F.go")

//base on source/player-loader-2.0.7.min.js L3313
//base on source/player-loader-2.0.7.min.js L3455
type header struct {
	PackL int32
	HeadL int16
	BodyV int16
	OpeaT int32
	Seque int32
}

//头部生成与检查
func HeadGen(datalenght,Opeation,Sequence int) []byte {
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.Write(Itob32(int32(datalenght + WS_PACKAGE_HEADER_TOTAL_LENGTH)))
	buffer.Write(Itob16(WS_PACKAGE_HEADER_TOTAL_LENGTH))
	buffer.Write(Itob16(WS_HEADER_DEFAULT_VERSION))
	buffer.Write(Itob32(int32(Opeation)))
	buffer.Write(Itob32(int32(Sequence)))

	return buffer.Bytes()
}

func HeadChe(head []byte) (header) {

	if len(head) != WS_PACKAGE_HEADER_TOTAL_LENGTH {flog.Base(1, "头部检查").E("输入头长度错误");return header{}}

	PackL := Btoi32(head, WS_PACKAGE_OFFSET)
	HeadL := Btoi16(head, WS_HEADER_OFFSET)
	BodyV := Btoi16(head, WS_VERSION_OFFSET)
	OpeaT := Btoi32(head, WS_OPERATION_OFFSET)
	Seque := Btoi32(head, WS_SEQUENCE_OFFSET)

	return header{
		PackL :PackL,
		HeadL :HeadL,
		BodyV :BodyV,
		OpeaT :OpeaT,
		Seque :Seque,
	}
}

//认证生成与检查
func HelloGen(roomid int, key string) []byte {
	flog.Base(-1, "认证生成")
	defer flog.Base(0)

	if roomid == 0 || key == "" {
		flog.E("roomid == 0 || key == \"\"")
		return []byte("")
	}
	
	//from player-loader-2.0.7.min.js
	/*
		customAuthParam
	*/
	const (
		_uid = 0
		_protover = 2
		_platform = "web"
		VERSION = "2.0.7"
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
	
	buffer.Write(HeadGen(len(obj), WS_OP_USER_AUTHENTICATION, WS_HEADER_DEFAULT_SEQUENCE))

	buffer.Write([]byte(obj))

	return buffer.Bytes()
}

func HelloChe(r []byte) bool {
	if len(r) == 0 {return false}

	var obj = `{"code":0}`

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.Write(HeadGen(len(obj), WS_OP_CONNECT_SUCCESS, WS_HEADER_DEFAULT_SEQUENCE))

	buffer.Write([]byte(obj))

	h := buffer.Bytes()

	if len(h) != len(r) {return false}

	for k, v := range r {
		if v != h[k] {return false}
	}
	return true
}

//获取人气生成
func Heartbeat() ([]byte, int) {
	//from player-loader-2.0.4.min.js
	const heartBeatInterval = 30

	var obj = `[object Object]`

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	
	buffer.Write(HeadGen(len(obj), WS_OP_HEARTBEAT, WS_HEADER_DEFAULT_SEQUENCE))

	buffer.Write([]byte(obj))

	return buffer.Bytes(), heartBeatInterval

}