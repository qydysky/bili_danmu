package F

import (
	"sync"
	"time"
	"fmt"
	"os"
	"strconv"
	"strings"
    "context"
	"encoding/json"
	"net/http"
	"net/url"

	qr "github.com/skip2/go-qrcode"
	c "github.com/qydysky/bili_danmu/CV"
	g "github.com/qydysky/part/get"
	p "github.com/qydysky/part"
	uuid "github.com/gofrs/uuid"
)

type api struct {
	Roomid int
	Uid int
	Url []string
	Live []string
	Live_status float64
	Locked bool
	Token string
	Parent_area_id int
	Area_id int
}

var apilog = c.Log.Base(`api`)
var api_limit = p.Limit(1,2000,15000)//频率限制1次/2s，最大等待时间15s

func New_api(Roomid int) (o *api) {
	apilog.Base_add(`新建`).L(`T: `,"ok")
	o = new(api)
	o.Roomid = Roomid
	o.Parent_area_id = -1
	o.Area_id = -1
	o.Get_info()
	return
}

func (i *api) Get_info() (o *api) {
	o = i
	apilog := apilog.Base_add(`获取房号`)	

	if o.Roomid == 0 {
		apilog.L(`E: `,"还未New_api")
		return
	}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消

	o.Get_LIVE_BUVID()
	
	Roomid := strconv.Itoa(o.Roomid)

	r := g.Get(p.Rval{
		Url:"https://live.bilibili.com/blanc/" + Roomid,
	})
	//uid
	if tmp := r.S(`"uid":`, `,`, 0, 0);tmp.Err != nil {
		// apilog.L(`E: `,"uid", tmp.Err)
	} else if i,err := strconv.Atoi(tmp.RS[0]); err != nil{
		apilog.L(`E: `,"uid", err)
	} else {
		o.Uid = i
	}
	//Title
	if e := r.S(`"title":"`, `",`, 0, 0).Err;e == nil {
		c.Title = r.RS[0]
	}
	//主播id
	if e := r.S(`"base_info":{"uname":"`, `",`, 0, 0).Err;e == nil {
		c.Uname = r.RS[0]
	}
	//分区
	if e := r.S(`"parent_area_id":`, `,`, 0, 0).Err;e == nil {
		if tmp,e := strconv.Atoi(r.RS[0]);e != nil{
			apilog.L(`E: `,"parent_area_id", e)
		} else {o.Parent_area_id = tmp}
	}
	if e := r.S(`"area_id":`, `,`, 0, 0).Err;e == nil {
		if tmp,e := strconv.Atoi(r.RS[0]);e != nil{
			apilog.L(`E: `,"area_id", e)
		} else {o.Area_id = tmp}
	}
	//roomid
	if tmp := r.S(`"room_id":`, `,`, 0, 0);tmp.Err != nil {
		// apilog.L(`E: `,"room_id", tmp.Err)
	} else if i,err := strconv.Atoi(tmp.RS[0]); err != nil{
		apilog.L(`E: `,"room_id", err)
	} else {
		o.Roomid = i
	}

	if	o.Area_id != -1 && 
		o.Parent_area_id != -1 &&
		o.Roomid != 0 &&
		o.Uid != 0 &&
		c.Title != ``{return}

	{//使用其他api
		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:"https://api.live.bilibili.com/xlive/web-room/v1/index/getInfoByRoom?room_id=" + Roomid,
			Header:map[string]string{
				`Referer`:"https://live.bilibili.com/" + Roomid,
				`Cookie`:p.Map_2_Cookies_String(c.Cookie),
			},
			Timeout:10,
			Retry:2,
		});err != nil {
			apilog.L(`E: `,err)
			return
		}
		var tmp struct{
			Code int `json:"code"`
			Message string `json:"message"`
			Data struct{
				Room_info struct{
					Uid int `json:"uid"`
					Room_id int `json:"room_id"`
					Title string `json:"title"`
					Lock_status int `json:"lock_status"`
					Area_id int `json:"area_id"`
					Parent_area_id int `json:"parent_area_id"`
				} `json:"room_info"`
				Anchor_info struct{
					Base_info struct{
						Uname string `json:"uname"`
					} `json:"base_info"`
				} `json:"anchor_info"`
			} `json:"data"`
		}
		if e := json.Unmarshal(req.Respon, &tmp);e != nil{
			apilog.L(`E: `,e)
			return
		}

		//错误响应
		if tmp.Code != 0 {
			apilog.L(`E: `,`code`,tmp.Message)
			return
		}

		//主播
		if tmp.Data.Anchor_info.Base_info.Uname != `` && c.Uname == ``{
			c.Uname = tmp.Data.Anchor_info.Base_info.Uname
		}

		//主播id
		if tmp.Data.Room_info.Uid != 0{
			o.Uid = tmp.Data.Room_info.Uid
		} else {
			apilog.L(`E: `,"data.room_info.parent_area_id = 0")
			return
		}

		//分区
		if tmp.Data.Room_info.Parent_area_id != 0{
			o.Parent_area_id = tmp.Data.Room_info.Parent_area_id
		} else {
			apilog.L(`E: `,"data.room_info.parent_area_id = 0")
			return
		}
		if tmp.Data.Room_info.Area_id != 0{
			o.Area_id = tmp.Data.Room_info.Area_id
		} else {
			apilog.L(`E: `,"data.room_info.Area_id = 0")
			return
		}

		//房间id
		if tmp.Data.Room_info.Room_id != 0{
			o.Roomid = tmp.Data.Room_info.Room_id
		} else {
			apilog.L(`E: `,"data.room_info.room_id = 0")
			return
		}
		
		//房间标题
		if tmp.Data.Room_info.Title != ``{
			c.Title = tmp.Data.Room_info.Title
		} else {
			apilog.L(`E: `,"data.room_info.title = ''")
			return
		}

		//直播间是否被封禁
		if tmp.Data.Room_info.Lock_status == 1{
			apilog.L(`W: `,"直播间封禁中")
			o.Locked = true
			return
		}
	}
	return
}

