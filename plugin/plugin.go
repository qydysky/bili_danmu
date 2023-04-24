package plugin

import (
	msgq "github.com/qydysky/part/msgq"
)

// Event
const (
	LoadKv = iota
)

var Plugin msgq.Msgq

type Danmu struct {
	Msg    string
	Color  string
	Border bool
	Mode   int
	Auth   any
	Uid    string
	Roomid int
}
