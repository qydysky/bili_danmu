package parsem3u8

import (
	"bytes"
	"encoding/base64"
	"errors"
	"iter"
	"strconv"
	"strings"

	comp "github.com/qydysky/part/component2"
	ps "github.com/qydysky/part/slice"
	unsafe "github.com/qydysky/part/unsafe"
)

type TargetInterface interface {
	Parse(respon []byte, lastNo int) (m4sLink iter.Seq[interface {
		IsHeader() bool
		M4sLink() string
	}], redirectUrl string, err error)
	IsErrRedirect(e error) bool
}

func init() {
	if e := comp.Register[TargetInterface]("parseM3u8", parseM3u8{}); e != nil {
		panic(e)
	}
}

var (
	ErrRedirect = errors.New(`ErrRedirect`)

	extXStreamInf = []byte("#EXT-X-STREAM-INF")
	extXMap       = []byte("#EXT-X-MAP")
)

type parseM3u8I struct {
	link   string
	header bool
}

func (t parseM3u8I) IsHeader() bool  { return t.header }
func (t parseM3u8I) M4sLink() string { return t.link }

type parseM3u8 struct{}

func (t parseM3u8) IsErrRedirect(e error) bool {
	return e != nil && errors.Is(e, ErrRedirect)
}

func (t parseM3u8) Parse(respon []byte, lastNo int) (m4sLink iter.Seq[interface {
	IsHeader() bool
	M4sLink() string
}], redirectUrl string, err error) {
	// base64解码
	if len(respon) != 0 && !bytes.Contains(respon, []byte{'#'}) {
		respon, err = base64.StdEncoding.DecodeString(unsafe.B2S(respon))
		if err != nil {
			return
		}
	}

	var (
		maxqn            = -1
		hasExtXStreamInf = false
	)

	ps.Split(respon, []byte{'\n'}, func(s []byte) bool {
		if len(s) == 0 {
			return true
		}
		if hasExtXStreamInf {
			tmp := strings.TrimSpace(unsafe.B2S(s))
			if qn, e := strconv.Atoi(ParseQuery(tmp, "qn=")); e == nil {
				if maxqn < qn {
					maxqn = qn
					redirectUrl = tmp
				}
			}
			err = ErrRedirect
			hasExtXStreamInf = false
		}
		if bytes.HasPrefix(s, extXStreamInf) {
			hasExtXStreamInf = true
		}
		return true
	})
	if t.IsErrRedirect(err) {
		return
	}

	m4sLink = func(yield func(interface {
		IsHeader() bool
		M4sLink() string
	}) bool) {
		ps.Split(respon, []byte{'\n'}, func(line []byte) bool {
			if len(line) == 0 {
				return true
			}

			var (
				m4sLink  string //切片文件名
				isHeader bool
			)

			if line[0] == '#' {
				if bytes.HasPrefix(line, extXMap) {
					if lastNo != 0 {
						return true
					}
					e := bytes.Index(line[16:], []byte{'"'}) + 16
					m4sLink = string(line[16:e])
					isHeader = true
				} else {
					return true
				}
			} else {
				m4sLink = string(line)
			}

			if !isHeader {
				// 只增加新的切片
				if no, _ := strconv.Atoi(m4sLink[:len(m4sLink)-4]); lastNo >= no {
					return true
				}
			}

			if !yield(parseM3u8I{
				header: isHeader,
				link:   m4sLink,
			}) {
				return false
			}
			return true
		})
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
