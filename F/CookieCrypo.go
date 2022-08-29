package F

import (
	"fmt"

	c "github.com/qydysky/bili_danmu/CV"
	p "github.com/qydysky/part"
	crypto "github.com/qydysky/part/crypto"
)

//公私钥加密
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
			c.C.Log.Block(1000) //等待所有日志输出完毕
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
			if d, e := crypto.FileLoad(`cookie.txt`); e != nil || string(d[:3]) != `nol` {
				clog.L(`E: `, e, `cookie保存格式:`, string(d[:3]))
				return []byte{}
			} else {
				return d[3:]
			}
		}
	}
	if source, e := crypto.FileLoad(`cookie.txt`); e != nil || string(source[:3]) != `pem` {
		clog.L(`E: `, e, `cookie保存格式:`, string(source[:3]))
		return []byte{}
	} else if s, e := crypto.Decrypt(source[3:], pri); e != nil {
		clog.L(`E: `, e)
		return []byte{}
	} else {
		return s
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
			f := p.File()
			f.FileWR(p.Filel{
				File:    `cookie.txt`,
				Loc:     0,
				Context: []interface{}{`nol`, source},
			})
			return
		}
	}
	if source, e := crypto.Encrypt(source, pub); e != nil {
		clog.L(`E: `, e)
		return
	} else {
		f := p.File()
		f.FileWR(p.Filel{
			File:    `cookie.txt`,
			Loc:     0,
			Context: []interface{}{`pem`, source},
		})
	}
}
