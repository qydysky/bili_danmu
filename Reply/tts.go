package reply

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os/exec"
	"strings"
	"time"

	c "github.com/qydysky/bili_danmu/CV"

	p "github.com/qydysky/part"
	funcCtrl "github.com/qydysky/part/funcCtrl"
	limit "github.com/qydysky/part/limit"
	msgq "github.com/qydysky/part/msgq"
	reqf "github.com/qydysky/part/reqf"
	pstrings "github.com/qydysky/part/strings"
	ws "github.com/qydysky/part/websocket"
)

var (
	tts_setting_string = map[string]string{
		"0buyguide":  "感谢{D}",
		"0gift":      "感谢{D}",
		"0superchat": "感谢{D}",
	}
	tts_setting_replace = map[string]string{
		"\n": " ",
	}
)
var tts_List = make(chan string, 20)

var tts_limit = limit.New(1, 5000, 15000) //频率限制1次/5s，最大等待时间15s

var tts_log = c.C.Log.Base_add(`TTS`)

var (
	tts_ser     = "baidu"
	tts_ser_map = map[string]func(string) error{
		`baidu`:  baidu,
		`youdao`: youdao,
		`xf`:     xf,
	}
	tts_prog     = "ffplay"
	tts_prog_set = "-autoexit -nodisp"
)

func init() {
	{ //tts配置

		if v, ok := c.C.K_v.LoadV(`TTS_总开关`).(bool); ok && !v {
			return
		}
		if v, ok := c.C.K_v.LoadV(`TTS_使用程序路径`).(string); ok && v != `` {
			tts_prog = v
		} else {
			tts_log.L(`E: `, `TTS_使用程序路径不是字符串或为空`)
		}
		if v, ok := c.C.K_v.LoadV(`TTS_使用程序参数`).(string); ok && v != `` {
			tts_prog_set = v
		} else {
			tts_log.L(`E: `, `TTS_使用程序参数不是字符串`)
		}
		if v, ok := c.C.K_v.LoadV(`TTS_服务器`).(string); ok && v != "" {
			if _, ok := tts_ser_map[v]; ok {
				tts_ser = v
			} else {
				tts_log.L(`I: `, `未支持设定服务提供商，使用baidu`)
				tts_ser = `baidu`
			}
		}

		bb, err := ioutil.ReadFile("config/config_tts.json")
		if err != nil {
			return
		}
		var buf map[string]interface{}
		json.Unmarshal(bb, &buf)
		if onoff, ok := buf[`onoff`]; ok {
			for k, v := range onoff.(map[string]interface{}) {
				tts_setting_string[k] = v.(string)
			}
		}
		if replace, ok := buf[`replace`]; ok {
			for k, v := range replace.(map[string]interface{}) {
				tts_setting_replace[k] = v.(string)
			}
		}
	}
	//启动程序
	p.Exec().Start(exec.Command(tts_prog))

	go func() {
		for {
			s := <-tts_List
			for len(tts_List) > 0 && len(s) < 100 {
				s += " " + <-tts_List
			}
			TTS(s)
		}
	}()

	//消息队列接收tts类消息，并传送到TTS朗读
	//使用带tag的消息队列在功能间传递消息
	c.C.Danmu_Main_mq.Pull_tag(msgq.FuncMap{
		`tts`: func(data interface{}) bool { //tts
			d, _ := data.(Danmu_mq_t)
			if s, ok := tts_setting_string[d.uid]; ok && len(d.m) != 0 && s != "" {

				for k, v := range d.m {
					s = strings.ReplaceAll(s, k, v)
				}
				for k, v := range tts_setting_replace {
					s = strings.ReplaceAll(s, k, v)
				}
				var (
					skip  bool
					runel []rune
				)
				for _, v := range s {
					if v == []rune("{")[0] {
						skip = true
					}
					if v == []rune("}")[0] {
						skip = false
						continue
					}
					if skip {
						continue
					}
					runel = append(runel, v)
				}

				tts_log.L(`I: `, d.uid, string(runel))
				tts_List <- string(runel)
			}
			return false
		},
		`change_room`: func(data interface{}) bool {
			for {
				select {
				case <-tts_List:
				default:
					return false
				}
			}
			return false
		},
	})
}

func TTS(msg string) {
	if tts_limit.TO() {
		return
	}

	var err error
	if f, ok := tts_ser_map[tts_ser]; ok {
		err = f(msg)
	} else {
		err = baidu(msg)
	}

	if err != nil {
		tts_log.L(`E: `, err)
		return
	}

	return
}

