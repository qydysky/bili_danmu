package savetojson

import (
	"sync/atomic"

	comp "github.com/qydysky/part/component2"
	file "github.com/qydysky/part/file"
	msgq "github.com/qydysky/part/msgq"
)

// 保存所有消息到json
func init() {
	comp.RegisterOrPanic[interface {
		Init(path any)
		Write(data *[]byte)
		Close()
	}](`saveToJson`, &saveToJson{})
}

type saveToJson struct {
	state atomic.Int32
	msg   *msgq.MsgType[[]byte]
}

func (t *saveToJson) Init(pathi any) {
	if path, ok := pathi.(string); !ok || pathi == "" {
		return
	} else if !t.state.CompareAndSwap(0, 1) {
		return
	} else {
		f := file.Open(path)
		_ = f.Delete()
		_, _ = f.Write([]byte("["))
		_ = f.Close()

		t.msg = msgq.NewType[[]byte]()
		t.msg.Pull_tag(map[string]func([]byte) (disable bool){
			`data`: func(b []byte) (disable bool) {
				f := file.New(path, -1, false)
				_, _ = f.Write(b)
				_, _ = f.Write([]byte(","))
				_ = f.Close()
				return false
			},
			`stop`: func(_ []byte) (disable bool) {
				f := file.New(path, -1, false)
				_ = f.SeekIndex(-1, file.AtEnd)
				_, _ = f.Write([]byte("]"))
				_ = f.Close()
				return true
			},
		})
	}
}

func (t *saveToJson) Write(data *[]byte) {
	if t.state.Load() == 1 {
		t.msg.PushLock_tag(`data`, *data)
	}
}

func (t *saveToJson) Close() {
	if t.state.Load() == 1 {
		t.msg.PushLock_tag(`stop`, nil)
	}
}
