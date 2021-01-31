package F

import (
	"net/http"
	"encoding/json"
    "time"
	"context"
	"strconv"
	"github.com/gorilla/websocket"
	"github.com/skratchdot/open-golang/open"
	p "github.com/qydysky/part"
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
	S string `json:"s"` //加密字符串
}

//全局对象
var (
	xinxinboot = make(chan struct{},1) //调用标记，仅调用一次
	wslog = c.Log.Base(`api`).Base_add(`小心心加密`) //日志
	send_chan = make(chan RT,1)//发通道
	rec_chan = make(chan S,1)//收通道
	port = p.Sys().GetFreePort()//随机端口
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
	
		wslog.L(`T: `,`被调用`)
	
		select{
		case xinxinboot <- struct{}{}: //没有启动实例
			wslog.L(`I: `,`启动`)
			web()
			<- xinxinboot
		default: //有启动实例
			wslog.L(`I: `,`已启动`)
		}
	}()
}

func web() {
	web :=  http.NewServeMux()

	var (
		server *http.Server
		upgrader = websocket.Upgrader{}
	)

	web.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		server.Shutdown(context.Background())
	})

	web.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			wslog.L(`E: `,"upgrade:", err)
			return
		}
		defer ws.Close()

		//测试 提示
		go func(){
			time.Sleep(time.Second*time.Duration(3))
			if s := Wasm(RT{
				R:R{
				Id: "[9,371,1,22613059]",
				Device: "[\"AUTO8216117272375373\",\"77bee604-b591-4664-845b-b69603f8c71c\"]",
				Ets: 1611836581,
				Benchmark: "seacasdgyijfhofiuxoannn",
				Time: 60,
				Ts: 1611836642190,
				},
				T: []int{2, 5, 1, 4},
			});s == `e4249b7657c2d4a44955548eb814797d41ddd99bfdfa5974462b8c387d701b8c83898f6d7dde1772c67fad6a113d20c20e454be1d1627e7ea99617a8a1f99bd0` {
				wslog.L(`I: `,`测试通过`)
			} else {
				wslog.L(`E: `,`测试未通过`,s)
			}
		}()

		var close_chan = make(chan struct{})
		//发送
		go func(){
			for {
				select {
				case <-close_chan:
					return
				case tmp := <- send_chan:
					b, e := json.Marshal(tmp)
					if e != nil {
						wslog.L(`E: `,e)
					}

					if e := ws.WriteMessage(websocket.TextMessage,b);e != nil {
						wslog.L(`E: `,e)
						return
					}
				}
			}
		}()

		//接收
		for {
			ws.SetReadDeadline(time.Now().Add(time.Second*time.Duration(300)))
			if _, message, e := ws.ReadMessage();e != nil {
				if websocket.IsCloseError(e,websocket.CloseGoingAway) {
					wslog.L(`I: `,e)
				} else {
					wslog.L(`E: `,e)
				}
				close(close_chan)
				break
			} else {
				var s S
				e := json.Unmarshal(message, &s)
				if e != nil {
					wslog.L(`E: `, e, string(message))
				}
				rec_chan <- s
			}
		}
	})

	//html js
	web.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var path string = r.URL.Path[1:]
		if path == `` {path = `index.html`}
        http.ServeFile(w, r, "html/"+path)
	})
	
	server = &http.Server{
		Addr: "127.0.0.1:"+strconv.Itoa(port),
		WriteTimeout: time.Second * time.Duration(10),
		Handler: web,
	}

	//测试 提示
	go func(){
		time.Sleep(time.Second*time.Duration(3))
		open.Run("http://127.0.0.1:"+strconv.Itoa(port))
		wslog.L(`I: `,`保持浏览器打开`,"http://127.0.0.1:"+strconv.Itoa(port),`以正常运行`)
	}()

	server.ListenAndServe()
}

func Wasm(s RT) (o string) {
	select{
	case send_chan <- s:
		select {
		case r :=<- rec_chan:
			return r.S
		case <- time.After(time.Second):
			wslog.L(`E: `,`超时！响应>1s，确认保持`,"http://127.0.0.1:"+strconv.Itoa(port),`开启`)
			return
		}
	default:
		return
	}
}