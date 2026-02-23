package decoder

import (
	"errors"
	"fmt"
	"io"
	"math"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
	F "github.com/qydysky/bili_danmu/F"
	pe "github.com/qydysky/part/errors"
	pio "github.com/qydysky/part/io"
	pool "github.com/qydysky/part/pool"
	slice "github.com/qydysky/part/slice"
	unsafe "github.com/qydysky/part/unsafe"
)

var (
	fmp4BoxLenSize  = 4
	fmp4BoxNameSize = 4
)

var (
	ActionInitFmp4        pe.Action = `InitFmp4`
	ActionGetIndexFmp4    pe.Action = `GetIndexFmp4`
	ActionGenFastSeedFmp4 pe.Action = `GenFastSeedFmp4`
	ActionSeekFmp4        pe.Action = `SeekFmp4`
	ActionOneFFmp4        pe.Action = `OneFFmp4`
	ActionCheckTFail      pe.Action = `CheckTFail`
)

var boxs map[string]bool

func init() {
	boxs = make(map[string]bool)
	//isPureBox? || need to skip?
	boxs["ftyp"] = true
	boxs["moov"] = false
	boxs["mvhd"] = true
	boxs["trak"] = false
	boxs["tkhd"] = true
	boxs["mdia"] = false
	boxs["mdhd"] = true
	boxs["hdlr"] = true
	boxs["minf"] = false || true
	boxs["mvex"] = false || true
	boxs["moof"] = false
	boxs["mfhd"] = true
	boxs["traf"] = false
	boxs["tfhd"] = true
	boxs["tfdt"] = true
	boxs["trun"] = true
	boxs["mdat"] = true
}

type ie struct {
	n string // box name
	i int    // start index(box size include)
	e int    // end index
}

type trak struct {
	// firstTimeStamp int
	// lastTimeStamp  int
	timescale   int
	trackID     int
	handlerType byte
}

type timeStamp struct {
	// firstTimeStamp int
	timeStamp   int
	timescale   int
	data        []byte
	handlerType byte
}

// func (t *timeStamp) resetTs() {
// 	t.timeStamp -= t.firstTimeStamp
// 	switch len(t.data) {
// 	case 4:
// 		copy(t.data, F.Itob32(int32(t.timeStamp)))
// 	case 8:
// 		copy(t.data, F.Itob64(int64(t.timeStamp)))
// 	}
// }

func (t *timeStamp) getT() float64 {
	return float64(t.timeStamp) / float64(t.timescale)
}

type Fmp4Decoder struct {
	traks   map[int]*trak
	buf     *slice.Buf[byte]
	buflock sync.Mutex

	AVTDiff float64 // 音视频时间戳容差
	Debug   bool
}

var Fmp4DecoderPool = pool.New(pool.PoolFunc[Fmp4Decoder]{
	New: func() (fd *Fmp4Decoder) {
		fd = &Fmp4Decoder{
			traks: make(map[int]*trak),
			buf:   slice.New[byte](),
		}
		fd.buf.ExpandCapTo(humanize.MByte * 5)
		return
	},
	Reuse: func(fd *Fmp4Decoder) *Fmp4Decoder {
		clear(fd.traks)
		fd.buf.Reset()
		return fd
	},
}, -1)

// Deprecated:use Fmp4DecoderPool
// func NewFmp4Decoder() (fd *Fmp4Decoder) {
// 	return Fmp4DecoderPool.Get()
// }

// Deprecated:use Fmp4DecoderPool
// func NewFmp4DecoderWithBufsize(size int) *Fmp4Decoder {
// 	f := Fmp4DecoderPool.Get()
// 	return f
// }

