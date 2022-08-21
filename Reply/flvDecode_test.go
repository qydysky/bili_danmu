package reply

import (
	"bytes"
	"testing"

	p "github.com/qydysky/part"
)

func Test_flv(t *testing.T) {
	f := p.File().FileWR(p.Filel{
		File: "1.flv",
	})
	Seach_stream_tag(bytes.NewBufferString(f).Bytes())
}
