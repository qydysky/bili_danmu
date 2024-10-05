package cv

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	crypto "github.com/qydysky/part/crypto"
	pctx "github.com/qydysky/part/ctx"
	file "github.com/qydysky/part/file"
	log "github.com/qydysky/part/log"
	mq "github.com/qydysky/part/msgq"
	pool "github.com/qydysky/part/pool"
	reqf "github.com/qydysky/part/reqf"
	psql "github.com/qydysky/part/sql"
	syncmap "github.com/qydysky/part/sync"
	sys "github.com/qydysky/part/sys"
	web "github.com/qydysky/part/web"
	_ "modernc.org/sqlite"
)

//go:embed VERSION
var version string

type StreamType struct {
	Protocol_name string
	Format_name   string
	Codec_name    string
}

type Common struct {
	InIdle            bool           `json:"-"`                //闲置中？
	PID               int            `json:"-"`                //进程id
	Version           string         `json:"-"`                //版本
	Uid               int            `json:"-"`                //client uid
	Live              []*LiveQn      `json:"live"`             //直播流链接
	Live_qn           int            `json:"liveQn"`           //当前直播流质量
	Live_want_qn      int            `json:"-"`                //期望直播流质量
	Roomid            int            `json:"-"`                //房间ID
	Cookie            syncmap.Map    `json:"-"`                //Cookie
	Title             string         `json:"title"`            //直播标题
	Uname             string         `json:"uname"`            //主播名
	UpUid             int            `json:"upUid"`            //主播uid
	Rev               float64        `json:"rev"`              //营收
	Renqi             int            `json:"renqi"`            //人气
	Watched           int            `json:"watched"`          //观看人数
	OnlineNum         int            `json:"onlineNum"`        //在线人数
	GuardNum          int            `json:"guardNum"`         //舰长数
	ParentAreaID      int            `json:"parentAreaID"`     //父分区
	AreaID            int            `json:"areaID"`           //子分区
	Locked            bool           `json:"locked"`           //直播间封禁
	Note              string         `json:"note"`             //分区排行
	Live_Start_Time   time.Time      `json:"-"`                //直播开始时间
	Liveing           bool           `json:"liveing"`          //是否在直播
	Wearing_FansMedal int            `json:"WearingFansMedal"` //当前佩戴的粉丝牌
	Token             string         `json:"-"`                //弹幕钥
	WSURL             []string       `json:"-"`                //弹幕链接
	LiveBuvidUpdated  time.Time      `json:"-"`                //LIVE_BUVID更新时间
	Stream_url        *url.URL       `json:"-"`                //直播Web服务
	Proxy             string         `json:"-"`                //全局代理
	SerLocation       int            `json:"-"`                //服务器时区
	AcceptQn          map[int]string `json:"-"`                //允许的直播流质量
	Qn                map[int]string `json:"-"`                //全部直播流质量
	// StreamType        StreamType            `json:"streamType"`    //当前直播流类型
	AllStreamType map[string]StreamType            `json:"-"` //直播流类型
	K_v           syncmap.Map                      `json:"-"` //配置文件
	Log           *log.Log_interface               `json:"-"` //日志
	Danmu_Main_mq *mq.Msgq                         `json:"-"` //消息
	ReqPool       *pool.Buf[reqf.Req]              `json:"-"` //请求池
	SerF          *web.WebPath                     `json:"-"` //web服务处理
	SerLimit      *web.Limits                      `json:"-"` //Web服务连接限制
	StartT        time.Time                        `json:"-"` //启动时间
	Cache         syncmap.MapExceeded[string, any] `json:"-"` //缓存
}

