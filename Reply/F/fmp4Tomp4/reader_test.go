package fmp4Tomp4

import (
	"errors"
	"io"
	"testing"

	file "github.com/qydysky/part/file"
)

func Test_parse(t *testing.T) {
	var read = reader{
		f: file.New("/codefile/testdata/0.mp4", 0, false),
	}
	if e := read.Parse(); e != nil && !errors.Is(e, io.EOF) {
		t.Fatal(e)
	}
	t.Log(len(read.Boxs))

	m, e := read.ParseTrun()
	t.Log(m)
	t.Log(e)
	// for i := 0; i < len(read.boxs); i++ {
	// 	t.Log(read.boxs[i].name)
	// }
}
