package F

import "testing"

func TestBtoi32(t *testing.T) {
	t.Log(Btoi64([]byte{0xd6, 0xfd, 0x62, 0x50}, 0))
}

func TestItob32(t *testing.T) {
	t.Log(Itob64(1131984000))
}
