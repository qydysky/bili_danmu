package reply

import (
	"bytes"
	"errors"

	// "math"

	"github.com/dustin/go-humanize"
	F "github.com/qydysky/bili_danmu/F"
	slice "github.com/qydysky/part/slice"
)

const (
	flv_header_size  = 9
	tag_header_size  = 11
	previou_tag_size = 4

	video_tag  = byte(0x09)
	audio_tag  = byte(0x08)
	script_tag = byte(0x12)
)

var (
	flv_header_sign = []byte{0x46, 0x4c, 0x56}
)

// this fuction read []byte and return flv header and all complete keyframe if possible.
// complete keyframe means the video and audio tags between two video key frames tag
func Search_stream_tag(buf []byte, keyframe *slice.Buf[byte]) (front_buf []byte, last_available_offset int, err error) {
	if len(buf) > humanize.MByte*100 {
		err = errors.New("buf too large")
		return
	}
	//get flv header(9byte) + FirstTagSize(4byte)
	if header_offset := bytes.Index(buf, flv_header_sign); header_offset != -1 {
		if header_offset+flv_header_size+previou_tag_size > len(buf) {
			err = errors.New(`no found available tag`)
			return
		}
		front_buf = buf[header_offset : header_offset+flv_header_size+previou_tag_size]
		last_available_offset = header_offset + flv_header_size + previou_tag_size
	}

	var (
		sign         = 0x00
		keyframe_num = -1
		confirm_num  = -1
		tag_num      = 0
	)

	defer func() {
		if sign != 0x07 {
			// if sign != 0x00 {
			// fmt.Printf("front_buf error:%x\n", sign)
			// }
			front_buf = []byte{}
		}
		if bufl := keyframe.Size(); confirm_num != bufl {
			_ = keyframe.RemoveBack(bufl - confirm_num)
		}
	}()

	for buf_offset := 0; buf_offset+tag_header_size < len(buf); {

		tag_offset := buf_offset + bytes.IndexAny(buf[buf_offset:], string([]byte{video_tag, audio_tag, script_tag}))
		if tag_offset == buf_offset-1 {
			err = errors.New(`no found available tag`)
			// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
			return //no found available video,audio,script tag
		}
		if tag_offset+tag_header_size > len(buf) {
			// err = errors.New(`reach end when get tag header`)
			// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
			return //buf end
		}

		streamid := int(F.Btoi32([]byte{0x00, buf[tag_offset+8], buf[tag_offset+9], buf[tag_offset+10]}, 0))
		if streamid != 0 {
			buf_offset = tag_offset + 1
			last_available_offset = buf_offset
			// fmt.Printf("streamid error %x\n",buf[tag_offset:tag_offset+tag_header_size])
			continue //streamid error
		}

		tag_size := int(F.Btoi32([]byte{0x00, buf[tag_offset+1], buf[tag_offset+2], buf[tag_offset+3]}, 0))
		if tag_offset+tag_header_size+tag_size+previou_tag_size > len(buf) {
			// err = errors.New(`reach end when get tag body`)
			// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
			return //buf end
		}
		if tag_size == 0 {
			buf_offset = tag_offset + 1
			last_available_offset = buf_offset
			// fmt.Printf("tag_size error %x\n",buf[tag_offset:tag_offset+tag_header_size])
			continue //tag_size error
		}

		tag_size_check := int(F.Btoi32(buf[tag_offset+tag_header_size+tag_size:tag_offset+tag_header_size+tag_size+previou_tag_size], 0))
		if tag_num+tag_size_check == 0 {
			tag_size_check = tag_size + tag_header_size
		}
		if tag_size_check != tag_size+tag_header_size {
			buf_offset = tag_offset + 1
			last_available_offset = buf_offset
			// fmt.Printf("tag_size_check error %x\n",buf[tag_offset:tag_offset+tag_header_size])
			continue //tag_size_check error
		}

		time_stamp := int(F.Btoi32([]byte{buf[tag_offset+7], buf[tag_offset+4], buf[tag_offset+5], buf[tag_offset+6]}, 0))

		// show tag header
		// fmt.Printf("%x\n", buf[tag_offset:tag_offset+tag_header_size])

		tag_num += 1

		if time_stamp == 0 || sign != 0x00 { // ignore first video audio time_stamp
			if len(front_buf) != 0 {
				//first video audio script tag
				if (buf[tag_offset] == video_tag) && (sign&0x04 == 0x00) {
					sign |= 0x04
					front_buf = append(front_buf, buf[tag_offset:tag_offset+tag_size_check+previou_tag_size]...)
				} else if (buf[tag_offset] == audio_tag) && (sign&0x02 == 0x00) {
					sign |= 0x02
					front_buf = append(front_buf, buf[tag_offset:tag_offset+tag_size_check+previou_tag_size]...)
				} else if (buf[tag_offset] == script_tag) && (sign&0x01 == 0x00) {
					sign |= 0x01
					front_buf = append(front_buf, buf[tag_offset:tag_offset+tag_size_check+previou_tag_size]...)
				}
			}
			buf_offset = tag_offset + tag_size_check + previou_tag_size
			last_available_offset = buf_offset
			continue
		}

		if buf[tag_offset] == video_tag {
			if buf[tag_offset+11]&0xf0 == 0x10 { //key frame
				keyframe_num += 1
				confirm_num = keyframe.Size()
				last_available_offset = tag_offset
			}

			if keyframe_num >= 0 {
				_ = keyframe.Append(buf[tag_offset : tag_offset+tag_size_check+previou_tag_size])
			}
		} else if buf[tag_offset] == audio_tag {
			if keyframe_num >= 0 {
				_ = keyframe.Append(buf[tag_offset : tag_offset+tag_size_check+previou_tag_size])
			}
		}

		buf_offset = tag_offset + tag_size_check + previou_tag_size
	}

	return
}
