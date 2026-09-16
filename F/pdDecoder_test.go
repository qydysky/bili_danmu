package F

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"testing"

	unsafe "github.com/qydysky/part/unsafe"
)

func Test_SendGiftV2(t *testing.T) {
	raw := []byte(`{"cmd":"SEND_GIFT_V2","danmu":{"area":1},"data":{"dmscore":140,"pb":"CM3igxASEuS7u+WlueiKseW8gOiKseiQvRpKaHR0cHM6Ly9pMi5oZHNsYi5jb20vYmZzL2ZhY2UvYWYzZDBhNzhiZWQ4M2QwMDEyMWFlOWY5YmFhNzk2NjA1MzdjZTVjYS5qcGdCAFKRBgi88wESD+eyieS4neWboueBr+eJjBgBIAEoZDBkOGRCBGdvbGRKEzQ4MTc2NTk3MDQxMjM3ODk4MjRQ7Kiq1QZYAWI4YmF0Y2g6Z2lmdDpjb21ib19pZDozMzYxNjIwNTo2NzE0MTozMTE2NDoxNzg5NTY0MDEyLjI5MjloCnBkeAWFAQAAgD+IAQGSAQbmipXlloLAAby1y6ED6gEYChLlgrLmhaLnmoTlsI/ogonljIUQxYwEigLNAgjFjAQSxgIKEuWCsuaFoueahOWwj+iCieWMhRJKaHR0cHM6Ly9pMS5oZHNsYi5jb20vYmZzL2dhcmIvNGQ5MzhkZjY5OWNiOWUzNjdjODcyNWYwOWFmYTI0OWU5NmE1YzUwNS5wbmcyYAoS5YKy5oWi55qE5bCP6IKJ5YyFEkpodHRwczovL2kxLmhkc2xiLmNvbS9iZnMvZ2FyYi80ZDkzOGRmNjk5Y2I5ZTM2N2M4NzI1ZjA5YWZhMjQ5ZTk2YTVjNTA1LnBuZzqBAQgBEhpiaWxpYmlsaSDnn6XlkI3muLjmiI9VUOS4uxph5Luj6KGo5L2c44CK44CQ55Sf5YyW5Y2x5py6N0RMQ+OAkeS8iuajruW/hemhu+atu+iuqeS8iuajruWTreedgOaJvuWmiOWmiOeahOmAmuWFs+aUu+eVpeinhumikeOAi5ICBwiM4NcCEAGaAuUBCkpodHRwczovL3MxLmhkc2xiLmNvbS9iZnMvbGl2ZS9lMDUxZGZkNDU1NzY3OGY4ZWRjYWM0OTkzZWQwMGEwOTM1Y2JkOWNjLnBuZxJLaHR0cHM6Ly9pMC5oZHNsYi5jb20vYmZzL2xpdmUvMzJiNzk5MTIwZTE2MTRmYTYyNzViNmQxNWRhN2E1MmIyMWRkMDE5ZC53ZWJwKkpodHRwczovL2kwLmhkc2xiLmNvbS9iZnMvbGl2ZS84MTZmOGI3YWEyMTMyODg4ZmNlOTI4Y2RmYjE3YjljZjIxY2MwODIzLmdpZqoCFAgBEgcI9bfwAhACEgcIjODXAhABWAFqAggIerYCCM3igxASzwEKEuS7u+WlueiKseW8gOiKseiQvRJKaHR0cHM6Ly9pMi5oZHNsYi5jb20vYmZzL2ZhY2UvYWYzZDBhNzhiZWQ4M2QwMDEyMWFlOWY5YmFhNzk2NjA1MzdjZTVjYS5qcGcyYAoS5Lu75aW56Iqx5byA6Iqx6JC9EkpodHRwczovL2kyLmhkc2xiLmNvbS9iZnMvZmFjZS9hZjNkMGE3OGJlZDgzZDAwMTIxYWU5ZjliYWE3OTY2MDUzN2NlNWNhLmpwZzoLIP///////////wEaXQoEQ+mFsRAdGMCBgwYgwIGDBijAgYMGMNWQtAFQxYwEYL2fAXoJIzkxOTI5OENDggEJIzkxOTI5OENDigEJIzkxOTI5OENDkgEHI0ZGRkZGRpoBCSM5MTkyOThFNg=="}}`)

	var s = struct {
		Cmd   string `json:"cmd"`
		Danmu struct {
			Area int `json:"area"`
		} `json:"danmu"`
		Data struct {
			Dmscore int    `json:"dmscore"`
			Pb      string `json:"pb"`
			PbS     struct {
				Uid   int    `pd:"1"`
				Uname string `pd:"2"`
				// Face       string `pd:"3"`
				// NameColor  string `pd:"4"`
				// GuardLevel int    `pd:"5"`
				// SvgaBlock  int    `pd:"6"`
				// SendMaster struct {
				// } `pd:"7"`
				// MedalInfo struct {
				// 	TargetId         int    `pd:"1"`
				// 	Special          string `pd:"2"`
				// 	AnchorUname      string `pd:"3"`
				// 	AnchorRoomid     int    `pd:"4"`
				// 	MedalLevel       int    `pd:"5"`
				// 	MedalName        string `pd:"6"`
				// 	MedalColor       int    `pd:"7"`
				// 	MedalColorStart  int    `pd:"8"`
				// 	MedalColorEnd    int    `pd:"9"`
				// 	MedalColorBorder int    `pd:"10"`
				// 	IsLighted        int    `pd:"11"`
				// 	GuardLevel       int    `pd:"12"`
				// } `pd:"8"`
				// BlindGift struct {
				// } `pd:"9"`
				GiftList []struct {
					// GiftId             int     `pd:"1"`
					GiftName string `pd:"2"`
					Num      int64  `pd:"3"`
					// Demarcation        int64   `pd:"4"`
					// Price              int64   `pd:"5"`
					// DiscountPrice      int64   `pd:"6"`
					TotalCoin int64  `pd:"7"`
					CoinType  string `pd:"8"`
					// Tid                string  `pd:"9"`
					// Timestamp          int64   `pd:"10"`
					// SuperBatchGiftNum  int64   `pd:"11"`
					// BatchComboId       string  `pd:"12"`
					// ComboResourcesId   int64   `pd:"13"`
					// ComboTotalCoin     int64   `pd:"14"`
					// ComboStayTime      int64   `pd:"15"`
					// Magnification      float32 `pd:"16"`
					// ShowBatchComboSend bool    `pd:"17"`
					// Action             string  `pd:"18"`
					// EffectBlock        int64   `pd:"19"`
					// IsSpecialBatch     int64   `pd:"20"`
					// FloatScResourceId  int64   `pd:"21"`
					// TagImage           string  `pd:"22"`
					// CritProb           int64   `pd:"23"`
					// Rcost              int64   `pd:"24"`
					// Test               int64   `pd:"25"`
					// FaceEffectType     int64   `pd:"26"`
					// FaceEffectId       int64   `pd:"27"`
					// IsNaming           bool    `pd:"28"`
					ReceiveUserInfo struct {
						Uname string `pd:"1"`
						Uid   int64  `pd:"2"`
					} `pd:"29"`
					// IsJoinReceiver bool `pd:"30"`
					// BagGift        struct{}   `pd:"31"`
					// GiftTag        []struct{} `pd:"32"`
					// ReceiverUinfo  struct{}   `pd:"33"`
					// FaceEffectV2   struct{}   `pd:"34"`
					// GiftInfo       struct{}   `pd:"35"`
					// GiftTipPrice   int64      `pd:"36"`
					// FaceEffect     struct{}   `pd:"37"`
					// Benefits       struct{}   `pd:"38"`
					// EffectConfig   struct{}   `pd:"39"`
				} `pd:"10"`
				// Switch bool `pd:"11"`
				// Test   int  `pd:"12"`
				// WealthInfo  struct{} `pd:"13"`
				// GroupMedal  struct{} `pd:"14"`
				// SenderUinfo struct{} `pd:"15"`
			}
		} `json:"data"`
	}{}

	if e := json.Unmarshal(raw, &s); e != nil {
		t.Fatal(e)
	} else {
		decoder := NewPdDecoder()
		for pd := range decoder.LoadBase64S(s.Data.Pb).Range() {
			switch pd.Type() {
			case 8:
			case 10:
				// for v := range pd.Slice() {
				// 	for pd := range v.Range() {
				// 		t.Log("", pd.Type())
				// 	}
				// }
			case 11:
			case 13:
			case 15:
			default:
				t.Log(pd.Type())
			}
		}
		if e := decoder.UnmarshalBase64S(s.Data.Pb, &s.Data.PbS); e != nil {
			t.Fatal(e)
		} else if s, e := json.MarshalIndent(s.Data.PbS, " ", " "); e == nil {
			t.Log(string(s))
		}
	}
}

