package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
	q "github.com/qydysky/bili_danmu"
	file "github.com/qydysky/part/file"
)

// go test -run ^TestMain$ github.com/qydysky/bili_danmu/demo -race -count=1 -v -r xxx
func TestMain(m *testing.T) {
	fl := file.Open("danmu.log")
	_ = fl.Delete()

	ctx, c := context.WithTimeout(context.Background(), time.Second*40)
	q.Start(ctx)
	c()

	var line = []byte{}
	for {
		if e := fl.ReadUntilV2(&line, []byte{'\n'}, humanize.KByte, humanize.MByte); e != nil {
			if !errors.Is(e, io.EOF) {
				m.Fatal(e)
			} else {
				return
			}
		} else if bytes.HasPrefix(line, []byte("E: ")) {
			m.Fatal(string(line))
		}
	}
}