func (i *api) Get_live(qn ...string) (o *api) {
	o = i
	apilog := apilog.Base_add(`获取直播流`)

	if o.Roomid == 0 {
		apilog.L(`E: `,"还未New_api")
		return
	}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消

	Cookie := p.Map_2_Cookies_String(c.Cookie)
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
			Url:"https://live.bilibili.com/blanc/" + strconv.Itoa(o.Roomid),
			Header:map[string]string{
				`Cookie`:Cookie,
			},
		})
		if e := r.S(`"durl":[`, `]`, 0, 0).Err;e == nil {
			if urls := p.Json().GetArrayFrom("[" + r.RS[0] + "]", "url");urls != nil {
				apilog.L(`W: `,"直播中")
				c.Liveing = true
				o.Live_status = 1
				for _,v := range urls {
					o.Live = append(o.Live, v.(string))
				}
				return
			}
		}
		if e := r.S(`"live_time":"`, `"`, 0, 0).Err;e == nil {
			c.Live_Start_Time,_ = time.Parse("2006-01-02 15:04:05", r.RS[0])
		}
	}

	cu_qn := "0"
	{//api获取
		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:"https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo?no_playurl=0&mask=1&qn=0&platform=web&protocol=0,1&format=0,2&codec=0,1&room_id=" + strconv.Itoa(o.Roomid),
			Header:map[string]string{
				`Referer`:"https://live.bilibili.com/" + strconv.Itoa(o.Roomid),
				`Cookie`:Cookie,
			},
			Timeout:10,
			Retry:2,
		});err != nil {
			apilog.L(`E: `,err)
			return
		}
		res := string(req.Respon)
		if code := p.Json().GetValFromS(res, "code");code == nil || code.(float64) != 0 {
			apilog.L(`E: `,"code", code)
			return
		}
		if is_locked := p.Json().GetValFrom(res, "data.is_locked");is_locked == nil {
			apilog.L(`E: `,"data.is_locked", is_locked)
			return
		} else if is_locked.(bool) {
			apilog.L(`W: `,"直播间封禁中")
			o.Locked = true
			return
		}
		if live_status := p.Json().GetValFrom(res, "data.live_status");live_status == nil {
			apilog.L(`E: `,"data.live_status", live_status)
			return
		} else {
			o.Live_status = live_status.(float64)
			switch live_status.(float64) {
			case 2:
				c.Liveing = false
				apilog.L(`W: `,"轮播中")
				return
			case 0: //未直播
				c.Liveing = false
				apilog.L(`W: `,"未在直播")
				return
			case 1:
				c.Liveing = true
				apilog.L(`W: `,"直播中")
			default:
				apilog.L(`W: `,"live_status:", live_status)
			}
		}
		if codec0 := p.Json().GetValFrom(res, "data.playurl_info.playurl.stream.[0].format.[0].codec.[0]");codec0 != nil {//直播流链接
			base_url := p.Json().GetValFrom(codec0, "base_url")
			if base_url == nil {return}
			url_info := p.Json().GetValFrom(codec0, "url_info")
			if v,ok := url_info.([]interface{});!ok || len(v) == 0 {return}
			for _,v := range url_info.([]interface{}) {
				host := p.Json().GetValFrom(v, "host")
				extra := p.Json().GetValFrom(v, "extra")
				if host == nil || extra == nil {continue}
				o.Live = append(o.Live, host.(string) + base_url.(string) + extra.(string))
			}
		}
		if len(o.Live) == 0 {apilog.L(`E: `,"live url is nil");return}

		if i := p.Json().GetValFrom(res, "data.playurl_info.playurl.stream.[0].format.[0].codec.[0].current_qn"); i != nil {
			cu_qn = strconv.Itoa(int(i.(float64)))
		}
		if i := p.Json().GetValFrom(res, "data.live_time"); i != nil {
			c.Live_Start_Time = time.Unix(int64(i.(float64)),0).In(time.FixedZone("UTC-8", -8*60*60))
		}

		if len(qn) != 0 && qn[0] != "0" && qn[0] != "" {
			var (
				accept_qn_request bool
				tmp_qn int
				e error
			)
			if tmp_qn,e = strconv.Atoi(qn[0]);e != nil {apilog.L(`E: `,`qn error`,e);return}
			if i,ok := p.Json().GetValFrom(res, "data.playurl_info.playurl.stream.[0].format.[0].codec.[0].accept_qn").([]interface{}); ok && len(i) != 0 {
				for _,v := range i {
					if o,ok := v.(float64);ok && int(o) == tmp_qn {accept_qn_request = true}
				}
			}
			if !accept_qn_request {
				apilog.L(`E: `,`qn不在accept_qn中`);
				return
			}
			if _,ok := c.Default_qn[qn[0]];!ok{
				apilog.L(`W: `,"清晰度未找到", qn[0], ",使用默认")
				return
			}
			if err := req.Reqf(p.Rval{
				Url:"https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo?no_playurl=0&mask=1&platform=web&protocol=0,1&format=0,2&codec=0,1&room_id=" + strconv.Itoa(o.Roomid) + "&qn=" + qn[0],
				Header:map[string]string{
					`Cookie`:Cookie,
					`Referer`:"https://live.bilibili.com/" + strconv.Itoa(o.Roomid),
				},
				Timeout:10,
				Retry:2,
			});err != nil {
				apilog.L(`E: `,err)
				return
			}
			res = string(req.Respon)
			if codec0 := p.Json().GetValFrom(res, "data.playurl_info.playurl.stream.[0].format.[0].codec.[0]");codec0 != nil {//直播流链接
				base_url := p.Json().GetValFrom(codec0, "base_url")
				if base_url == nil {return}
				url_info := p.Json().GetValFrom(codec0, "url_info")
				if v,ok := url_info.([]interface{});!ok || len(v) == 0 {return}
				for _,v := range url_info.([]interface{}) {
					host := p.Json().GetValFrom(v, "host")
					extra := p.Json().GetValFrom(v, "extra")
					if host == nil || extra == nil {continue}
					o.Live = append(o.Live, host.(string) + base_url.(string) + extra.(string))
				}
			}
			if len(o.Live) == 0 {apilog.L(`E: `,"live url is nil");return}
	
			if i := p.Json().GetValFrom(res, "data.playurl_info.playurl.stream.[0].format.[0].codec.[0].current_qn"); i != nil {
				cu_qn = strconv.Itoa(int(i.(float64)))
			}
		}
	}

	if v,ok := c.Default_qn[cu_qn];ok {
		apilog.L(`W: `,"当前清晰度:", v)
	}
	return
}