func play() {
	var prog = []string{}
	prog = append(prog, p.Sys().Cdir()+"/tts.mp3")
	prog = append(prog, strings.Split(tts_prog_set, " ")...)
	p.Exec().Run(false, tts_prog, prog...)

}

func baidu(msg string) error {
	req := reqf.New()
	if err := req.Reqf(reqf.Rval{
		Url:        `https://fanyi.baidu.com/gettts?lan=zh&text=` + url.PathEscape(msg) + `&spd=5&source=web`,
		SaveToPath: p.Sys().Cdir() + `/tts.mp3`,
		Timeout:    3 * 1000,
		Retry:      1,
		SleepTime:  5000,
		Proxy:      c.C.Proxy,
	}); err != nil {
		return err
	}
	play()
	return nil
}

var (
	youdaoId     string
	youdaoappKey string
)

func init() {
	if v, ok := c.C.K_v.LoadV(`TTS_服务器_youdaoId`).(string); ok && v != `` {
		youdaoId = v
	}
	if v, ok := c.C.K_v.LoadV(`TTS_服务器_youdaoKey`).(string); ok && v != `` {
		youdaoappKey = v
	}
	if tts_ser == `youdao` && (youdaoId == `` || youdaoappKey == ``) {
		tts_log.L(`W: `, `未提供youdaoId、Key，使用baidu`)
		tts_ser = `baidu`
	}
}
func youdao(msg string) error {
	if youdaoId == "" || youdaoappKey == "" {
		return baidu(msg)
	}

	//https://ai.youdao.com/gw.s#/
	var (
		api = map[string]string{
			`q`:            msg,
			`langType`:     "zh-CHS",
			`youdaoappKey`: youdaoId,
			`salt`:         pstrings.Rand(1, 8),
		}
		postS string
	)
	api[`sign`] = strings.ToUpper(p.Md5().Md5String(api[`youdaoappKey`] + api[`q`] + api[`salt`] + youdaoappKey))
	for k, v := range api {
		if postS != "" {
			postS += "&"
		}
		postS += k + `=` + v
	}

	req := reqf.New()
	if err := req.Reqf(reqf.Rval{
		Url:        `https://openapi.youdao.com/ttsapi`,
		PostStr:    url.PathEscape(postS),
		SaveToPath: p.Sys().Cdir() + `/tts.mp3`,
		Timeout:    3 * 1000,
		Retry:      1,
		SleepTime:  5000,
		Proxy:      c.C.Proxy,
	}); err != nil {
		return err
	}
	if req.Response.Header.Get(`Content-type`) == `application/json` {
		return errors.New(`错误 ` + req.Response.Status + string(req.Respon))
	}
	play()
	return nil
}

var (
	xfId     string
	xfKey    string
	xfSecret string
	xfVoice  = "random"
	xfVmap   = map[string]bool{
		`xiaoyan`:   true,
		`aisjiuxu`:  true,
		`aisxping`:  true,
		`aisjinger`: true,
		`aisbabyxu`: true,
	}
	xfwsClient   *ws.Client
	xf_req       func()
	xf_req_block funcCtrl.BlockFunc
)

