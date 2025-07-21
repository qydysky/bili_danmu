package F

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	J "github.com/qydysky/bili_danmu/Json"

	_ "github.com/qydysky/biliApi"
	cmp "github.com/qydysky/part/component2"
	limit "github.com/qydysky/part/limit"
	reqf "github.com/qydysky/part/reqf"
	psync "github.com/qydysky/part/sync"
)

const id = "github.com/qydysky/bili_danmu/F.biliApi"

var apilog = c.C.Log.Base(`api`)
var api_limit = limit.New(2, "1s", "30s") //频率限制2次/s，最大等待时间30s

var biliApi = cmp.Get(id, cmp.PreFuncCu[BiliApiInter]{
	Initf: func(ba BiliApiInter) BiliApiInter {
		ba.SetLocation(c.C.SerLocation)
		ba.SetProxy(c.C.Proxy)
		ba.SetReqPool(c.C.ReqPool)

		savepath := "./cookie.txt"
		if tmp, ok := c.C.K_v.LoadV("cookie路径").(string); ok && tmp != "" {
			savepath = tmp
		}
		ba.SetCookiesCallback(func(cookies []*http.Cookie) {
			CookieSet(savepath, []byte(reqf.Cookies_List_2_String(cookies))) //cookie 存入文件
			psync.StoreAll(c.C.Cookie, reqf.Cookies_List_2_Map(cookies))     //cookie 存入全局变量
		})
		return ba
	},
})

// 获取当前佩戴的牌子
func Get_weared_medal(uid, upUid int) (item J.GetWearedMedal_Data, e error) {
	apilog := apilog.Base_add(`获取佩戴牌子`)
	if err, res := biliApi.GetWearedMedal(uid, upUid); err != nil {
		apilog.L(`E: `, err)
		e = err
	} else {
		item.Roominfo.RoomID = res.RoomID
		item.TargetID = res.TargetID
		item.TodayIntimacy = res.TodayIntimacy
	}
	return
}

// 礼物列表
func Gift_list() (list []struct {
	Bag_id    int
	Gift_id   int
	Gift_name string
	Gift_num  int
	Expire_at int
}) {
	apilog := apilog.Base_add(`礼物列表`)
	if c.C.Roomid == 0 {
		apilog.L(`E: `, `失败！无Roomid`)
		return
	}
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	if !biliApi.IsLogin() {
		apilog.L(`W: `, `未登录`)
		return
	}
	if err, res := biliApi.GetBagList(c.C.Roomid); err != nil {
		apilog.L(`E: `, err)
		return
	} else {
		apilog.L(`T: `, `成功`)
		return res
	}
}

var ErrNoCookiesSave = errors.New("ErrNoCookiesSave")

// 正在直播主播
type UpItem struct {
	Uname      string `json:"uname"`
	Title      string `json:"title"`
	Roomid     int    `json:"roomid"`
	LiveStatus int    `json:"live_status"`
}

// 获取历史观看 直播
func GetHisStream() (Uplist []struct {
	Uname      string
	Title      string
	Roomid     int
	LiveStatus int
}) {
	apilog := apilog.Base_add(`历史直播主播`).L(`T: `, `获取中`)
	defer apilog.L(`T: `, `完成`)
	//验证登录
	if !biliApi.IsLogin() {
		apilog.L(`T: `, `未登录`)
		return
	}
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	if e, res := biliApi.GetHisStream(); e != nil {
		apilog.L(`E: `, e)
		return
	} else {
		Uplist = res
	}
	return
}

// 进入房间
func RoomEntryAction(roomId int) {
	apilog := apilog.Base_add(`进入房间`)
	//验证登录
	if !biliApi.IsLogin() {
		apilog.L(`T: `, `未登录`)
		return
	}
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	if e := biliApi.RoomEntryAction(roomId); e != nil {
		apilog.L(`E: `, e)
		return
	}
}

func Feed_list() (Uplist []struct {
	Roomid     int
	Uname      string
	Title      string
	LiveStatus int
}) {
	apilog := apilog.Base_add(`正在直播主播`).L(`T: `, `获取中`)
	defer apilog.L(`T: `, `完成`)
	//验证登录
	if !biliApi.IsLogin() {
		apilog.L(`T: `, `未登录`)
		return
	}
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	if e, res := biliApi.GetFollowing(); e != nil {
		apilog.L(`E: `, e)
		return
	} else {
		Uplist = res
	}

	return
}

func SearchUP(s string) (list []struct {
	Roomid  int
	Uname   string
	Is_live bool
}) {
	apilog := apilog.Base_add(`搜索主播`)
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	fmt.Println(s)
	if e, res := biliApi.SearchUP(s); e != nil {
		fmt.Println(e)
		apilog.L(`E: `, e)
		return
	} else {
		return res
	}
}

func KeepConnect() {
	for !IsConnected() {
		time.Sleep(time.Duration(30) * time.Second)
	}
}

func IsConnected() bool {
	apilog := apilog.Base_add(`IsConnected`)

	v, ok := c.C.K_v.LoadV(`网络中断不退出`).(bool)
	if !ok || !v {
		return true
	}

	if err := biliApi.IsConnected(); err != nil {
		apilog.L(`W: `, `网络中断`, err)
		return false
	}

	apilog.L(`T: `, `已连接`)
	return true
}
