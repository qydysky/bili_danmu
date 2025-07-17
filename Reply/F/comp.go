package f

import (
	"context"
	"io/fs"
	"iter"
	"net/http"
	"time"

	_ "github.com/qydysky/bili_danmu/Reply/F/ass"              //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuCountPerMin" //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuEmotes"      //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuMerge"       //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuji"          //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/keepMedalLight"   //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/lessDanmu"        //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/parseM3u8"
	_ "github.com/qydysky/bili_danmu/Reply/F/rev"           //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/tts"           //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/videoFastSeed" //removable
	comp "github.com/qydysky/part/component2"
	log "github.com/qydysky/part/log"
)

type TTSI interface {
	Init(ctx context.Context, l *log.Log_interface, config any)
	Deal(uid string, m map[string]string)
	Clear()
	Stop()
}

var TTS = comp.GetV2(`tts`, comp.PreFuncErr[TTSI]{})

type LessDanmuI interface {
	Init(maxNumSec int)
	InitRoom(roomid int)
	Unset()
	Do(s string) (show bool)
}

var LessDanmu = comp.GetV2(`lessDanmu`, comp.PreFuncErr[LessDanmuI]{})

type DanmuMergeI interface {
	InitSend(f func(roomid int, num uint, msg string)) bool
	Init(ctx context.Context, roomid int)
	Unset()
	Do(s string) uint
}

var DanmuMerge = comp.GetV2(`danmuMerge`, comp.PreFuncErr[DanmuMergeI]{})

type KeepMedalLightI interface {
	Init(L *log.Log_interface, Roomid int, SendDanmu func(danmu string, RoomID int) error, PreferDanmu any)
	Clear()
	Do(prefer ...string)
}

var KeepMedalLight = comp.GetV2(`keepMedalLight`, comp.PreFuncErr[KeepMedalLightI]{})

type RevI interface {
	Init(l *log.Log_interface)
	ShowRev(roomid int, rev float64)
}

var Rev = comp.GetV2(`rev`, comp.PreFuncErr[RevI]{})

type DanmuCountPerMinI interface {
	// will WriteHeader
	GetRec(savePath string, r *http.Request, w http.ResponseWriter) error
	CheckRoot(dir string)
	Rec(ctx context.Context, roomid int, savePath string) func(any)
	Do(roomid int, msg string, uid string)
}

var DanmuCountPerMin = comp.GetV2(`danmuCountPerMin`, comp.PreFuncErr[DanmuCountPerMinI]{})

type AssI interface {
	ToAss(savePath string, filename ...string)
	Init(cfg any)
}

var Ass = comp.GetV2(`ass`, comp.PreFuncErr[AssI]{})

type DanmujiI interface {
	Danmujif(s string, then func(string))
	Danmuji_auto(ctx context.Context, danmus []any, waitSec float64, then func(string))
}

var Danmuji = comp.GetV2(`danmuji`, comp.PreFuncErr[DanmujiI]{})

type VideoFastSeedI interface {
	InitGet(fastSeedFilePath string) (getIndex func(seedTo time.Duration) (int64, error), e error)
	InitSav(fastSeedFilePath string) (savIndex func(seedTo time.Duration, cuIndex int64) error, e error)
}

var VideoFastSeed = comp.GetV2(`videoFastSeed`, comp.PreFuncErr[VideoFastSeedI]{})

type ParseM3u8I interface {
	Parse(respon []byte, lastNo int) (m4sLink iter.Seq[interface {
		IsHeader() bool
		M4sLink() string
	}], redirectUrl string, err error)
	IsErrRedirect(e error) bool
}

var ParseM3u8 = comp.GetV2(`parseM3u8`, comp.PreFuncPanic[ParseM3u8I]{})

type DanmuEmotesS struct {
	Logg *log.Log_interface
	Info []any
	Msg  *string
}

type DanmuEmotesI interface {
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
}

var DanmuEmotes = comp.GetV2(`danmuEmotes`, comp.PreFuncErr[DanmuEmotesI]{})
