package Reply

import (
	"errors"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
	comp "github.com/qydysky/part/component2"
	perrors "github.com/qydysky/part/errors"
	file "github.com/qydysky/part/file"
	slice "github.com/qydysky/part/slice"
)

func Test_FLVdeal(t *testing.T) {
	f := file.Open("testdata/0.flv")
	defer f.CloseErr()

	if f.IsDir() || !f.IsExist() {
		t.Fatal("file not support")
	}

	buf := make([]byte, humanize.MByte)
	buff := slice.New[byte](10 * humanize.MByte)
	init := false
	flvDecoder := NewFlvDecoder()
	keyframe := slice.New[byte]()

	for c := 0; true; c++ {
		n, e := f.Read(buf)
		if n == 0 && errors.Is(e, io.EOF) {
			t.Log("reach end")
			break
		}
		_ = buff.Append(buf[:n])
		if !init {
			_, last_available_offset, e := flvDecoder.Init(buff.GetPureBuf())
			_ = buff.RemoveFront(last_available_offset)
			if e != nil {
				t.Fatal(e)
			}
			init = true
		} else {
			last_available_offset, e := flvDecoder.SearchStreamFrame(buff.GetPureBuf(), keyframe)
			if !keyframe.IsEmpty() {
				keyframe.Reset()
			}
			_ = buff.RemoveFront(last_available_offset)
			if e != nil {
				t.Fatal(e)
			}
		}
	}
}

// 10s-30s 184.852917ms
// 10m-10m20s 3.278605875s
func Test_FLVCut(t *testing.T) {
	{
		st := time.Now()
		defer func() {
			fmt.Println(time.Since(st))
		}()
	}
	cutf := file.Open("testdata/0.cut.flv")
	defer cutf.CloseErr()
	_ = cutf.Delete()

	f := file.Open("testdata/0.flv")
	defer f.CloseErr()

	if f.IsDir() || !f.IsExist() {
		t.Log("test file not exist")
	}

	e := NewFlvDecoder().Cut(f, time.Minute*10, time.Second*20, cutf.File(), false, false)
	t.Log(perrors.ErrorFormat(e))
}

func Test_FLVGenFastSeed(t *testing.T) {
	var VideoFastSeed = comp.GetV3[interface {
		InitGet(fastSeedFilePath string) (getIndex func(seedTo time.Duration) (int64, error), e error)
		InitSav(fastSeedFilePath string) (savIndex func(seedTo time.Duration, cuIndex int64) error, e error)
	}](`videoFastSeed`).Inter()

	f := file.Open("testdata/0.flv")
	defer f.CloseErr()
	sf, e := VideoFastSeed.InitSav("testdata/0.flv.fastSeed")
	if e != nil {
		t.Fatal(e)
	}

	if f.IsDir() || !f.IsExist() {
		t.Log("test file not exist")
	}

	e = NewFlvDecoder().GenFastSeed(f, func(seedTo time.Duration, cuIndex int64) error {
		return sf(seedTo, cuIndex)
	})
	if e != nil && !errors.Is(e, io.EOF) {
		t.Fatal(e)
	}
	// t.Log(perrors.ErrorFormat(e))
}

// 10s-30s 215.815423ms
// 10m-10m20s 174.918508ms
func Test_FLVCutSeed(t *testing.T) {
	{
		st := time.Now()
		defer func() {
			fmt.Println(time.Since(st))
		}()
	}
	cutf := file.Open("testdata/0.cut.flv")
	defer cutf.CloseErr()
	_ = cutf.Delete()

	f := file.Open("testdata/0.flv")
	defer f.CloseErr()

	if f.IsDir() || !f.IsExist() {
		t.Log("test file not exist")
	}

	var VideoFastSeed = comp.Get[interface {
		InitGet(fastSeedFilePath string) (getIndex func(seedTo time.Duration) (int64, error), e error)
		InitSav(fastSeedFilePath string) (savIndex func(seedTo time.Duration, cuIndex int64) error, e error)
	}](`videoFastSeed`)

	gf, e := VideoFastSeed.InitGet("testdata/0.flv.fastSeed")
	if e != nil {
		t.Fatal(e)
	}

	e = NewFlvDecoder().CutSeed(f, time.Minute*10, time.Second*20, cutf.File(), f, gf, false, false)
	t.Log(perrors.ErrorFormat(e))
}
