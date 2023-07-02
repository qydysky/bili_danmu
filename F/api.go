package F

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
	c "github.com/qydysky/bili_danmu/CV"
	J "github.com/qydysky/bili_danmu/Json"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/exp/slices"

	p "github.com/qydysky/part"
	file "github.com/qydysky/part/file"
	funcCtrl "github.com/qydysky/part/funcCtrl"
	g "github.com/qydysky/part/get"
	limit "github.com/qydysky/part/limit"
	reqf "github.com/qydysky/part/reqf"

	"github.com/mdp/qrterminal/v3"
	qr "github.com/skip2/go-qrcode"
)

var apilog = c.C.Log.Base(`api`)
var api_limit = limit.New(2, "1s", "30s") //频率限制2次/s，最大等待时间30s

type GetFunc struct {
	*c.Common
	l sync.RWMutex
}

type cacheItem struct {
	data     any
	exceeded time.Time
}

func Get(c *c.Common) *GetFunc {
	return &GetFunc{Common: c}
}

func (c *GetFunc) Get(key string) {
	apilog := apilog.Base_add(`Get`)

	if api_limit.TO() {
		return
	} //超额请求阻塞，超时将取消

	var (
		api_can_get = map[string][]func() (missKey []string){
			`Cookie`: { //Cookie
				c.Get_cookie,
			},
			`Uid`: { //用戶uid
				c.GetUid,
			},
			`UpUid`: { //主播uid
				c.getInfoByRoom,
				c.getRoomPlayInfo,
				c.Html,
			},
			`Live_Start_Time`: { //直播开始时间
				c.getInfoByRoom,
				c.getRoomPlayInfo,
				c.Html,
			},
			`Liveing`: { //是否在直播
				c.getInfoByRoom,
				c.getRoomPlayInfo,
				c.Html,
			},
			`Title`: { //直播间标题
				c.getInfoByRoom,
				c.Html,
			},
			`Uname`: { //主播名
				c.getInfoByRoom,
				c.Html,
			},
			`ParentAreaID`: { //分区
				c.getInfoByRoom,
				c.Html,
			},
			`AreaID`: { //子分区
				c.getInfoByRoom,
				c.Html,
			},
			`Roomid`: { //房间id
				c.missRoomId,
			},
			`GuardNum`: { //舰长数
				c.Get_guardNum,
				c.getInfoByRoom,
				c.getRoomPlayInfo,
				c.Html,
			},
			`Note`: { //分区排行
				c.getPopularAnchorRank,
				// c.Get_HotRank,
				c.getInfoByRoom,
				c.Html,
			},
			`Locked`: { //直播间是否被封禁
				c.getInfoByRoom,
				c.Html,
			},
			`Live_qn`: { //当前直播流质量
				c.getRoomPlayInfo,
				c.Html,
			},
			`AcceptQn`: { //允许的清晰度
				c.getRoomPlayInfo,
				c.Html,
			},
			`Live`: { //直播流链接
				c.getRoomPlayInfoByQn,
				c.getRoomPlayInfo,
				c.Html,
			},
			`Token`: { //弹幕钥
				c.getDanmuInfo,
			},
			`WSURL`: { //弹幕链接
				c.getDanmuInfo,
			},
			// `VERSION`:[]func()([]string){//客户版本  不再需要
			// 	Get_Version,
			// },
			`LIVE_BUVID`: { //LIVE_BUVID
				c.Get_LIVE_BUVID,
			},
			`CheckSwitch_FansMedal`: { //切换粉丝牌
				c.CheckSwitch_FansMedal,
			},
			`getOnlineGoldRank`: { //切换粉丝牌
				c.getOnlineGoldRank,
			},
		}
		// 验证是否有效
		check = map[string]func() (valid bool){
			`Uid`: func() bool { //用戶uid
				return c.Uid != 0
			},
			`UpUid`: func() bool { //主播uid
				return c.UpUid != 0
			},
			`Live_Start_Time`: func() bool { //直播开始时间
				return c.Live_Start_Time != time.Time{}
			},
			`Liveing`: func() bool { //是否在直播
				return true
			},
			`Title`: func() bool { //直播间标题
				return c.Title != ``
			},
			`Uname`: func() bool { //主播名
				return c.Uname != ``
			},
			`ParentAreaID`: func() bool { //分区
				return c.ParentAreaID != 0
			},
			`AreaID`: func() bool { //子分区
				return c.AreaID != 0
			},
			`Roomid`: func() bool { //房间id
				return c.Roomid != 0
			},
			`GuardNum`: func() bool { //舰长数
				return c.GuardNum != 0
			},
			`Note`: func() bool { //分区排行
				return c.Note != ``
			},
			`Locked`: func() bool { //直播间是否被封禁
				return true
			},
			`Live_qn`: func() bool { //当前直播流质量
				return c.Live_qn != 0
			},
			`AcceptQn`: func() bool { //允许的清晰度
				return len(c.AcceptQn) != 0
			},
			`Live`: func() bool { //直播流链接
				return len(c.Live) != 0
			},
			`Token`: func() bool { //弹幕钥
				return c.Token != ``
			},
			`WSURL`: func() bool { //弹幕链接
				return len(c.WSURL) != 0
			},
			// `VERSION`:func()(bool){//客户版本  不再需要
			// 	return c.VERSION != `2.0.11`
			// },
			`LIVE_BUVID`: func() bool { //LIVE_BUVID
				return c.LiveBuvidUpdated.After(time.Now().Add(-time.Hour))
			},
			`CheckSwitch_FansMedal`: func() bool { //切换粉丝牌
				return true
			},
			`Cookie`: func() bool { //Cookie
				return true
			},
		}
	)

	if fList, ok := api_can_get[key]; ok {
		for _, fItem := range fList {
			apilog.Log_show_control(false).L(`T: `, `Get`, key)

			c.l.Lock()
			missKey := fItem()
			c.l.Unlock()

			if len(missKey) > 0 {
				apilog.L(`T: `, `missKey when get`, key, missKey)
				for _, misskeyitem := range missKey {
					if checkf, ok := check[misskeyitem]; ok {
						c.l.RLock()
						if checkf() {
							c.l.RUnlock()
							continue
						}
						c.l.RUnlock()
					}
					if misskeyitem == key {
						apilog.L(`W: `, `missKey equrt key`, key, missKey)
						continue
					}
					c.Get(misskeyitem)
				}

				c.l.Lock()
				missKey := fItem()
				c.l.Unlock()

				if len(missKey) > 0 {
					apilog.L(`W: `, `missKey when get`, key, missKey)
					continue
				}
			}
			if checkf, ok := check[key]; ok {
				c.l.RLock()
				if checkf() {
					c.l.RUnlock()
					break
				}
				c.l.RUnlock()
			}
		}
	}
}

func (c *GetFunc) GetUid() (missKey []string) {
	if uid, ok := c.Cookie.LoadV(`DedeUserID`).(string); !ok { //cookie中无DedeUserID
		missKey = append(missKey, `Cookie`)
	} else if uid, e := strconv.Atoi(uid); e != nil {
		missKey = append(missKey, `Cookie`)
	} else {
		c.Uid = uid
	}
	return
}

