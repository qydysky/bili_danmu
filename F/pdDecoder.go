package F

import (
	"encoding/base64"
	"iter"
)

type Pd struct {
	p      *PdDecoder
	dealed bool
	ty     uint32
}

func (t *Pd) set(ty uint32) *Pd {
	t.ty = ty
	t.dealed = false
	return t
}
func (t *Pd) skipIfNotDeal() {
	if !t.dealed {
		t.p.skipType(t.ty)
	}
}

func (t *Pd) Type() uint32 {
	return t.ty >> 3
}
func (t *Pd) Uint32() (r uint32) {
	t.dealed = true
	return t.p.uint32()
}
func (t *Pd) Bytes() (r []byte) {
	t.dealed = true
	return t.p.bytes()
}

type PdDecoder struct {
	pos int
	buf []byte
}

func NewPdDecoder() *PdDecoder {
	return &PdDecoder{}
}
func (t *PdDecoder) LoadBase64(buf string) *PdDecoder {
	t.buf, _ = base64.StdEncoding.DecodeString(buf)
	t.pos = 0
	return t
}
func (t *PdDecoder) LoadBuf(buf []byte) *PdDecoder {
	t.buf = buf
	t.pos = 0
	return t
}

// var uid uint32
//
//	for pd := range NewPdDecoder(b).Range() {
//		switch pd.Type() {
//		case 1:
//			uid = pd.Uint32()
//		}
//	}
//
//	if 689754432 != uid {
//		t.Fatal(uid)
//	}
func (t *PdDecoder) Range() iter.Seq[*Pd] {
	return func(yield func(*Pd) bool) {
		pd := Pd{p: t}
		for t.pos < len(t.buf) && yield(pd.set(t.uint32())) {
			pd.skipIfNotDeal()
		}
	}
}

func (t *PdDecoder) skip(r ...uint32) {
	if len(r) > 0 {
		if t.pos+int(r[0]) < len(t.buf) {
			t.pos += int(r[0])
		}
	} else {
		for t.pos < len(t.buf) && t.buf[t.pos]&0x80 == 0x80 {
			t.pos += 1
		}
		t.pos += 1
	}
}

func (t *PdDecoder) skipType(ty uint32) {
	if t.pos >= len(t.buf) {
		return
	}
	switch r := uint32(7 & ty); r {
	case 0:
		t.skip()
	case 1:
		t.skip(8)
	case 2:
		t.skip(t.uint32())
	case 3:
		for {
			r = 7 & t.uint32()
			if 4 == r {
				break
			}
			t.skipType(r)
		}
	case 5:
		t.skip(4)
	default:
	}
}

func (t *PdDecoder) uint32() (r uint32) {
	for s := 0; s < 32 && t.pos < len(t.buf); {
		if s <= 28 {
			r |= uint32(t.buf[t.pos]&0b01111111) << s
			s += 7
		} else {
			r |= uint32(t.buf[t.pos]&0b1111) << s
			s += 4
		}
		if t.buf[t.pos]&0x80 == 0x00 {
			s = 32
		}
		t.pos += 1
	}
	return
}

func (t *PdDecoder) bytes() (r []byte) {
	if t.pos >= len(t.buf) {
		return
	}
	size := t.uint32()
	r = t.buf[t.pos : t.pos+int(size)]
	t.pos += int(size)
	return
}
