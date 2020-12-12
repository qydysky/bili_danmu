package F

import (
	"time"
	"fmt"
	"strconv"
	"strings"

	c "github.com/qydysky/bili_danmu/CV"
	g "github.com/qydysky/part/get"
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

	r := g.Get(p.Rval{
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
			Header:map[string]string{
				`Referer`:"https://live.bilibili.com/" + Roomid,
			},
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

	if len(qn) == 0 || qn[0] == "0" || qn[0] == "" {//html获取
		r := g.Get(p.Rval{
			Url:"https://live.bilibili.com/" + strconv.Itoa(o.Roomid),
			Header:map[string]string{
				`Cookie`:Cookie,
			},
		})
		if e := r.S(`"durl":[`, `]`, 0, 0).Err;e == nil {
			if urls := p.Json().GetArrayFrom("[" + r.RS + "]", "url");urls != nil {
				apilog.W("直播中")
				c.Liveing = true
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
		if e := r.S(`"live_time":"`, `"`, 0, 0).Err;e == nil {
			c.Live_Start_Time,_ = time.Parse("2006-01-02 15:04:05", r.RS)
		}
		apilog.W("api version", c.VERSION)
	}

	cu_qn := "0"
	{//api获取
		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:"https://api.live.bilibili.com/xlive/web-room/v1/index/getRoomPlayInfo?play_url=1&mask=1&qn=0&platform=web&ptype=16&room_id=" + strconv.Itoa(o.Roomid),
			Header:map[string]string{
				`Referer`:"https://live.bilibili.com/" + strconv.Itoa(o.Roomid),
				`Cookie`:Cookie,
			},
			Timeout:10,
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
				c.Liveing = false
				apilog.W("轮播中")
				return
			case 0: //未直播
				c.Liveing = false
				apilog.W("未在直播")
				return
			case 1:
				c.Liveing = true
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
		if i := p.Json().GetValFrom(res, "data.live_time"); i != nil {
			c.Live_Start_Time = time.Unix(int64(i.(float64)),0).In(time.FixedZone("UTC-8", -8*60*60))
		}
		if len(qn) != 0 && qn[0] != "0" && qn[0] != "" {
			if _,ok := c.Default_qn[qn[0]];!ok{
				apilog.W("清晰度未找到", qn[0], ",使用默认")
				return
			}
			if err := req.Reqf(p.Rval{
				Url:"https://api.live.bilibili.com/xlive/web-room/v1/playUrl/playUrl?cid=" + strconv.Itoa(o.Roomid) + "&qn=" + qn[0] + "&platform=web&https_url_req=1&ptype=16",
				Header:map[string]string{
					`Cookie`:Cookie,
					`Referer`:"https://live.bilibili.com/" + strconv.Itoa(o.Roomid),
				},
				Timeout:10,
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
		Header:map[string]string{
			`Referer`:"https://live.bilibili.com/" + Roomid,
		},
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

func Get_face_src(uid string) (string) {
	if uid == "" {return ""}

	req := p.Req()
	if err := req.Reqf(p.Rval{
		Url:"https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuMedalAnchorInfo?ruid=" + uid,
		Header:map[string]string{
			`Referer`:"https://live.bilibili.com/" + strconv.Itoa(c.Roomid),
		},
		Timeout:10,
		Retry:2,
	});err != nil {
		apilog.Base(1, "获取face").E(err)
		return ""
	}
	res := string(req.Respon)
	if msg := p.Json().GetValFromS(res, "message");msg == nil || msg != "0" {
		apilog.Base(1, "获取face").E("message", msg)
		return ""
	}

	rface := p.Json().GetValFromS(res, "data.rface")
	if rface == nil {
		apilog.Base(1, "获取face").E("data.rface", rface)
		return ""
	}
	return rface.(string) + `@58w_58h`
}

func (i *api) Get_OnlineGoldRank() {
	if i.Uid == 0 || c.Roomid == 0 {
		apilog.Base(1, "Get_OnlineGoldRank").E("i.Uid == 0 || c.Roomid == 0")
		return
	}
	var session_roomid = c.Roomid
	var self_loop func(page int)
	self_loop = func(page int){
		if page <= 0 || session_roomid != c.Roomid{return}
		// apilog.Base(1, "self_loop").E(page)

		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:`https://api.live.bilibili.com/xlive/general-interface/v1/rank/getOnlineGoldRank?ruid=`+strconv.Itoa(i.Uid)+`&roomId=`+strconv.Itoa(c.Roomid)+`&page=`+strconv.Itoa(page)+`&pageSize=20`,
			Header:map[string]string{
				`Host`: `api.live.bilibili.com`,
				`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
				`Accept`: `application/json, text/plain, */*`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`: `https://live.bilibili.com`,
				`Connection`: `keep-alive`,
				`Pragma`: `no-cache`,
				`Cache-Control`: `no-cache`,
				`Referer`:"https://live.bilibili.com/" + strconv.Itoa(c.Roomid),
			},
			Timeout:3,
			Retry:2,
		});err != nil {
			apilog.Base(1, "获取OnlineGoldRank").E(err)
			return
		}
		res := string(req.Respon)
		if msg := p.Json().GetValFromS(res, "message");msg == nil || msg != "0" {
			apilog.Base(1, "获取OnlineGoldRank").E("message", msg)
			return
		}
		if onlineNum := p.Json().GetValFromS(res, "data.onlineNum");onlineNum == nil {
			apilog.Base(1, "获取onlineNum").E("onlineNum", onlineNum)
			return
		} else {
			tmp_onlineNum := onlineNum.(float64)
			if tmp_onlineNum == 0 {
				// apilog.Base(1, "获取tmp_onlineNum").E("tmp_onlineNum", tmp_onlineNum)
				return
			}

			var score = 0.0
			if tmp_score_list := p.Json().GetArrayFrom(p.Json().GetValFromS(res, "data.OnlineRankItem"), "score");len(tmp_score_list) != 0 {
				for _,v := range tmp_score_list {
					score += v.(float64)/10
				}
			}
			c.Danmu_Main_mq.Push(c.Danmu_Main_mq_item{//传入消息队列
				Class:`c.Rev_add`,
				Data:score,
			})

			if rank_list := p.Json().GetArrayFrom(p.Json().GetValFromS(res, "data.OnlineRankItem"), "userRank");rank_list == nil {
				apilog.Base(1, "获取 rank_list").E("rank_list", len(rank_list))
				return
			} else if len(rank_list) == 0 {
				// apilog.Base(1, "获取 rank_list").E("rank_list == tmp_onlineNum")
				return
			} else {
				p.Sys().Timeoutf(1)
				// apilog.Base(1, "获取page").E(page, score)
				self_loop(page+1)
				return
			}
		}
	}
	// apilog.Base(1, "Get_OnlineGoldRank").E("start")

	// apilog.Base(1, "获取score").E("score", self_loop(1))
	self_loop(1)
	apilog.Base(1, "获取score").W("以往营收获取成功", fmt.Sprintf("%.2f", c.Rev))
	// c.Danmu_Main_mq.Push(c.Danmu_Main_mq_item{//传入消息队列
	// 	Class:`c.Rev_add`,
	// 	Data:self_loop(1),
	// })
	return
}

var guard_num_get_limit = p.Limit(1,1000,2000)//频率限制1次/1s，最大等待时间2s
func (i *api) Get_guardNum() {
	if i.Uid == 0 || c.Roomid == 0 {
		apilog.Base(1, "Get_guardNum").E("i.Uid == 0 || c.Roomid == 0")
		return
	}
	if guard_num_get_limit.TO() {return}//超额请求阻塞，超时将取消

	req := p.Req()
	if err := req.Reqf(p.Rval{
		Url:`https://api.live.bilibili.com/xlive/app-room/v2/guardTab/topList?roomid=`+strconv.Itoa(c.Roomid)+`&page=1&ruid=`+strconv.Itoa(i.Uid)+`&page_size=29`,
		Header:map[string]string{
			`Host`: `api.live.bilibili.com`,
			`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
			`Accept`: `application/json, text/plain, */*`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Origin`: `https://live.bilibili.com`,
			`Connection`: `keep-alive`,
			`Pragma`: `no-cache`,
			`Cache-Control`: `no-cache`,
			`Referer`:"https://live.bilibili.com/" + strconv.Itoa(c.Roomid),
		},
		Timeout:3,
		Retry:2,
	});err != nil {
		apilog.Base(1, "获取guardNum").E(err)
		return
	}
	res := string(req.Respon)
	if msg := p.Json().GetValFromS(res, "message");msg == nil || msg != "0" {
		apilog.Base(1, "获取guardNum").E("message", msg)
		return
	}
	if num := p.Json().GetValFromS(res, "data.info.num");num == nil {
		apilog.Base(1, "获取num").E("num", num)
		return
	} else {
		c.GuardNum = int(num.(float64))
		apilog.Base(1, "获取guardNum").W("舰长数获取成功", c.GuardNum)
	}
	return
}
