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

var (
	tts_setting_string = map[string]string{
		"0buyguide":"感谢{D}",
		"0gift":"感谢{D}",
		"0superchat":"感谢{D}",
	}
	tts_setting_replace = map[string]string{
		"\n":" ",
	}
)
var tts_List = make(chan string,20)

var tts_limit = limit.New(1,5000,15000)//频率限制1次/5s，最大等待时间15s

var tts_log = c.Log.Base_add(`TTS`)

var (
	tts_ser = "baidu"
	tts_ser_map = map[string]func(string)reqf.Rval{
		`baidu`:baidu,
		`youdao`:youdao,
	}
	tts_prog = "ffplay"
	tts_prog_set = "-autoexit -nodisp"
)

func init(){
	{//tts配置

		if v, ok := c.K_v.LoadV(`TTS_总开关`).(bool);ok && !v {
			return
		}
		if v, ok := c.K_v.LoadV(`TTS_使用程序路径`).(string);ok && v != ``{
			tts_prog = v
		} else {
			tts_log.L(`E: `,`TTS_使用程序路径不是字符串或为空`)
		}
		if v, ok := c.K_v.LoadV(`TTS_使用程序参数`).(string);ok && v != ``{
			tts_prog_set = v
		} else {
			tts_log.L(`E: `,`TTS_使用程序参数不是字符串`)
		}
		if v, ok := c.K_v.LoadV(`TTS_服务器`).(string);ok && v != "" {
			if _,ok := tts_ser_map[v];ok{
				tts_ser = v
			} else {
				tts_log.L(`I: `,`未支持设定服务提供商，使用baidu`)
				tts_ser = `baidu`
			}
		}
		
		buf := s.New()
		buf.Load("config/config_tts.json")
		if onoff,ok := buf.Get(`onoff`);ok {
			for k,v := range onoff.(map[string]interface{}) {
				tts_setting_string[k] = v.(string)
			}
		}
		if replace,ok := buf.Get(`replace`);ok {
			for k,v := range replace.(map[string]interface{}) {
				tts_setting_replace[k] = v.(string)
			}
		}
	}
	//启动程序
	p.Exec().Start(exec.Command(tts_prog))

	go func(){
		for{
			s := <- tts_List
			for len(tts_List) > 0 && len(s) < 100 {
				s += " " + <- tts_List
			}
			TTS(<- tts_List)
		}
	}()
	
	//消息队列接收tts类消息，并传送到TTS朗读
	//使用带tag的消息队列在功能间传递消息
	c.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
		`tts`:func(data interface{})(bool){//tts
			d,_ := data.(Danmu_mq_t)
			if s,ok := tts_setting_string[d.uid];ok && len(d.m) != 0 && s != "" {
				
				for k,v := range d.m {
					s = strings.ReplaceAll(s, k, v)
				}
				for k,v := range tts_setting_replace {
					s = strings.ReplaceAll(s, k, v)
				}
				var (
					skip bool
					runel []rune
				)
				for _,v := range s {
					if v == []rune("{")[0] {skip = true}
					if v == []rune("}")[0] {skip = false;continue}
					if skip {continue}
					runel = append(runel,v)
				}

				tts_log.L(`I: `, d.uid, string(runel))
				tts_List <- string(runel)
			}
			return false
		},
		`change_room`:func(data interface{})(bool){
			for {
				select {
				case <- tts_List:;
				default:return false;
				}
			}
			return false
		},
	})
}


func TTS(msg string) {
	if tts_limit.TO() {return}

	var (
		req = reqf.New()
		rval reqf.Rval
	)
	if f,ok := tts_ser_map[tts_ser];ok{
		rval = f(msg)
	} else {
		rval = baidu(msg)
	}
	if rval.Url == `` {
		tts_log.L(`E: `, `rval.Url`)
		return
	}

	if err := req.Reqf(rval);err != nil {
		tts_log.L(`E: `,err)
		return
	}
	if req.Response.Header.Get(`Content-type`) == `application/json` {
		tts_log.L(`W: `, `错误`, req.Response.StatusCode, string(req.Respon))
		return
	}
	
	var prog = []string{}
	prog = append(prog, p.Sys().Cdir()+"/tts.mp3")
	prog = append(prog, strings.Split(tts_prog_set," ")...)
	p.Exec().Run(false, tts_prog, prog...)
	return
}

func baidu(msg string) reqf.Rval {
	return reqf.Rval{
		Url:`https://fanyi.baidu.com/gettts?lan=zh&text=`+ url.PathEscape(msg) +`&spd=5&source=web`,
		SaveToPath:p.Sys().Cdir()+`/tts.mp3`,
		Timeout:3*1000,
		Retry:1,
		SleepTime:5000,
		Proxy:c.Proxy,
	}
}

var (
	appId string
	appKey string
)
func init(){
	if v, ok := c.K_v.LoadV(`TTS_服务器_youdaoId`).(string);ok && v != ``{
		appId = v
	}
	if v, ok := c.K_v.LoadV(`TTS_服务器_youdaoKey`).(string);ok && v != ``{
		appKey = v
	}
	if tts_ser == `youdao` && (appId == `` || appKey == ``) {
		tts_log.L(`W: `, `未提供youdaoId、Key，使用baidu`)
		tts_ser = `baidu`
	}
}
func youdao(msg string) reqf.Rval {
	if appId == "" || appKey == "" {
		return baidu(msg)
	}

	//https://ai.youdao.com/gw.s#/
	var (
		api = map[string]string{
			`q`:msg,
			`langType`:"zh-CHS",
			`appKey`:appId,
			`salt`:p.Stringf().Rand(1, 8),
		}
		postS string
	)
	api[`sign`] = strings.ToUpper(p.Md5().Md5String(api[`appKey`]+api[`q`]+api[`salt`]+appKey))
	for k,v := range api {
		if postS != "" {postS += "&"}
		postS += k+`=`+v
	}
	return reqf.Rval{
		Url:`https://openapi.youdao.com/ttsapi`,
		PostStr:url.PathEscape(postS),
		SaveToPath:p.Sys().Cdir()+`/tts.mp3`,
		Timeout:3*1000,
		Retry:1,
		SleepTime:5000,
		Proxy:c.Proxy,
	}
}