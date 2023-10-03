package fmp4Tomp4

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	comp "github.com/qydysky/part/component"
	file "github.com/qydysky/part/file"
)

// 直接保存下来的mp4在chrome上无法直接播放
//
// https://serverfault.com/questions/738881/chrome-makes-way-too-many-requests-22000-while-downloading-mp4-video-34mb
type Sign struct {
	// 重设mp4的时间戳
	conver func(ctx context.Context, ptr *string) error
}

func init() {
	sign := Sign{
		conver: conver,
	}
	if e := comp.Put[string](comp.Sign[Sign](`conver`), sign.conver); e != nil {
		panic(e)
	}
}

var (
	ErrParse = errors.New("ErrParse")
)

func conver(ctx context.Context, ptr *string) error {
	be := time.Now()
	fmt.Println("conver")
	defer fmt.Printf("conver fin (%v)\n", time.Since(be))

	sf := file.New(*ptr+"0.mp4", 0, false)
	if !sf.IsExist() {
		return nil
	}
	defer sf.Close()

	r, _ := NewReader(sf)
	if e := r.Parse(); e != nil || len(r.Boxs) == 0 {
		return e
	}
	traks, e := r.ParseTrun()
	if e != nil {
		return e
	}
	fmt.Printf("conver parse ok (%v)\n", time.Since(be))

	boxReader := r.Read(r.getN("mvhd", 1)[0], 30)
	timescale := boxReader.I32(20)
	scaleDur := boxReader.I32(24)
	mainDuration := float64(scaleDur) / float64(timescale)

	tf := file.New(*ptr+"1.mp4", 0, false)
	if tf.IsExist() {
		_ = tf.Delete()
	}
	defer tf.Close()

	w, e := NewBoxWriter(tf)
	if e != nil {
		return e
	}

	// ftyp
	{
		ftyp := w.Box("ftyp")
		ftyp.Write([]byte("isom"))
		ftyp.Write(itob32(512))
		ftyp.Write([]byte("isom"))
		ftyp.Write([]byte("iso2"))
		ftyp.Write([]byte("avc1"))
		ftyp.Write([]byte("mp41"))
		if e := ftyp.Close(); e != nil {
			return e
		}
	}

	// moov
	{
		moov := w.Box("moov")
		// mvhd
		{
			mvhd := moov.Box("mvhd")
			mvhd.Write(make([]byte, 12))
			mvhd.Write(itob32(1000))
			mvhd.Write(itob32(int32(mainDuration * 1000)))
			mvhd.Write(r.Read(r.getN("mvhd", 1)[0], -1).Buf[28:])
			if e := mvhd.Close(); e != nil {
				return e
			}
		}
		// trak
		var trakCount = -1
		for trakId, trakSum := range traks {
			fmt.Printf("conver traks (%v)(%v)\n", trakId, time.Since(be))

			trakCount++
			trak := moov.Box("trak")
			// tkhd
			{
				if boxs := r.getN("trak", 2); len(boxs) != 2 {
					return errors.Join(ErrParse, fmt.Errorf("trak"))
				} else {
					tkhd := trak.Box("tkhd")
					tkhd.Write([]byte{0, 0, 0, 3})
					tkhd.Write(make([]byte, 8))
					tkhd.Write(itob32(trakId))
					tkhd.Write(make([]byte, 4))
					tkhd.Write(itob32(int32(mainDuration * 1000)))
					tkhd.Write(r.Read(boxs[trakCount], -1).Buf[32:])
					if e := tkhd.Close(); e != nil {
						return e
					}
				}
			}
			// mdia
			{
				mdia := trak.Box("mdia")
				// mdhd
				{
					if boxs := r.getN("mdhd", 2); len(boxs) != 2 {
						return errors.Join(ErrParse, fmt.Errorf("mdhd"))
					} else {
						mdhd := mdia.Box("mdhd")
						mdhd.Write(r.Read(boxs[trakCount], -1).Buf)
						if e := mdhd.Close(); e != nil {
							return e
						}
					}
				}
				// hdlr
				var handlerType = make([]byte, 4)
				{
					if boxs := r.getN("hdlr", 2); len(boxs) != 2 {
						return errors.Join(ErrParse, fmt.Errorf("hdlr"))
					} else {
						hdlr := mdia.Box("hdlr")
						boxReader := r.Read(boxs[trakCount], -1)
						copy(handlerType, boxReader.Buf[16:20])
						hdlr.Write(boxReader.Buf)
						if e := hdlr.Close(); e != nil {
							return e
						}
					}
				}
				// minf
				{
					minf := mdia.Box("minf")
					// vmhd
					if bytes.Equal(handlerType, []byte("vide")) {
						if boxs := r.getN("vmhd", 1); len(boxs) != 1 {
							return errors.Join(ErrParse, fmt.Errorf("vmhd"))
						} else {
							vmhd := minf.Box("vmhd")
							vmhd.Write(r.Read(boxs[0], -1).Buf)
							if e := vmhd.Close(); e != nil {
								return e
							}
						}
					}
					// dinf
					{
						if boxs := r.getN("dinf", 2); len(boxs) != 2 {
							return errors.Join(ErrParse, fmt.Errorf("dinf"))
						} else {
							dinf := minf.Box("dinf")
							dinf.Write(r.Read(boxs[trakCount], -1).Buf)
							if e := dinf.Close(); e != nil {
								return e
							}
						}
					}
					// stbl
					{
						stbl := minf.Box("stbl")
						// stsd
						{
							if boxs := r.getN("stsd", 2); len(boxs) != 2 {
								return errors.Join(ErrParse, fmt.Errorf("stsd"))
							} else {
								stsd := stbl.Box("stsd")
								stsd.Write(r.Read(boxs[trakCount], -1).Buf)
								if e := stsd.Close(); e != nil {
									return e
								}
							}
						}
						// stts
						{
							stts := stbl.Box("stts")
							stts.Write([]byte{0, 0, 0, 0})
							stts.Write(itob32(int32(len(trakSum.dur))))
							for k, v := range trakSum.dur {
								stts.Write(itob32(int32(k)))
								stts.Write(itob32(v))
							}
							if e := stts.Close(); e != nil {
								return e
							}
						}
						// stsc
						{
							stsc := stbl.Box("stsc")
							stsc.Write([]byte{0, 0, 0, 0})
							stsc.Write(itob32(int32(len(trakSum.sampleCount))))
							for k, v := range trakSum.sampleCount {
								stsc.Write(itob32(int32(k)))
								stsc.Write(itob32(v))
								stsc.Write([]byte{0, 0, 0, 1})
							}
							if e := stsc.Close(); e != nil {
								return e
							}
						}
						// stsz
						{
							stsz := stbl.Box("stsz")
							stsz.Write([]byte{0, 0, 0, 0})
							stsz.Write(itob32(int32(len(trakSum.size))))
							for _, v := range trakSum.size {
								stsz.Write(itob32(v))
							}
							if e := stsz.Close(); e != nil {
								return e
							}
						}
						// co64
						{
							co64 := stbl.Box("co64")
							co64.Write([]byte{0, 0, 0, 0})
							co64.Write(itob32(int32(len(trakSum.chunkSize))))

							cuIndex, _ := tf.CurIndex()
							cuIndex += int64(8*len(trakSum.chunkSize)) + 8

							co64.Write(itob64(cuIndex))
							for i := 0; i < len(trakSum.chunkSize)-1; i++ {
								co64.Write(itob64(cuIndex + trakSum.chunkSize[i]))
							}
							if e := co64.Close(); e != nil {
								return e
							}
						}
						if e := stbl.Close(); e != nil {
							return e
						}
					}
					if e := minf.Close(); e != nil {
						return e
					}
				}
				if e := mdia.Close(); e != nil {
					return e
				}
			}
			if e := trak.Close(); e != nil {
				return e
			}
		}
		if e := moov.Close(); e != nil {
			return e
		}
		fmt.Printf("conver moov fin (%v)\n", time.Since(be))
	}

	// mdat
	{
		mdat := w.Box("mdat")
		for i := 0; i < len(r.Boxs); i++ {
			box := r.Boxs[i]
			if box.Name == "mdat" {
				if e := sf.SeekIndex(box.Index+box.HeaderSize, file.AtOrigin); e != nil {
					return e
				}
				mdat.CopyFrom(sf, uint64(box.Size-box.HeaderSize))
			}
		}
		if e := mdat.Close(); e != nil {
			return e
		}
	}
	return nil
}

