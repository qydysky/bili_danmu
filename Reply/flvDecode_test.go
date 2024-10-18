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
		keyframe := slice.New[byte]()
		front_buf, last_available_offset, e := flvDecoder.InitFlv(buff.GetPureBuf())
		if e != nil {
			t.Fatal(e)
		}
		last_available_offset, e = flvDecoder.SearchStreamTag(buff.GetPureBuf()[last_available_offset:], keyframe)
		if e != nil {
			t.Fatal(e)
		}
		_, _ = flog.Write([]byte(fmt.Sprintf("%d %d %d %d\n", c, len(front_buf), keyframe.Size(), last_available_offset)), true)
		t.Log(c, len(front_buf), keyframe.Size())
		_ = buff.RemoveFront(last_available_offset)
	}
	t.Log("max", humanize.Bytes(uint64(max)))
}
