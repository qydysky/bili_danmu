package F

import (
	"net/http"
	"time"

	pool "github.com/qydysky/part/pool"
	reqf "github.com/qydysky/part/reqf"
)

type biliApiInter interface {
	SetReqPool(pool *pool.Buf[reqf.Req])
	SetProxy(proxy string)
	SetCookies(cookies []*http.Cookie)

	LoginQrCode() (err error, imgUrl string, QrcodeKey string)
	LoginQrPoll(QrcodeKey string) (err error, cookies []*http.Cookie)
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
	GetRoomPlayInfo(Roomid int) (err error, res struct {
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
}
