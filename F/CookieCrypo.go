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
)

// 公私钥加密
var (
	clog       = c.C.Log.Base(`cookie加密`)
	pub        []byte
	pri        []byte
	cookie     []byte
	cookieLock sync.RWMutex
	sym        = pcs.Chacha20poly1305F
)

func CookieGet(path string) []byte {
	clog := clog.Base_add(`获取`)

	cookieLock.RLock()
	defer cookieLock.RUnlock()

	if len(cookie) > 0 {
		clog.L(`T: `, `从内存中获取cookie`)
		return cookie
	} else {
		clog.L(`T: `, `从文件中获取cookie`)
	}

	if len(pri) == 0 {
		if priS, ok := c.C.K_v.LoadV(`cookie解密私钥`).(string); ok && priS != `` {
			if d, e := FileLoad(priS); e != nil {
				clog.L(`E: `, e)
				return []byte{}
			} else {
				pri = d
			}
		} else if pubS, ok := c.C.K_v.LoadV(`cookie加密公钥`).(string); ok && pubS != `` {
			priS := ``
			fmt.Printf("cookie密钥路径: ")
			_, err := fmt.Scanln(&priS)
			if err != nil {
				clog.L(`E: `, "输入错误", err)
				return []byte{}
			}
			if d, e := FileLoad(priS); e != nil {
				clog.L(`E: `, e)
				return []byte{}
			} else {
				pri = d
			}
		} else {
			if d, e := FileLoad(path); e != nil {
				clog.L(`E: `, e, `cookie保存格式`)
				return []byte{}
			} else if string(d[:6]) == `t=nol;` {
				return d[6:]
			} else if string(d[:3]) == `nol` {
				return d[3:]
			} else {
				clog.L(`E: `, e, `cookie保存格式:`, string(d[:6]))
				return []byte{}
			}
		}
	}
	if d, e := FileLoad(path); e != nil {
		clog.L(`E: `, e, `cookie保存格式`)
		return []byte{}
	} else if string(d[:6]) == `t=pem;` {
		priKey, _ := pem.Decode(pri)
		if dec, e := pca.ChoseAsymmetricByPem(priKey).Decrypt(priKey); e != nil {
			clog.L(`E: `, e)
			return []byte{}
		} else {
			b, ext := pca.Unpack(d[6:])
			if s, e := dec(sym, b, ext); e != nil {
				clog.L(`E: `, e)
				return []byte{}
			} else {
				return s
			}
		}
	} else {
		clog.L(`E: `, e, `cookie保存格式:`, string(d[:6]))
		return []byte{}
	}
}

func CookieSet(path string, source []byte) {
	clog := clog.Base_add(`设置`)

	cookieLock.Lock()
	defer cookieLock.Unlock()

	cookie = source
	clog.L(`T: `, `保存cookie到文件`)

	if len(pub) == 0 {
		if pubS, ok := c.C.K_v.LoadV(`cookie加密公钥`).(string); ok && pubS != `` {
			if d, e := FileLoad(pubS); e != nil {
				clog.L(`E: `, e)
				return
			} else {
				pub = d
			}
		} else {
			f := file.New(path, 0, true)
			_ = f.Delete()
			_, _ = f.Write(append([]byte("t=nol;"), source...), true)
			return
		}
	}
	pubKey, _ := pem.Decode(pub)
	if enc, e := pca.ChoseAsymmetricByPem(pubKey).Encrypt(pubKey); e != nil {
		clog.L(`E: `, e)
		return
	} else {
		if b, ext, e := enc(sym, source); e != nil {
			clog.L(`E: `, e)
			return
		} else {
			f := file.New(path, 0, true)
			_ = f.Delete()
			_, _ = f.Write(append([]byte("t=pem;"), pca.Pack(b, ext)...), true)
		}
	}
}

func FileLoad(path string) (data []byte, err error) {
	fileObject, e := os.OpenFile(path, os.O_RDONLY, 0644)
	if e != nil {
		err = e
		return
	}
	defer fileObject.Close()
	data, e = io.ReadAll(fileObject)
	if e != nil {
		err = e
		return
	}
	return
}
