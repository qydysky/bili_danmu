package reply

import (
	"bytes"
	"errors"

	// "math"

	F "github.com/qydysky/bili_danmu/F"
)

const (
	flv_header_size  = 9
	tag_header_size  = 11
	previou_tag_size = 4

	video_tag  = byte(0x09)
	audio_tag  = byte(0x08)
	script_tag = byte(0x12)

	//custom define
	// eof_tag = byte(0x00)
)

var (
	flv_header_sign = []byte{0x46, 0x4c, 0x56}
	// flvlog          = c.C.Log.Base(`flv解码`)
	// send_sign       = []byte{0x00}
)

// type flv_source struct {
// 	buf   []byte
// 	Rval  reqf.Rval
// 	Inuse bool
// }

// func Stream_sync(sources []*flv_source) (buf_p *[]byte,err error) {
// 	var (
// 		sources sync.Map
// 		buf = []byte{}
// 	)
// 	buf_p = &buf

// 	for i:=0;i<len(sources);i+=1 {
// 		item := *sources[i]
// 		if _,ok := sources.Load(item.req.Url);ok{
// 			time.Sleep(time.Second*5)
// 			continue
// 		}

// 		if len(sources)

// 		go func(m *sync.Map,source_p *flv_source){
// 			m.Store(source_p.Rval.Url,interface{})
// 			defer m.Delete(source_p.Rval.Url)

// 			*source_p.Inuse = true
// 			rval := *source_p.Rval
// 			rval.ResponChan = make(chan []byte,1024*1024)

// 			go func(){
// 				for {
// 					*source_p.buf = append(*source_p.buf, <-rval.ResponChan)

// 				}
// 			}()

// 			if e := rr.Reqf(rval); e != nil{l.L(`W: `,e)}

// 			*source_p.Inuse = false
// 		}(&sources,sources[i])
// 	}
// }

// type flv_tag struct {
// 	Tag       byte
// 	Offset    int64
// 	Timestamp int32
// 	PreSize   int32
// 	FirstByte byte
// 	Buf       *[]byte
// }

// func Tag_stream(c chan []byte,co chan []byte) ( bool) {
// 	//check channel
// 	if c == nil || co == nil {return}

// 	var (
// 		buf []byte
// 		seach_stream_tag = func(buf []byte)(front_buf []byte,available_offset int64){
// 			//get flv header(9byte) + FirstTagSize(4byte)
// 			{
// 				if bytes.Index(buf,flv_header_sign) == 0 {
// 					front_buf = buf[:flv_header_size+previou_tag_size]
// 				} else {
// 					defer front_buf = []byte{}
// 				}
// 			}

// 			var sign = 0x00
// 			for buf_offset:=0;buf_offset<len(buf); {
// 				if tag_offset := buf_offset+bytes.IndexAny(buf[buf_offset:], string([]byte{video_tag,audio_tag,script_tag}));tag_offset == buf_offset-1 {
// 					return//no found available video,audio,script tag
// 				} else if streamid_offset := tag_offset+bytes.Index(buf[tag_offset:], []byte{0x00,0x00,0x00});streamid_offset == tag_offset-1 {
// 					return//no found available streamid
// 				} else if streamid_offset-8 != tag_offset {
// 					buf_offset = tag_offset + 1
// 					continue//streamid offset error
// 				} else if time_offset := tag_offset+4;buf[time_offset:time_offset+2] == []byte{0x00,0x00,0x00}{

// 					size := F.Btoi32(append([]byte{0x00},buf[tag_offset+1:tag_offset+3]...),0)+7
// 					if (buf[tag_offset] == video_tag) && (sign & 0x04 == 0x00) {
// 						sign |= 0x04
// 						front_buf = append(front_buf, buf[tag_offset:tag_offset+size])
// 					} else if (buf[tag_offset] == audio_tag) && (sign & 0x02 == 0x00) {
// 						sign |= 0x02
// 						front_buf = append(front_buf, buf[tag_offset:tag_offset+size])
// 					} else if (buf[tag_offset] == script_tag) && (sign & 0x01 == 0x00) {
// 						sign |= 0x01
// 						front_buf = append(front_buf, buf[tag_offset:tag_offset+size])
// 					}

// 					buf_offset = tag_offset + 1
// 					continue//time error

// 				} else {
// 					available_offset = tag_offset
// 					return
// 				}
// 			}
// 			return
// 		}
// 	)

