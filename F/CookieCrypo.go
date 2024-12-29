package F

import (
	"fmt"
	"io"
	"os"
	"sync"

	c "github.com/qydysky/bili_danmu/CV"
	crypto "github.com/qydysky/part/crypto"
	file "github.com/qydysky/part/file"
)

// 公私钥加密
var (
	clog       = c.C.Log.Base(`cookie加密`)
	pub        []byte
	pri        []byte
	cookieLock sync.RWMutex
)

func CookieGet(path string) []byte {
	clog := clog.Base_add(`获取`)

	cookieLock.RLock()
	defer cookieLock.RUnlock()

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
		if s, e := crypto.Decrypt(d[6:], pri); e != nil {
			clog.L(`E: `, e)
			return []byte{}
		} else {
			return s
		}
	} else if string(d[:3]) == `pem` {
		if s, e := crypto.Decrypt(d[3:], pri); e != nil {
			clog.L(`E: `, e)
			return []byte{}
		} else {
			return s
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
	if source, e := crypto.Encrypt(source, pub); e != nil {
		clog.L(`E: `, e)
		return
	} else {
		f := file.New(path, 0, true)
		_ = f.Delete()
		_, _ = f.Write(append([]byte("t=pem;"), source...), true)
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
