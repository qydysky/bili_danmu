package bili_danmu

import (
	"time"

	"github.com/gorilla/websocket"

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
	l := p.Logf().New().Level(LogLevel).T("New_ws")
	defer l.Block()

	l.T("->", "ok")
	o = new(ws)
	o.url = url
	o.SendChan = make(chan []byte, 1e4)
	o.RecvChan = make(chan []byte, 1e4)
	return 
}

func (i *ws) Handle() (o *ws) {
	o = i
	l := p.Logf().New().Level(LogLevel).T("*ws.handle")
	defer l.Block()

	if o.used {
		l.E("->", "o.used")
		return
	}

	if o.url == "" {
		l.E("->", "o.url == \"\"")
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
			l.E("->", err)
			return
		}
		defer c.Close()

		l.T("->", "ok")
		o.interrupt = make(chan struct{})
		done := make(chan struct{})

		go func() {
			defer close(done)

			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
						l.E("->", err)
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
					l.E("->", "write:", err)
					return
				}
			case <- o.interrupt:
				l.I("->", "interrupt")
				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					l.E("->", err)
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
	l := p.Logf().New().Level(LogLevel).T("*ws.heartbeat")
	defer l.Block()

	if !o.used {
		l.E("->", "!o.used")
		return
	}
	o.SendChan <- msg
	l.T("->", "ok")

	go func(){
		ticker := time.NewTicker(time.Duration(Millisecond)*time.Millisecond)
		defer ticker.Stop()

		for {
			select {
				case <-ticker.C:
					o.SendChan <- msg
				case <- o.interrupt:
					l.I("->", "fin")
					return
				}
		}
	}()

	return
}

func (o *ws) Close() {
	l := p.Logf().New().Level(LogLevel).I("*ws.Close")
	defer l.Block()

	if !o.used {
		l.I("->", "!o.used")
		return
	}
	o.used = false

	close(o.interrupt)
	l.I("->", "ok")
}

func (o *ws) Isclose() bool {
	return !o.used
}