// 	for {
// 		if b:=<- c;len(b) == 0 {
// 			return
// 		} else {
// 			buf = append(buf, b...)
// 		}
// 		if _,offset := seach_stream_tag(buf);offset != -1{
// 			co <- buf[offset:]
// 			buf = []byte{}
// 			break
// 		}
// 	}
// 	for len(co) != 0 {runtime.Gosched()}
// 	co = c
// 	return true
// }

// func Stream(path string, front_buf *[]byte, streamChan chan []byte, cancel chan struct{}) error {
// 	flvlog.L(`T: `, path)
// 	defer flvlog.L(`T: `, `退出`)
// 	//file
// 	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	defer close(streamChan)

// 	//get flv header(9byte) + FirstTagSize(4byte)
// 	{
// 		buf := make([]byte, flv_header_size+previou_tag_size)
// 		if _, err := f.Read(buf); err != nil {
// 			return err
// 		}
// 		if bytes.Index(buf, flv_header_sign) != 0 {
// 			return errors.New(`no flv`)
// 		}
// 		*front_buf = append(*front_buf, buf...)
// 	}

// 	type flv_tag struct {
// 		Tag       byte
// 		Offset    int64
// 		Timestamp int32
// 		PreSize   int32
// 		FirstByte byte
// 		Buf       *[]byte
// 	}

// 	var seachtag = func(f *os.File, begin_offset int64) (available_offset int64) {
// 		available_offset += begin_offset
// 		f.Seek(begin_offset, 0) //seek to begin

// 		buf := make([]byte, 1024*1024*10)
// 		if size, _ := f.Read(buf); size == 0 {
// 			return
// 		} else {
// 			for buf_offset := 0; buf_offset < size; {
// 				if tag_offset := bytes.IndexAny(buf[buf_offset:], string([]byte{video_tag, audio_tag, script_tag})); tag_offset == -1 {
// 					return
// 				} else if streamid_offset := bytes.Index(buf[tag_offset:], []byte{0x00, 0x00, 0x00}); streamid_offset == -1 {
// 					return
// 				} else if streamid_offset == 8 {
// 					available_offset += int64(tag_offset + buf_offset)
// 					return
// 				} else {
// 					buf_offset += tag_offset + 1
// 				}
// 			}
// 		}
// 		return
// 	}

// 	//get tag func
// 	var getTag = func(f *os.File) (t flv_tag) {
// 		t.Offset, _ = f.Seek(0, 1)
// 		Buf := []byte{}
// 		t.Buf = &Buf

// 		buf := make([]byte, tag_header_size)
// 		if size, err := f.Read(buf); err != nil || size == 0 {
// 			t.Tag = eof_tag
// 			return
// 		}
// 		Buf = append(Buf, buf...)
// 		t.Tag = buf[0]
// 		t.Timestamp = F.Btoi32([]byte{buf[7], buf[4], buf[5], buf[6]}, 0)

// 		size := F.Btoi32(append([]byte{0x00}, buf[1:4]...), 0)

// 		data := make([]byte, size)
// 		if size, err := f.Read(data); err != nil || size == 0 {
// 			t.Tag = eof_tag
// 			return
// 		}
// 		t.FirstByte = data[0]

// 		pre_tag := make([]byte, previou_tag_size)
// 		if size, err := f.Read(pre_tag); err != nil || size == 0 {
// 			t.Tag = eof_tag
// 			return
// 		}
// 		t.PreSize = F.Btoi32(pre_tag, 0)

// 		Buf = append(Buf, append(data, pre_tag...)...)
// 		// if t.PreSize == 0{fmt.Println(t.Tag,size,data[size:])}

// 		return
// 	}

// 	//get first video and audio tag
// 	//find last_keyframe_video_offset
// 	var (
// 		last_keyframe_video_offsets []int64
// 		first_video_tag             bool
// 		first_audio_tag             bool
// 		// last_timestamps []int32
// 	)
// 	for {
// 		t := getTag(f)
// 		if t.Tag == script_tag {
// 			*front_buf = append(*front_buf, *t.Buf...)
// 		} else if t.Tag == video_tag {
// 			if !first_video_tag {
// 				first_video_tag = true
// 				*front_buf = append(*front_buf, *t.Buf...)
// 			}

