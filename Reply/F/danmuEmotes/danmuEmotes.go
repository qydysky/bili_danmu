package danmuemotes

import (
	"context"
	"encoding/json"
	"strings"

	c "github.com/qydysky/bili_danmu/CV"
	comp "github.com/qydysky/part/component"
	file "github.com/qydysky/part/file"
	log "github.com/qydysky/part/log"
	ppool "github.com/qydysky/part/pool"
	reqf "github.com/qydysky/part/reqf"
	pslice "github.com/qydysky/part/slice"
)

// path
var (
	SaveEmote = comp.NewComp(saveEmote)
	Hashr     = comp.NewComp(func(ctx context.Context, s string) (r string, e error) {
		return hashr(s), nil
	})
	danmuPool = ppool.New(ppool.PoolFunc[pslice.Buf[byte]]{
		New: func() *pslice.Buf[byte] {
			return pslice.New[byte]()
		},
		InUse: func(b *pslice.Buf[byte]) bool {
			return !b.IsEmpty()
		},
		Reuse: func(b *pslice.Buf[byte]) *pslice.Buf[byte] {
			return b
		},
		Pool: func(b *pslice.Buf[byte]) *pslice.Buf[byte] {
			b.Reset()
			return b
		},
	}, 100)
)

type Danmu struct {
	Logg *log.Log_interface
	Info []any
	Msg  *string
}

func init() {
	_, _ = file.New("emots/README.md", 0, true).Write([]byte(""), false)
}

func saveEmote(ctx context.Context, ptr Danmu) (ret any, err error) {
	if m, ok := ptr.Info[13].(map[string]any); ok {
		if url, ok := m[`url`].(string); ok {
			if !strings.Contains(*ptr.Msg, "[") {
				if emoticon_unique, ok := m[`emoticon_unique`].(string); ok {
					*ptr.Msg = "[" + *ptr.Msg + emoticon_unique + "]"
				}
			}
			*ptr.Msg = hashr(*ptr.Msg)
			savePath := "emots/" + *ptr.Msg + ".png"
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
					*ptr.Msg = hashr(*ptr.Msg)
					for k, v := range E.Emots {
						m, ok := v.(map[string]any)
						if !ok {
							continue
						}

						url, ok := m[`url`].(string)
						if !ok {
							continue
						}

						savePath := "emots/" + hashr(k) + ".png"
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

func hashr(s string) (r string) {
	buf := danmuPool.Get()
	defer danmuPool.Put(buf)

	emoteB := false
	for i := 0; i < len(s); i++ {
		if !emoteB {
			_ = buf.Append([]byte{s[i]})
			emoteB = s[i] == '['
			continue
		} else if s[i] == ']' {
			_ = buf.Append([]byte{s[i]})
			emoteB = false
			continue
		}
		switch s[i] {
		case '\\':
			_ = buf.Append([]byte("_1"))
		case '/':
			_ = buf.Append([]byte("_2"))
		case '*':
			_ = buf.Append([]byte("_3"))
		case '<':
			_ = buf.Append([]byte("_4"))
		case '>':
			_ = buf.Append([]byte("_5"))
		case '|':
			_ = buf.Append([]byte("_6"))
		case ':':
			_ = buf.Append([]byte("："))
		case '?':
			_ = buf.Append([]byte("？"))
		case '"':
			_ = buf.Append([]byte("“"))
		default:
			_ = buf.Append([]byte{s[i]})
		}
	}
	return string(buf.GetCopyBuf())
}