func TestS(t *testing.T) {
	base64S := `CMCi88gCEhLliKvmrbvlnKjngavmmJ/kuIoiAgMBKAEwxdMFOMPRj8MGQOO56dq2NEooCPZlEBYaBeW4hVBpIMuoaSjLqGkwkrvKAjjLqGlAAWDF0wVoqoLsF2IAeIaW3o3Zo4mnGJoBALIB0wEIwKLzyAISYQoS5Yir5q275Zyo54Gr5pif5LiKEktodHRwczovL2kwLmhkc2xiLmNvbS9iZnMvZmFjZS9kODM4ZjZkYTVkZDE4NTdhYWU3NzkzYTIwM2ZmNTdlYTkwYjNlMGUwLndlYnAaYgoF5biFUGkQFhjLqGkgkrvKAijLqGkwy6hpOPoNSAFQ9mVgqoLsF3oJIzQzQjNFM0NDggEJIzQzQjNFM0NDigEJIzVGQzdGNEZGkgEJI0ZGRkZGRkZGmgEJIzAwMzA4Qzk5IgIIGjIAugEA`

	var (
		msgType uint32
		uname   string
	)
	for pd := range NewPdDecoder().LoadBase64S(base64S).Range() {
		fmt.Println(pd.Type())
		switch pd.Type() {
		case 22:
			for pd1 := range pd.Child().Range() {
				fmt.Println("=>", pd1.Type())
				switch pd1.Type() {
				case 2:
					for pd2 := range pd1.Child().Range() {
						fmt.Println("=> =>", pd2.Type())
						switch pd2.Type() {
						case 2:
							fmt.Println("=> ", string(pd2.Bytes()))
						case 4:
							fmt.Println("=> ", pd2.Bool())
						}
					}
				}
			}
		case 5:
			msgType = pd.Uint32()
		case 2:
			uname = string(pd.Bytes())
		default:
		}
	}
	if msgType != 1 || uname != "别死在火星上" {
		t.Fatal()
	}
}
func TestU(t *testing.T) {
	base64S := `CMCi88gCEhLliKvmrbvlnKjngavmmJ/kuIoiAgMBKAEwxdMFOMPRj8MGQOO56dq2NEooCPZlEBYaBeW4hVBpIMuoaSjLqGkwkrvKAjjLqGlAAWDF0wVoqoLsF2IAeIaW3o3Zo4mnGJoBALIB0wEIwKLzyAISYQoS5Yir5q275Zyo54Gr5pif5LiKEktodHRwczovL2kwLmhkc2xiLmNvbS9iZnMvZmFjZS9kODM4ZjZkYTVkZDE4NTdhYWU3NzkzYTIwM2ZmNTdlYTkwYjNlMGUwLndlYnAaYgoF5biFUGkQFhjLqGkgkrvKAijLqGkwy6hpOPoNSAFQ9mVgqoLsF3oJIzQzQjNFM0NDggEJIzQzQjNFM0NDigEJIzVGQzdGNEZGkgEJI0ZGRkZGRkZGmgEJIzAwMzA4Qzk5IgIIGjIAugEA`

	type S struct {
		MsgType int    `json:"msgType" pd:"5"`
		Uname   string `pd:"2"`
	}

	ss := S{}

	decoder := NewPdDecoder()
	if e := decoder.UnmarshalBase64B(unsafe.S2B(base64S), &ss); e != nil {
		t.Fatal(e)
	}
	if ss.MsgType != 1 || ss.Uname != "别死在火星上" {
		t.Fatal()
	}
}

