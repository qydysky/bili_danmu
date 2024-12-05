package reply

import (
	"bytes"
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
	i int    // start index
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
	traks map[int]*trak
	buf   *slice.Buf[byte]

	AVTDiff float64 // 音视频时间戳容差
	Debug   bool
}

func NewFmp4Decoder() *Fmp4Decoder {
	return &Fmp4Decoder{
		traks: make(map[int]*trak),
		buf:   slice.New[byte](),
	}
}

func (t *Fmp4Decoder) Init_fmp4(buf []byte) (b []byte, err error) {
	var ftypI, ftypE, moovI, moovE int

	ies, e := decode(buf, "ftyp")
	if e != nil {
		return
	}

	err = deal(ies,
		[]string{"ftyp", "moov"},
		func(m []ie) (bool, error) {
			ftypI = m[0].i
			ftypE = m[0].e
			moovI = m[1].i
			moovE = m[1].e
			return true, nil
		})

	if err != nil {
		return nil, err
	}

	err = deal(ies,
		[]string{"tkhd", "mdia", "mdhd", "hdlr"},
		func(m []ie) (bool, error) {
			tackId := int(F.Btoi(buf, m[0].i+20, 4))
			t.traks[tackId] = &trak{
				trackID: tackId,
				// firstTimeStamp: -1,
				// lastTimeStamp:  -1,
				timescale:   int(F.Btoi(buf, m[2].i+20, 4)),
				handlerType: buf[m[3].i+16],
			}
			return false, nil
		})

	if err != nil {
		return nil, err
	}

	if len(t.traks) == 0 {
		return nil, errors.New("未找到任何trak包")
	}

	b = make([]byte, ftypE-ftypI+moovE-moovI)
	copy(b[:ftypE-ftypI], buf[ftypI:ftypE])
	copy(b[ftypE-ftypI:], buf[moovI:moovE])
	return b, nil
}

var (
	ErrBufTooLarge = errors.New("ErrBufTooLarge")
	ErrMisTraks    = errors.New("ErrMisTraks")
)

