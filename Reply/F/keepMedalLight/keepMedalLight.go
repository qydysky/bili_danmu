package keepMedalLight

import (
	"container/ring"
	"sync"
	"time"

	p "github.com/qydysky/part"

	comp2 "github.com/qydysky/part/component2"

	log "github.com/qydysky/part/log/v2"
)

// 点赞30次、发送弹幕10条，可点亮勋章3天
// 点赞 = 1P ； 弹幕 = 3P

func init() {
	comp2.RegisterOrPanic[interface {
		Init(L *log.Log, Roomid int, SendDanmu func(danmu string, RoomID int) error, PreferDanmu any)
		Clear()
		// 在所有可以发送点赞/弹幕的地方都加上，会评估是否需要点赞/弹幕，当prefer存在时，必然发送一条
		Do(prefer ...string)
	}](`keepMedalLight`, &keepMedalLight{})
}

type keepMedalLight struct {
	roomid       int
	log          *log.Log
	sendDanmu    func(danmu string, RoomID int) error
	preferDanmu  []any
	hisPointTime *ring.Ring
	l            sync.RWMutex
}

func (t *keepMedalLight) Init(L *log.Log, Roomid int, SendDanmu func(danmu string, RoomID int) error, PreferDanmu any) {
	t.l.Lock()
	defer t.l.Unlock()

	t.roomid = Roomid
	t.log = L
	t.sendDanmu = SendDanmu
	t.hisPointTime = ring.New(30)
	if ds, ok := PreferDanmu.([]any); ok {
		t.preferDanmu = append(t.preferDanmu, ds...)
	}
}

func (t *keepMedalLight) Clear() {
	t.l.Lock()
	defer t.l.Unlock()

	t.roomid = 0
}

// 发送的时机
func (t *keepMedalLight) Do(prefer ...string) {
	t.l.Lock()
	defer t.l.Unlock()

	if t.roomid == 0 || t.sendDanmu == nil {
		return
	}

	var waitToSend string

	if len(prefer) > 0 {
		waitToSend = prefer[0]
	} else if d, ok := t.hisPointTime.Value.(time.Time); ok && time.Since(d) < time.Hour*24 {
		// 环中最后一个时间在1天内
		return
	} else if d, ok := t.hisPointTime.Prev().Value.(time.Time); ok && time.Since(d) < time.Second*100 {
		// 100s最多发一次
		return
	} else if len(t.preferDanmu) > 0 {
		if s, ok := t.preferDanmu[p.Rand().MixRandom(0, int64(len(t.preferDanmu)-1))].(string); ok {
			waitToSend = s
		}
	}

	if waitToSend == `` {
		waitToSend = `点赞`
	}

	if waitToSend == `点赞` {
		t.getPoint(1)
	} else {
		t.getPoint(3)
	}

	t.log.T(`保持亮牌`)
	if e := t.sendDanmu(waitToSend, t.roomid); e != nil {
		t.log.E(e)
	}
}

func (t *keepMedalLight) getPoint(n int) {
	for range n {
		t.hisPointTime.Value = time.Now()
		t.hisPointTime = t.hisPointTime.Next()
	}
}
