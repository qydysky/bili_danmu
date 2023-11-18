package danmuReLiveTriger

import (
	"context"
	"regexp"
	"strings"
	"sync/atomic"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	comp "github.com/qydysky/part/component"
	log "github.com/qydysky/part/log"
)

// path
var (
	Init  = comp.NewComp(initf)
	Check = comp.NewComp(check)
)

// 指定弹幕重启录制
var (
	logg      *log.Log_interface
	common    *c.Common
	streamCut func(i int, title ...string)
	reload    atomic.Bool
	l         map[string]*regexp.Regexp
)

type DanmuReLiveTriger struct {
	StreamCut func(i int, title ...string)
	C         *c.Common
}

func initf(ctx context.Context, ptr DanmuReLiveTriger) error {
	l = make(map[string]*regexp.Regexp)
	if v, ok := ptr.C.K_v.LoadV(`指定弹幕重启录制`).([]any); ok && len(v) > 0 {
		logg = ptr.C.Log.Base("指定弹幕重启录制")
		streamCut = ptr.StreamCut
		common = ptr.C
		for i := 0; i < len(v); i++ {
			var item = v[i].(map[string]any)
			var uid = strings.TrimSpace(item["uid"].(string))
			var danmu = strings.TrimSpace(item["danmu"].(string))
			if uid != "" && danmu != "" {
				if reg, e := regexp.Compile(danmu); e != nil {
					clear(l)
					return e
				} else {
					l[uid] = reg
				}
			}
		}
		if len(l) != 0 {
			logg.L(`T: `, `加载`, len(l), `条规则`)
		}
	}
	return nil
}

type Danmu struct {
	Uid, Msg string
}

func check(ctx context.Context, item Danmu) error {
	if reg, ok := l[item.Uid]; ok {
		if ss := reg.FindStringSubmatch(item.Msg); len(ss) > 0 {
			if reload.CompareAndSwap(false, true) {
				switch len(ss) {
				case 1:
					logg.L(`I: `, item.Uid, item.Msg, "请求重启录制")
					streamCut(common.Roomid)
				case 2:
					logg.L(`I: `, item.Uid, ss[1], "请求重启录制带标题")
					streamCut(common.Roomid, ss[1])
				}
				time.AfterFunc(time.Minute, func() {
					reload.Store(false)
				})
			}
		}
	}
	return nil
}