// func bufChange(buf []byte, size int) {
// 	if n := size - len(buf); n > 0 {
// 		if size <= cap(buf) {
// 			buf = buf[:size]
// 		} else {
// 			buf = append(buf, make([]byte, n)...)
// 		}
// 	} else if n < 0 {
// 		buf = buf[:size]
// 	}
// 	clear(buf)
// }

// type track struct {
// 	chunkSize   []int64
// 	sampleCount []int32
// 	dur         []int32
// 	size        []int32
// }

// type wt struct {
// 	sf  *reader
// 	f   *file.File
// 	buf []byte
// 	m   map[string]float64
// }

// func NewWt(sf *reader, tf *file.File) *wt {
// 	return &wt{
// 		sf:  sf,
// 		f:   tf,
// 		buf: make([]byte, 1<<20),
// 		m:   make(map[string]float64),
// 	}
// }

// func (t *wt) start(boxName string, wrongPanic bool) (wSize int) {
// 	if _, e := t.f.Write([]byte{0, 0, 0, 1}, false); e != nil {
// 		panic(e)
// 	}
// 	if _, e := t.f.Write([]byte(boxName), false); e != nil {
// 		panic(e)
// 	}
// 	if _, e := t.f.Write(make([]byte, 8), false); e != nil {
// 		panic(e)
// 	}
// 	wSize = 16
// 	return
// }

