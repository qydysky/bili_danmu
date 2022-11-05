package reply

import (
	"bytes"
	"errors"

	F "github.com/qydysky/bili_danmu/F"
)

type trak struct {
	timescale   int
	trackID     int
	handlerType byte
}

type Fmp4Decoder struct {
	traks map[int]trak
}

func (t *Fmp4Decoder) Init_fmp4(buf []byte) error {
	var cu int
	for cu < len(buf) {
		//moov
		moovI := bytes.Index(buf[cu:], []byte("moov"))
		if moovI == -1 {
			break
		}
		moovI = cu + moovI - 4
		cu = moovI
		moovE := moovI + int(F.Btoi(buf, moovI, 4))
		if moovE > len(buf) {
			return errors.New("moov包破损")
		}

		for cu < moovE {
			//trak
			trakI := bytes.Index(buf[cu:], []byte("trak"))
			if trakI == -1 {
				break
			}
			trakI = cu + trakI - 4
			cu = trakI
			trakE := trakI + int(F.Btoi(buf, trakI, 4))
			if trakE > moovE {
				return errors.New("trak包破损")
			}

			//tkhd
			tkhdI := bytes.Index(buf[cu:], []byte("tkhd"))
			if tkhdI == -1 {
				return errors.New("未找到tkhd包")
			}
			tkhdI = cu + tkhdI - 4
			cu = tkhdI
			tkhdE := tkhdI + int(F.Btoi(buf, tkhdI, 4))
			if tkhdE > trakE {
				return errors.New("tkhd包破损")
			}

			//mdia
			mdiaI := bytes.Index(buf[cu:], []byte("mdia"))
			if mdiaI == -1 {
				return errors.New("未找到mdia包")
			}
			mdiaI = cu + mdiaI - 4
			cu = mdiaI
			mdiaE := mdiaI + int(F.Btoi(buf, mdiaI, 4))
			if mdiaE > trakE {
				return errors.New("mdia包破损")
			}

			//mdhd
			mdhdI := bytes.Index(buf[cu:], []byte("mdhd"))
			if mdhdI == -1 {
				return errors.New("未找到mdhd包")
			}
			mdhdI = cu + mdhdI - 4
			cu = mdhdI
			mdhdE := mdhdI + int(F.Btoi(buf, mdhdI, 4))
			if mdhdE > mdiaE {
				return errors.New("mdhd包破损")
			}

			//hdlr
			hdlrI := bytes.Index(buf[cu:], []byte("hdlr"))
			if hdlrI == -1 {
				return errors.New("未找到hdlr包")
			}
			hdlrI = cu + hdlrI - 4
			cu = hdlrI
			hdlrE := hdlrI + int(F.Btoi(buf, hdlrI, 4))
			if hdlrE > mdiaE {
				return errors.New("hdlr包破损")
			}

			tackId := int(F.Btoi(buf, tkhdI+20, 4))
			if t.traks == nil {
				t.traks = make(map[int]trak)
			}
			t.traks[tackId] = trak{
				trackID:     tackId,
				timescale:   int(F.Btoi(buf, mdhdI+20, 4)),
				handlerType: buf[hdlrI+16],
			}
		}
	}
	if len(t.traks) == 0 {
		return errors.New("未找到trak包")
	}
	return nil
}

