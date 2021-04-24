package reply

import (
	"os"
	// "fmt"
	"bytes"
	"time"
	"errors"
	// "math"

	c "github.com/qydysky/bili_danmu/CV"
	F "github.com/qydysky/bili_danmu/F"
)

const (
	flv_header_size = 9
	tag_header_size = 11
	previou_tag_size = 4

	video_tag = byte(0x09)
	audio_tag = byte(0x08)
	script_tag = byte(0x12)

	//custom define
	eof_tag = byte(0x00)
)

var (
	flv_header_sign = []byte{0x46,0x4c,0x56}
	flvlog = c.Log.Base(`flv解码`)
	send_sign = []byte{0x00}
)

func Stream(path string,front_buf *[]byte,streamChan chan []byte,cancel chan struct{}) (error) {
	flvlog.L(`T: `,path)
	defer flvlog.L(`T: `,`退出`)
	//file
	f,err := os.OpenFile(path,os.O_RDONLY,0644)
	if err != nil {
		return err
	}
	defer f.Close()
	defer close(streamChan)

	//get flv header(9byte) + FirstTagSize(4byte)
	{
		buf := make([]byte, flv_header_size+previou_tag_size)
		if _,err := f.Read(buf);err != nil {return err}
		if bytes.Index(buf,flv_header_sign) != 0 {return errors.New(`no flv`)}
		*front_buf = append(*front_buf, buf...)
	}

	type flv_tag struct {
		Tag byte
		Offset int64
		Timestamp int32
		PreSize int32
		FirstByte byte
		Buf *[]byte
	}

	var seachtag = func(f *os.File, begin_offset int64)(available_offset int64){
		available_offset += begin_offset
		f.Seek(begin_offset, 0)//seek to begin

		buf := make([]byte, 1024*1024*10)
		if size,_ := f.Read(buf);size == 0 {
			return
		} else {
			for buf_offset:=0;buf_offset<size; {
				if tag_offset := bytes.IndexAny(buf[buf_offset:], string([]byte{video_tag,audio_tag,script_tag}));tag_offset == -1 {
					return
				} else if streamid_offset := bytes.Index(buf[tag_offset:], []byte{0x00,0x00,0x00});streamid_offset == -1 {
					return
				} else if streamid_offset == 8 {
					available_offset += int64(tag_offset + buf_offset)
					return
				} else {
					buf_offset += tag_offset + 1
				}
			}
		}
		return
	}

	//get tag func
	var getTag = func(f *os.File)(t flv_tag){
		t.Offset,_ = f.Seek(0,1)
		Buf := []byte{}
		t.Buf = &Buf

		buf := make([]byte, tag_header_size)
		if size,err := f.Read(buf);err != nil || size == 0 {
			t.Tag = eof_tag
			return
		}
		Buf = append(Buf, buf...)
		t.Tag = buf[0]
		t.Timestamp = F.Btoi32([]byte{buf[7],buf[4],buf[5],buf[6]},0)

		size := F.Btoi32(append([]byte{0x00},buf[1:4]...),0)

		data := make([]byte, size)
		if size,err := f.Read(data);err != nil || size == 0 {
			t.Tag = eof_tag
			return
		}
		t.FirstByte = data[0]

		pre_tag := make([]byte, previou_tag_size)
		if size,err := f.Read(pre_tag);err != nil || size == 0 {
			t.Tag = eof_tag
			return
		} 
		t.PreSize = F.Btoi32(pre_tag,0)
		
		Buf = append(Buf, append(data, pre_tag...)...)
		// if t.PreSize == 0{fmt.Println(t.Tag,size,data[size:])}

		return
	}

	//get first video and audio tag
	//find last_keyframe_video_offset
	var (
		last_keyframe_video_offsets []int64
		first_video_tag bool
		first_audio_tag bool
		// last_timestamps []int32
	)
	for {
		t := getTag(f)
		if t.Tag == script_tag {
			*front_buf = append(*front_buf, *t.Buf...)
		} else if t.Tag == video_tag {
			if !first_video_tag {
				first_video_tag = true
				*front_buf = append(*front_buf, *t.Buf...)
			}

			if t.FirstByte & 0xf0 == 0x10 {
				if len(last_keyframe_video_offsets) > 2 {
					// last_timestamps = append(last_timestamps[1:], t.Timestamp)
					last_keyframe_video_offsets = append(last_keyframe_video_offsets[1:], t.Offset)
				} else {
					// last_timestamps = append(last_timestamps, t.Timestamp)
					last_keyframe_video_offsets = append(last_keyframe_video_offsets, t.Offset)
				}
			}
		} else if t.Tag == audio_tag {
			if !first_audio_tag {
				first_audio_tag = true
				*front_buf = append(*front_buf, *t.Buf...)
			}
		} else {//eof_tag 
			break;
		}
	}

	//seed to the second last tag
	if len(last_keyframe_video_offsets) == 0 {flvlog.L(`W: `,`no keyframe`);return errors.New(`no keyframe`)}
	f.Seek(last_keyframe_video_offsets[0],0)


	// var (
	// 	last_video_keyframe_timestramp int32
	// 	video_keyframe_speed int32
	// )
	//copy when key frame
	{
		last_available_offset := last_keyframe_video_offsets[0]
		var buf []byte
		// last_Timestamp := last_timestamps[0]
		for {
			//退出
			select {
			case <-cancel:return nil;
			default:;
			}
			t := getTag(f)
			if t.Tag == eof_tag {
				f.Seek(last_available_offset,0)
				time.Sleep(time.Second)
				continue
			} else if t.PreSize == 0 {
				f.Seek(seachtag(f, last_available_offset),0)
				continue
			} else if t.Tag == video_tag {
				if t.FirstByte & 0xf0 == 0x10 {
					streamChan <- buf
					buf = []byte{}
				}
				buf = append(buf, *t.Buf...)
			} else if t.Tag == audio_tag {
				buf = append(buf, *t.Buf...)
			} else if t.Tag != script_tag {
				;
			}
			
			last_available_offset = t.Offset

		}
	}

	return nil
}