func (t *Fmp4Decoder) Search_stream_fmp4(buf []byte, keyframe *slice.Buf[byte]) (cu int, err error) {
	if len(buf) > humanize.MByte*100 {
		return 0, ErrBufTooLarge
	}
	if len(t.traks) == 0 {
		return 0, ErrMisTraks
	}
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
				ts.timeStamp = int(F.Btoi32(buf, tfdt+16))
			case 1:
				ts.data = buf[tfdt+12 : tfdt+20]
				ts.timeStamp = int(F.Btoi64(buf, tfdt+12))
			}
			return
		}

		//get track type
		get_track_type = func(tfhd, tfdt int) (ts timeStamp, handlerType byte) {
			track, ok := t.traks[int(F.Btoi(buf, tfhd+12, 4))]
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
					if F.Btoi(buf, i+4, 4) < 1000 {
						return errors.New("find sample size less then 1000")
					}
				}
			}
			return nil
		}

		//is t error?
		check_set_maxT = func(ts timeStamp, equal func(ts timeStamp) error, larger func(ts timeStamp) error) (err error) {
			switch ts.handlerType {
			case 'v':
				if maxVT == 0 {
					maxVT = ts.getT()
				} else if maxVT == ts.getT() && equal != nil {
					err = equal(ts)
				} else if maxVT > ts.getT() && larger != nil {
					err = larger(ts)
				} else {
					maxVT = ts.getT()
				}
			case 'a':
				if maxAT == 0 {
					maxAT = ts.getT()
				} else if maxAT == ts.getT() && equal != nil {
					err = equal(ts)
				} else if maxAT > ts.getT() && larger != nil {
					err = larger(ts)
				} else {
					maxAT = ts.getT()
				}
			default:
			}
			return
		}
	)

	ies, e := decode(buf, "moof")
	if e != nil {
		return 0, e
	}

	err = deals(ies,
		[][]string{
			{"moof", "mfhd", "traf", "tfhd", "tfdt", "trun", "mdat"},
			{"moof", "mfhd", "traf", "tfhd", "tfdt", "trun", "traf", "tfhd", "tfdt", "trun", "mdat"}},
		[]func(m []ie) (bool, error){
			func(m []ie) (bool, error) {
				var (
					keyframeMoof = buf[m[5].i+20] == byte(0x02)
					// moofSN       = int(F.Btoi(buf, m[1].i+12, 4))
				)

				{
					ts, _ := get_track_type(m[3].i, m[4].i)
					if ts.handlerType == 'v' {
						if e := checkSampleEntries(m[5].i, m[6].i); e != nil {
							//skip
							t.buf.Reset()
							haveKeyframe = false
							cu = m[0].i
							return false, e
						}
					}
					if e := check_set_maxT(ts, func(_ timeStamp) error {
						return errors.New("skip")
					}, func(_ timeStamp) error {
						t.buf.Reset()
						haveKeyframe = false
						cu = m[0].i
						return errors.New("skip")
					}); e != nil {
						return false, e
					}
				}

				// fmt.Println(ts.getT(), "frame0", keyframeMoof, t.buf.size(), m[0].i, m[6].n, m[6].e)

				//deal frame
				if keyframeMoof {
					if v, e := t.buf.HadModified(bufModified); e == nil && v && !t.buf.IsEmpty() {
						if e := t.buf.AppendTo(keyframe); e != nil {
							return false, e
						}
						cu = m[0].i
						t.buf.Reset()
					}
					haveKeyframe = true
				} else if !haveKeyframe {
					cu = m[6].e
				}
				if haveKeyframe {
					if e := t.buf.Append(buf[m[0].i:m[6].e]); e != nil {
						return false, e
					}
				}
				return false, nil
			},
			func(m []ie) (bool, error) {
				var (
					keyframeMoof = buf[m[5].i+20] == byte(0x02) || buf[m[9].i+20] == byte(0x02)
					// moofSN       = int(F.Btoi(buf, m[1].i+12, 4))
					video timeStamp
					audio timeStamp
				)

				// fmt.Println(moofSN, "frame1", keyframeMoof, t.buf.size(), m[0].i, m[10].n, m[10].e)

				{
					ts, handlerType := get_track_type(m[3].i, m[4].i)
					if handlerType == 'v' {
						if e := checkSampleEntries(m[5].i, m[6].i); e != nil {
							//skip
							t.buf.Reset()
							haveKeyframe = false
							cu = m[0].i
							return false, e
						}
					}
					switch handlerType {
					case 'v':
						video = ts
					case 's':
						audio = ts
					}
					if e := check_set_maxT(ts, func(_ timeStamp) error {
						return errors.New("skip")
					}, func(_ timeStamp) error {
						t.buf.Reset()
						haveKeyframe = false
						cu = m[0].i
						return errors.New("skip")
					}); e != nil {
						return false, e
					}
				}
				{
					ts, handlerType := get_track_type(m[7].i, m[8].i)
					if handlerType == 'v' {
						if e := checkSampleEntries(m[9].i, m[10].i); e != nil {
							//skip
							t.buf.Reset()
							haveKeyframe = false
							cu = m[0].i
							return false, e
						}
					}
					switch handlerType {
					case 'v':
						video = ts
					case 's':
						audio = ts
					}
					if e := check_set_maxT(ts, func(_ timeStamp) error {
						return errors.New("skip")
					}, func(_ timeStamp) error {
						t.buf.Reset()
						haveKeyframe = false
						cu = m[0].i
						return errors.New("skip")
					}); e != nil {
						return false, e
					}
				}

				//sync audio timeStamp
				if t.AVTDiff <= 0.1 {
					t.AVTDiff = 0.1
				}
				if diff := math.Abs(video.getT() - audio.getT()); diff > t.AVTDiff {
					return false, fmt.Errorf("时间戳不匹配 %v %v (或许应调整fmp4音视频时间戳容差s>%.2f)", video.timeStamp, audio.timeStamp, diff)
					// copy(video.data, F.Itob64(int64(audio.getT()*float64(video.timescale))))
				}

				//deal frame
				if keyframeMoof {
					if v, e := t.buf.HadModified(bufModified); e == nil && v && !t.buf.IsEmpty() {
						if e := t.buf.AppendTo(keyframe); e != nil {
							return false, e
						}
						cu = m[0].i
						t.buf.Reset()
					}
					haveKeyframe = true
				} else if !haveKeyframe {
					cu = m[10].e
				}
				if haveKeyframe {
					if e := t.buf.Append(buf[m[0].i:m[10].e]); e != nil {
						return false, e
					}
				}
				return false, nil
			}})
	return
}