func (t *Fmp4Decoder) Init(buf []byte) (b []byte, dropOffset int, err error) {
	var ftypI, ftypE, moovI, moovE int

	ies, recycle, e := decode(buf)
	defer recycle(ies)
	if e != nil {
		return
	}

	err = deal(ies, dealIE{
		boxNames: []string{"ftyp", "moov"},
		fs: func(m []ie) error {
			ftypI = m[0].i
			ftypE = m[0].e
			moovI = m[1].i
			moovE = m[1].e
			return nil
		},
	})

	if err != nil {
		return
	}

	err = deal(ies, dealIE{
		boxNames: []string{"tkhd", "mdia", "mdhd", "hdlr"},
		fs: func(m []ie) error {
			tackId := int(F.Btoiv2(buf, m[0].i+20, 4))
			t.traks[tackId] = &trak{
				trackID: tackId,
				// firstTimeStamp: -1,
				// lastTimeStamp:  -1,
				timescale:   int(F.Btoiv2(buf, m[2].i+20, 4)),
				handlerType: buf[m[3].i+16],
			}
			return nil
		},
	})

	if err != nil {
		return
	}

	if len(t.traks) == 0 {
		err = errors.New("未找到任何trak包")
		return
	}

	b = make([]byte, ftypE-ftypI+moovE-moovI)
	copy(b[:ftypE-ftypI], buf[ftypI:ftypE])
	copy(b[ftypE-ftypI:], buf[moovI:moovE])
	return
}

var (
	ErrBufTooLarge = errors.New("ErrBufTooLarge")
	ErrMisTraks    = errors.New("ErrMisTraks")
)