func (t *GetFunc) Html() (missKey []string) {
	apilog := apilog.Base_add(`html`)

	if t.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
		return
	}

	Roomid := strconv.Itoa(t.Roomid)

	//html
	{
		r := g.Get(reqf.Rval{
			Url:   "https://live.bilibili.com/" + Roomid,
			Proxy: t.Proxy,
		})

		if tmp := r.S(`<script>window.__NEPTUNE_IS_MY_WAIFU__=`, `</script>`, 0, 0); tmp.Err != nil {
			apilog.L(`E: `, `不存在<script>window.__NEPTUNE_IS_MY_WAIFU__= </script>`)
		} else {
			s := tmp.RS[0]

			var j J.NEPTUNE_IS_MY_WAIFU
			if e := json.Unmarshal([]byte(s), &j); e != nil {
				apilog.L(`E: `, e)
				return
			} else if j.RoomInitRes.Code != 0 {
				apilog.L(`E: `, j.RoomInitRes.Message)
				return
			}
			//Roominitres
			{
				//主播uid
				t.UpUid = j.RoomInitRes.Data.UID
				//房间号（完整）
				if j.RoomInitRes.Data.RoomID != 0 {
					t.Roomid = j.RoomInitRes.Data.RoomID
				}
				//直播开始时间
				t.Live_Start_Time = time.Unix(int64(j.RoomInitRes.Data.LiveTime), 0)
				//是否在直播
				t.Liveing = j.RoomInitRes.Data.LiveStatus == 1

				//未在直播，不获取直播流
				if !t.Liveing {
					t.Live_qn = 0
					t.AcceptQn = t.Qn
					t.Live = []c.LiveQn{}
					return
				}

				//当前直播流
				t.configStreamType(j.RoomInitRes.Data.PlayurlInfo.Playurl.Stream)
			}

			//Roominfores
			{
				//直播间标题
				t.Title = j.RoomInfoRes.Data.RoomInfo.Title
				//主播名
				t.Uname = j.RoomInfoRes.Data.AnchorInfo.BaseInfo.Uname
				//分区
				t.ParentAreaID = j.RoomInfoRes.Data.RoomInfo.ParentAreaID
				//子分区
				t.AreaID = j.RoomInfoRes.Data.RoomInfo.AreaID
				//舰长数
				t.GuardNum = j.RoomInfoRes.Data.GuardInfo.Count
				//分区排行
				t.Note = j.RoomInfoRes.Data.PopularRankInfo.RankName + " "
				if rank := j.RoomInfoRes.Data.PopularRankInfo.Rank; rank > 50 || rank == 0 {
					t.Note += "100+"
				} else {
					t.Note += strconv.Itoa(rank)
				}
				//直播间是否被封禁
				if j.RoomInfoRes.Data.RoomInfo.LockStatus == 1 {
					apilog.L(`W: `, "直播间封禁中")
					t.Locked = true
					return
				}
			}
		}
	}
	return
}

// 配置直播流
func (t *GetFunc) configStreamType(sts []J.StreamType) {
	defer apilog.Base_add(`configStreamType`).L(`T: `, fmt.Sprintf("使用直播流 %s %s %s", t.Qn[t.Live_qn], t.StreamType.Format_name, t.StreamType.Codec_name))

	if v, ok := t.Common.K_v.LoadV(`直播流类型`).(string); ok {
		if st, ok := t.AllStreamType[v]; ok {
			t.StreamType = st
		}
	}

	// 查找配置类型是否存在
	for _, v := range sts {
		if v.ProtocolName != t.StreamType.Protocol_name {
			continue
		}

		for _, v := range v.Format {
			if v.FormatName != t.StreamType.Format_name {
				continue
			}

			for _, v := range v.Codec {
				if v.CodecName != t.StreamType.Codec_name {
					continue
				}

				//当前直播流质量
				t.Live_qn = v.CurrentQn
				if t.Live_want_qn == 0 {
					t.Live_want_qn = v.CurrentQn
				}
				//允许的清晰度
				{
					var tmp = make(map[int]string)
					for _, v := range v.AcceptQn {
						if s, ok := t.Qn[v]; ok {
							tmp[v] = s
						}
					}
					t.AcceptQn = tmp
				}
				//直播流链接
				t.Live = []c.LiveQn{}
				for _, v1 := range v.URLInfo {
					item := c.LiveQn{
						Url: v1.Host + v.BaseURL + v1.Extra,
					}

					if query, e := url.ParseQuery(v1.Extra); e == nil {
						if expires, e := strconv.Atoi(query.Get("expires")); e == nil {
							item.Expires = expires
						}
					}

					t.Live = append(t.Live, item)
				}

				return
			}
		}
	}

	apilog.Base_add(`configStreamType`).L(`W: `, "未找到配置的直播流类型，使用默认flv、fmp4")

	// 默认使用flv、fmp4
	for _, streamType := range []c.StreamType{
		t.AllStreamType[`flv`],
		t.AllStreamType[`fmp4`],
	} {

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

					//当前直播流质量
					t.Live_qn = v.CurrentQn
					if t.Live_want_qn == 0 {
						t.Live_want_qn = v.CurrentQn
					}
					//允许的清晰度
					{
						var tmp = make(map[int]string)
						for _, v := range v.AcceptQn {
							if s, ok := t.Qn[v]; ok {
								tmp[v] = s
							}
						}
						t.AcceptQn = tmp
					}
					//直播流链接
					t.Live = []c.LiveQn{}
					for _, v1 := range v.URLInfo {
						item := c.LiveQn{
							Url: v1.Host + v.BaseURL + v1.Extra,
						}

						if query, e := url.ParseQuery(v1.Extra); e == nil {
							if expires, e := strconv.Atoi(query.Get("expires")); e == nil {
								item.Expires = expires
							}
						}

						t.Live = append(t.Live, item)
					}
				}
			}
		}
	}
}

func (c *GetFunc) missRoomId() (missKey []string) {
	apilog.Base_add(`missRoomId`).L(`E: `, `missRoomId`)
	return
}

func (c *GetFunc) getInfoByRoom() (missKey []string) {

	fkey := `getInfoByRoom`

	if v, ok := c.Cache.LoadV(fkey).(cacheItem); ok && v.exceeded.After(time.Now()) {
		return
	}

	apilog := apilog.Base_add(`getInfoByRoom`)

	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
		return
	}

	Roomid := strconv.Itoa(c.Roomid)

	{ //使用其他api
		req := c.Common.ReqPool.Get()
		defer c.Common.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: "https://api.live.bilibili.com/xlive/web-room/v1/index/getInfoByRoom?room_id=" + Roomid,
			Header: map[string]string{
				`Referer`: "https://live.bilibili.com/" + Roomid,
			},
			Proxy:   c.Proxy,
			Timeout: 10 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		//Roominfores
		{
			var j J.Roominfores

			if e := json.Unmarshal(req.Respon, &j); e != nil {
				apilog.L(`E: `, e)
				return
			} else if j.Code != 0 {
				apilog.L(`E: `, j.Message)
				return
			}

			//直播开始时间
			c.Live_Start_Time = time.Unix(int64(j.Data.RoomInfo.LiveStartTime), 0)
			//是否在直播
			c.Liveing = j.Data.RoomInfo.LiveStatus == 1
			//直播间标题
			c.Title = j.Data.RoomInfo.Title
			//主播名
			c.Uname = j.Data.AnchorInfo.BaseInfo.Uname
			//分区
			c.ParentAreaID = j.Data.RoomInfo.ParentAreaID
			//子分区
			c.AreaID = j.Data.RoomInfo.AreaID
			//主播id
			c.UpUid = j.Data.RoomInfo.UID
			//房间id
			if j.Data.RoomInfo.RoomID != 0 {
				c.Roomid = j.Data.RoomInfo.RoomID
			}
			//舰长数
			c.GuardNum = j.Data.GuardInfo.Count
			//分区排行
			c.Note = j.Data.PopularRankInfo.RankName + " "
			if rank := j.Data.PopularRankInfo.Rank; rank > 50 || rank == 0 {
				c.Note += "100+"
			} else {
				c.Note += strconv.Itoa(rank)
			}
			//直播间是否被封禁
			if j.Data.RoomInfo.LockStatus == 1 {
				apilog.L(`W: `, "直播间封禁中")
				c.Locked = true
				return
			}
		}
	}

	c.Cache.Store(fkey, cacheItem{
		data:     nil,
		exceeded: time.Now().Add(time.Second * 2),
	})
	return
}

