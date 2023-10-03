package fmp4Tomp4

import (
	"errors"
	"fmt"

	file "github.com/qydysky/part/file"
	pio "github.com/qydysky/part/io"
)

var (
	ErrFileAutoClose = errors.New("file AutoClose must false")
	ErrSeed          = errors.New("ErrSeed")
	ErrWrite         = errors.New("ErrWrite")
)

type boxWriter struct {
	f  *file.File
	wn int64
	e  error
	p  *boxWriter
}

func NewBoxWriter(f *file.File) (t *boxWriter, err error) {
	if f.Config.AutoClose {
		return nil, ErrFileAutoClose
	}
	t = &boxWriter{f: f}
	return
}

func (t *boxWriter) Box(name string) (tc *boxWriter) {
	if t.e != nil {
		return
	}
	tc = &boxWriter{f: t.f, p: t, e: t.e}
	tc.Write([]byte{0, 0, 0, 1})
	tc.Write([]byte(name))
	tc.wn = 0
	tc.Write(make([]byte, 8))
	return
}

func (t *boxWriter) Write(b []byte) (tc *boxWriter) {
	if t.e != nil {
		return t
	}
	n := 0
	n, t.e = t.f.Write(b, false)
	t.wn += int64(n)
	return t
}

func (t *boxWriter) CopyFrom(f *file.File, size uint64) (tc *boxWriter) {
	if t.e != nil {
		return t
	}
	t.e = f.CopyTo(t.f, pio.CopyConfig{MaxByte: size}, false)
	t.wn += int64(size)
	return t
}

func (t *boxWriter) Close() error {
	if t.e != nil {
		return t.e
	}
	t.e = t.f.SeekIndex(-t.wn, file.AtCurrent)
	if t.e != nil {
		return errors.Join(ErrSeed, t.e, fmt.Errorf("Arg %v", -t.wn))
	}
	_, t.e = t.f.Write(itob64(t.wn), false)
	if t.e != nil {
		return errors.Join(ErrWrite, t.e, fmt.Errorf("Arg %v", -t.wn))
	}
	t.e = t.f.SeekIndex(t.wn-8, file.AtCurrent)
	if t.p != nil {
		t.p.wn += t.wn
	}
	return t.e
}
