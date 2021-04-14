package reply

import (
	"os"
	"bytes"
	"time"
	"errors"

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
	copy_buf_size = 1024*1024*10
)

var (
	flv_header_sign = []byte{0x46,0x4c,0x56}
	flvlog = c.Log.Base(`flv解码`)
)

func Stream(path string,streamChan chan []byte,cancel chan struct{}) (error) {
	flvlog.L(`I: `,path)
	defer flvlog.L(`I: `,`退出`)
	//file
	f,err := os.OpenFile(path,os.O_RDONLY,0644)
	if err != nil {
		return err
	}
	defer f.Close()
	defer close(streamChan)

	//init
	f.Seek(0,0)

	//get flv header(9byte) + FirstTagSize(4byte)
	{
		buf := make([]byte, flv_header_size+previou_tag_size)
		if _,err := f.Read(buf);err != nil {return err}
		if !bytes.Contains(buf,flv_header_sign) {return errors.New(`no flv`)}
		streamChan <- buf
	}

	//get tag func
	var getTag = func(f *os.File)(tag byte,offset int64,buf_p *[]byte,data_p *[]byte){
		buf := make([]byte, tag_header_size)
		buf_p = &buf
		if _,err := f.Read(buf);err != nil {tag = eof_tag;return}
		tag = buf[0]
		size := F.Btoi32(append([]byte{0x00},buf[1:4]...),0)

		data := make([]byte, size+previou_tag_size)
		data_p = &data
		if _,err := f.Read(data);err != nil {tag = eof_tag;return}
		
		offset,_ = f.Seek(0,1)
		offset -= tag_header_size+int64(size)+previou_tag_size
		return
	}

	//get first video and audio tag
	//find last_keyframe_video_offset
	var last_keyframe_video_offsets []int64
	first_video_tag,first_audio_tag := false,false
	for {
		tag,offset,buf_p,data_p := getTag(f)
		if tag == script_tag {
			streamChan <- *buf_p
			streamChan <- *data_p
		} else if tag == video_tag {
			if !first_video_tag {
				first_video_tag = true
				streamChan <- *buf_p
				streamChan <- *data_p
			}

			if (*data_p)[0] & 0xf0 == 0x10 {
				if len(last_keyframe_video_offsets) > 2 {
					last_keyframe_video_offsets = append(last_keyframe_video_offsets[1:], offset)
				} else {
					last_keyframe_video_offsets = append(last_keyframe_video_offsets, offset)
				}
			}
		} else if tag == audio_tag {
			if !first_audio_tag {
				first_audio_tag = true
				streamChan <- *buf_p
				streamChan <- *data_p
			}
		} else {//eof_tag
			break;
		}
	}

	//seed to the second last tag
	f.Seek(last_keyframe_video_offsets[0],0)

	//copy
	{
		buf := make([]byte, copy_buf_size)
		preOffset,_ := f.Seek(0,1)
		for {

			//退出
			select {
			case <-cancel:return nil;
			default:;
			}

			size,err := f.Read(buf)
			if err != nil {
				if err.Error() != `EOF` {
					return err
				} else if offset,_ := f.Seek(0,1);offset == preOffset {
					break
				}
			}

			if size > 0 {
				streamChan <- buf[:size]
			}

			if err != nil {
				preOffset,_ = f.Seek(0,1)
				time.Sleep(time.Duration(1) * time.Second)
			}

		}
	}

	return nil
}
