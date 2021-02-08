package F

import (
	"net/http"
	"encoding/json"
    "time"
	"github.com/skratchdot/open-golang/open"
	websocket "github.com/qydysky/part/websocket"
	msgq "github.com/qydysky/part/msgq"
	web "github.com/qydysky/part/web"
	c "github.com/qydysky/bili_danmu/CV"
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
)

func init() {
	go func() {
		for {
			v,ok := c.K_v["get_xiao_xinxin"]
			if !ok {
				time.Sleep(time.Second)
				continue
			}
			if t,ok := v.(bool);!ok || !t {return}
			break
		}
	
		//初始化web服务器，初始化websocket
		server()
		wslog.L(`I: `,`启动`)
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
				wslog.L(`E: `,data)
				return false
			},
		})
	}

	w := web.New(&http.Server{})//新建web实例
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
	wslog.L(`I: `,`如需加密，会自动打开`,webpath)
}

func Wasm(maxloop int, uid uintptr,s RT) (o string) {//maxloop 超时重试
	if maxloop <= 0 {return}

	for try:=5;try > 0 && ws.Len() == 0;try-=1 {//没有从池中取出
		open.Run(webpath)
		wslog.L(`I: `,`浏览器打开`,webpath)
		time.Sleep(time.Second*time.Duration(3))
	}

	b, e := json.Marshal(s)
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
			if r.Id != s.R.Id {break}//或许接收到之前的请求，校验Id字段
			return r.S
		case <- time.After(time.Second*time.Duration(3)):
			wslog.L(`E: `,`超时！响应>1s，确认保持`,webpath,`开启`)
			o = Wasm(maxloop-1, uid, s)
			return
		}
	}
}

func test(uid uintptr) bool {
	time.Sleep(time.Second*time.Duration(3))
	if s := Wasm(3, uid, RT{
		R:R{
		Id: "[9,371,1,22613059]",
		Device: "[\"AUTO8216117272375373\",\"77bee604-b591-4664-845b-b69603f8c71c\"]",
		Ets: 1611836581,
		Benchmark: "seacasdgyijfhofiuxoannn",
		Time: 60,
		Ts: 1611836642190,
		},
		T: []int{2, 5, 1, 4},
	});s != `e4249b7657c2d4a44955548eb814797d41ddd99bfdfa5974462b8c387d701b8c83898f6d7dde1772c67fad6a113d20c20e454be1d1627e7ea99617a8a1f99bd0` {
		wslog.L(`E: `,`测试未通过`,s)
		return false
	}
	return true
}