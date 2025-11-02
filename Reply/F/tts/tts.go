package reply

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"

	c "github.com/qydysky/bili_danmu/CV"

	p "github.com/qydysky/part"
	comp "github.com/qydysky/part/component2"
	pctx "github.com/qydysky/part/ctx"
	file "github.com/qydysky/part/file"
	phash "github.com/qydysky/part/hash"
	limit "github.com/qydysky/part/limit"
	log "github.com/qydysky/part/log"
	reqf "github.com/qydysky/part/reqf"
	pstrings "github.com/qydysky/part/strings"
	sync "github.com/qydysky/part/sync"
	sys "github.com/qydysky/part/sys"
	ws "github.com/qydysky/part/websocket"
)

type ttsI interface {
	Init(ctx context.Context, l *log.Log_interface, config any)
	Deal(uid string, m map[string]string)
	Clear()
	Stop()
}

type tts struct {
	ctx               context.Context
	cancle            func()
	l                 *log.Log_interface
	ttsSer            string
	ttsSerMap         map[string]func(string) error
	ttsProg           string
	ttsProgSet        string
	ttsList           chan string
	ttsSettingString  map[string]string
	ttsSettingReplace map[string]string
	youdaoId          string
	youdaoappKey      string
	xfId              string
	xfKey             string
	xfSecret          string
	xfVoice           string
	xfVmap            map[string]bool
	lock              sync.RWMutex
}

var tts_limit = limit.New(1, "5s", "15s") //频率限制1次/5s，最大等待时间15s

func init() {
	t := &tts{
		cancle:     func() {},
		ttsSer:     "baidu",
		ttsSerMap:  make(map[string]func(string) error),
		ttsProg:    "ffplay",
		ttsProgSet: "-autoexit -nodisp",
		ttsSettingString: map[string]string{
			"0buyguide":  "感谢{D}",
			"0gift":      "感谢{D}",
			"0superchat": "感谢{D}",
		},
		ttsSettingReplace: map[string]string{
			"\n": " ",
		},
		ttsList: make(chan string, 20),
		xfVoice: "random",
		xfVmap: map[string]bool{
			`xiaoyan`:   true,
			`aisjiuxu`:  true,
			`aisxping`:  true,
			`aisjinger`: true,
			`aisbabyxu`: true,
		},
	}
	t.ttsSerMap[`baidu`] = t.baidu
	t.ttsSerMap[`youdao`] = t.youdao
	t.ttsSerMap[`xf`] = t.xf
	comp.RegisterOrPanic[ttsI](`tts`, t)
}

