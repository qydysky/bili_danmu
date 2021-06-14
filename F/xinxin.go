package F

import (
	"net/http"
	"strconv"
	"encoding/json"
    "time"
	"strings"

	c "github.com/qydysky/bili_danmu/CV"

	p "github.com/qydysky/part"
	websocket "github.com/qydysky/part/websocket"
	msgq "github.com/qydysky/part/msgq"
	web "github.com/qydysky/part/web"
	reqf "github.com/qydysky/part/reqf"

	"github.com/skratchdot/open-golang/open"
)

/*
	小心心加密golang-ws-js-webassembly工具
*/

//需要加密的数据
type R struct {
	Id string `json:"id"` 
	Device string `json:"device"`
	Ets int `json:"ets"`
	Benchmark string `json:"benchmark"`
	Time int `json:"time"`
	Ts int `json:"ts"`
	Ua string `json:"ua"`
}

//发送的原始对象
type RT struct {
	R R `json:"r"`
	T []int `json:"t"` //加密方法
}

//返回的加密对象
type S struct {
	Id string `json:"id"`//发送的数据中的Id项，以确保是对应的返回
	S string `json:"s"` //加密字符串
}


//全局对象
var (
	wslog = c.Log.Base(`api`).Base_add(`小心心加密`) //日志
	rec_chan = make(chan S,5)//收通道
	webpath string//web地址,由于实时获取空闲端口，故将在稍后web启动后赋值
	ws = websocket.New_server()//新建websocket实例
	nodeJsUrl string
)

func init() {
	go func() {
		for {
			v,ok := c.K_v.Load("get_xiao_xinxin")
			if !ok {
				time.Sleep(time.Second)
				continue
			}
			if t,ok := v.(bool);!ok || !t {return}
			break
		}
	
		//初始化web服务器，初始化websocket
		NodeJsUrl,ok := c.K_v.LoadV("小心心nodjs加密服务地址").(string)
		if ok && NodeJsUrl != "" {
			nodeJsUrl = NodeJsUrl
			if test(0) {
				wslog.L(`I: `,`使用NodeJs`,NodeJsUrl,`进行加密`)
			} else {
				wslog.L(`E: `,`发生错误！`)
			}
		} else {
			server()
		}
		wslog.L(`T: `,`启动`)
	}()
}

func server() {
	{
		ws_mq := ws.Interface()//获取websocket操作对象
		ws_mq.Pull_tag(msgq.FuncMap{
			`recv`:func(data interface{})(bool){
				if tmp,ok := data.(websocket.Uinterface);ok {//websocket接收并响应
					//websocket.Uinterface{
					// 	Id	uintptr 会话id
					// 	Data []byte 接收的websocket数据
					// }

					var s S
					e := json.Unmarshal(tmp.Data, &s)
					if e != nil {
						wslog.L(`E: `, e, string(tmp.Data))
					}

					select{
					case rec_chan <- s:
					default:
					}

				}
				return false
			},
			`error`:func(data interface{})(bool){//websocket错误
				wslog.L(`W: `, data)
				return false
			},
		})
	}

	w := web.New(&http.Server{
		Addr: "127.0.0.1:"+strconv.Itoa(p.Sys().GetFreePort()),
	})//新建web实例
	w.Handle(map[string]func(http.ResponseWriter,*http.Request){//路径处理函数
		`/`:func(w http.ResponseWriter,r *http.Request){
			var path string = r.URL.Path[1:]
			if path == `` {path = `index.html`}
			http.ServeFile(w, r, "html/"+path)
		},
		`/ws`:func(w http.ResponseWriter,r *http.Request){
			//获取通道
			conn := ws.WS(w,r)
			//由通道获取本次会话id，并测试 提示
			go test(<-conn)
			//等待会话结束，通道释放
			<-conn
		},
	})
	webpath = `http://`+w.Server.Addr
	//提示
	wslog.L(`I: `,`使用WebJs`,webpath,`进行加密`)
}

