package F

import (
	"time"
	"fmt"
	"os"
	"strconv"
	"strings"
    "context"
	"net/http"

	qr "github.com/skip2/go-qrcode"
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

var apilog = c.Log.Base(`api`)
var api_limit = p.Limit(1,2000,15000)//频率限制1次/2s，最大等待时间15s

func New_api(Roomid int) (o *api) {
	apilog.Base_add(`新建`).L(`T: `,"ok")
	o = new(api)
	o.Roomid = Roomid
	o.Get_info()

	return
}

func (i *api) Get_info() (o *api) {
	o = i
	apilog := apilog.Base_add(`获取房号`)
	defer apilog.L(`T: `,"ok")
	

	if o.Roomid == 0 {
		apilog.L(`E: `,"还未New_api")
		return
	}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消

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
	//排行
	if e := r.S(`"rank_desc":"`, `",`, 0, 0).Err;e == nil {
		c.Note = r.RS[0]
	}
	//roomid
	if tmp := r.S(`"room_id":`, `,`, 0, 0);tmp.Err != nil {
		// apilog.L(`E: `,"room_id", tmp.Err)
	} else if i,err := strconv.Atoi(tmp.RS[0]); err != nil{
		apilog.L(`E: `,"room_id", err)
	} else {
		o.Roomid = i
	}

	if o.Roomid != 0 && o.Uid != 0 && c.Title != ``{return}

	{//使用其他api
		req := p.Req()
		if err := req.Reqf(p.Rval{
			Url:"https://api.live.bilibili.com/xlive/web-room/v1/index/getInfoByRoom?room_id=" + Roomid,
			Header:map[string]string{
				`Referer`:"https://live.bilibili.com/" + Roomid,
				`Cookie`:c.Cookie,
			},
			Timeout:10,
			Retry:2,
		});err != nil {
			apilog.L(`E: `,err)
			return
		}
		res := string(req.Respon)
		if code := p.Json().GetValFrom(res, "code");code == nil || code.(float64) != 0 {
			apilog.L(`E: `,"code", code, p.Json().GetValFrom(res, "message"))
			return
		}
		//主播id
		if Uname,ok := p.Json().GetValFrom(res, "data.anchor_info.base_info.uname").(string);ok && c.Uname == `` {
			c.Uname = Uname
		}
		//排行
		if rank_desc,ok := p.Json().GetValFrom(res, "data.rankdb_info.rank_desc").(string);ok && c.Note == `` {//有时会返回`小时总榜`
			c.Note = rank_desc
		}
		if Uid := p.Json().GetValFrom(res, "data.room_info.uid");Uid == nil {
			apilog.L(`E: `,"data.room_info.uid", Uid)
			return
		} else {
			o.Uid = int(Uid.(float64))
		}

		if room_id := p.Json().GetValFrom(res, "data.room_info.room_id");room_id == nil {
			apilog.L(`E: `,"data.room_info.room_id", room_id)
			return
		} else {
			o.Roomid = int(room_id.(float64))
		}
		if title := p.Json().GetValFrom(res, "data.room_info.title");title == nil {
			apilog.L(`E: `,"data.room_info.title", title)
			return
		} else {
			c.Title = title.(string)
		}
		if is_locked := p.Json().GetValFrom(res, "data.room_info.lock_status");is_locked == nil {
			apilog.L(`E: `,"data.room_info.is_locked", is_locked)
			return
		} else if is_locked.(float64) == 1 {
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
	defer apilog.L(`T: `,"ok")

	if o.Roomid == 0 {
		apilog.L(`E: `,"还未New_api")
		return
	}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消

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
			`Cookie`:c.Cookie,
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
			`Cookie`:c.Cookie,
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
				`Cookie`:c.Cookie,
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

func (i *api) Get_guardNum() {
	if i.Uid == 0 || c.Roomid == 0 {
		apilog.Base_add("Get_guardNum").L(`E: `,"i.Uid == 0 || c.Roomid == 0")
		return
	}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消
	apilog := apilog.Base_add(`获取舰长数`)
	defer apilog.L(`T: `,"ok")

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
			`Cookie`:c.Cookie,
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
	defer apilog.L(`T: `,"ok")

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
			if strings.Contains(tmp,`player-loader`) {
				player_js_url = tmp
				break
			}
		}
		if player_js_url == `` {
			apilog.L(`E: `,`no found player-loader js`)
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

		r.S(`version={html5:{web:"`,`"`,0,0)
		if r.Err != nil {
			apilog.L(`E: `,r.Err)
			return
		}
		c.VERSION = r.RS[0]
		apilog.L(`W: `,"api version", c.VERSION)
	}
}

func Get_cookie() {
	if api_limit.TO() {return}//超额请求阻塞，超时将取消
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
					`Cookie`:c.Cookie,
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
					apilog.L(`W: `,`登录中`,v)
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
		c.Cookie = cookie
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
	if c.Cookie == `` {return}
	if api_limit.TO() {return}//超额请求阻塞，超时将取消
	apilog := apilog.Base_add(`切换粉丝牌`)
	defer apilog.L(`T: `,"ok")

	{//验证是否本直播间牌子
		r := p.Req()
		if e := r.Reqf(p.Rval{
			Url:`https://api.live.bilibili.com/live_user/v1/UserInfo/get_weared_medal`,
			Header:map[string]string{
				`Cookie`:c.Cookie,
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
				`Cookie`:c.Cookie,
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
		var csrf string
		if i := strings.Index(c.Cookie, "bili_jct="); i == -1 {
			apilog.L(`E: `,"Cookie错误,无bili_jct=")
			return
		} else {
			if d := strings.Index(c.Cookie[i + 9:], ";"); d == -1 {
				csrf = c.Cookie[i + 9:]
			} else {
				csrf = c.Cookie[i + 9:][:d]
			}
		}
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
				`Cookie`:c.Cookie,
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