package danmuemotes

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"iter"
	"regexp"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	c "github.com/qydysky/bili_danmu/CV"
	comp "github.com/qydysky/part/component2"
	pe "github.com/qydysky/part/errors"
	file "github.com/qydysky/part/file"
	phash "github.com/qydysky/part/hash"
	log "github.com/qydysky/part/log/v2"
	reqf "github.com/qydysky/part/reqf"
	unsafe "github.com/qydysky/part/unsafe"
)

type TargetInterface interface {
	SaveEmote(ctx context.Context, ptr struct {
		Logg *log.Log
		Info []any
		Msg  *string
	}) (ret any, err error)
	Hashr(s string) (r string)
	SetLayerN(n int)
	IsErrNoEmote(e error) bool
	PackEmotes(dir string) error
	GetEmotesDir(dir string) interface {
		fs.FS
		io.Closer
	}
}

func init() {
	i := danmuEmotes{
		Dir:    "emots/",
		LayerN: 0,
	}
	if e := comp.Register[TargetInterface]("danmuEmotes", &i); e != nil {
		panic(e)
	}
	_, _ = file.New(i.Dir+"README.md", 0, true).WriteRaw([]byte(""), false)
}

var errNoEmote = errors.New("errNoEmote")

type danmuEmotes struct {
	Dir    string
	LayerN int
}

func (t *danmuEmotes) IsErrNoEmote(e error) bool {
	return errors.Is(e, errNoEmote)
}

func (t *danmuEmotes) SetLayerN(n int) {
	t.LayerN = n
}

func (t *danmuEmotes) SaveEmote(ctx context.Context, ptr struct {
	Logg *log.Log
	Info []any
	Msg  *string
}) (ret any, err error) {
	isEmote := false
	if m, ok := ptr.Info[13].(map[string]any); ok {
		if url, ok := m[`url`].(string); ok {
			if !strings.Contains(*ptr.Msg, "[") {
				if emoticon_unique, ok := m[`emoticon_unique`].(string); ok {
					*ptr.Msg = "[" + *ptr.Msg + emoticon_unique + "]"
				}
			}
			isEmote = true
			savePath := t.Dir + t.Hashr(*ptr.Msg) + ".png"
			if !file.IsExist(savePath) {
				go func() {
					req := c.C.ReqPool.Get()
					defer c.C.ReqPool.Put(req)
					if e := req.Reqf(reqf.Rval{
						SaveToPath: savePath,
						Url:        url,
						Proxy:      c.C.Proxy,
						Timeout:    5000,
						Header: map[string]string{
							`User-Agent`:      c.UA,
							`Accept`:          `*/*`,
							`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
							`Origin`:          `https://live.bilibili.com`,
							`Connection`:      `keep-alive`,
							`Pragma`:          `no-cache`,
							`Cache-Control`:   `no-cache`,
							`Referer`:         "https://live.bilibili.com/",
						},
					}); e != nil {
						ptr.Logg.E(`表情下载失败`, pe.ErrorFormat(e, pe.ErrActionInLineFunc))
					}
				}()
			}
		}
	} else if m, ok := ptr.Info[15].(map[string]any); ok {
		if extra, ok := m[`extra`]; ok {
			if extrab, ok := extra.(string); ok {
				var E struct {
					Emots map[string]any `json:"emots"`
				}
				if e := json.Unmarshal(unsafe.S2B(extrab), &E); e != nil {
					return nil, e
				} else {
					for k, v := range E.Emots {
						m, ok := v.(map[string]any)
						if !ok {
							continue
						}

						url, ok := m[`url`].(string)
						if !ok {
							continue
						}

						isEmote = true
						savePath := t.Dir + t.Hashr(k) + ".png"
						if file.IsExist(savePath) {
							continue
						}
						go func() {
							req := c.C.ReqPool.Get()
							defer c.C.ReqPool.Put(req)
							if e := req.Reqf(reqf.Rval{
								SaveToPath: savePath,
								Url:        url,
								Proxy:      c.C.Proxy,
								Timeout:    5000,
								Header: map[string]string{
									`User-Agent`:      c.UA,
									`Accept`:          `*/*`,
									`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
									`Origin`:          `https://live.bilibili.com`,
									`Connection`:      `keep-alive`,
									`Pragma`:          `no-cache`,
									`Cache-Control`:   `no-cache`,
									`Referer`:         "https://live.bilibili.com/",
								},
							}); e != nil {
								ptr.Logg.E(e)
							}
						}()
					}
				}
			}
		}
	}
	if !isEmote {
		err = errNoEmote
	}
	return
}

