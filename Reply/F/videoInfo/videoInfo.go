package videoInfo

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	component "github.com/qydysky/part/component"
	pctx "github.com/qydysky/part/ctx"
	file "github.com/qydysky/part/file"
)

var (
	Save = component.NewComp(save)
	Get  = component.NewComp(get)
)

type Info interface {
	Common() *c.Common
	GetStreamType() string
	GetSavePath() string
}

type Paf struct {
	Uname           string        `json:"uname"`
	UpUid           int           `json:"upUid"`
	Roomid          int           `json:"roomid"`
	Qn              string        `json:"qn"`
	Name            string        `json:"name"`
	StartT          string        `json:"startT"`
	StartTS         int64         `json:"-"`
	EndT            string        `json:"endT"`
	EndTS           int64         `json:"-"`
	Dur             time.Duration `json:"-"`
	Path            string        `json:"path"`
	CurrentSavePath string        `json:"-"`
	Format          string        `json:"format"`
	StartLiveT      string        `json:"startLiveT"`
	OnlinesPerMin   []int         `json:"onlinesPerMin"`
}

func save(ctx context.Context, i Info) (*Paf, error) {
	infop := _newPaf(i.Common(), i.GetSavePath(), i.GetStreamType())
	if e := _save(infop); e != nil {
		return nil, e
	}

	go func(ctx context.Context, info Info, infop *Paf) {
		logger := info.Common().Log.Base_add("videoInfo")
		ctx1, done1 := pctx.WaitCtx(ctx)
		defer done1()
		for {
			select {
			case <-ctx1.Done():
				infop.EndT = time.Now().Format(time.DateTime)
				if e := _save(infop); e != nil {
					logger.L(`E: `, e)
					return
				}
				return
			case <-time.After(time.Minute):
				infop.OnlinesPerMin = append(infop.OnlinesPerMin, info.Common().OnlineNum)

				if infop.Name != info.Common().Title {
					infop.Name = info.Common().Title
					if e := _save(infop); e != nil {
						logger.L(`E: `, e)
						return
					}
				}
			}
		}
	}(ctx, i, infop)

	return nil, nil
}

func get(ctx context.Context, savepath string) (*Paf, error) {
	var d Paf
	dirf := file.New(savepath, 0, true)
	defer dirf.CloseErr()
	if dirf.IsDir() {
		// 从文件夹获取信息
		// {
		// 	dirfName := path.Base(dirf.File().Name())
		// 	if len(dirfName) > 20 {
		// 		d = Paf{Name: dirfName[20:], StartT: dirfName[:19], Path: dirfName}
		// 	}
		// 	mp4f := file.New(savepath+string(os.PathSeparator)+"0.mp4", 0, true)
		// 	if mp4f.IsExist() {
		// 		d.Format = "mp4"
		// 	} else {
		// 		d.Format = "flv"
		// 	}
		// }
		// 从0.json获取信息
		{
			json0 := file.New(savepath+string(os.PathSeparator)+"0.json", 0, true).CheckRoot(savepath)
			if !json0.IsExist() {
				return &d, os.ErrNotExist
			}
			defer json0.CloseErr()
			if data, e := json0.ReadAll(1<<8, 1<<16); e != nil && !errors.Is(e, io.EOF) {
				return &d, e
			} else {
				if e := json.Unmarshal(data, &d); e != nil {
					return &d, e
				}
				if t, e := time.Parse("2006_01_02-15_04_05", d.StartT); e == nil {
					d.StartT = t.Format(time.DateTime)
				}
				st, se := time.Parse(time.DateTime, d.StartT)
				if se == nil {
					d.StartTS = st.Unix()
				}
				et, ee := time.Parse(time.DateTime, d.EndT)
				if ee == nil {
					d.EndTS = et.Unix()
				}
				if se == nil && ee == nil {
					d.Dur = et.Sub(st)
				}
				d.CurrentSavePath = d.Path + "/0." + d.Format
			}
		}
	}
	return &d, nil
}

func _save(pathInfo *Paf) error {
	fj := file.New(pathInfo.CurrentSavePath+"0.json", 0, true)
	if fj.IsExist() {
		if err := fj.Delete(); err != nil {
			return err
		}
	}
	if pathInfoJson, err := json.Marshal(pathInfo); err != nil {
		return err
	} else if _, err := fj.Write(pathInfoJson); err != nil {
		return err
	}
	return nil
}

func _newPaf(common *c.Common, savePath, streamType string) *Paf {
	return &Paf{
		Uname:           common.Uname,
		UpUid:           common.UpUid,
		Roomid:          common.Roomid,
		Qn:              c.C.Qn[common.Live_qn],
		Name:            common.Title,
		StartT:          time.Now().Format(time.DateTime),
		Path:            path.Base(savePath),
		CurrentSavePath: savePath,
		Format:          streamType,
		OnlinesPerMin:   []int{common.OnlineNum},
		StartLiveT:      common.Live_Start_Time.Format(time.DateTime),
	}
}
