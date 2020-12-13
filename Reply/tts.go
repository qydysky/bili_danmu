//+build tts

package reply

import (
	"log"
	"net/url"
	p "github.com/qydysky/part"
)

var tts_on = map[string]bool{
	"0buyguide":true,
	"0gift":true,
	"0superchat":true,
}
var tts_List = make(chan interface{},20)

var tts_limit = p.Limit(1,5000,15000)//频率限制1次/5s，最大等待时间15s

func init(){
	go func(){
		var (
			sig = Danmu_mq.Sig()
			data interface{}
		)
		go func(){
			for{
				e := <- tts_List
				TTS(e.(Danmu_mq_t).uid, e.(Danmu_mq_t).msg)
			}
		}()

		for {
			data,sig = Danmu_mq.Pull(sig)
			if v,ok := tts_on[data.(Danmu_mq_t).uid];!ok || !v {continue}
			tts_List <- data
		}
	}()
}


func TTS(uid,msg string) {
	if tts_limit.TO() {return}
	log.Println(`TTS:`, uid, msg)
	req := p.Req()
	if err := req.Reqf(p.Rval{
		Url:`https://fanyi.baidu.com/gettts?lan=zh&text=`+ url.QueryEscape(msg) +`&spd=5&source=web`,
		SaveToPath:p.Sys().Cdir()+`/tts.mp3`,
		Timeout:3,
		Retry:1,
		SleepTime:500,
	});err != nil {
		log.Println(`TTS:`, err)
		return
	}
	p.Exec().Run(false, "ffplay", p.Sys().Cdir()+"/tts.mp3","-autoexit","-nodisp")
	return
}