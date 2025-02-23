package ass

import (
	"testing"

	comp "github.com/qydysky/part/component2"
)

func TestMain(t *testing.T) {
	var ass = comp.Get[interface {
		ToAss(savePath string)
		Init(cfg any)
	}](`ass`)
	ass.ToAss("./testdata/")
}

func TestStos(t *testing.T) {
	t.Log(stos(3661))
}
