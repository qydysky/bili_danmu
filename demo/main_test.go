package main

import (
	"context"
	"testing"
	"time"

	q "github.com/qydysky/bili_danmu"
)

// go test -run ^TestMain$ github.com/qydysky/bili_danmu/demo -race -count=1 -v -r xxx
func TestMain(m *testing.T) {
	ctx, c := context.WithTimeout(context.Background(), time.Second*40)
	defer c()
	q.Start(ctx)
}
