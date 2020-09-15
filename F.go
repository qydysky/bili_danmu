package bili_danmu

import (
	"bytes"

	p "github.com/qydysky/part"
) 

type Autoban struct {
	buf []byte
	Inuse bool
}

var autoban = Autoban {
	Inuse:true,
}

func Autobanf(s string) float32 {
	if autoban.Inuse {return 0}

	if len(autoban.buf) == 0 {
		f := p.File().FileWR(p.Filel{
			File:"Autoban.txt",
			Write:false,
		})
		autoban.buf = []byte(f)
	}

	var scop int
	for _, v := range []byte(s) {
		if bytes.Contains(autoban.buf, []byte{v}) {scop += 1}
	}
	return float32(scop) / float32(len(s))
}

func Autoban_add(s string) {
	if autoban.Inuse {return}

	autoban.buf = append(autoban.buf, []byte(s)...)
	p.File().FileWR(p.Filel{
		File:"Autoban.txt",
		Write:true,
		Loc:-1,
		Context:[]interface{}{s},
	})
}