package bili_danmu

import (
	"strconv"

	p "github.com/qydysky/part"
)

type api struct {
	Roomid int
	Uid int
	Url []string
	Token string
}

func New_api(Roomid int) (o *api) {
	l := p.Logf().New().Level(LogLevel).T("New_api")
	defer l.Block()

	l.T("->", "ok")
	o = new(api)
	o.Roomid = Roomid
	o.Get_info()

	return
}

func (i *api) Get_info() (o *api) {
	o = i 
	l := p.Logf().New().Level(LogLevel).T("*api.Get_info")
	defer l.Block()

	if o.Roomid == 0 {
		l.E("->", "还未New_api")
		return
	}
	Roomid := strconv.Itoa(o.Roomid)

	req := p.Req()
	if err := req.Reqf(p.Rval{
		Url:"https://api.live.bilibili.com/room/v1/Room/room_init?id=" + Roomid,
		Referer:"https://live.bilibili.com/" + Roomid,
		Timeout:10,
		Retry:2,
	});err != nil {
		l.E("->", err)
		return
	}
	res := string(req.Respon)
	if msg := p.Json().GetValFrom(res, "msg");msg == nil || msg != "ok" {
		l.E("->", "msg", msg)
		return
	}
	if Uid := p.Json().GetValFrom(res, "data.uid");Uid == nil {
		l.E("->", "data.uid", Uid)
		return
	} else {
		o.Uid = int(Uid.(float64))
	}

	if room_id := p.Json().GetValFrom(res, "data.room_id");room_id == nil {
		l.E("->", "data.room_id", room_id)
		return
	} else {
		l.T("->", "ok")
		o.Roomid = int(room_id.(float64))
	}
	return
}

func (i *api) Get_host_Token() (o *api) {
	o = i
	l := p.Logf().New().Level(LogLevel).T("*api.Get_host_Token")
	defer l.Block()

	if o.Roomid == 0 {
		l.E("->", "还未New_api")
		return
	}
	Roomid := strconv.Itoa(o.Roomid)


	req := p.Req()
	if err := req.Reqf(p.Rval{
		Url:"https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?type=0&id=" + Roomid,
		Referer:"https://live.bilibili.com/" + Roomid,
		Timeout:10,
		Retry:2,
	});err != nil {
		l.E("->", err)
		return
	}
	res := string(req.Respon)
	if msg := p.Json().GetValFrom(res, "message");msg == nil || msg != "0" {
		l.E("->", "message", msg)
		return
	}

	_Token := p.Json().GetValFrom(res, "data.token")
	if _Token == nil {
		l.E("->", "data.token", _Token, res)
		return
	}
	o.Token = _Token.(string)

	if host_list := p.Json().GetValFrom(res, "data.host_list");host_list == nil {
		l.E("->", "data.host_list", host_list)
		return
	} else {
		for k, v := range host_list.([]interface{}) {
			if _host := p.Json().GetValFrom(v, "host");_host == nil {
				l.E("->", "data.host_list[", k, "].host", _host)
				continue
			} else {
				o.Url = append(o.Url, "wss://" + _host.(string) + "/sub")
			}			
		}
		l.T("->", "ok")
	}

	return
}