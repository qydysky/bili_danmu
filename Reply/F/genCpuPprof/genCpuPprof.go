package genCpuPprof

import (
	"context"
	"runtime/pprof"

	comp "github.com/qydysky/part/component"
	pctx "github.com/qydysky/part/ctx"
	pfile "github.com/qydysky/part/file"
)

var Start = comp.NewComp(start)

func start(ctx context.Context, file string) error {
	if file == "" {
		return nil
	}
	pgo := pfile.New(file, 0, false)
	if pgo.IsExist() {
		_ = pgo.Delete()
	}
	pgo.Create()
	if err := pprof.StartCPUProfile(pgo.File()); err != nil {
		return err
	}
	go func() {
		ctx1, done1 := pctx.WaitCtx(ctx)
		defer done1()
		<-ctx1.Done()
		pprof.StopCPUProfile()
		pgo.Close()
	}()
	return nil
}
