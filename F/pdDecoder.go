package F

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"iter"
	"math"
	"reflect"

	unsafe "github.com/qydysky/part/unsafe"
)

type Pd struct {
	p      *PdDecoder
	dealed bool
	ty     uint32
}

func Type(ty uint32) uint32 {
	return ty >> 3
}
func WireType(ty uint32) uint32 {
	return ty & 0b00000111
}
func (t *Pd) Type() uint32 {
	return Type(t.ty)
}
func (t *Pd) WireType() uint32 {
	return WireType(t.ty)
}
func (t *Pd) Bool() bool {
	t.dealed = true
	return t.p.uint32() != 0
}
func (t *Pd) Uint64() (r uint64) {
	t.dealed = true
	return t.p.uint64()
}
func (t *Pd) Uint32() (r uint32) {
	t.dealed = true
	return t.p.uint32()
}
func (t *Pd) Float32() (r float32) {
	t.dealed = true
	return t.p.float32()
}
func (t *Pd) Float64() (r float64) {
	t.dealed = true
	return t.p.float64()
}
func (t *Pd) Double() (r float64) {
	return t.Float64()
}
func (t *Pd) Bytes() (r []byte) {
	t.dealed = true
	return t.p.bytes()
}
func (t *Pd) String() (r string) {
	t.dealed = true
	return unsafe.B2S(t.p.bytes())
}
func (t *Pd) Child() *PdDecoder {
	t.dealed = true
	return t.p.child()
}

// func (t *Pd) Slice() iter.Seq[*PdDecoder] {
// 	t.dealed = true
// 	return t.p.slice()
// }

type PdDecoder struct {
	pos     int
	endOfPd int
	buf     []byte
	pdTag   string
	pdM     map[uint32]reflect.Value
}

func NewPdDecoder() *PdDecoder {
	return &PdDecoder{
		pdTag: "pd",
		pdM:   make(map[uint32]reflect.Value),
	}
}
func (t *PdDecoder) LoadBase64S(buf string) *PdDecoder {
	return t.LoadBase64B(unsafe.S2B(buf))
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
	// defer func() {
	// 	if panicM := recover(); panicM != nil {
	// 		panic(panicM)
	// 	}
	// }()
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
	ErrValNoStructPointer    = errors.New(`ErrValNoStructPointer`)
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
		return ErrValNoStructPointer
	} else if rv.IsNil() {
		return ErrValNil
	} else if rt := rv.Type().Elem(); rt.Kind() != reflect.Struct {
		return ErrValNoStructPointer
	} else {
		rv = reflect.Indirect(rv)
		for i := 0; i < rt.NumField(); i++ {
			pdk := unsafe.S2B(rt.Field(i).Tag.Get(t.pdTag))
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
	}

	for pd := range t.Range() {
		if rv, ok := t.pdM[Type(pd.ty)]; ok {
			if e := setFunc(pd, rv); e != nil {
				return e
			}
		}
	}

	return nil
}

func setFunc(pd *Pd, rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.Struct:
		if e := pd.Child().UnmarshalRaw(rv.Addr()); e != nil {
			return e
		}
	case reflect.Slice:
		switch rv.Type().Elem().Kind() {
		case reflect.Uint8:
			rv.SetBytes(pd.Bytes())
		default:
			rvv := reflect.New(rv.Type().Elem()).Elem()
			setFunc(pd, rvv)
			rvv = reflect.Append(rv, rvv)
			rv.Set(rvv)
		}
	case reflect.String:
		rv.SetString(unsafe.B2S(pd.Bytes()))
	case reflect.Int, reflect.Int32, reflect.Int64:
		rv.SetInt(int64(pd.Uint64()))
	case reflect.Uint, reflect.Uint32, reflect.Uint64:
		rv.SetUint(pd.Uint64())
	case reflect.Bool:
		rv.SetBool(pd.Bool())
	case reflect.Float64:
		rv.SetFloat(pd.Float64())
	case reflect.Float32:
		rv.SetFloat(float64(pd.Float32()))
	default:
		return ErrValUnSupportFieldType
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
		pd := &Pd{p: t}
		for t.pos < len(t.buf) {
			pd = t.next(pd)
			if !yield(pd) {
				break
			}
			if !pd.dealed {
				t.skipType(pd.ty)
			}
		}
	}
}

func (t *PdDecoder) skipType(ty uint32) {
	if t.pos >= len(t.buf) {
		return
	}
	switch r := WireType(ty); r {
	case 0:
		for t.pos < len(t.buf) && t.buf[t.pos]&0x80 == 0x80 {
			t.pos += 1
		}
		t.pos += 1
	case 1:
		t.pos += 8
	case 2:
		t.pos = t.endOfPd
	case 5:
		t.pos += 4
	default:
	}
}

func (t *PdDecoder) uint32() (r uint32) {
	for s := 0; s < 32 && t.pos < len(t.buf); {
		if s < 28 {
			r |= uint32(t.buf[t.pos]&0b01111111) << s
			s += 7
		} else {
			r |= uint32(t.buf[t.pos]&0b1111) << s
			s += 4
		}
		if t.buf[t.pos]&0b10000000 == 0x00 {
			s = 32
		}
		t.pos += 1
	}
	return
}

func (t *PdDecoder) uint64() (r uint64) {
	for s := 0; s < 64 && t.pos < len(t.buf); {
		if s < 60 {
			r |= uint64(t.buf[t.pos]&0b01111111) << s
			s += 7
		} else {
			r |= uint64(t.buf[t.pos]&0b1111) << s
			s += 4
		}
		if t.buf[t.pos]&0b10000000 == 0x00 {
			s = 64
		}
		t.pos += 1
	}
	return
}

func (t *PdDecoder) float32() (r float32) {
	r = math.Float32frombits(binary.LittleEndian.Uint32(t.buf[t.pos : t.pos+4]))
	t.pos += 4
	return
}

func (t *PdDecoder) float64() (r float64) {
	r = math.Float64frombits(binary.LittleEndian.Uint64(t.buf[t.pos : t.pos+8]))
	t.pos += 8
	return
}

func (t *PdDecoder) bytes() (r []byte) {
	if t.pos >= len(t.buf) {
		return
	}
	r = t.buf[t.pos:t.endOfPd]
	t.pos = t.endOfPd
	return
}

func (t *PdDecoder) child() *PdDecoder {
	if t.pos >= len(t.buf) {
		return nil
	}
	p := NewPdDecoder().LoadBuf(t.buf[t.pos:t.endOfPd])
	t.pos = t.endOfPd
	return p
}

func (t *PdDecoder) next(prePd *Pd) (pd *Pd) {
	pd = prePd
	pd.dealed = false
	if t.pos >= t.endOfPd {
		pd.ty = t.uint32()
		if WireType(pd.ty) == 2 {
			t.endOfPd = int(t.uint32()) + t.pos
		}
	}
	return
}
