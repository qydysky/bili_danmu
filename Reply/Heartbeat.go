package reply

import (
	"strconv"

	p "github.com/qydysky/part"
	F "github.com/qydysky/bili_danmu/F"
)

var heartlog = p.Logf().New().Base(-1, "Heart.go").Open("danmu.log").Fileonly(true)
var Heart_map = map[string]func(replyF, string) {
	"heartbeat":replyF.heartbeat,//人气
}

func Heart(b []byte){
	s := strconv.Itoa(int(F.Btoi32(b, 0)))
	if F,ok := Heart_map["heartbeat"]; ok {
		var f replyF
		F(f, s);
	}
}