// 			if t.FirstByte&0xf0 == 0x10 {
// 				if len(last_keyframe_video_offsets) > 2 {
// 					// last_timestamps = append(last_timestamps[1:], t.Timestamp)
// 					last_keyframe_video_offsets = append(last_keyframe_video_offsets[1:], t.Offset)
// 				} else {
// 					// last_timestamps = append(last_timestamps, t.Timestamp)
// 					last_keyframe_video_offsets = append(last_keyframe_video_offsets, t.Offset)
// 				}
// 			}
// 		} else if t.Tag == audio_tag {
// 			if !first_audio_tag {
// 				first_audio_tag = true
// 				*front_buf = append(*front_buf, *t.Buf...)
// 			}
// 		} else { //eof_tag
// 			break
// 		}
// 	}

// 	//seed to the second last tag
// 	if len(last_keyframe_video_offsets) == 0 {
// 		flvlog.L(`W: `, `no keyframe`)
// 		return errors.New(`no keyframe`)
// 	}
// 	f.Seek(last_keyframe_video_offsets[0], 0)

// 	// var (
// 	// 	last_video_keyframe_timestramp int32
// 	// 	video_keyframe_speed int32
// 	// )
// 	//copy when key frame
// 	{
// 		last_available_offset := last_keyframe_video_offsets[0]
// 		var buf []byte
// 		// last_Timestamp := last_timestamps[0]
// 		for {
// 			//退出
// 			select {
// 			case <-cancel:
// 				return nil
// 			default:
// 			}
// 			t := getTag(f)
// 			if t.Tag == eof_tag {
// 				f.Seek(last_available_offset, 0)
// 				time.Sleep(time.Second)
// 				continue
// 			} else if t.PreSize == 0 {
// 				f.Seek(seachtag(f, last_available_offset), 0)
// 				continue
// 			} else if t.Tag == video_tag {
// 				if t.FirstByte&0xf0 == 0x10 {
// 					streamChan <- buf
// 					buf = []byte{}
// 				}
// 				buf = append(buf, *t.Buf...)
// 			} else if t.Tag == audio_tag {
// 				buf = append(buf, *t.Buf...)
// 			}

// 			last_available_offset = t.Offset

// 		}
// 	}

// 	return nil
// }

// func TimeStramp_Check(path string) error {
// 	//file
// 	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	//get flv header(9byte) + FirstTagSize(4byte)
// 	{
// 		buf := make([]byte, flv_header_size+previou_tag_size)
// 		if _, err := f.Read(buf); err != nil {
// 			return err
// 		}
// 		if bytes.Index(buf, flv_header_sign) != 0 {
// 			return errors.New(`no flv`)
// 		}
// 	}

// 	type flv_tag struct {
// 		Tag       byte
// 		Offset    int64
// 		Timestamp int32
// 		PreSize   int32
// 		FirstByte byte
// 	}

// 	//get tag func
// 	var getTag = func(f *os.File) (t flv_tag) {
// 		t.Offset, _ = f.Seek(0, 1)

// 		buf := make([]byte, tag_header_size)
// 		if size, err := f.Read(buf); err != nil || size == 0 {
// 			t.Tag = eof_tag
// 			return
// 		}
// 		t.Tag = buf[0]
// 		t.Timestamp = F.Btoi32([]byte{buf[7], buf[4], buf[5], buf[6]}, 0)

// 		size := F.Btoi32(append([]byte{0x00}, buf[1:4]...), 0)

// 		data := make([]byte, size)
// 		if size, err := f.Read(data); err != nil || size == 0 {
// 			t.Tag = eof_tag
// 			return
// 		}
// 		t.FirstByte = data[0]

// 		pre_tag := make([]byte, previou_tag_size)
// 		if size, err := f.Read(pre_tag); err != nil || size == 0 {
// 			t.Tag = eof_tag
// 			return
// 		}
// 		t.PreSize = F.Btoi32(pre_tag, 0)

// 		// if t.PreSize == 0{fmt.Println(t.Tag,size,data[size:])}

// 		return
// 	}

// 	//get first video and audio tag
// 	//find last_keyframe_video_offset
// 	var (
// 		lasttimestramp int32
// 		// last_timestamps []int32
// 	)
// 	for {
// 		t := getTag(f)
// 		if t.Tag == script_tag && t.Timestamp == 0 {
// 			continue
// 		} else if t.Tag == video_tag || t.Tag == audio_tag {
// 			if t.Timestamp < lasttimestramp {
// 				fmt.Printf("error: now %d < pre %d\n", t.Timestamp, lasttimestramp)
// 				lasttimestramp = t.Timestamp
// 				continue
// 			}
// 			fmt.Printf("%d\n", t.Timestamp)
// 			lasttimestramp = t.Timestamp
// 			if lasttimestramp > 10000 {
// 				return nil
// 			}
// 		} else { //eof_tag
// 			break
// 		}
// 	}
// 	fmt.Printf("ok\n")
// 	return nil
// }