func TestU2(t *testing.T) {
	base64S := `CMCi88gCEhLliKvmrbvlnKjngavmmJ/kuIoiAgMBKAEwxdMFOMPRj8MGQOO56dq2NEooCPZlEBYaBeW4hVBpIMuoaSjLqGkwkrvKAjjLqGlAAWDF0wVoqoLsF2IAeIaW3o3Zo4mnGJoBALIB0wEIwKLzyAISYQoS5Yir5q275Zyo54Gr5pif5LiKEktodHRwczovL2kwLmhkc2xiLmNvbS9iZnMvZmFjZS9kODM4ZjZkYTVkZDE4NTdhYWU3NzkzYTIwM2ZmNTdlYTkwYjNlMGUwLndlYnAaYgoF5biFUGkQFhjLqGkgkrvKAijLqGkwy6hpOPoNSAFQ9mVgqoLsF3oJIzQzQjNFM0NDggEJIzQzQjNFM0NDigEJIzVGQzdGNEZGkgEJI0ZGRkZGRkZGmgEJIzAwMzA4Qzk5IgIIGjIAugEA`

	type InteractWord struct {
		FansMedalInfo struct {
			TargetId     int `pd:"1"`
			AnchorRoomid int `pd:"12"`
		} `pd:"9"`
		UserInfo struct {
			Base struct {
				IsMystery bool `pd:"4"`
			} `pd:"2"`
		} `pd:"22"`
		MsgType uint   `json:"msgType" pd:"5"`
		Uname   []byte `pd:"2"`
	}

	ss := InteractWord{}

	decoder := NewPdDecoder()
	if e := decoder.UnmarshalBase64B(unsafe.S2B(base64S), &ss); e != nil {
		t.Fatal(e)
	}
	if ss.MsgType != 1 {
		t.Fatal()
	}
	if string(ss.Uname) != "别死在火星上" {
		t.Fatal()
	}
	if ss.FansMedalInfo.TargetId != 13046 {
		t.Fatal()
	}
	if ss.FansMedalInfo.AnchorRoomid != 92613 {
		t.Fatal()
	}
}

