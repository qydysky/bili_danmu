package bili_danmu

import (
	"time"

	"github.com/gorilla/websocket"

	c "github.com/qydysky/bili_danmu/CV"
	p "github.com/qydysky/part"
)

type ws struct {
	used bool
	
	SendChan chan []byte
	RecvChan chan []byte
	interrupt chan struct{}
	url string
}

func New_ws(url string) (o *ws) {
	l := p.Logf().New().Base(-1, "ws.go>新建").Level(c.LogLevel).T("New_ws")
	defer l.Block()

	l.T("ok")
	o = new(ws)
	o.url = url
	o.SendChan = make(chan []byte, 1e4)
	o.RecvChan = make(chan []byte, 1e4)
	return 
}

func (i *ws) Handle() (o *ws) {
	o = i
	l := p.Logf().New().Base(-1, "ws.go>处理").Level(c.LogLevel).T("*ws.handle")
	defer l.Block()

	if o.used {
		l.E("o.used")
		return
	}

	if o.url == "" {
		l.E("o.url == \"\"")
		return
	}

	started := make(chan struct{})

	go func() {
		defer func(){
			close(o.RecvChan)
			o.used = false
		}()

		c, _, err := websocket.DefaultDialer.Dial(o.url, nil)
		if err != nil {
			l.E(err)
			return
		}
		defer c.Close()

		l.T("ok")
		o.interrupt = make(chan struct{})
		done := make(chan struct{})

		go func() {
			defer close(done)

			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					if e, ok := err.(*websocket.CloseError); ok {
						switch e.Code {
						case websocket.CloseNormalClosure:l.E("服务器关闭连接")
						case websocket.CloseAbnormalClosure:l.E("服务器中断连接")
						default:l.E(err);
						}
					}
					return
				}
				o.RecvChan <- message
			}
		}()

		close(started)

		for {
			select {
			case <- done:
				return
			case t := <- o.SendChan:
				err := c.WriteMessage(websocket.TextMessage, t)
				if err != nil {
					l.E("write:", err)
					return
				}
			case <- o.interrupt:
				l.I("捕获到中断")
				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					l.E(err)
				}
				select {
				case <- done:
				case <- time.After(time.Second):
				}
				return
			}
		}
	}()

	<- started
	o.used = true
	return
}

func (i *ws) Heartbeat(Millisecond int, msg []byte) (o *ws) {
	o = i
	l := p.Logf().New().Base(-1, "ws.go>心跳").Level(c.LogLevel).T("*ws.heartbeat")
	defer l.Block()

	if !o.used {
		l.E("!o.used")
		return
	}
	o.SendChan <- msg
	l.T("ok")

	go func(){
		ticker := time.NewTicker(time.Duration(Millisecond)*time.Millisecond)
		defer ticker.Stop()

		for {
			select {
				case <-ticker.C:
					o.SendChan <- msg
				case <- o.interrupt:
					l.I("停止！")
					return
				}
		}
	}()

	return
}

func (o *ws) Close() {
	l := p.Logf().New().Base(-1, "ws.go>关闭").Level(c.LogLevel)
	defer l.Block()

	if !o.used {
		l.E("未在使用的连接")
		return
	}
	o.used = false

	close(o.interrupt)
	l.I("关闭!")
}

func (o *ws) Isclose() bool {
	return !o.used
}