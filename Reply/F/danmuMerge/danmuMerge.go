package danmumerge

import (
	"context"
	"sync"
	"time"

	comp "github.com/qydysky/part/component2"
)

type TargetInterface interface {
	InitSend(f func(roomid int, num uint, msg string)) bool
	Init(ctx context.Context, roomid int)
	Unset()
	Do(s string) uint
}

func init() {
	if e := comp.Register[TargetInterface]("danmuMerge", &danmuMerge{
		buf:    make(map[string]*danmuMergeItem),
		ticker: time.NewTicker(time.Duration(2) * time.Second),
	}); e != nil {
		panic(e)
	}
}

type danmuMerge struct {
	f      func(roomid int, num uint, msg string)
	roomid int
	buf    map[string]*danmuMergeItem
	now    uint
	ticker *time.Ticker
	sync.RWMutex
}

type danmuMergeItem struct {
	Exprie uint
	Num    uint
}

// Unset implements TargetInterface.
func (t *danmuMerge) Unset() {
	t.Lock()
	defer t.Unlock()
	t.roomid = 0
}

// InitSend implements TargetInterface.
func (t *danmuMerge) InitSend(f func(roomid int, num uint, msg string)) bool {
	t.Lock()
	defer t.Unlock()
	t.f = f
	return false
}

func (t *danmuMerge) Init(ctx context.Context, roomid int) {
	t.Lock()
	defer t.Unlock()
	clear(t.buf)
	t.roomid = roomid
	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.now += 1
			case <-ctx.Done():
				return
			}

			t.Lock()
			if len(t.buf) != 0 {
				for k, v := range t.buf {
					if v.Exprie <= t.now {
						//超时显示
						delete(t.buf, k)
						t.f(roomid, v.Num, k)
					}
				}
			}
			t.Unlock()
		}
	}()
}

// Do implements TargetInterface.
func (t *danmuMerge) Do(s string) uint {
	t.RLock()
	defer t.RUnlock()

	if t.roomid == 0 || t.f == nil {
		return 0
	}

	{ //验证是否已经存在
		if v, ok := t.buf[s]; ok && t.now < v.Exprie {
			t.buf[s].Num += 1
			return v.Num
		}
	}
	{ //设置
		t.buf[s] = &danmuMergeItem{
			Exprie: t.now + 8,
			Num:    1,
		}
	}
	return 0
}
