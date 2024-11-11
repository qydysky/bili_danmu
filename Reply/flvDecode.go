package reply

import (
	"errors"
	"fmt"
	"io"
	"math"
	"time"

	"github.com/dustin/go-humanize"
	F "github.com/qydysky/bili_danmu/F"
	perrors "github.com/qydysky/part/errors"
	slice "github.com/qydysky/part/slice"
)

const (
	flvHeaderSize  = 9
	tagHeaderSize  = 11
	previouTagSize = 4

	streamId  = byte(0x00)
	videoTag  = byte(0x09)
	audioTag  = byte(0x08)
	scriptTag = byte(0x12)
)

var (
	flvHeaderSign = []byte{0x46, 0x4c, 0x56}

	ErrInit             = errors.New("ErrInit")
	ErrNoInit           = errors.New("ErrNoInit")
	ErrNoFoundFlvHeader = errors.New("ErrNoFoundFlvHeader")
	ErrNoFoundTagHeader = errors.New("ErrNoFoundTagHeader")
	ErrTagSizeZero      = errors.New("ErrTagSizeZero")
	ErrStreamId         = errors.New("ErrStreamId")
	ErrTagSize          = errors.New("ErrTagSize")
	ErrSignLost         = errors.New("ErrSignLost")
)

type FlvDecoder struct {
	Diff float64
	init bool
}

func NewFlvDecoder() *FlvDecoder {
	return &FlvDecoder{Diff: 100}
}

func (t *FlvDecoder) InitFlv(buf []byte) (frontBuf []byte, dropOffset int, err error) {
	if t.init {
		err = ErrInit
		return
	}

	var sign = 0x00
	var tagNum = 0

	if len(buf) < flvHeaderSize+previouTagSize {
		return
	}

	if buf[0] != flvHeaderSign[0] || buf[1] != flvHeaderSign[1] || buf[2] != flvHeaderSign[2] {
		err = ErrNoFoundFlvHeader
		return
	}

	if buf[flvHeaderSize]|buf[flvHeaderSize+1]|buf[flvHeaderSize+2]|buf[flvHeaderSize+3] != 0 {
		err = ErrTagSize
		return
	}

	for bufOffset := flvHeaderSize + previouTagSize; bufOffset >= flvHeaderSize && bufOffset+tagHeaderSize < len(buf); {

		if buf[bufOffset]&0b11000000 != 0 &&
			buf[bufOffset]&0b00011111 != videoTag &&
			buf[bufOffset]&0b00011111 != audioTag &&
			buf[bufOffset]&0b00011111 != scriptTag {
			err = ErrNoFoundTagHeader
			return
		}

		if buf[bufOffset+8]|buf[bufOffset+9]|buf[bufOffset+10] != streamId {
			err = ErrStreamId
			return
		}

		tagSize := int(F.Btoi32([]byte{0x00, buf[bufOffset+1], buf[bufOffset+2], buf[bufOffset+3]}, 0))
		if tagSize == 0 {
			err = ErrTagSizeZero
			return
		}

		if bufOffset+tagHeaderSize+tagSize+previouTagSize > len(buf) {
			return
		}

		tagSizeCheck := int(F.Btoi32(buf[bufOffset+tagHeaderSize+tagSize:bufOffset+tagHeaderSize+tagSize+previouTagSize], 0))
		if tagNum != 0 && tagSizeCheck != tagSize+tagHeaderSize {
			err = ErrTagSize
			return
		}

		tagNum += 1

		if sign != 0x07 { // ignore first video audio time_stamp
			if (buf[bufOffset] == videoTag) && (sign&0x04 == 0x00) {
				sign |= 0x04
			} else if (buf[bufOffset] == audioTag) && (sign&0x02 == 0x00) {
				sign |= 0x02
			} else if (buf[bufOffset] == scriptTag) && (sign&0x01 == 0x00) {
				sign |= 0x01
			} else {
				err = ErrSignLost
				return
			}
			bufOffset += tagSizeCheck + previouTagSize

			if sign == 0x07 {
				frontBuf = append(frontBuf, buf[0:bufOffset]...) // copy
				dropOffset = bufOffset
				t.init = true
				return
			}
		}
	}
	return
}