func (t *danmuEmotes) Hashr(s string) (r string) {
	rs := phash.Md5String(s)
	if t.LayerN <= 0 {
		return rs
	}
	rr := []byte{}
	layer := t.LayerN
	if layer > len(rs)-1 {
		layer = layer - 1
	}
	for i := 0; i < len(rs); i++ {
		rr = append(rr, rs[i])
		if layer > 0 {
			rr = append(rr, '/')
			layer -= 1
		}
	}
	return unsafe.B2S(rr)
}

func (t *danmuEmotes) PackEmotes(dir string) error {
	var (
		w    *zip.Writer
		set  = make(map[string]struct{})
		r, _ = regexp.Compile(`\[.*?\]`)
	)

	for line := range loadCsv(dir, "0.csv") {
		for _, v := range r.FindAllString(line.Text, -1) {
			key := t.Hashr(v)
			if _, ok := set[key]; ok {
				continue
			} else {
				set[key] = struct{}{}
			}

			f := file.Open(t.Dir + key + ".png")
			if f.IsExist() {
				if w == nil {
					f := file.Open(dir + "emotes.zip")
					if f.IsExist() {
						_ = f.Delete()
					}
					w = zip.NewWriter(f.File())
					defer func() { _ = w.Close() }()
				}
				if iw, e := w.Create(key + ".png"); e != nil {
					return e
				} else if _, e := io.Copy(iw, f); e != nil {
					return e
				}
			}
		}
	}
	return nil
}

type wrapperFs struct {
	f fs.FS
	c func() error
}

func (t wrapperFs) Open(name string) (fs.File, error) {
	return t.f.Open(name)
}

func (t wrapperFs) Close() error {
	return t.c()
}

func (t *danmuEmotes) GetEmotesDir(dir string) interface {
	fs.FS
	io.Closer
} {
	if dir != "" && file.IsExist(dir+"/emotes.zip") {
		if rc, e := zip.OpenReader(dir + "/emotes.zip"); e == nil {
			return wrapperFs{
				f: rc,
				c: rc.Close,
			}
		}
	}
	return wrapperFs{
		f: file.DirFS(t.Dir),
		c: func() error { return nil },
	}
}

func loadCsv(savePath string, filename ...string) iter.Seq[Data] {
	return func(yield func(Data) bool) {
		csvf := file.New(savePath+append(filename, "0.csv")[0], 0, false)
		defer func() { _ = csvf.Close() }()

		if !csvf.IsExist() {
			return
		}

		var (
			data = Data{}
			line = []byte{}
		)
		for i := 0; true; i += 1 {
			if e := csvf.ReadUntilV2(&line, []byte{'\n'}, humanize.KByte, humanize.MByte); len(line) != 0 {
				lined := bytes.SplitN(line, []byte{','}, 3)
				if len(lined) == 3 {
					if t, e := strconv.ParseFloat(unsafe.B2S(lined[0]), 64); e == nil {
						if e := json.Unmarshal(lined[2], &data); e == nil {
							data.Time = t
							if data.Style.Color == "" {
								data.Style.Color = "#FFFFFF"
							}
							if !yield(data) {
								return
							}
						}
					}
				}
			} else if e != nil {
				break
			}
		}
	}
}

type DataStyle struct {
	Color  string `json:"color"`
	Border bool   `json:"border"`
	Mode   int    `json:"mode"`
}

type Data struct {
	Text  string    `json:"text"`
	Style DataStyle `json:"style"`
	Time  float64   `json:"time"`
}
