package recStartEnd

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	comp "github.com/qydysky/part/component"
	log "github.com/qydysky/part/log/v2"
	psync "github.com/qydysky/part/sync"
	"golang.org/x/exp/slices"
)

var (
	InitF         = comp.NewComp(initf)
	RecStartCheck = comp.NewComp(recStartCheck)
	LoopCheck     = comp.NewComp(loopCheck)
)

type dur struct {
	start int
	end   int
}

var (
	logg        *log.Log
	roomSetting map[int][]dur
	timePoints  []int
)

func initf(ctx context.Context, ptr *c.Common) (_ any, err error) {
	if list, ok := ptr.K_v.LoadV("指定房间录制区间").([]any); ok {
		logg = ptr.Log.Base("功能", "指定房间录制区间")
		defer func() {
			if err != nil {
				clear(roomSetting)
				clear(timePoints)
			}
		}()
		if roomSetting == nil {
			roomSetting = make(map[int][]dur)
		}
		clear(roomSetting)
		for _, v := range list {
			if vm, ok := v.(map[string]any); ok {
				if roomid, ok := vm["roomid"].(float64); ok && int(roomid) > 0 {
					var durs []dur
					if sts, ok := vm["fromTo"].([]any); ok {
						for _, v := range sts {
							if vm, ok := v.(map[string]any); ok {
								var durv dur
								if start, ok := vm["start"].(string); ok {
									if tt, e := time.Parse(time.TimeOnly, start); e != nil {
										err = e
										return
									} else {
										durv.start = tt.Hour()*3600 + tt.Minute()*60 + tt.Second() + 1
										timePoints = append(timePoints, durv.start)
									}
								}
								if end, ok := vm["end"].(string); ok {
									if tt, e := time.Parse(time.TimeOnly, end); e != nil {
										err = e
										return
									} else {
										durv.end = tt.Hour()*3600 + tt.Minute()*60 + tt.Second() + 1
										timePoints = append(timePoints, durv.end)
									}
								}
								durs = append(durs, durv)
							}
						}
					}
					logg.T("加载规则", fmt.Sprintf("%d %d条", int(roomid), len(durs)))
					roomSetting[int(roomid)] = durs
				}
			}
		}
		slices.Sort(timePoints)
	}
	return nil, nil
}

func recStartCheck(ctx context.Context, ptr *c.Common) (any, error) {
	if setting, ok := roomSetting[ptr.Roomid]; ok {
		now := time.Now()
		t := now.Hour()*3600 + now.Minute()*60 + now.Second() + 1
		var hasSpace = false
		for _, v := range setting {
			if v.start != 0 && v.end != 0 {
				hasSpace = true
				if t <= v.end && t >= v.start {
					return nil, nil
				}
			}
		}
		if hasSpace {
			return nil, errors.New("当前不在设定时间段内")
		}
	}
	return nil, nil
}

type StreamCtl struct {
	C       *c.Common
	Commons *psync.Map
	State   func(int) bool
	Start   func(int)
	End     func(int)
	Cut     func(int)
}

var streamCtl StreamCtl

func loopCheck(ctx context.Context, ptr StreamCtl) (any, error) {
	streamCtl = ptr
	setNextFunc()
	return nil, nil
}

func setNextFunc() {
	if len(timePoints) == 0 {
		return
	}

	now := time.Now()
	t := now.Hour()*3600 + now.Minute()*60 + now.Second() + 1

	var tmp []int
	for i := 0; i < len(timePoints); i++ {
		if t > timePoints[i] {
			tmp = append(tmp, timePoints[i]+60*60*24-t)
		} else {
			tmp = append(tmp, timePoints[i]-t)
		}
	}
	slices.Sort(tmp)

	// logg.T("下个时间点", time.Second*time.Duration(tmp[0]))

	time.AfterFunc(time.Second*time.Duration(tmp[0]), func() {
		roomId := streamCtl.C.Roomid
		common := streamCtl.Commons.LoadV(roomId).(*c.Common)
		if common.Liveing {
			if setting, ok := roomSetting[roomId]; ok {
				now := time.Now()
				t := now.Hour()*3600 + now.Minute()*60 + now.Second() + 1
				for _, v := range setting {
					if v.start != 0 && math.Abs(float64(t-v.start)) < 5 {
						if streamCtl.State(roomId) {
							logg.T("切片", roomId)
							streamCtl.Cut(roomId)
						} else {
							logg.T("开始", roomId)
							streamCtl.Start(roomId)
						}
						time.Sleep(time.Second * 5)
						break
					}
					if v.end != 0 && math.Abs(float64(t-v.end)) < 5 {
						if streamCtl.State(roomId) {
							logg.T("结束", roomId)
							streamCtl.End(roomId)
						}
						time.Sleep(time.Second * 5)
						break
					}
				}
			}
		}
		setNextFunc()
	})
}
