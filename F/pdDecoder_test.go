package F

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	unsafe "github.com/qydysky/part/unsafe"
)

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
