package reply

import (
	"bytes"
	"errors"

	F "github.com/qydysky/bili_danmu/F"
)

type fmp4KeyFrame struct {
	videoTime int
	audioTime int
	data      []byte
}

func Seach_stream_fmp4(buf []byte) (keyframes []fmp4KeyFrame, last_avilable_offset int, err error) {
	var (
		cu       int
		keyframe fmp4KeyFrame
	)

	for cu < len(buf) {
		//moof
		moofI := cu + bytes.Index(buf[cu:], []byte("moof")) - 4
		if moofI == -1 {
			err = errors.New("未找到moof包")
			break
		}
		moofE := moofI + int(F.Btoi32(buf, moofI))
		if moofE > len(buf) {
			break
		}
		cu = moofI

		var (
			iskeyFrame bool
			videoTime  int
			audioTime  int
		)

		for trafCount := 0; trafCount < 2 && cu < moofE; trafCount += 1 {

			//traf
			trafI := cu + bytes.Index(buf[cu:], []byte("traf")) - 4
			if trafI == -1 {
				err = errors.New("未找到traf包")
				break
			}
			cu = trafI
			trafE := trafI + int(F.Btoi32(buf, trafI))
			if trafE > moofE {
				err = errors.New("traf包破损")
				break
			}

			//tfdt
			tfdtI := cu + bytes.Index(buf[cu:], []byte("tfdt")) - 4
			if tfdtI == -1 {
				err = errors.New("未找到tfdt包")
				break
			}
			cu = tfdtI
			tfdtE := tfdtI + int(F.Btoi32(buf, tfdtI))
			if tfdtE > trafE {
				err = errors.New("tfdt包破损")
				break
			}

			//trun
			trunI := cu + bytes.Index(buf[cu:], []byte("trun")) - 4
			if trunI == -1 {
				err = errors.New("未找到trun包")
				break
			}
			cu = trunI
			trunE := trunI + int(F.Btoi32(buf, trunI))
			if trunE > trafE {
				err = errors.New("trun包破损")
				break
			}

			timeStamp := int(F.Btoi32(buf, tfdtI+16))
			if trafCount == 0 {
				videoTime = timeStamp
			} else if trafCount == 1 {
				audioTime = timeStamp
			}

			if !iskeyFrame && buf[trunI+20] == byte(0x02) {
				iskeyFrame = true
			}
		}

		if err != nil {
			break
		}

		if iskeyFrame {
			last_avilable_offset = moofI - 1
			if len(keyframe.data) != 0 {
				keyframes = append(keyframes, keyframe)
			}
			keyframe = fmp4KeyFrame{
				videoTime: videoTime,
				audioTime: audioTime,
			}
		}

		//mdat
		mdatI := cu + bytes.Index(buf[cu:], []byte("mdat")) - 4
		if moofI == -1 {
			err = errors.New("未找到mdat包")
			break
		}
		cu = mdatI
		mdatE := mdatI + int(F.Btoi32(buf, mdatI))
		if mdatE > len(buf) {
			err = errors.New("mdat包破损")
			break
		}

		keyframe.data = append(keyframe.data, buf[moofI:mdatE]...)
		cu = mdatE
	}

	return
}