func (t *Fmp4Decoder) SearchStreamFrame(buf []byte, keyframe *slice.Buf[byte]) (cu int, err error) {
	if len(buf) > humanize.MByte*100 {
		return 0, ErrBufTooLarge
	}
	if len(t.traks) == 0 {
		return 0, ErrMisTraks
	}

	t.buflock.Lock()
	defer t.buflock.Unlock()

	t.buf.Reset()
	keyframe.Reset()

	defer func() {
		if err != nil {
			keyframe.Reset()
			cu = 0
		}
	}()

	var (
		haveKeyframe bool
		bufModified  = t.buf.GetModified()
		// maxSequenceNumber int //有时并不是单调增加
		maxVT float64
		maxAT float64

		//get timeStamp
		get_timeStamp = func(tfdt int) (ts timeStamp) {
			switch buf[tfdt+8] {
			case 0:
				ts.data = buf[tfdt+16 : tfdt+20]
				ts.timeStamp = int(F.Btoi32v2(buf, tfdt+16))
			case 1:
				ts.data = buf[tfdt+12 : tfdt+20]
				ts.timeStamp = int(F.Btoi64(buf, tfdt+12))
			}
			return
		}

		//get track type
		get_track_type = func(tfhd, tfdt int) (ts timeStamp, handlerType byte) {
			track, ok := t.traks[int(F.Btoiv2(buf, tfhd+12, 4))]
			if ok {
				ts := get_timeStamp(tfdt)
				// if track.firstTimeStamp == -1 {
				// 	track.firstTimeStamp = ts.timeStamp
				// }

				// ts.firstTimeStamp = track.firstTimeStamp
				ts.handlerType = track.handlerType
				ts.timescale = track.timescale

				// if ts.timeStamp > track.lastTimeStamp {
				// 	track.lastTimeStamp = ts.timeStamp
				// 	ts.resetTs()
				// }

				return ts, track.handlerType
			}
			return
		}

		// getSampleEntries = func(trun, mdat int) (sampleChunkI, sampleChunkE int64) {
		// 	sampleChunkI = F.Btoiv2(buf, trun+16, 4)
		// 	sampleChunkE = sampleChunkI
		// 	// https://blog.qydysky.org/post/20221025fmp4/#trun
		// 	if buf[trun+10]&0x02 == 0x02 {
		// 		for i := trun + 24; i < mdat; i += 12 {
		// 			sampleChunkE += F.Btoiv2(buf, i+4, 4)
		// 		}
		// 	}
		// 	return
		// }

		//is t error?
		checkAndSetMaxT = func(ts timeStamp) (err error) {
			switch ts.handlerType {
			case 'v':
				if maxVT == 0 {
					maxVT = ts.getT()
				} else if maxVT == ts.getT() {
					err = ActionCheckTFail.New("equal VT detect")
				} else if maxVT > ts.getT() {
					err = ActionCheckTFail.New("lower VT detect")
				} else {
					maxVT = ts.getT()
				}
			case 'a':
				if maxAT == 0 {
					maxAT = ts.getT()
				} else if maxAT == ts.getT() {
					err = ActionCheckTFail.New("equal AT detect")
				} else if maxAT > ts.getT() {
					err = ActionCheckTFail.New("lower AT detect")
				} else {
					maxAT = ts.getT()
				}
			default:
			}
			return
		}

		dropKeyFrame = func(index int) {
			t.buf.Reset()
			haveKeyframe = false
			cu = index
		}
	)

	ies, recycle, e := decode(buf)
	defer recycle(ies)
	if e != nil {
		return 0, e
	}

	err = deals(ies,
		[]dealIE{
			{
				boxNames: []string{"moof", "mfhd", "traf", "tfhd", "tfdt", "trun", "mdat"},
				fs: func(m []ie) error {
					var (
						keyframeMoof = buf[m[5].i+20] == byte(0x02)
						// moofSN       = int(F.Btoiv2(buf, m[1].i+12, 4))
					)

					{
						ts, _ := get_track_type(m[3].i, m[4].i)
						// if ts.handlerType == 'v' {
						// 	if e := getSampleEntries(m[5].i, m[6].i); e != nil {
						// 		//skip
						// 		dropKeyFrame(m[0].e)
						// 		return pe.Join(ErrDecode, e)
						// 	}
						// }
						if e := checkAndSetMaxT(ts); e != nil {
							dropKeyFrame(m[0].e)
							return pe.Join(ErrDecode, e)
						}
					}

					// fmt.Println(ts.getT(), "frame0", keyframeMoof, t.buf.size(), m[0].i, m[6].n, m[6].e)

					//deal frame
					if keyframeMoof {
						if v, e := t.buf.HadModified(bufModified); e == nil && v && !t.buf.IsEmpty() {
							if e := t.buf.AppendTo(keyframe); e != nil {
								return e
							}
							dropKeyFrame(m[0].i)
						}
						haveKeyframe = true
					} else if !haveKeyframe {
						cu = m[6].e
					}
					if haveKeyframe {
						if e := t.buf.Append(buf[m[0].i:m[6].e]); e != nil {
							return e
						}
					}
					return nil
				},
			},
			{
				boxNames: []string{"moof", "mfhd", "traf", "tfhd", "tfdt", "trun", "traf", "tfhd", "tfdt", "trun", "mdat"},
				fs: func(m []ie) error {
					var (
						keyframeMoof = buf[m[5].i+20] == byte(0x02) || buf[m[9].i+20] == byte(0x02)
						// moofSN       = int(F.Btoiv2(buf, m[1].i+12, 4))
						video timeStamp
						audio timeStamp
					)

					// fmt.Println(moofSN, "frame1", keyframeMoof, t.buf.size(), m[0].i, m[10].n, m[10].e)

					{
						ts, handlerType := get_track_type(m[3].i, m[4].i)
						// if handlerType == 'v' {
						// 	if e := checkSampleEntries(m[5].i, m[6].i); e != nil {
						// 		//skip
						// 		dropKeyFrame(m[0].e)
						// 		return pe.Join(ErrDecode, e)
						// 	}
						// }
						switch handlerType {
						case 'v':
							video = ts
						case 's':
							audio = ts
						}
						if e := checkAndSetMaxT(ts); e != nil {
							dropKeyFrame(m[0].e)
							return pe.Join(ErrDecode, e)
						}
					}
					{
						ts, handlerType := get_track_type(m[7].i, m[8].i)
						// if handlerType == 'v' {
						// 	if e := checkSampleEntries(m[9].i, m[10].i); e != nil {
						// 		//skip
						// 		dropKeyFrame(m[0].e)
						// 		return pe.Join(ErrDecode, e)
						// 	}
						// }
						switch handlerType {
						case 'v':
							video = ts
						case 's':
							audio = ts
						}
						if e := checkAndSetMaxT(ts); e != nil {
							dropKeyFrame(m[0].e)
							return pe.Join(ErrDecode, e)
						}
					}

					//sync audio timeStamp
					if t.AVTDiff <= 0.1 {
						t.AVTDiff = 0.1
					}
					if diff := math.Abs(video.getT() - audio.getT()); diff > t.AVTDiff {
						return pe.Join(ErrDecode, fmt.Errorf("时间戳不匹配 lastVT(%v) lastAT(%v) (或许应调整fmp4音视频时间戳容差s>%.2f)", video.timeStamp, audio.timeStamp, diff))
						// copy(video.data, F.Itob64(int64(audio.getT()*float64(video.timescale))))
					}

					//deal frame
					if keyframeMoof {
						if v, e := t.buf.HadModified(bufModified); e == nil && v && !t.buf.IsEmpty() {
							if e := t.buf.AppendTo(keyframe); e != nil {
								return e
							}
							dropKeyFrame(m[0].i)
						}
						haveKeyframe = true
					} else if !haveKeyframe {
						cu = m[10].e
					}
					if haveKeyframe {
						if e := t.buf.Append(buf[m[0].i:m[10].e]); e != nil {
							return e
						}
					}
					return nil
				},
			},
		})
	return
}