func (i *api) Get_host_Token() (o *api) {
	o = i
	apilog := apilog.Base_add(`获取Token`)

	if o.Roomid == 0 {
		apilog.L(`E: `,"还未New_api")
		return
	}
	Roomid := strconv.Itoa(o.Roomid)
	if api_limit.TO() {return}//超额请求阻塞，超时将取消


	req := p.Req()
	if err := req.Reqf(p.Rval{
		Url:"https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?type=0&id=" + Roomid,
		Header:map[string]string{
			`Referer`:"https://live.bilibili.com/" + Roomid,
			`Cookie`:p.Map_2_Cookies_String(c.Cookie),
		},
		Timeout:10,
		Retry:2,
	});err != nil {
		apilog.L(`E: `,err)
		return
	}
	res := string(req.Respon)
	if msg := p.Json().GetValFromS(res, "message");msg == nil || msg != "0" {
		apilog.L(`E: `,"message", msg)
		return
	}

	_Token := p.Json().GetValFromS(res, "data.token")
	if _Token == nil {
		apilog.L(`E: `,"data.token", _Token, res)
		return
	}
	o.Token = _Token.(string)

	if host_list := p.Json().GetValFromS(res, "data.host_list");host_list == nil {
		apilog.L(`E: `,"data.host_list", host_list)
		return
	} else {
		for k, v := range host_list.([]interface{}) {
			if _host := p.Json().GetValFrom(v, "host");_host == nil {
				apilog.L(`E: `,"data.host_list[", k, "].host", _host)
				continue
			} else {
				o.Url = append(o.Url, "wss://" + _host.(string) + "/sub")
			}			
		}
		apilog.L(`T: `,"ok")
	}

	return
}

func Get_face_src(uid string) (string) {
	if uid == "" {return ""}
	if api_limit.TO() {return ""}//超额请求阻塞，超时将取消
	apilog := apilog.Base_add(`获取头像`)

	req := p.Req()
	if err := req.Reqf(p.Rval{
		Url:"https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuMedalAnchorInfo?ruid=" + uid,
		Header:map[string]string{
			`Referer`:"https://live.bilibili.com/" + strconv.Itoa(c.Roomid),
			`Cookie`:p.Map_2_Cookies_String(c.Cookie),
		},
		Timeout:10,
		Retry:2,
	});err != nil {
		apilog.L(`E: `,err)
		return ""
	}
	res := string(req.Respon)
	if msg := p.Json().GetValFromS(res, "message");msg == nil || msg != "0" {
		apilog.L(`E: `,"message", msg)
		return ""
	}

	rface := p.Json().GetValFromS(res, "data.rface")
	if rface == nil {
		apilog.L(`E: `,"data.rface", rface)
		return ""
	}
	return rface.(string) + `@58w_58h`
}