func (t *Fmp4Decoder) oneF(buf []byte, ifWrite func(t float64) bool, w ...io.Writer) (cu int, err error) {
	if len(buf) > humanize.MByte*100 {
		return 0, ErrBufTooLarge
	}
	if len(t.traks) == 0 {
		return 0, ErrMisTraks
	}
	t.buf.Reset()

	defer func() {
		if err != nil {
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
				ts.timeStamp = int(F.Btoi32(buf, tfdt+16))
			case 1:
				ts.data = buf[tfdt+12 : tfdt+20]
				ts.timeStamp = int(F.Btoi64(buf, tfdt+12))
			}
			return
		}

		//get track type
		get_track_type = func(tfhd, tfdt int) (ts timeStamp, handlerType byte) {
			track, ok := t.traks[int(F.Btoi(buf, tfhd+12, 4))]
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
					if F.Btoi(buf, i+4, 4) < 1000 {
						return errors.New("find sample size less then 1000")
					}
				}
			}
			return nil
		}

		//is t error?
		check_set_maxT = func(ts timeStamp, equal func(ts timeStamp) error, larger func(ts timeStamp) error) (err error) {
			switch ts.handlerType {
			case 'v':
				if maxVT == 0 {
					maxVT = ts.getT()
				} else if maxVT == ts.getT() && equal != nil {
					err = equal(ts)
				} else if maxVT > ts.getT() && larger != nil {
					err = larger(ts)
				} else {
					maxVT = ts.getT()
				}
			case 'a':
				if maxAT == 0 {
					maxAT = ts.getT()
				} else if maxAT == ts.getT() && equal != nil {
					err = equal(ts)
				} else if maxAT > ts.getT() && larger != nil {
					err = larger(ts)
				} else {
					maxAT = ts.getT()
				}
			default:
			}
			return
		}
	)

	ies, e := decode(buf, "moof")
	if e != nil {
		return 0, e
	}

	var ErrNormal = perrors.New("ErrNormal", "ErrNormal")

	err = deals(ies,
		[][]string{
			{"moof", "mfhd", "traf", "tfhd", "tfdt", "trun", "mdat"},
			{"moof", "mfhd", "traf", "tfhd", "tfdt", "trun", "traf", "tfhd", "tfdt", "trun", "mdat"}},
		[]func(m []ie) (bool, error){
			func(m []ie) (bool, error) {
				var (
					keyframeMoof = buf[m[5].i+20] == byte(0x02)
					// moofSN       = int(F.Btoi(buf, m[1].i+12, 4))
					video timeStamp
				)

				{
					ts, handlerType := get_track_type(m[3].i, m[4].i)
					if ts.handlerType == 'v' {
						if e := checkSampleEntries(m[5].i, m[6].i); e != nil {
							//skip
							t.buf.Reset()
							haveKeyframe = false
							cu = m[0].i
							return false, e
						}
					}
					if handlerType == 'v' {
						video = ts
					}
					if e := check_set_maxT(ts, func(_ timeStamp) error {
						return errors.New("skip")
					}, func(_ timeStamp) error {
						t.buf.Reset()
						haveKeyframe = false
						cu = m[0].i
						return errors.New("skip")
					}); e != nil {
						return false, e
					}
				}

				// fmt.Println(ts.getT(), "frame0", keyframeMoof, t.buf.size(), m[0].i, m[6].n, m[6].e)

				//deal frame
				if keyframeMoof {
					if v, e := t.buf.HadModified(bufModified); e == nil && v && !t.buf.IsEmpty() {
						cu = m[0].i
						if haveKeyframe && len(w) > 0 {
							if ifWrite(video.getT()) {
								_, err = w[0].Write(t.buf.GetPureBuf())
							}
							t.buf.Reset()
							return true, ErrNormal
						}
						t.buf.Reset()
					}
					haveKeyframe = true
				} else if !haveKeyframe {
					cu = m[6].e
				}
				if haveKeyframe {
					if e := t.buf.Append(buf[m[0].i:m[6].e]); e != nil {
						return false, e
					}
				}
				return false, nil
			},
			func(m []ie) (bool, error) {
				var (
					keyframeMoof = buf[m[5].i+20] == byte(0x02) || buf[m[9].i+20] == byte(0x02)
					// moofSN       = int(F.Btoi(buf, m[1].i+12, 4))
					video timeStamp
					audio timeStamp
				)

				// fmt.Println(moofSN, "frame1", keyframeMoof, t.buf.size(), m[0].i, m[10].n, m[10].e)

				{
					ts, handlerType := get_track_type(m[3].i, m[4].i)
					if handlerType == 'v' {
						if e := checkSampleEntries(m[5].i, m[6].i); e != nil {
							//skip
							t.buf.Reset()
							haveKeyframe = false
							cu = m[0].i
							return false, e
						}
					}
					switch handlerType {
					case 'v':
						video = ts
					case 's':
						audio = ts
					}
					if e := check_set_maxT(ts, func(_ timeStamp) error {
						return errors.New("skip")
					}, func(_ timeStamp) error {
						t.buf.Reset()
						haveKeyframe = false
						cu = m[0].i
						return errors.New("skip")
					}); e != nil {
						return false, e
					}
				}
				{
					ts, handlerType := get_track_type(m[7].i, m[8].i)
					if handlerType == 'v' {
						if e := checkSampleEntries(m[9].i, m[10].i); e != nil {
							//skip
							t.buf.Reset()
							haveKeyframe = false
							cu = m[0].i
							return false, e
						}
					}
					switch handlerType {
					case 'v':
						video = ts
					case 's':
						audio = ts
					}
					if e := check_set_maxT(ts, func(_ timeStamp) error {
						return errors.New("skip")
					}, func(_ timeStamp) error {
						t.buf.Reset()
						haveKeyframe = false
						cu = m[0].i
						return errors.New("skip")
					}); e != nil {
						return false, e
					}
				}

				//sync audio timeStamp
				if t.AVTDiff <= 0.1 {
					t.AVTDiff = 0.1
				}
				if diff := math.Abs(video.getT() - audio.getT()); diff > t.AVTDiff {
					return false, fmt.Errorf("时间戳不匹配 %v %v (或许应调整fmp4音视频时间戳容差s>%.2f)", video.timeStamp, audio.timeStamp, diff)
					// copy(video.data, F.Itob64(int64(audio.getT()*float64(video.timescale))))
				}

				//deal frame
				if keyframeMoof {
					if v, e := t.buf.HadModified(bufModified); e == nil && v && !t.buf.IsEmpty() {
						cu = m[0].i
						if haveKeyframe && len(w) > 0 {
							if ifWrite(video.getT()) {
								_, err = w[0].Write(t.buf.GetPureBuf())
							}
							return true, ErrNormal
						}
						t.buf.Reset()
					}
					haveKeyframe = true
				} else if !haveKeyframe {
					cu = m[10].e
				}
				if haveKeyframe {
					if e := t.buf.Append(buf[m[0].i:m[10].e]); e != nil {
						return false, e
					}
				}
				return false, nil
			},
		},
	)

	if errors.Is(err, ErrNormal) {
		err = nil
	}

	return
}