func (t *Common) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Live          []*LiveQn
		LiveQn        int
		Title         string
		Uname         string
		UpUid         int
		Rev           float64
		Watched       int
		OnlineNum     int
		GuardNum      int
		ParentAreaID  int
		AreaID        int
		Locked        bool
		Note          string
		LiveStartTime string
		Liveing       bool
	}{
		Live:          append([]*LiveQn{}, t.Live...),
		LiveQn:        t.Live_qn,
		Title:         t.Title,
		Uname:         t.Uname,
		UpUid:         t.UpUid,
		Rev:           t.Rev,
		Watched:       t.Watched,
		OnlineNum:     t.OnlineNum,
		GuardNum:      t.GuardNum,
		ParentAreaID:  t.ParentAreaID,
		AreaID:        t.AreaID,
		Locked:        t.Locked,
		Note:          t.Note,
		LiveStartTime: t.Live_Start_Time.Format(time.DateTime),
		Liveing:       t.Liveing,
	})
}

type LiveQn struct {
	Url          string `json:"-"`
	Uuid         string `json:"-"`
	Codec        string
	ReUpTime     time.Time
	CreateTime   time.Time
	DisableCount int
	Expires      time.Time //流到期时间
}

func (t LiveQn) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Host         string
		Up           bool
		Codec        string
		CreateTime   string
		ReUpTime     string
		Expires      string
		DisableCount int
	}{
		Host:         t.Host(),
		Up:           time.Now().After(t.ReUpTime),
		Codec:        t.Codec,
		CreateTime:   t.CreateTime.Format(time.DateTime),
		ReUpTime:     t.ReUpTime.Format(time.DateTime),
		Expires:      t.Expires.Format(time.DateTime),
		DisableCount: t.DisableCount,
	})
}

func (t *LiveQn) SetUrl(url string) {
	t.Url = url
}

func (t *LiveQn) Host() string {
	if liveUrl, e := url.Parse(t.Url); e == nil {
		return liveUrl.Host
	} else {
		panic(e)
	}
}

func (t *LiveQn) Valid() bool {
	return time.Now().After(t.ReUpTime)
}

// 自动停用机制
func (t *LiveQn) DisableAuto() (hadDisable bool) {
	if !t.Valid() {
		return true
	}
	if time.Now().After(t.ReUpTime.Add(time.Minute).Add(time.Second * time.Duration(20*t.DisableCount))) {
		t.DisableCount = 0
	}
	t.DisableCount += 1
	t.ReUpTime = time.Now().Add(time.Minute).Add(time.Second * time.Duration(20*t.DisableCount))
	return
}

func (t *LiveQn) Disable(reUpTime time.Time) {
	t.ReUpTime = reUpTime
}

func (t *Common) IsOn(key string) bool {
	v, ok := t.K_v.LoadV(key).(bool)
	return ok && v
}

func (t *Common) Copy() *Common {
	var c = Common{
		InIdle:            t.InIdle,
		PID:               t.PID,
		Version:           t.Version,
		Uid:               t.Uid,
		Live:              t.Live,
		Live_qn:           t.Live_qn,
		Live_want_qn:      t.Live_want_qn,
		Roomid:            t.Roomid,
		Cookie:            t.Cookie.Copy(),
		Title:             t.Title,
		Uname:             t.Uname,
		UpUid:             t.UpUid,
		Rev:               t.Rev,
		Renqi:             t.Renqi,
		Watched:           t.Watched,
		OnlineNum:         t.OnlineNum,
		GuardNum:          t.GuardNum,
		ParentAreaID:      t.ParentAreaID,
		AreaID:            t.AreaID,
		Locked:            t.Locked,
		Note:              t.Note,
		Live_Start_Time:   t.Live_Start_Time,
		Liveing:           t.Liveing,
		Wearing_FansMedal: t.Wearing_FansMedal,
		Token:             t.Token,
		WSURL:             t.WSURL,
		LiveBuvidUpdated:  t.LiveBuvidUpdated,
		Stream_url:        t.Stream_url,
		Proxy:             t.Proxy,
		SerLocation:       t.SerLocation,
		AcceptQn:          syncmap.Copy(t.AcceptQn),
		Qn:                syncmap.Copy(t.Qn),
		// StreamType:        t.StreamType,
		AllStreamType: syncmap.Copy(t.AllStreamType),
		K_v:           t.K_v.Copy(),
		Log:           t.Log,
		Danmu_Main_mq: t.Danmu_Main_mq,
		ReqPool:       t.ReqPool,
		SerF:          t.SerF,
		SerLimit:      t.SerLimit,
		StartT:        t.StartT,
		Cache:         *t.Cache.Copy(),
	}

	return &c
}