// this fuction read []byte and return flv header and all complete keyframe if possible.
// complete keyframe means the video and audio tags between two video key frames tag
func (t *FlvDecoder) SearchStreamTag(buf []byte, keyframe *slice.Buf[byte]) (dropOffset int, err error) {
	if !t.init {
		err = ErrNoInit
		return
	}

	keyframe.Reset()

	var (
		keyframeOp = -1
		lastAT     int
		lastVT     int
	)

	for bufOffset := 0; bufOffset >= 0 && bufOffset+tagHeaderSize < len(buf); {

		if buf[bufOffset]&0b11000000 != 0 &&
			buf[bufOffset]&0b00011111 != videoTag &&
			buf[bufOffset]&0b00011111 != audioTag &&
			buf[bufOffset]&0b00011111 != scriptTag {
			err = ErrNoFoundTagHeader
			return
		}

		if buf[bufOffset+8]|buf[bufOffset+9]|buf[bufOffset+10] != streamId {
			err = ErrStreamId
			return
		}

		tagSize := int(F.Btoi32([]byte{0x00, buf[bufOffset+1], buf[bufOffset+2], buf[bufOffset+3]}, 0))
		if tagSize == 0 {
			err = ErrTagSizeZero
			return
		}
		if bufOffset+tagHeaderSize+tagSize+previouTagSize > len(buf) {
			return
		}

		tagSizeCheck := int(F.Btoi32(buf[bufOffset+tagHeaderSize+tagSize:bufOffset+tagHeaderSize+tagSize+previouTagSize], 0))
		if tagSizeCheck != tagSize+tagHeaderSize {
			err = ErrTagSize
			return
		}

		timeStamp := int(F.Btoi32([]byte{buf[bufOffset+7], buf[bufOffset+4], buf[bufOffset+5], buf[bufOffset+6]}, 0))
		switch {
		case buf[bufOffset] == videoTag:
			lastVT = timeStamp
		case buf[bufOffset] == audioTag:
			lastAT = timeStamp
		default:
		}
		if lastAT != 0 && lastVT != 0 {
			diff := math.Abs(float64(lastVT - lastAT))
			if diff > t.Diff {
				err = fmt.Errorf("时间戳不匹配 %v %v (或许应调整flv音视频时间戳容差ms>%f)", lastVT, lastAT, diff)
				return
			}
		}

		if buf[bufOffset] == videoTag && buf[bufOffset+11]&0xf0 == 0x10 { //key frame
			if keyframeOp >= 0 {
				err = keyframe.Append(buf[keyframeOp:bufOffset])
				dropOffset = bufOffset
			}
			keyframeOp = bufOffset
		}
		bufOffset += tagSizeCheck + previouTagSize
	}

	return
}

