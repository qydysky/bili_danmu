package F

import (
	"errors"
	"fmt"

	c "github.com/qydysky/bili_danmu/CV"
	J "github.com/qydysky/bili_danmu/Json"

	limit "github.com/qydysky/part/limit"
)

var apilog = c.C.Log.Base(`api`)
var api_limit = limit.New(2, "1s", "30s") //频率限制2次/s，最大等待时间30s

// 获取当前佩戴的牌子
func Get_weared_medal(uid, upUid int) (item J.GetWearedMedal_Data, e error) {
	apilog := apilog.Base_add(`获取佩戴牌子`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		apilog.L(`E: `, `biliApi组件未构建`, ce)
		e = ce
		return nil
	})
	if biliApi == nil {
		return
	}

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

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		apilog.L(`E: `, `biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

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

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		apilog.L(`E: `, `biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

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

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		apilog.L(`E: `, `biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

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

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		apilog.L(`E: `, `biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

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

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		apilog.L(`E: `, `biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if e, res := biliApi.SearchUP(s); e != nil {
		fmt.Println(e)
		apilog.L(`E: `, e)
		return
	} else {
		return res
	}
}

func IsConnected() (ok bool) {
	apilog := apilog.Base_add(`IsConnected`)

	v, ok := c.C.K_v.LoadV(`网络中断不退出`).(bool)
	if !ok || !v {
		return true
	}

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		apilog.L(`E: `, `biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return false
	}

	if err := biliApi.IsConnected(); err != nil {
		apilog.L(`W: `, `网络中断`, err)
		return false
	}

	apilog.L(`T: `, `已连接`)
	return true
}
