package F

import (
	"bytes"
	"testing"
)

// 143.9 ns/op            16 B/op          1 allocs/op
func Benchmark(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Btoi([]byte{1, 2, 3, 4, 5, 5}, 0, 4)
	}
}

func TestItob32(t *testing.T) {
	t.Log(Btoi([]byte{1, 2, 3, 4, 5}, 0, 4))
	t.Log(Btoiv2([]byte{1, 2, 3, 4, 5}, 0, 4))
	t.Log(Btoi16v2([]byte{1, 2}, 0))
	t.Log(Btoi16([]byte{1, 2}, 0))
}

func Test1(t *testing.T) {
	if Btoi([]byte{1, 2, 3, 4, 5}, 0, 4) != int64(Btoi32([]byte{1, 2, 3, 4, 5}, 0)) {
		t.Fatal()
	}
	if !bytes.Equal(Itob64(Btoi([]byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, 0, 8)), []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}) {
		t.Fatal()
	}
}