func (t *FlvDecoder) oneF(buf []byte, w ...io.Writer) (dropT int, dropOffset int, err error) {

	if !t.init {
		err = ErrNoInit
		return
	}

	var (
		keyframeOp = -1
		lastAT     int
		lastVT     int
	)

	for bufOffset := 0; bufOffset >= 0 && bufOffset+tagHeaderSize < len(buf); {

		if buf[bufOffset]&0b11000000 != 0 &&
			buf[bufOffset]&0b00011111 != videoTag &&
			buf[bufOffset]&0b00011111 != audioTag &&
			buf[bufOffset]&0b00011111 != scriptTag {
			err = ErrNoFoundTagHeader
			return
		}

		if buf[bufOffset+8]|buf[bufOffset+9]|buf[bufOffset+10] != streamId {
			err = ErrStreamId
			return
		}

		tagSize := int(F.Btoi32([]byte{0x00, buf[bufOffset+1], buf[bufOffset+2], buf[bufOffset+3]}, 0))
		if tagSize == 0 {
			err = ErrTagSizeZero
			return
		}
		if bufOffset+tagHeaderSize+tagSize+previouTagSize > len(buf) {
			return
		}

		tagSizeCheck := int(F.Btoi32(buf[bufOffset+tagHeaderSize+tagSize:bufOffset+tagHeaderSize+tagSize+previouTagSize], 0))
		if tagSizeCheck != tagSize+tagHeaderSize {
			err = ErrTagSize
			return
		}

		timeStamp := int(F.Btoi32([]byte{buf[bufOffset+7], buf[bufOffset+4], buf[bufOffset+5], buf[bufOffset+6]}, 0))
		switch {
		case buf[bufOffset] == videoTag:
			lastVT = timeStamp
		case buf[bufOffset] == audioTag:
			lastAT = timeStamp
		default:
		}
		if lastAT != 0 && lastVT != 0 {
			diff := math.Abs(float64(lastVT - lastAT))
			if diff > t.Diff {
				err = fmt.Errorf("时间戳不匹配 %v %v (或许应调整flv音视频时间戳容差ms>%f)", lastVT, lastAT, diff)
				return
			}
		}

		if buf[bufOffset] == videoTag && buf[bufOffset+11]&0xf0 == 0x10 { //key frame
			if keyframeOp >= 0 {
				dropOffset = bufOffset
				if len(w) > 0 {
					_, err = w[0].Write(buf[keyframeOp:bufOffset])
				}
				return
			}
			dropT = timeStamp
			keyframeOp = bufOffset
		}
		bufOffset += tagSizeCheck + previouTagSize
	}

	return
}

func (t *FlvDecoder) Cut(reader io.Reader, startT, duration time.Duration, w io.Writer) (err error) {
	buf := make([]byte, humanize.MByte)
	buff := slice.New[byte](10 * humanize.MByte)
	skiped := false
	startTM := startT.Milliseconds()
	durationM := duration.Milliseconds()
	firstFT := -1
	for c := 0; err == nil; c++ {
		n, e := reader.Read(buf)
		if n == 0 && errors.Is(e, io.EOF) {
			return io.EOF
		}
		err = buff.Append(buf[:n])

		if !t.init {
			if frontBuf, dropOffset, e := t.InitFlv(buf); e != nil {
				return perrors.New("InitFlv", e.Error())
			} else {
				if dropOffset != 0 {
					_ = buff.RemoveFront(dropOffset)
				}
				if len(frontBuf) == 0 {
					continue
				} else {
					_, err = w.Write(frontBuf)
				}
			}
		} else if !skiped {
			if dropT, dropOffset, e := t.oneF(buff.GetPureBuf()); e != nil {
				return perrors.New("skip", e.Error())
			} else {
				if dropOffset != 0 {
					_ = buff.RemoveFront(dropOffset)
				}
				if firstFT == -1 {
					firstFT = dropT
				} else if startTM < int64(dropT-firstFT) {
					skiped = true
				}
			}
		} else {
			if dropT, dropOffset, e := t.oneF(buff.GetPureBuf(), w); e != nil {
				return perrors.New("w", e.Error())
			} else {
				if dropOffset != 0 {
					_ = buff.RemoveFront(dropOffset)
				}
				if durationM+startTM < int64(dropT-firstFT) {
					return
				}
			}
		}
	}
	return
}

func (t *FlvDecoder) Parse(buf []byte, keyframe *slice.Buf[byte]) (frontBuf []byte, dropOffset int, err error) {
	if !t.init {
		frontBuf, dropOffset, err = t.InitFlv(buf)
	} else {
		dropOffset, err = t.SearchStreamTag(buf, keyframe)
	}
	return
}
