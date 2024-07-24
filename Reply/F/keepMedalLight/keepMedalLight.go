package keepMedalLight

import (
	"context"
	"sync/atomic"
	"time"

	p "github.com/qydysky/part"
	comp "github.com/qydysky/part/component"
)

var (
	KeepMedalLight = comp.NewComp(keepMedalLight)
	skip           atomic.Bool
)

type Func struct {
	LightedRoomID   []int // 熄灭的徽章只能通过送礼物点亮
	SendDanmu       func(danmu string, RoomID int) error
	GetHistoryDanmu func(RoomID int) []string
	PreferDanmu     []any
}

func keepMedalLight(ctx context.Context, ptr Func) (ret any, err error) {
	if !skip.CompareAndSwap(false, true) {
		return
	}
	defer skip.Store(false)

	for i := 0; i < len(ptr.LightedRoomID); i++ {
		time.Sleep(time.Second * 5)

		if len(ptr.PreferDanmu) > 0 {
			rand := p.Rand().MixRandom(0, int64(len(ptr.PreferDanmu)-1))
			if s, ok := ptr.PreferDanmu[rand].(string); ok {
				if e := ptr.SendDanmu(s, ptr.LightedRoomID[i]); e == nil {
					continue
				} else {
					err = e
				}
			}
		}

		if e := ptr.SendDanmu(`点赞`, ptr.LightedRoomID[i]); e == nil {
			continue
		} else {
			err = e
		}

		his := ptr.GetHistoryDanmu(ptr.LightedRoomID[i])

		if len(his) > 0 {
			if e := ptr.SendDanmu(his[0], ptr.LightedRoomID[i]); e == nil {
				continue
			} else {
				err = e
			}
		}
	}

	return
}
