package F

/*
整数 字节转换区
64 8字节
32 4字节
16 2字节
*/
func Itob64(v int64) []byte {
	//binary.BigEndian.PutUint64
	b := make([]byte, 8)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b
}

func Itob32(v int32) []byte {
	//binary.BigEndian.PutUint32
	b := make([]byte, 4)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	return b
}

func Itob16(v int16) []byte {
	//binary.BigEndian.PutUint16
	b := make([]byte, 2)
	b[0] = byte(v >> 8)
	b[1] = byte(v)
	return b
}

func Btoi(b []byte, offset int, size int) int64 {
	if size > 8 {
		panic("最大8位")
	}

	bu := make([]byte, 8)
	l := len(b) - offset
	if l > size {
		l = size
	}
	for i := 0; i < size && i < l; i++ {
		bu[i+8-size] = b[offset+i]
	}

	//binary.BigEndian.Uint64
	return int64(uint64(bu[7]) | uint64(bu[6])<<8 | uint64(bu[5])<<16 | uint64(bu[4])<<24 |
		uint64(bu[3])<<32 | uint64(bu[2])<<40 | uint64(bu[1])<<48 | uint64(bu[0])<<56)
}

func Btoui32(b []byte, offset int) uint32 {
	s := 4
	bu := make([]byte, s)
	l := len(b) - offset
	if l > s {
		l = s
	}
	for i := 0; i < s && i < l; i++ {
		bu[i+s-l] = b[offset+i]
	}

	//binary.BigEndian.Uint32
	return uint32(bu[3]) | uint32(bu[2])<<8 | uint32(bu[1])<<16 | uint32(bu[0])<<24
}

func Btoi32(b []byte, offset int) int32 {
	s := 4
	bu := make([]byte, s)
	l := len(b) - offset
	if l > s {
		l = s
	}
	for i := 0; i < s && i < l; i++ {
		bu[i+s-l] = b[offset+i]
	}

	//binary.BigEndian.Uint32
	return int32((uint32(bu[3]) | uint32(bu[2])<<8 | uint32(bu[1])<<16 | uint32(bu[0])<<24))
}

func Btoui16(b []byte, offset int) uint16 {
	s := 2
	bu := make([]byte, s)
	l := len(b) - offset
	if l > s {
		l = s
	}
	for i := 0; i < s && i < l; i++ {
		bu[i+s-l] = b[offset+i]
	}

	//binary.BigEndian.Uint16
	return uint16(bu[1]) | uint16(bu[0])<<8
}

func Btoi16(b []byte, offset int) int16 {
	s := 2
	bu := make([]byte, s)
	l := len(b) - offset
	if l > s {
		l = s
	}
	for i := 0; i < s && i < l; i++ {
		bu[i+s-l] = b[offset+i]
	}

	//binary.BigEndian.Uint16
	return int16(uint16(bu[1]) | uint16(bu[0])<<8)
}

func Btoi64(b []byte, offset int) int64 {
	s := 8
	bu := make([]byte, s)
	l := len(b) - offset
	if l > s {
		l = s
	}
	for i := 0; i < s && i < l; i++ {
		bu[i+s-l] = b[offset+i]
	}

	//binary.BigEndian.Uint64
	return int64(uint64(bu[7]) | uint64(bu[6])<<8 | uint64(bu[5])<<16 | uint64(bu[4])<<24 |
		uint64(bu[3])<<32 | uint64(bu[2])<<40 | uint64(bu[1])<<48 | uint64(bu[0])<<56)
}
