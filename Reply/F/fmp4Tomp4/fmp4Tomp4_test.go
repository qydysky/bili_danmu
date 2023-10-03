package fmp4Tomp4

import (
	"context"
	"testing"
)

func Test_conver(t *testing.T) {
	path := "/codefile/testdata/"
	if e := conver(context.Background(), &path); e != nil {
		t.Fatal(e)
	}
}
