//+build tts

package reply

import (
	"fmt"
	"net/url"
	"strings"
	p "github.com/qydysky/part"
	c "github.com/qydysky/bili_danmu/CV"
	s "github.com/qydysky/part/buf"
)

var tts_setting = map[string]string{
	"0buyguide":"感谢{D}",
	"0gift":"感谢{D}",
	"0superchat":"感谢{D}",
}
var tts_List = make(chan interface{},20)

var tts_limit = p.Limit(1,5000,15000)//频率限制1次/5s，最大等待时间15s

func init(){
	{//tts配置
		buf := s.New()
		buf.Load("config/config_tts.json")
		for k,v := range buf.B {
			tts_setting[k] = v.(string)
		}
	}
	go func(){
		for{
			e := <- tts_List
			TTS(e.(Danmu_mq_t).uid, e.(Danmu_mq_t).msg)
		}
	}()
	
	//消息队列接收tts类消息，并传送到TTS朗读
	//使用带tag的消息队列在功能间传递消息
	c.Danmu_Main_mq.Pull_tag(map[string]func(interface{})(bool){
		`tts`:func(data interface{})(bool){//tts
			if _,ok := tts_setting[data.(Danmu_mq_t).uid];ok {
				tts_List <- data
			}
			return false
		},
	})
}


func TTS(uid,msg string) {
	if tts_limit.TO() {return}
	fmt.Println(`TTS:`, uid, msg)
	req := p.Req()
	if v,ok := tts_setting[uid];ok{
		msg = strings.ReplaceAll(v, "{D}", msg)
	}
	if err := req.Reqf(p.Rval{
		Url:`https://fanyi.baidu.com/gettts?lan=zh&text=`+ url.QueryEscape(msg) +`&spd=5&source=web`,
		SaveToPath:p.Sys().Cdir()+`/tts.mp3`,
		Timeout:3,
		Retry:1,
		SleepTime:500,
	});err != nil {
		fmt.Println(`TTS:`, err)
		return
	}
	p.Exec().Run(false, "ffplay", p.Sys().Cdir()+"/tts.mp3","-autoexit","-nodisp")
	return
}