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
	copy_buf_size = 1024*1024
)

var (
	flv_header_sign = []byte{0x46,0x4c,0x56}
	flvlog = c.Log.Base(`flv解码`)
)

func Stream(path string,streamChan chan []byte,cancel chan struct{}) (error) {
	//
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
		f.Seek(0,0)
		buf := make([]byte, flv_header_size+previou_tag_size)
		if _,err := f.Read(buf);err != nil {return err}
		if !bytes.Contains(buf,flv_header_sign) {return errors.New(`no flv`)}
		streamChan <- buf
	}

	//get tag func
	var getTag = func(f *os.File)(tag byte,offset int64,buf_p *[]byte,data_p *[]byte){
		buf := make([]byte, tag_header_size)
		if _,err := f.Read(buf);err != nil {tag = eof_tag;return}
		tag = buf[0]
		size := F.Btoi32(append([]byte{0x00},buf[1:4]...),0)

		data := make([]byte, size+previou_tag_size)
		if _,err := f.Read(data);err != nil {tag = eof_tag;return}
		
		offset,_ = f.Seek(0,1)
		offset -= tag_header_size+int64(size)+previou_tag_size
		return tag,offset,&buf,&data
	}

	//get first video and audio tag
	//find last_keyframe_video_offset
	var last_keyframe_video_offset int64
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
				last_keyframe_video_offset = offset
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

	//seed to last tag
	f.Seek(last_keyframe_video_offset,0)

	//copy
	{
		buf := make([]byte, copy_buf_size)
		eof_wait_turn := 3
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
				}
				if eof_wait_turn < 0 {break}
				eof_wait_turn -= 1
			}

			if size > 0 {
				streamChan <- buf[:size]
			}

			if eof_wait_turn > 0 {time.Sleep(time.Second*3)}
		}
	}

	return nil
}
