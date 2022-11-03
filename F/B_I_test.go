package F

import "testing"

func TestBtoi32(t *testing.T) {
	t.Log(int32(Btoui32([]byte{0xff, 0x0f, 0xff, 0xff}, 0)))
	t.Log(Btoi32([]byte{0xff, 0xff, 0xff, 0xff}, 0))
	t.Log(Btoi32([]byte{0x00, 0xff, 0xff, 0xff}, 0))
}

func TestItob32(t *testing.T) {
	t.Log(Itob64(1131984000))
}
