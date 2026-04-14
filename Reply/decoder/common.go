package decoder

import (
	pe "github.com/qydysky/part/errors/v2"

	_ "github.com/qydysky/bili_danmu/Reply/F/videoFastSeed" //removable
)

var ActDecoder = pe.Action[struct {
	Decode      pe.Error
	BufOverflow pe.Error
}](`ActDecoder`)
