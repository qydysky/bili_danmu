package F

import (
	"net/http"

	_ "github.com/qydysky/biliApi" //removable
	c "github.com/qydysky/bili_danmu/CV"
	cmp "github.com/qydysky/part/component2"
	reqf "github.com/qydysky/part/reqf"
	psync "github.com/qydysky/part/sync"
	"github.com/qydysky/part/unsafe"
)

var biliApi = cmp.GetV3("github.com/qydysky/bili_danmu/F.biliApi", cmp.PreFuncCu[BiliApiInter]{
	Initf: func(ba BiliApiInter) BiliApiInter {
		ba.SetLocation(c.C.SerLocation)
		ba.SetProxy(c.C.Proxy)
		ba.SetReqPool(c.C.ReqPool)

		savepath := "./cookie.txt"
		if tmp, ok := c.C.K_v.LoadV("cookie路径").(string); ok && tmp != "" {
			savepath = tmp
		}
		ba.SetCookiesCallback(func(cookies []*http.Cookie) {
			CookieSet(savepath, unsafe.S2B(reqf.Cookies_List_2_String(cookies))) //cookie 存入文件
			psync.StoreAll(c.C.Cookie, reqf.Cookies_List_2_Map(cookies))         //cookie 存入全局变量
		})
		return ba
	},
})