// this fuction read []byte and return flv header and all complete keyframe if possible.
// complete keyframe means the video and audio tags between two video key frames tag
func Seach_stream_tag(buf []byte) (front_buf []byte, keyframe [][]byte, last_avilable_offset int, err error) {
	//get flv header(9byte) + FirstTagSize(4byte)
	if header_offset := bytes.Index(buf, flv_header_sign); header_offset != -1 {
		front_buf = buf[header_offset : header_offset+flv_header_size+previou_tag_size]
		last_avilable_offset = header_offset + flv_header_size + previou_tag_size
	}

	var (
		sign         = 0x00
		keyframe_num = -1
		tag_num      = 0
	)

	defer func() {
		if sign != 0x07 {
			// if sign != 0x00 {
			// fmt.Printf("front_buf error:%x\n", sign)
			// }
			front_buf = []byte{}
		}
		if len(keyframe) > 0 {
			keyframe = keyframe[:len(keyframe)-1]
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
			err = errors.New(`reach end when get tag header`)
			// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
			return //buf end
		}

		streamid := int(F.Btoi32([]byte{0x00, buf[tag_offset+8], buf[tag_offset+9], buf[tag_offset+10]}, 0))
		if streamid != 0 {
			buf_offset = tag_offset + 1
			last_avilable_offset = buf_offset
			// fmt.Printf("streamid error %x\n",buf[tag_offset:tag_offset+tag_header_size])
			continue //streamid error
		}

		tag_size := int(F.Btoi32([]byte{0x00, buf[tag_offset+1], buf[tag_offset+2], buf[tag_offset+3]}, 0))
		if tag_offset+tag_header_size+tag_size+previou_tag_size > len(buf) {
			err = errors.New(`reach end when get tag body`)
			// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
			return //buf end
		}
		if tag_size == 0 {
			buf_offset = tag_offset + 1
			last_avilable_offset = buf_offset
			// fmt.Printf("tag_size error %x\n",buf[tag_offset:tag_offset+tag_header_size])
			continue //tag_size error
		}

		tag_size_check := int(F.Btoi32(buf[tag_offset+tag_header_size+tag_size:tag_offset+tag_header_size+tag_size+previou_tag_size], 0))
		if tag_num+tag_size_check == 0 {
			tag_size_check = tag_size + tag_header_size
		}
		if tag_size_check != tag_size+tag_header_size {
			buf_offset = tag_offset + 1
			last_avilable_offset = buf_offset
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
			last_avilable_offset = buf_offset
			continue
		}

		if buf[tag_offset] == video_tag {
			if buf[tag_offset+11]&0xf0 == 0x10 { //key frame
				keyframe_num += 1
				keyframe = append(keyframe, []byte{})
				last_avilable_offset = tag_offset
			}

			if keyframe_num >= 0 {
				keyframe[keyframe_num] = append(keyframe[keyframe_num], buf[tag_offset:tag_offset+tag_size_check+previou_tag_size]...)
			}
		} else if buf[tag_offset] == audio_tag {
			if keyframe_num >= 0 {
				keyframe[keyframe_num] = append(keyframe[keyframe_num], buf[tag_offset:tag_offset+tag_size_check+previou_tag_size]...)
			}
		}

		buf_offset = tag_offset + tag_size_check + previou_tag_size
	}

	return
}

// same as Seach_stream_tag but faster
// func Seach_keyframe_tag(buf []byte) (front_buf []byte, keyframe [][]byte, err error) {

// 	var (
// 		sign = 0x00
// 		// keyframe_num = -1
// 		tag_num    = 0
// 		buf_offset = 0
// 	)

// 	defer func() {
// 		if sign != 0x07 {
// 			front_buf = []byte{}
// 		}
// 	}()

// 	//front_buf
// 	if header_offset := bytes.Index(buf, flv_header_sign); header_offset != -1 {
// 		front_buf = buf[header_offset : header_offset+flv_header_size+previou_tag_size]

// 		for buf_offset+tag_header_size < len(buf) {

// 			fmt.Println(`front_buf`, buf_offset)

// 			tag_offset := buf_offset + bytes.IndexAny(buf[buf_offset:], string([]byte{video_tag, audio_tag, script_tag}))
// 			if tag_offset == buf_offset-1 {
// 				err = errors.New(`no found available tag`)
// 				// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				return //no found available video,audio,script tag
// 			}
// 			if tag_offset+tag_header_size > len(buf) {
// 				err = errors.New(`reach end when get tag header`)
// 				// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				return //buf end
// 			}

