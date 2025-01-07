package f

import (
	"context"
	"net/http"
	"time"

	_ "github.com/qydysky/bili_danmu/Reply/F/ass"
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuCountPerMin"
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuEmotes"
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuji"
	comp "github.com/qydysky/part/component2"
	log "github.com/qydysky/part/log"
)

var DanmuCountPerMin = comp.Get[interface {
	// will WriteHeader
	GetRec(savePath string, r *http.Request, w http.ResponseWriter) error
	Rec(ctx context.Context, roomid int, savePath string)
	Do(roomid int)
}](`danmuCountPerMin`)

var Ass = comp.Get[interface {
	Assf(s string) error
	Ass_f(ctx context.Context, enc, savePath string, st time.Time)
}](`ass`)

var Danmuji = comp.Get[interface {
	Danmujif(s string, then func(string))
	Danmuji_auto(ctx context.Context, danmus []any, waitSec float64, then func(string))
}](`danmuji`)

type DanmuEmotesS struct {
	Logg *log.Log_interface
	Info []any
	Msg  *string
}

var DanmuEmotes = comp.Get[interface {
	SaveEmote(ctx context.Context, ptr struct {
		Logg *log.Log_interface
		Info []any
		Msg  *string
	}) (ret any, err error)
	Hashr(s string) (r string)
	SetLayerN(n int)
}](`danmuEmotes`)