func init() {
	if v, ok := c.C.K_v.LoadV(`TTS_服务器_xfId`).(string); ok && v != `` {
		xfId = v
	}
	if v, ok := c.C.K_v.LoadV(`TTS_服务器_xfKey`).(string); ok && v != `` {
		xfKey = v
	}
	if v, ok := c.C.K_v.LoadV(`TTS_服务器_xfSecret`).(string); ok && v != `` {
		xfSecret = v
	}
	if v, ok := c.C.K_v.LoadV(`TTS_服务器_xfVoice`).(string); ok && v != `` {
		if _, ok := xfVmap[v]; ok || v == `random` {
			xfVoice = v
		} else {
			tts_log.L(`I: `, `未支持设定发音，使用随机`)
		}
	}

	//	设置了非讯飞tts
	if tts_ser != `xf` {
		return
	}

	if xfId == `` || xfKey == `` || xfSecret == `` {
		tts_log.L(`W: `, `未提供讯飞Id、Key、Secret，使用baidu`)
		tts_ser = `baidu`
		return
	}

	//@hosturl :  like  wss://tts-api.xfyun.cn/v2/tts
	//@apikey : apiKey
	//@apiSecret : apiSecret
	assembleAuthUrl := func(hosturl string, apiKey, apiSecret string) (string, error) {
		ul, err := url.Parse(hosturl)
		if err != nil {
			return "", err
		}
		//签名时间
		date := time.Now().UTC().Format(time.RFC1123)
		//参与签名的字段 host ,date, request-line
		signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
		//拼接签名字符串
		sgin := strings.Join(signString, "\n")
		//签名
		mac := hmac.New(sha256.New, []byte(apiSecret))
		mac.Write([]byte(sgin))
		sha := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		//构建请求参数 此时不需要urlencoding
		authUrl := fmt.Sprintf("api_key=\"%s\",algorithm=\"%s\",headers=\"%s\",signature=\"%s\"", apiKey, "hmac-sha256", "host date request-line", sha)
		//将请求参数使用base64编码
		authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))
		v := url.Values{}
		v.Add("host", ul.Host)
		v.Add("date", date)
		v.Add("authorization", authorization)
		//将编码后的字符串url encode后添加到url后面
		callurl := hosturl + "?" + v.Encode()
		return callurl, nil
	}

	wsUrl, err := assembleAuthUrl("wss://tts-api.xfyun.cn/v2/tts", xfKey, xfSecret)

	if err != nil {
		tts_log.L(`E: `, `错误,使用百度`, err)
		tts_ser = `baidu`
		return
	}

	xf_req = func() {
		xf_req_block.Block() //cant call in same time
		defer xf_req_block.UnBlock()

		xfwsClient = ws.New_client(ws.Client{
			Url:   wsUrl,
			Proxy: c.C.Proxy,
			Header: map[string]string{
				`User-Agent`:      `Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0`,
				`Accept`:          `*/*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
			},
		}).Handle()
		if xfwsClient.Isclose() {
			tts_log.L(`E: `, "连接错误,使用百度", xfwsClient.Error())
			tts_ser = `baidu`
		} else {
			go func() {
				var buf []byte
				for !xfwsClient.Isclose() {
					data := <-xfwsClient.RecvChan
					if len(data) == 0 {
						break
					}

					var partS struct {
						Code    int    `json:"code"`
						Message string `json:"message"`
						Sid     string `json:"sid"`
						Data    struct {
							Audio  string `json:"audio"`
							Ced    string `json:"ced"`
							Status int    `json:"status"`
						} `json:"data"`
					}
					if e := json.Unmarshal(data, &partS); e != nil {
						tts_log.L(`E: `, "错误", e, data)
						xfwsClient.Close()
						return
					} else {
						if partS.Code != 0 {
							tts_log.L(`W: `, fmt.Sprintf("code:%d msg:%s", partS.Code, partS.Message))
							break
						}
						if partS.Data.Audio != "" {
							if part, e := base64.StdEncoding.DecodeString(partS.Data.Audio); e != nil {
								tts_log.L(`E: `, "错误", e)
								break
							} else {
								buf = append(buf, part...)
							}
						}
						if partS.Data.Status == 2 {
							break
						}
					}
				}
				if len(buf) != 0 {
					p.File().FileWR(p.Filel{
						File:    p.Sys().Cdir() + `/tts.mp3`,
						Context: []interface{}{buf},
					})
					play()
				}
				xfwsClient.Close()
			}()
		}

	}
	xf_req()
}
func xf(msg string) error {
	if xfId == `` || xfKey == `` || xfSecret == `` {
		tts_log.L(`T: `, "参数不足,使用百度")
		return baidu(msg)
	}

	voice := xfVoice
	if voice == `random` {
		for k, _ := range xfVmap {
			voice = k
			break
		}
	}

	type rec struct {
		Common struct {
			AppID string `json:"app_id"`
		} `json:"common"`
		Business struct {
			Aue string `json:"aue"`
			Vcn string `json:"vcn"`
			Tte string `json:"tte"`
			Sfl int    `json:"sfl"`
		} `json:"business"`
		Data struct {
			Status int    `json:"status"`
			Text   string `json:"text"`
		} `json:"data"`
	}

	{ //msg
		var postS = rec{}
		postS.Common.AppID = xfId
		postS.Business.Aue = "lame"
		postS.Business.Sfl = 1
		postS.Business.Tte = "UTF8"
		postS.Business.Vcn = voice
		postS.Data.Status = 2
		postS.Data.Text = base64.StdEncoding.EncodeToString([]byte(msg))

		if b, e := json.Marshal(postS); e != nil {
			return e
		} else {
			if xfwsClient.Isclose() {
				xf_req()
			}
			xfwsClient.SendChan <- b
		}
	}
	return nil
}