// 			if buf[tag_offset+8]|buf[tag_offset+9]|buf[tag_offset+10] != 0 {
// 				buf_offset = tag_offset + 1
// 				// fmt.Printf("streamid error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				continue //streamid error
// 			}

// 			tag_size := int(F.Btoi32([]byte{0x00, buf[tag_offset+1], buf[tag_offset+2], buf[tag_offset+3]}, 0))
// 			if tag_offset+tag_header_size+tag_size+previou_tag_size > len(buf) {
// 				err = errors.New(`reach end when get tag body`)
// 				// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				return //buf end
// 			}
// 			if tag_size == 0 {
// 				buf_offset = tag_offset + 1
// 				// fmt.Printf("tag_size error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				continue //tag_size error
// 			}

// 			tag_size_check := int(F.Btoi32(buf[tag_offset+tag_header_size+tag_size:tag_offset+tag_header_size+tag_size+previou_tag_size], 0))
// 			if tag_num+tag_size_check == 0 {
// 				tag_size_check = tag_size + tag_header_size
// 			}
// 			if tag_size_check != tag_size+tag_header_size {
// 				buf_offset = tag_offset + 1
// 				// fmt.Printf("tag_size_check error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				continue //tag_size_check error
// 			}

// 			tag_num += 1

// 			if buf[tag_offset+7]|buf[tag_offset+4]|buf[tag_offset+5]|buf[tag_offset+6] == 0 {

// 				if len(front_buf) != 0 {
// 					if (buf[tag_offset] == video_tag) && (sign&0x04 == 0x00) {
// 						sign |= 0x04
// 						front_buf = append(front_buf, buf[tag_offset:tag_offset+tag_size_check+previou_tag_size]...)
// 					} else if (buf[tag_offset] == audio_tag) && (sign&0x02 == 0x00) {
// 						sign |= 0x02
// 						front_buf = append(front_buf, buf[tag_offset:tag_offset+tag_size_check+previou_tag_size]...)
// 					} else if (buf[tag_offset] == script_tag) && (sign&0x01 == 0x00) {
// 						sign |= 0x01
// 						front_buf = append(front_buf, buf[tag_offset:tag_offset+tag_size_check+previou_tag_size]...)
// 					}
// 				}
// 				buf_offset = tag_offset + tag_size_check + previou_tag_size
// 			}
// 			if sign == 0x07 {
// 				break
// 			}
// 		}
// 	}

// 	//keyframe
// 	var last_keyframe_offset int
// 	for buf_offset+tag_header_size < len(buf) {
// 		fmt.Println(`keyframe`, buf_offset)
// 		tag_offset := buf_offset + bytes.Index(buf[buf_offset:], []byte{video_tag})
// 		if tag_offset == buf_offset-1 {
// 			err = errors.New(`no found available tag`)
// 			// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 			return //no found available video,audio,script tag
// 		}
// 		if tag_offset+tag_header_size > len(buf) {
// 			err = errors.New(`reach end when get tag header`)
// 			// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 			return //buf end
// 		}

// 		if buf[tag_offset+8]|buf[tag_offset+9]|buf[tag_offset+10] != 0 {
// 			buf_offset = tag_offset + 1
// 			// fmt.Printf("streamid error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 			continue //streamid error
// 		}

// 		tag_size := int(F.Btoi32([]byte{0x00, buf[tag_offset+1], buf[tag_offset+2], buf[tag_offset+3]}, 0))
// 		if tag_offset+tag_header_size+tag_size+previou_tag_size > len(buf) {
// 			err = errors.New(`reach end when get tag body`)
// 			// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 			return //buf end
// 		}
// 		if tag_size == 0 {
// 			buf_offset = tag_offset + 1
// 			// fmt.Printf("tag_size error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 			continue //tag_size error
// 		}

// 		tag_size_check := int(F.Btoi32(buf[tag_offset+tag_header_size+tag_size:tag_offset+tag_header_size+tag_size+previou_tag_size], 0))
// 		if tag_num+tag_size_check == 0 {
// 			tag_size_check = tag_size + tag_header_size
// 		}
// 		if tag_size_check != tag_size+tag_header_size {
// 			buf_offset = tag_offset + 1
// 			// fmt.Printf("tag_size_check error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 			continue //tag_size_check error
// 		}

