package F

import (
	"fmt"

	c "github.com/qydysky/bili_danmu/CV"
	crypto "github.com/qydysky/part/crypto"
	file "github.com/qydysky/part/file"
)

// 公私钥加密
var (
	clog = c.C.Log.Base(`cookie加密`)
	pub  []byte
	pri  []byte
)

func CookieGet() []byte {
	clog := clog.Base_add(`获取`)
	if len(pri) == 0 {
		if priS, ok := c.C.K_v.LoadV(`cookie解密私钥`).(string); ok && priS != `` {
			if d, e := crypto.FileLoad(priS); e != nil {
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
			if d, e := crypto.FileLoad(priS); e != nil {
				clog.L(`E: `, e)
				return []byte{}
			} else {
				pri = d
			}
		} else {
			if d, e := crypto.FileLoad(`cookie.txt`); e != nil {
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
	if d, e := crypto.FileLoad(`cookie.txt`); e != nil {
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

func CookieSet(source []byte) {
	clog := clog.Base_add(`设置`)
	if len(pub) == 0 {
		if pubS, ok := c.C.K_v.LoadV(`cookie加密公钥`).(string); ok && pubS != `` {
			if d, e := crypto.FileLoad(pubS); e != nil {
				clog.L(`E: `, e)
				return
			} else {
				pub = d
			}
		} else {
			f := file.New(`cookie.txt`, 0, true)
			_ = f.Delete()
			_, _ = f.Write(append([]byte("t=nol;"), source...), true)
			return
		}
	}
	if source, e := crypto.Encrypt(source, pub); e != nil {
		clog.L(`E: `, e)
		return
	} else {
		f := file.New(`cookie.txt`, 0, true)
		_ = f.Delete()
		_, _ = f.Write(append([]byte("t=pem;"), source...), true)
	}
}
