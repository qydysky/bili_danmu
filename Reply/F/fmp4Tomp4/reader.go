package fmp4Tomp4

import (
	"bytes"
	"errors"
	"fmt"

	file "github.com/qydysky/part/file"
	"golang.org/x/exp/slices"
)

type Box struct {
	Name       string
	Size       int64
	HeaderSize int64
	Index      int64
}

type reader struct {
	f    *file.File
	Boxs []Box
	buf  []byte
}

func NewReader(f *file.File) (*reader, error) {
	if f.Config.AutoClose {
		return nil, errors.New("file AutoClose must false")
	}
	return &reader{f: f}, nil
}

func (t *reader) getN(name string, n int) (ls []Box) {
	for i := 0; i < len(t.Boxs); i++ {
		box := t.Boxs[i]
		if box.Name == name {
			ls = append(ls, box)
			if len(ls) >= n {
				return
			}
		}
	}
	return
}

type BoxReader struct {
	Buf []byte
}

func (t *BoxReader) I32(offset int) int32 {
	return btoi32(t.Buf, offset)
}
func (t *BoxReader) I64(offset int) int64 {
	return btoi64(t.Buf, offset)
}

func (t *reader) Read(box Box, size int) BoxReader {
	if size == -1 {
		t.bufChange(int(box.Size))
	} else {
		t.bufChange(size)
	}
	_ = t.f.SeekIndex(box.Index, file.AtOrigin)
	_, _ = t.f.Read(t.buf)
	return BoxReader{t.buf}
}

func (t *reader) bufChange(size int) {
	if n := size - len(t.buf); n > 0 {
		if size <= cap(t.buf) {
			t.buf = t.buf[:size]
		} else {
			t.buf = append(t.buf, make([]byte, n)...)
		}
	} else if n < 0 {
		t.buf = t.buf[:size]
	}
	clear(t.buf)
}

// 正常将返回io.EOF
func (t *reader) Parse() error {
	stat, e := t.f.Stat()
	if e != nil {
		return e
	}

	var (
		b4          = make([]byte, 4)
		b8          = make([]byte, 8)
		parseInside = []string{"moov", "trak", "mdia", "minf", "stbl", "moof", "traf"}
	)
	for i := int64(0); i < stat.Size(); {
		boxHeaderSize := 0
		if n, e := t.f.Read(b4); e != nil {
			return e
		} else {
			boxHeaderSize += n
		}
		size := int64(btoi32(b4, 0))
		if n, e := t.f.Read(b4); e != nil {
			return e
		} else {
			boxHeaderSize += n
		}
		name := string(b4)
		if size == 1 {
			if n, e := t.f.Read(b8); e != nil {
				return e
			} else {
				boxHeaderSize += n
			}
			size = btoi64(b8, 0)
		}
		t.Boxs = append(t.Boxs, Box{
			Name:       name,
			Size:       size,
			Index:      i,
			HeaderSize: int64(boxHeaderSize),
		})
		if !slices.Contains(parseInside, name) {
			seedSize := size - int64(boxHeaderSize)
			if e := t.f.SeekIndex(seedSize, file.AtCurrent); e != nil {
				return e
			}
			i += size
		} else {
			i += int64(boxHeaderSize)
		}
	}
	return nil
}

type Track struct {
	chunkSize   []int64
	sampleCount []int32
	dur         []int32
	size        []int32
}

func (t *reader) ParseTrun() (tracks map[int32]*Track, err error) {
	tracks = make(map[int32]*Track)

	var (
		b4  = make([]byte, 4)
		b8  = make([]byte, 8)
		b12 = make([]byte, 12)
	)

	for i := 0; i < len(t.Boxs); i++ {
		box := t.Boxs[i]
		if box.Name == "tfhd" {
			_ = t.f.SeekIndex(box.Index+box.HeaderSize, file.AtOrigin)

			if _, e := t.f.Read(b8); e != nil {
				err = e
				return
			}
			trackID := btoi32(b8, 4)
			defaultSampleDuration := int32(0)
			{
				var offset int64
				if b8[3]&0x01 == 0x01 {
					offset += 8
				}
				if b8[3]&0x02 == 0x02 {
					offset += 4
				}
				if b8[3]&0x08 == 0x08 {
					_ = t.f.SeekIndex(offset, file.AtCurrent)
					if _, e := t.f.Read(b4); e != nil {
						err = e
						return
					}
				}
				defaultSampleDuration = btoi32(b4, 0)
			}

			trackO, ok := tracks[trackID]
			if !ok {
				tracks[trackID] = &Track{}
				trackO = tracks[trackID]
			}

			for ; i < len(t.Boxs); i++ {
				box := t.Boxs[i]
				if box.Name == "trun" {
					_ = t.f.SeekIndex(box.Index+box.HeaderSize, file.AtOrigin)
					if _, e := t.f.Read(b8); e != nil {
						err = e
						return
					}

					_ = t.f.SeekIndex(8, file.AtCurrent)

					var chunkSize int64
					if bytes.Equal(b8[:4], []byte{0x01, 0x00, 0x0b, 0x05}) {
						sampleCount := btoi32(b8, 4)
						if sampleCount*12 != int32(box.Size-24) {
							err = errors.New("wrong trun trunSize not match sampleCount")
							return
						}
						trackO.sampleCount = append(trackO.sampleCount, sampleCount)
						for i := int32(0); i < sampleCount; i++ {
							if _, e := t.f.Read(b12); e != nil {
								err = e
								return
							}
							trackO.dur = append(trackO.dur, btoi32(b12, 0))
							trackO.size = append(trackO.size, btoi32(b12, 4))
							chunkSize += int64(btoi32(b12, 4))
						}
					} else if bytes.Equal(b8[:4], []byte{0x01, 0x00, 0x02, 0x05}) {
						sampleCount := btoi32(b8, 4)
						if sampleCount*4 != int32(box.Size-24) {
							err = errors.New("wrong trun trunSize not match sampleCount")
							return
						}
						trackO.sampleCount = append(trackO.sampleCount, sampleCount)
						for i := int32(0); i < sampleCount; i++ {
							if _, e := t.f.Read(b4); e != nil {
								err = e
								return
							}
							trackO.dur = append(trackO.dur, defaultSampleDuration)
							trackO.size = append(trackO.size, btoi32(b4, 0))
							chunkSize += int64(btoi32(b4, 0))
						}
					} else {
						err = fmt.Errorf("wrong trun tr_flag(%v)", b8[:4])
						return
					}
					trackO.chunkSize = append(trackO.chunkSize, chunkSize)

					break
				}
			}
		}
	}

	return
}
