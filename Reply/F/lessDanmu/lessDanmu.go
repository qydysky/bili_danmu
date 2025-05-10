package lessdanmu

import (
	"sync"

	comp "github.com/qydysky/part/component2"
	limit "github.com/qydysky/part/limit"
	slice "github.com/qydysky/part/slice"
)

type TargetInterface interface {
	Init(maxNumSec int)
	InitRoom(roomid int)
	Unset()
	Do(s string) (show bool)
}

func init() {
	if e := comp.Register[TargetInterface]("lessDanmu", &lessDanmu{
		threshold: 0.7,
	}); e != nil {
		panic(e)
	}
}

type lessDanmu struct {
	roomid    int
	buf       []string
	limit     *limit.Limit
	threshold float32
	maxNumSec int
	sync.RWMutex
}

// InitRoom implements TargetInterface.
func (t *lessDanmu) InitRoom(roomid int) {
	t.Lock()
	defer t.Unlock()
	t.roomid = roomid
	clear(t.buf)
	t.buf = t.buf[:0]
}

// Unset implements TargetInterface.
func (t *lessDanmu) Unset() {
	t.Lock()
	defer t.Unlock()
	t.roomid = 0
}

func (t *lessDanmu) Init(maxNumSec int) {
	t.Lock()
	defer t.Unlock()
	t.limit = limit.New(maxNumSec, "1s", "0s")
}

// Do implements TargetInterface.
func (t *lessDanmu) Do(s string) (show bool) {
	t.Lock()
	defer t.Unlock()

	if t.roomid == 0 {
		return
	}

	if len(t.buf) < 20 {
		t.buf = append(t.buf, s)
		return true
	}
	o := cross(s, t.buf)
	if o == 1 {
		return false
	} //完全无用的弹幕

	slice.DelFront(&t.buf, 1)

	show = o < t.threshold

	if show && t.maxNumSec > 0 {
		return !t.limit.TO()
	}

	return
}

// 字符重复度检查
// a在buf中出现的字符占a的百分数
func cross(a string, buf []string) float32 {
	var s float32
	var matched bool
	for _, v1 := range a {
		for _, v2 := range buf {
			for _, v3 := range v2 {
				if v3 == v1 {
					matched = true
					break
				}
			}
			if matched {
				break
			}
		}
		if matched {
			s += 1
		}
		matched = false
	}
	return s / float32(len([]rune(a)))
}
