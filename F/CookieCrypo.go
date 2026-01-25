package F

import (
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"sync"

	c "github.com/qydysky/bili_danmu/CV"
	pca "github.com/qydysky/part/crypto/asymmetric"
	pcs "github.com/qydysky/part/crypto/symmetric"
	file "github.com/qydysky/part/file"
	unsafe "github.com/qydysky/part/unsafe"
)

// 公私钥加密
var (
	clog       = c.C.Log.Base(`cookie加密`)
	pub        []byte
	cookie     []byte
	cookieLock sync.RWMutex
	sym        = pcs.Chacha20poly1305F
)

// 需要先判断path存在cookie文件
//
// 从文件获取cookie到缓存
func CookieGet(path string) []byte {
	clog := clog.BaseAdd(`获取`)

	cookieLock.Lock()
	defer cookieLock.Unlock()

	if len(cookie) > 0 {
		clog.T(`从内存中获取cookie`)
		return cookie
	} else {
		clog.T(`从文件中获取cookie`)
	}

	var cookieB []byte
	defer clear(cookieB)

	if d, e := FileLoad(path); e != nil {
		clog.E(`cookie读取错误`, e)
		return []byte{}
	} else {
		switch unsafe.B2S(d[:6]) {
		case `t=nol;`:
			cookie = d[6:]
			return d[6:]
		case `t=pem;`:
			cookieB = d[6:]
		default:
			clog.E(`未知的cookie保存格式:`, unsafe.B2S(d[:6]))
			return []byte{}
		}
	}

	var pri []byte
	defer clear(pri)

	if priS, ok := c.C.K_v.LoadV(`cookie解密私钥`).(string); ok && priS != `` {
		if d, e := FileLoad(priS); e != nil {
			clog.E(`cookie私钥读取错误`, e)
			return []byte{}
		} else {
			pri = d
		}
	} else {
		priS = ``
		fmt.Printf("cookie密钥路径: ")
		_, err := fmt.Scanln(&priS)
		if err != nil {
			clog.E("输入错误", err)
			return []byte{}
		}
		if d, e := FileLoad(priS); e != nil {
			clog.E(`cookie私钥读取错误`, e)
			return []byte{}
		} else {
			pri = d
		}
	}

	priKey, _ := pem.Decode(pri)
	defer clear(priKey.Bytes)

	if dec, e := pca.ChoseAsymmetricByPem(priKey).Decrypt(priKey); e != nil {
		clog.E(`cookie私钥错误`, e)
		return []byte{}
	} else {
		b, ext := pca.Unpack(cookieB)
		defer clear(b)
		defer clear(ext)

		if s, e := dec(sym, b, ext); e != nil {
			clog.E(`cookie私钥解密错误`, e)
			return []byte{}
		} else {
			cookie = s
			return s
		}
	}
}

// 保存到cookie到缓存及文件
func CookieSet(path string, source []byte) {
	clog := clog.BaseAdd(`设置`)

	cookieLock.Lock()
	defer cookieLock.Unlock()

	cookie = append(cookie[:0], source...)
	clog.T(`保存cookie到文件`)

	if len(pub) == 0 {
		if pubS, ok := c.C.K_v.LoadV(`cookie加密公钥`).(string); ok && pubS != `` {
			if d, e := FileLoad(pubS); e != nil {
				clog.E(e)
				return
			} else {
				pub = d
			}
		} else {
			f := file.Open(path)
			_ = f.Delete()
			_, _ = f.Write(unsafe.S2B("t=nol;"))
			_, _ = f.Write(source)
			_ = f.Close()
			return
		}
	}
	pubKey, _ := pem.Decode(pub)
	if enc, e := pca.ChoseAsymmetricByPem(pubKey).Encrypt(pubKey); e != nil {
		clog.E(e)
		return
	} else {
		if b, ext, e := enc(sym, source); e != nil {
			clog.E(e)
			return
		} else {
			f := file.Open(path)
			_ = f.Delete()
			_, _ = f.Write(unsafe.S2B("t=pem;"))
			_, _ = f.Write(pca.Pack(b, ext))
			_ = f.Close()
		}
	}
}

func FileLoad(path string) (data []byte, err error) {
	fileObject, e := os.OpenFile(path, os.O_RDONLY, 0644)
	if e != nil {
		err = e
		return
	}
	defer func() { _ = fileObject.Close() }()
	data, e = io.ReadAll(fileObject)
	if e != nil {
		err = e
		return
	}
	return
}