func (i *api) Get_OnlineGoldRank() {
	if i.Uid == 0 || c.Roomid == 0 {
		apilog.Base_add("Get_OnlineGoldRank").L(`E: `,"i.Uid == 0 || c.Roomid == 0")
		return
	}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消
	apilog := apilog.Base_add(`获取贡献榜`)

	var session_roomid = c.Roomid
	var self_loop func(page int)
	self_loop = func(page int){
		if page <= 0 || session_roomid != c.Roomid{return}

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
				`Cookie`:p.Map_2_Cookies_String(c.Cookie),
			},
			Timeout:3,
			Retry:2,
		});err != nil {
			apilog.L(`E: `,err)
			return
		}
		res := string(req.Respon)
		if msg := p.Json().GetValFromS(res, "message");msg == nil || msg != "0" {
			apilog.L(`E: `,"message", msg)
			return
		}
		if onlineNum := p.Json().GetValFromS(res, "data.onlineNum");onlineNum == nil {
			apilog.L(`E: `,"onlineNum", onlineNum)
			return
		} else {
			tmp_onlineNum := onlineNum.(float64)
			if tmp_onlineNum == 0 {
				return
			}

			var score = 0.0
			if tmp_score_list := p.Json().GetArrayFrom(p.Json().GetValFromS(res, "data.OnlineRankItem"), "score");len(tmp_score_list) != 0 {
				for _,v := range tmp_score_list {
					score += v.(float64)/10
				}
			}
			//传入消息队列
			c.Danmu_Main_mq.Push_tag(`c.Rev_add`,score)

			if rank_list := p.Json().GetArrayFrom(p.Json().GetValFromS(res, "data.OnlineRankItem"), "userRank");rank_list == nil {
				apilog.L(`E: `,"rank_list", len(rank_list))
				return
			} else if len(rank_list) == 0 {
				// apilog.L(`E: `,"rank_list == tmp_onlineNum")
				return
			} else {
				p.Sys().Timeoutf(1)
				self_loop(page+1)
				return
			}
		}
	}

	self_loop(1)
	apilog.Base("获取score").L(`W: `,"以往营收获取成功", fmt.Sprintf("%.2f", c.Rev))
	// c.Danmu_Main_mq.Push(c.Danmu_Main_mq_item{//传入消息队列
	// 	Class:`c.Rev_add`,
	// 	Data:self_loop(1),
	// })
	return
}

//获取热门榜
func (i *api) Get_HotRank() {
	if i.Uid == 0 || c.Roomid == 0 {
		apilog.Base_add("Get_HotRank").L(`E: `,"i.Uid == 0 || c.Roomid == 0")
		return
	}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消
	apilog := apilog.Base_add(`获取热门榜`)

	req := p.Req()
	if err := req.Reqf(p.Rval{
		Url:`https://api.live.bilibili.com/xlive/general-interface/v1/rank/getHotRank?ruid=`+strconv.Itoa(i.Uid)+`&room_id=`+strconv.Itoa(c.Roomid)+`&is_pre=0&page_size=50&source=2&area_id=`+strconv.Itoa(i.Parent_area_id),
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
			`Cookie`:p.Map_2_Cookies_String(c.Cookie),
		},
		Timeout:3,
		Retry:2,
	});err != nil {
		apilog.L(`E: `,err)
		return
	}
	
	var type_item struct{
		Code int `json:"code"`
		Message string `json:"message"`
		Data struct{
			Own struct{
				Rank int `json:"rank"`
				Area_parent_name string `json:"area_parent_name"`
			} `json:"own"`
		} `json:"data"`
	}
	if e := json.Unmarshal(req.Respon, &type_item);e != nil {
		apilog.L(`E: `, e)
	}
	if type_item.Code != 0 {
		apilog.L(`E: `,type_item.Message)
		return
	}
	c.Note = type_item.Data.Own.Area_parent_name + " "
	if type_item.Data.Own.Rank == 0 {
		c.Note += `50+`
	} else {
		c.Note += strconv.Itoa(type_item.Data.Own.Rank)
	}
	apilog.L(`I: `,`热门榜:`,c.Note)
}

func (i *api) Get_guardNum() {
	if i.Uid == 0 || c.Roomid == 0 {
		apilog.Base_add("Get_guardNum").L(`E: `,"i.Uid == 0 || c.Roomid == 0")
		return
	}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消
	apilog := apilog.Base_add(`获取舰长数`)

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
			`Cookie`:p.Map_2_Cookies_String(c.Cookie),
		},
		Timeout:3,
		Retry:2,
	});err != nil {
		apilog.L(`E: `,err)
		return
	}
	res := string(req.Respon)
	if msg := p.Json().GetValFromS(res, "message");msg == nil || msg != "0" {
		apilog.L(`E: `,"message", msg)
		return
	}
	if num := p.Json().GetValFromS(res, "data.info.num");num == nil {
		apilog.L(`E: `,"num", num)
		return
	} else {
		c.GuardNum = int(num.(float64))
		apilog.L(`W: `,"舰长数获取成功", c.GuardNum)
	}
	return
}

func (i *api) Get_Version() {
	Roomid := strconv.Itoa(i.Roomid)
	if api_limit.TO() {return}//超额请求阻塞，超时将取消
	apilog := apilog.Base_add(`获取客户版本`)

	var player_js_url string
	{//获取player_js_url
		r := g.Get(p.Rval{
			Url:"https://live.bilibili.com/blanc/" + Roomid,
		})

		if r.Err != nil {
			apilog.L(`E: `,r.Err)
			return
		}

		r.S2(`<script src=`,`.js`)
		if r.Err != nil {
			apilog.L(`E: `,r.Err)
			return
		}

		for _,v := range r.RS {
			tmp := string(v) + `.js`
			if strings.Contains(tmp,`http`) {continue}
			tmp = `https:` + tmp
			if strings.Contains(tmp,`player`) {
				player_js_url = tmp
				break
			}
		}
		if player_js_url == `` {
			apilog.L(`E: `,`no found player js`)
			return
		}
	}

	{//获取VERSION
		r := g.Get(p.Rval{
			Url:player_js_url,
		})

		if r.Err != nil {
			apilog.L(`E: `,r.Err)
			return
		}

		r.S(`Bilibili HTML5 Live Player v`,` `,0,0)
		if r.Err != nil {
			apilog.L(`E: `,r.Err)
			return
		}
		c.VERSION = r.RS[0]
		apilog.L(`W: `,"api version", c.VERSION)
	}
}