func (t *Fmp4Decoder) Cut(reader io.Reader, startT, duration time.Duration, w io.Writer) (err error) {
	bufSize := humanize.KByte * 1100
	buf := make([]byte, humanize.MByte)
	buff := slice.New[byte]()
	init := false
	over := false
	startTM := startT.Seconds()
	durationM := duration.Seconds()
	firstFT := -1.0

	ifWriteF := func(t float64) bool {
		if firstFT == -1 {
			firstFT = t
		}
		cu := t - firstFT
		over = cu > durationM+startTM
		if startTM <= cu && !over {
			return true
		}
		return false
	}

	for c := 0; err == nil && !over; c++ {
		n, e := reader.Read(buf)
		if n == 0 && errors.Is(e, io.EOF) {
			return io.EOF
		}
		err = buff.Append(buf[:n])
		if buff.Size() < bufSize {
			continue
		}

		if !init {
			if frontBuf, e := t.Init_fmp4(buff.GetPureBuf()); e != nil {
				return perrors.New("Init_fmp4", e.Error())
			} else {
				if len(frontBuf) == 0 {
					bufSize *= 2
					continue
				} else {
					if t.Debug {
						fmt.Printf("write frontBuf: frontBufSize: %d", len(frontBuf))
					}
					init = true
					_, err = w.Write(frontBuf)
				}
			}
		} else {
			if dropOffset, e := t.oneF(buff.GetPureBuf(), ifWriteF, w); e != nil {
				return perrors.New("w", e.Error())
			} else {
				if dropOffset != 0 {
					_ = buff.RemoveFront(dropOffset)
				} else {
					bufSize *= 2
				}
			}
		}
	}
	return
}