// 		// fmt.Printf("%x\n",buf[tag_offset:tag_offset+tag_header_size])

// 		tag_num += 1

// 		if buf[tag_offset] == video_tag {
// 			if buf[tag_offset+11]&0xf0 == 0x10 { //key frame
// 				if last_keyframe_offset != 0 {
// 					keyframe = append(keyframe, buf[last_keyframe_offset:tag_offset])
// 				}

// 				last_keyframe_offset = tag_offset
// 			}
// 		}

// 		buf_offset = tag_offset + tag_size_check + previou_tag_size
// 	}

// 	return
// }

// this fuction merge two stream and return the merge buffer,which has the newest frame.
// once len(merge_buf) isn't 0,old_buf can be drop and new_buf can be used from now on.or it's still need to keep buf until find the same tag.
// func Merge_stream(keyframe_lists [][][]byte, last_keyframe_timestramp int) (keyframe_timestamp int, merge_buf []byte, merged int) {

// 	if len(keyframe_lists) == 0 {
// 		return
// 	}

// 	// var keyframe_lists [][][]byte
// 	// for i:=0;i<len(bufs);i+=1 {
// 	// 	_,keyframe_list,_ := Seach_stream_tag(bufs[i])
// 	// 	keyframe_lists = append(keyframe_lists, keyframe_list)
// 	// }

// 	var (
// 		buf   [][]byte
// 		buf_o int
// 	)
// 	buf = keyframe_lists[0]

// 	// fmt.Println(`buf:`,len(buf[0]),buf[0][:tag_header_size])
// 	// fmt.Println(`buf:`,buf[len(buf)-1][:tag_header_size])
// 	// fmt.Println(`buf1:`,len(keyframe_lists[1]),keyframe_lists[1][0][:tag_header_size])
// 	// fmt.Println(`buf1:`,keyframe_lists[1][len(keyframe_lists[1])-1][:tag_header_size])

// 	for i := 1; i < len(keyframe_lists); i += 1 {
// 		for n := buf_o; n < len(buf); n += 1 {
// 			for o := 0; o < len(keyframe_lists[i]); o += 1 {
// 				// fmt.Println(keyframe_lists[o])
// 				// keyframe_list_i_header := fmt.Sprintf("%x",keyframe_lists[i][o][:tag_header_size-3])
// 				// old_buf_o := buf_o
// 				if bytes.Index(buf[n][1:4], keyframe_lists[i][o][1:4]) != -1 {

// 					// last_time_stamp := int(F.Btoi32([]byte{buf[n][7], buf[n][4], buf[n][5], buf[n][6]},0))

// 					// tmp_kfs := make([][]byte,len(keyframe_lists[i][o:]))

// 					keyframe_timestamp, _ = Keyframe_timebase(keyframe_lists[i][o:], last_keyframe_timestramp)

// 					buf = append(buf[:n], keyframe_lists[i][o:]...)
// 					merged = i
// 					break
// 				}

// 				// if keyframe_list_i_header == fmt.Sprintf("%x",buf[n][:tag_header_size-3]) {
// 				// 	// buf_o = n
// 				// 	buf = append(buf[:buf_o], keyframe_lists[i][o:]...)
// 				// 	merged = true
// 				// 	break
// 				// }
// 			}
// 			// if old_buf_o != buf_o {break}
// 		}
// 	}

// 	// merged = len(buf) != len(keyframe_lists[0])

// 	for n := 0; n < len(buf); n += 1 {
// 		merge_buf = append(merge_buf, buf[n]...)
// 	}
// 	return

// 	// for i:=0;i<len(old_list);i+=1 {
// 	// 	old_tag_header := fmt.Sprintf("%x",old_list[i][:tag_header_size])
// 	// 	for n:=0;n<len(new_list);n+=1 {
// 	// 		new_tag_header := fmt.Sprintf("%x",new_list[n][:tag_header_size])
// 	// 		if old_tag_header == new_tag_header {
// 	// 			old_offset := bytes.Index(old_buf, old_list[i][:tag_header_size])
// 	// 			new_offset := bytes.Index(new_buf, new_list[n][:tag_header_size])
// 	// 			merge_buf = append(old_buf[:old_offset], new_buf[new_offset:]...)
// 	// 			return
// 	// 		}
// 	// 	}
// 	// }

// 	// return
// }

// func Keyframe_timebase(buf [][]byte, last_keyframe_timestamp int) (keyframe_timestamp int, err error) {
// 	var (
// 		tag_num            int
// 		base_keyframe_time int
// 		keyframe_interval  int
// 	)

