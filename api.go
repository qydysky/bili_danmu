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

var apilog = p.Logf().New().Base(-1, "api.go").Level(LogLevel)
func New_api(Roomid int) (o *api) {
	apilog.Base(-1, "新建")
	defer apilog.Base(0)

	apilog.T("ok")
	o = new(api)
	o.Roomid = Roomid
	o.Get_info()

	return
}

func (i *api) Get_info() (o *api) {
	o = i
	apilog.Base(-1, "获取房号")
	defer apilog.Base(0)

	if o.Roomid == 0 {
		apilog.E("还未New_api")
		return
	}
	Roomid := strconv.Itoa(o.Roomid)

	r := p.Get(p.Rval{
		Url:"https://live.bilibili.com/" + Roomid,
	})
	//uid
	if tmp := r.S(`"uid":`, `,`, 0, 0);tmp.Err != nil {
		apilog.E("uid", tmp.Err)
		return
	} else if i,err := strconv.Atoi(tmp.RS); err != nil{
		apilog.E("uid", err)
		return
	} else {
		o.Uid = i
	}
	//roomid
	if tmp := r.S(`"room_id":`, `,`, 0, 0);tmp.Err != nil {
		apilog.E("room_id", tmp.Err)
		return
	} else if i,err := strconv.Atoi(tmp.RS); err != nil{
		apilog.E("room_id", err)
		return
	} else {
		apilog.T("ok")
		o.Roomid = i
	}
	return
}

func (i *api) Get_host_Token() (o *api) {
	o = i
	apilog.Base(-1, "获取host key")
	defer apilog.Base(0)

	if o.Roomid == 0 {
		apilog.E("还未New_api")
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
		apilog.E(err)
		return
	}
	res := string(req.Respon)
	if msg := p.Json().GetValFromS(res, "message");msg == nil || msg != "0" {
		apilog.E("message", msg)
		return
	}

	_Token := p.Json().GetValFromS(res, "data.token")
	if _Token == nil {
		apilog.E("data.token", _Token, res)
		return
	}
	o.Token = _Token.(string)

	if host_list := p.Json().GetValFromS(res, "data.host_list");host_list == nil {
		apilog.E("data.host_list", host_list)
		return
	} else {
		for k, v := range host_list.([]interface{}) {
			if _host := p.Json().GetValFrom(v, "host");_host == nil {
				apilog.E("data.host_list[", k, "].host", _host)
				continue
			} else {
				o.Url = append(o.Url, "wss://" + _host.(string) + "/sub")
			}			
		}
		apilog.T("ok")
	}

	return
}