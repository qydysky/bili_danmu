package f

import (
	"context"
	"io/fs"
	"iter"
	"net/http"
	"time"

	_ "github.com/qydysky/bili_danmu/Reply/F/ass"
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuCountPerMin"
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuEmotes"
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuMerge"
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuji"
	_ "github.com/qydysky/bili_danmu/Reply/F/keepMedalLight"
	_ "github.com/qydysky/bili_danmu/Reply/F/lessDanmu"
	_ "github.com/qydysky/bili_danmu/Reply/F/parseM3u8"
	_ "github.com/qydysky/bili_danmu/Reply/F/rev"
	_ "github.com/qydysky/bili_danmu/Reply/F/videoFastSeed"
	comp "github.com/qydysky/part/component2"
	log "github.com/qydysky/part/log"
)

var LessDanmu = comp.Get[interface {
	Init(maxNumSec int)
	InitRoom(roomid int)
	Unset()
	Do(s string) (show bool)
}](`lessDanmu`)

var DanmuMerge = comp.Get[interface {
	InitSend(f func(roomid int, num uint, msg string)) bool
	Init(ctx context.Context, roomid int)
	Unset()
	Do(s string) uint
}](`danmuMerge`)

var KeepMedalLight = comp.Get[interface {
	Init(L *log.Log_interface, Roomid int, SendDanmu func(danmu string, RoomID int) error, PreferDanmu any)
	Clear()
	Do(prefer ...string)
}](`keepMedalLight`)

var Rev = comp.Get[interface {
	Init(l *log.Log_interface)
	ShowRev(roomid int, rev float64)
}](`rev`)

var DanmuCountPerMin = comp.Get[interface {
	// will WriteHeader
	GetRec(savePath string, r *http.Request, w http.ResponseWriter) error
	CheckRoot(dir string)
	Rec(ctx context.Context, roomid int, savePath string) func(any)
	Do(roomid int, msg string, uid string)
}](`danmuCountPerMin`)

var Ass = comp.Get[interface {
	ToAss(savePath string, filename ...string)
	Init(cfg any)
}](`ass`)

var Danmuji = comp.Get[interface {
	Danmujif(s string, then func(string))
	Danmuji_auto(ctx context.Context, danmus []any, waitSec float64, then func(string))
}](`danmuji`)

var VideoFastSeed = comp.Get[interface {
	InitGet(fastSeedFilePath string) (getIndex func(seedTo time.Duration) (int64, error), e error)
	InitSav(fastSeedFilePath string) (savIndex func(seedTo time.Duration, cuIndex int64) error, e error)
}](`videoFastSeed`)

var ParseM3u8 = comp.Get[interface {
	Parse(respon []byte, lastNo int) (m4sLink iter.Seq[interface {
		IsHeader() bool
		M4sLink() string
	}], redirectUrl string, err error)
	IsErrRedirect(e error) bool
}](`parseM3u8`)

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
	IsErrNoEmote(e error) bool
	PackEmotes(dir string) error
	GetEmotesDir(dir string) fs.FS
}](`danmuEmotes`)