func Wasm(uid uintptr,rt RT) (so RT, o string) {//maxloop 超时重试
	so = rt

	{//nodejs
		if nodeJsUrl != "" {
			req := reqf.New()
			if err := req.Reqf(reqf.Rval{
				Header:map[string]string{
					`Content-Type`: `application/json`,
				},
				Url:nodeJsUrl,
				PostStr:toNodeJsString(so),
				Proxy:c.Proxy,
				Timeout:3*1000,
			});err != nil {
				wslog.L(`E: `,err)
				so,o = Wasm(uid, so)
				return
			}

			var res struct{
				Code int `json:"code"`
				S string `json:"s"`
				Message string `json:"message"`
			}

			if e := json.Unmarshal(req.Respon, &res);e != nil {
				wslog.L(`E: `,e)
			} else if res.Code != 0 {
				wslog.L(`E: `,res.Message)
			} else {
				o = res.S
			}
			return
		}
	}

	{//web
		for try:=5;try > 0 && ws.Len() == 0;try-=1 {//没有从池中取出
			open.Run(webpath)
			wslog.L(`I: `,`浏览器打开`,webpath)
			time.Sleep(time.Second*time.Duration(10))
		}

		if ws.Len() == 0 {
			wslog.L(`W: `,`浏览器打开`,webpath,`失败，请手动打开`)
			return
		}

		if !strings.Contains(so.R.Ua, `Test`) {
			so.R.Ts = int(p.Sys().GetMTime())
		}
		b, e := json.Marshal(so)
		if e != nil {
			wslog.L(`E: `,e)
		}

		//获取websocket操作对象 发送
		ws.Interface().Push_tag(`send`,websocket.Uinterface{
			Id:uid,
			Data:b,
		})

		for {
			select {
			case r :=<- rec_chan:
				if r.Id != so.R.Id {break}//或许接收到之前的请求，校验Id字段
				return so, r.S
			case <- time.After(time.Second*time.Duration(1)):
				wslog.L(`E: `,`超时！响应>1s，确认保持`,webpath,`开启`)
				return
			}
		}
	}
}

func Close(uid uintptr){
	//nodejs不需要关闭
	if nodeJsUrl != "" {
		return
	}
	//获取websocket操作对象 关闭
	ws.Interface().Push_tag(`close`,websocket.Uinterface{
		Id:uid,
		Data:[]byte(`获取结束、关闭连接`),
	})
}

func test(uid uintptr) bool {
	time.Sleep(time.Second*time.Duration(3))
	if _,s := Wasm(uid, RT{
		R:R{
		Id: "[9,371,1,22613059]",
		Device: "[\"AUTO8216117272375373\",\"77bee604-b591-4664-845b-b69603f8c71c\"]",
		Ets: 1611836581,
		Benchmark: "seacasdgyijfhofiuxoannn",
		Time: 60,
		Ts: 1611836642190,
		Ua:`Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0 Test`,
		},
		T: []int{2, 5, 1, 4},
	});s != `e4249b7657c2d4a44955548eb814797d41ddd99bfdfa5974462b8c387d701b8c83898f6d7dde1772c67fad6a113d20c20e454be1d1627e7ea99617a8a1f99bd0` {
		wslog.L(`E: `,`测试未通过`,s)
		return false
	}
	return true
}

func toNodeJsString(r RT) (o string) {
	o += `{"t":{"id":`
	o += r.R.Id+`,`
	o += `"device":`+r.R.Device+`,`
	o += `"ets":`+strconv.Itoa(r.R.Ets)+`,`
	o += `"benchmark":"`+r.R.Benchmark+`",`
	o += `"time":`+strconv.Itoa(r.R.Time)+`,`
	o += `"ts":`+strconv.Itoa(r.R.Ts)+`,`
	o += `"ua":"`+r.R.Ua+`"},"r":[`
	o += strconv.Itoa(r.T[0])+`,`
	o += strconv.Itoa(r.T[1])+`,`
	o += strconv.Itoa(r.T[2])+`,`
	o += strconv.Itoa(r.T[3])
	o += `]}`
	return
}