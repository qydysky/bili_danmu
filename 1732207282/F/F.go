package F

import (
	"bytes"
	"strconv"

	c "github.com/qydysky/bili_danmu/CV"
)

var flog = c.C.Log.Base(`F/F.go`)

// base on source/player-loader-2.0.7.min.js L3313
// base on source/player-loader-2.0.7.min.js L3455
type header struct {
	PackL int32
	HeadL int16
	BodyV int16
	OpeaT int32
	Seque int32
}

// 头部生成与检查
func HeadGen(datalenght, Opeation, Sequence int) []byte {
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.Write(Itob32(int32(datalenght + c.WS_PACKAGE_HEADER_TOTAL_LENGTH)))
	buffer.Write(Itob16(c.WS_PACKAGE_HEADER_TOTAL_LENGTH))
	buffer.Write(Itob16(c.WS_HEADER_DEFAULT_VERSION))
	buffer.Write(Itob32(int32(Opeation)))
	buffer.Write(Itob32(int32(Sequence)))

	return buffer.Bytes()
}

func HeadChe(head []byte) header {

	if len(head) != c.WS_PACKAGE_HEADER_TOTAL_LENGTH {
		flog.Base_add("头部检查").L(`E: `, "输入头长度错误")
		return header{}
	}

	PackL := Btoi32(head, c.WS_PACKAGE_OFFSET)
	HeadL := Btoi16(head, c.WS_HEADER_OFFSET)
	BodyV := Btoi16(head, c.WS_VERSION_OFFSET)
	OpeaT := Btoi32(head, c.WS_OPERATION_OFFSET)
	Seque := Btoi32(head, c.WS_SEQUENCE_OFFSET)

	return header{
		PackL: PackL,
		HeadL: HeadL,
		BodyV: BodyV,
		OpeaT: OpeaT,
		Seque: Seque,
	}
}

// 认证生成与检查
func HelloGen(roomid int, key string) []byte {
	flog := flog.Base_add("认证生成")

	if roomid == 0 || key == "" {
		flog.L(`E: `, "roomid == 0 || key == \"\"")
		return []byte("")
	}

	var obj = `{"uid":` + strconv.Itoa(c.C.Uid) +
		`,"roomid":` + strconv.Itoa(roomid) +
		`,"protover":` + strconv.Itoa(c.Protover)

	if buvid, ok := c.C.Cookie.LoadV("buvid3").(string); ok {
		obj += `,"buvid":"` + buvid + `"`
	}

	obj += `,"platform":"` + c.Platform +
		// `","clientver":"` + c.VERSION + //delete at 2021 4 14
		`","type":` + strconv.Itoa(c.Type) +
		`,"key":"` + key + `"}`

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.Write(HeadGen(len(obj), c.WS_OP_USER_AUTHENTICATION, c.WS_HEADER_DEFAULT_SEQUENCE))

	buffer.Write([]byte(obj))

	return buffer.Bytes()
}

func HelloChe(r []byte) bool {
	if len(r) == 0 {
		return false
	}

	var obj = `{"code":0}`

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.Write(HeadGen(len(obj), c.WS_OP_CONNECT_SUCCESS, c.WS_HEADER_DEFAULT_SEQUENCE))

	buffer.Write([]byte(obj))

	h := buffer.Bytes()

	if len(h) != len(r) {
		return false
	}

	for k, v := range r {
		if v != h[k] {
			return false
		}
	}
	return true
}

// 获取人气生成
func Heartbeat() ([]byte, int) {
	//from player-loader-2.0.4.min.js
	const heartBeatInterval = 30

	var obj = `[object Object]`

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.Write(HeadGen(len(obj), c.WS_OP_HEARTBEAT, c.WS_HEADER_DEFAULT_SEQUENCE))

	buffer.Write([]byte(obj))

	return buffer.Bytes(), heartBeatInterval

}

// cookie检查
func CookieCheck(key []string) (missKey []string) {
	for _, tk := range key {
		if tk == `` {
			continue
		}
		if _, ok := c.C.Cookie.Load(tk); !ok {
			missKey = append(missKey, tk)
		}
	}
	return
}

// just faster, use in right way
//
// eg. ParseQuery(`http://1.com/2?workspace=1`, "workspace=") => `1`
func ParseQuery(rawURL, key string) string {
	s := 0
	for i := 0; i < len(rawURL); i++ {
		if rawURL[i] == '?' {
			s = i + 1
			break
		}
	}

	for i := s; i < len(rawURL); i++ {
		for j := 0; i < len(rawURL) && j < len(key); j, i = j+1, i+1 {
			if rawURL[i] != key[j] {
				break
			} else if j == len(key)-1 {
				s = i + 1
				i = len(rawURL)
				break
			}
		}
	}

	d := s
	for ; d < len(rawURL); d++ {
		if rawURL[d] == '&' || rawURL[d] == '#' {
			break
		}
	}

	return rawURL[s:d]
}

// just faster, use in right way
//
// eg. ParseHost(`http://1.com/2`) => `1.com`
func ParseHost(rawURL string) string {
	s := 0
	for i := 0; i < len(rawURL); i++ {
		for j := 0; i < len(rawURL) && j < len("//"); j, i = j+1, i+1 {
			if rawURL[i] != "//"[j] {
				break
			} else if j == len("//")-1 {
				s = i + 1
				i = len(rawURL)
				break
			}
		}
	}

	d := s
	for i := s; i < len(rawURL); i++ {
		if rawURL[i] == '/' {
			d = i
			break
		}
	}

	return rawURL[s:d]
}

// just faster, use in right way
//
// eg. ResolveReferenceLast(`http://1.com/2`, `1`) => `http://1.com/1`
func ResolveReferenceLast(rawURL, ref string) string {
	s := 0
	for i := 0; i < len(rawURL); i++ {
		if rawURL[i] == '/' {
			s = i
		}
		if rawURL[i] == '?' || rawURL[i] == '#' {
			break
		}
	}

	return rawURL[:s+1] + ref
}
