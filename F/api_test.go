package F

import (
	"testing"

	c "github.com/qydysky/bili_danmu/CV"
)

func TestCookie(t *testing.T) {
	//获取cookie
	Api.Get(c.C, `Cookie`)
	//获取LIVE_BUVID
	// Api.Get(c.C, `LIVE_BUVID`)

	// if _, ok := c.C.Cookie.Load("LIVE_BUVID"); !ok {
	// 	t.Fatal()
	// }
	if _, ok := c.C.Cookie.Load("buvid3"); !ok {
		t.Fatal()
	}
}

func Test_SearchUP(t *testing.T) {
	//获取cookie
	Api.Get(c.C, `Cookie`)
	//获取LIVE_BUVID
	// Api.Get(c.C, `LIVE_BUVID`)

	if v := SearchUP("qydysky"); len(v) == 0 || v[0].Roomid != 394988 {
		t.Fatal()
	}
}

func Test_Title(t *testing.T) {
	//获取cookie
	Api.Get(c.C, `Cookie`)
	//获取LIVE_BUVID
	// Api.Get(c.C, `LIVE_BUVID`)
	c.C.Roomid = 394988
	Api.Get(c.C, `Title`)
	if c.C.Title == `` {
		t.Fatal()
	}
}

func Test_Html(t *testing.T) {
	//获取cookie
	Api.Get(c.C, `Cookie`)
	//获取LIVE_BUVID
	// Api.Get(c.C, `LIVE_BUVID`)
	c.C.Roomid = 394988
	c.C.UpUid = 0
	Api.common = c.C
	if _, e := Api.htmlLive(); e != nil {
		t.Fatal(e)
	}
	if c.C.UpUid == 0 {
		t.Fatal()
	}
}

func Test_getRoomPlayInfo(t *testing.T) {
	//获取cookie
	Api.Get(c.C, `Cookie`)
	//获取LIVE_BUVID
	// Api.Get(c.C, `LIVE_BUVID`)
	c.C.Roomid = 394988
	c.C.UpUid = 0
	Api.common = c.C
	if _, e := Api.getRoomPlayInfoLive(); e != nil {
		t.Fatal(e)
	}
	if c.C.UpUid == 0 {
		t.Fatal()
	}
}

func Test_getRoomPlayInfoByQn(t *testing.T) {
	//获取cookie
	Api.Get(c.C, `Cookie`)
	//获取LIVE_BUVID
	// Api.Get(c.C, `LIVE_BUVID`)
	c.C.Roomid = 394988
	c.C.UpUid = 0
	Api.common = c.C
	if _, e := Api.getRoomPlayInfoByQnLive(); e != nil {
		t.Fatal(e)
	}
	if c.C.UpUid == 0 {
		t.Fatal()
	}
}

func Test_getDanmuInfo(t *testing.T) {
	//获取cookie
	Api.Get(c.C, `Cookie`)
	//获取LIVE_BUVID
	// Api.Get(c.C, `LIVE_BUVID`)
	c.C.Roomid = 394988
	c.C.WSURL = []string{}
	Api.Get(c.C, `WSURL`)
	if len(c.C.WSURL) == 0 {
		t.Fatal()
	}
}

func Test_Get_guardNum(t *testing.T) {
	//获取cookie
	Api.Get(c.C, `Cookie`)
	//获取LIVE_BUVID
	// Api.Get(c.C, `LIVE_BUVID`)
	c.C.Roomid = 394988
	c.C.GuardNum = -1
	Api.Get(c.C, `GuardNum`)
	if c.C.GuardNum == -1 {
		t.Fatal()
	}
}
