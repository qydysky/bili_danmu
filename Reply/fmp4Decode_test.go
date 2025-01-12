package reply

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

func Test_deal(t *testing.T) {
	flog := file.New("0.mp4.log", 0, false)
	_ = flog.Delete()
	defer flog.Close()
	f := file.New("testdata/0.mp4", 0, false)
	defer f.Close()

	if f.IsDir() || !f.IsExist() {
		t.Fatal("file not support")
	}

	buf := make([]byte, humanize.MByte)
	buff := slice.New[byte]()
	max := 0
	fmp4Decoder := NewFmp4Decoder()

	for c := 0; true; c++ {
		n, e := f.Read(buf)
		if n == 0 && errors.Is(e, io.EOF) {
			t.Log("reach end")
			break
		}
		_ = buff.Append(buf[:n])
		if s := buff.Size(); max < s {
			max = s
		}
		if max > humanize.MByte*100 {
			t.Log("reach max")
			break
		}

		front_buf, e := fmp4Decoder.Init_fmp4(buff.GetCopyBuf())
		if e != nil {
			t.Fatal(e)
		}
		last_available_offset, e := fmp4Decoder.Search_stream_fmp4(buff.GetPureBuf(), slice.New[byte]())
		if e != nil && e.Error() != "未初始化traks" {
			t.Fatal(e)
		}
		if len(front_buf) != 0 {
			t.Log("front_buf")
			break
		}
		_, _ = flog.Write([]byte(fmt.Sprintf("%d %d\n", c, len(front_buf))), true)
		t.Log(c, len(front_buf))
		_ = buff.RemoveFront(last_available_offset)
	}
	t.Log("max", humanize.Bytes(uint64(max)))
}

// 10s-30s 110.962896ms
// 10m-10m20s 1.955749395s
// 30m-30m20s 5.791614855s
func Test_Mp4Cut(t *testing.T) {
	{
		st := time.Now()
		defer func() {
			fmt.Println(time.Since(st))
		}()
	}

	cutf := file.New("testdata/0.cut.mp4", 0, false)
	defer cutf.Close()
	_ = cutf.Delete()

	f := file.New("testdata/0.mp4", 0, false)
	defer f.Close()

	if f.IsDir() || !f.IsExist() {
		t.Log("test file not exist")
	}

	e := NewFmp4Decoder().Cut(f, time.Minute*30, time.Second*20, cutf.File())
	if perrors.Catch(e, "Read") {
		t.Log("err Read", e)
	}
	if perrors.Catch(e, "Init_fmp4") {
		t.Log("err Init_fmp4", e)
	}
	if perrors.Catch(e, "skip") {
		t.Log("err skip", e)
	}
	if perrors.Catch(e, "cutW") {
		t.Log("err cutW", e)
	}
	t.Log(e)
}

func Test_Mp4GenFastSeed(t *testing.T) {
	var VideoFastSeed = comp.Get[interface {
		InitGet(fastSeedFilePath string) (getIndex func(seedTo time.Duration) (int64, error), e error)
		InitSav(fastSeedFilePath string) (savIndex func(seedTo time.Duration, cuIndex int64) error, e error)
	}](`videoFastSeed`)

	f := file.New("testdata/0.mp4", 0, false)
	defer f.Close()
	sf, e := VideoFastSeed.InitSav("testdata/0.fastSeed")
	if e != nil {
		t.Fatal(e)
	}

	e = NewFmp4Decoder().GenFastSeed(f, func(seedTo time.Duration, cuIndex int64) error {
		return sf(seedTo, cuIndex)
	})
	if perrors.Catch(e, "Read") {
		t.Log("err Read", e)
	}
	if perrors.Catch(e, "Init_fmp4") {
		t.Log("err Init_fmp4", e)
	}
	if perrors.Catch(e, "skip") {
		t.Log("err skip", e)
	}
	if perrors.Catch(e, "cutW") {
		t.Log("err cutW", e)
	}
	t.Log(e)

	// VideoFastSeed.BeforeGet("testdata/1.fastSeed")
	// {
	// 	index, _ := VideoFastSeed.GetIndex(0)
	// 	f.SeekIndex(index, file.AtOrigin)
	// 	buf := make([]byte, 8)
	// 	f.Read(buf)
	// 	fmt.Println(string(buf[4:8]))
	// }
}

// 10s-30s 90.05983ms
// 10m-10m20s 88.769475ms
// 30m-30m20s 104.381225ms
func Test_Mp4CutSeed(t *testing.T) {
	{
		st := time.Now()
		defer func() {
			fmt.Println(time.Since(st))
		}()
	}

	cutf := file.New("testdata/0.cut.mp4", 0, false)
	defer cutf.Close()
	_ = cutf.Delete()

	f := file.New("testdata/0.mp4", 0, false)
	defer f.Close()

	if f.IsDir() || !f.IsExist() {
		t.Log("test file not exist")
	}

	var VideoFastSeed = comp.Get[interface {
		InitGet(fastSeedFilePath string) (getIndex func(seedTo time.Duration) (int64, error), e error)
		InitSav(fastSeedFilePath string) (savIndex func(seedTo time.Duration, cuIndex int64) error, e error)
	}](`videoFastSeed`)

	gf, e := VideoFastSeed.InitGet("testdata/0.fastSeed")
	if e != nil {
		t.Fatal(e)
	}

	e = NewFmp4Decoder().CutSeed(f, time.Minute*30, time.Second*20, cutf.File(), f, gf)
	if perrors.Catch(e, "Read") {
		t.Log("err Read", e)
	}
	if perrors.Catch(e, "Init_fmp4") {
		t.Log("err Init_fmp4", e)
	}
	if perrors.Catch(e, "skip") {
		t.Log("err skip", e)
	}
	if perrors.Catch(e, "cutW") {
		t.Log("err cutW", e)
	}
	t.Log(e)
}
