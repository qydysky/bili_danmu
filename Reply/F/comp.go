package f

import (
	"context"
	"io"
	"io/fs"
	"iter"
	"time"

	_ "github.com/qydysky/bili_danmu/Reply/F/ass"              //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuCountPerMin" //removable //replay
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuEmotes"      //removable //replay
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuMerge"       //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/danmuji"          //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/genCpuPprof"      //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/keepMedalLight"   //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/lessDanmu"        //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/parseM3u8"
	_ "github.com/qydysky/bili_danmu/Reply/F/rev"           //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/saveDanmuToDB" //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/saveToJson"    //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/shortDanmu"    //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/tts"           //removable
	_ "github.com/qydysky/bili_danmu/Reply/F/videoFastSeed" //removable
	comp "github.com/qydysky/part/component2"
	log "github.com/qydysky/part/log/v2"
)

var GenCpuPprof = comp.GetV3[interface {
	Start(ctx context.Context, file string) (any, error)
}](`genCpuPprof`)

type TTSI interface {
	Init(ctx context.Context, l *log.Log, config any)
	Deal(uid string, m map[string]string)
	Clear()
	Stop()
}

var TTS = comp.GetV3[TTSI](`tts`)

type LessDanmuI interface {
	Init(maxNumSec int)
	InitRoom(roomid int)
	Unset()
	Do(s string) (show bool)
}

var LessDanmu = comp.GetV3[LessDanmuI](`lessDanmu`)

type DanmuMergeI interface {
	InitSend(f func(roomid int, num uint, msg string)) bool
	Init(ctx context.Context, roomid int)
	Unset()
	Do(s string) uint
}

var DanmuMerge = comp.GetV3[DanmuMergeI](`danmuMerge`)

type KeepMedalLightI interface {
	Init(L *log.Log, Roomid int, SendDanmu func(danmu string, RoomID int) error, PreferDanmu any)
	Clear()
	Do(prefer ...string)
}

var KeepMedalLight = comp.GetV3[KeepMedalLightI](`keepMedalLight`)

type RevI interface {
	Init(l *log.Log)
	ShowRev(roomid int, rev float64)
}

var Rev = comp.GetV3[RevI](`rev`)

type DanmuCountPerMinI interface {
	// will WriteHeader
	GetRec4(savePath string, points *[]int) (e error)
	CheckRoot(dir string)
	Rec(ctx context.Context, roomid int, savePath string) func(any)
	Do(roomid int, msg string, uid string)
}

var DanmuCountPerMin = comp.GetV3[DanmuCountPerMinI](`danmuCountPerMin`)

type AssI interface {
	ToAss(savePath string, filename ...string)
	Init(cfg any)
}

var Ass = comp.GetV3[AssI](`ass`)

type DanmujiI interface {
	Danmujif(s string, then func(string))
	Danmuji_auto(ctx context.Context, danmus []any, waitSec float64, then func(string))
}

var Danmuji = comp.GetV3[DanmujiI](`danmuji`)

type VideoFastSeedI interface {
	InitGet(fastSeedFilePath string) (getIndex func(seedTo time.Duration) (int64, error), e error)
	InitSav(fastSeedFilePath string) (savIndex func(seedTo time.Duration, cuIndex int64) error, delete func(), e error)
}

var VideoFastSeed = comp.GetV3[VideoFastSeedI](`videoFastSeed`)

type ParseM3u8I interface {
	Parse(respon []byte, lastNo int) (m4sLink iter.Seq[interface {
		IsHeader() bool
		M4sLink() string
	}], redirectUrl string, err error)
	IsErrRedirect(e error) bool
}

var ParseM3u8 = comp.GetV3(`parseM3u8`, comp.PreFuncPanic[ParseM3u8I]{})

type DanmuEmotesS struct {
	Logg *log.Log
	Info []any
	Msg  *string
}

type DanmuEmotesI interface {
	SaveEmote(ctx context.Context, ptr struct {
		Logg *log.Log
		Info []any
		Msg  *string
	}) (ret any, err error)
	Hashr(s string) (r string)
	SetLayerN(n int)
	IsErrNoEmote(e error) bool
	PackEmotes(dir string) error
	GetEmotesDir(dir string) interface {
		fs.FS
		io.Closer
	}
}

var DanmuEmotes = comp.GetV3[DanmuEmotesI](`danmuEmotes`)

var ShortDanmu = comp.GetV3[interface {
	Deal(string) string
}](`shortDanmu`)

type SaveToJsonI interface {
	Close()
	Init(path any)
	Write(data *[]byte)
}

var SaveToJson = comp.GetV3[SaveToJsonI](`saveToJson`)

type SaveDanmuToDBI interface {
	Init(config any, fl *log.Log)
	Danmu(Msg string, Color string, Auth any, Uid string, Roomid int64)
	Close() error
}

var SaveDanmuToDB = comp.GetV3[SaveDanmuToDBI](`saveDanmuToDB`)