// 	//search keyframe
// 	{
// 		var first_t, last_t int
// 		for buf_offset := 0; buf_offset+tag_header_size < len(buf[0]); {

// 			tag_offset := buf_offset + bytes.IndexAny(buf[0][buf_offset:], string([]byte{video_tag, audio_tag, script_tag}))
// 			if tag_offset == buf_offset-1 {
// 				err = errors.New(`no found available tag`)
// 				// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				return //no found available video,audio,script tag
// 			}
// 			if tag_offset+tag_header_size > len(buf[0]) {
// 				err = errors.New(`reach end when get tag header`)
// 				// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				return //buf end
// 			}

// 			if buf[0][tag_offset+8]|buf[0][tag_offset+9]|buf[0][tag_offset+10] != 0 {
// 				buf_offset = tag_offset + 1
// 				// fmt.Printf("streamid error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				continue //streamid error
// 			}

// 			tag_size := int(F.Btoi32([]byte{0x00, buf[0][tag_offset+1], buf[0][tag_offset+2], buf[0][tag_offset+3]}, 0))
// 			if tag_offset+tag_header_size+tag_size+previou_tag_size > len(buf[0]) {
// 				err = errors.New(`reach end when get tag body`)
// 				// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				return //buf end
// 			}
// 			if tag_size == 0 {
// 				buf_offset = tag_offset + 1
// 				// fmt.Printf("tag_size error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				continue //tag_size error
// 			}

// 			tag_size_check := int(F.Btoi32(buf[0][tag_offset+tag_header_size+tag_size:tag_offset+tag_header_size+tag_size+previou_tag_size], 0))
// 			if tag_num+tag_size_check == 0 {
// 				tag_size_check = tag_size + tag_header_size
// 			}
// 			if tag_size_check != tag_size+tag_header_size {
// 				buf_offset = tag_offset + 1
// 				// fmt.Printf("tag_size_check error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				continue //tag_size_check error
// 			}

// 			tag_num += 1

// 			time_stamp := int(F.Btoi32([]byte{buf[0][tag_offset+7], buf[0][tag_offset+4], buf[0][tag_offset+5], buf[0][tag_offset+6]}, 0))

// 			// if tag_num == 1 && last_keyframe_timestamp != 0 {
// 			// 	diff_time = last_keyframe_timestamp + 3000 - time_stamp
// 			// 	fmt.Printf("时间戳调整 last:%d now:%d diff:%d\n",last_keyframe_timestamp,time_stamp,diff_time)

// 			if buf[0][tag_offset] == video_tag && buf[0][tag_offset+11]&0xf0 == 0x10 {
// 				first_t = time_stamp
// 			} else {
// 				last_t = time_stamp
// 			}
// 			// }

// 			buf_offset = tag_offset + tag_size_check + previou_tag_size
// 		}
// 		for keyframe_interval = 100; keyframe_interval <= last_t-first_t; keyframe_interval += 100 {
// 		}
// 	}

// 	tag_num = 0
// 	base_keyframe_time = 0

// 	for i := 0; i < len(buf); i += 1 {
// 		keyframe_timestamp = last_keyframe_timestamp + keyframe_interval

// 		for buf_offset := 0; buf_offset+tag_header_size < len(buf[i]); {

// 			tag_offset := buf_offset + bytes.IndexAny(buf[i][buf_offset:], string([]byte{video_tag, audio_tag, script_tag}))
// 			if tag_offset == buf_offset-1 {
// 				err = errors.New(`no found available tag`)
// 				// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				return //no found available video,audio,script tag
// 			}
// 			if tag_offset+tag_header_size > len(buf[i]) {
// 				err = errors.New(`reach end when get tag header`)
// 				// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				return //buf end
// 			}

// 			if buf[i][tag_offset+8]|buf[i][tag_offset+9]|buf[i][tag_offset+10] != 0 {
// 				buf_offset = tag_offset + 1
// 				// fmt.Printf("streamid error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				continue //streamid error
// 			}

// 			tag_size := int(F.Btoi32([]byte{0x00, buf[i][tag_offset+1], buf[i][tag_offset+2], buf[i][tag_offset+3]}, 0))
// 			if tag_offset+tag_header_size+tag_size+previou_tag_size > len(buf[i]) {
// 				err = errors.New(`reach end when get tag body`)
// 				// fmt.Printf("last %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				return //buf end
// 			}
// 			if tag_size == 0 {
// 				buf_offset = tag_offset + 1
// 				// fmt.Printf("tag_size error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				continue //tag_size error
// 			}

