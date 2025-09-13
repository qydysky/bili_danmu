package F

import (
	"encoding/base64"
	"errors"
	"iter"
	"math"
	"reflect"
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
func (t *Pd) Bool() bool {
	t.dealed = true
	return t.p.uint32() != 0
}
func (t *Pd) Uint32() (r uint32) {
	t.dealed = true
	return t.p.uint32()
}
func (t *Pd) Bytes() (r []byte) {
	t.dealed = true
	return t.p.bytes()
}
func (t *Pd) Child() *PdDecoder {
	t.dealed = true
	return t.p.child()
}

type PdDecoder struct {
	pos   int
	buf   []byte
	pdTag string
	pdM   map[uint32]reflect.Value
}

func NewPdDecoder() *PdDecoder {
	return &PdDecoder{
		pdTag: "pd",
		pdM:   make(map[uint32]reflect.Value),
	}
}
func (t *PdDecoder) LoadBase64S(buf string) *PdDecoder {
	return t.LoadBase64B([]byte(buf))
}
func (t *PdDecoder) LoadBase64B(buf []byte) *PdDecoder {
	if l := base64.StdEncoding.DecodedLen(len(buf)); cap(t.buf) < l {
		t.buf = append(t.buf[:0], make([]byte, l)...)
	}
	n, _ := base64.StdEncoding.Decode(t.buf, buf)
	return t.LoadBuf(t.buf[:n])
}
func (t *PdDecoder) LoadBuf(buf []byte) *PdDecoder {
	t.buf = buf
	t.pos = 0
	clear(t.pdM)
	return t
}

func UnmarshalBase64S(data string, v any) error {
	if data == "" {
		return nil
	}
	return NewPdDecoder().LoadBase64S(data).UnmarshalRaw(v)
}

func (t *PdDecoder) UnmarshalBase64S(data string, v any) error {
	t.LoadBase64S(data)
	return t.UnmarshalRaw(v)
}

func (t *PdDecoder) UnmarshalBase64B(data []byte, v any) error {
	t.LoadBase64B(data)
	return t.UnmarshalRaw(v)
}

func (t *PdDecoder) Unmarshal(data []byte, v any) error {
	t.LoadBuf(data)
	return t.UnmarshalRaw(v)
}

var (
	ErrValNoPointer          = errors.New(`ErrValNoPointer`)
	ErrValNil                = errors.New(`ErrValNil`)
	ErrValUnSupportFieldType = errors.New(`ErrValUnSupportFieldType`)
)

func (t *PdDecoder) UnmarshalRaw(v any) error {
	var rv reflect.Value
	switch v := v.(type) {
	case reflect.Value:
		rv = v
	default:
		rv = reflect.ValueOf(v)
	}
	if rv.Kind() != reflect.Pointer {
		return ErrValNoPointer
	} else if rv.IsNil() {
		return ErrValNil
	}
	rt := rv.Type().Elem()
	rv = reflect.Indirect(rv)
	for i := 0; i < rt.NumField(); i++ {
		pdk := []byte(rt.Field(i).Tag.Get(t.pdTag))
		pdki := uint32(0)
		for j := 0; j < len(pdk); j++ {
			if pdk[j] < 48 || pdk[j] > 57 {
				break
			}
			pdki *= 10
			pdki += (uint32(pdk[j]) - 48)
		}
		if pdki == 0 {
			continue
		}
		t.pdM[pdki] = rv.Field(i)
	}
	for pd := range t.Range() {
		if rv, ok := t.pdM[pd.Type()]; ok {
			switch rv.Kind() {
			case reflect.Struct:
				if e := pd.Child().UnmarshalRaw(rv.Addr()); e != nil {
					return e
				}
			case reflect.Slice:
				switch rv.Type().String() {
				case "[]uint8":
					rv.SetBytes(pd.Bytes())
				default:
					return ErrValUnSupportFieldType
				}
			case reflect.String:
				rv.SetString(string(pd.Bytes()))
			case reflect.Int, reflect.Int32, reflect.Int64:
				rv.SetInt(int64(pd.Uint32()))
			case reflect.Uint, reflect.Uint32, reflect.Uint64:
				rv.SetUint(uint64(pd.Uint32()))
			case reflect.Bool:
				rv.SetBool(pd.Bool())
			case reflect.Float64:
				rv.SetFloat(math.Float64frombits(uint64(pd.Uint32())))
			default:
				return ErrValUnSupportFieldType
			}
		}
	}

	return nil
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
		if t.pos+int(r[0]) <= len(t.buf) {
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
			if r == 4 {
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

func (t *PdDecoder) child() *PdDecoder {
	if t.pos >= len(t.buf) {
		return nil
	}
	size := t.uint32()
	p := NewPdDecoder().LoadBuf(t.buf[t.pos : t.pos+int(size)])
	t.pos += int(size)
	return p
}
