package videofastseed

import (
	"errors"
	"os"
	"time"

	comp "github.com/qydysky/part/component2"
	file "github.com/qydysky/part/file"
)

type TargetInterface interface {
	InitGet(fastSeedFilePath string) (getIndex func(seedTo time.Duration) (int64, error), e error)
	InitSav(fastSeedFilePath string) (savIndex func(seedTo time.Duration, cuIndex int64) error, e error)
}

func init() {
	if e := comp.Register[TargetInterface]("videoFastSeed", videoFastSeed{}); e != nil {
		panic(e)
	}
}

var (
	ErrFormat    = errors.New("ErrFormat")
	ErrNoInitSav = errors.New("ErrNoInitSav")
	ErrNoInitGet = errors.New("ErrNoInitGet")
)

type videoFastSeed struct {
	initSav  bool
	initGet  bool
	filepath string
}

func (_ videoFastSeed) InitGet(fastSeedFilePath string) (getIndex func(seedTo time.Duration) (int64, error), e error) {
	t := videoFastSeed{}
	t.filepath = fastSeedFilePath
	f := file.New(t.filepath, -1, false)
	defer f.Close()
	if !f.IsExist() {
		return nil, os.ErrNotExist
	}
	t.initGet = true
	return t.GetIndex, nil
}

func (_ videoFastSeed) InitSav(fastSeedFilePath string) (savIndex func(seedTo time.Duration, cuIndex int64) error, e error) {
	t := videoFastSeed{}
	t.filepath = fastSeedFilePath
	f := file.New(t.filepath, -1, false)
	defer f.Close()
	if f.IsExist() {
		_ = f.Delete()
	}
	t.initSav = true
	return t.SavIndex, nil
}

func (t *videoFastSeed) SavIndex(ms time.Duration, cuIndex int64) error {
	if !t.initSav {
		return ErrNoInitSav
	}
	f := file.New(t.filepath, -1, false)
	defer f.Close()
	if _, e := f.Write(Itob64(ms.Milliseconds()), false); e != nil {
		return e
	}
	if _, e := f.Write(Itob64(cuIndex), false); e != nil {
		return e
	}
	return nil
}

func (t *videoFastSeed) GetIndex(seedTo time.Duration) (int64, error) {
	if !t.initGet {
		return -1, ErrNoInitGet
	}
	f := file.New(t.filepath, 0, false)
	defer f.Close()
	if !f.IsExist() {
		return -1, os.ErrNotExist
	}
	buf := make([]byte, 16)
	lastIndex := int64(0)
	for {
		if n, e := f.Read(buf); e != nil {
			return -1, e
		} else if n != 16 {
			return -1, ErrFormat
		} else {
			if ms := Btoi(buf[:8], 0, 8); ms > seedTo.Milliseconds() {
				return lastIndex, nil
			} else {
				lastIndex = Btoi(buf[8:], 0, 8)
			}
		}
	}
}

func Btoi(b []byte, offset int, size int) int64 {
	if size > 8 {
		panic("最大8位")
	}

	bu := make([]byte, 8)
	l := len(b) - offset
	if l > size {
		l = size
	}
	for i := 0; i < size && i < l; i++ {
		bu[i+8-size] = b[offset+i]
	}

	//binary.BigEndian.Uint64
	return int64(uint64(bu[7]) | uint64(bu[6])<<8 | uint64(bu[5])<<16 | uint64(bu[4])<<24 |
		uint64(bu[3])<<32 | uint64(bu[2])<<40 | uint64(bu[1])<<48 | uint64(bu[0])<<56)
}

func Itob64(v int64) []byte {
	//binary.BigEndian.PutUint64
	b := make([]byte, 8)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b
}