// func (t *wt) skipt(size int) (n int) {
// 	if size == 0 {
// 		return
// 	}
// 	t.bufChange(size)
// 	if n, e := io.Writer(t.f.File()).Write(t.buf); e != nil {
// 		panic(e)
// 	} else {
// 		return n
// 	}
// }

// func (t *wt) skips(size int64) (n int64) {
// 	if size == 0 {
// 		return
// 	}
// 	if e := t.sf.f.SeekIndex(size, file.AtCurrent); e != nil {
// 		panic(e)
// 	}
// 	return size
// }

// func (t *wt) w(p []byte) (n int) {
// 	if n, e := t.f.Write(p, false); e != nil {
// 		panic(e)
// 	} else {
// 		return n
// 	}
// }

// func (t *wt) r(size int) (n int64) {
// 	t.bufChange(size)
// 	if n, e := t.sf.f.Read(t.buf); e != nil {
// 		panic(e)
// 	} else {
// 		return int64(n)
// 	}
// }

// func (t *wt) bufChange(size int) {
// 	if n := size - len(t.buf); n > 0 {
// 		if size <= cap(t.buf) {
// 			t.buf = t.buf[:size]
// 		} else {
// 			t.buf = append(t.buf, make([]byte, n)...)
// 		}
// 	} else if n < 0 {
// 		t.buf = t.buf[:size]
// 	}
// 	clear(t.buf)
// }

// func (t *wt) fin(size int) int {
// 	if e := t.f.SeekIndex(-int64(size), file.AtCurrent); e != nil {
// 		panic(e)
// 	}
// 	if e := t.f.SeekIndex(8, file.AtCurrent); e != nil {
// 		panic(e)
// 	}
// 	if _, e := t.f.Write(itob64(int64(size)), false); e != nil {
// 		panic(e)
// 	}
// 	if e := t.f.SeekIndex(int64(size-16), file.AtCurrent); e != nil {
// 		panic(e)
// 	}
// 	return size
// }

// func (t *wt) copyBox(boxName string) int {
// 	wSize := t.start(boxName, true)
// 	t.r(int(lSize))
// 	wSize += t.w(t.buf)
// 	return t.fin(wSize)
// }

// func (t *wt) ftyp() {
// 	n := t.start("ftyp", true)
// 	n += t.w([]byte("isom"))
// 	n += t.w(itob32(512))
// 	n += t.w([]byte("isom"))
// 	n += t.w([]byte("iso2"))
// 	n += t.w([]byte("avc1"))
// 	n += t.w([]byte("mp41"))
// 	t.fin(n)
// }

// func (t *wt) moov() {
// 	n, _ := t.start("moov", true)
// 	n += t.mvhd()
// 	n += t.trak()
// 	t.fin(n)
// }

// func (t *wt) mvhd() int {
// 	wSize, lSize := t.start("mvhd", true)
// 	lSize -= t.skips(12)
// 	lSize -= t.r(4)
// 	timescale := btoi32(t.buf, 0)
// 	lSize -= t.r(4)
// 	duration := btoi32(t.buf, 0)
// 	t.m["mainDuration"] = float64(duration) / float64(timescale)
// 	wSize += t.skipt(12)
// 	wSize += t.w(itob32(1000))
// 	wSize += t.w(itob32(timescale * duration / 1000))
// 	lSize -= t.r(56)
// 	wSize += t.w(t.buf)
// 	t.skips(lSize)
// 	return t.fin(wSize)
// }