type dealFMp4 func(t float64, index int, re io.Reader) error

func (t *Fmp4Decoder) oneF(buf []byte, w ...dealFMp4) (cu int, err error) {
	if len(buf) > humanize.MByte*100 {
		return 0, ErrBufTooLarge
	}
	if len(t.traks) == 0 {
		return 0, ErrMisTraks
	}

	sbuf := slice.NewSliceIndexNoLock(buf)

	defer func() {
		if err != nil {
			cu = 0
		}
	}()

	var (
		haveKeyframe bool
		bufModified  = sbuf.GetModified()
		// maxSequenceNumber int //有时并不是单调增加
		maxVT float64
		maxAT float64

		//get timeStamp
		get_timeStamp = func(tfdt int) (ts timeStamp) {
			switch buf[tfdt+8] {
			case 0:
				ts.data = buf[tfdt+16 : tfdt+20]
				ts.timeStamp = int(F.Btoi32v2(buf, tfdt+16))
			case 1:
				ts.data = buf[tfdt+12 : tfdt+20]
				ts.timeStamp = int(F.Btoi64(buf, tfdt+12))
			}
			return
		}

		//get track type
		get_track_type = func(tfhd, tfdt int) (ts timeStamp, handlerType byte) {
			track, ok := t.traks[int(F.Btoiv2(buf, tfhd+12, 4))]
			if ok {
				ts := get_timeStamp(tfdt)
				// if track.firstTimeStamp == -1 {
				// 	track.firstTimeStamp = ts.timeStamp
				// }

				// ts.firstTimeStamp = track.firstTimeStamp
				ts.handlerType = track.handlerType
				ts.timescale = track.timescale

				// if ts.timeStamp > track.lastTimeStamp {
				// 	track.lastTimeStamp = ts.timeStamp
				// 	ts.resetTs()
				// }

				return ts, track.handlerType
			}
			return
		}

		//is SampleEntries error?
		checkSampleEntries = func(trun, mdat int) error {
			if buf[trun+11] == 'b' {
				for i := trun + 24; i < mdat; i += 12 {
					if F.Btoiv2(buf, i+4, 4) < 1000 {
						return errors.New("find sample size less then 1000")
					}
				}
			}
			return nil
		}

		//is t error?
		checkAndSetMaxT = func(ts timeStamp) (err error) {
			switch ts.handlerType {
			case 'v':
				if maxVT == 0 {
					maxVT = ts.getT()
				} else if maxVT == ts.getT() {
					err = ActionCheckTFail.New("equal VT detect")
				} else if maxVT > ts.getT() {
					err = ActionCheckTFail.New("lower VT detect")
				} else {
					maxVT = ts.getT()
				}
			case 'a':
				if maxAT == 0 {
					maxAT = ts.getT()
				} else if maxAT == ts.getT() {
					err = ActionCheckTFail.New("equal AT detect")
				} else if maxAT > ts.getT() {
					err = ActionCheckTFail.New("lower AT detect")
				} else {
					maxAT = ts.getT()
				}
			default:
			}
			return
		}

		dropKeyFrame = func(index int) {
			sbuf.Reset()
			haveKeyframe = false
			cu = index
		}
	)

	ies, recycle, e := decode(buf)
	defer recycle(ies)
	if e != nil {
		return 0, e
	}

	var ErrNormal pe.Action = "ErrNormal"

	err = deals(ies,
		[]dealIE{
			{
				boxNames: []string{"moof", "mfhd", "traf", "tfhd", "tfdt", "trun", "mdat"},
				fs: func(m []ie) error {
					var (
						keyframeMoof = buf[m[5].i+20] == byte(0x02)
						// moofSN       = int(F.Btoiv2(buf, m[1].i+12, 4))
						video timeStamp
					)

					{
						ts, handlerType := get_track_type(m[3].i, m[4].i)
						if ts.handlerType == 'v' {
							if e := checkSampleEntries(m[5].i, m[6].i); e != nil {
								//skip
								dropKeyFrame(m[0].e)
								return pe.Join(ErrDecode, e)
							}
						}
						if handlerType == 'v' {
							video = ts
						}
						if e := checkAndSetMaxT(ts); e != nil {
							dropKeyFrame(m[0].e)
							return pe.Join(ErrDecode, e)
						}
					}

					//deal frame
					if keyframeMoof {
						if v, e := sbuf.HadModified(bufModified); e == nil && v && !sbuf.IsEmpty() {
							if haveKeyframe && len(w) > 0 {
								// pio.WrapIoWriteTo 实现io.WriteTo避免分配复制缓存
								err = w[0](video.getT(), cu, pio.WrapIoWriteTo(sbuf))
								dropKeyFrame(m[0].i)
								return ErrNormal
							}
							dropKeyFrame(m[0].i)
						}
						haveKeyframe = true
					} else if !haveKeyframe {
						cu = m[6].e
					}
					if haveKeyframe {
						sbuf.Append(m[0].i, m[6].e)
					}
					return nil
				},
			},
			{
				boxNames: []string{"moof", "mfhd", "traf", "tfhd", "tfdt", "trun", "traf", "tfhd", "tfdt", "trun", "mdat"},
				fs: func(m []ie) error {
					var (
						keyframeMoof = buf[m[5].i+20] == byte(0x02) || buf[m[9].i+20] == byte(0x02)
						// moofSN       = int(F.Btoiv2(buf, m[1].i+12, 4))
						video timeStamp
						audio timeStamp
					)

					{
						ts, handlerType := get_track_type(m[3].i, m[4].i)
						if handlerType == 'v' {
							if e := checkSampleEntries(m[5].i, m[6].i); e != nil {
								//skip
								dropKeyFrame(m[0].e)
								return pe.Join(ErrDecode, e)
							}
						}
						switch handlerType {
						case 'v':
							video = ts
						case 's':
							audio = ts
						}
						if e := checkAndSetMaxT(ts); e != nil {
							dropKeyFrame(m[0].e)
							return pe.Join(ErrDecode, e)
						}
					}
					{
						ts, handlerType := get_track_type(m[7].i, m[8].i)
						if handlerType == 'v' {
							if e := checkSampleEntries(m[9].i, m[10].i); e != nil {
								//skip
								dropKeyFrame(m[0].e)
								return pe.Join(ErrDecode, e)
							}
						}
						switch handlerType {
						case 'v':
							video = ts
						case 's':
							audio = ts
						}
						if e := checkAndSetMaxT(ts); e != nil {
							dropKeyFrame(m[0].e)
							return pe.Join(ErrDecode, e)
						}
					}

					//sync audio timeStamp
					if t.AVTDiff <= 0.1 {
						t.AVTDiff = 0.1
					}
					if diff := math.Abs(video.getT() - audio.getT()); diff > t.AVTDiff {
						return pe.Join(ErrDecode, fmt.Errorf("时间戳不匹配 lastVT(%v) lastAT(%v) (或许应调整fmp4音视频时间戳容差s>%.2f)", video.timeStamp, audio.timeStamp, diff))
						// copy(video.data, F.Itob64(int64(audio.getT()*float64(video.timescale))))
					}

					//deal frame
					if keyframeMoof {
						if v, e := sbuf.HadModified(bufModified); e == nil && v && !sbuf.IsEmpty() {
							if haveKeyframe && len(w) > 0 {
								// pio.WrapIoWriteTo 实现io.WriteTo避免分配复制缓存
								err = w[0](video.getT(), cu, pio.WrapIoWriteTo(sbuf))
								dropKeyFrame(m[0].i)
								return ErrNormal
							}
							dropKeyFrame(m[0].i)
						}
						haveKeyframe = true
					} else if !haveKeyframe {
						cu = m[10].e
					}
					if haveKeyframe {
						sbuf.Append(m[0].i, m[10].e)
					}
					return nil
				},
			},
		},
	)

	if errors.Is(err, ErrNormal) {
		err = nil
	}

	return
}

