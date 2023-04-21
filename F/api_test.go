package F

import "testing"

func Test_getWridWts(t *testing.T) {
	w_rid, _ := new(GetFunc).getWridWts(
		"mid=11280430&token=&platform=web&web_location=1550101",
		"https://i0.hdslb.com/bfs/wbi/e1be084baf3b4663b2465fca5bf1d889.png",
		"https://i0.hdslb.com/bfs/wbi/0ae7d656b8114fe1901717dd092b7ee9.png",
		"1682105747",
	)
	if w_rid != "766c2091b69102edb391bf16ef917d6c" {
		t.Fatal()
	}
}
