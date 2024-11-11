package reply

import (
	"errors"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
	perrors "github.com/qydysky/part/errors"
	file "github.com/qydysky/part/file"
	slice "github.com/qydysky/part/slice"
)

func Test_FLVdeal(t *testing.T) {
	flog := file.New("0.flv.log", 0, false)
	_ = flog.Delete()
	defer flog.Close()
	f := file.New("testdata/0.flv", 0, false)
	defer f.Close()

	if f.IsDir() || !f.IsExist() {
		t.Fatal("file not support")
	}

	buf := make([]byte, humanize.MByte)
	buff := slice.New[byte](10 * humanize.MByte)
	max := 0
	flvDecoder := NewFlvDecoder()
	keyframe := slice.New[byte]()

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
		front_buf, last_available_offset, e := flvDecoder.Parse(buff.GetPureBuf(), keyframe)
		if e != nil {
			t.Fatal(e)
		}
		_, _ = flog.Write([]byte(fmt.Sprintf("%d %d %d %d\n", c, len(front_buf), keyframe.Size(), last_available_offset)), true)
		t.Log(c, len(front_buf), keyframe.Size())
		_ = buff.RemoveFront(last_available_offset)
	}
	t.Log("max", humanize.Bytes(uint64(max)))
}

func _Test_FLVCut(t *testing.T) {

	cutf := file.New("testdata/1.cut.flv", 0, false)
	defer cutf.Close()
	cutf.Delete()

	f := file.New("testdata/1.flv", 0, false)
	defer f.Close()

	if f.IsDir() || !f.IsExist() {
		t.Fatal("file not support")
	}

	e := NewFlvDecoder().Cut(f, time.Second*10, time.Second*20, cutf.File())
	if perrors.Catch(e, "Read") {
		t.Log("err Read", e)
	}
	if perrors.Catch(e, "InitFlv") {
		t.Log("err InitFlv", e)
	}
	if perrors.Catch(e, "skip") {
		t.Log("err skip", e)
	}
	if perrors.Catch(e, "cutW") {
		t.Log("err cutW", e)
	}
	t.Log(e)
}