// Deprecated:效率低于GenFastSeed+CutSeed
func (t *Fmp4Decoder) Cut(reader io.Reader, startT, duration time.Duration, w io.Writer, skipHeader, writeLastBuf bool) (err error) {
	return t.CutSeed(reader, startT, duration, w, nil, nil, skipHeader, writeLastBuf)
}

func (t *Fmp4Decoder) CutSeed(reader io.Reader, startT, duration time.Duration, w io.Writer, seeker io.Seeker, getIndex func(seedTo time.Duration) (int64, error), skipHeader, writeLastBuf bool) (err error) {
	init := false
	seek := false
	over := false
	startTM := startT.Seconds()
	durationM := duration.Seconds()
	firstFT := -1.0

	t.buflock.Lock()
	defer t.buflock.Unlock()
	t.buf.Reset()

	wf := func(t float64, index int, re io.Reader) (e error) {
		if firstFT == -1 {
			firstFT = t
		}
		cu := t - firstFT
		over = duration != 0 && cu > durationM+startTM
		if startTM <= cu && !over {
			_, e = io.Copy(w, re)
		}
		return
	}

	if t.Debug {
		fmt.Printf("cut startT: %v duration: %v\n", startT, duration)
	}
	for c := 0; err == nil && !over; c++ {
		if t.buf.Size() == t.buf.Cap() {
			return ErrBufOverflow
		}
		n, e := reader.Read(t.buf.GetRawBuf(t.buf.Size(), min(t.buf.Size()+humanize.MByte, t.buf.Cap())))
		t.buf.AddSize(n)
		if n == 0 && errors.Is(e, io.EOF) {
			if t.buf.Size() > 0 {
				_, _ = w.Write(t.buf.GetPureBuf())
				t.buf.Reset()
			}
			return io.EOF
		}

		if !init {
			if frontBuf, dropOffset, e := t.Init(t.buf.GetPureBuf()); e != nil {
				return pe.New(e.Error(), ActionInitFmp4)
			} else {
				if len(frontBuf) != 0 {
					init = true
					if !skipHeader {
						if t.Debug {
							fmt.Printf("write frontBuf: frontBufSize: %d\n", len(frontBuf))
						}
						_, err = w.Write(frontBuf)
					}
				}
				if dropOffset > 0 {
					t.buf.RemoveFront(dropOffset)
				}
			}
		} else {
			if !seek && seeker != nil && getIndex != nil {
				if index, e := getIndex(startT); e != nil {
					return pe.New(e.Error(), ActionGetIndexFmp4)
				} else {
					if _, e := seeker.Seek(index, io.SeekStart); e != nil {
						return pe.New(e.Error(), ActionSeekFmp4)
					}
				}
				seek = true
				startTM = 0
				t.buf.Reset()
			}
			for {
				if dropOffset, e := t.oneF(t.buf.GetPureBuf(), wf); e != nil {
					return pe.New(e.Error(), ActionOneFFmp4)
				} else {
					if dropOffset != 0 {
						_ = t.buf.RemoveFront(dropOffset)
					} else {
						break
					}
				}
			}
		}
	}
	return
}

