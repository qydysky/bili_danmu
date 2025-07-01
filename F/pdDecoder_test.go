package F

import (
	"testing"
)

func TestS(t *testing.T) {
	base64S := `CMCi88gCEhLliKvmrbvlnKjngavmmJ/kuIoiAgMBKAEwxdMFOMPRj8MGQOO56dq2NEooCPZlEBYaBeW4hVBpIMuoaSjLqGkwkrvKAjjLqGlAAWDF0wVoqoLsF2IAeIaW3o3Zo4mnGJoBALIB0wEIwKLzyAISYQoS5Yir5q275Zyo54Gr5pif5LiKEktodHRwczovL2kwLmhkc2xiLmNvbS9iZnMvZmFjZS9kODM4ZjZkYTVkZDE4NTdhYWU3NzkzYTIwM2ZmNTdlYTkwYjNlMGUwLndlYnAaYgoF5biFUGkQFhjLqGkgkrvKAijLqGkwy6hpOPoNSAFQ9mVgqoLsF3oJIzQzQjNFM0NDggEJIzQzQjNFM0NDigEJIzVGQzdGNEZGkgEJI0ZGRkZGRkZGmgEJIzAwMzA4Qzk5IgIIGjIAugEA`

	var (
		msgType uint32
		uname   string
	)
	for pd := range NewPdDecoder().LoadBase64(base64S).Range() {
		switch pd.Type() {
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
