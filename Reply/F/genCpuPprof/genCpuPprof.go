package genCpuPprof

import (
	"bytes"
	"context"
	"runtime/pprof"
	"sync"
	"time"

	comp "github.com/qydysky/part/component2"
	pctx "github.com/qydysky/part/ctx"
	pfile "github.com/qydysky/part/file"
	pio "github.com/qydysky/part/io"
)

func init() {
	comp.RegisterOrPanic[interface {
		Start(ctx context.Context, file string) (any, error)
	}](`genCpuPprof`, &pp{})
}

type pp struct {
	once sync.Once
}

func (t *pp) Start(ctx context.Context, file string) (any, error) {
	if file == "" {
		return nil, nil
	}
	go t.once.Do(
		func() {
			ctx1, done1 := pctx.WaitCtx(ctx)
			defer done1()

			var buf bytes.Buffer
			var bufB = pio.RWC{
				R: buf.Read,
				W: buf.Write,
				C: func() error {
					buf.Reset()
					return nil
				},
			}

			for {
				_ = bufB.C()

				if err := pprof.StartCPUProfile(bufB); err != nil {
					return
				}
				select {
				case <-time.After(time.Minute):
					pprof.StopCPUProfile()
				case <-ctx1.Done():
					pprof.StopCPUProfile()
					return
				}
				pgo := pfile.Open(file)
				if pgo.IsExist() {
					_ = pgo.Delete()
				}
				pgo.Create()

				_ = pgo.CopyFromIoReader(bufB, pio.CopyConfig{})

				_ = pgo.Close()
			}
		})
	return nil, nil
}