func (t *Fmp4Decoder) GenFastSeed(reader io.Reader, save func(seedTo time.Duration, cuIndex int64) error) (err error) {
	t.buflock.Lock()
	defer t.buflock.Unlock()

	t.buf.Reset()

	totalRead := 0
	init := false
	firstFT := -1.0

	for c := 0; err == nil; c++ {
		if t.buf.Size() == t.buf.Cap() {
			return ErrBufOverflow
		}
		n, e := reader.Read(t.buf.GetRawBuf(t.buf.Size(), min(t.buf.Size()+humanize.MByte, t.buf.Cap())))
		t.buf.AddSize(n)
		if n == 0 && errors.Is(e, io.EOF) {
			return io.EOF
		}
		totalRead += int(n)
		if !init {
			if frontBuf, dropOffset, e := t.Init(t.buf.GetPureBuf()); e != nil {
				return pe.New(e.Error(), ActionInitFmp4)
			} else if len(frontBuf) == 0 {
				continue
			} else {
				_ = t.buf.RemoveFront(dropOffset)
				init = true
			}
		} else {
			for {
				if dropOffset, e := t.oneF(t.buf.GetPureBuf(), func(ts float64, index int, _ io.Reader) error {
					if firstFT == -1 {
						firstFT = ts
					}
					if e := save(time.Second*time.Duration(ts-firstFT), int64(totalRead-t.buf.Size()+index)); e != nil {
						return pe.Join(ActionGenFastSeedFmp4, e)
					}
					return nil
				}); e != nil {
					return pe.Join(ActionGenFastSeedFmp4, ActionOneFFmp4, e)
				} else {
					if dropOffset != 0 {
						_ = t.buf.RemoveFront(dropOffset)
					} else {
						break
					}
				}
			}
		}
	}
	return
}