// 自动停用机制
func (t *Common) DisableLiveAutoByUuid(uuid string) (hadDisable bool) {
	for i := 0; i < len(t.Live); i++ {
		if t.Live[i].Uuid == uuid {
			return t.Live[i].DisableAuto()
		}
	}
	return
}

// Deprecated: 存在缺陷
func (t *Common) DisableLiveAuto(host string) (hadDisable bool) {
	for i := 0; i < len(t.Live); i++ {
		if liveUrl, e := url.Parse(t.Live[i].Url); e == nil {
			if host == liveUrl.Host {
				return t.Live[i].DisableAuto()
			}
		}
	}
	return
}

// Deprecated: 存在缺陷
func (t *Common) DisableLive(host string, reUpTime time.Time) {
	for i := 0; i < len(t.Live); i++ {
		if liveUrl, e := url.Parse(t.Live[i].Url); e == nil {
			if host == liveUrl.Host {
				t.Live[i].ReUpTime = reUpTime
				break
			}
		}
	}
}

func (t *Common) ValidNum() (num int) {
	for i := 0; i < len(t.Live); i++ {
		if time.Now().After(t.Live[i].ReUpTime) {
			num += 1
		}
	}
	return
}

func (t *Common) ValidLive() *LiveQn {
	for i := 0; i < len(t.Live); i++ {
		if time.Now().Before(t.Live[i].ReUpTime) {
			continue
		}
		return t.Live[i]
	}
	return nil
}

