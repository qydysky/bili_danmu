package danmucoutpermin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	comp "github.com/qydysky/part/component2"
	file "github.com/qydysky/part/file"
	part "github.com/qydysky/part/io"
	msgq "github.com/qydysky/part/msgq"
)

type TargetInterface interface {
	GetRec(savePath string, w http.ResponseWriter) error
	Rec(ctx context.Context, roomid int, savePath string)
	Do(roomid int)
}

func init() {
	if e := comp.Register[TargetInterface]("danmuCountPerMin", &danmuCountPerMin{
		m: msgq.NewType[int](),
	}); e != nil {
		panic(e)
	}
}

const filename = "danmuCountPerMin.json"

type danmuCountPerMin struct {
	m *msgq.MsgType[int]
}

func (t *danmuCountPerMin) GetRec(savePath string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	f := file.New(savePath+filename, 0, true)
	if f.IsDir() || !f.IsExist() {
		_, _ = w.Write([]byte("[]"))
		return os.ErrNotExist
	}
	return file.New(savePath+filename, 0, true).CopyToIoWriter(w, part.CopyConfig{})
}

func (t *danmuCountPerMin) Rec(ctx context.Context, rid int, savePath string) {
	go func() {
		var cpm []int
		var startT = time.Now()

		cancel := t.m.Pull_tag_only(`do`, func(roomid int) (disable bool) {
			if rid == roomid {
				cu := int(time.Since(startT).Minutes())
				if len(cpm) < cu+1 {
					cpm = append(cpm, make([]int, cu+1-len(cpm))...)
				}
				cpm[cu] += 1
			}
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

func (t *danmuCountPerMin) Do(roomid int) {
	t.m.Push_tag(`do`, roomid)
}
