package reply

import (
	F "github.com/qydysky/bili_danmu/F"
)

/*
	HeartBeat数据分派
*/

var Heart_map = map[string]func(replyF, int){
	"heartbeat": replyF.heartbeat, //人气
}

// HeartBeat类型，将人气4位byte转为字符串，并送到上述map指定的方法
func Heart(replyFS replyF, b []byte) {
	s := int(F.Btoi32(b, 0))
	if F, ok := Heart_map["heartbeat"]; ok {
		F(replyFS, s)
	}
}