func BenchmarkXxx(b *testing.B) {
	base64S, _ := base64.StdEncoding.DecodeString(`CMCi88gCEhLliKvmrbvlnKjngavmmJ/kuIoiAgMBKAEwxdMFOMPRj8MGQOO56dq2NEooCPZlEBYaBeW4hVBpIMuoaSjLqGkwkrvKAjjLqGlAAWDF0wVoqoLsF2IAeIaW3o3Zo4mnGJoBALIB0wEIwKLzyAISYQoS5Yir5q275Zyo54Gr5pif5LiKEktodHRwczovL2kwLmhkc2xiLmNvbS9iZnMvZmFjZS9kODM4ZjZkYTVkZDE4NTdhYWU3NzkzYTIwM2ZmNTdlYTkwYjNlMGUwLndlYnAaYgoF5biFUGkQFhjLqGkgkrvKAijLqGkwy6hpOPoNSAFQ9mVgqoLsF3oJIzQzQjNFM0NDggEJIzQzQjNFM0NDigEJIzVGQzdGNEZGkgEJI0ZGRkZGRkZGmgEJIzAwMzA4Qzk5IgIIGjIAugEA`)
	type InteractWord struct {
		MsgType uint `json:"msgType" pd:"5"`
	}

	d := NewPdDecoder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.Unmarshal(base64S, new(InteractWord))
	}
}

func BenchmarkJson(b *testing.B) {
	var base64S = `{"msgType":1}`
	type InteractWord struct {
		MsgType uint `json:"msgType" pd:"5"`
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(unsafe.S2B(base64S), new(InteractWord))
	}
}

func Test_Float(t *testing.T) {
	raw := []byte{24, 45, 68, 84, 251, 33, 9, 64}
	t.Log(math.Float64frombits(binary.LittleEndian.Uint64(raw)))
}