func deal(ies []ie, boxNames []string, fs func([]ie) (breakloop bool, err error)) (err error) {
	return deals(ies, [][]string{boxNames}, []func([]ie) (breakloop bool, err error){fs})
}

func deals(ies []ie, boxNames [][]string, fs []func([]ie) (breakloop bool, e error)) (err error) {
	if len(boxNames) != len(fs) {
		panic("boxNames与fs数量不相等")
	}
	var matchCounts = make([]int, len(boxNames))
	for cu := 0; cu < len(ies) && len(boxNames) != 0; cu++ {
		for i := 0; i < len(boxNames); i++ {
			if ies[cu].n == boxNames[i][matchCounts[i]] {
				matchCounts[i] += 1
				if matchCounts[i] == len(boxNames[i]) {
					matchCounts[i] = 0
					if breakloop, e := fs[i](ies[cu-len(boxNames[i])+1 : cu+1]); e != nil {
						return e
					} else if breakloop {
						boxNames = append(boxNames[:i], boxNames[i+1:]...)
						fs = append(fs[:i], fs[i+1:]...)
						matchCounts = append(matchCounts[:i], matchCounts[i+1:]...)
						i -= 1
					}
				}
			} else {
				matchCounts[i] = 0
			}
		}
	}
	return
}

var (
	ErrMisBox     = perrors.New("decode", "ErrMisBox")
	ErrCantResync = perrors.New("decode")
)

func decode(buf []byte, reSyncboxName string) (m []ie, err error) {
	var cu int

	for cu < len(buf)-3 {
		boxName, i, e, E := searchBox(buf, &cu)
		if E != nil {
			if errors.Is(E, io.EOF) {
				if len(m) == 0 {
					err = ErrMisBox
				}
				return
			}
			err = E
			if reSyncI := bytes.Index(buf[cu:], []byte(reSyncboxName)); reSyncI != -1 {
				cu += reSyncI - 4
				m = m[:0]
				continue
			}
			err = ErrCantResync.WithReason(E.Error() + "> 未能reSync")
			return
		}

		m = append(m, ie{
			n: boxName,
			i: i,
			e: e,
		})
	}

	return
}

var (
	ErrUnkownBox = perrors.New("ErrUnkownBox")
)

func searchBox(buf []byte, cu *int) (boxName string, i int, e int, err error) {
	i = *cu
	e = i + int(F.Btoi(buf, *cu, 4))
	boxName = string(buf[*cu+4 : *cu+8])
	isPureBoxOrNeedSkip, ok := boxs[boxName]
	if !ok {
		err = ErrUnkownBox.WithReason("未知包: " + boxName)
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
// 			moofSN := int(F.Btoi(buf, m[1].i+12, 4))
// 			keyframeMoof := buf[m[5].i+20] == byte(0x02) || buf[m[9].i+20] == byte(0x02)
// 			fmt.Println(moofSN, "frame", keyframeMoof, m[0].i, m[10].n, m[10].e)
// 			return false
// 		})
// 	fmt.Println("err", err)
// 	fmt.Println("testBox", "<===")
// }