func (t *Common) Init() *Common {
	t.PID = os.Getpid()
	t.Version = strings.TrimSpace(version)
	t.StartT = time.Now()

	t.AllStreamType = map[string]StreamType{
		`fmp4`: {
			Protocol_name: "http_hls",
			Format_name:   "fmp4",
			Codec_name:    "avc",
		},
		`flv`: {
			Protocol_name: "http_stream",
			Format_name:   "flv",
			Codec_name:    "avc",
		},
		`fmp4H`: {
			Protocol_name: "http_hls",
			Format_name:   "fmp4",
			Codec_name:    "hevc",
		},
		`flvH`: {
			Protocol_name: "http_stream",
			Format_name:   "flv",
			Codec_name:    "hevc",
		},
	}

	t.Qn = map[int]string{ // no change
		20000: "4K",
		10000: "原画",
		400:   "蓝光",
		250:   "超清",
		150:   "高清",
		80:    "流畅",
	}

	t.AcceptQn = map[int]string{ // no change
		20000: "4K",
		10000: "原画",
		400:   "蓝光",
		250:   "超清",
		150:   "高清",
		80:    "流畅",
	}

	t.Danmu_Main_mq = mq.New(time.Second*5, time.Second*10)

	go func() { //日期变化
		var old = time.Now().Hour()
		for {
			if now := time.Now().Hour(); now == 0 && old != now {
				t.Danmu_Main_mq.Push_tag(`new day`, nil)
				old = now
			}
			t.Danmu_Main_mq.Push_tag(`every100s`, nil)
			time.Sleep(time.Second * time.Duration(100))
		}
	}()
	t.ReqPool = pool.New(
		pool.PoolFunc[reqf.Req]{
			New: func() *reqf.Req {
				return reqf.New()
			},
			InUse: func(r *reqf.Req) bool {
				return r.IsLive()
			},
			Reuse: func(r *reqf.Req) *reqf.Req {
				return r
			},
			Pool: func(r *reqf.Req) *reqf.Req {
				return r
			},
		},
		100,
	)

	var (
		ckv     = flag.String("ckv", "", "自定义配置KV文件，将会覆盖config_K_v配置")
		roomIdP = flag.Int("r", 0, "roomid")
		genKey  = flag.Bool("genKey", false, "生成cookie加密公私钥")
	)
	testing.Init()
	flag.Parse()

	if *genKey {
		if pri, pub, e := crypto.NewKey(); e != nil {
			panic(e)
		} else {
			fmt.Println("公钥：")
			fmt.Println(string(pub))
			fmt.Println("私钥：")
			fmt.Println(string(pri))
			fmt.Println("请复制以上公私钥并另存为文件,可以在cookie加密公钥、cookie解密私钥中使用")
			os.Exit(0)
		}
	}

	t.Roomid = *roomIdP

	if e := t.loadConf(*ckv); e != nil {
		if os.IsNotExist(e) {
			fmt.Println("未能加载配置文件")
		} else {
			panic(e)
		}
	}

	go func() {
		for {
			v, ok := t.K_v.LoadV("几秒后重载").(float64)
			if !ok || v < 0 {
				break
			} else if v < 60 {
				v = 60
			}
			time.Sleep(time.Duration(int(v)) * time.Second)

			if e := t.loadConf(*ckv); e != nil {
				fmt.Println(e)
			}
		}
	}()

	t.SerF = new(web.WebPath)
	t.Stream_url = &url.URL{}
	t.SerLimit = &web.Limits{}

	if serAdress, ok := t.K_v.LoadV("Web服务地址").(string); ok && serAdress != "" {
		serUrl, e := url.Parse("http://" + serAdress)
		if e != nil {
			panic(e)
		}

		port := `80`
		if s := strings.Split(serAdress, ":"); len(s) > 1 {
			port = s[1]
		}

		{
			// 启动时显示ip
			showIpOnBoot, _ := t.K_v.LoadV("启动时显示ip").(bool)
			for ip := range sys.GetIpByCidr() {
				if showIpOnBoot {
					if ip.To4() != nil {
						fmt.Printf("当前地址 http://%s:%s\n", ip.String(), port)
					} else {
						fmt.Printf("当前地址 http://[%s]:%s\n", ip.String(), port)
					}
				}
			}
		}

		var (
			readTimeout       = 3
			readHeaderTimeout = 3
			idleTimeout       = 3
		)

		if v, ok := t.K_v.LoadV("Web服务超时配置").(map[string]any); ok {
			if v1, ok := v["ReadTimeout"].(float64); ok && v1 > 3 {
				readTimeout = int(v1)
			}
			if v1, ok := v["ReadHeaderTimeout"].(float64); ok && v1 > 3 {
				readHeaderTimeout = int(v1)
			}
			if v1, ok := v["IdleTimeout"].(float64); ok && v1 > 3 {
				idleTimeout = int(v1)
			}
		}

		web.NewSyncMap(&http.Server{
			Addr:              serUrl.Host,
			ReadTimeout:       time.Duration(int(time.Second) * readTimeout),
			ReadHeaderTimeout: time.Duration(int(time.Second) * readHeaderTimeout),
			IdleTimeout:       time.Duration(int(time.Second) * idleTimeout),
		}, t.SerF, t.SerF.LoadPerfix)

		if limits, ok := t.K_v.LoadV(`Web服务连接限制`).([]any); ok {
			for i := 0; i < len(limits); i++ {
				if vm, ok := limits[i].(map[string]any); ok {
					if cidr, ok := vm["cidr"].(string); !ok {
						continue
					} else if max, ok := vm["max"].(float64); !ok {
						continue
					} else {
						t.SerLimit.AddLimitItem(web.NewLimitItem(int(max)).Cidr(cidr))
					}
				}
			}
		}

		if val, ok := t.K_v.LoadV("ip路径").(string); ok && val != "" {
			t.SerF.Store(val, func(w http.ResponseWriter, r *http.Request) {
				if DefaultHttpFunc(t, w, r, http.MethodGet) {
					return
				}
				for ip := range sys.GetIpByCidr() {
					if _, e := w.Write([]byte(ip.String() + "\n")); e != nil {
						return
					}
				}
			})
		}

		if val, ok := t.K_v.LoadV("性能路径").(string); ok && val != "" {
			var cache web.Cache
			t.SerF.Store(val, func(w http.ResponseWriter, r *http.Request) {
				if DefaultHttpFunc(t, w, r, http.MethodGet) {
					return
				}

				//cache
				if bp, ok := cache.IsCache(val); ok {
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Cache-Control", "max-age=5")
					_, _ = w.Write(*bp)
					return
				}
				w = cache.Cache(val, time.Second*5, w)

				var memStats runtime.MemStats
				runtime.ReadMemStats(&memStats)

				var gcAvgS float64

				if memStats.NumGC != 0 {
					gcAvgS = time.Since(t.StartT).Seconds() / float64(memStats.NumGC)
				}

				streams := make(map[int]any)

				Commons.Range(func(key, value any) bool {
					if common, ok := value.(*Common); ok && common.InIdle {
						return true
					}
					streams[key.(int)] = value
					return true
				})

				reqState := t.ReqPool.State()

				ResStruct{0, "ok", map[string]any{
					"version":     t.Version,
					"startTime":   t.StartT.Format(time.DateTime),
					"currentTime": time.Now().Format(time.DateTime),
					"state": map[string]any{
						"base": map[string]any{
							"reqPoolState": map[string]any{
								"pooled":   reqState.Pooled,
								"nopooled": reqState.Nopooled,
								"inuse":    reqState.Inuse,
								"nouse":    reqState.Nouse,
								"sum":      reqState.Sum,
								"qts":      math.Round(reqState.GetPerSec*100) / 100,
							},
							"pid":          t.PID,
							"numGoroutine": runtime.NumGoroutine(),
							"goVersion":    runtime.Version(),
						},
						"mem": map[string]any{
							"memInUse": humanize.Bytes(memStats.HeapInuse + memStats.StackInuse),
						},
						"gc": map[string]any{
							"numGC":            memStats.NumGC,
							"lastGC":           time.UnixMicro(int64(memStats.LastGC / 1000)).Format(time.DateTime),
							"gcCPUFractionPpm": float64(int(memStats.GCCPUFraction*100000000)) / 100,
							"gcAvgS":           float64(int(gcAvgS*100)) / 100,
						},
						"common": streams,
					},
				},
				}.Write(w)
			})
		}

		t.Stream_url, _ = url.Parse(`http://` + serAdress)
	}

	if val, exist := t.K_v.Load("代理地址"); exist {
		t.Proxy = val.(string)
	}

	if val, exist := t.K_v.LoadV("服务器时区").(float64); exist && val != 0 {
		t.SerLocation = int(val)
	}

	// 配置直播流类型
	if val, exist := t.K_v.Load("直播流类型"); exist {
		if _, ok := t.AllStreamType[val.(string)]; !ok {
			panic("未找到设定类型" + val.(string))
		}
	}

	{
		v, _ := t.K_v.LoadV("日志文件输出").(string)
		t.Log = log.New(log.Config{
			File:   v,
			Stdout: true,
			Prefix_string: map[string]struct{}{
				`T: `: log.On,
				`I: `: log.On,
				`N: `: log.On,
				`W: `: log.On,
				`E: `: log.On,
			},
		})

		if v, ok := t.K_v.LoadV(`保存日志至db`).(map[string]any); ok && len(v) != 0 {
			dbname, dbnameok := v["dbname"].(string)
			url, urlok := v["url"].(string)
			create, createok := v["create"].(string)
			insert, insertok := v["insert"].(string)
			if dbnameok && urlok && insertok && dbname != "" && url != "" && insert != "" {
				db, e := sql.Open(dbname, url)
				if e != nil {
					panic("保存日志至db打开连接错误 " + e.Error())
				}
				if createok {
					tx := psql.BeginTx[any](db, pctx.GenTOCtx(time.Second*5))
					tx.Do(psql.SqlFunc[any]{
						Query:      create,
						SkipSqlErr: true,
					})
					if _, e := tx.Fin(); e != nil {
						panic("保存日志至db打开连接错误 " + e.Error())
					}
				}
				t.Log = t.Log.LDB(db, insert, time.Second*5)
			}
		}

		logmap := make(map[string]struct{})
		if array, ok := t.K_v.Load(`日志显示`); ok {
			for _, v := range array.([]interface{}) {
				logmap[v.(string)] = log.On
			}
		}
		t.Log = t.Log.Level(logmap)
	}

	return t
}

