package streamPredeal

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	comp "github.com/qydysky/part/component2"
)

func init() {
	comp.RegisterOrPanic[interface {
		Init(any)
		Deal(mode, streamType string, ctx context.Context, w http.ResponseWriter) (err error, cmdI io.WriteCloser)
	}](`streamPredeal`, new(a))
}

type a struct {
	config map[string]any
}

func (t *a) Init(c any) {
	if modes, ok := c.(map[string]any); ok {
		t.config = modes
	}
}

func (t *a) Deal(mode, streamType string, ctx context.Context, w http.ResponseWriter) (err error, cmdI io.WriteCloser) {
	if mode, ok := t.config[mode]; ok {
		if cmd, ok := mode.(map[string]any); ok {
			if args, ok := cmd[streamType]; ok {
				if tmp, ok := args.([]any); ok {
					var arg []string
					for i := range tmp {
						arg = append(arg, tmp[i].(string))
					}
					cmd := exec.CommandContext(ctx, arg[0], arg[1:]...)
					cmd.Stderr = os.Stdout
					cmd.Stdout = w
					cmdI, err = cmd.StdinPipe()
					if err != nil {
						return
					}
					err = cmd.Start()
					if err != nil {
						return
					}
					go func() {
						fmt.Println(cmd.Wait())
					}()
				}
			}
		}
	}
	return
}
