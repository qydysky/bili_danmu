package F

import (
	"errors"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mdp/qrterminal/v3"
	c "github.com/qydysky/bili_danmu/CV"
	pe "github.com/qydysky/part/errors"
	file "github.com/qydysky/part/file"
	pkf "github.com/qydysky/part/keyFunc"
	reqf "github.com/qydysky/part/reqf"
	qr "github.com/skip2/go-qrcode"
	"github.com/skratchdot/open-golang/open"
)

var Api = NewGetFuncV2()

type GetFuncV2 struct {
	common *c.Common
	api    *pkf.KeyFunc
	l      sync.Mutex
}

// api管理器
//
//	使用Get(key)获取需要的key，会尝试调用可以获取到该key的接口
func NewGetFuncV2() *GetFuncV2 {
	t := &GetFuncV2{api: pkf.NewKeyFunc()}
	t.api.Reg(`CookieNoBlock`, t.isValid(`Cookie`), t.getCookieNoBlock)
	t.api.Reg(`Cookie`, t.isValid(`Cookie`), t.getCookie)
	t.api.Reg(`UpUid`, t.isValid(`UpUid`), t.getRoomBaseInfo, t.getInfoByRoom, t.getRoomPlayInfo, t.html)
	t.api.Reg(`Live_Start_Time`, t.isValid(`Live_Start_Time`), t.getRoomBaseInfo, t.getInfoByRoom, t.getRoomPlayInfo, t.html)
	t.api.Reg(`Liveing`, t.isValid(`Liveing`), t.getRoomBaseInfo, t.getInfoByRoom, t.getRoomPlayInfo, t.html)
	t.api.Reg(`Title`, t.isValid(`Title`), t.getRoomBaseInfo, t.getInfoByRoom, t.html)
	t.api.Reg(`Uname`, t.isValid(`Uname`), t.getRoomBaseInfo, t.getInfoByRoom, t.html)
	t.api.Reg(`ParentAreaID`, t.isValid(`ParentAreaID`), t.getRoomBaseInfo, t.getInfoByRoom, t.html)
	t.api.Reg(`AreaID`, t.isValid(`AreaID`), t.getRoomBaseInfo, t.getInfoByRoom, t.html)
	t.api.Reg(`Roomid`, t.isValid(`Roomid`), t.getRoomBaseInfo, t.getInfoByRoom)
	t.api.Reg(`GuardNum`, t.isValid(`GuardNum`), t.getGuardNum, t.getInfoByRoom, t.getRoomPlayInfo, t.html)
	t.api.Reg(`Note`, t.isValid(`Note`), t.getPopularAnchorRank, t.getInfoByRoom, t.html)
	t.api.Reg(`Locked`, t.isValid(`Locked`), t.getInfoByRoom, t.html)
	t.api.Reg(`Live_qn`, t.isValid(`Live_qn`), t.getRoomPlayInfo, t.html)
	t.api.Reg(`AcceptQn`, t.isValid(`AcceptQn`), t.getRoomPlayInfo, t.html)
	t.api.Reg(`Live`, t.isValid(`Live`), t.getRoomPlayInfoByQn, t.getRoomPlayInfo, t.html)
	t.api.Reg(`Token`, t.isValid(`Token`), t.getDanmuInfo)
	t.api.Reg(`WSURL`, t.isValid(`WSURL`), t.getDanmuInfo)
	// t.api.Reg(`LIVE_BUVID`, t.isValid(`LIVE_BUVID`), t.getLiveBuvid)
	t.api.Reg(`CheckSwitch_FansMedal`, t.isValid(`CheckSwitch_FansMedal`), t.checkSwitchFansMedal)
	t.api.Reg(`getOnlineGoldRank`, t.isValid(`getOnlineGoldRank`), t.queryContributionRank, t.getOnlineGoldRank)
	t.api.Reg(`Silver2Coin`, t.isValid(`Silver2Coin`), t.silver2Coin)
	return t
}

func (t *GetFuncV2) Get(common *c.Common, key string) {
	if api_limit.TO() {
		return
	} //超额请求阻塞，超时将取消

	t.l.Lock()
	defer t.l.Unlock()
	defer common.Lock()()

	t.common = common
	for node := range t.api.GetTrace(key).Asc() {
		if node.Err != nil {
			apilog.BaseAdd(`Get`).E(node.Key, node.MethodIndex, pe.ErrorFormat(node.Err, pe.ErrActionInLineFunc))
		}
	}
}

