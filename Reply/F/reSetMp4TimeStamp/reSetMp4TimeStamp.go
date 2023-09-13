package reSetMp4TimeStamp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	comp "github.com/qydysky/part/component"
	file "github.com/qydysky/part/file"
)

// 直接保存下来的mp4在chrome上无法直接播放
//
// https://serverfault.com/questions/738881/chrome-makes-way-too-many-requests-22000-while-downloading-mp4-video-34mb
type Sign struct {
	// 重设mp4的时间戳
	resetTS func(ctx context.Context, ptr *string) error
}

func init() {
	sign := Sign{
		resetTS: resetTS,
	}
	if e := comp.Put[string](comp.Sign[Sign](`resetTS`), sign.resetTS); e != nil {
		panic(e)
	}
}

func resetTS(ctx context.Context, ptr *string) error {
	be := time.Now()
	fmt.Println("resetTS")
	defer fmt.Printf("resetTS fin (%v)\n", time.Since(be))

	f := file.New(*ptr+"0.mp4", 0, false)
	if !f.IsExist() {
		return nil
	}
	defer f.Close()

	var (
		byte4     = make([]byte, 4)
		byte16    = make([]byte, 16)
		bgdts     = make(map[int32]int64)
		eddts     = make(map[int32]int64)
		timescale = make(map[int32]int64)
	)

	for {
		if e := f.SeekUntil([]byte("tkhd"), file.AtCurrent, 1<<17, 1<<22); e != nil {
			if errors.Is(e, file.ErrMaxReadSizeReach) {
				break
			}
			if errors.Is(e, io.EOF) {
				break
			}
			return e
		}
		_ = f.SeekIndex(16, file.AtCurrent)
		if _, e := f.Read(byte4); e != nil {
			return e
		}
		trackId := btoi32(byte4, 0)

		bgdts[trackId] = -1
		eddts[trackId] = 0
	}

	// rewrite dts
	_ = f.SeekIndex(0, file.AtOrigin)
	for {
		if e := f.SeekUntil([]byte("tfhd"), file.AtCurrent, 1<<17, 1<<20); e != nil {
			if errors.Is(e, file.ErrMaxReadSizeReach) {
				continue
			}
			if errors.Is(e, io.EOF) {
				break
			}
			return e
		}
		_ = f.SeekIndex(8, file.AtCurrent)
		if _, e := f.Read(byte4); e != nil {
			return e
		}
		trackID := btoi32(byte4, 0)

		if e := f.SeekUntil([]byte("tfdt"), file.AtCurrent, 1<<17, 1<<20); e != nil {
			if errors.Is(e, file.ErrMaxReadSizeReach) {
				continue
			}
			if errors.Is(e, io.EOF) {
				break
			}
			return e
		}
		if _, e := f.Read(byte16); e != nil {
			return e
		}
		switch byte16[4] {
		case 0:
			ts := int64(btoi32(byte16, 12))
			eddts[trackID] = ts
			if e := f.SeekIndex(-4, file.AtCurrent); e != nil {
				return e
			}
			if bgdts[trackID] == -1 {
				bgdts[trackID] = ts
			}
			if _, e := f.Write(itob32(int32(ts-bgdts[trackID])), false); e != nil {
				return e
			}
		case 1:
			ts := btoi64(byte16, 8)
			eddts[trackID] = ts
			if e := f.SeekIndex(-8, file.AtCurrent); e != nil {
				return e
			}
			if bgdts[trackID] == -1 {
				bgdts[trackID] = ts
			}
			if _, e := f.Write(itob64(ts-bgdts[trackID]), false); e != nil {
				return e
			}
		default:
			return fmt.Errorf("unknow tfdt version %x", byte16[8])
		}
	}

	_ = f.SeekIndex(0, file.AtOrigin)
	for {
		if e := f.SeekUntil([]byte("tkhd"), file.AtCurrent, 1<<17, 1<<22); e != nil {
			if errors.Is(e, file.ErrMaxReadSizeReach) {
				break
			}
			if errors.Is(e, io.EOF) {
				break
			}
			return e
		}
		_ = f.SeekIndex(16, file.AtCurrent)
		if _, e := f.Read(byte4); e != nil {
			return e
		}
		trackId := btoi32(byte4, 0)

		if e := f.SeekUntil([]byte("mdhd"), file.AtCurrent, 1<<17, 1<<22); e != nil {
			if errors.Is(e, file.ErrMaxReadSizeReach) {
				break
			}
			if errors.Is(e, io.EOF) {
				break
			}
			return e
		}
		_ = f.SeekIndex(16, file.AtCurrent)
		if _, e := f.Read(byte4); e != nil {
			return e
		}

		timescale[trackId] = int64(btoi32(byte4, 0))
	}

	var duration int32
	for k, v := range bgdts {
		fmt.Println(eddts[k], v)
		duration = int32((eddts[k] - v) / timescale[k])
		break
	}

	_ = f.SeekIndex(0, file.AtOrigin)
	{
		if e := f.SeekUntil([]byte("moov"), file.AtCurrent, 1<<17, 1<<22); e != nil {
			return e
		}
	}

	// write mvhd
	_ = f.SeekIndex(0, file.AtOrigin)
	{
		if e := f.SeekUntil([]byte("mvhd"), file.AtCurrent, 1<<17, 1<<22); e != nil {
			return e
		}
		_ = f.SeekIndex(20, file.AtCurrent)
		if _, e := f.Write(itob32(duration), false); e != nil {
			return e
		}
	}

	// write tkhd mdhd
	_ = f.SeekIndex(0, file.AtOrigin)
	for {
		if e := f.SeekUntil([]byte("tkhd"), file.AtCurrent, 1<<17, 1<<20); e != nil {
			if errors.Is(e, file.ErrMaxReadSizeReach) {
				break
			}
			if errors.Is(e, io.EOF) {
				break
			}
			return e
		}
		_ = f.SeekIndex(24, file.AtCurrent)
		if _, e := f.Write(itob32(duration), false); e != nil {
			return e
		}

		if e := f.SeekUntil([]byte("mdhd"), file.AtCurrent, 1<<17, 1<<22); e != nil {
			if errors.Is(e, file.ErrMaxReadSizeReach) {
				continue
			}
			if errors.Is(e, io.EOF) {
				break
			}
			return e
		}
		_ = f.SeekIndex(20, file.AtCurrent)
		if _, e := f.Write(itob32(duration), false); e != nil {
			return e
		}
	}
	return nil
}

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