type cookie_lock_item struct{
	sync.RWMutex
}
var cookies_lock = new(cookie_lock_item)
func Get_cookie() {
	if api_limit.TO() {return}//超额请求阻塞，超时将取消
	cookies_lock.Lock()
	defer cookies_lock.Unlock()
	apilog := apilog.Base_add(`获取Cookie`)

	var img_url string
	var oauth string
	{//获取二维码
		r := p.Req()
		if e := r.Reqf(p.Rval{
			Url:`https://passport.bilibili.com/qrcode/getLoginUrl`,
			Timeout:10,
			Retry:2,
		});e != nil {
			apilog.L(`E: `,e)
			return
		}
		res := string(r.Respon)
		if v,ok := p.Json().GetValFromS(res, "status").(bool);!ok || !v {
			apilog.L(`E: `,`getLoginUrl status failed!`)
			return
		} else {
			if v,ok := p.Json().GetValFromS(res, "data.url").(string);ok {
				img_url = v
			}
			if v,ok := p.Json().GetValFromS(res, "data.oauthKey").(string);ok {
				oauth = v
			}
		}
		if img_url == `` || oauth == `` {
			apilog.L(`E: `,`img_url:`,img_url,` oauth:`,oauth)
			return
		}
	}
	var server *http.Server
	{//生成二维码
		qr.WriteFile(img_url,qr.Medium,256,`qr.png`)
		if !p.Checkfile().IsExist(`qr.png`) {
			apilog.L(`E: `,`qr error`)
			return
		}
		go func(){//启动web
			web :=  http.NewServeMux()
			web.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
				w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
				var path string = r.URL.Path[1:]
				http.ServeFile(w, r,`./`+path)
			})
			server = &http.Server{
				Addr:         `127.0.0.1:`+strconv.Itoa(p.Sys().GetFreePort()),
				WriteTimeout: time.Second * 10,
				Handler:      web,
			}
			apilog.L(`W: `,`打开链接扫码登录：`,`http://`+server.Addr+`/qr.png`)
			server.ListenAndServe()
		}()
		p.Sys().Timeoutf(1)
	}
	var cookie string
	{//3s刷新查看是否通过
		max_try := 20
		change_room_sign := false
		c.Danmu_Main_mq.Pull_tag(map[string]func(interface{})(bool){
			`change_room`:func(data interface{})(bool){//房间改变
				change_room_sign = true
				return true
			},
		})

		for max_try > 0 && !change_room_sign {
			max_try -= 1
			p.Sys().Timeoutf(3)
			r := p.Req()
			if e := r.Reqf(p.Rval{
				Url:`https://passport.bilibili.com/qrcode/getLoginInfo`,
				PostStr:`oauthKey=`+oauth,
				Header:map[string]string{
					`Content-Type`:`application/x-www-form-urlencoded; charset=UTF-8`,
					`Referer`: `https://passport.bilibili.com/login`,
					`Cookie`:p.Map_2_Cookies_String(c.Cookie),
				},
				Timeout:10,
				Retry:2,	
			});e != nil {
				apilog.L(`E: `,e)
				return
			}
			res := string(r.Respon)
			if v,ok := p.Json().GetValFromS(res, "status").(bool);!ok {
				apilog.L(`E: `,`getLoginInfo status false`)
				return
			} else if !v {
				if v,ok := p.Json().GetValFromS(res, "message").(string);ok {
					apilog.L(`W: `,`登录中`,v,max_try,`s`)
				}
				continue
			} else {
				apilog.L(`W: `,`登录，并保存了cookie`)
				if v := r.Response.Cookies();len(v) == 0 {
					apilog.L(`E: `,`getLoginInfo cookies len == 0`)
					return
				} else {
					cookie = p.Map_2_Cookies_String(p.Cookies_List_2_Map(v))//cookie to string
				}
				if cookie == `` {
					apilog.L(`E: `,`getLoginInfo cookies ""`)
					return
				} else {break}
			}
		}
		if max_try <= 0 {
			apilog.L(`W: `,`登录取消`)
			return
		}
	}
	{//写入cookie.txt
		for k,v := range p.Cookies_String_2_Map(cookie){
			c.Cookie[k] = v
		}
		f := p.File()
		f.FileWR(p.Filel{
			File:`cookie.txt`,
			Write:true,
			Loc:0,
			Context:[]interface{}{cookie},
		})
	}
	{//关闭web
		if err := server.Shutdown(context.Background()); err != nil {
			apilog.L(`E: `,"HTTP server Shutdown:", err.Error())
		}
		if p.Checkfile().IsExist(`qr.png`) {
			os.RemoveAll(`qr.png`)
			return
		}
	}
}