func (t *GetFuncV2) isValid(key string) func() bool {
	return func() bool {
		switch key {
		case `UpUid`:
			return t.common.UpUid > 0
		case `Live_Start_Time`:
			return t.common.Live_Start_Time != time.Time{}
		case `Title`:
			return t.common.Title != ``
		case `Uname`:
			return t.common.Uname != ``
		case `ParentAreaID`:
			return t.common.ParentAreaID > 0
		case `AreaID`:
			return t.common.AreaID > 0
		case `GuardNum`:
			return t.common.GuardNum > 0
		case `Note`:
			return t.common.Note != ``
		case `Live_qn`:
			return t.common.Live_qn > 0
		case `AcceptQn`:
			return len(t.common.AcceptQn) > 0
		case `Live`:
			return len(t.common.Live) > 0
		case `Token`:
			return t.common.Token != ``
		case `WSURL`:
			return len(t.common.WSURL) > 0
			// case `LIVE_BUVID`:
			// 	return t.common.LiveBuvidUpdated.After(time.Now().Add(-time.Hour))
		}
		return true
	}
}

// 扫码登录
func (t *GetFuncV2) getCookie() (missKey string, err error) {
	apilog := apilog.BaseAdd(`获取Cookie`)

	savepath := "./cookie.txt"
	if tmp, ok := t.common.K_v.LoadV("cookie路径").(string); ok && tmp != "" {
		savepath = tmp
	}

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	//获取其他Cookie
	defer func() {
		if err := biliApi.GetOtherCookies(); err != nil {
			apilog.E(err)
		}
	}()

	if file.IsExist(savepath) { //读取cookie文件
		if cookieString := string(CookieGet(savepath)); cookieString != `` {
			biliApi.SetCookies(reqf.Cookies_String_2_List(cookieString)) //cookie 存入biliApi
			if biliApi.IsLogin() {
				if e, res := biliApi.GetNav(); e != nil {
					apilog.E(e)
				} else if res.IsLogin {
					// uid
					if e, uid := biliApi.GetCookie(`DedeUserID`); e == nil { //cookie中无DedeUserID
						if uid, e := strconv.Atoi(uid); e == nil {
							t.common.Uid = uid
						}
					}

					t.common.Login = true
					apilog.I(`已登录`)
					return
				}
			}
		}
	}

	t.common.Login = false
	t.common.Uid = 0
	apilog.I(`未登录`)

	if v, ok := t.common.K_v.LoadV(`扫码登录`).(bool); !ok || !v {
		apilog.W(`配置文件已禁止扫码登录，如需登录，修改配置文件"扫码登录"为true`)
		return
	} else {
		apilog.I(`"扫码登录"为true，开始登录`)
	}

	//获取id
	// id := boot_Get_cookie.Flash()
	// defer boot_Get_cookie.UnFlash()

	var img_url string
	var oauth string
	//获取二维码
	if err, imgUrl, QrcodeKey := biliApi.LoginQrCode(); err != nil {
		apilog.E(err)
		return "", pkf.ErrNextMethod.NewErr(err)
	} else {
		img_url = imgUrl
		oauth = QrcodeKey
	}

	//有新实例，退出
	// if boot_Get_cookie.NeedExit(id) {
	// 	return
	// }

	{ //生成二维码
		if e := qr.WriteFile(img_url, qr.Medium, 256, `qr.png`); e != nil || !file.IsExist("qr.png") {
			apilog.E(`qr error`)
			return
		}
		defer func() {
			_ = os.RemoveAll(`qr.png`)
		}()
		//启动web
		if scanPath, ok := t.common.K_v.LoadV("扫码登录路径").(string); ok && scanPath != "" {
			if t.common.K_v.LoadV(`扫码登录自动打开标签页`).(bool) {
				_ = open.Run(`http://127.0.0.1:` + t.common.Stream_url.Port() + scanPath)
			}
			apilog.W(`扫描命令行二维码或打开链接扫码登录：` + t.common.Stream_url.String() + scanPath)
		}

		c := qrterminal.Config{
			Level:     qrterminal.L,
			Writer:    os.Stdout,
			BlackChar: `  `,
			WhiteChar: `OO`,
		}
		if white, ok := t.common.K_v.LoadV(`登录二维码-白`).(string); ok && len(white) != 0 {
			c.WhiteChar = white
		}
		if black, ok := t.common.K_v.LoadV(`登录二维码-黑`).(string); ok && len(black) != 0 {
			c.BlackChar = black
		}
		//show qr code in cmd
		qrterminal.GenerateWithConfig(img_url, c)
		apilog.I(`手机扫命令行二维码登录。如不登录，修改配置文件"扫码登录"为false`)
		time.Sleep(time.Second)
	}

	//有新实例，退出
	// if boot_Get_cookie.NeedExit(id) {
	// 	return
	// }

	{ //循环查看是否通过
		// r := t.common.ReqPool.Get()
		// defer t.common.ReqPool.Put(r)
		for pollC := 10; pollC > 0; pollC-- {
			//3s刷新查看是否通过
			time.Sleep(time.Duration(3) * time.Second)

			//有新实例，退出
			// if boot_Get_cookie.NeedExit(id) {
			// 	return
			// }

			if err, code := biliApi.LoginQrPoll(oauth); err != nil {
				apilog.E(err)
				return "", pkf.ErrNextMethod.NewErr(err)
			} else if code == 0 {
				if cookies := biliApi.GetCookies(); len(cookies) != 0 {
					if e, uid := biliApi.GetCookie(`DedeUserID`); e == nil { //cookie中无DedeUserID
						if uid, e := strconv.Atoi(uid); e == nil {
							t.common.Uid = uid
						}
					}
					apilog.I(`登录,并保存了cookie`)
					return "", nil
				}
			}
		}
		apilog.W(`扫码超时`)
		return "", errors.New(`扫码超时`)
	}
}