func (t *tts) Init(ctx context.Context, l *log.Log_interface, config any) {
	if m, ok := config.(map[string]any); !ok {
		return
	} else {
		defer t.lock.Lock()()

		t.l = l.Base_add(`TTS`)

		{ //tts配置
			if v, ok := m[`TTS_总开关`].(bool); ok && !v {
				return
			}
			if v, ok := m[`TTS_使用程序路径`].(string); ok && v != `` {
				t.ttsProg = v
			} else {
				t.l.L(`E: `, `TTS_使用程序路径不是字符串或为空`)
			}
			if v, ok := m[`TTS_使用程序参数`].(string); ok && v != `` {
				t.ttsProgSet = v
			} else {
				t.l.L(`E: `, `TTS_使用程序参数不是字符串`)
			}
			if v, ok := m[`TTS_服务器`].(string); ok && v != "" {
				if _, ok := t.ttsSerMap[v]; ok {
					t.ttsSer = v
				} else {
					t.ttsSer = `baidu`
				}
			}
			if v, ok := m[`TTS_服务器_youdaoId`].(string); ok && v != `` {
				t.youdaoId = v
			}
			if v, ok := m[`TTS_服务器_youdaoKey`].(string); ok && v != `` {
				t.youdaoappKey = v
			}
			if t.ttsSer == `youdao` && (t.youdaoId == `` || t.youdaoappKey == ``) {
				t.l.L(`W: `, `未提供youdaoId、Key，使用baidu`)
				t.ttsSer = `baidu`
			}
			if v, ok := m[`TTS_服务器_xfId`].(string); ok && v != `` {
				t.xfId = v
			}
			if v, ok := m[`TTS_服务器_xfKey`].(string); ok && v != `` {
				t.xfKey = v
			}
			if v, ok := m[`TTS_服务器_xfSecret`].(string); ok && v != `` {
				t.xfSecret = v
			}
			if v, ok := m[`TTS_服务器_xfVoice`].(string); ok && v != `` {
				if _, ok := t.xfVmap[v]; ok || v == `random` {
					t.xfVoice = v
				} else {
					t.l.L(`I: `, `未支持设定发音，使用随机`)
				}
			}
			if t.ttsSer == `xf` && (t.xfId == `` || t.xfKey == `` || t.xfSecret == ``) {
				t.l.L(`W: `, `未提供讯飞Id、Key、Secret，使用baidu`)
				t.ttsSer = `baidu`
			}

			f := file.New("config/config_tts.json", 0, true)
			if !f.IsExist() {
				return
			}
			bb, err := f.ReadAll(100, 1<<16)
			if !errors.Is(err, io.EOF) {
				return
			}
			var buf map[string]any
			_ = json.Unmarshal(bb, &buf)
			if onoff, ok := buf[`onoff`]; ok {
				for k, v := range onoff.(map[string]any) {
					t.ttsSettingString[k] = v.(string)
				}
			}
			if replace, ok := buf[`replace`]; ok {
				for k, v := range replace.(map[string]any) {
					t.ttsSettingReplace[k] = v.(string)
				}
			}
		}
		//启动程序
		p.Exec().Start(exec.Command(t.ttsProg))

		t.ctx, t.cancle = pctx.WaitCtx(ctx)
		go func() {
			defer t.cancle()
			for {
				select {
				case <-t.ctx.Done():
					return
				case s := <-t.ttsList:
					if !tts_limit.TO() {
						for len(t.ttsList) > 0 && len(s) < 100 {
							s += " " + <-t.ttsList
						}
						var err error

						ulock := t.lock.RLock()
						if f, ok := t.ttsSerMap[t.ttsSer]; ok {
							err = f(s)
						} else {
							err = t.baidu(s)
						}
						ulock()

						if err != nil {
							t.l.L(`E: `, err)
							return
						}
					}
				}
			}
		}()
	}
}

