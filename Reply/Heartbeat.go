package reply

import (
	"strconv"

	p "github.com/qydysky/part"
	F "github.com/qydysky/bili_danmu/F"
)

/*
	HeartBeat数据分派
*/

var heartlog = p.Logf().New().Base(-1, "Heart.go").Open("danmu.log").Fileonly(true)

//HeartBeat类型处理方法map
var Heart_map = map[string]func(replyF, string) {
	"heartbeat":replyF.heartbeat,//人气
}

//HeartBeat类型，将人气4位byte转为字符串，并送到上述map指定的方法
func Heart(b []byte){
	s := strconv.Itoa(int(F.Btoi32(b, 0)))
	if F,ok := Heart_map["heartbeat"]; ok {
		var f replyF
		F(f, s);
	}
}