func (t *GetFuncV2) getCookieNoBlock() (missKey string, err error) {
	apilog := apilog.BaseAdd(`获取Cookie`)

	savepath := "./cookie.txt"
	if tmp, ok := t.common.K_v.LoadV("cookie路径").(string); ok && tmp != "" {
		savepath = tmp
	}

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if file.IsExist(savepath) { //读取cookie文件
		if cookieString := string(CookieGet(savepath)); cookieString != `` {
			biliApi.SetCookies(reqf.Cookies_String_2_List(cookieString)) //cookie 存入biliApi
			if biliApi.IsLogin() {
				if e, res := biliApi.GetNav(); e != nil {
					apilog.E(e)
				} else if res.IsLogin {
					// uid
					if e, uid := biliApi.GetCookie(`DedeUserID`); e == nil { //cookie中无DedeUserID
						if uid, e := strconv.Atoi(uid); e == nil {
							t.common.Uid = uid
						}
					}

					t.common.Login = true
					apilog.I(`已登录`)
					return
				}
			}
		}
	}

	t.common.Login = false
	t.common.Uid = 0
	apilog.I(`未登录`)

	if v, ok := t.common.K_v.LoadV(`扫码登录`).(bool); !ok || !v {
		apilog.W(`配置文件已禁止扫码登录，如需登录，修改配置文件"扫码登录"为true`)
		return
	} else {
		apilog.I(`"扫码登录"为true，开始登录`)
	}

	var img_url string
	var oauth string
	//获取二维码
	if err, imgUrl, QrcodeKey := biliApi.LoginQrCode(); err != nil {
		apilog.E(err)
		return "", pkf.ErrNextMethod.NewErr(err)
	} else {
		img_url = imgUrl
		oauth = QrcodeKey
	}

	{ //生成二维码
		if e := qr.WriteFile(img_url, qr.Medium, 256, `qr.png`); e != nil || !file.IsExist("qr.png") {
			apilog.E(`qr error`)
			return
		}
		//启动web
		if scanPath, ok := t.common.K_v.LoadV("扫码登录路径").(string); ok && scanPath != "" {
			if t.common.K_v.LoadV(`扫码登录自动打开标签页`).(bool) {
				_ = open.Run(`http://127.0.0.1:` + t.common.Stream_url.Port() + scanPath)
			}
			apilog.W(`扫描命令行二维码或打开链接扫码登录：` + t.common.Stream_url.String() + scanPath)
		}

		c := qrterminal.Config{
			Level:     qrterminal.L,
			Writer:    os.Stdout,
			BlackChar: `  `,
			WhiteChar: `OO`,
		}
		if white, ok := t.common.K_v.LoadV(`登录二维码-白`).(string); ok && len(white) != 0 {
			c.WhiteChar = white
		}
		if black, ok := t.common.K_v.LoadV(`登录二维码-黑`).(string); ok && len(black) != 0 {
			c.BlackChar = black
		}
		//show qr code in cmd
		qrterminal.GenerateWithConfig(img_url, c)
		apilog.I(`手机扫命令行二维码登录。如不登录，修改配置文件"扫码登录"为false`)
	}

	{ //循环查看是否通过
		go func() {
			//获取其他Cookie
			defer func() {
				if err := biliApi.GetOtherCookies(); err != nil {
					apilog.E(err)
				}
				_ = os.RemoveAll(`qr.png`)
			}()
			for pollC := 10; pollC > 0; pollC-- {
				//3s刷新查看是否通过
				time.Sleep(time.Duration(3) * time.Second)

				if err, code := biliApi.LoginQrPoll(oauth); err != nil {
					apilog.E(err)
					return
				} else if code == 0 {
					if cookies := biliApi.GetCookies(); len(cookies) != 0 {
						if e, uid := biliApi.GetCookie(`DedeUserID`); e == nil { //cookie中无DedeUserID
							if uid, e := strconv.Atoi(uid); e == nil {
								t.common.Uid = uid
							}
						}
						apilog.I(`登录,并保存了cookie`)
						return
					}
				}
			}
			apilog.W(`扫码超时`)
		}()
	}
	return "", nil
}