func (i *api) Switch_FansMedal() {
	if len(c.Cookie) == 0 {return}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消
	apilog := apilog.Base_add(`切换粉丝牌`)
	Cookie := p.Map_2_Cookies_String(c.Cookie)
	{//验证是否本直播间牌子
		r := p.Req()
		if e := r.Reqf(p.Rval{
			Url:`https://api.live.bilibili.com/live_user/v1/UserInfo/get_weared_medal`,
			Header:map[string]string{
				`Cookie`:Cookie,
			},
			Timeout:10,
			Retry:2,
		});e != nil {
			apilog.L(`E: `,e)
			return
		}
		res := string(r.Respon)
		if v,ok := p.Json().GetValFromS(res, "data.roominfo.room_id").(float64);ok && int(v) == c.Roomid {
			return
		}
	}
	var medal_id int
	{//获取牌子列表
		r := p.Req()
		if e := r.Reqf(p.Rval{
			Url:`https://api.live.bilibili.com/fans_medal/v1/FansMedal/get_list_in_room`,
			Header:map[string]string{
				`Cookie`:p.Map_2_Cookies_String(c.Cookie),
			},
			Timeout:10,
			Retry:2,
		});e != nil {
			apilog.L(`E: `,e)
			return
		}
		res := string(r.Respon)
		if v,ok := p.Json().GetValFromS(res, "code").(float64);!ok || v != 0 {
			apilog.L(`E: `,`Get_FansMedal get_list_in_room code`, v)
			return
		} else {
			if v,ok := p.Json().GetValFromS(res, "data").([]interface{});ok{
				for _,item := range v {
					if room_id,ok := p.Json().GetValFrom(item, "room_id").(float64);!ok || int(room_id) != c.Roomid {
						continue
					} else {
						if tmp_medal_id,ok := p.Json().GetValFrom(item, "medal_id").(float64);!ok {
							apilog.L(`E: `,`medal_id error`)
							return
						} else {
							medal_id = int(tmp_medal_id)
							break
						}
					}
				}
			} else {
				apilog.L(`E: `,`data error`)
				return
			}
		}
	}
	var (
		post_url string
		post_str string
	)
	{//生成佩戴信息
		csrf := c.Cookie[`bili_jct`]
		if csrf == `` {apilog.L(`E: `,"Cookie错误,无bili_jct=");return}
		
		post_str = `csrf_token=`+csrf+`&csrf=`+csrf
		if medal_id == 0 {//无牌，不佩戴牌子
			post_url = `https://api.live.bilibili.com/xlive/web-room/v1/fansMedal/take_off`
		} else {
			post_url = `https://api.live.bilibili.com/xlive/web-room/v1/fansMedal/wear`
			post_str = `medal_id=`+strconv.Itoa(medal_id)+`&`+post_str
		}
	}
	{//切换牌子
		r := p.Req()
		if e := r.Reqf(p.Rval{
			Url:post_url,
			PostStr:post_str,
			Header:map[string]string{
				`Cookie`:Cookie,
				`Content-Type`:`application/x-www-form-urlencoded; charset=UTF-8`,
				`Referer`: `https://passport.bilibili.com/login`,
			},
			Timeout:10,
			Retry:2,
		});e != nil {
			apilog.L(`E: `,e)
			return
		}
		res := string(r.Respon)
		if v,ok := p.Json().GetValFromS(res, "code").(float64);ok && v == 0 {
			apilog.L(`W: `,`自动切换粉丝牌`)
			return
		}
		if v,ok := p.Json().GetValFromS(res, "message").(string);ok {
			apilog.L(`E: `,`Get_FansMedal wear message`, v)
		} else {
			apilog.L(`E: `,`Get_FansMedal wear message nil`)
		}
	}
}

//签到
func Dosign() {
	apilog := apilog.Base_add(`签到`).L(`T: `,`签到`)
	if len(c.Cookie) == 0 {apilog.L(`E: `,`失败！无cookie`);return}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消

	{//检查是否签到
		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:`https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/WebGetSignInfo`,
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
				`Referer`:"https://live.bilibili.com/all",
				`Cookie`:p.Map_2_Cookies_String(c.Cookie),
			},
			Timeout:3,
			Retry:2,
		});err != nil {
			apilog.L(`E: `,err)
			return
		}
	
		var msg struct {
			Code int `json:"code"`
			Message string `json:"message"`
			Data struct {
				Status int `json:"status"`
			} `json:"data"`
		}
		if e := json.Unmarshal(req.Respon,&msg);e != nil{
			apilog.L(`E: `,e)
		}
		if msg.Code != 0 {apilog.L(`E: `,msg.Message);return}
		if msg.Data.Status == 1 {//今日已签到
			return
		}
	}

	{//签到
		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:`https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/DoSign`,
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
				`Referer`:"https://live.bilibili.com/all",
				`Cookie`:p.Map_2_Cookies_String(c.Cookie),
			},
			Timeout:3,
			Retry:2,
		});err != nil {
			apilog.L(`E: `,err)
			return
		}
	
		var msg struct {
			Code int `json:"code"`
			Message string `json:"message"`
			Data struct {
				HadSignDays int `json:"hadSignDays"`
			} `json:"data"`
		}
		if e := json.Unmarshal(req.Respon,&msg);e != nil{
			apilog.L(`E: `,e)
		}
		if msg.Code == 0 {apilog.L(`I: `,`签到成功!本月已签到`, msg.Data.HadSignDays,`天`);return}
		apilog.L(`E: `,msg.Message)
	}
}