func (t *GetFunc) getRoomPlayInfo() (missKey []string) {
	apilog := apilog.Base_add(`getRoomPlayInfo`)

	if t.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if t.LiveBuvidUpdated.Before(time.Now().Add(-time.Hour)) {
		missKey = append(missKey, `LIVE_BUVID`)
	}
	if len(missKey) > 0 {
		return
	}

	Roomid := strconv.Itoa(t.Roomid)

	//Roominitres
	{
		Cookie := make(map[string]string)
		t.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := t.Common.ReqPool.Get()
		defer t.Common.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: "https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo?no_playurl=0&mask=1&qn=0&platform=web&protocol=0,1&format=0,2&codec=0,1&room_id=" + Roomid,
			Header: map[string]string{
				`Referer`: "https://live.bilibili.com/" + Roomid,
				`Cookie`:  reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   t.Proxy,
			Timeout: 10 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var j J.GetRoomPlayInfo

		if e := json.Unmarshal([]byte(req.Respon), &j); e != nil {
			apilog.L(`E: `, e)
			return
		} else if j.Code != 0 {
			apilog.L(`E: `, j.Message)
			return
		}

		//主播uid
		t.UpUid = j.Data.UID
		//房间号（完整）
		if j.Data.RoomID != 0 {
			t.Roomid = j.Data.RoomID
		}
		//直播开始时间
		t.Live_Start_Time = time.Unix(int64(j.Data.LiveTime), 0)
		//是否在直播
		t.Liveing = j.Data.LiveStatus == 1

		//未在直播，不获取直播流
		if !t.Liveing {
			t.Live_qn = 0
			t.AcceptQn = t.Qn
			t.Live = []c.LiveQn{}
			return
		}

		//当前直播流
		t.configStreamType(j.Data.PlayurlInfo.Playurl.Stream)
	}
	return
}

func (t *GetFunc) getRoomPlayInfoByQn() (missKey []string) {
	apilog := apilog.Base_add(`getRoomPlayInfoByQn`)

	if t.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if t.LiveBuvidUpdated.Before(time.Now().Add(-time.Hour)) {
		missKey = append(missKey, `LIVE_BUVID`)
	}
	if len(missKey) > 0 {
		return
	}

	{
		AcceptQn := []int{}
		for k := range t.AcceptQn {
			if k <= t.Live_want_qn {
				AcceptQn = append(AcceptQn, k)
			}
		}
		MaxQn := 0
		for i := 0; len(AcceptQn) > i; i += 1 {
			if AcceptQn[i] > MaxQn {
				MaxQn = AcceptQn[i]
			}
		}
		if MaxQn == 0 {
			apilog.L(`W: `, "使用默认")
		}
		t.Live_qn = MaxQn
	}

	Roomid := strconv.Itoa(t.Roomid)

	//Roominitres
	{
		Cookie := make(map[string]string)
		t.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := t.Common.ReqPool.Get()
		defer t.Common.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: "https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo?no_playurl=0&mask=1&qn=" + strconv.Itoa(t.Live_qn) + "&platform=web&protocol=0,1&format=0,2&codec=0,1&room_id=" + Roomid,
			Header: map[string]string{
				`Referer`: "https://live.bilibili.com/" + Roomid,
				`Cookie`:  reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   t.Proxy,
			Timeout: 10 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var j J.GetRoomPlayInfo

		if e := json.Unmarshal([]byte(req.Respon), &j); e != nil {
			apilog.L(`E: `, e)
			return
		} else if j.Code != 0 {
			apilog.L(`E: `, j.Message)
			return
		}

		//主播uid
		t.UpUid = j.Data.UID
		//房间号（完整）
		if j.Data.RoomID != 0 {
			t.Roomid = j.Data.RoomID
		}
		//直播开始时间
		t.Live_Start_Time = time.Unix(int64(j.Data.LiveTime), 0)
		//是否在直播
		t.Liveing = j.Data.LiveStatus == 1

		//未在直播，不获取直播流
		if !t.Liveing {
			t.Live_qn = 0
			t.AcceptQn = t.Qn
			t.Live = []c.LiveQn{}
			return
		}

		//当前直播流
		t.configStreamType(j.Data.PlayurlInfo.Playurl.Stream)
	}
	return
}

func (c *GetFunc) getDanmuInfo() (missKey []string) {
	apilog := apilog.Base_add(`getDanmuInfo`)

	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if c.LiveBuvidUpdated.Before(time.Now().Add(-time.Hour)) {
		missKey = append(missKey, `LIVE_BUVID`)
	}
	if len(missKey) > 0 {
		return
	}

	Roomid := strconv.Itoa(c.Roomid)

	//GetDanmuInfo
	{
		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := c.Common.ReqPool.Get()
		defer c.Common.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?type=0&id=" + Roomid,
			Header: map[string]string{
				`Referer`: "https://live.bilibili.com/" + Roomid,
				`Cookie`:  reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.Proxy,
			Timeout: 10 * 1000,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var j J.GetDanmuInfo

		if e := json.Unmarshal([]byte(req.Respon), &j); e != nil {
			apilog.L(`E: `, e)
			return
		} else if j.Code != 0 {
			apilog.L(`E: `, j.Message)
			return
		}

		//弹幕钥
		c.Token = j.Data.Token
		//弹幕链接
		var tmp []string
		for _, v := range j.Data.HostList {
			if v.WssPort != 443 {
				tmp = append(tmp, "wss://"+v.Host+":"+strconv.Itoa(v.WssPort)+"/sub")
			} else {
				tmp = append(tmp, "wss://"+v.Host+"/sub")
			}
		}
		c.WSURL = tmp
	}
	return
}

func Get_face_src(uid string) string {
	if uid == "" {
		return ""
	}
	if api_limit.TO() {
		return ""
	} //超额请求阻塞，超时将取消
	apilog := apilog.Base_add(`获取头像`)

	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	req := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(req)
	if err := req.Reqf(reqf.Rval{
		Url: "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuMedalAnchorInfo?ruid=" + uid,
		Header: map[string]string{
			`Referer`: "https://live.bilibili.com/" + strconv.Itoa(c.C.Roomid),
			`Cookie`:  reqf.Map_2_Cookies_String(Cookie),
		},
		Proxy:   c.C.Proxy,
		Timeout: 10 * 1000,
		Retry:   2,
	}); err != nil {
		apilog.L(`E: `, err)
		return ""
	}
	res := string(req.Respon)
	if msg := p.Json().GetValFromS(res, "message"); msg == nil || msg != "0" {
		apilog.L(`E: `, "message", msg)
		return ""
	}

	rface := p.Json().GetValFromS(res, "data.rface")
	if rface == nil {
		apilog.L(`E: `, "data.rface", rface)
		return ""
	}
	return rface.(string) + `@58w_58h`
}

func (c *GetFunc) getPopularAnchorRank() (missKey []string) {
	apilog := apilog.Base_add(`Get_HotRank`)

	if c.Uid == 0 {
		missKey = append(missKey, `Uid`)
	}
	if c.UpUid == 0 {
		missKey = append(missKey, `UpUid`)
	}
	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if len(missKey) > 0 {
		return
	}

	Roomid := strconv.Itoa(c.Roomid)

	//getHotRank
	{
		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := c.ReqPool.Get()
		defer c.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: `https://api.live.bilibili.com/xlive/general-interface/v1/rank/getPopularAnchorRank?uid=` + strconv.Itoa(c.Uid) + `&ruid=` + strconv.Itoa(c.UpUid) + `&clientType=2`,
			Header: map[string]string{
				`Host`:            `api.live.bilibili.com`,
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
				`Accept`:          `application/json, text/plain, */*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`:          `https://live.bilibili.com`,
				`Connection`:      `keep-alive`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Referer`:         "https://live.bilibili.com/" + Roomid,
				`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.Proxy,
			Timeout: 3 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var j J.GetPopularAnchorRank

		if e := json.Unmarshal(req.Respon, &j); e != nil {
			apilog.L(`E: `, e)
			return
		} else if j.Code != 0 {
			apilog.L(`E: `, j.Message)
			return
		}

		//获取排名
		c.Note = "人气榜 "
		if j.Data.Anchor.Rank == 0 {
			c.Note += "100+"
		} else {
			c.Note += strconv.Itoa(j.Data.Anchor.Rank)
		}
	}

	return
}

// Deprecated: 2023-01-15
func (c *GetFunc) Get_HotRank() (missKey []string) {
	// apilog := apilog.Base_add(`Get_HotRank`)

	// if c.UpUid == 0 {
	// 	missKey = append(missKey, `UpUid`)
	// }
	// if c.Roomid == 0 {
	// 	missKey = append(missKey, `Roomid`)
	// }
	// if c.ParentAreaID == 0 {
	// 	missKey = append(missKey, `ParentAreaID`)
	// }
	// if !c.LIVE_BUVID {
	// 	missKey = append(missKey, `LIVE_BUVID`)
	// }
	// if len(missKey) > 0 {
	// 	return
	// }

	// Roomid := strconv.Itoa(c.Roomid)

	// //getHotRank
	// {
	// 	Cookie := make(map[string]string)
	// 	c.Cookie.Range(func(k, v interface{}) bool {
	// 		Cookie[k.(string)] = v.(string)
	// 		return true
	// 	})

	// 	reqi := c.ReqPool.Get()
	// 	defer c.ReqPool.Put(reqi)
	// 	req := reqi.Item.(*reqf.Req)
	// 	if err := req.Reqf(reqf.Rval{
	// 		Url: `https://api.live.bilibili.com/xlive/general-interface/v1/rank/getHotRank?ruid=` + strconv.Itoa(c.UpUid) + `&room_id=` + Roomid + `&is_pre=0&page_size=50&source=2&area_id=` + strconv.Itoa(c.ParentAreaID),
	// 		Header: map[string]string{
	// 			`Host`:            `api.live.bilibili.com`,
	// 			`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
	// 			`Accept`:          `application/json, text/plain, */*`,
	// 			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
	// 			`Accept-Encoding`: `gzip, deflate, br`,
	// 			`Origin`:          `https://live.bilibili.com`,
	// 			`Connection`:      `keep-alive`,
	// 			`Pragma`:          `no-cache`,
	// 			`Cache-Control`:   `no-cache`,
	// 			`Referer`:         "https://live.bilibili.com/" + Roomid,
	// 			`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
	// 		},
	// 		Proxy:   c.Proxy,
	// 		Timeout: 3 * 1000,
	// 		Retry:   2,
	// 	}); err != nil {
	// 		apilog.L(`E: `, err)
	// 		return
	// 	}

	// 	var j J.GetHotRank

	// 	if e := json.Unmarshal([]byte(req.Respon), &j); e != nil {
	// 		apilog.L(`E: `, e)
	// 		return
	// 	} else if j.Code != 0 {
	// 		apilog.L(`E: `, j.Message)
	// 		return
	// 	}

	// 	//获取排名
	// 	c.Note = j.Data.Own.AreaName + " "
	// 	if j.Data.Own.Rank == 0 {
	// 		c.Note += "50+"
	// 	} else {
	// 		c.Note += strconv.Itoa(j.Data.Own.Rank)
	// 	}
	// }

	return
}

func (c *GetFunc) Get_guardNum() (missKey []string) {
	apilog := apilog.Base_add(`Get_guardNum`)

	if c.UpUid == 0 {
		missKey = append(missKey, `UpUid`)
	}
	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if c.LiveBuvidUpdated.Before(time.Now().Add(-time.Hour)) {
		missKey = append(missKey, `LIVE_BUVID`)
	}
	if len(missKey) > 0 {
		return
	}

	Roomid := strconv.Itoa(c.Roomid)

	//Get_guardNum
	{
		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := c.ReqPool.Get()
		defer c.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: `https://api.live.bilibili.com/xlive/app-room/v2/guardTab/topList?roomid=` + Roomid + `&page=1&ruid=` + strconv.Itoa(c.UpUid) + `&page_size=29`,
			Header: map[string]string{
				`Host`:            `api.live.bilibili.com`,
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
				`Accept`:          `application/json, text/plain, */*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`:          `https://live.bilibili.com`,
				`Connection`:      `keep-alive`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Referer`:         "https://live.bilibili.com/" + Roomid,
				`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.Proxy,
			Timeout: 3 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var j J.GetGuardNum

		if e := json.Unmarshal([]byte(req.Respon), &j); e != nil {
			apilog.L(`E: `, e)
			return
		} else if j.Code != 0 {
			apilog.L(`E: `, j.Message)
			return
		}

		//获取舰长数
		c.GuardNum = j.Data.Info.Num
	}

	return
}

// func Get_Version() (missKey []string) {  不再需要
// 	if c.Roomid == 0 {
// 		missKey = append(missKey, `Roomid`)
// 	}
// 	if len(missKey) != 0 {return}

// 	Roomid := strconv.Itoa(c.Roomid)

// 	apilog := apilog.Base_add(`获取客户版本`)

// 	var player_js_url string
// 	{//获取player_js_url
// 		r := g.Get(reqf.Rval{
// 			Url:"https://live.bilibili.com/blanc/" + Roomid,
// 		})

// 		if r.Err != nil {
// 			apilog.L(`E: `,r.Err)
// 			return
// 		}

// 		r.S2(`<script src=`,`.js`)
// 		if r.Err != nil {
// 			apilog.L(`E: `,r.Err)
// 			return
// 		}

// 		for _,v := range r.RS {
// 			tmp := string(v) + `.js`
// 			if strings.Contains(tmp,`http`) {continue}
// 			tmp = `https:` + tmp
// 			if strings.Contains(tmp,`player`) {
// 				player_js_url = tmp
// 				break
// 			}
// 		}
// 		if player_js_url == `` {
// 			apilog.L(`E: `,`no found player js`)
// 			return
// 		}
// 	}

// 	{//获取VERSION
// 		r := g.Get(reqf.Rval{
// 			Url:player_js_url,
// 		})

// 		if r.Err != nil {
// 			apilog.L(`E: `,r.Err)
// 			return
// 		}

// 		r.S(`Bilibili HTML5 Live Player v`,` `,0,0)
// 		if r.Err != nil {
// 			apilog.L(`E: `,r.Err)
// 			return
// 		}
// 		c.VERSION = r.RS[0]
// 		apilog.L(`T: `,"api version", c.VERSION)
// 	}
// 	return
// }

func (c *GetFunc) Info(UpUid int) (J.Info, error) {
	fkey := `Info`

	if v, ok := c.Cache.LoadV(fkey).(cacheItem); ok && v.exceeded.After(time.Now()) {
		return (v.data).(J.Info), nil
	}

	// 超额请求阻塞，超时将取消
	apilog := apilog.Base_add(`Info`)
	if api_limit.TO() {
		return J.Info{}, os.ErrDeadlineExceeded
	}

	query := fmt.Sprintf("mid=%d&token=&platform=web&web_location=1550101", UpUid)
	// wbi
	{
		v, e := c.GetNav()
		if e != nil {
			return J.Info{}, e
		}
		wrid, wts := c.getWridWts(query, v.Data.WbiImg.ImgURL, v.Data.WbiImg.SubURL)
		query += "&w_rid=" + wrid + "&wts=" + wts
	}

	// html
	{
		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})
		req := c.ReqPool.Get()
		defer c.ReqPool.Put(req)

		if err := req.Reqf(reqf.Rval{
			Url:     `https://api.bilibili.com/x/space/wbi/acc/info?` + query,
			Proxy:   c.Proxy,
			Timeout: 10 * 1000,
			Retry:   2,
			Header: map[string]string{
				`Accept`: "application/json, text/plain, */*",
				`Cookie`: reqf.Map_2_Cookies_String(Cookie),
			},
		}); err != nil {
			apilog.L(`E: `, err)
			return J.Info{}, err
		}

		var info J.Info

		//Info
		if e := json.Unmarshal(req.Respon, &info); e != nil {
			apilog.L(`E: `, e)
			return J.Info{}, e
		}

		c.Cache.Store(fkey, cacheItem{
			data:     info,
			exceeded: time.Now().Add(time.Hour),
		})
		return info, nil
	}
}

// 调用记录
var boot_Get_cookie funcCtrl.FlashFunc //新的替代旧的

// 是否登录
func (c *GetFunc) GetNav() (J.Nav, error) {
	fkey := `GetNav`

	if v, ok := c.Cache.LoadV(fkey).(cacheItem); ok && v.exceeded.After(time.Now()) {
		return (v.data).(J.Nav), nil
	}

	apilog := apilog.Base_add(`是否登录`)
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return J.Nav{}, os.ErrDeadlineExceeded
	} //超额请求阻塞，超时将取消

	Cookie := make(map[string]string)
	c.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	req := c.ReqPool.Get()
	defer c.ReqPool.Put(req)
	if err := req.Reqf(reqf.Rval{
		Url: `https://api.bilibili.com/x/web-interface/nav`,
		Header: map[string]string{
			`Host`:            `api.bilibili.com`,
			`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
			`Accept`:          `application/json, text/plain, */*`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Origin`:          `https://t.bilibili.com`,
			`Connection`:      `keep-alive`,
			`Pragma`:          `no-cache`,
			`Cache-Control`:   `no-cache`,
			`Referer`:         `https://t.bilibili.com/pages/nav/index_new`,
			`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
		},
		Proxy:   c.Proxy,
		Timeout: 3 * 1000,
		Retry:   2,
	}); err != nil {
		apilog.L(`E: `, err)
		return J.Nav{}, err
	}

	var res J.Nav

	if e := json.Unmarshal(req.Respon, &res); e != nil {
		apilog.L(`E: `, e)
		return J.Nav{}, e
	}

	c.Cache.Store(fkey, cacheItem{
		data:     res,
		exceeded: time.Now().Add(time.Hour),
	})

	return res, nil
}

// 扫码登录
func (t *GetFunc) Get_cookie() (missKey []string) {
	apilog := apilog.Base_add(`获取Cookie`)
	//获取其他Cookie
	defer t.Get_other_cookie()

	if p.Checkfile().IsExist("cookie.txt") { //读取cookie文件
		if cookieString := string(CookieGet()); cookieString != `` {
			for k, v := range reqf.Cookies_String_2_Map(cookieString) { //cookie存入全局变量syncmap
				t.Cookie.Store(k, v)
			}
			if miss := CookieCheck([]string{
				`bili_jct`,
				`DedeUserID`,
			}); len(miss) == 0 {
				if v, e := t.GetNav(); e != nil {
					apilog.L(`E: `, e)
				} else if v.Data.IsLogin {
					apilog.L(`I: `, `已登录`)
					return
				}
			}
		}
	}

	if v, ok := t.K_v.LoadV(`扫码登录`).(bool); !ok || !v {
		apilog.L(`W: `, `配置文件已禁止扫码登录，如需登录，修改配置文件"扫码登录"为true`)
		return
	}

	//获取id
	id := boot_Get_cookie.Flash()
	defer boot_Get_cookie.UnFlash()

	var img_url string
	var oauth string
	{ //获取二维码
		r := t.ReqPool.Get()
		defer t.ReqPool.Put(r)
		if e := r.Reqf(reqf.Rval{
			Url:     `https://passport.bilibili.com/qrcode/getLoginUrl`,
			Proxy:   t.Proxy,
			Timeout: 10 * 1000,
			Retry:   2,
		}); e != nil {
			apilog.L(`E: `, e)
			return
		}
		var res struct {
			Code   int  `json:"code"`
			Status bool `json:"status"`
			Data   struct {
				Url      string `json:"url"`
				OauthKey string `json:"oauthKey"`
			} `json:"data"`
		}
		if e := json.Unmarshal(r.Respon, &res); e != nil {
			apilog.L(`E: `, e)
			return
		}
		if res.Code != 0 {
			apilog.L(`E: `, `code != 0`)
			return
		}
		if !res.Status {
			apilog.L(`E: `, `status == false`)
			return
		}

		if res.Data.Url == `` {
			apilog.L(`E: `, `Data.Urls == ""`)
			return
		} else {
			img_url = res.Data.Url
		}
		if res.Data.OauthKey == `` {
			apilog.L(`E: `, `Data.OauthKey == ""`)
			return
		} else {
			oauth = res.Data.OauthKey
		}
	}

	//有新实例，退出
	if boot_Get_cookie.NeedExit(id) {
		return
	}

	{ //生成二维码
		if e := qr.WriteFile(img_url, qr.Medium, 256, `qr.png`); e != nil || !p.Checkfile().IsExist(`qr.png`) {
			apilog.L(`E: `, `qr error`)
			return
		}
		defer os.RemoveAll(`qr.png`)
		//启动web
		if scanPath, ok := t.K_v.LoadV("扫码登录路径").(string); ok && scanPath != "" {
			t.SerF.Store(scanPath, func(w http.ResponseWriter, r *http.Request) {
				if c.DefaultHttpCheck(t.Common, w, r, http.MethodGet) {
					return
				}
				_ = file.New("qr.png", 0, true).CopyToIoWriter(w, humanize.MByte, true)
			})
			if t.K_v.LoadV(`扫码登录自动打开标签页`).(bool) {
				_ = open.Run(`http://127.0.0.1:` + t.Stream_url.Port() + scanPath)
			}
			apilog.L(`W: `, `或打开链接扫码登录：`+t.Stream_url.String()+scanPath)
		}

		apilog.Block(1000)
		//show qr code in cmd
		qrterminal.GenerateWithConfig(img_url, qrterminal.Config{
			Level:     qrterminal.L,
			Writer:    os.Stdout,
			BlackChar: `  `,
			WhiteChar: `OO`,
		})
		apilog.L(`W: `, `手机扫命令行二维码登录。如不登录，修改配置文件"扫码登录"为false`)
		time.Sleep(time.Second)
	}

	//有新实例，退出
	if boot_Get_cookie.NeedExit(id) {
		return
	}

	{ //循环查看是否通过
		Cookie := make(map[string]string)
		t.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		for {
			//3s刷新查看是否通过
			time.Sleep(time.Duration(3) * time.Second)

			//有新实例，退出
			if boot_Get_cookie.NeedExit(id) {
				return
			}

			r := t.ReqPool.Get()
			defer t.ReqPool.Put(r)
			if e := r.Reqf(reqf.Rval{
				Url:     `https://passport.bilibili.com/qrcode/getLoginInfo`,
				PostStr: `oauthKey=` + oauth,
				Header: map[string]string{
					`Content-Type`: `application/x-www-form-urlencoded; charset=UTF-8`,
					`Referer`:      `https://passport.bilibili.com/login`,
					`Cookie`:       reqf.Map_2_Cookies_String(Cookie),
				},
				Proxy:   t.Proxy,
				Timeout: 10 * 1000,
				Retry:   2,
			}); e != nil {
				apilog.L(`E: `, e)
				return
			}

			var res struct {
				Status  bool   `josn:"status"`
				Message string `json:"message"`
			}

			if e := json.Unmarshal(r.Respon, &res); e != nil {
				apilog.L(`E: `, e.Error(), string(r.Respon))
			}

			if !res.Status {
				if res.Message == `Can't Match oauthKey~` {
					apilog.L(`W: `, `登录超时`)
					return
				}
			} else {
				apilog.L(`W: `, `登录，并保存了cookie`)
				if e := save_cookie(r.Response.Cookies()); e != nil {
					apilog.L(`E: `, e)
				}
				break
			}
		}
	}
	return
}

// 获取其他Cookie
func (c *GetFunc) Get_other_cookie() {
	apilog := apilog.Base_add(`获取其他Cookie`)

	r := c.ReqPool.Get()
	defer c.ReqPool.Put(r)

	Cookie := make(map[string]string)
	c.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	if e := r.Reqf(reqf.Rval{
		Url: `https://www.bilibili.com/`,
		Header: map[string]string{
			`Cookie`: reqf.Map_2_Cookies_String(Cookie),
		},
		Proxy:   c.Proxy,
		Timeout: 10 * 1000,
		Retry:   2,
	}); e != nil {
		apilog.L(`E: `, e)
		return
	}

	if e := save_cookie(r.Response.Cookies()); e != nil {
		apilog.L(`E: `, e)
	}
}

// 短信登录
func Get_cookie_by_msg() {
	/*https://passport.bilibili.com/x/passport-login/web/sms/send*/
}

// 获取牌子信息
func Get_list_in_room() (array []J.GetMyMedals_Items) {

	apilog := apilog.Base_add(`获取牌子`)
	//验证cookie
	if missKey := CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	}); len(missKey) != 0 {
		apilog.L(`T: `, `Cookie无Key:`, missKey)
		return
	}
	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	{ //获取牌子列表
		var medalList []J.GetMyMedals_Items
		for pageNum := 1; true; pageNum += 1 {
			r := c.C.ReqPool.Get()
			defer c.C.ReqPool.Put(r)
			if e := r.Reqf(reqf.Rval{
				Url: `https://api.live.bilibili.com/xlive/app-ucenter/v1/user/GetMyMedals?page=` + strconv.Itoa(pageNum) + `&page_size=10`,
				Header: map[string]string{
					`Cookie`: reqf.Map_2_Cookies_String(Cookie),
				},
				Proxy:   c.C.Proxy,
				Timeout: 10 * 1000,
				Retry:   2,
			}); e != nil {
				apilog.L(`E: `, e)
				return
			}

			var res J.GetMyMedals

			if e := json.Unmarshal(r.Respon, &res); e != nil {
				apilog.L(`E: `, e)
			}

			if res.Code != 0 {
				apilog.L(`E: `, `返回code`, res.Code, res.Message)
				return
			}

			medalList = append(medalList, res.Data.Items...)

			if res.Data.PageInfo.CurPage == res.Data.PageInfo.TotalPage {
				break
			}

			time.Sleep(time.Second)
		}

		return medalList
	}
}

// 获取当前佩戴的牌子
func Get_weared_medal() (item J.GetWearedMedal_Data) {

	apilog := apilog.Base_add(`获取牌子`)
	//验证cookie
	if missKey := CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	}); len(missKey) != 0 {
		apilog.L(`T: `, `Cookie无Key:`, missKey)
		return
	}
	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	{ //获取
		r := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(r)
		if e := r.Reqf(reqf.Rval{
			Url: `https://api.live.bilibili.com/live_user/v1/UserInfo/get_weared_medal`,
			Header: map[string]string{
				`Cookie`: reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.C.Proxy,
			Timeout: 10 * 1000,
			Retry:   2,
		}); e != nil {
			apilog.L(`E: `, e)
			return
		}

		var res J.GetWearedMedal
		if e := json.Unmarshal(r.Respon, &res); e != nil {
			apilog.L(`W: `, e)
			return
		}

		if res.Code != 0 {
			apilog.L(`E: `, `返回code`, res.Code, res.Msg)
			return
		}

		switch res.Data.(type) {
		case []interface{}:
		default:
			if data, err := json.Marshal(res.Data); err == nil {
				_ = json.Unmarshal(data, &item)
			}
		}

		return
	}

}

func (c *GetFunc) CheckSwitch_FansMedal() (missKey []string) {

	if c.LiveBuvidUpdated.Before(time.Now().Add(-time.Hour)) {
		missKey = append(missKey, `LIVE_BUVID`)
	}
	if c.UpUid == 0 {
		missKey = append(missKey, `UpUid`)
	}
	if len(missKey) > 0 {
		return
	}

	apilog := apilog.Base_add(`切换粉丝牌`)
	//验证cookie
	if missCookie := CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	}); len(missCookie) != 0 {
		apilog.L(`T: `, `Cookie无Key:`, missCookie)
		return
	}

	Cookie := make(map[string]string)
	c.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})
	{ //获取当前牌子，验证是否本直播间牌子
		res := Get_weared_medal()

		c.Wearing_FansMedal = res.Roominfo.RoomID //更新佩戴信息
		if res.TargetID == c.UpUid {
			return
		}
	}

	var medal_id int //将要使用的牌子id
	//检查是否有此直播间的牌子
	{
		medal_list := Get_list_in_room()
		for _, v := range medal_list {
			if v.TargetID != c.UpUid {
				continue
			}
			medal_id = v.MedalID
		}
		if medal_id == 0 { //无牌
			apilog.L(`I: `, `无主播粉丝牌`)
			if c.Wearing_FansMedal == 0 { //当前没牌
				return
			}
		}
	}

	var (
		post_url string
		post_str string
	)
	{ //生成佩戴信息
		csrf, _ := c.Cookie.LoadV(`bili_jct`).(string)
		if csrf == `` {
			apilog.L(`E: `, "Cookie错误,无bili_jct=")
			return
		}

		post_str = `csrf_token=` + csrf + `&csrf=` + csrf

		if medal_id == 0 { //无牌，不佩戴牌子
			post_url = `https://api.live.bilibili.com/xlive/web-room/v1/fansMedal/take_off`
		} else {
			post_url = `https://api.live.bilibili.com/xlive/web-room/v1/fansMedal/wear`
			post_str = `medal_id=` + strconv.Itoa(medal_id) + `&` + post_str
		}
	}
	{ //切换牌子
		r := c.ReqPool.Get()
		defer c.ReqPool.Put(r)
		if e := r.Reqf(reqf.Rval{
			Url:     post_url,
			PostStr: post_str,
			Header: map[string]string{
				`Cookie`:       reqf.Map_2_Cookies_String(Cookie),
				`Content-Type`: `application/x-www-form-urlencoded; charset=UTF-8`,
				`Referer`:      `https://passport.bilibili.com/login`,
			},
			Proxy:   c.Proxy,
			Timeout: 10 * 1000,
			Retry:   2,
		}); e != nil {
			apilog.L(`E: `, e)
			return
		}

		var res J.FansMedal

		if e := json.Unmarshal(r.Respon, &res); e != nil {
			apilog.L(`E: `, e)
			return
		} else if res.Code != 0 {
			apilog.L(`E: `, res.Message)
			return
		}

		if res.Message == "佩戴成功" {
			if medal_id == 0 {
				apilog.L(`I: `, `已取下粉丝牌`)
			} else {
				apilog.L(`I: `, `自动切换粉丝牌 id:`, medal_id)
			}
			c.Wearing_FansMedal = medal_id //更新佩戴信息
			return
		}
	}
	return
}