func (t *GetFuncV2) getRoomBaseInfo() (missKey string, err error) {
	fkey := `getRoomBaseInfo`

	if _, ok := t.common.Cache.Load(fkey); ok {
		return
	}

	apilog := apilog.BaseAdd(`getRoomBaseInfo`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.Roomid == 0 {
		return `Roomid`, nil
	}

	//使用其他api
	if err, res := biliApi.GetRoomBaseInfo(t.common.Roomid); err != nil {
		apilog.E(err)
		return "", pkf.ErrNextMethod.NewErr(err)
	} else {
		t.common.UpUid = res.UpUid
		t.common.Uname = res.Uname
		t.common.ParentAreaID = res.ParentAreaID
		t.common.AreaID = res.AreaID
		t.common.Title = res.Title
		t.common.Live_Start_Time = res.LiveStartTime
		t.common.Liveing = res.Liveing
		t.common.Roomid = res.RoomID
	}

	t.common.Cache.Store(fkey, nil, time.Second*2)
	return
}

func (t *GetFuncV2) getInfoByRoom() (missKey string, err error) {

	fkey := `getInfoByRoom`

	if _, ok := t.common.Cache.Load(fkey); ok {
		return
	}

	apilog := apilog.BaseAdd(`getInfoByRoom`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.Roomid == 0 {
		return `Roomid`, nil
	}

	//使用其他api
	if err, res := biliApi.GetInfoByRoom(t.common.Roomid); err != nil {
		apilog.E(err)
		return "", pkf.ErrNextMethod.NewErr(err)
	} else {
		t.common.UpUid = res.UpUid
		t.common.Uname = res.Uname
		t.common.ParentAreaID = res.ParentAreaID
		t.common.AreaID = res.AreaID
		t.common.Title = res.Title
		t.common.Live_Start_Time = res.LiveStartTime
		t.common.Liveing = res.Liveing
		t.common.Roomid = res.RoomID
		t.common.GuardNum = res.GuardNum
		t.common.Note = res.Note
		t.common.Locked = res.Locked
	}

	t.common.Cache.Store(fkey, nil, time.Second*2)

	return
}

func (t *GetFuncV2) getRoomPlayInfo() (missKey string, err error) {
	apilog := apilog.BaseAdd(`getRoomPlayInfo`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.Roomid == 0 {
		return `Roomid`, nil
	}

	//Roominitres
	{
		if err, res := biliApi.GetRoomPlayInfo(t.common.Roomid, 0); err != nil {
			apilog.E(err)
			return "", pkf.ErrNextMethod.NewErr(err)
		} else {
			//主播uid
			t.common.UpUid = res.UpUid
			//房间号（完整）
			t.common.Roomid = res.RoomID
			//直播开始时间
			t.common.Live_Start_Time = res.LiveStartTime
			//是否在直播
			t.common.Liveing = res.Liveing

			//未在直播，不获取直播流
			if !t.common.Liveing {
				t.common.Live_qn = 0
				t.common.AcceptQn = t.common.Qn
				t.common.Live = t.common.Live[:0]
				return "", nil
			}

			//当前直播流
			var s = make([]struct {
				ProtocolName string
				Format       []struct {
					FormatName string
					Codec      []struct {
						CodecName string
						CurrentQn int
						AcceptQn  []int
						BaseURL   string
						URLInfo   []struct {
							Host      string
							Extra     string
							StreamTTL int
						}
						HdrQn     any
						DolbyType int
						AttrName  string
					}
				}
			}, len(res.Streams))
			for i := 0; i < len(res.Streams); i++ {
				s[i] = struct {
					ProtocolName string
					Format       []struct {
						FormatName string
						Codec      []struct {
							CodecName string
							CurrentQn int
							AcceptQn  []int
							BaseURL   string
							URLInfo   []struct {
								Host      string
								Extra     string
								StreamTTL int
							}
							HdrQn     any
							DolbyType int
							AttrName  string
						}
					}
				}(res.Streams[i])
			}
			t.configStreamType(s)
		}
	}
	return
}

func (t *GetFuncV2) html() (missKey string, err error) {
	apilog := apilog.BaseAdd(`html`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.Roomid == 0 {
		return `Roomid`, nil
	}

	//html
	{

		if err, j := biliApi.LiveHtml(t.common.Roomid); err != nil {
			apilog.E(err)
			return "", pkf.ErrNextMethod.NewErr(err)
		} else {
			//Roominitres
			{
				//主播uid
				t.common.UpUid = j.RoomInitRes.Data.UID
				//房间号（完整）
				if j.RoomInitRes.Data.RoomID != 0 {
					t.common.Roomid = j.RoomInitRes.Data.RoomID
				}
				//直播开始时间
				if j.RoomInitRes.Data.LiveTime != 0 {
					t.common.Live_Start_Time = time.Unix(int64(j.RoomInitRes.Data.LiveTime), 0)
				}
				//是否在直播
				t.common.Liveing = j.RoomInitRes.Data.LiveStatus == 1

				//未在直播，不获取直播流
				if !t.common.Liveing {
					t.common.Live_qn = 0
					t.common.AcceptQn = t.common.Qn
					t.common.Live = t.common.Live[:0]
					return "", nil
				}

				//当前直播流
				t.configStreamType(j.RoomInitRes.Data.PlayurlInfo.Playurl.Stream)
			}

			//Roominfores
			{
				//直播间标题
				t.common.Title = j.RoomInfoRes.Data.RoomInfo.Title
				//主播名
				t.common.Uname = j.RoomInfoRes.Data.AnchorInfo.BaseInfo.Uname
				//分区
				t.common.ParentAreaID = j.RoomInfoRes.Data.RoomInfo.ParentAreaID
				//子分区
				t.common.AreaID = j.RoomInfoRes.Data.RoomInfo.AreaID
				//舰长数
				t.common.GuardNum = j.RoomInfoRes.Data.GuardInfo.Count
				//分区排行
				t.common.Note = j.RoomInfoRes.Data.PopularRankInfo.RankName + " "
				if rank := j.RoomInfoRes.Data.PopularRankInfo.Rank; rank > 50 || rank == 0 {
					t.common.Note += "100+"
				} else {
					t.common.Note += strconv.Itoa(rank)
				}
				//直播间是否被封禁
				t.common.Locked = j.RoomInfoRes.Data.RoomInfo.LockStatus == 1
				if t.common.Locked {
					apilog.W("直播间封禁中")
				}
			}
		}
	}
	return
}

func (t *GetFuncV2) getGuardNum() (missKey string, err error) {
	apilog := apilog.BaseAdd(`getGuardNum`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.UpUid == 0 {
		return `UpUid`, nil
	}
	if t.common.Roomid == 0 {
		return `Roomid`, nil
	}

	//Get_guardNum
	if err, GuardNum := biliApi.GetGuardNum(t.common.UpUid, t.common.Roomid); err != nil {
		apilog.E(err)
	} else {
		t.common.GuardNum = GuardNum
	}

	return
}

func (t *GetFuncV2) getPopularAnchorRank() (missKey string, err error) {
	apilog := apilog.BaseAdd(`Get_HotRank`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.UpUid == 0 {
		return `UpUid`, nil
	}
	if t.common.Roomid == 0 {
		return `Roomid`, nil
	}

	//getHotRank
	if err, note := biliApi.GetPopularAnchorRank(t.common.Uid, t.common.UpUid, t.common.Roomid); err != nil {
		apilog.E(err)
	} else {
		t.common.Note = note
	}

	return
}

func (t *GetFuncV2) getRoomPlayInfoByQn() (missKey string, err error) {
	apilog := apilog.BaseAdd(`getRoomPlayInfoByQn`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.Roomid == 0 {
		return `Roomid`, nil
	}

	// 挑选最大的画质
	{
		MaxQn := 0
		for k := range t.common.AcceptQn {
			if k <= t.common.Live_want_qn && k > MaxQn {
				MaxQn = k
			}
		}
		if MaxQn == 0 {
			apilog.W("使用默认")
		} else if t.common.Live_want_qn != MaxQn {
			apilog.W("期望清晰度不可用，使用", t.common.Qn[MaxQn])
		}
		t.common.Live_qn = MaxQn
	}

	//Roominitres
	{
		if err, res := biliApi.GetRoomPlayInfo(t.common.Roomid, t.common.Live_qn); err != nil {
			apilog.E(err)
			return ``, err
		} else {
			//主播uid
			t.common.UpUid = res.UpUid
			//房间号（完整）
			t.common.Roomid = res.RoomID
			//直播开始时间
			t.common.Live_Start_Time = res.LiveStartTime
			//是否在直播
			t.common.Liveing = res.Liveing

			//未在直播，不获取直播流
			if !t.common.Liveing {
				t.common.Live_qn = 0
				t.common.AcceptQn = t.common.Qn
				t.common.Live = t.common.Live[:0]
				return "", nil
			}

			//当前直播流
			var s = make([]struct {
				ProtocolName string
				Format       []struct {
					FormatName string
					Codec      []struct {
						CodecName string
						CurrentQn int
						AcceptQn  []int
						BaseURL   string
						URLInfo   []struct {
							Host      string
							Extra     string
							StreamTTL int
						}
						HdrQn     any
						DolbyType int
						AttrName  string
					}
				}
			}, len(res.Streams))
			for i := 0; i < len(res.Streams); i++ {
				s[i] = struct {
					ProtocolName string
					Format       []struct {
						FormatName string
						Codec      []struct {
							CodecName string
							CurrentQn int
							AcceptQn  []int
							BaseURL   string
							URLInfo   []struct {
								Host      string
								Extra     string
								StreamTTL int
							}
							HdrQn     any
							DolbyType int
							AttrName  string
						}
					}
				}(res.Streams[i])
			}
			t.configStreamType(s)
		}
	}
	return
}

func (t *GetFuncV2) getDanmuInfo() (missKey string, err error) {
	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.Roomid == 0 {
		return `Roomid`, nil
	}

	//GetDanmuInfo
	if err, res := biliApi.GetDanmuInfo(t.common.Roomid); err != nil {
		t.common.Token = ""
		t.common.WSURL = t.common.WSURL[:0]
		apilog.E(err)
		return ``, err
	} else {
		t.common.Token = res.Token
		t.common.WSURL = res.WSURL
	}
	return
}

// LIVE_BUVID
// func (t *GetFuncV2) getLiveBuvid() (missKey string, err error) {
// 	apilog := apilog.BaseAdd(`LIVE_BUVID`)

// 	//当房间处于特殊活动状态时，将会获取不到，此处使用了若干著名up主房间进行尝试
// 	roomIdList := []int{
// 		3, //哔哩哔哩音悦台
// 		2, //直播姬
// 		1, //哔哩哔哩直播
// 	}

// 	req := t.common.ReqPool.Get()
// 	defer t.common.ReqPool.Put(req)
// 	for _, roomid := range roomIdList { //获取
// 		err := biliApi.GetLiveBuvid(roomid)
// 		if err != nil {
// 			apilog.E(err)
// 			return ``, err
// 		}
// 		psync.StoreAll(t.common.Cookie, reqf.Cookies_List_2_Map(biliApi.GetCookies()))
// 		if e, _ := biliApi.GetCookie(`LIVE_BUVID`); e == nil {
// 			apilog.I(`获取到LIVE_BUVID，保存cookie`)
// 			break
// 		} else {
// 			apilog.I(roomid, `未获取到，重试`)
// 			time.Sleep(time.Second)
// 		}
// 	}

// 	t.common.LiveBuvidUpdated = time.Now()

// 	return
// }

func (t *GetFuncV2) checkSwitchFansMedal() (missKey string, err error) {
	apilog := apilog.BaseAdd(`切换粉丝牌`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.UpUid == 0 {
		return `UpUid`, nil
	}

	//验证登录
	if !biliApi.IsLogin() {
		apilog.T(`未登录`)
		return
	}

	{ //获取当前牌子，验证是否本直播间牌子
		res, e := Get_weared_medal(t.common.Uid, t.common.UpUid)
		if e != nil {
			return
		}

		t.common.Wearing_FansMedal = res.Roominfo.RoomID //更新佩戴信息
		if res.TargetID == t.common.UpUid {
			return
		}
	}

	var medal_id int //将要使用的牌子id
	//检查是否有此直播间的牌子
	{
		if err, medal_list := biliApi.GetFansMedal(t.common.Roomid, t.common.UpUid); err != nil {
			apilog.E(err)
			return "", pkf.ErrNextMethod.NewErr(err)
		} else {
			for _, v := range medal_list {
				if v.TargetID != t.common.UpUid {
					continue
				}
				medal_id = v.MedalID
			}
			if medal_id == 0 { //无牌
				apilog.I(`无主播粉丝牌`)
				if t.common.Wearing_FansMedal == 0 { //当前没牌
					return "", nil
				}
			}
		}
	}
	{ //切换牌子
		err := biliApi.SetFansMedal(medal_id)

		if err == nil {
			if medal_id == 0 {
				apilog.I(`已取下粉丝牌`)
			} else {
				apilog.I(`自动切换粉丝牌 id:`, medal_id)
			}
			t.common.Wearing_FansMedal = medal_id //更新佩戴信息
			return "", nil
		}
	}
	return
}

// 获取在线人数
func (t *GetFuncV2) queryContributionRank() (missKey string, err error) {
	apilog := apilog.BaseAdd(`获取在线人数`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.UpUid == 0 {
		return `UpUid`, nil
	}
	if t.common.Roomid == 0 {
		return `Roomid`, nil
	}

	if e, OnlineNum := biliApi.QueryContributionRank(t.common.UpUid, t.common.Roomid); e != nil {
		apilog.E(e)
		return
	} else {
		t.common.OnlineNum = OnlineNum
		apilog.LShow(false).I(`在线人数:`, t.common.OnlineNum)
	}

	return
}

// 获取在线人数
func (t *GetFuncV2) getOnlineGoldRank() (missKey string, err error) {
	apilog := apilog.BaseAdd(`获取在线人数`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	if t.common.UpUid == 0 {
		return `UpUid`, nil
	}
	if t.common.Roomid == 0 {
		return `Roomid`, nil
	}

	if e, OnlineNum := biliApi.GetOnlineGoldRank(t.common.UpUid, t.common.Roomid); e != nil {
		apilog.E(e)
		return
	} else {
		t.common.OnlineNum = OnlineNum
		apilog.LShow(false).I(`在线人数:`, t.common.OnlineNum)
	}

	return
}

// 银瓜子2硬币
func (t *GetFuncV2) silver2Coin() (missKey string, err error) {
	apilog := apilog.BaseAdd(`银瓜子=>硬币`)

	biliApi := biliApi.Inter(func(ce error) BiliApiInter {
		err = ce
		apilog.E(`biliApi组件未构建`, ce)
		return nil
	})
	if biliApi == nil {
		return
	}

	//验证登录
	if !biliApi.IsLogin() {
		apilog.T(`未登录`)
		return
	}

	var Silver int
	//验证是否还有机会
	if e, res := biliApi.GetWalletStatus(); e != nil {
		apilog.E(e)
		return
	} else {
		if res.Silver2CoinLeft == 0 {
			apilog.I(`今天次数已用完`)
			return
		}
		apilog.T(`现在有银瓜子`, res.Silver, `个`)
		Silver = res.Silver
	}

	//获取交换规则，验证数量足够
	if e, Silver2CoinPrice := biliApi.GetWalletRule(); e != nil {
		apilog.E(e)
		return
	} else if Silver < Silver2CoinPrice {
		apilog.I(`当前银瓜子数量不足`)
		return
	}

	//交换
	if e, msg := biliApi.Silver2coin(); e != nil {
		apilog.E(e)
	} else {
		apilog.I(msg)
	}
	return
}

// 配置直播流
func (t *GetFuncV2) configStreamType(sts []struct {
	ProtocolName string
	Format       []struct {
		FormatName string
		Codec      []struct {
			CodecName string
			CurrentQn int
			AcceptQn  []int
			BaseURL   string
			URLInfo   []struct {
				Host      string
				Extra     string
				StreamTTL int
			}
			HdrQn     any
			DolbyType int
			AttrName  string
		}
	}
}) {
	var (
		wantTypes []c.StreamType
		chosen    = -1
	)

	defer func() {
		apilog := apilog.BaseAdd(`configStreamType`)
		if chosen == -1 {
			apilog.E(`未能选择到流`)
			return
		}
		if _, ok := t.common.Qn[t.common.Live_qn]; !ok {
			apilog.W(`未知的清晰度`, t.common.Live_qn)
		}
		apilog.TF("当前有效 %d/%d 条直播流 %s %s %s", t.common.ValidNum(), len(t.common.Live), t.common.Qn[t.common.Live_qn], wantTypes[chosen].Format_name, wantTypes[chosen].Codec_name)
	}()

	// 期望类型
	if v, ok := t.common.K_v.LoadV(`直播流类型`).(string); ok {
		if st, ok := t.common.AllStreamType[v]; ok {
			wantTypes = append(wantTypes, st)
		}
	}
	// 默认类型
	wantTypes = append(wantTypes, t.common.AllStreamType[`fmp4`], t.common.AllStreamType[`flv`])

	// t.common.Live = t.common.Live[:0]
	// for i := 0; i < len(t.common.Live); i++ {
	// 	if time.Now().Add(time.Minute).Before(t.common.Live[i].ReUpTime) {
	// 		t.common.Live = append(t.common.Live[:i], t.common.Live[i+1:]...)
	// 	}
	// }

	for k, streamType := range wantTypes {
		for _, v := range sts {
			if v.ProtocolName != streamType.Protocol_name {
				continue
			}

			for _, v := range v.Format {
				if v.FormatName != streamType.Format_name {
					continue
				}

				for _, v := range v.Codec {
					if v.CodecName != streamType.Codec_name {
						continue
					}

					chosen = k
					//当前直播流质量
					t.common.Live_qn = v.CurrentQn
					if t.common.Live_want_qn == 0 {
						t.common.Live_want_qn = v.CurrentQn
					}
					//允许的清晰度
					{
						var tmp = make(map[int]string)
						for _, v := range v.AcceptQn {
							if s, ok := t.common.Qn[v]; ok {
								tmp[v] = s
							}
						}
						t.common.AcceptQn = tmp
					}
					//直播流链接
					for _, v1 := range v.URLInfo {
						item := c.LiveQn{
							Uuid:       uuid.NewString(),
							Codec:      v.CodecName,
							Url:        v1.Host + v.BaseURL + v1.Extra,
							CreateTime: time.Now(),
						}
						if query, e := url.ParseQuery(v1.Extra); e == nil {
							if expires, e := strconv.Atoi(query.Get("expires")); e == nil {
								item.Expires = time.Now().Add(time.Duration(expires * int(time.Second)))
							}
						}
						t.common.Live = append(t.common.Live, &item)
					}

					// 已选定并设置好参数 退出
					return
				}
			}
		}
	}
}