// 			tag_size_check := int(F.Btoi32(buf[i][tag_offset+tag_header_size+tag_size:tag_offset+tag_header_size+tag_size+previou_tag_size], 0))
// 			if tag_num+tag_size_check == 0 {
// 				tag_size_check = tag_size + tag_header_size
// 			}
// 			if tag_size_check != tag_size+tag_header_size {
// 				buf_offset = tag_offset + 1
// 				// fmt.Printf("tag_size_check error %x\n",buf[tag_offset:tag_offset+tag_header_size])
// 				continue //tag_size_check error
// 			}

// 			tag_num += 1

// 			time_stamp := int(F.Btoi32([]byte{buf[i][tag_offset+7], buf[i][tag_offset+4], buf[i][tag_offset+5], buf[i][tag_offset+6]}, 0))

// 			// if tag_num == 1 && last_keyframe_timestamp != 0 {
// 			// 	diff_time = last_keyframe_timestamp + 3000 - time_stamp
// 			// 	fmt.Printf("时间戳调整 last:%d now:%d diff:%d\n",last_keyframe_timestamp,time_stamp,diff_time)

// 			if buf[i][tag_offset] == video_tag && buf[i][tag_offset+11]&0xf0 == 0x10 {
// 				// if  {//key frame
// 				base_keyframe_time = time_stamp
// 				time_stamp = keyframe_timestamp
// 				last_keyframe_timestamp = keyframe_timestamp
// 				// fmt.Printf("当前关键帧时间戳 %d %d=>%d\n",last_keyframe_timestamp,base_keyframe_time,keyframe_timestamp)
// 				// }
// 			} else {
// 				time_stamp += keyframe_timestamp - base_keyframe_time
// 			}
// 			// }

// 			time_stamp_byte := F.Itob32(int32(time_stamp))

// 			buf[i][tag_offset+7] = time_stamp_byte[0]
// 			buf[i][tag_offset+4] = time_stamp_byte[1]
// 			buf[i][tag_offset+5] = time_stamp_byte[2]
// 			buf[i][tag_offset+6] = time_stamp_byte[3]

// 			buf_offset = tag_offset + tag_size_check + previou_tag_size
// 		}
// 	}
// 	return
// }

// func SearchStreamOffset(buf []byte) (front_buf []byte, available_offset int) {
// 	//get flv header(9byte) + FirstTagSize(4byte)
// 	{
// 		if bytes.Index(buf, flv_header_sign) == 0 {
// 			front_buf = buf[:flv_header_size+previou_tag_size]
// 		}
// 	}

// 	var sign = 0x00
// 	for buf_offset := 0; buf_offset < len(buf); {
// 		if tag_offset := buf_offset + bytes.IndexAny(buf[buf_offset:], string([]byte{video_tag, audio_tag, script_tag})); tag_offset == buf_offset-1 {
// 			return //no found available video,audio,script tag
// 		} else if streamid_offset := tag_offset + bytes.Index(buf[tag_offset:], []byte{0x00, 0x00, 0x00}); streamid_offset == tag_offset-1 {
// 			return //no found available streamid
// 		} else if streamid_offset-8 != tag_offset {
// 			buf_offset = tag_offset + 1
// 			continue //streamid offset error
// 		} else if time_offset := tag_offset + 4; bytes.Index(buf[time_offset:time_offset+2], []byte{0x00, 0x00, 0x00}) == 0 {

// 			size := int(F.Btoi32(append([]byte{0x00}, buf[tag_offset+1:tag_offset+3]...), 0) + 7)
// 			if (buf[tag_offset] == video_tag) && (sign&0x04 == 0x00) {
// 				sign |= 0x04
// 				front_buf = append(front_buf, buf[tag_offset:tag_offset+size]...)
// 			} else if (buf[tag_offset] == audio_tag) && (sign&0x02 == 0x00) {
// 				sign |= 0x02
// 				front_buf = append(front_buf, buf[tag_offset:tag_offset+size]...)
// 			} else if (buf[tag_offset] == script_tag) && (sign&0x01 == 0x00) {
// 				sign |= 0x01
// 				front_buf = append(front_buf, buf[tag_offset:tag_offset+size]...)
// 			}

// 			buf_offset = tag_offset + 1
// 			continue //time error

// 		} else {
// 			available_offset = tag_offset
// 			return
// 		}
// 	}
// 	return
// }