// func (t *wt) trak() int {
// 	wSize, _ := t.start("trak", false)
// 	if wSize == 0 {
// 		return 0
// 	}

// 	wSize += t.tkhd()
// 	wSize += t.mdia()

// 	return t.fin(wSize)
// }

// func (t *wt) tkhd() int {
// 	wSize, lSize := t.start("mvhd", true)
// 	lSize -= t.r(12)
// 	wSize += t.w(t.buf)
// 	lSize -= t.r(4)
// 	wSize += t.w(t.buf)
// 	wSize += t.skipt(4)
// 	wSize += t.w(itob32(int32(t.m["mainDuration"] / 1000)))
// 	wSize += t.skipt(10)
// 	lSize -= t.skips(18)
// 	lSize -= t.r(52)
// 	wSize += t.w(t.buf)
// 	t.skips(lSize)
// 	return t.fin(wSize)
// }

// func (t *wt) mdia() int {
// 	wSize, _ := t.start("mdia", false)
// 	if wSize == 0 {
// 		return 0
// 	}

// 	wSize += t.mdhd()
// 	wSize += t.copyBox(`hdlr`)
// 	wSize += t.minf()

// 	return t.fin(wSize)
// }

// func (t *wt) mdhd() int {
// 	wSize, lSize := t.start("mdhd", true)
// 	lSize -= t.r(16)
// 	wSize += t.w(t.buf)
// 	lSize -= t.skips(8)
// 	wSize += t.w(itob32(1000))
// 	wSize += t.w(itob32(int32(t.m["mainDuration"] / 1000)))
// 	lSize -= t.r(4)
// 	wSize += t.w(t.buf)
// 	t.skips(lSize)
// 	return t.fin(wSize)
// }

// func (t *wt) minf() int {
// 	wSize, _ := t.start("mdia", false)
// 	if wSize == 0 {
// 		return 0
// 	}

// 	wSize += t.copyBox(`vmhd`)
// 	wSize += t.copyBox(`dinf`)
// 	wSize += t.stbl()

// 	return t.fin(wSize)
// }

// func (t *wt) stbl() int {
// 	wSize, _ := t.start("stbl", false)
// 	if wSize == 0 {
// 		return 0
// 	}

// 	wSize += t.copyBox(`stsd`)
// 	wSize += t.stts()

// 	return t.fin(wSize)
// }

// func (t *wt) stts() int {
// 	wSize, lSize := t.start("stts", true)
// 	for {
// 		lSize -= t.r(1 << 20)
// 	}
// 	// lSize -= t.r(16)
// 	// wSize += t.w(t.buf)
// 	// lSize -= t.skips(8)
// 	// wSize += t.w(itob32(1000))
// 	// wSize += t.w(itob32(int32(t.m["mainDuration"] / 1000)))
// 	// lSize -= t.r(4)
// 	// wSize += t.w(t.buf)
// 	t.skips(lSize)
// 	return t.fin(wSize)
// }

func btoi64(b []byte, offset int) int64 {
	s := 8
	bu := make([]byte, s)
	l := len(b) - offset
	if l > s {
		l = s
	}
	for i := 0; i < s && i < l; i++ {
		bu[i+s-l] = b[offset+i]
	}

	//binary.BigEndian.Uint64
	return int64(uint64(bu[7]) | uint64(bu[6])<<8 | uint64(bu[5])<<16 | uint64(bu[4])<<24 |
		uint64(bu[3])<<32 | uint64(bu[2])<<40 | uint64(bu[1])<<48 | uint64(bu[0])<<56)
}

func btoi32(b []byte, offset int) int32 {
	s := 4
	bu := make([]byte, s)
	l := len(b) - offset
	if l > s {
		l = s
	}
	for i := 0; i < s && i < l; i++ {
		bu[i+s-l] = b[offset+i]
	}

	//binary.BigEndian.Uint32
	return int32((uint32(bu[3]) | uint32(bu[2])<<8 | uint32(bu[1])<<16 | uint32(bu[0])<<24))
}

func itob64(v int64) []byte {
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

func itob32(v int32) []byte {
	//binary.BigEndian.PutUint32
	b := make([]byte, 4)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	return b
}