func (t *Common) loadConf(customConf string) error {
	var data map[string]interface{}

	// 64k
	f := file.New("config/config_K_v.json", 0, true)
	if !f.IsExist() {
		return os.ErrNotExist
	}
	if bb, e := f.ReadAll(100, 1<<16); e != nil {
		if !errors.Is(e, io.EOF) {
			return e
		} else {
			_ = json.Unmarshal(bb, &data)
		}
	}

	if customConf != "" {
		if strings.Contains(customConf, "http:") || strings.Contains(customConf, "https:") {
			//从网址读取
			req := t.ReqPool.Get()
			defer t.ReqPool.Put(req)
			if e := req.Reqf(reqf.Rval{
				Url: customConf,
				Header: map[string]string{
					`User-Agent`:      UA,
					`Accept`:          `*/*`,
					`Accept-Language`: `zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2`,
					`Accept-Encoding`: `gzip, deflate, br`,
					`Pragma`:          `no-cache`,
					`Cache-Control`:   `no-cache`,
					`Connection`:      `close`,
				},
				Timeout: 10 * 1000,
			}); e != nil {
				return errors.New("无法获取自定义配置文件 " + e.Error())
			}
			if req.Response == nil {
				return errors.New("无法获取自定义配置文件 响应为空")
			} else if req.Response.StatusCode&200 != 200 {
				return fmt.Errorf("无法获取自定义配置文件 %d", req.Response.StatusCode)
			} else {
				var tmp map[string]interface{}
				_ = json.Unmarshal(req.Respon, &tmp)
				for k, v := range tmp {
					data[k] = v
				}
			}
		} else {
			//从文件读取
			if bb, err := file.New(customConf, 0, true).ReadAll(100, 1<<16); err != nil {
				if errors.Is(err, io.EOF) {
					var tmp map[string]interface{}
					_ = json.Unmarshal(bb, &tmp)
					for k, v := range tmp {
						data[k] = v
					}
				} else {
					return err
				}
			}
		}
	}

	for k, v := range data {
		t.K_v.Store(k, v)
	}

	return nil
}

