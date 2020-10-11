package F

import (
	"strconv"
	"strings"

	c "github.com/qydysky/bili_danmu/CV"
	p "github.com/qydysky/part"
)

type api struct {
	Roomid int
	Uid int
	Url []string
	Live []string
	Live_status float64
	Locked bool
	Token string
}

var apilog = p.Logf().New().Base(-1, "api.go").Level(2)
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
		// apilog.E("uid", tmp.Err)
	} else if i,err := strconv.Atoi(tmp.RS); err != nil{
		apilog.E("uid", err)
	} else {
		o.Uid = i
	}
	//Title
	if e := r.S(`"title":"`, `",`, 0, 0).Err;e == nil {
		c.Title = r.RS
	}
	//VERSION
	if e := r.S(`player-loader-`, `.min`, 0, 0).Err;e == nil {
		c.VERSION = r.RS
	}
	apilog.W("api version", c.VERSION)
	//roomid
	if tmp := r.S(`"room_id":`, `,`, 0, 0);tmp.Err != nil {
		// apilog.E("room_id", tmp.Err)
	} else if i,err := strconv.Atoi(tmp.RS); err != nil{
		apilog.E("room_id", err)
	} else {
		apilog.T("ok")
		o.Roomid = i
	}

	if o.Roomid != 0 && o.Uid != 0 {return}

	{
		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:"https://api.live.bilibili.com/room/v1/Room/room_init?id=" + Roomid,
			Referer:"https://live.bilibili.com/" + Roomid,
			Timeout:10,
			Retry:2,
		});err != nil {
			apilog.E(err)
			return
		}
		res := string(req.Respon)
		if msg := p.Json().GetValFrom(res, "msg");msg == nil || msg != "ok" {
			apilog.E("msg", msg)
			return
		}
		if Uid := p.Json().GetValFrom(res, "data.uid");Uid == nil {
			apilog.E("data.uid", Uid)
			return
		} else {
			o.Uid = int(Uid.(float64))
		}

		if room_id := p.Json().GetValFrom(res, "data.room_id");room_id == nil {
			apilog.E("data.room_id", room_id)
			return
		} else {
			apilog.T("ok")
			o.Roomid = int(room_id.(float64))
		}
		if is_locked := p.Json().GetValFrom(res, "data.is_locked");is_locked == nil {
			apilog.E("data.is_locked", is_locked)
			return
		} else if is_locked.(bool) {
			apilog.W("直播间封禁中")
			o.Locked = true
			return
		}
	}
	return
}

func (i *api) Get_live(qn ...string) (o *api) {
	o = i
	if o.Roomid == 0 {
		apilog.E("还未New_api")
		return
	}

	Cookie := c.Cookie
	if i := strings.Index(Cookie, "PVID="); i != -1 {
		if d := strings.Index(Cookie[i:], ";"); d == -1 {
			Cookie = Cookie[:i]
		} else {
			Cookie = Cookie[:i] + Cookie[i + d + 1:]
		}
	} else {
		qn = []string{}
	}

	if len(qn) == 0 || qn[0] == "0" {//html获取
		r := p.Get(p.Rval{
			Url:"https://live.bilibili.com/" + strconv.Itoa(o.Roomid),
			Cookie:Cookie,
		})
		if e := r.S(`"durl":[`, `]`, 0, 0).Err;e == nil {
			if urls := p.Json().GetArrayFrom("[" + r.RS + "]", "url");urls != nil {
				apilog.W("直播中")
				o.Live_status = 1
				for _,v := range urls {
					o.Live = append(o.Live, v.(string))
				}
				return
			}
		}
		if e := r.S(`player-loader-`, `.min`, 0, 0).Err;e == nil {
			c.VERSION = r.RS
		}
		apilog.W("api version", c.VERSION)
	}

	cu_qn := "0"
	{//api获取
		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:"https://api.live.bilibili.com/xlive/web-room/v1/index/getRoomPlayInfo?play_url=1&mask=1&qn=0&platform=web&ptype=16&room_id=" + strconv.Itoa(o.Roomid),
			Referer:"https://live.bilibili.com/" + strconv.Itoa(o.Roomid),
			Timeout:10,
			Cookie:Cookie,
			Retry:2,
		});err != nil {
			apilog.E(err)
			return
		}
		res := string(req.Respon)
		if code := p.Json().GetValFromS(res, "code");code == nil || code.(float64) != 0 {
			apilog.E("code", code)
			return
		}
		if is_locked := p.Json().GetValFrom(res, "data.is_locked");is_locked == nil {
			apilog.E("data.is_locked", is_locked)
			return
		} else if is_locked.(bool) {
			apilog.W("直播间封禁中")
			o.Locked = true
			return
		}
		if live_status := p.Json().GetValFrom(res, "data.live_status");live_status == nil {
			apilog.E("data.live_status", live_status)
			return
		} else {
			o.Live_status = live_status.(float64)
			switch live_status.(float64) {
			case 2:
				apilog.W("轮播中")
				return
			case 0: //未直播
				apilog.W("未在直播")
				return
			case 1:
				apilog.W("直播中")
			default:
				apilog.W("live_status:", live_status)
			}
		}
		if urls := p.Json().GetArrayFrom(p.Json().GetValFrom(res, "data.play_url.durl"), "url");urls == nil {
			apilog.E("url", urls)
			return
		} else {
			apilog.T("ok")
			o.Live = []string{}
			for _,v := range urls {
				o.Live = append(o.Live, v.(string))
			}
		}
		if i := p.Json().GetValFrom(res, "data.play_url.current_qn"); i != nil {
			cu_qn = strconv.Itoa(int(i.(float64)))
		}

		if len(qn) != 0 && qn[0] != "0" {
			if _,ok := c.Default_qn[qn[0]];!ok{
				apilog.W("清晰度未找到", qn[0], ",使用默认")
				return
			}
			if err := req.Reqf(p.Rval{
				Url:"https://api.live.bilibili.com/xlive/web-room/v1/playUrl/playUrl?cid=" + strconv.Itoa(o.Roomid) + "&qn=" + qn[0] + "&platform=web&https_url_req=1&ptype=16",
				Referer:"https://live.bilibili.com/" + strconv.Itoa(o.Roomid),
				Timeout:10,
				Cookie:Cookie,
				Retry:2,
			});err != nil {
				apilog.E(err)
				return
			}
			res = string(req.Respon)
			if urls := p.Json().GetArrayFrom(p.Json().GetValFrom(res, "data.durl"), "url");urls == nil {
				apilog.E("url", urls)
				return
			} else {
				apilog.T("ok")
				o.Live = []string{}
				for _,v := range urls {
					o.Live = append(o.Live, v.(string))
				}
			}
			if i := p.Json().GetValFrom(res, "data.current_qn"); i != nil {
				cu_qn = strconv.Itoa(int(i.(float64)))
				c.Live_qn = cu_qn
			}
		}
	}

	if v,ok := c.Default_qn[cu_qn];ok {
		apilog.W("当前清晰度:", v)
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