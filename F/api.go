package F

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	J "github.com/qydysky/bili_danmu/Json"
	"github.com/skratchdot/open-golang/open"

	p "github.com/qydysky/part"
	funcCtrl "github.com/qydysky/part/funcCtrl"
	g "github.com/qydysky/part/get"
	limit "github.com/qydysky/part/limit"
	reqf "github.com/qydysky/part/reqf"
	sys "github.com/qydysky/part/sys"
	web "github.com/qydysky/part/web"

	"github.com/mdp/qrterminal/v3"
	qr "github.com/skip2/go-qrcode"
)

var apilog = c.C.Log.Base(`api`)
var api_limit = limit.New(1, 2000, 30000) //频率限制1次/2s，最大等待时间30s

type GetFunc struct {
	*c.Common
}

func Get(c *c.Common) *GetFunc {
	return &GetFunc{c}
}

func (c *GetFunc) Get(key string) {
	apilog := apilog.Base_add(`Get`)

	if api_limit.TO() {
		return
	} //超额请求阻塞，超时将取消

	var (
		api_can_get = map[string][]func() []string{
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

			`Silver_2_coin`: { //银瓜子2硬币
				c.Silver_2_coin,
			},
			`CheckSwitch_FansMedal`: { //切换粉丝牌
				c.CheckSwitch_FansMedal,
			},
		}
		check = map[string]func() bool{
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
				return c.LIVE_BUVID
			},
			`Silver_2_coin`: func() bool { //银瓜子2硬币
				return true
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
			apilog.L(`T: `, `Get`, key)
			missKey := fItem()
			if len(missKey) > 0 {
				apilog.L(`T: `, `missKey when get`, key, missKey)
				for _, misskeyitem := range missKey {
					if checkf, ok := check[misskeyitem]; ok && checkf() {
						continue
					}
					if misskeyitem == key {
						apilog.L(`W: `, `missKey equrt key`, key, missKey)
						continue
					}
					c.Get(misskeyitem)
				}
				missKey := fItem()
				if len(missKey) > 0 {
					apilog.L(`W: `, `missKey when get`, key, missKey)
					continue
				}
			}
			if checkf, ok := check[key]; ok && checkf() {
				break
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

func (c *GetFunc) Html() (missKey []string) {
	apilog := apilog.Base_add(`html`)

	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
		return
	}

	Roomid := strconv.Itoa(c.Roomid)

	//html
	{
		r := g.Get(reqf.Rval{
			Url:   "https://live.bilibili.com/" + Roomid,
			Proxy: c.Proxy,
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
				c.UpUid = j.RoomInitRes.Data.UID
				//房间号（完整）
				if j.RoomInitRes.Data.RoomID != 0 {
					c.Roomid = j.RoomInitRes.Data.RoomID
				}
				//直播开始时间
				c.Live_Start_Time = time.Unix(int64(j.RoomInitRes.Data.LiveTime), 0)
				//是否在直播
				c.Liveing = j.RoomInitRes.Data.LiveStatus == 1

				//未在直播，不获取直播流
				if !c.Liveing {
					c.Live_qn = 0
					c.AcceptQn = c.Qn
					c.Live = []string{}
					return
				}

				//当前直播流
				{
					type Stream_name struct {
						Protocol_name string
						Format_name   string
						Codec_name    string
					}
					var name_map = map[string]Stream_name{
						`flv`: {
							Protocol_name: "http_stream",
							Format_name:   "flv",
							Codec_name:    "avc",
						},
						`hls`: {
							Protocol_name: "http_hls",
							Format_name:   "fmp4",
							Codec_name:    "avc",
						},
					}

					want_type := name_map[`hls`]
					if v, ok := c.K_v.LoadV(`直播流类型`).(string); ok {
						if v, ok := name_map[v]; ok {
							want_type = v
						} else {
							apilog.L(`I: `, `未找到`, v, `,默认hls`)
						}
					} else {
						apilog.L(`T: `, `默认flv`)
					}

					for _, v := range j.RoomInitRes.Data.PlayurlInfo.Playurl.Stream {
						if v.ProtocolName != want_type.Protocol_name {
							continue
						}

						for _, v := range v.Format {
							if v.FormatName != want_type.Format_name {
								continue
							}

							for _, v := range v.Codec {
								if v.CodecName != want_type.Codec_name {
									continue
								}

								//当前直播流质量
								c.Live_qn = v.CurrentQn
								if c.Live_want_qn == 0 {
									c.Live_want_qn = v.CurrentQn
								}
								//允许的清晰度
								{
									var tmp = make(map[int]string)
									for _, v := range v.AcceptQn {
										if s, ok := c.Qn[v]; ok {
											tmp[v] = s
										}
									}
									c.AcceptQn = tmp
								}
								//直播流链接
								c.Live = []string{}
								for _, v1 := range v.URLInfo {
									c.Live = append(c.Live, v1.Host+v.BaseURL+v1.Extra)
								}
							}
						}
					}
				}
			}

			//Roominfores
			{
				//直播间标题
				c.Title = j.RoomInfoRes.Data.RoomInfo.Title
				//主播名
				c.Uname = j.RoomInfoRes.Data.AnchorInfo.BaseInfo.Uname
				//分区
				c.ParentAreaID = j.RoomInfoRes.Data.RoomInfo.ParentAreaID
				//子分区
				c.AreaID = j.RoomInfoRes.Data.RoomInfo.AreaID
				//舰长数
				c.GuardNum = j.RoomInfoRes.Data.GuardInfo.Count
				//分区排行
				c.Note = j.RoomInfoRes.Data.PopularRankInfo.RankName + " "
				if rank := j.RoomInfoRes.Data.PopularRankInfo.Rank; rank > 50 || rank == 0 {
					c.Note += "100+"
				} else {
					c.Note += strconv.Itoa(rank)
				}
				//直播间是否被封禁
				if j.RoomInfoRes.Data.RoomInfo.LockStatus == 1 {
					apilog.L(`W: `, "直播间封禁中")
					c.Locked = true
					return
				}
			}
		}
	}
	return
}

func (c *GetFunc) missRoomId() (missKey []string) {
	apilog.Base_add(`missRoomId`).L(`E: `, `missRoomId`)
	return
}

func (c *GetFunc) getInfoByRoom() (missKey []string) {
	apilog := apilog.Base_add(`getInfoByRoom`)

	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
		return
	}

	Roomid := strconv.Itoa(c.Roomid)

	{ //使用其他api
		reqi := c.Common.ReqPool.Get()
		defer c.Common.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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
	return
}

func (c *GetFunc) getRoomPlayInfo() (missKey []string) {
	apilog := apilog.Base_add(`getRoomPlayInfo`)

	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if !c.LIVE_BUVID {
		missKey = append(missKey, `LIVE_BUVID`)
	}
	if len(missKey) > 0 {
		return
	}

	Roomid := strconv.Itoa(c.Roomid)

	//Roominitres
	{
		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		reqi := c.Common.ReqPool.Get()
		defer c.Common.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
		if err := req.Reqf(reqf.Rval{
			Url: "https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo?no_playurl=0&mask=1&qn=0&platform=web&protocol=0,1&format=0,2&codec=0,1&room_id=" + Roomid,
			Header: map[string]string{
				`Referer`: "https://live.bilibili.com/" + Roomid,
				`Cookie`:  reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.Proxy,
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
		c.UpUid = j.Data.UID
		//房间号（完整）
		if j.Data.RoomID != 0 {
			c.Roomid = j.Data.RoomID
		}
		//直播开始时间
		c.Live_Start_Time = time.Unix(int64(j.Data.LiveTime), 0)
		//是否在直播
		c.Liveing = j.Data.LiveStatus == 1

		//未在直播，不获取直播流
		if !c.Liveing {
			c.Live_qn = 0
			c.AcceptQn = c.Qn
			c.Live = []string{}
			return
		}

		//当前直播流
		{
			type Stream_name struct {
				Protocol_name string
				Format_name   string
				Codec_name    string
			}

			//所有支持的格式
			var name_map = map[string]Stream_name{
				`flv`: {
					Protocol_name: "http_stream",
					Format_name:   "flv",
					Codec_name:    "avc",
				},
				`hls`: {
					Protocol_name: "http_hls",
					Format_name:   "fmp4",
					Codec_name:    "avc",
				},
			}

			// 默认使用hls
			want_type := name_map[`hls`]

			//从配置文件选取格式
			if v, ok := c.K_v.LoadV(`直播流类型`).(string); ok {
				if v, ok := name_map[v]; ok {
					want_type = v
				} else {
					apilog.L(`I: `, `未找到`, v, `,默认hls`)
				}
			} else {
				apilog.L(`T: `, `默认hls`)
			}

			no_found_type := true
			for {
				//返回的所有支持的格式
				for _, v := range j.Data.PlayurlInfo.Playurl.Stream {
					//选取配置中的格式
					if v.ProtocolName != want_type.Protocol_name {
						continue
					}

					for _, v := range v.Format {
						//选取配置中的格式
						if v.FormatName != want_type.Format_name {
							continue
						}

						no_found_type = false

						for _, v := range v.Codec {
							//选取配置中的格式
							if v.CodecName != want_type.Codec_name {
								continue
							}

							//当前直播流质量
							c.Live_qn = v.CurrentQn
							if c.Live_want_qn == 0 {
								c.Live_want_qn = v.CurrentQn
							}
							//允许的清晰度
							{
								var tmp = make(map[int]string)
								for _, v := range v.AcceptQn {
									if s, ok := c.Qn[v]; ok {
										tmp[v] = s
									}
								}
								c.AcceptQn = tmp
							}
							//直播流链接
							c.Live = []string{}
							for _, v1 := range v.URLInfo {
								c.Live = append(c.Live, v1.Host+v.BaseURL+v1.Extra)
							}

							//找到配置格式，跳出
							break
						}
					}
				}
				if no_found_type {
					if want_type.Protocol_name == "http_stream" {
						apilog.L(`I: `, `不支持flv，使用hls`)
						want_type = name_map[`hls`]
					} else {
						apilog.L(`I: `, `不支持hls，使用flv`)
						want_type = name_map[`flv`]
					}
					no_found_type = false
				} else {
					break
				}
			}
		}
	}
	return
}

func (c *GetFunc) getRoomPlayInfoByQn() (missKey []string) {
	apilog := apilog.Base_add(`getRoomPlayInfoByQn`)

	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if !c.LIVE_BUVID {
		missKey = append(missKey, `LIVE_BUVID`)
	}
	if len(missKey) > 0 {
		return
	}

	{
		AcceptQn := []int{}
		for k := range c.AcceptQn {
			if k <= c.Live_want_qn {
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
		c.Live_qn = MaxQn
	}

	Roomid := strconv.Itoa(c.Roomid)

	//Roominitres
	{
		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		reqi := c.Common.ReqPool.Get()
		defer c.Common.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
		if err := req.Reqf(reqf.Rval{
			Url: "https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo?no_playurl=0&mask=1&qn=" + strconv.Itoa(c.Live_qn) + "&platform=web&protocol=0,1&format=0,2&codec=0,1&room_id=" + Roomid,
			Header: map[string]string{
				`Referer`: "https://live.bilibili.com/" + Roomid,
				`Cookie`:  reqf.Map_2_Cookies_String(Cookie),
			},
			Proxy:   c.Proxy,
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
		c.UpUid = j.Data.UID
		//房间号（完整）
		if j.Data.RoomID != 0 {
			c.Roomid = j.Data.RoomID
		}
		//直播开始时间
		c.Live_Start_Time = time.Unix(int64(j.Data.LiveTime), 0)
		//是否在直播
		c.Liveing = j.Data.LiveStatus == 1

		//未在直播，不获取直播流
		if !c.Liveing {
			c.Live_qn = 0
			c.AcceptQn = c.Qn
			c.Live = []string{}
			return
		}

		//当前直播流
		{
			type Stream_name struct {
				Protocol_name string
				Format_name   string
				Codec_name    string
			}
			var name_map = map[string]Stream_name{
				`flv`: {
					Protocol_name: "http_stream",
					Format_name:   "flv",
					Codec_name:    "avc",
				},
				`hls`: {
					Protocol_name: "http_hls",
					Format_name:   "fmp4",
					Codec_name:    "avc",
				},
			}

			want_type := name_map[`hls`]
			if v, ok := c.K_v.LoadV(`直播流类型`).(string); ok {
				if v, ok := name_map[v]; ok {
					want_type = v
				} else {
					apilog.L(`I: `, `未找到`, v, `,默认hls`)
				}
			} else {
				apilog.L(`T: `, `默认hls`)
			}

			no_found_type := true
			for {
				for _, v := range j.Data.PlayurlInfo.Playurl.Stream {
					if v.ProtocolName != want_type.Protocol_name {
						continue
					}

					for _, v := range v.Format {
						if v.FormatName != want_type.Format_name {
							continue
						}

						for _, v := range v.Codec {
							if v.CodecName != want_type.Codec_name {
								continue
							}

							no_found_type = false

							//当前直播流质量
							c.Live_qn = v.CurrentQn
							if c.Live_want_qn == 0 {
								c.Live_want_qn = v.CurrentQn
							}
							//允许的清晰度
							{
								var tmp = make(map[int]string)
								for _, v := range v.AcceptQn {
									if s, ok := c.Qn[v]; ok {
										tmp[v] = s
									}
								}
								c.AcceptQn = tmp
							}
							//直播流链接
							c.Live = []string{}
							for _, v1 := range v.URLInfo {
								c.Live = append(c.Live, v1.Host+v.BaseURL+v1.Extra)
							}
						}
					}
				}
				if no_found_type {
					if want_type.Protocol_name == "http_stream" {
						apilog.L(`I: `, `不支持flv，使用hls`)
						want_type = name_map[`hls`]
					} else {
						apilog.L(`I: `, `不支持hls，使用flv`)
						want_type = name_map[`flv`]
					}
					no_found_type = false
				} else {
					break
				}
			}
			if s, ok := c.Qn[c.Live_qn]; !ok {
				apilog.L(`W: `, `未知清晰度`, c.Live_qn)
			} else {
				apilog.L(`I: `, s)
			}
		}
	}
	return
}

func (c *GetFunc) getDanmuInfo() (missKey []string) {
	apilog := apilog.Base_add(`getDanmuInfo`)

	if c.Roomid == 0 {
		missKey = append(missKey, `Roomid`)
	}
	if !c.LIVE_BUVID {
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

		reqi := c.Common.ReqPool.Get()
		defer c.Common.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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
		for _, v := range j.Data.HostList {
			if v.WssPort != 443 {
				c.WSURL = append(c.WSURL, "wss://"+v.Host+":"+strconv.Itoa(v.WssPort)+"/sub")
			}
			c.WSURL = append(c.WSURL, "wss://"+v.Host+"/sub")
		}
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

	reqi := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(reqi)
	req := reqi.Item.(*reqf.Req)
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

		reqi := c.ReqPool.Get()
		defer c.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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
	if !c.LIVE_BUVID {
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

		reqi := c.ReqPool.Get()
		defer c.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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

func Info(UpUid int) (info J.Info) {
	apilog := apilog.Base_add(`Info`)
	if api_limit.TO() {
		return
	} //超额请求阻塞，超时将取消

	//html
	{
		Cookie := make(map[string]string)
		c.C.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})
		reqi := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
		if err := req.Reqf(reqf.Rval{
			Url:     `https://api.bilibili.com/x/space/acc/info?mid=` + strconv.Itoa(UpUid) + `&token=&platform=web&jsonp=jsonp`,
			Proxy:   c.C.Proxy,
			Timeout: 10 * 1000,
			Retry:   2,
			Header: map[string]string{
				`Cookie`: reqf.Map_2_Cookies_String(Cookie),
			},
		}); err != nil {
			apilog.L(`E: `, err)
			return
		}

		//Info
		{
			if e := json.Unmarshal(req.Respon, &info); e != nil {
				apilog.L(`E: `, e)
				return
			} else if info.Code != 0 {
				apilog.L(`E: `, info.Message)
				return
			}
		}
	}
	return
}

// 调用记录
var boot_Get_cookie funcCtrl.FlashFunc //新的替代旧的

// 扫码登录
func (c *GetFunc) Get_cookie() (missKey []string) {
	if v, ok := c.K_v.LoadV(`扫码登录`).(bool); !ok || !v {
		return
	}

	apilog := apilog.Base_add(`获取Cookie`)

	if p.Checkfile().IsExist("cookie.txt") { //读取cookie文件
		if cookieString := string(CookieGet()); cookieString != `` {
			for k, v := range reqf.Cookies_String_2_Map(cookieString) { //cookie存入全局变量syncmap
				c.Cookie.Store(k, v)
			}
			if miss := CookieCheck([]string{
				`bili_jct`,
				`DedeUserID`,
			}); len(miss) == 0 {
				return
			}
		}
	}

	//获取id
	id := boot_Get_cookie.Flash()
	defer boot_Get_cookie.UnFlash()

	var img_url string
	var oauth string
	{ //获取二维码
		reqi := c.ReqPool.Get()
		defer c.ReqPool.Put(reqi)
		r := reqi.Item.(*reqf.Req)
		if e := r.Reqf(reqf.Rval{
			Url:     `https://passport.bilibili.com/qrcode/getLoginUrl`,
			Proxy:   c.Proxy,
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

	scanPort := int(c.K_v.LoadV("扫码登录端口").(float64))
	if scanPort <= 0 {
		scanPort = sys.Sys().GetFreePort()
	}
	var server = &http.Server{
		Addr: "0.0.0.0:" + strconv.Itoa(scanPort),
	}

	{ //生成二维码
		qr.WriteFile(img_url, qr.Medium, 256, `qr.png`)
		if !p.Checkfile().IsExist(`qr.png`) {
			apilog.L(`E: `, `qr error`)
			return
		}
		//启动web
		s := web.New(server)
		s.Handle(map[string]func(http.ResponseWriter, *http.Request){
			`/`: func(w http.ResponseWriter, r *http.Request) {
				var path string = r.URL.Path[1:]
				if path == `` {
					path = `index.html`
				}
				http.ServeFile(w, r, path)
			},
			`/exit`: func(_ http.ResponseWriter, _ *http.Request) {
				s.Server.Shutdown(context.Background())
			},
		})
		defer server.Shutdown(context.Background())

		if c.K_v.LoadV(`扫码登录自动打开标签页`).(bool) {
			open.Run(`http://` + server.Addr + `/qr.png`)
		}
		apilog.Block(1000)
		//show qr code in cmd
		qrterminal.GenerateWithConfig(img_url, qrterminal.Config{
			Level:     qrterminal.L,
			Writer:    os.Stdout,
			BlackChar: `  `,
			WhiteChar: `OO`,
		})
		apilog.L(`W: `, `手机扫命令行二维码登录`)
		apilog.L(`W: `, `或打开链接扫码登录： http://`+s.Server.Addr+`/qr.png`)
		sys.Sys().Timeoutf(1)
	}

	//有新实例，退出
	if boot_Get_cookie.NeedExit(id) {
		return
	}

	var cookie string
	{ //循环查看是否通过
		Cookie := make(map[string]string)
		c.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		for {
			//3s刷新查看是否通过
			sys.Sys().Timeoutf(3)

			//有新实例，退出
			if boot_Get_cookie.NeedExit(id) {
				return
			}

			reqi := c.ReqPool.Get()
			defer c.ReqPool.Put(reqi)
			r := reqi.Item.(*reqf.Req)
			if e := r.Reqf(reqf.Rval{
				Url:     `https://passport.bilibili.com/qrcode/getLoginInfo`,
				PostStr: `oauthKey=` + oauth,
				Header: map[string]string{
					`Content-Type`: `application/x-www-form-urlencoded; charset=UTF-8`,
					`Referer`:      `https://passport.bilibili.com/login`,
					`Cookie`:       reqf.Map_2_Cookies_String(Cookie),
				},
				Proxy:   c.Proxy,
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
				if v := r.Response.Cookies(); len(v) == 0 {
					apilog.L(`E: `, `getLoginInfo cookies len == 0`)
					return
				} else {
					cookie = reqf.Map_2_Cookies_String(reqf.Cookies_List_2_Map(v)) //cookie to string
				}
				if cookie == `` {
					apilog.L(`E: `, `getLoginInfo cookies ""`)
					return
				} else {
					break
				}
			}
		}
		if len(cookie) == 0 {
			return
		}
	}

	//有新实例，退出
	if boot_Get_cookie.NeedExit(id) {
		return
	}

	{ //写入cookie.txt
		for k, v := range reqf.Cookies_String_2_Map(cookie) {
			c.Cookie.Store(k, v)
		}
		//生成cookieString
		cookieString := ``
		{
			c.Cookie.Range(func(k, v interface{}) bool {
				cookieString += k.(string) + `=` + v.(string) + `; `
				return true
			})
			t := []rune(cookieString)
			cookieString = string(t[:len(t)-2])
		}

		CookieSet([]byte(cookieString))
	}

	//有新实例，退出
	if boot_Get_cookie.NeedExit(id) {
		return
	}

	{ //清理
		if p.Checkfile().IsExist(`qr.png`) {
			os.RemoveAll(`qr.png`)
			return
		}
	}
	return
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
			reqi := c.C.ReqPool.Get()
			defer c.C.ReqPool.Put(reqi)
			r := reqi.Item.(*reqf.Req)
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
		reqi := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(reqi)
		r := reqi.Item.(*reqf.Req)
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
				json.Unmarshal(data, &item)
			}
		}

		return
	}

}

func (c *GetFunc) CheckSwitch_FansMedal() (missKey []string) {

	if !c.LIVE_BUVID {
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
		reqi := c.ReqPool.Get()
		defer c.ReqPool.Put(reqi)
		r := reqi.Item.(*reqf.Req)
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
		res := string(r.Respon)
		if v, ok := p.Json().GetValFromS(res, "code").(float64); ok && v == 0 {
			apilog.L(`I: `, `自动切换粉丝牌 id:`, medal_id)
			c.Wearing_FansMedal = medal_id //更新佩戴信息
			return
		}
		if v, ok := p.Json().GetValFromS(res, "message").(string); ok {
			apilog.L(`E: `, `Get_FansMedal wear message`, v)
		} else {
			apilog.L(`E: `, `Get_FansMedal wear message nil`)
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

		reqi := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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

		reqi := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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
	apilog := apilog.Base_add(`LIVE_BUVID`).L(`T: `, `获取`)

	if live_buvid, ok := c.Cookie.LoadV(`LIVE_BUVID`).(string); ok && live_buvid != `` {
		apilog.L(`T: `, `存在`)
		c.LIVE_BUVID = true
		return
	}

	//当房间处于特殊活动状态时，将会获取不到，此处使用了若干著名up主房间进行尝试
	roomIdList := []string{
		"3", //哔哩哔哩音悦台
		"2", //直播姬
		"1", //哔哩哔哩直播
	}

	for _, roomid := range roomIdList { //获取
		reqi := c.ReqPool.Get()
		defer c.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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
		var has bool
		for k, v := range reqf.Cookies_List_2_Map(req.Response.Cookies()) {
			c.Cookie.Store(k, v)
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

	Cookie := make(map[string]string)
	c.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})

	CookieSet([]byte(reqf.Map_2_Cookies_String(Cookie)))

	c.LIVE_BUVID = true

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

	reqi := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(reqi)
	req := reqi.Item.(*reqf.Req)
	if err := req.Reqf(reqf.Rval{
		Url: `https://api.live.bilibili.com/xlive/web-room/v1/gift/bag_list?t=` + strconv.Itoa(int(sys.Sys().GetMTime())) + `&room_id=` + strconv.Itoa(c.C.Roomid),
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

	if !c.LIVE_BUVID {
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

		reqi := c.ReqPool.Get()
		defer c.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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

		reqi := c.ReqPool.Get()
		defer c.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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

		reqi := c.ReqPool.Get()
		defer c.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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

		save_cookie(req.Response.Cookies())

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

func save_cookie(Cookies []*http.Cookie) {
	for k, v := range reqf.Cookies_List_2_Map(Cookies) {
		c.C.Cookie.Store(k, v)
	}

	Cookie := make(map[string]string)
	c.C.Cookie.Range(func(k, v interface{}) bool {
		Cookie[k.(string)] = v.(string)
		return true
	})
	CookieSet([]byte(reqf.Map_2_Cookies_String(Cookie)))
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

	reqi := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(reqi)
	req := reqi.Item.(*reqf.Req)
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

	reqi := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(reqi)

	req := reqi.Item.(*reqf.Req)
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

	reqi := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(reqi)
	req := reqi.Item.(*reqf.Req)
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
		reqi := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)
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

func SearchUP(s string) (list []searchresult) {
	apilog := apilog.Base_add(`搜索主播`)
	if api_limit.TO() {
		apilog.L(`E: `, `超时！`)
		return
	} //超额请求阻塞，超时将取消

	{ //使用其他api
		reqi := c.C.ReqPool.Get()
		defer c.C.ReqPool.Put(reqi)
		req := reqi.Item.(*reqf.Req)

		Cookie := make(map[string]string)
		c.C.Cookie.Range(func(k, v interface{}) bool {
			Cookie[k.(string)] = v.(string)
			return true
		})

		if err := req.Reqf(reqf.Rval{
			Url:   "https://api.bilibili.com/x/web-interface/wbi/search/type?page=1&page_size=10&order=online&platform=pc&search_type=live_user&keyword=" + url.PathEscape(s),
			Proxy: c.C.Proxy,
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

	reqi := c.C.ReqPool.Get()
	defer c.C.ReqPool.Put(reqi)
	req := reqi.Item.(*reqf.Req)
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
