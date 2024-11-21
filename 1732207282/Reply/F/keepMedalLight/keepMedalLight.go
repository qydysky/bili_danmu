package keepMedalLight

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/qydysky/bili_danmu/F"
	p "github.com/qydysky/part"

	comp "github.com/qydysky/part/component"

	log "github.com/qydysky/part/log"
)

var (
	Main = comp.NewComp(main)
	rand = p.Rand()
	skip atomic.Bool
)

type Func struct {
	Uid         int
	Logg        *log.Log_interface
	BiliApi     F.BiliApiInter
	SendDanmu   func(danmu string, RoomID int) error
	PreferDanmu []any
}

func main(ctx context.Context, ptr Func) (ret any, err error) {
	if !skip.CompareAndSwap(false, true) {
		return
	}

	ptr.Logg.L(`T: `, `开始`)
	t := time.NewTicker(time.Second * 5)

	defer func() {
		t.Stop()
		ptr.Logg.L(`I: `, `完成`)
		skip.Store(false)
	}()

	if e, list := ptr.BiliApi.GetFansMedal(0, 0); e != nil {
		ptr.Logg.L(`E: `, e)
	} else {
		for _, i := range list {

			select {
			case <-ctx.Done():
				return
			case <-t.C:
			}

			if len(ptr.PreferDanmu) > 0 {
				if s, ok := ptr.PreferDanmu[rand.MixRandom(0, int64(len(ptr.PreferDanmu)-1))].(string); ok {
					if e := ptr.SendDanmu(s, i.RoomID); e != nil {
						ptr.Logg.L(`E: `, e)
					}
				}
			} else if e := ptr.SendDanmu(`点赞`, i.RoomID); e != nil {
				ptr.Logg.L(`E: `, e)
			}
		}
	}

	if e, list := ptr.BiliApi.GetFansMedal(0, 0); e != nil {
		ptr.Logg.L(`E: `, e)
	} else {
		for _, i := range list {

			select {
			case <-ctx.Done():
				return
			case <-t.C:
			}

			if i.IsLighted == 0 {
				if e, his := ptr.BiliApi.GetHisDanmu(i.RoomID); e != nil {
					ptr.Logg.L(`E: `, e)
				} else if len(his) > 0 {
					if e := ptr.SendDanmu(his[0], i.RoomID); e != nil {
						ptr.Logg.L(`E: `, e)
					}
				}
			}
		}
	}

	return
}
