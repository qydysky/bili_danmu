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

	ErrNoFoundFlvHeader = errors.New("ErrNoFoundFlvHeader")
	ErrNoFoundTagHeader = errors.New("ErrNoFoundTagHeader")
	ErrTagSizeZero      = errors.New("ErrTagSizeZero")
	ErrStreamId         = errors.New("ErrStreamId")
	ErrTagSize          = errors.New("ErrTagSize")
	ErrSignLost         = errors.New("ErrSignLost")

	ActionInitFlv        perrors.Action = `InitFlv`
	ActionGetIndexFlv    perrors.Action = `GetIndexFlv`
	ActionGenFastSeedFlv perrors.Action = `GenFastSeedFlv`
	ActionSeekFlv        perrors.Action = `SeekFlv`
	ActionOneFFlv        perrors.Action = `OneFFlv`
)

type FlvDecoder struct {
	Diff float64
}

func NewFlvDecoder() *FlvDecoder {
	return &FlvDecoder{Diff: 100}
}

func (t *FlvDecoder) InitFlv(buf []byte) (frontBuf []byte, dropOffset int, err error) {

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

		if buf[bufOffset]&0b11000000 != 0 ||
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

		tagSize := int(F.Btoi32v2(buf[bufOffset+1:bufOffset+4], 0))
		if tagSize == 0 {
			err = ErrTagSizeZero
			return
		}

		if bufOffset+tagHeaderSize+tagSize+previouTagSize > len(buf) {
			return
		}

		tagSizeCheck := int(F.Btoi32v2(buf[bufOffset+tagHeaderSize+tagSize:bufOffset+tagHeaderSize+tagSize+previouTagSize], 0))
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
				return
			}
		}
	}
	return
}

