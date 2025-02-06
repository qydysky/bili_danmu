package danmucoutpermin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	comp "github.com/qydysky/part/component2"
	file "github.com/qydysky/part/file"
	part "github.com/qydysky/part/io"
	msgq "github.com/qydysky/part/msgq"
	pweb "github.com/qydysky/part/web"
)

type TargetInterface interface {
	// will WriteHeader
	GetRec(savePath string, r *http.Request, w http.ResponseWriter) error
	Rec(ctx context.Context, roomid int, savePath string) func(map[string]any)
	Do(roomid int, msg string, uid string)
}

func init() {
	if e := comp.Register[TargetInterface]("danmuCountPerMin", &danmuCountPerMin{
		m: msgq.NewType[mi](),
	}); e != nil {
		panic(e)
	}
}

const filename = "danmuCountPerMin.json"

var noFoundModT, _ = time.Parse(time.DateTime, "2006-01-02 15:04:05")

type mi struct {
	msg string
	uid string
}

type danmuCountPerMin struct {
	m *msgq.MsgType[mi]
}

func (t *danmuCountPerMin) GetRec(savePath string, r *http.Request, w http.ResponseWriter) error {
	f := file.New(savePath+filename, 0, true)

	if f.IsDir() || !f.IsExist() {
		if pweb.NotModified(r, w, noFoundModT) {
			return nil
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("[]"))
		return os.ErrNotExist
	}

	if mod, e := f.GetFileModTimeT(); e != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return e
	} else if pweb.NotModified(r, w, mod) {
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	return f.CopyToIoWriter(w, part.CopyConfig{})
}

func (t *danmuCountPerMin) Rec(ctx context.Context, rid int, savePath string) func(map[string]any) {
	return func(cfg map[string]any) {
		cfgMsg := make(map[*regexp.Regexp]int)
		cfgUid := make(map[string]int)

		if m, ok := cfg["danmu"].(map[string]any); ok {
			for k, v := range m {
				if point, ok := v.(float64); ok && point != 0 {
					if reg, err := regexp.Compile(k); err == nil {
						cfgMsg[reg] = int(point)
					}
				}
			}
		}

		if m, ok := cfg["uid"].(map[string]any); ok {
			for k, v := range m {
				if point, ok := v.(float64); ok && point != 0 {
					cfgUid[k] = int(point)
				}
			}
		}

		if len(cfgMsg)+len(cfgUid) == 0 {
			return
		}

		go func() {
			var cpm []int
			var startT = time.Now()

			cancel := t.m.Pull_tag_only(fmt.Sprintf("do%d", rid), func(i mi) (disable bool) {
				point := 0
				for k, v := range cfgMsg {
					if k.MatchString(i.msg) {
						point += v
					}
				}
				for k, v := range cfgUid {
					if k == i.uid {
						point += v
					}
				}
				if point == 0 {
					return false
				}
				cu := int(time.Since(startT).Minutes())
				if len(cpm) < cu+1 {
					cpm = append(cpm, make([]int, cu+1-len(cpm))...)
				}
				cpm[cu] += point
				return false
			})

			<-ctx.Done()
			cancel()

			cu := int(time.Since(startT).Minutes())
			if len(cpm) < cu+1 {
				cpm = append(cpm, make([]int, cu+1-len(cpm))...)
			}

			if data, e := json.MarshalIndent(cpm, "", " "); e != nil {
				fmt.Println(e)
			} else {
				f := file.New(savePath+filename, 0, true)
				defer f.Close()
				_ = f.Delete()
				if _, e = f.Write(data, false); e != nil {
					fmt.Println(e)
				}
			}
		}()
	}
}

func (t *danmuCountPerMin) Do(roomid int, msg string, uid string) {
	t.m.Push_tag(fmt.Sprintf("do%d", roomid), mi{msg, uid})
}
