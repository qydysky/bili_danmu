package F

import "testing"

func TestBtoi32(t *testing.T) {
	t.Log(Btoi64([]byte{0xff, 0x0f, 0xff, 0xff}, 0))
}

func TestItob32(t *testing.T) {
	t.Log(Itob64(1131984000))
}
