package shortdanmu

import (
	"sync"

	comp "github.com/qydysky/part/component2"
)

func init() {
	comp.RegisterOrPanic[interface {
		Deal(string) string
	}](`shortDanmu`, new(shortdanmu))
}

type shortdanmu struct {
	lastdanmu []rune
	l         sync.Mutex
}

func (t *shortdanmu) Deal(s string) string {
	t.l.Lock()
	defer t.l.Unlock()

	if len(t.lastdanmu) == 0 {
		t.lastdanmu = []rune(s)
		return s
	}

	var new string

	for k, v := range []rune(s) {
		if k >= len(t.lastdanmu) {
			new += string([]rune(s)[k:])
			break
		}
		if v != t.lastdanmu[k] {
			switch k {
			case 0, 1, 2:
				new = s
			default:
				new = "..." + string([]rune(s)[k-1:])
			}
			break
		}
	}
	// if new == "" {new = "...."}
	t.lastdanmu = []rune(s)
	return new
}