//LIVE_BUVID
func (i *api) Get_LIVE_BUVID() (o *api){
	o = i
	apilog := apilog.Base_add(`LIVE_BUVID`).L(`T: `,`获取LIVE_BUVID`)
	if len(c.Cookie) == 0 {apilog.L(`E: `,`失败！无cookie`);return}
	if api_limit.TO() {apilog.L(`E: `,`超时！`);return}//超额请求阻塞，超时将取消

	{//获取
		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:`https://live.bilibili.com/`+ strconv.Itoa(o.Roomid),
			Header:map[string]string{
				`Host`: `api.live.bilibili.com`,
				`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
				`Accept`: `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`: `https://live.bilibili.com`,
				`Connection`: `keep-alive`,
				`Pragma`: `no-cache`,
				`Cache-Control`: `no-cache`,
				`Referer`:"https://live.bilibili.com/all",
				`Upgrade-Insecure-Requests`: `1`,
				`Cookie`:p.Map_2_Cookies_String(c.Cookie),
			},
			Timeout:3,
			Retry:2,
		});err != nil {
			apilog.L(`E: `,err)
			return
		}

		//cookie
		for k,v := range p.Cookies_List_2_Map(req.Response.Cookies()){
			c.Cookie[k] = v
		}
	}
	return
}

//小心心
type E_json struct{
	Code int `json:"code"`
	Message string `json:"message"`
	Ttl int `json:"ttl"`
	Data struct{
		Timestamp int `json:"timestamp"`
		Heartbeat_interval int `json:"heartbeat_interval"`
		Secret_key string `json:"secret_key"`
		Secret_rule []int `json:"secret_rule"`
		Patch_status int `json:"patch_status"`
	} `json:"data"`
}

func (i *api) F_x25Kn() (o *api) {
	o = i
	apilog := apilog.Base_add(`小心心`).L(`T: `,`获取小心心`)
	if len(c.Cookie) == 0 {apilog.L(`E: `,`失败！无cookie`);return}
	if c.Cookie[`LIVE_BUVID`] == `` {apilog.L(`E: `,`失败！无LIVE_BUVID`);return}
	if o.Parent_area_id == -1 {apilog.L(`E: `,`失败！未获取Parent_area_id`);return}
	if o.Area_id == -1 {apilog.L(`E: `,`失败！未获取Area_id`);return}
	if api_limit.TO() {apilog.L(`E: `,`超时！`);return}//超额请求阻塞，超时将取消

	//查看今天小心心数量
	for _,v := range Gift_list() {
		if v.Gift_id == 30607 && v.Expire_at - int(p.Sys().GetSTime()) > 6 * 86400 {
			if v.Gift_num == 24 {
				apilog.L(`I: `,`今天小心心已满！`);return
			} else {
				apilog.L(`I: `,`今天已有`,v.Gift_num,`个小心心`)
			}
		}
	}

	var (
		res E_json
		loop_num = 0
	)

	csrf := c.Cookie[`bili_jct`]
	if csrf == `` {apilog.L(`E: `,"Cookie错误,无bili_jct");return}

	LIVE_BUVID := c.Cookie[`LIVE_BUVID`]
	if LIVE_BUVID == `` {apilog.L(`E: `,"Cookie错误,无LIVE_BUVID");return}

	var new_uuid string
	{
		if tmp_uuid,e := uuid.NewV4();e == nil {
			new_uuid = tmp_uuid.String()
		} else {
			apilog.L(`E: `,e)
			return
		}
	}

	{//初始化
		PostStr := `id=[`+strconv.Itoa(o.Parent_area_id)+`,`+strconv.Itoa(o.Area_id)+`,`+strconv.Itoa(loop_num)+`,`+strconv.Itoa(o.Roomid)+`]&`
		PostStr += `device=["`+LIVE_BUVID+`","`+new_uuid+`"]&`
		PostStr += `ts=`+strconv.Itoa(int(p.Sys().GetMTime()))
		PostStr += `&is_patch=0&`
		PostStr += `heart_beat=[]&`
		PostStr += `ua=Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0&`
		PostStr += `csrf_token=`+csrf+`&csrf=`+csrf+`&`
		PostStr += `visit_id=`

		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:`https://live-trace.bilibili.com/xlive/data-interface/v1/x25Kn/E`,
			Header:map[string]string{
				`Host`: `api.live.bilibili.com`,
				`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
				`Accept`: `application/json, text/plain, */*`,
				`Content-Type`: `application/x-www-form-urlencoded`,
				`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
				`Accept-Encoding`: `gzip, deflate, br`,
				`Origin`: `https://live.bilibili.com`,
				`Connection`: `keep-alive`,
				`Pragma`: `no-cache`,
				`Cache-Control`: `no-cache`,
				`Referer`:"https://live.bilibili.com/"+strconv.Itoa(o.Roomid),
				`Cookie`:p.Map_2_Cookies_String(c.Cookie),
			},
			PostStr:url.PathEscape(PostStr),
			Timeout:3,
			Retry:2,
		});err != nil {
			apilog.L(`E: `,err)
			return
		}

		if e := json.Unmarshal(req.Respon,&res);e != nil {
			apilog.L(`E: `,e)
			return
		}

		if res.Code != 0{
			apilog.L(`E: `,res.Message)
			return
		}
	}

	{//loop
		var cancle = make(chan struct{},1)
		//使用带tag的消息队列在功能间传递消息
		c.Danmu_Main_mq.Pull_tag(map[string]func(interface{})(bool){
			`change_room`:func(data interface{})(bool){//换房时退出当前房间
				close(cancle)
				return true
			},
		})
		
		defer apilog.L(`I: `,`退出`)

		for loop_num < 24*5 {
			select{
			case <- time.After(time.Second*time.Duration(res.Data.Heartbeat_interval)):
			case <- cancle:
				return
			}

			loop_num += 1
			
			var rt_obj = RT{
				R:R{
					Id:`[`+strconv.Itoa(o.Parent_area_id)+`,`+strconv.Itoa(o.Area_id)+`,`+strconv.Itoa(loop_num)+`,`+strconv.Itoa(o.Roomid)+`]`,
					Device:`["`+LIVE_BUVID+`","`+new_uuid+`"]`,
					Ets:res.Data.Timestamp,
					Benchmark:res.Data.Secret_key,
					Time:res.Data.Heartbeat_interval,
					Ts:int(p.Sys().GetMTime()),
				},
				T:res.Data.Secret_rule,
			}

			PostStr := `id=`+rt_obj.R.Id+`&`
			PostStr += `device=["`+LIVE_BUVID+`","`+new_uuid+`"]&`
			PostStr += `ets=`+strconv.Itoa(res.Data.Timestamp)
			PostStr += `&benchmark=`+res.Data.Secret_key
			PostStr += `&time=`+strconv.Itoa(res.Data.Heartbeat_interval)
			PostStr += `&ts=`+strconv.Itoa(rt_obj.R.Ts)
			PostStr += `&is_patch=0&`
			PostStr += `heart_beat=[]&`
			PostStr += `ua=Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0&`
			PostStr += `csrf_token=`+csrf+`&csrf=`+csrf+`&`
			PostStr += `visit_id=`
			
			if wasm := Wasm(rt_obj);wasm == `` {
				apilog.L(`E: `,`发生错误`)
				return
			} else {
				PostStr = `s=`+wasm+`&`+PostStr
			}
			
			req := p.Req()
			if err := req.Reqf(p.Rval{
				Url:`https://live-trace.bilibili.com/xlive/data-interface/v1/x25Kn/X`,
				Header:map[string]string{
					`Host`: `api.live.bilibili.com`,
					`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
					`Accept`: `application/json, text/plain, */*`,
					`Content-Type`: `application/x-www-form-urlencoded`,
					`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
					`Accept-Encoding`: `gzip, deflate, br`,
					`Origin`: `https://live.bilibili.com`,
					`Connection`: `keep-alive`,
					`Pragma`: `no-cache`,
					`Cache-Control`: `no-cache`,
					`Referer`:"https://live.bilibili.com/"+strconv.Itoa(o.Roomid),
					`Cookie`:p.Map_2_Cookies_String(c.Cookie),
				},
				PostStr:url.PathEscape(PostStr),
				Timeout:3,
				Retry:2,
			});err != nil {
				apilog.L(`E: `,err)
				return
			}

			if e := json.Unmarshal(req.Respon,&res);e != nil {
				apilog.L(`E: `,e)
				return
			}
	
			if res.Code != 0{
				apilog.L(`E: `,res.Message)
				return
			}

			//查看今天小心心数量
			if loop_num%5 == 0 {//每5min
				for _,v := range Gift_list() {
					if v.Gift_id == 30607 && v.Expire_at - int(p.Sys().GetSTime()) > 6 * 86400 {
						if v.Gift_num == 24 {
							apilog.L(`I: `,`今天小心心已满！`);return
						} else {
							apilog.L(`I: `,`获取到第`,v.Gift_num,`个小心心`)
						}
					}
				}
			}
		}
	}
	return
}

