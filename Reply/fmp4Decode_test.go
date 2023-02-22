package reply

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/dustin/go-humanize"
	file "github.com/qydysky/part/file"
	slice "github.com/qydysky/part/slice"
)

func Test_deal(t *testing.T) {
	flog := file.New("E:\\test\\0.flv.log", 0, false)
	flog.Delete()
	defer flog.Close()
	f := file.New("E:\\test\\0.mp4", 0, false)
	defer f.Close()

	if f.IsDir() || !f.IsExist() {
		t.Fatal("file not support")
	}

	buf := make([]byte, humanize.MByte)
	buff := slice.New[byte]()
	max := 0
	fmp4Decoder := new(Fmp4Decoder)
	fmp4KeyFrames := slice.New[byte]()

	for c := 0; true; c++ {
		n, e := f.Read(buf)
		if n == 0 && errors.Is(e, io.EOF) {
			t.Log("reach end")
			break
		}
		buff.Append(buf[:n])
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
		last_available_offset, e := fmp4Decoder.Search_stream_fmp4(buff.GetPureBuf(), fmp4KeyFrames)
		if e != nil && e.Error() != "未初始化traks" {
			t.Fatal(e)
		}
		if len(front_buf) != 0 {
			t.Log("front_buf")
			break
		}
		flog.Write([]byte(fmt.Sprintf("%d %d\n", c, len(front_buf))), true)
		t.Log(c, len(front_buf))
		buff.RemoveFront(last_available_offset)
	}
	t.Log("max", humanize.Bytes(uint64(max)))
}
