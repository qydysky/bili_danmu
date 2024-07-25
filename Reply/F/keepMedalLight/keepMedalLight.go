package keepMedalLight

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/qydysky/bili_danmu/F"
	p "github.com/qydysky/part"

	comp "github.com/qydysky/part/component"
	reqf "github.com/qydysky/part/reqf"

	log "github.com/qydysky/part/log"
)

var (
	Main            = comp.NewComp(main)
	rand            = p.Rand()
	roomI           = make(map[int]*room)
	skip            atomic.Bool
	ErrNotFoundRoom = errors.New(`ErrNotFoundRoom`)
	ErrNotLight     = errors.New(`ErrNotLight`)
)

type Func struct {
	Uid         int
	Logg        *log.Log_interface
	BiliApi     F.BiliApiInter
	SendDanmu   func(danmu string, RoomID int) error
	PreferDanmu []any
}

type room struct {
	medalID  int
	targetID int
	danmu    int
	like     int
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
		return nil, e
	} else {
		for _, v := range list {
			// 熄灭的徽章只能通过送礼物点亮
			if v.TodayFeed > 0 {
				delete(roomI, v.RoomID)
				continue
			} else if v.IsLighted == 1 {
				roomI[v.RoomID] = &room{
					targetID: v.TargetID,
					medalID:  v.MedalID,
				}
			}
		}
	}

	ptr.Logg.L(`I: `, "等待保持点亮数量", len(roomI))

	for len(roomI) != 0 && err == nil {
		// deal roomI
		for roomid, v := range roomI {

			select {
			case <-ctx.Done():
				return
			case <-t.C:
			}

			if e, tmp := ptr.BiliApi.GetFansMedal(roomid, v.targetID); e != nil && !reqf.IsTimeout(e) {
				err = e
			} else if len(tmp) == 0 {
				err = ErrNotFoundRoom
			} else if tmp[0].TodayFeed > 0 {
				delete(roomI, roomid)
				ptr.Logg.L(`I: `, roomid, "已获得亲密度", tmp[0].TodayFeed, "剩余", len(roomI))
				continue
			} else if tmp[0].LivingStatus == 1 {
				if e := ptr.BiliApi.LikeReport(15, ptr.Uid, roomid, v.targetID); e == nil {
					v.like += 15
				} else if !reqf.IsTimeout(e) {
					err = e
				}
			} else {
				if len(ptr.PreferDanmu) > 0 {
					if s, ok := ptr.PreferDanmu[rand.MixRandom(0, int64(len(ptr.PreferDanmu)-1))].(string); ok {
						if e := ptr.SendDanmu(s, roomid); e == nil {
							v.danmu += 1
						} else if !reqf.IsTimeout(e) {
							err = e
						}
					}
				} else if e := ptr.SendDanmu(`点赞`, roomid); e == nil {
					v.danmu += 1
				} else if !reqf.IsTimeout(e) {
					err = e
				}
			}

			if v.danmu < 10 && v.like < 50 {
				// 发送弹幕：每日首次发送弹幕达10条可获得70亲密度
				// 给主播点赞：每日首次点满50个赞可获得50亲密度
				continue
			} else if v.danmu > 20 || v.like > 70 {
				delete(roomI, roomid)
				ptr.Logg.L(`I: `, roomid, "未获得亲密度")
				break
			}
		}
	}

	return
}