var C = new(Common).Init()

// StreamRec
// fmp4
// https://datatracker.ietf.org/doc/html/draft-pantos-http-live-streaming
var StreamO = new(sync.Map)
var Commons = new(syncmap.Map)
var CommonsLoadOrStore = syncmap.LoadOrStoreFunc[Common]{
	Init: func() *Common {
		return C.Copy()
	},
}

// 消息队列
type Danmu_Main_mq_item struct {
	Class string
	Data  interface{}
}

// Web服务响应格式
type ResStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (t ResStruct) Write(w http.ResponseWriter) []byte {
	w.Header().Set("Content-Type", "application/json")
	data, e := json.MarshalIndent(t, "", "    ")
	if e != nil {
		t.Code = -1
		t.Data = nil
		t.Message = e.Error()
		data, _ = json.Marshal(t)
	}
	_, _ = w.Write(data)
	return data
}

func DefaultHttpFunc(c *Common, w http.ResponseWriter, r *http.Request, method ...string) bool {
	if strings.Contains(r.URL.Path, "../") {
		web.WithStatusCode(w, http.StatusForbidden)
		return true
	}
	//method
	if !web.IsMethod(r, method...) {
		web.WithStatusCode(w, http.StatusMethodNotAllowed)
		return true
	}
	//limit
	if c.SerLimit.AddCount(r) {
		web.WithStatusCode(w, http.StatusTooManyRequests)
		return true
	}
	//Web自定义响应头
	if resHeaders, ok := c.K_v.LoadV("Web自定义响应头").(map[string]any); ok && len(resHeaders) > 0 {
		for k, v := range resHeaders {
			if k == "" {
				continue
			} else if vs, ok := v.(string); !ok || vs == "" {
				continue
			} else {
				w.Header().Set(k, vs)
			}
		}
	}
	return false
}
