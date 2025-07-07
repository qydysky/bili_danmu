package danmucoutpermin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"sync"
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
	CheckRoot(dir string)
	Rec(ctx context.Context, roomid int, savePath string) func(any)
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
	roomid int
	msg    string
	uid    string
}

type danmuCountPerMin struct {
	root string
	m    *msgq.MsgType[mi]
}

func (t *danmuCountPerMin) CheckRoot(dir string) {
	t.root = dir
}

func (t *danmuCountPerMin) GetRec(savePath string, r *http.Request, w http.ResponseWriter) error {
	f := file.New(savePath+filename, 0, true).CheckRoot(t.root)
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

func (t *danmuCountPerMin) Rec(ctx context.Context, rid int, savePath string) func(any) {
	return func(cfg any) {
		cfgM := make(map[int][]func(mi) int)
		// 默认统计弹幕数
		cfgM[-1] = []func(mi) int{
			func(m mi) int {
				return 1
			},
		}

		if cfgs, ok := cfg.([]any); ok {
			// 配置特定房间
			for i := 0; i < len(cfgs); i++ {
				cfg, ok := cfgs[i].(map[string]any)
				if !ok {
					continue
				}

				roomid := 0
				if m, ok := cfg["roomid"].(float64); !ok || roomid == 0 {
					continue
				} else {
					roomid = int(m)
				}

				var funcs []func(mi) int

				if m, ok := cfg["danmu"].(map[string]any); ok {
					for k, v := range m {
						if point, ok := v.(float64); ok && point != 0 {
							if reg, err := regexp.Compile(k); err == nil {
								funcs = append(funcs, func(m mi) int {
									if reg.MatchString(m.msg) {
										return int(point)
									}
									return 0
								})
							}
						}
					}
				}

				if m, ok := cfg["uid"].(map[string]any); ok {
					for k, v := range m {
						if point, ok := v.(float64); ok && point != 0 {
							funcs = append(funcs, func(m mi) int {
								if k == m.uid {
									return int(point)
								}
								return 0
							})
						}
					}
				}

				if len(funcs) == 0 {
					continue
				}

				cfgM[roomid] = funcs

				fmt.Println(roomid, len(funcs))
			}
		}
		go func() {
			var cpm []int
			var cpmLock sync.Mutex
			var startT = time.Now()

			cancel := t.m.Pull_tag_only(`do`, func(m mi) (disable bool) {
				point := 0
				if fs, ok := cfgM[m.roomid]; ok {
					for i := 0; i < len(fs); i++ {
						point += fs[i](m)
					}
				} else if fs, ok := cfgM[-1]; ok {
					for i := 0; i < len(fs); i++ {
						point += fs[i](m)
					}
				}
				cu := int(time.Since(startT).Minutes())
				cpmLock.Lock()
				if len(cpm) < cu+1 {
					cpm = append(cpm, make([]int, cu+1-len(cpm))...)
				}
				cpm[cu] += point
				cpmLock.Unlock()
				return false
			})

			<-ctx.Done()
			cancel()

			cu := int(time.Since(startT).Minutes())
			cpmLock.Lock()
			defer cpmLock.Unlock()
			if len(cpm) < cu+1 {
				cpm = append(cpm, make([]int, cu+1-len(cpm))...)
			}

			if data, e := json.MarshalIndent(cpm, "", " "); e != nil {
				fmt.Println(e)
			} else {
				f := file.New(savePath+filename, 0, true)
				defer f.Close()
				_ = f.Delete()
				if _, e = f.WriteRaw(data, false); e != nil {
					fmt.Println(e)
				}
			}
		}()
	}
}

func (t *danmuCountPerMin) Do(roomid int, msg string, uid string) {
	t.m.Push_tag(`do`, mi{roomid, msg, uid})
}
