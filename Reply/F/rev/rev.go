package rev

import (
	"fmt"
	"sync"
	"time"

	comp "github.com/qydysky/part/component2"
	log "github.com/qydysky/part/log/v2"
)

func init() {
	comp.RegisterOrPanic[interface {
		Init(l *log.Log)
		ShowRev(roomid int, rev float64)
	}](`rev`, &rev{})
}

type rev struct {
	l           *log.Log
	currentRoom int
	currentRev  float64
	lastShow    time.Time
	sync.Mutex
}

func (t *rev) Init(l *log.Log) {
	t.Lock()
	defer t.Unlock()

	t.l = l.Base(`营收`)
}

func (t *rev) ShowRev(roomid int, rev float64) {
	t.Lock()
	defer t.Unlock()

	if t.l == nil {
		return
	}

	if roomid != t.currentRoom {
		if t.currentRoom != 0 {
			t.l.I(fmt.Sprintf("%d ￥%.2f", t.currentRoom, t.currentRev))
		}
		t.l.I(fmt.Sprintf("%d ￥%.2f", roomid, rev))
	} else if rev != t.currentRev && time.Since(t.lastShow).Minutes() > 1 {
		t.lastShow = time.Now()
		t.l.I(fmt.Sprintf("%d ￥%.2f", roomid, rev))
	}
	t.currentRev = rev
	t.currentRoom = roomid
}
