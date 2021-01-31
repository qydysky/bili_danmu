package reply

import (
	s "github.com/qydysky/part/buf"
)

//从config.json初始化
func init(){
	buf := s.New()
	buf.Load("config/config_F.json")
	for k,v := range buf.B {
		AllF[k] = v.(bool)
	}
} 