//礼物列表
type Gift_list_type struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data Gift_list_type_Data `json:"data"`
}

type Gift_list_type_Data struct {
	List []Gift_list_type_Data_List `json:"list"`
}

type Gift_list_type_Data_List struct{
	Gift_id int `json:"gift_id"`
	Gift_num int `json:"gift_num"`
	Expire_at int `json:"expire_at"`
}

func Gift_list() (list []Gift_list_type_Data_List) {
	apilog := apilog.Base_add(`小心心`).L(`T: `,`获取礼物列表`)
	if len(c.Cookie) == 0 {apilog.L(`E: `,`失败！无cookie`);return}
	if c.Roomid == 0 {apilog.L(`E: `,`失败！无Roomid`);return}
	if api_limit.TO() {apilog.L(`E: `,`超时！`);return}//超额请求阻塞，超时将取消

	req := p.Req()
	if err := req.Reqf(p.Rval{
		Url:`https://api.live.bilibili.com/xlive/web-room/v1/gift/bag_list?t=`+strconv.Itoa(int(p.Sys().GetMTime()))+`&room_id=`+strconv.Itoa(c.Roomid),
		Header:map[string]string{
			`Host`: `api.live.bilibili.com`,
			`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
			`Accept`: `application/json, text/plain, */*`,
			`Content-Type`: `application/x-www-form-urlencoded`,
			`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
			`Accept-Encoding`: `gzip, deflate, br`,
			`Origin`: `https://live.bilibili.com`,
			`Connection`: `keep-alive`,
			`Pragma`: `no-cache`,
			`Cache-Control`: `no-cache`,
			`Referer`:"https://live.bilibili.com/"+strconv.Itoa(c.Roomid),
			`Cookie`:p.Map_2_Cookies_String(c.Cookie),
		},
		Timeout:3,
		Retry:2,
	});err != nil {
		apilog.L(`E: `,err)
		return
	}

	var res Gift_list_type

	if e := json.Unmarshal(req.Respon,&res);e != nil {
		apilog.L(`E: `,e)
		return
	}

	if res.Code != 0{
		apilog.L(`E: `,res.Message)
		return
	}

	return res.Data.List
}