func (t *Fmp4Decoder) Seach_stream_fmp4(buf []byte) (keyframes [][]byte, last_avilable_offset int, err error) {
	if len(t.traks) == 0 {
		err = errors.New("未初始化traks")
		return
	}

	var (
		cu       int
		keyframe []byte
	)

	for cu < len(buf) {
		//moof
		moofI := bytes.Index(buf[cu:], []byte("moof"))
		if moofI == -1 {
			break
		}
		moofI = cu + moofI - 4
		cu = moofI
		moofE := moofI + int(F.Btoi(buf, moofI, 4))
		if moofE > len(buf) {
			break
		}

		var (
			iskeyFrame     bool
			videoTime      float64
			audioTime      float64
			audioTimeIndex int
			audioTimeSize  int
			audioTimeScale int
		)

		for cu < moofE {
			//traf
			trafI := bytes.Index(buf[cu:], []byte("traf"))
			if trafI == -1 {
				break
			}
			trafI = cu + trafI - 4
			cu = trafI
			trafE := trafI + int(F.Btoi(buf, trafI, 4))
			if trafE > moofE {
				break
			}

			//tfhd
			tfhdI := bytes.Index(buf[cu:], []byte("tfhd"))
			if tfhdI == -1 {
				err = errors.New("未找到tfhd包")
				break
			}
			tfhdI = cu + tfhdI - 4
			cu = tfhdI
			tfhdE := tfhdI + int(F.Btoi(buf, tfhdI, 4))
			if tfhdE > trafE {
				err = errors.New("tfhd包破损")
				break
			}

			//tfdt
			tfdtI := bytes.Index(buf[cu:], []byte("tfdt"))
			if tfdtI == -1 {
				err = errors.New("未找到tfdt包")
				break
			}
			tfdtI = cu + tfdtI - 4
			cu = tfdtI
			tfdtE := tfdtI + int(F.Btoi(buf, tfdtI, 4))
			if tfdtE > trafE {
				err = errors.New("tfdt包破损")
				break
			}

			//trun
			trunI := bytes.Index(buf[cu:], []byte("trun"))
			if trunI == -1 {
				err = errors.New("未找到trun包")
				break
			}
			trunI = cu + trunI - 4
			cu = trunI
			trunE := trunI + int(F.Btoi(buf, trunI, 4))
			if trunE > trafE {
				err = errors.New("trun包破损")
				break
			}

			var (
				timeStamp      int
				timeStampIndex int
				timeSize       int
			)
			switch buf[tfdtI+8] {
			case 0:
				timeSize = 4
				timeStampIndex = tfdtI + 16
				timeStamp = int(F.Btoi(buf, tfdtI+16, 4))
			case 1:
				timeSize = 8
				timeStampIndex = tfdtI + 12
				timeStamp = int(F.Btoi64(buf, tfdtI+12))
			}

			track, ok := t.traks[int(F.Btoi(buf, tfhdI+12, 4))]
			if !ok {
				err = errors.New("找不到trak")
				// log.Default().Println(`cant find trak`, int(F.Btoi(buf, tfhdI+12)))
				continue
			}

			switch track.handlerType {
			case 'v':
				videoTime = float64(timeStamp) / float64(track.timescale)
			case 's':
				audioTimeIndex = timeStampIndex
				audioTimeSize = timeSize
				audioTimeScale = track.timescale
				audioTime = float64(timeStamp) / float64(track.timescale)
			}

			if !iskeyFrame && buf[trunI+20] == byte(0x02) {
				iskeyFrame = true
			}
		}

		if err != nil {
			break
		}

		//change audio timeStamp
		if audioTime != videoTime {
			// err = errors.New("重新设置音频时间戳")
			switch audioTimeSize {
			case 4:
				// log.Default().Println("set audio to:", int32(videoTime*float64(audioTimeScale)))
				date := F.Itob32(int32(videoTime * float64(audioTimeScale)))
				for i := 0; i < audioTimeSize; i += 1 {
					buf[audioTimeIndex+i] = date[i]
				}
			case 8:
				// log.Default().Println("set audio to:", int64(videoTime*float64(audioTimeScale)))
				date := F.Itob64(int64(videoTime * float64(audioTimeScale)))
				for i := 0; i < audioTimeSize; i += 1 {
					buf[audioTimeIndex+i] = date[i]
				}
			}
		}

		if iskeyFrame {
			last_avilable_offset = moofI - 1
			if len(keyframe) != 0 {
				keyframes = append(keyframes, keyframe)
			}
			keyframe = []byte{}
		}

		//mdat
		mdatI := bytes.Index(buf[cu:], []byte("mdat"))
		if moofI == -1 {
			err = errors.New("未找到mdat包")
			break
		}
		mdatI = cu + mdatI - 4
		cu = mdatI
		mdatE := mdatI + int(F.Btoi(buf, mdatI, 4))
		if mdatE > len(buf) {
			// err = errors.New("mdat包破损")
			break
		}

		keyframe = append(keyframe, buf[moofI:mdatE]...)
	}

	if cu == 0 {
		err = errors.New("未找到moof")
	}
	if last_avilable_offset == 0 && len(buf) > 1024*1024*20 {
		err = errors.New("buf超过20M")
	}
	return
}