// 签到
func Dosign() {
	apilog := apilog.Base_add(`签到`).L(`T: `, `签到`)
	//验证cookie
	if missKey := CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	}); len(missKey) != 0 {
		apilog.L(`T: `, `Cookie无Key:`, missKey)
		return
	}

	{ //检查是否签到
		Cookie := make(map[string]string)
		c.C.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: `https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/WebGetSignInfo`,
			Header: map[string]string{
				`Host`:            `api.live.bilibili.com`,
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
				`Accept`:          `application/json, text/plain, */*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`:          `https://live.bilibili.com`,
				`Connection`:      `keep-alive`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Referer`:         "https://live.bilibili.com/all",
				`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.C.Proxy,
			Timeout: 3 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var msg struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				Status int `json:"status"`
			} `json:"data"`
		}
		if e := json.Unmarshal(req.Respon, &msg); e != nil {
			apilog.L(`E: `, e)
		}
		if msg.Code != 0 {
			apilog.L(`E: `, msg.Message)
			return
		}
		if msg.Data.Status == 1 { //今日已签到
			return
		}
	}

	{ //签到
		Cookie := make(map[string]string)
		c.C.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: `https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/DoSign`,
			Header: map[string]string{
				`Host`:            `api.live.bilibili.com`,
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
				`Accept`:          `application/json, text/plain, */*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`:          `https://live.bilibili.com`,
				`Connection`:      `keep-alive`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Referer`:         "https://live.bilibili.com/all",
				`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.C.Proxy,
			Timeout: 3 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var msg struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				HadSignDays int `json:"hadSignDays"`
			} `json:"data"`
		}
		if e := json.Unmarshal(req.Respon, &msg); e != nil {
			apilog.L(`E: `, e)
		}
		if msg.Code == 0 {
			apilog.L(`I: `, `签到成功!本月已签到`, msg.Data.HadSignDays, `天`)
			return
		}
		apilog.L(`E: `, msg.Message)
	}
}

