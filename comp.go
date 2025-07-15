package bili_danmu

import (
	_ "github.com/qydysky/bili_danmu/cmd" //removable
	comp "github.com/qydysky/part/component2"
)

type CmdI interface {
	Cmd()
}

var Cmd = comp.GetV2(`cmd`, comp.PreFuncErr[CmdI]{})
