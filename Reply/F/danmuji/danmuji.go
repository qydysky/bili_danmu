package danmuji

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	comp "github.com/qydysky/part/component2"
	file "github.com/qydysky/part/file"
	limit "github.com/qydysky/part/limit"
)

var bbuf = make(map[string]string)

type i interface {
	Danmujif(s string, then func(string))
	Danmuji_auto(ctx context.Context, danmus []any, waitSec float64, then func(string))
}

func init() {
	f := file.New("config/config_auto_reply.json", 0, true)
	if !f.IsExist() {
		return
	}
	bb, err := f.ReadAll(100, 1<<16)
	if !errors.Is(err, io.EOF) {
		return
	}
	var buf map[string]interface{}
	_ = json.Unmarshal(bb, &buf)

	for k, v := range buf {
		if k == v {
			continue
		}
		bbuf[k] = v.(string)
	}

	if e := comp.Register[i]("danmuji", &Danmuji{
		reflect_limit: limit.New(1, "4s", "8s"),
	}); e != nil {
		panic(e)
	}
}

type Danmuji struct {
	reflect_limit *limit.Limit
}

func (t *Danmuji) Danmujif(s string, then func(string)) {
	if !t.reflect_limit.TO() {
		for k, v := range bbuf {
			if strings.Contains(s, k) {
				then(v)
				break
			}
		}
	}
}

func (t *Danmuji) Danmuji_auto(ctx context.Context, danmus []any, waitSec float64, then func(string)) {
	if waitSec < 5 {
		waitSec = 5
	}

	go func() {
		for i := 0; true; i++ {
			if i >= len(danmus) {
				i = 0
			}
			if msg, ok := danmus[i].(string); ok && msg != `` {
				then(msg)
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(waitSec) * time.Second):
			}
		}
	}()
}
