package F

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	J "github.com/qydysky/bili_danmu/Json"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/exp/slices"

	_ "github.com/qydysky/biliApi"
	p "github.com/qydysky/part"
	cmp "github.com/qydysky/part/component2"
	file "github.com/qydysky/part/file"
	funcCtrl "github.com/qydysky/part/funcCtrl"
	pio "github.com/qydysky/part/io"
	limit "github.com/qydysky/part/limit"
	reqf "github.com/qydysky/part/reqf"

	"github.com/mdp/qrterminal/v3"
	qr "github.com/skip2/go-qrcode"
)

const id = "github.com/qydysky/bili_danmu/F.biliApi"

var apilog = c.C.Log.Base(`api`)
var api_limit = limit.New(2, "1s", "30s") //频率限制2次/s，最大等待时间30s

var biliApi = cmp.Get(id, func(ba biliApiInter) biliApiInter {
	ba.SetProxy(c.C.Proxy)
	ba.SetReqPool(c.C.ReqPool)
	return ba
})

type GetFunc struct {
	*c.Common
	count atomic.Int32
	l     sync.RWMutex
}

// type cacheItem struct {
// 	data     any
// 	exceeded time.Time
// }

func Get(c *c.Common) *GetFunc {
	return &GetFunc{Common: c}
}

func (c *GetFunc) Get(key string) {
	apilog := apilog.Base_add(`Get`)

	current := c.count.Add(1)
	defer c.count.Add(-1)

	if current > 10 {
		apilog.L(`E: `, `max loop`)
		return
	}

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
				c.getRoomBaseInfo,
				c.getInfoByRoom,
				c.getRoomPlayInfo,
				c.Html,
			},
			`Live_Start_Time`: { //直播开始时间
				c.getRoomBaseInfo,
				c.getInfoByRoom,
				c.getRoomPlayInfo,
				c.Html,
			},
			`Liveing`: { //是否在直播
				c.getRoomBaseInfo,
				c.getInfoByRoom,
				c.getRoomPlayInfo,
				c.Html,
			},
			`Title`: { //直播间标题
				c.getRoomBaseInfo,
				c.getInfoByRoom,
				c.Html,
			},
			`Uname`: { //主播名
				c.getRoomBaseInfo,
				c.getInfoByRoom,
				c.Html,
			},
			`ParentAreaID`: { //分区
				c.getRoomBaseInfo,
				c.getInfoByRoom,
				c.Html,
			},
			`AreaID`: { //子分区
				c.getRoomBaseInfo,
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

	if fList, ok := api_can_get[key]; !ok {
		apilog.L(`E: `, `no api`, key)
	} else {
		for i := 0; i < len(fList); i++ {
			apilog.Log_show_control(false).L(`T: `, `Get`, key)

			c.l.Lock()
			missKey := fList[i]()
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
				missKey := fList[i]()
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
				} else {
					c.l.RUnlock()
					apilog.L(`W: `, `check fail`, key)
					c.Get(key)
				}
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

	//html
	{

		if err, j := biliApi.LiveHtml(t.Roomid); err != nil {
			apilog.L(`E: `, err)
			return
		} else {
			//Roominitres
			{
				//主播uid
				t.UpUid = j.RoomInitRes.Data.UID
				//房间号（完整）
				if j.RoomInitRes.Data.RoomID != 0 {
					t.Roomid = j.RoomInitRes.Data.RoomID
				}
				//直播开始时间
				if j.RoomInitRes.Data.LiveTime != 0 {
					t.Live_Start_Time = time.Unix(int64(j.RoomInitRes.Data.LiveTime), 0)
				}
				//是否在直播
				t.Liveing = j.RoomInitRes.Data.LiveStatus == 1

				//未在直播，不获取直播流
				if !t.Liveing {
					t.Live_qn = 0
					t.AcceptQn = t.Qn
					t.Live = t.Live[:0]
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
				t.Locked = j.RoomInfoRes.Data.RoomInfo.LockStatus == 1
				if t.Locked {
					apilog.L(`W: `, "直播间封禁中")
				}
			}
		}
	}
	return
}

// 配置直播流
func (t *GetFunc) configStreamType(sts []struct {
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
		chosen    int = -1
	)

	defer func() {
		apilog := apilog.Base_add(`configStreamType`)
		if chosen == -1 {
			apilog.L(`E: `, `未能选择到流`)
			return
		}
		if _, ok := t.Qn[t.Live_qn]; !ok {
			apilog.L(`W: `, `未知的清晰度`, t.Live_qn)
		}
		apilog.L(`T: `, fmt.Sprintf("获取到 %d 条直播流 %s %s %s", len(t.Live), t.Qn[t.Live_qn], wantTypes[chosen].Format_name, wantTypes[chosen].Codec_name))
	}()

	// 期望类型
	if v, ok := t.Common.K_v.LoadV(`直播流类型`).(string); ok {
		if st, ok := t.AllStreamType[v]; ok {
			wantTypes = append(wantTypes, st)
		}
	}
	// 默认类型
	wantTypes = append(wantTypes, t.AllStreamType[`fmp4`], t.AllStreamType[`flv`])

	t.Live = t.Live[:0]

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
					for _, v1 := range v.URLInfo {
						item := c.LiveQn{
							Url: v1.Host + v.BaseURL + v1.Extra,
						}

						if query, e := url.ParseQuery(v1.Extra); e == nil {
							if expires, e := strconv.Atoi(query.Get("expires")); e == nil {
								item.Expires = expires
							}
						}

						t.Live = append(t.Live, &item)
					}

					// 已选定并设置好参数 退出
					return
				}
			}
		}
	}
}

