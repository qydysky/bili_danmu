package json

import (
	"encoding/json"

	cmp "github.com/qydysky/part/component2"
	pio "github.com/qydysky/part/io"
	pool "github.com/qydysky/part/pool"
)

type I interface {
	Unmarshal(data []byte, v any) error
}

func init() {
	reuseJ := &j{
		pool: pool.New(pool.PoolFunc[i]{
			New: func() (s *i) {
				s = &i{}
				s.bufReader = pio.NewBufReader()
				s.decoder = json.NewDecoder(s.bufReader)
				return
			},
		}, 100),
	}
	cmp.RegisterOrPanic[I](`json`, reuseJ)
}

type j struct {
	pool *pool.Buf[i]
}

type i struct {
	bufReader *pio.BufReader
	decoder   *json.Decoder
}

func (t *j) Unmarshal(data []byte, v any) error {
	i := t.pool.Get()
	defer t.pool.Put(i)
	if e := i.bufReader.Put(data); e != nil {
		i.bufReader.Clear()
		return e
	} else if e := i.decoder.Decode(v); e != nil {
		i.bufReader.Clear()
		i.decoder = json.NewDecoder(i.bufReader)
		return e
	}
	return nil
}