func (t *tts) Deal(uid string, m map[string]string) {
	defer t.lock.RLock()()

	if t.ctx == nil {
		return
	}

	if s, ok := t.ttsSettingString[uid]; ok && len(m) != 0 && s != "" {
		for k, v := range m {
			s = strings.ReplaceAll(s, k, v)
		}
		for k, v := range t.ttsSettingReplace {
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

		t.l.L(`I: `, uid, string(runel))
		t.ttsList <- string(runel)
	}
}

func (t *tts) Clear() {
	for {
		select {
		case <-t.ttsList:
		default:
			return
		}
	}
}

func (t *tts) Stop() {
	t.cancle()
}

func (t *tts) play() {
	var prog = []string{}
	prog = append(prog, sys.Sys().Cdir()+"/tts.mp3")
	prog = append(prog, strings.Split(t.ttsProgSet, " ")...)
	p.Exec().Run(false, t.ttsProg, prog...)
}

func (t *tts) baidu(msg string) error {
	req := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(req)
	if err := req.Reqf(reqf.Rval{
		Url:        `https://fanyi.baidu.com/gettts?lan=zh&text=` + url.PathEscape(msg) + `&spd=5&source=web`,
		SaveToPath: sys.Sys().Cdir() + `/tts.mp3`,
		Timeout:    3 * 1000,
		Retry:      1,
		SleepTime:  5000,
		Proxy:      c.C.Proxy,
	}); err != nil {
		return err
	}
	t.play()
	return nil
}

func (t *tts) youdao(msg string) error {
	//https://ai.youdao.com/gw.s#/
	var (
		api = map[string]string{
			`q`:            msg,
			`langType`:     "zh-CHS",
			`youdaoappKey`: t.youdaoId,
			`salt`:         pstrings.Rand(1, 8),
		}
		postS string
	)
	api[`sign`] = strings.ToUpper(phash.Md5String(api[`youdaoappKey`] + api[`q`] + api[`salt`] + t.youdaoappKey))
	for k, v := range api {
		if postS != "" {
			postS += "&"
		}
		postS += k + `=` + v
	}

	req := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(req)
	if err := req.Reqf(reqf.Rval{
		Url:        `https://openapi.youdao.com/ttsapi`,
		PostStr:    url.PathEscape(postS),
		SaveToPath: sys.Sys().Cdir() + `/tts.mp3`,
		Timeout:    3 * 1000,
		Retry:      1,
		SleepTime:  5000,
		Proxy:      c.C.Proxy,
	}); err != nil {
		return err
	}
	if req.ResHeader().Get(`Content-type`) == `application/json` {
		return req.Response(func(r *http.Response) error {
			return req.Respon(func(b []byte) error {
				return errors.New(`错误 ` + r.Status + string(b))
			})
		})
	}
	t.play()
	return nil
}

func (t *tts) xf(msg string) error {
	voice := t.xfVoice
	if voice == `random` {
		for k := range t.xfVmap {
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

	{
		//msg
		// @hosturl :  like  wss://tts-api.xfyun.cn/v2/tts
		// @apikey : apiKey
		// @apiSecret : apiSecret
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

		var postS = rec{}
		postS.Common.AppID = t.xfId
		postS.Business.Aue = "lame"
		postS.Business.Sfl = 1
		postS.Business.Tte = "UTF8"
		postS.Business.Vcn = voice
		postS.Data.Status = 2
		postS.Data.Text = base64.StdEncoding.EncodeToString([]byte(msg))

		if b, e := json.Marshal(postS); e != nil {
			return e
		} else {

			if wsUrl, err := assembleAuthUrl("wss://tts-api.xfyun.cn/v2/tts", t.xfKey, t.xfSecret); err != nil {
				return err
			} else {
				xfwsClient, _ := ws.New_client(&ws.Client{
					Url:   wsUrl,
					Proxy: c.C.Proxy,
					Header: map[string]string{
						`User-Agent`:      c.UA,
						`Accept`:          `*/*`,
						`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
						`Pragma`:          `no-cache`,
						`Cache-Control`:   `no-cache`,
					},
				})
				wsc, _ := xfwsClient.Handle()
				if xfwsClient.Isclose() {
					return xfwsClient.Error()
				} else {
					var buf []byte
					wait, cancel := context.WithCancel(t.ctx)
					defer cancel()

					var someErr = errors.New(`someErr`)
					wsc.Pull_tag_only(`recv`, func(wm *ws.WsMsg) (disable bool) {
						return wm.Msg(func(b []byte) error {
							if len(b) == 0 {
								cancel()
								return someErr
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
							if e := json.Unmarshal(b, &partS); e != nil {
								xfwsClient.Close()
								return e
							} else {
								if partS.Code != 0 {
									cancel()
									return someErr
								}
								if partS.Data.Audio != "" {
									if part, e := base64.StdEncoding.DecodeString(partS.Data.Audio); e != nil {
										cancel()
										return someErr
									} else {
										buf = append(buf, part...)
									}
								}
								if partS.Data.Status == 2 {
									cancel()
									return someErr
								}
							}
							return nil
						}) != nil
					})

					wsc.Push_tag(`send`, &ws.WsMsg{
						Msg: func(f func([]byte) error) error {
							return f(b)
						},
					})

					<-wait.Done()
					if len(buf) != 0 {
						_, _ = file.New(sys.Sys().Cdir()+`/tts.mp3`, 0, true).Write(buf)
						t.play()
					}
					xfwsClient.Close()
				}
			}
		}
	}
	return nil
}