func (c *GetFunc) missRoomId() (missKey []string) {
	apilog.Base_add(`missRoomId`).L(`E: `, `missRoomId`)
	return
}

func (c *GetFunc) getRoomBaseInfo() (missKey []string) {
	fkey := `getRoomBaseInfo`

	if _, ok := c.Cache.Load(fkey); ok {
		return
	}

	apilog := apilog.Base_add(`getRoomBaseInfo`)

	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
		return
	}

	//使用其他api
	if err, res := biliApi.GetRoomBaseInfo(c.Roomid); err != nil {
		apilog.L(`E: `, err)
		return
	} else {
		c.UpUid = res.UpUid
		c.Uname = res.Uname
		c.ParentAreaID = res.ParentAreaID
		c.AreaID = res.AreaID
		c.Title = res.Title
		c.Live_Start_Time = res.LiveStartTime
		c.Liveing = res.Liveing
		c.Roomid = res.RoomID
	}

	c.Cache.Store(fkey, nil, time.Second*2)
	return
}

func (c *GetFunc) getInfoByRoom() (missKey []string) {

	fkey := `getInfoByRoom`

	if _, ok := c.Cache.Load(fkey); ok {
		return
	}

	apilog := apilog.Base_add(`getInfoByRoom`)

	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
		return
	}

	//使用其他api
	if err, res := biliApi.GetInfoByRoom(c.Roomid); err != nil {
		apilog.L(`E: `, err)
		return
	} else {
		c.UpUid = res.UpUid
		c.Uname = res.Uname
		c.ParentAreaID = res.ParentAreaID
		c.AreaID = res.AreaID
		c.Title = res.Title
		c.Live_Start_Time = res.LiveStartTime
		c.Liveing = res.Liveing
		c.Roomid = res.RoomID
		c.GuardNum = res.GuardNum
		c.Note = res.Note
		c.Locked = res.Locked
	}

	c.Cache.Store(fkey, nil, time.Second*2)

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

	//Roominitres
	{
		if err, res := biliApi.GetRoomPlayInfo(t.Roomid, 0); err != nil {
			apilog.L(`E: `, err)
			return
		} else {
			//主播uid
			t.UpUid = res.UpUid
			//房间号（完整）
			t.Roomid = res.RoomID
			//直播开始时间
			t.Live_Start_Time = res.LiveStartTime
			//是否在直播
			t.Liveing = res.Liveing

			//未在直播，不获取直播流
			if !t.Liveing {
				t.Live_qn = 0
				t.AcceptQn = t.Qn
				t.Live = t.Live[:0]
				return
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

	//Roominitres
	{
		if err, res := biliApi.GetRoomPlayInfo(t.Roomid, MaxQn); err != nil {
			apilog.L(`E: `, err)
			return
		} else {
			//主播uid
			t.UpUid = res.UpUid
			//房间号（完整）
			t.Roomid = res.RoomID
			//直播开始时间
			t.Live_Start_Time = res.LiveStartTime
			//是否在直播
			t.Liveing = res.Liveing

			//未在直播，不获取直播流
			if !t.Liveing {
				t.Live_qn = 0
				t.AcceptQn = t.Qn
				t.Live = t.Live[:0]
				return
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

func (c *GetFunc) getDanmuInfo() (missKey []string) {
	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if c.LiveBuvidUpdated.Before(time.Now().Add(-time.Hour)) {
		missKey = append(missKey, `LIVE_BUVID`)
	}
	if len(missKey) > 0 {
		return
	}

	//GetDanmuInfo
	if err, res := biliApi.GetDanmuInfo(c.Roomid); err != nil {
		apilog.L(`E: `, err)
		return
	} else {
		c.Token = res.Token
		c.WSURL = res.WSURL
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

	if e, rface := biliApi.GetDanmuMedalAnchorInfo(uid, c.C.Roomid); e != nil {
		apilog.L(`E: `, e)
		return ""
	} else {
		return rface
	}
}

func (t *GetFunc) getPopularAnchorRank() (missKey []string) {
	if t.Uid == 0 {
		missKey = append(missKey, `Uid`)
	}
	if t.UpUid == 0 {
		missKey = append(missKey, `UpUid`)
	}
	if t.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if len(missKey) > 0 {
		return
	}

	apilog := apilog.Base_add(`Get_HotRank`)

	//getHotRank
	if err, note := biliApi.GetPopularAnchorRank(t.Uid, t.UpUid, t.Roomid); err != nil {
		apilog.L(`E: `, err)
	} else {
		t.Note = note
	}

	return
}

func (t *GetFunc) Get_guardNum() (missKey []string) {
	apilog := apilog.Base_add(`Get_guardNum`)

	if t.UpUid == 0 {
		missKey = append(missKey, `UpUid`)
	}
	if t.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if t.LiveBuvidUpdated.Before(time.Now().Add(-time.Hour)) {
		missKey = append(missKey, `LIVE_BUVID`)
	}
	if len(missKey) > 0 {
		return
	}

	//Get_guardNum
	if err, GuardNum := biliApi.GetGuardNum(t.UpUid, t.Roomid); err != nil {
		apilog.L(`E: `, err)
	} else {
		t.GuardNum = GuardNum
	}

	return
}

// func (t *GetFunc) Info(UpUid int) (J.Info, error) {
// 	fkey := `Info`

// 	if v, ok := t.Cache.LoadV(fkey).(cacheItem); ok && v.exceeded.After(time.Now()) {
// 		return (v.data).(J.Info), nil
// 	}

// 	// 超额请求阻塞，超时将取消
// 	apilog := apilog.Base_add(`Info`)
// 	if api_limit.TO() {
// 		return J.Info{}, os.ErrDeadlineExceeded
// 	}

// 	query := fmt.Sprintf("mid=%d&token=&platform=web&web_location=1550101", UpUid)
// 	// wbi
// 	if e, queryE := biliApi.Wbi(query); e != nil {
// 		return J.Info{}, e
// 	} else {
// 		query = queryE
// 	}

// 	// html
// 	{
// 		Cookie := make(map[string]string)
// 		t.Cookie.Range(func(k, v interface{}) bool {
// 			Cookie[k.(string)] = v.(string)
// 			return true
// 		})
// 		req := t.ReqPool.Get()
// 		defer t.ReqPool.Put(req)

// 		if err := req.Reqf(reqf.Rval{
// 			Url:     `https://api.bilibili.com/x/space/wbi/acc/info?` + query,
// 			Proxy:   t.Proxy,
// 			Timeout: 10 * 1000,
// 			Retry:   2,
// 			Header: map[string]string{
// 				`Accept`: "application/json, text/plain, */*",
// 				`Cookie`: reqf.Map_2_Cookies_String(Cookie),
// 			},
// 		}); err != nil {
// 			apilog.L(`E: `, err)
// 			return J.Info{}, err
// 		}

// 		var info J.Info

// 		//Info
// 		if e := json.Unmarshal(req.Respon, &info); e != nil {
// 			apilog.L(`E: `, e)
// 			return J.Info{}, e
// 		}

// 		t.Cache.Store(fkey, cacheItem{
// 			data:     info,
// 			exceeded: time.Now().Add(time.Hour),
// 		})
// 		return info, nil
// 	}
// }

// 调用记录
var boot_Get_cookie funcCtrl.FlashFunc //新的替代旧的

// 扫码登录
func (t *GetFunc) Get_cookie() (missKey []string) {
	apilog := apilog.Base_add(`获取Cookie`)
	//获取其他Cookie
	defer func() {
		if err := biliApi.GetOtherCookies(); err != nil {
			apilog.L(`E: `, err)
		} else if cookies := biliApi.GetCookies(); len(cookies) != 0 {
			if err := save_cookie(cookies); err != nil && !errors.Is(err, ErrNoCookiesSave) {
				apilog.L(`E: `, err)
			}
		}
	}()

	if p.Checkfile().IsExist("cookie.txt") { //读取cookie文件
		if cookieString := string(CookieGet()); cookieString != `` {
			for k, v := range reqf.Cookies_String_2_Map(cookieString) { //cookie存入全局变量syncmap
				t.Cookie.Store(k, v)
			}

			if miss := CookieCheck([]string{
				`bili_jct`,
				`DedeUserID`,
			}); len(miss) == 0 {
				biliApi.SetCookies(reqf.Cookies_String_2_List(cookieString))
				if e, _ := biliApi.GetNav(); e != nil {
					apilog.L(`E: `, e)
				} else {
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
	//获取二维码
	if err, imgUrl, QrcodeKey := biliApi.LoginQrCode(); err != nil {
		apilog.L(`E: `, err)
		return
	} else {
		img_url = imgUrl
		oauth = QrcodeKey
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
				_ = file.New("qr.png", 0, true).CopyToIoWriter(w, pio.CopyConfig{})
			})
			if t.K_v.LoadV(`扫码登录自动打开标签页`).(bool) {
				_ = open.Run(`http://127.0.0.1:` + t.Stream_url.Port() + scanPath)
			}
			apilog.L(`W: `, `或打开链接扫码登录：`+t.Stream_url.String()+scanPath)
		}

		c := qrterminal.Config{
			Level:     qrterminal.L,
			Writer:    os.Stdout,
			BlackChar: `  `,
			WhiteChar: `OO`,
		}
		if white, ok := t.K_v.LoadV(`登陆二维码-白`).(string); ok && len(white) != 0 {
			c.WhiteChar = white
		}
		if black, ok := t.K_v.LoadV(`登陆二维码-黑`).(string); ok && len(black) != 0 {
			c.BlackChar = black
		}
		//show qr code in cmd
		qrterminal.GenerateWithConfig(img_url, c)
		apilog.L(`I: `, `手机扫命令行二维码登录。如不登录，修改配置文件"扫码登录"为false`)
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

		r := t.ReqPool.Get()
		defer t.ReqPool.Put(r)
		for pollC := 10; pollC > 0; pollC-- {
			//3s刷新查看是否通过
			time.Sleep(time.Duration(3) * time.Second)

			//有新实例，退出
			if boot_Get_cookie.NeedExit(id) {
				return
			}

			if err := biliApi.LoginQrPoll(oauth); err != nil {
				apilog.L(`E: `, err)
				return
			} else if cookies := biliApi.GetCookies(); len(cookies) != 0 {
				if err := save_cookie(cookies); err != nil {
					apilog.L(`E: `, err)
					return
				}
				apilog.L(`I: `, `登录,并保存了cookie`)
				return
			}
		}
		apilog.L(`W: `, `扫码超时`)
	}
	return
}

// 获取其他Cookie
// func (c *GetFunc) Get_other_cookie() {
// 	apilog := apilog.Base_add(`获取其他Cookie`)

// 	r := c.ReqPool.Get()
// 	defer c.ReqPool.Put(r)

// 	Cookie := make(map[string]string)
// 	c.Cookie.Range(func(k, v interface{}) bool {
// 		Cookie[k.(string)] = v.(string)
// 		return true
// 	})

// 	if e := r.Reqf(reqf.Rval{
// 		Url: `https://www.bilibili.com/`,
// 		Header: map[string]string{
// 			`Cookie`: reqf.Map_2_Cookies_String(Cookie),
// 		},
// 		Proxy:   c.Proxy,
// 		Timeout: 10 * 1000,
// 		Retry:   2,
// 	}); e != nil {
// 		apilog.L(`E: `, e)
// 		return
// 	}

// 	if e := save_cookie(r.Response.Cookies()); e != nil && !errors.Is(e, ErrNoCookiesSave) {
// 		apilog.L(`E: `, e)
// 	}
// }

// 短信登录
func Get_cookie_by_msg() {
	/*https://passport.bilibili.com/x/passport-login/web/sms/send*/
}

// 牌子字段
// 获取牌子信息
func Get_list_in_room(RoomID, TargetID int) (array []struct {
	TargetID  int
	IsLighted int
	MedalID   int
	RoomID    int
}) {
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

	//getHotRank
	if err, res := biliApi.GetFansMedal(RoomID, TargetID); err != nil {
		apilog.L(`E: `, err)
	} else {
		return res
	}

	return
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

	if err, res := biliApi.GetWearedMedal(); err != nil {
		apilog.L(`E: `, err)
	} else {
		item.Roominfo.RoomID = res.RoomID
		item.TargetID = res.TargetID
		item.TodayIntimacy = res.TodayIntimacy
	}
	return
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
		medal_list := Get_list_in_room(c.Roomid, c.UpUid)
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
	{ //切换牌子
		err := biliApi.SetFansMedal(medal_id)

		if err == nil {
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

	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	//检查是否签到
	if err, status := biliApi.GetWebGetSignInfo(); err != nil {
		apilog.L(`E: `, err)
		return
	} else if status == 1 { //今日已签到
		apilog.L(`T: `, `今日已签到`)
		return
	}

	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	{ //签到
		if err, HadSignDays := biliApi.DoSign(); err != nil {
			apilog.L(`E: `, err)
		} else {
			apilog.L(`I: `, `签到成功!本月已签到`, HadSignDays, `天`)
		}
		return
	}
}

// LIVE_BUVID
func (t *GetFunc) Get_LIVE_BUVID() (missKey []string) {
	apilog := apilog.Base_add(`LIVE_BUVID`)

	//当房间处于特殊活动状态时，将会获取不到，此处使用了若干著名up主房间进行尝试
	roomIdList := []int{
		3, //哔哩哔哩音悦台
		2, //直播姬
		1, //哔哩哔哩直播
	}

	req := t.ReqPool.Get()
	defer t.ReqPool.Put(req)
	for _, roomid := range roomIdList { //获取
		err := biliApi.GetLiveBuvid(roomid)
		if err != nil {
			apilog.L(`E: `, err)
			return
		}
		cookies := biliApi.GetCookies()

		//cookie
		_ = save_cookie(cookies)
		var has bool
		for k := range reqf.Cookies_List_2_Map(cookies) {
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

	t.LiveBuvidUpdated = time.Now()

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

	if err, res := biliApi.GetBagList(c.C.Roomid); err != nil {
		apilog.L(`E: `, err)
		return
	} else {
		apilog.L(`T: `, `成功`)
		return res
	}
}

// 银瓜子2硬币
func (t *GetFunc) Silver_2_coin() (missKey []string) {
	apilog := apilog.Base_add(`银瓜子=>硬币`)

	if t.LiveBuvidUpdated.Before(time.Now().Add(-time.Hour)) {
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
	//验证是否还有机会
	if e, res := biliApi.GetWalletStatus(); e != nil {
		apilog.L(`E: `, e)
		return
	} else {
		if res.Silver2CoinLeft == 0 {
			apilog.L(`I: `, `今天次数已用完`)
			return
		}
		apilog.L(`T: `, `现在有银瓜子`, res.Silver, `个`)
		Silver = res.Silver
	}

	//获取交换规则，验证数量足够
	if e, Silver2CoinPrice := biliApi.GetWalletRule(); e != nil {
		apilog.L(`E: `, e)
		return
	} else if Silver < Silver2CoinPrice {
		apilog.L(`I: `, `当前银瓜子数量不足`)
		return
	}

	//交换
	if e, msg := biliApi.Silver2coin(); e != nil {
		apilog.L(`E: `, e)
		return
	} else {
		apilog.L(`I: `, msg)
		if cookies := biliApi.GetCookies(); len(cookies) != 0 {
			_ = save_cookie(cookies)
		}
	}
	return
}

var ErrNoCookiesSave = errors.New("ErrNoCookiesSave")

func save_cookie(Cookies []*http.Cookie) error {
	if len(Cookies) == 0 {
		return ErrNoCookiesSave
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
	biliApi.SetCookies(Cookies)
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
func GetHisStream() (Uplist []struct {
	Uname      string
	Title      string
	Roomid     int
	LiveStatus int
}) {
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

	if e := biliApi.RoomEntryAction(roomId); e != nil {
		apilog.L(`E: `, e)
		return
	}
}

// 获取在线人数
func (t *GetFunc) getOnlineGoldRank() (misskey []string) {
	apilog := apilog.Base_add(`获取在线人数`)
	if t.UpUid == 0 {
		misskey = append(misskey, `UpUid`)
		return
	}
	if t.Roomid == 0 {
		misskey = append(misskey, `Roomid`)
		return
	}

	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	if e, OnlineNum := biliApi.GetOnlineGoldRank(t.UpUid, t.Roomid); e != nil {
		apilog.L(`E: `, e)
		return
	} else {
		t.OnlineNum = OnlineNum
		apilog.Log_show_control(false).L(`I: `, `在线人数:`, t.OnlineNum)
	}

	return
}

func Feed_list() (Uplist []struct {
	Roomid     int
	Uname      string
	Title      string
	LiveStatus int
}) {
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

	if e, res := biliApi.GetFollowing(); e != nil {
		apilog.L(`E: `, e)
		return
	} else {
		Uplist = res
	}

	return
}

func GetHistory(Roomid_int int) (j []string) {
	apilog := apilog.Base_add(`GetHistory`)

	if e, res := biliApi.GetHisDanmu(Roomid_int); e != nil {
		apilog.L(`E: `, e)
		return
	} else {
		return res
	}
}

func (c *GetFunc) SearchUP(s string) (list []struct {
	Roomid  int
	Uname   string
	Is_live bool
}) {
	apilog := apilog.Base_add(`搜索主播`)
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	if e, res := biliApi.SearchUP(s); e != nil {
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
