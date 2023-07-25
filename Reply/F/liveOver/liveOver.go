package liveOver

import (
	"context"
	"fmt"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	comp "github.com/qydysky/part/component"
)

func init() {
	if e := comp.Put[c.Common](`bili_danmu.Reply.wsmsg.preparing.sumup`, sumup); e != nil {
		panic(e)
	}
}

func sumup(ctx context.Context, ptr *c.Common) error {
	dura := time.Since(ptr.Live_Start_Time).Round(time.Second)
	if ptr.Live_Start_Time.IsZero() {
		dura = 0
	}
	ptr.Log.Base(`功能`, `下播总结`).L(`I: `, fmt.Sprintf("%d 时长 %s 总营收 ￥%.2f 观看人数 %d", ptr.Roomid, dura, ptr.Rev, ptr.Watched))
	return nil
}
