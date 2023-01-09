package reply

import (
	"bytes"
	"errors"
	"io"

	F "github.com/qydysky/bili_danmu/F"
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
	timescale   int
	trackID     int
	handlerType byte
}

type timeStamp struct {
	timeStamp   int
	data        []byte
	timescale   int
	handlerType byte
}

func (t *timeStamp) getT() float64 {
	return float64(t.timeStamp) / float64(t.timescale)
}

type Fmp4Decoder struct {
	traks map[int]trak
	buf   bufB
}

func (t *Fmp4Decoder) Init_fmp4(buf []byte) (b []byte, err error) {
	var ftypI, ftypE, moovI, moovE int

	err = deal(buf,
		[]string{"ftyp", "moov"},
		func(m []*ie) bool {
			ftypI = m[0].i
			ftypE = m[0].e
			moovI = m[1].i
			moovE = m[1].e
			return true
		})

	if err != nil {
		return nil, err
	}

	err = deal(buf,
		[]string{"tkhd", "mdia", "mdhd", "hdlr"},
		func(m []*ie) bool {
			tackId := int(F.Btoi(buf, m[0].i+20, 4))
			if t.traks == nil {
				t.traks = make(map[int]trak)
			}
			t.traks[tackId] = trak{
				trackID:     tackId,
				timescale:   int(F.Btoi(buf, m[2].i+20, 4)),
				handlerType: buf[m[3].i+16],
			}
			return false
		})

	if err != nil {
		return nil, err
	}

	if len(t.traks) == 0 {
		return nil, errors.New("未找到任何trak包")
	}

	return append(buf[ftypI:ftypE], buf[moovI:moovE]...), nil
}

func (t *Fmp4Decoder) Seach_stream_fmp4(buf []byte, keyframes *bufB) (cu int, err error) {
	if len(t.traks) == 0 {
		err = errors.New("未初始化traks")
		return
	}

	t.buf.reset()
	var (
		haveKeyframe      bool
		bufModified       = t.buf.getModifiedTime()
		maxSequenceNumber int

		//get timeStamp
		get_timeStamp = func(tfdt int) (ts timeStamp) {
			switch buf[tfdt+8] {
			case 0:
				ts.data = buf[tfdt+16 : tfdt+20]
				ts.timeStamp = int(F.Btoi(buf, tfdt+16, 4))
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
				ts.handlerType = track.handlerType
				ts.timescale = track.timescale
				return ts, track.handlerType
			}
			return
		}
	)

	err = deal(buf,
		[]string{"moof", "mfhd",
			"traf", "tfhd", "tfdt", "trun",
			"traf", "tfhd", "tfdt", "trun",
			"mdat"},
		func(m []*ie) bool {
			var (
				keyframeMoof = buf[m[5].i+20] == byte(0x02) || buf[m[9].i+20] == byte(0x02)
				moofSN       = int(F.Btoi(buf, m[1].i+12, 4))
				video        timeStamp
				audio        timeStamp
			)

			// fmt.Println(moofSN, "frame", keyframeMoof, t.buf.size(), m[0].i, m[10].n, m[10].e)

			//is sn error?
			if maxSequenceNumber == 0 {
				maxSequenceNumber = moofSN
			} else if moofSN == maxSequenceNumber {
				return false
			} else if moofSN != maxSequenceNumber+1 {
				t.buf.reset()
				haveKeyframe = false
				cu = m[0].i
				return false
			} else {
				maxSequenceNumber = moofSN
			}
			{
				ts, handlerType := get_track_type(m[3].i, m[4].i)
				switch handlerType {
				case 'v':
					video = ts
				case 's':
					audio = ts
				}
			}
			{
				ts, handlerType := get_track_type(m[7].i, m[8].i)
				switch handlerType {
				case 'v':
					video = ts
				case 's':
					audio = ts
				}
			}

			//deal frame
			if keyframeMoof {
				//sync audio timeStamp
				if audio.getT() != video.getT() {
					date := F.Itob64(int64(video.getT() * float64(audio.timescale)))
					copy(audio.data, date)
				}
				if t.buf.hadModified(bufModified) && !t.buf.isEmpty() {
					keyframes.append(t.buf.getPureBuf())
					cu = m[0].i
					t.buf.reset()
				}
				haveKeyframe = true
			} else if !haveKeyframe {
				cu = m[10].e
			}
			if haveKeyframe {
				t.buf.append(buf[m[0].i:m[10].e])
			}
			return false
		})

	if len(buf) > 1024*1024*20 {
		err = errors.New("buf超过20M")
	}

	return
}

func deal(buf []byte, boxName []string, f func([]*ie) (breakloop bool)) (err error) {

	m, e := decode(buf, boxName[0])
	if len(m) == 0 {
		return errors.New("未找到box")
	}
	if e != nil {
		err = e
	}

	var matchCount = 0
	for cu := 0; cu < len(m); cu++ {
		if m[cu].n == boxName[matchCount] {
			matchCount += 1
			if matchCount == len(boxName) {
				var ies []*ie

				for k, v := range boxName {
					ies = append(ies, &ie{
						n: v,
						i: m[cu-(matchCount-1)+k].i,
						e: m[cu-(matchCount-1)+k].e,
					})
				}

				if f(ies) {
					break
				}
				matchCount = 0
			}
		} else {
			matchCount = 0
		}
	}

	return
}

func decode(buf []byte, reSyncboxName string) (m []ie, err error) {
	var cu int

	for cu < len(buf) {
		boxName, i, e, E := searchBox(buf, &cu)
		if E != nil {
			err = E
			if reSyncI := bytes.Index(buf[cu:], []byte(reSyncboxName)); reSyncI != -1 {
				cu += reSyncI - 4
				m = []ie{}
				continue
			}
			err = errors.New(E.Error() + " > 未能reSync")
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

func searchBox(buf []byte, cu *int) (boxName string, i int, e int, err error) {
	i = *cu
	e = i + int(F.Btoi(buf, *cu, 4))
	boxName = string(buf[*cu+4 : *cu+8])
	isPureBoxOrNeedSkip, ok := boxs[boxName]
	if !ok {
		err = errors.New("未知包: " + boxName)
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
