package liveOver

import (
	"context"
	"fmt"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	comp "github.com/qydysky/part/component"
)

// *c.Common
var Sumup = comp.NewComp(sumup)

func sumup(ctx context.Context, ptr *c.Common) (any, error) {
	dura := time.Since(ptr.Live_Start_Time).Round(time.Second)
	if ptr.Live_Start_Time.IsZero() {
		ptr.Log.Base(`功能`, `下播总结`).I(fmt.Sprintf("%d 未直播", ptr.Roomid))
	} else {
		var pperm = float64(ptr.Watched) / float64(dura/time.Minute)
		var yperm = float64(ptr.Rev) / float64(dura/time.Minute)
		// 若是中途录制，则使用启动时间
		if ptr.StartT.After(ptr.Live_Start_Time) {
			yperm = float64(ptr.Rev) / float64(time.Since(ptr.StartT).Round(time.Second)/time.Minute)
		}
		ptr.Log.Base(`功能`, `下播总结`).I(fmt.Sprintf("%d 时长 %s 营收 %.2f元 %.2f元/分 人数 %d人 %.2f人/分", ptr.Roomid, dura, ptr.Rev, yperm, ptr.Watched, pperm))

	}
	return nil, nil
}