type dealIE struct {
	matchCounts int
	boxNames    []string
	fs          func([]ie) (err error)
}

func (t *dealIE) deal(ies []ie, cu int) (err error) {
	if t.boxNames[t.matchCounts] == ies[cu].n {
		t.matchCounts += 1
		if t.matchCounts == len(t.boxNames) {
			t.matchCounts = 0
			return t.fs(ies[cu+1-len(t.boxNames) : cu+1])
		}
	} else {
		t.matchCounts = 0
	}
	return nil
}

func deal(ies *[]ie, dealIEf dealIE) (err error) {
	return deals(ies, []dealIE{dealIEf})
}

func deals(ies *[]ie, dealIEs []dealIE) (err error) {
	for cu := 0; cu < len(*ies) && len(dealIEs) != 0; cu++ {
		for i := 0; i < len(dealIEs); i++ {
			if e := dealIEs[i].deal(*ies, cu); e != nil {
				return e
			}
		}
	}
	return
}

var (
	ErrMisBox     = pe.New("decode", "ErrMisBox")
	ErrCantResync = pe.New("decode")
	iesPool       = pool.NewPoolBlocks[ie]()
)

// buf can not mod until recycle
func decode(buf []byte) (m *[]ie, recycle func(*[]ie), err error) {
	var cu int

	m = iesPool.Get()
	recycle = iesPool.Put
	*m = (*m)[:0]

	for cu < len(buf)-fmp4BoxLenSize-fmp4BoxNameSize {
		boxName, i, e, E := searchBox(buf, &cu)
		if E != nil {
			if errors.Is(E, io.EOF) {
				if len(*m) == 0 {
					err = ErrMisBox
				}
				return
			}
			err = E
			return
		}

		if cu := len(*m); cu < cap(*m) {
			*m = (*m)[:cu+1]
			(*m)[cu].n = boxName
			(*m)[cu].i = i
			(*m)[cu].e = e
		} else {
			*m = append(*m, ie{
				n: boxName,
				i: i,
				e: e,
			})
		}
	}

	return
}

var (
	ErrUnkownBox = pe.New("ErrUnkownBox")
)

// buf can not mod
func searchBox(buf []byte, cu *int) (boxName string, i int, e int, err error) {
	i = *cu
	e = i + int(F.Btoiv2(buf, *cu, fmp4BoxLenSize))
	boxName = unsafe.B2S(buf[*cu+fmp4BoxLenSize : *cu+fmp4BoxLenSize+fmp4BoxNameSize])
	isPureBoxOrNeedSkip, ok := boxs[boxName]
	if !ok {
		err = ErrUnkownBox.WithReason(fmt.Sprintf("未知包: hex(%x%x%x%x)", boxName[0], boxName[1], boxName[2], boxName[3]))
	} else if e > len(buf) {
		err = io.EOF
	} else if isPureBoxOrNeedSkip {
		*cu = e
	} else {
		*cu += 8
	}
	return
}

// func testBox(buf []byte, no string) {
// 	fmt.Println("testBox", "===>")
// 	err := deal(buf,
// 		[]string{"moof", "mfhd",
// 			"traf", "tfhd", "tfdt", "trun",
// 			"traf", "tfhd", "tfdt", "trun",
// 			"mdat"},
// 		func(m []*ie) bool {
// 			moofSN := int(F.Btoiv2(buf, m[1].i+12, 4))
// 			keyframeMoof := buf[m[5].i+20] == byte(0x02) || buf[m[9].i+20] == byte(0x02)
// 			fmt.Println(moofSN, "frame", keyframeMoof, m[0].i, m[10].n, m[10].e)
// 			return false
// 		})
// 	fmt.Println("err", err)
// 	fmt.Println("testBox", "<===")
// }
