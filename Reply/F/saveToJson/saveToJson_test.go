package savetojson

import (
	"errors"
	"io"
	"slices"
	"testing"

	"github.com/dustin/go-humanize"
	comp "github.com/qydysky/part/component2"
	pf "github.com/qydysky/part/file"
)

func Benchmark(b *testing.B) {
	i := comp.GetV3[interface {
		Init(path any)
		Write(data []byte)
		Close()
	}](`saveToJson`)

	_ = i.Run(func(i interface {
		Close()
		Init(path any)
		Write(data []byte)
	}) error {
		// i.Init(`1.json`)
		return nil
	})

	data := make([]byte, humanize.MByte)

	for b.Loop() {
		_ = i.Run(func(i interface {
			Close()
			Init(path any)
			Write(data []byte)
		}) error {
			i.Write(data)
			return nil
		})
	}
	_ = i.Run(func(i interface {
		Close()
		Init(path any)
		Write(data []byte)
	}) error {
		i.Close()
		return nil
	})

}

func Test(t *testing.T) {
	i := comp.GetV3(`saveToJson`, comp.PreFuncErr[interface {
		Init(path any)
		Write(data *[]byte)
		Close()
	}]{})

	_ = i.Run(func(i interface {
		Close()
		Init(path any)
		Write(data *[]byte)
	}) error {
		i.Init(`1.json`)
		if !pf.IsExist(`1.json`) {
			t.Fatal()
		}
		data := []byte(`123`)
		i.Write(&data)
		if tmp, e := pf.Open(`1.json`).ReadAll(10, 10); e != nil && !errors.Is(e, io.EOF) {
			t.Fatal(e)
		} else if !slices.Equal(tmp, []byte(`[123,`)) {
			t.Fatal(string(tmp))
		}
		i.Close()
		if tmp, e := pf.Open(`1.json`).ReadAll(10, 10); e != nil && !errors.Is(e, io.EOF) {
			t.Fatal(e)
		} else if !slices.Equal(tmp, []byte(`[123]`)) {
			t.Fatal()
		}
		_ = pf.Open(`1.json`).Delete()
		return nil
	})
}
