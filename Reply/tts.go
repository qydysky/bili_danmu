package reply

import (
	"os/exec"
	"net/url"
	"strings"
	c "github.com/qydysky/bili_danmu/CV"

	p "github.com/qydysky/part"
	msgq "github.com/qydysky/part/msgq"
	s "github.com/qydysky/part/buf"
	reqf "github.com/qydysky/part/reqf"
	limit "github.com/qydysky/part/limit"
)

var tts_setting = map[string]string{
	"0buyguide":"感谢{D}",
	"0gift":"感谢{D}",
	"0superchat":"感谢{D}",
}
var tts_List = make(chan interface{},20)

var tts_limit = limit.New(1,5000,15000)//频率限制1次/5s，最大等待时间15s

var tts_log = c.Log.Base_add(`TTS`)

var (
	tts_prog = "ffplay"
	tts_prog_set = "-autoexit -nodisp"
)

func init(){
	{//tts配置
		buf := s.New()
		buf.Load("config/config_K_v.json")
		for k,v := range buf.B {
			if k == `TTS_使用程序路径` {
				if tmp,ok := v.(string);ok && tmp != ``{
					tts_prog = tmp
				} else{tts_log.L(`E: `,`TTS_使用程序路径不是字符串或为空`)}
			}else if k == `TTS_使用程序参数` {
				if tmp,ok := v.(string);ok{
					tts_prog_set = tmp
				} else{tts_log.L(`E: `,`TTS_使用程序参数不是字符串`)}
			}else if k == `TTS_总开关` {
				if tmp,ok := v.(bool);ok && !tmp{
					return
				}
			}
		}
		buf.Load("config/config_tts.json")
		for k,v := range buf.B {
			tts_setting[k] = v.(string)
		}
	}
	//启动程序
	p.Exec().Start(exec.Command(tts_prog))

	go func(){
		for{
			e := <- tts_List
			TTS(e.(Danmu_mq_t).uid, e.(Danmu_mq_t).msg)
		}
	}()
	
	//消息队列接收tts类消息，并传送到TTS朗读
	//使用带tag的消息队列在功能间传递消息
	c.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
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

	v,ok := tts_setting[uid]
	if !ok || v == ``{return}

	tts_log.L(`I: `,uid, strings.ReplaceAll(msg, "\n", " "))

	msg = strings.ReplaceAll(v, "{D}", msg)

	req := reqf.Req()
	if err := req.Reqf(reqf.Rval{
		Url:`https://fanyi.baidu.com/gettts?lan=zh&text=`+ url.QueryEscape(msg) +`&spd=5&source=web`,
		SaveToPath:p.Sys().Cdir()+`/tts.mp3`,
		Timeout:3,
		Retry:1,
		SleepTime:500,
	});err != nil {
		tts_log.L(`E: `,err)
		return
	}
	var prog = []string{}
	prog = append(prog, p.Sys().Cdir()+"/tts.mp3")
	prog = append(prog, strings.Split(tts_prog_set," ")...)
	p.Exec().Run(false, tts_prog, prog...)
	return
}