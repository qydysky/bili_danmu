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
	flog := file.New("E:\\test\\0.flv.log", 0, false)
	flog.Delete()
	defer flog.Close()
	f := file.New("E:\\test\\0.flv", 0, false)
	defer f.Close()

	if f.IsDir() || !f.IsExist() {
		t.Fatal("file not support")
	}

	buf := make([]byte, humanize.MByte)
	buff := slice.New[byte](10 * humanize.MByte)
	max := 0

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
		front_buf, keyframe, last_available_offset, e := Search_stream_tag(buff.GetPureBuf())
		if e != nil {
			t.Fatal(e)
		}
		flog.Write([]byte(fmt.Sprintf("%d %d %d %d\n", c, len(front_buf), len(keyframe), last_available_offset)), true)
		t.Log(c, len(front_buf), len(keyframe))
		buff.RemoveFront(last_available_offset)
	}
	t.Log("max", humanize.Bytes(uint64(max)))
}