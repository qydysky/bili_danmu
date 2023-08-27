package reSetMp4TimeStamp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

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
	fmt.Println("resetTS")
	defer fmt.Println("resetTS fin")

	f := file.New(*ptr+"0.mp4", 0, false)
	if !f.IsExist() {
		return nil
	}
	defer f.Close()
	var tfdtBuf = make([]byte, 16)
	var tfhdBuf = make([]byte, 12)
	var boxBuf = make([]byte, 4)
	var trackBuf = make([]byte, 4)
	var timescaleBuf = make([]byte, 4)
	var timescale int32
	var opTs = make(map[int32]int)
	var cuTs = make(map[int32]int)

	if e := f.SeekUntil([]byte("mvhd"), file.AtCurrent, 1<<17, 1<<22); e != nil && !errors.Is(e, file.ErrMaxReadSizeReach) {
		return e
	}
	if _, e := f.Read(boxBuf); e != nil {
		return e
	} else if !bytes.Equal(boxBuf, []byte("mvhd")) {
		return fmt.Errorf("wrong box:%v", string(boxBuf))
	}
	_ = f.SeekIndex(12, file.AtCurrent)
	if _, e := f.Read(timescaleBuf); e != nil {
		return e
	}
	timescale = btoi32(timescaleBuf, 0)
	fmt.Printf("resetTS timescale:%v\n", timescale)

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
		if _, e := f.Read(boxBuf); e != nil {
			return e
		} else if !bytes.Equal(boxBuf, []byte("tkhd")) {
			return fmt.Errorf("wrong box:%v", string(boxBuf))
		}
		_ = f.SeekIndex(12, file.AtCurrent)
		if _, e := f.Read(trackBuf); e != nil {
			return e
		}
		trackId := btoi32(trackBuf, 0)
		fmt.Printf("trackId %v\n", trackId)
		opTs[trackId] = -1
		cuTs[trackId] = 0
	}

	// fmt.Println("resetTimeStamp Druation %v(%v-%v)", time.Duration(int(time.Second)*(cuTS-opTs)), cuTS, opTs)

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
		if _, e := f.Read(tfhdBuf); e != nil {
			return e
		}

		trackID := btoi32(tfhdBuf, 8)

		if e := f.SeekUntil([]byte("tfdt"), file.AtCurrent, 1<<17, 1<<20); e != nil {
			if errors.Is(e, file.ErrMaxReadSizeReach) {
				continue
			}
			if errors.Is(e, io.EOF) {
				break
			}
			return e
		}
		if _, e := f.Read(tfdtBuf); e != nil {
			return e
		}
		switch tfdtBuf[4] {
		case 0:
			ts := int(btoi32(tfdtBuf, 12))
			cuTs[trackID] = ts
			if e := f.SeekIndex(-4, file.AtCurrent); e != nil {
				return e
			}
			if opTs[trackID] == -1 {
				opTs[trackID] = ts
			}
			if _, e := f.Write(itob32(int32(ts-opTs[trackID])), false); e != nil {
				return e
			}
		case 1:
			ts := int(btoi64(tfdtBuf, 8))
			cuTs[trackID] = ts
			if e := f.SeekIndex(-8, file.AtCurrent); e != nil {
				return e
			}
			if opTs[trackID] == -1 {
				opTs[trackID] = ts
			}
			if _, e := f.Write(itob64(int64(ts-opTs[trackID])), false); e != nil {
				return e
			}
		default:
			return fmt.Errorf("unknow tfdt version %x", tfdtBuf[8])
		}
	}

	for k, v := range opTs {
		fmt.Printf("track %v opTs:%v cuTS:%v\n", k, v, cuTs[k])
	}
	// fmt.Println("resetTimeStamp Druation %v(%v-%v)", time.Duration(int(time.Second)*(cuTS-opTs)), cuTS, opTs)

	var duration int32
	for k, v := range opTs {
		duration = int32(cuTs[k]-v) / timescale
		break
	}
	fmt.Printf("resetTS dur:%v\n", duration)

	_ = f.SeekIndex(0, file.AtOrigin)

	if e := f.SeekUntil([]byte("mvhd"), file.AtCurrent, 1<<17, 1<<22); !errors.Is(e, file.ErrMaxReadSizeReach) {
		return e
	}
	_ = f.SeekIndex(20, file.AtCurrent)
	if _, e := f.Write(itob32(duration), false); e != nil {
		return e
	}

	// write tkhd
	_ = f.SeekIndex(0, file.AtOrigin)
	for i := 0; i < len(opTs); i++ {
		if e := f.SeekUntil([]byte("tkhd"), file.AtCurrent, 1<<17, 1<<20); e != nil {
			if errors.Is(e, file.ErrMaxReadSizeReach) {
				continue
			}
			if errors.Is(e, io.EOF) {
				break
			}
			return e
		}
		_ = f.SeekIndex(16, file.AtCurrent)
		if _, e := f.Read(trackBuf); e != nil {
			return e
		}
		trackID := btoi32(trackBuf, 0)
		_ = f.SeekIndex(4, file.AtCurrent)
		if _, e := f.Write(itob32(int32(cuTs[trackID]-opTs[trackID])/timescale), false); e != nil {
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
