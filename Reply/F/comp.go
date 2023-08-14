package f

import (
	"github.com/qydysky/bili_danmu/Reply/F/danmuXml"
	"github.com/qydysky/bili_danmu/Reply/F/liveOver"
	comp "github.com/qydysky/part/component"
)

func init() {
	var linkMap = map[string][]string{
		"github.com/qydysky/bili_danmu/Reply.startRecDanmu.stop": {
			comp.Sign[danmuXml.Sign](),
		},
		"github.com/qydysky/bili_danmu/Reply.preparing": {
			comp.Sign[liveOver.Sign](),
		},
	}
	if e := comp.Link(linkMap); e != nil {
		panic(e)
	}
}
