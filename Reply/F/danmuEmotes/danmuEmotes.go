package danmuemotes

import (
	"context"
	"encoding/json"
	"strings"

	c "github.com/qydysky/bili_danmu/CV"
	comp "github.com/qydysky/part/component2"
	file "github.com/qydysky/part/file"
	phash "github.com/qydysky/part/hash"
	log "github.com/qydysky/part/log"
	reqf "github.com/qydysky/part/reqf"
)

type TargetInterface interface {
	SaveEmote(ctx context.Context, ptr struct {
		Logg *log.Log_interface
		Info []any
		Msg  *string
	}) (ret any, err error)
	Hashr(s string) (r string)
	SetLayerN(n int)
}

func init() {
	i := danmuEmotes{
		Dir:    "emots/",
		LayerN: 0,
	}
	if e := comp.Register[TargetInterface]("danmuEmotes", &i); e != nil {
		panic(e)
	}
	_, _ = file.New(i.Dir+"README.md", 0, true).Write([]byte(""), false)
}

type danmuEmotes struct {
	Dir    string
	LayerN int
}

func (t *danmuEmotes) SetLayerN(n int) {
	t.LayerN = n
}

func (t *danmuEmotes) SaveEmote(ctx context.Context, ptr struct {
	Logg *log.Log_interface
	Info []any
	Msg  *string
}) (ret any, err error) {
	if m, ok := ptr.Info[13].(map[string]any); ok {
		if url, ok := m[`url`].(string); ok {
			if !strings.Contains(*ptr.Msg, "[") {
				if emoticon_unique, ok := m[`emoticon_unique`].(string); ok {
					*ptr.Msg = "[" + *ptr.Msg + emoticon_unique + "]"
				}
			}
			savePath := t.Dir + t.Hashr(*ptr.Msg) + ".png"
			if !file.New(savePath, 0, true).IsExist() {
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
						ptr.Logg.L(`E: `, e)
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
				if e := json.Unmarshal([]byte(extrab), &E); e != nil {
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

						savePath := t.Dir + t.Hashr(k) + ".png"
						if file.New(savePath, 0, true).IsExist() {
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
								ptr.Logg.L(`E: `, e)
							}
						}()
					}
				}
			}
		}
	}
	return
}

func (t *danmuEmotes) Hashr(s string) (r string) {
	rs := phash.Md5String(s)
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
	return string(rr)
}