// this fuction read []byte and return flv header and all complete keyframe if possible.
// complete keyframe means the video and audio tags between two video key frames tag
func (t *FlvDecoder) SearchStreamTag(buf []byte, keyframe *slice.Buf[byte]) (dropOffset int, err error) {

	keyframe.Reset()

	var (
		keyframeOp = -1
		lastAT     int
		lastVT     int
	)

	for bufOffset := 0; bufOffset >= 0 && bufOffset+tagHeaderSize < len(buf); {

		if buf[bufOffset]&0b11000000 != 0 ||
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

		tagSize := int(F.Btoi32v2(buf[bufOffset+1:bufOffset+4], 0))
		if tagSize == 0 {
			err = ErrTagSizeZero
			return
		}
		if bufOffset+tagHeaderSize+tagSize+previouTagSize > len(buf) {
			return
		}

		tagSizeCheck := int(F.Btoi32v2(buf[bufOffset+tagHeaderSize+tagSize:bufOffset+tagHeaderSize+tagSize+previouTagSize], 0))
		if tagSizeCheck != tagSize+tagHeaderSize {
			err = ErrTagSize
			return
		}

		timeStamp := int(F.Btoi32v2([]byte{buf[bufOffset+7], buf[bufOffset+4], buf[bufOffset+5], buf[bufOffset+6]}, 0))
		switch  buf[bufOffset]{
		case videoTag:
			lastVT = timeStamp
		case audioTag:
			lastAT = timeStamp
		default:
		}
		if lastAT != 0 && lastVT != 0 {
			diff := math.Abs(float64(lastVT - lastAT))
			if diff > t.Diff {
				err = fmt.Errorf("时间戳不匹配 lastVT(%v) lastAT(%v) (或许应调整flv音视频时间戳容差ms>%f)", lastVT, lastAT, diff)
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

type dealFFlv func(t int, index int, buf []byte) error

func (t *FlvDecoder) oneF(buf []byte, w ...dealFFlv) (dropOffset int, err error) {

	var (
		keyframeOp = -1
		lastAT     int
		lastVT     int
	)

	for bufOffset := 0; bufOffset >= 0 && bufOffset+tagHeaderSize < len(buf); {

		if buf[bufOffset]&0b11000000 != 0 ||
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

		tagSize := int(F.Btoi32v2(buf[bufOffset+1:bufOffset+4], 0))
		if tagSize == 0 {
			err = ErrTagSizeZero
			return
		}
		if bufOffset+tagHeaderSize+tagSize+previouTagSize > len(buf) {
			return
		}

		tagSizeCheck := int(F.Btoi32v2(buf[bufOffset+tagHeaderSize+tagSize:bufOffset+tagHeaderSize+tagSize+previouTagSize], 0))
		if tagSizeCheck != tagSize+tagHeaderSize {
			err = ErrTagSize
			return
		}

		timeStamp := int(F.Btoi32v2([]byte{buf[bufOffset+7], buf[bufOffset+4], buf[bufOffset+5], buf[bufOffset+6]}, 0))
		switch  buf[bufOffset]{
		case videoTag:
			lastVT = timeStamp
		case audioTag:
			lastAT = timeStamp
		default:
		}
		if lastAT != 0 && lastVT != 0 {
			diff := math.Abs(float64(lastVT - lastAT))
			if diff > t.Diff {
				err = fmt.Errorf("时间戳不匹配 lastVT(%v) lastAT(%v) (或许应调整flv音视频时间戳容差ms>%f)", lastVT, lastAT, diff)
				return
			}
		}

		if buf[bufOffset] == videoTag && buf[bufOffset+11]&0xf0 == 0x10 { //key frame
			if keyframeOp >= 0 && len(w) > 0 {
				dropOffset = bufOffset
				err = w[0](timeStamp, keyframeOp, buf[keyframeOp:bufOffset])
				return
			}
			keyframeOp = bufOffset
		}
		bufOffset += tagSizeCheck + previouTagSize
	}

	return
}

// Deprecated: 效率低于GenFastSeed+CutSeed
func (t *FlvDecoder) Cut(reader io.Reader, startT, duration time.Duration, w io.Writer) (err error) {
	return t.CutSeed(reader, startT, duration, w, nil, nil)
}

func (t *FlvDecoder) CutSeed(reader io.Reader, startT, duration time.Duration, w io.Writer, seeker io.Seeker, getIndex func(seedTo time.Duration) (int64, error)) (err error) {
	buf := make([]byte, humanize.KByte*500)
	buff := slice.New[byte]()
	over := false
	seek := false
	init := false
	startTM := startT.Milliseconds()
	durationM := duration.Milliseconds()
	firstFT := -1

	wf := func(t int, index int, buf []byte) (e error) {
		if firstFT == -1 {
			firstFT = t
		}
		cu := int64(t - firstFT)
		over = duration != 0 && cu > durationM+startTM
		if startTM <= cu && !over {
			_, e = w.Write(buf)
		}
		return
	}

	for c := 0; err == nil && !over; c++ {
		n, e := reader.Read(buf)
		if n == 0 && errors.Is(e, io.EOF) {
			return io.EOF
		}
		err = buff.Append(buf[:n])

		if !init {
			if frontBuf, dropOffset, e := t.InitFlv(buff.GetPureBuf()); e != nil {
				return perrors.New(e.Error(), ActionInitFlv)
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
			init = true
		} else {
			if !seek && seeker != nil && getIndex != nil {
				if index, e := getIndex(startT); e != nil {
					return perrors.New(e.Error(), ActionGetIndexFlv)
				} else {
					if _, e := seeker.Seek(index, io.SeekStart); e != nil {
						return perrors.New(e.Error(), ActionSeekFlv)
					}
				}
				seek = true
				startTM = 0
				buff.Reset()
			}
			for {
				if dropOffset, e := t.oneF(buff.GetPureBuf(), wf); e != nil {
					return perrors.New(e.Error(), ActionOneFFlv)
				} else {
					if dropOffset != 0 {
						_ = buff.RemoveFront(dropOffset)
					} else {
						break
					}
				}
			}
		}
	}
	return
}

func (t *FlvDecoder) GenFastSeed(reader io.Reader, save func(seedTo time.Duration, cuIndex int64) error) (err error) {
	totalRead := 0
	buf := make([]byte, humanize.KByte*500)
	buff := slice.New[byte]()
	over := false
	firstFT := -1
	init := false

	for c := 0; err == nil && !over; c++ {
		n, e := reader.Read(buf)
		if n == 0 && errors.Is(e, io.EOF) {
			return io.EOF
		}
		totalRead += n
		err = buff.Append(buf[:n])

		if !init {
			if frontBuf, dropOffset, e := t.InitFlv(buff.GetPureBuf()); e != nil {
				return perrors.New(e.Error(), ActionInitFlv)
			} else {
				if dropOffset != 0 {
					_ = buff.RemoveFront(dropOffset)
				}
				if len(frontBuf) == 0 {
					continue
				}
			}
			init = true
		} else {
			for {
				if dropOffset, e := t.oneF(buff.GetPureBuf(), func(t, index int, buf []byte) error {
					if firstFT == -1 {
						firstFT = t
					}
					if e := save(time.Millisecond*time.Duration(t-firstFT), int64(totalRead-buff.Size()+index)); e != nil {
						return perrors.Join(ActionGenFastSeedFlv, e)
					}
					return nil
				}); e != nil {
					return perrors.Join(ActionGenFastSeedFlv, ActionOneFFlv, e)
				} else {
					if dropOffset != 0 {
						_ = buff.RemoveFront(dropOffset)
					} else {
						break
					}
				}
			}
		}
	}
	return
}
