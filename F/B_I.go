package F

import (
	"bytes"
	"encoding/binary"

	p "github.com/qydysky/part"
)

/*
整数 字节转换区
64 8字节
32 4字节
16 2字节
*/
func Itob64(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer.Bytes()
}

func Itob32(num int32) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer.Bytes()
}

func Itob16(num int16) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer.Bytes()
}

func btoi64(b []byte) int64 {
	var buffer int64
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &buffer)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer
}

func btoi32(b []byte) int32 {
	var buffer int32
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &buffer)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer
}

func btoi16(b []byte) int16 {
	var buffer int16
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &buffer)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer
}

func Btoi64(b []byte, offset int) int64 {
	for len(b) < 8 {
		b = append([]byte{0x00}, b...)
	}
	return btoi64(b[offset : offset+8])
}

func Btoi32(b []byte, offset int) int32 {
	for len(b) < 4 {
		b = append([]byte{0x00}, b...)
	}
	return btoi32(b[offset : offset+4])
}

func Btoi16(b []byte, offset int) int16 {
	for len(b) < 2 {
		b = append([]byte{0x00}, b...)
	}
	return btoi16(b[offset : offset+2])
}