// LIVE_BUVID
func (c *GetFunc) Get_LIVE_BUVID() (missKey []string) {
	apilog := apilog.Base_add(`LIVE_BUVID`)

	//当房间处于特殊活动状态时，将会获取不到，此处使用了若干著名up主房间进行尝试
	roomIdList := []string{
		"3", //哔哩哔哩音悦台
		"2", //直播姬
		"1", //哔哩哔哩直播
	}

	for _, roomid := range roomIdList { //获取
		req := c.ReqPool.Get()
		defer c.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: `https://api.live.bilibili.com/live/getRoomKanBanModel?roomid=` + roomid,
			Header: map[string]string{
				`Host`:                      `live.bilibili.com`,
				`User-Agent`:                `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
				`Accept`:                    `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8`,
				`Accept-Language`:           `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`:           `gzip, deflate, br`,
				`Connection`:                `keep-alive`,
				`Cache-Control`:             `no-cache`,
				`Referer`:                   "https://live.bilibili.com",
				`DNT`:                       `1`,
				`Upgrade-Insecure-Requests`: `1`,
			},
			Proxy:   c.Proxy,
			Timeout: 3 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		//cookie
		_ = save_cookie(req.Response.Cookies())
		var has bool
		for k := range reqf.Cookies_List_2_Map(req.Response.Cookies()) {
			if k == `LIVE_BUVID` {
				has = true
			}
		}
		if has {
			apilog.L(`I: `, `获取到LIVE_BUVID，保存cookie`)
			break
		} else {
			apilog.L(`I: `, roomid, `未获取到，重试`)
			time.Sleep(time.Second)
		}
	}

	c.LiveBuvidUpdated = time.Now()

	return
}

// 礼物列表
type Gift_list_type struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    Gift_list_type_Data `json:"data"`
}

type Gift_list_type_Data struct {
	List []Gift_list_type_Data_List `json:"list"`
}

type Gift_list_type_Data_List struct {
	Bag_id    int    `json:"bag_id"`
	Gift_id   int    `json:"gift_id"`
	Gift_name string `json:"gift_name"`
	Gift_num  int    `json:"gift_num"`
	Expire_at int    `json:"expire_at"`
}

func Gift_list() (list []Gift_list_type_Data_List) {
	apilog := apilog.Base_add(`礼物列表`)
	//验证cookie
	if missKey := CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	}); len(missKey) != 0 {
		apilog.L(`T: `, `Cookie无Key:`, missKey)
		return
	}
	if c.C.Roomid == 0 {
		apilog.L(`E: `, `失败！无Roomid`)
		return
	}
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	req := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(req)
	if err := req.Reqf(reqf.Rval{
		Url: `https://api.live.bilibili.com/xlive/web-room/v1/gift/bag_list?t=` + strconv.Itoa(int(time.Now().UnixNano()/int64(time.Millisecond))) + `&room_id=` + strconv.Itoa(c.C.Roomid),
		Header: map[string]string{
			`Host`:            `api.live.bilibili.com`,
			`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
			`Accept`:          `application/json, text/plain, */*`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Origin`:          `https://live.bilibili.com`,
			`Connection`:      `keep-alive`,
			`Pragma`:          `no-cache`,
			`Cache-Control`:   `no-cache`,
			`Referer`:         "https://live.bilibili.com/" + strconv.Itoa(c.C.Roomid),
			`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
		},
		Proxy:   c.C.Proxy,
		Timeout: 3 * 1000,
		Retry:   2,
	}); err != nil {
		apilog.L(`E: `, err)
		return
	}

	var res Gift_list_type

	if e := json.Unmarshal(req.Respon, &res); e != nil {
		apilog.L(`E: `, e)
		return
	}

	if res.Code != 0 {
		apilog.L(`E: `, res.Message)
		return
	}

	apilog.L(`T: `, `成功`)
	return res.Data.List
}

// 银瓜子2硬币
func (c *GetFunc) Silver_2_coin() (missKey []string) {
	apilog := apilog.Base_add(`银瓜子=>硬币`)

	if c.LiveBuvidUpdated.Before(time.Now().Add(-time.Hour)) {
		missKey = append(missKey, `LIVE_BUVID`)
	}
	if len(missKey) > 0 {
		return
	}

	//验证cookie
	if miss := CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	}); len(miss) != 0 {
		apilog.L(`T: `, `Cookie无Key:`, miss)
		return
	}

	var Silver int
	{ //验证是否还有机会
		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := c.ReqPool.Get()
		defer c.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: `https://api.live.bilibili.com/xlive/revenue/v1/wallet/getStatus`,
			Header: map[string]string{
				`Host`:            `api.live.bilibili.com`,
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
				`Accept`:          `application/json, text/plain, */*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`:          `https://link.bilibili.com`,
				`Connection`:      `keep-alive`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Referer`:         `https://link.bilibili.com/p/center/index`,
				`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.Proxy,
			Timeout: 3 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var j J.ApiXliveRevenueV1WalletGetStatus

		if e := json.Unmarshal([]byte(req.Respon), &j); e != nil {
			apilog.L(`E: `, e)
			return
		} else if j.Code != 0 {
			apilog.L(`E: `, j.Message)
			return
		}

		if j.Data.Silver2CoinLeft == 0 {
			apilog.L(`I: `, `今天次数已用完`)
			return
		}

		apilog.L(`T: `, `现在有银瓜子`, j.Data.Silver, `个`)
		Silver = j.Data.Silver
	}

	{ //获取交换规则，验证数量足够
		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := c.ReqPool.Get()
		defer c.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: `https://api.live.bilibili.com/xlive/revenue/v1/wallet/getRule`,
			Header: map[string]string{
				`Host`:            `api.live.bilibili.com`,
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
				`Accept`:          `application/json, text/plain, */*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`:          `https://link.bilibili.com`,
				`Connection`:      `keep-alive`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Referer`:         `https://link.bilibili.com/p/center/index`,
				`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.Proxy,
			Timeout: 3 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var j J.ApixliveRevenueV1WalletGetRule

		if e := json.Unmarshal([]byte(req.Respon), &j); e != nil {
			apilog.L(`E: `, e)
			return
		} else if j.Code != 0 {
			apilog.L(`E: `, j.Message)
			return
		}

		if Silver < j.Data.Silver2CoinPrice {
			apilog.L(`W: `, `当前银瓜子数量不足`)
			return
		}
	}

	{ //交换
		csrf, _ := c.Cookie.LoadV(`bili_jct`).(string)
		if csrf == `` {
			apilog.L(`E: `, "Cookie错误,无bili_jct=")
			return
		}

		post_str := `csrf_token=` + csrf + `&csrf=` + csrf

		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		req := c.ReqPool.Get()
		defer c.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url:     `https://api.live.bilibili.com/xlive/revenue/v1/wallet/silver2coin`,
			PostStr: url.PathEscape(post_str),
			Header: map[string]string{
				`Host`:            `api.live.bilibili.com`,
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
				`Accept`:          `application/json, text/plain, */*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`:          `https://link.bilibili.com`,
				`Connection`:      `keep-alive`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Content-Type`:    `application/x-www-form-urlencoded`,
				`Referer`:         `https://link.bilibili.com/p/center/index`,
				`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.Proxy,
			Timeout: 3 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		_ = save_cookie(req.Response.Cookies())

		var res struct {
			Code    int    `json:"code"`
			Msg     string `json:"msg"`
			Message string `json:"message"`
		}

		if e := json.Unmarshal(req.Respon, &res); e != nil {
			apilog.L(`E: `, e)
			return
		}

		if res.Code != 0 {
			apilog.L(`E: `, res.Message)
			return
		}
		apilog.L(`I: `, res.Message)
	}
	return
}

func save_cookie(Cookies []*http.Cookie) error {
	if len(Cookies) == 0 {
		return errors.New("no cookie")
	}

	for k, v := range reqf.Cookies_List_2_Map(Cookies) {
		c.C.Cookie.Store(k, v)
	}

	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})
	CookieSet([]byte(reqf.Map_2_Cookies_String(Cookie)))
	return nil
}

// 正在直播主播
type UpItem struct {
	Uname      string `json:"uname"`
	Title      string `json:"title"`
	Roomid     int    `json:"roomid"`
	LiveStatus int    `json:"live_status"`
}

// 获取历史观看 直播
func GetHisStream() (Uplist []UpItem) {
	apilog := apilog.Base_add(`历史直播主播`).L(`T: `, `获取中`)
	defer apilog.L(`T: `, `完成`)
	//验证cookie
	if missKey := CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	}); len(missKey) != 0 {
		apilog.L(`T: `, `Cookie无Key:`, missKey)
		return
	}
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	req := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(req)
	if err := req.Reqf(reqf.Rval{
		Url: `https://api.bilibili.com/x/web-interface/history/cursor?type=live&ps=10`,
		Header: map[string]string{
			`Host`:            `api.live.bilibili.com`,
			`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
			`Accept`:          `application/json, text/plain, */*`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Origin`:          `https://t.bilibili.com`,
			`Connection`:      `keep-alive`,
			`Pragma`:          `no-cache`,
			`Cache-Control`:   `no-cache`,
			`Referer`:         `https://t.bilibili.com/pages/nav/index_new`,
			`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
		},
		Proxy:   c.C.Proxy,
		Timeout: 3 * 1000,
		Retry:   2,
	}); err != nil {
		apilog.L(`E: `, err)
		return
	}

	var res J.History

	if e := json.Unmarshal(req.Respon, &res); e != nil {
		apilog.L(`E: `, e)
		return
	}

	if res.Code != 0 {
		apilog.L(`E: `, res.Message)
		return
	}

	// 提前结束获取，仅获取当前正在直播的主播
	for _, item := range res.Data.List {
		Uplist = append(Uplist, UpItem{
			Uname:      item.AuthorName,
			Title:      item.Title,
			Roomid:     item.Kid,
			LiveStatus: item.LiveStatus,
		})
	}
	return
}

// 进入房间
func RoomEntryAction(roomId int) {
	apilog := apilog.Base_add(`进入房间`)
	//验证cookie
	if missKey := CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	}); len(missKey) != 0 {
		apilog.L(`T: `, `Cookie无Key:`, missKey)
		return
	}
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	csrf := Cookie[`bili_jct`]
	if csrf == `` {
		apilog.L(`E: `, "Cookie错误,无bili_jct=")
		return
	}

	req := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(req)
	if err := req.Reqf(reqf.Rval{
		Url:     `https://api.live.bilibili.com/xlive/web-room/v1/index/roomEntryAction`,
		PostStr: fmt.Sprintf("room_id=%d&platform=pc&csrf_token=%s&csrf=%s&visit_id=", roomId, csrf, csrf),
		Header: map[string]string{
			`Host`:            `api.live.bilibili.com`,
			`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
			`Accept`:          `application/json, text/plain, */*`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Origin`:          `https://live.bilibili.com`,
			`Connection`:      `keep-alive`,
			`Pragma`:          `no-cache`,
			`Cache-Control`:   `no-cache`,
			`Referer`:         fmt.Sprintf("https://live.bilibili.com/%d", roomId),
			`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
		},
		Proxy:   c.C.Proxy,
		Timeout: 3 * 1000,
		Retry:   2,
	}); err != nil {
		apilog.L(`E: `, err)
		return
	}

	var res J.RoomEntryAction

	if e := json.Unmarshal(req.Respon, &res); e != nil {
		apilog.L(`E: `, e)
		return
	}

	if res.Code != 0 {
		apilog.L(`E: `, res.Message)
		return
	}
}

// 获取在线人数
func (c *GetFunc) getOnlineGoldRank() (misskey []string) {
	apilog := apilog.Base_add(`获取在线人数`)
	if c.UpUid == 0 {
		misskey = append(misskey, `UpUid`)
		return
	}
	if c.Roomid == 0 {
		misskey = append(misskey, `Roomid`)
		return
	}

	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	req := c.ReqPool.Get()
	defer c.ReqPool.Put(req)

	if err := req.Reqf(reqf.Rval{
		Url: fmt.Sprintf("https://api.live.bilibili.com/xlive/general-interface/v1/rank/getOnlineGoldRank?ruid=%d&roomId=%d&page=1&pageSize=10", c.UpUid, c.Roomid),
		Header: map[string]string{
			`Host`:            `api.live.bilibili.com`,
			`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
			`Accept`:          `application/json, text/plain, */*`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Origin`:          `https://live.bilibili.com`,
			`Connection`:      `keep-alive`,
			`Pragma`:          `no-cache`,
			`Cache-Control`:   `no-cache`,
			`Referer`:         fmt.Sprintf("https://live.bilibili.com/%d", c.Roomid),
		},
		Proxy:   c.Proxy,
		Timeout: 3 * 1000,
	}); err != nil {
		apilog.L(`E: `, err)
		return
	}

	var res J.GetOnlineGoldRank

	if e := json.Unmarshal(req.Respon, &res); e != nil {
		apilog.L(`E: `, e)
		return
	}

	if res.Code != 0 {
		apilog.L(`E: `, res.Message)
		return
	}

	c.OnlineNum = res.Data.OnlineNum
	apilog.Log_show_control(false).L(`I: `, `在线人数:`, c.OnlineNum)
	return
}

func Feed_list() (Uplist []J.FollowingDataList) {
	apilog := apilog.Base_add(`正在直播主播`).L(`T: `, `获取中`)
	defer apilog.L(`T: `, `完成`)
	//验证cookie
	if missKey := CookieCheck([]string{
		`bili_jct`,
		`DedeUserID`,
		`LIVE_BUVID`,
	}); len(missKey) != 0 {
		apilog.L(`T: `, `Cookie无Key:`, missKey)
		return
	}
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	req := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(req)
	for pageNum := 1; true; pageNum += 1 {
		if err := req.Reqf(reqf.Rval{
			Url: `https://api.live.bilibili.com/xlive/web-ucenter/user/following?page=` + strconv.Itoa(pageNum) + `&page_size=10`,
			Header: map[string]string{
				`Host`:            `api.live.bilibili.com`,
				`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0`,
				`Accept`:          `application/json, text/plain, */*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`:          `https://t.bilibili.com`,
				`Connection`:      `keep-alive`,
				`Pragma`:          `no-cache`,
				`Cache-Control`:   `no-cache`,
				`Referer`:         `https://t.bilibili.com/pages/nav/index_new`,
				`Cookie`:          reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.C.Proxy,
			Timeout: 3 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var res J.Following

		if e := json.Unmarshal(req.Respon, &res); e != nil {
			apilog.L(`E: `, e)
			return
		}

		if res.Code != 0 {
			apilog.L(`E: `, res.Message)
			return
		}

		// 提前结束获取，仅获取当前正在直播的主播
		for _, item := range res.Data.List {
			if item.LiveStatus == 0 {
				break
			} else {
				Uplist = append(Uplist, item)
			}
		}

		if pageNum*10 > res.Data.TotalPage {
			break
		}
		time.Sleep(time.Second)
	}

	return
}

func GetHistory(Roomid_int int) (j J.GetHistory) {
	apilog := apilog.Base_add(`GetHistory`)

	Roomid := strconv.Itoa(Roomid_int)

	{ //使用其他api
		req := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(req)
		if err := req.Reqf(reqf.Rval{
			Url: "https://api.live.bilibili.com/xlive/web-room/v1/dM/gethistory?roomid=" + Roomid,
			Header: map[string]string{
				`Referer`: "https://live.bilibili.com/" + Roomid,
			},
			Proxy:   c.C.Proxy,
			Timeout: 10 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		//GetHistory
		{
			if e := json.Unmarshal(req.Respon, &j); e != nil {
				apilog.L(`E: `, e)
				return
			} else if j.Code != 0 {
				apilog.L(`E: `, j.Message)
				return
			}
		}
	}
	return
}

type searchresult struct {
	Roomid  int
	Uname   string
	Is_live bool
}

func (c *GetFunc) SearchUP(s string) (list []searchresult) {
	apilog := apilog.Base_add(`搜索主播`)
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	{ //使用其他api
		req := c.ReqPool.Get()
		defer c.ReqPool.Put(req)

		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		query := "page=1&page_size=10&order=online&platform=pc&search_type=live_user&keyword=" + url.PathEscape(s)
		// wbi
		{
			v, e := c.GetNav()
			if e != nil {
				apilog.L(`E: `, e)
				return
			}
			wrid, wts := c.getWridWts(query, v.Data.WbiImg.ImgURL, v.Data.WbiImg.SubURL)
			query += "&w_rid=" + wrid + "&wts=" + wts
		}

		if err := req.Reqf(reqf.Rval{
			Url:   "https://api.bilibili.com/x/web-interface/wbi/search/type?" + query,
			Proxy: c.Proxy,
			Header: map[string]string{
				`Cookie`: reqf.Map_2_Cookies_String(Cookie),
			},
			Timeout: 10 * 1000,
			Retry:   2,
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		var j J.Search

		//Search
		{
			if e := json.Unmarshal(req.Respon, &j); e != nil {
				apilog.L(`E: `, e)
				return
			} else if j.Code != 0 {
				apilog.L(`E: `, j.Message)
				return
			}
		}

		if j.Data.NumResults == 0 {
			apilog.L(`I: `, `没有匹配`)
			return
		}

		for i := 0; i < len(j.Data.Result); i += 1 {
			uname := strings.ReplaceAll(j.Data.Result[i].Uname, `<em class="keyword">`, ``)
			uname = strings.ReplaceAll(uname, `</em>`, ``)
			list = append(list, searchresult{
				Roomid:  j.Data.Result[i].Roomid,
				Uname:   uname,
				Is_live: j.Data.Result[i].IsLive,
			})
		}

	}

	return
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

	req := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(req)
	if err := req.Reqf(reqf.Rval{
		Url:              "https://www.bilibili.com",
		Proxy:            c.C.Proxy,
		Timeout:          10 * 1000,
		JustResponseCode: true,
	}); err != nil {
		apilog.L(`W: `, `网络中断`, err)
		return false
	}

	apilog.L(`T: `, `已连接`)
	return true
}

// bilibili wrid wts 计算
func (c *GetFunc) getWridWts(query string, imgURL, subURL string, customWts ...string) (w_rid, wts string) {
	wbi := imgURL[strings.LastIndex(imgURL, "/")+1:strings.LastIndex(imgURL, ".")] +
		subURL[strings.LastIndex(subURL, "/")+1:strings.LastIndex(subURL, ".")]

	code := []int{46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5,
		49, 33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55,
		40, 61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57,
		62, 11, 36, 20, 34, 44, 52}

	s := []byte{}

	for i := 0; i < len(code); i++ {
		if code[i] < len(wbi) {
			s = append(s, wbi[code[i]])
			if len(s) >= 32 {
				break
			}
		}
	}

	object := strings.Split(query, "&")

	if len(customWts) == 0 {
		wts = fmt.Sprintf("%d", time.Now().Unix())
	} else {
		wts = customWts[0]
	}
	object = append(object, "wts="+wts)

	slices.Sort(object)

	for i := 0; i < len(object); i++ {
		object[i] = url.PathEscape(object[i])
	}

	w_rid = fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(object, "&")+string(s))))

	return
}
