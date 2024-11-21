package F

import (
	"net/http"
	"time"

	pool "github.com/qydysky/part/pool"
	reqf "github.com/qydysky/part/reqf"
)

type BiliApiInter interface {
	SetReqPool(pool *pool.Buf[reqf.Req])
	SetProxy(proxy string)
	SetLocation(secOfTimeZone int) // east positive
	SetCookies(cookies []*http.Cookie)
	GetCookies() (cookies []*http.Cookie)
	GetCookie(name string) (error, string)

	LikeReport(hitCount, uid, roomid, upUid int) (err error)
	LoginQrCode() (err error, imgUrl string, QrcodeKey string)
	LoginQrPoll(QrcodeKey string) (err error, code int)
	GetOtherCookies() (err error)
	GetLiveBuvid(Roomid int) (err error)
	GetRoomBaseInfo(Roomid int) (err error, res struct {
		UpUid         int
		Uname         string
		ParentAreaID  int
		AreaID        int
		Title         string
		LiveStartTime time.Time
		Liveing       bool
		RoomID        int
	})
	GetInfoByRoom(Roomid int) (err error, res struct {
		UpUid         int
		Uname         string
		ParentAreaID  int
		AreaID        int
		Title         string
		LiveStartTime time.Time
		Liveing       bool
		RoomID        int
		GuardNum      int
		Note          string
		Locked        bool
	})
	GetRoomPlayInfo(Roomid int, Qn int) (err error, res struct {
		UpUid         int
		RoomID        int
		LiveStartTime time.Time
		Liveing       bool
		Streams       []struct {
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
		}
	})
	GetDanmuInfo(Roomid int) (err error, res struct {
		Token string
		WSURL []string
	})
	GetDanmuMedalAnchorInfo(uid string, Roomid int) (err error, rface string)
	GetPopularAnchorRank(uid, upUid, roomid int) (err error, note string)
	GetGuardNum(upUid, roomid int) (err error, GuardNum int)
	GetNav() (err error, res struct {
		IsLogin bool
		WbiImg  struct {
			ImgURL string
			SubURL string
		}
	})
	Wbi(query string, WbiImg struct {
		ImgURL string
		SubURL string
	}) (err error, queryEnc string)
	GetWearedMedal(uid, upUid int) (err error, res struct {
		TodayIntimacy int
		RoomID        int
		TargetID      int
	})
	GetFansMedal(RoomID, TargetID int) (err error, res []struct {
		TodayFeed    int
		TargetID     int
		IsLighted    int
		MedalID      int
		RoomID       int
		LivingStatus int
	})
	SetFansMedal(medalId int) (err error)
	GetWebGetSignInfo() (err error, Status int)
	DoSign() (err error, HadSignDays int)
	GetBagList(Roomid int) (err error, res []struct {
		Bag_id    int
		Gift_id   int
		Gift_name string
		Gift_num  int
		Expire_at int
	})
	GetWalletStatus() (err error, res struct {
		Silver          int
		Silver2CoinLeft int
	})
	GetWalletRule() (err error, Silver2CoinPrice int)
	Silver2coin() (err error, Message string)
	GetHisStream() (err error, res []struct {
		Uname      string
		Title      string
		Roomid     int
		LiveStatus int
	})
	RoomEntryAction(Roomid int) (err error)
	GetOnlineGoldRank(upUid, roomid int) (err error, OnlineNum int)
	GetFollowing() (err error, res []struct {
		Roomid     int
		Uname      string
		Title      string
		LiveStatus int
	})
	IsConnected() (err error)
	GetHisDanmu(Roomid int) (err error, res []string)
	SearchUP(s string) (err error, res []struct {
		Roomid  int
		Uname   string
		Is_live bool
	})
	LiveHtml(Roomid int) (err error, res struct {
		RoomInitRes struct {
			Code    int
			Message string
			TTL     int
			Data    struct {
				RoomID      int
				UID         int
				LiveStatus  int
				LiveTime    int
				PlayurlInfo struct {
					ConfJSON string
					Playurl  struct {
						Stream []struct {
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
						}
					}
				}
			}
		}
		RoomInfoRes struct {
			Code    int
			Message string
			TTL     int
			Data    struct {
				RoomInfo struct {
					Title        string
					LockStatus   int
					AreaID       int
					ParentAreaID int
				}
				AnchorInfo struct {
					BaseInfo struct {
						Uname string
					}
				}
				PopularRankInfo struct {
					Rank     int
					RankName string
				}
				GuardInfo struct {
					Count int
				}
			}
		}
	})
}
