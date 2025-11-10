package ass

import (
	"testing"

	comp "github.com/qydysky/part/component2"
)

func TestMain(t *testing.T) {
	var ass = comp.GetV3[interface {
		ToAss(savePath string, filename ...string)
		Init(cfg any)
	}](`ass`)
	ass.Inter().ToAss("./testdata/", "1.ass")
}

func TestStos(t *testing.T) {
	t.Log(stos(3661))
}
