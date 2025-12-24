package cv

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io"
	"maps"
	"math"
	"net"
	"net/http"
	"net/http/pprof"
	"net/url"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
	_ "github.com/go-sql-driver/mysql" //removable
	_ "github.com/jackc/pgx/v5/stdlib" //removable
	pca "github.com/qydysky/part/crypto/asymmetric"
	pctx "github.com/qydysky/part/ctx"
	file "github.com/qydysky/part/file"
	pio "github.com/qydysky/part/io"
	plog "github.com/qydysky/part/log/v2"
	mq "github.com/qydysky/part/msgq"
	pool "github.com/qydysky/part/pool"
	reqf "github.com/qydysky/part/reqf"
	psql "github.com/qydysky/part/sql"
	syncmap "github.com/qydysky/part/sync"
	sys "github.com/qydysky/part/sys"
	web "github.com/qydysky/part/web"
	_ "modernc.org/sqlite" //removable
)

//go:embed VERSION
var version string

type StreamType struct {
	Protocol_name string
	Format_name   string
	Codec_name    string
}

type Common struct {
	InIdle            bool                             `json:"-"`            //闲置中？
	PID               int                              `json:"-"`            //进程id
	Version           string                           `json:"-"`            //版本
	Uid               int                              `json:"-"`            //client uid
	Login             bool                             `json:"login"`        //登录
	Live              []*LiveQn                        `json:"live"`         //直播流链接
	Live_qn           int                              `json:"liveQn"`       //当前直播流质量
	Live_want_qn      int                              `json:"-"`            //期望直播流质量
	Roomid            int                              `json:"-"`            //房间ID
	Cookie            *syncmap.Map                     `json:"-"`            //Cookie
	Title             string                           `json:"title"`        //直播标题
	Uname             string                           `json:"uname"`        //主播名
	UpUid             int                              `json:"upUid"`        //主播uid
	Rev               float64                          `json:"rev"`          //营收
	Renqi             int                              `json:"renqi"`        //人气
	Watched           int                              `json:"watched"`      //观看人数
	OnlineNum         int                              `json:"onlineNum"`    //在线人数
	GuardNum          int                              `json:"guardNum"`     //舰长数
	ParentAreaID      int                              `json:"parentAreaID"` //父分区
	AreaID            int                              `json:"areaID"`       //子分区
	Locked            bool                             `json:"locked"`       //直播间封禁
	Note              string                           `json:"note"`         //分区排行
	Live_Start_Time   time.Time                        `json:"-"`            //直播开始时间
	Liveing           bool                             `json:"liveing"`      //是否在直播
	Wearing_FansMedal int                              `json:"-"`            //当前佩戴的粉丝牌
	Token             string                           `json:"-"`            //弹幕钥
	WSURL             []string                         `json:"-"`            //弹幕链接
	LiveBuvidUpdated  time.Time                        `json:"-"`            //LIVE_BUVID更新时间
	Stream_url        *url.URL                         `json:"-"`            //直播Web服务
	Proxy             string                           `json:"-"`            //全局代理
	SerLocation       int                              `json:"-"`            //服务器时区
	AcceptQn          map[int]string                   `json:"-"`            //允许的直播流质量
	Qn                map[int]string                   `json:"-"`            //全部直播流质量
	AllStreamType     map[string]StreamType            `json:"-"`            //直播流类型
	K_v               syncmap.Map                      `json:"-"`            //配置文件
	Log               *plog.Log                        `json:"-"`            //日志
	Danmu_Main_mq     *mq.Msgq                         `json:"-"`            //消息
	ReqPool           *pool.Buf[reqf.Req]              `json:"-"`            //请求池
	SerF              *web.WebPath                     `json:"-"`            //web服务处理
	SerLimit          *web.Limits                      `json:"-"`            //Web服务连接限制
	StartT            time.Time                        `json:"-"`            //启动时间
	Cache             syncmap.MapExceeded[string, any] `json:"-"`            //缓存
	l                 sync.RWMutex                     `json:"-"`
	buf               []byte                           `json:"-"`
}

func (t *Common) Lock() func() {
	t.l.Lock()
	return t.l.Unlock
}

func (t *Common) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Live          []*LiveQn `json:"live"`
		LiveQn        int       `json:"liveQn"`
		Title         string    `json:"title"`
		Uname         string    `json:"uname"`
		UpUid         int       `json:"upUid"`
		Rev           float64   `json:"rev"`
		Watched       int       `json:"watched"`
		OnlineNum     int       `json:"onlineNum"`
		GuardNum      int       `json:"guardNum"`
		ParentAreaID  int       `json:"parentAreaID"`
		AreaID        int       `json:"areaID"`
		Locked        bool      `json:"locked"`
		Login         bool      `json:"login"`
		Note          string    `json:"note"`
		LiveStartTime string    `json:"liveStartTime"`
		Liveing       bool      `json:"liveing"`
	}{
		Live:          append([]*LiveQn{}, t.Live...),
		LiveQn:        t.Live_qn,
		Title:         t.Title,
		Uname:         t.Uname,
		UpUid:         t.UpUid,
		Rev:           math.Round(t.Rev*100) / 100,
		Watched:       t.Watched,
		OnlineNum:     t.OnlineNum,
		GuardNum:      t.GuardNum,
		ParentAreaID:  t.ParentAreaID,
		AreaID:        t.AreaID,
		Locked:        t.Locked,
		Login:         t.Login,
		Note:          t.Note,
		LiveStartTime: t.Live_Start_Time.Format(time.RFC3339),
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
		Host         string `json:"host"`
		Up           bool   `json:"up"`
		Codec        string `json:"codec"`
		CreateTime   string `json:"createTime"`
		ReUpTime     string `json:"reUpTime"`
		Expires      string `json:"expires"`
		DisableCount int    `json:"disableCount"`
	}{
		Host:         t.Host(),
		Up:           time.Now().After(t.ReUpTime),
		Codec:        t.Codec,
		CreateTime:   t.CreateTime.Format(time.RFC3339),
		ReUpTime:     t.ReUpTime.Format(time.RFC3339),
		Expires:      t.Expires.Format(time.RFC3339),
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

func (t *Common) GenReqCookie() string {
	return reqf.Iter_2_Cookies_String(func(yield func(string, string) bool) {
		t.Cookie.Range(func(k, v any) bool {
			yield(k.(string), v.(string))
			return true
		})
	})
}

func (t *Common) IsLogin() bool {
	for _, n := range []string{`bili_jct`, `DedeUserID`} {
		if _, ok := t.Cookie.Load(n); !ok {
			return false
		}
	}
	return true
}

func (t *Common) IsOn(key string) bool {
	v, ok := t.K_v.LoadV(key).(bool)
	return ok && v
}

func (t *Common) Copy() *Common {
	var c = Common{
		Login:             t.Login,
		InIdle:            t.InIdle,
		PID:               t.PID,
		Version:           t.Version,
		Uid:               t.Uid,
		Live:              t.Live,
		Live_qn:           t.Live_qn,
		Live_want_qn:      t.Live_want_qn,
		Roomid:            t.Roomid,
		Cookie:            t.Cookie,
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
	t.Cookie = &syncmap.Map{}

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
		`fmp4A`: {
			Protocol_name: "http_hls",
			Format_name:   "fmp4",
			Codec_name:    "av1",
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

	t.Danmu_Main_mq = mq.New(time.Second*5, time.Second*90)

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
		stop    = flag.Bool("stop", false, "向当前配置发送退出信号")
	)

	// 支持test命令
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-test.") {
			testing.Init()
			break
		}
	}

	flag.Parse()

	if *genKey {
		if pri, pub, e := pca.MlkemF.NewKey(); e != nil {
			panic(e)
		} else {
			fmt.Println("公钥：")
			fmt.Println(string(pem.EncodeToMemory(pub)))
			fmt.Println("私钥：")
			fmt.Println(string(pem.EncodeToMemory(pri)))
			fmt.Println("请复制以上公私钥并另存为文件,可以在cookie加密公钥、cookie解密私钥中使用")
			os.Exit(0)
		}
	}

	// load from env
	if tmp := os.Getenv("ckv"); *ckv == "" && tmp != "" {
		fmt.Println("加载环境变量ckv:", tmp)
		*ckv = tmp
	}
	if tmp := os.Getenv("r"); *roomIdP == 0 && tmp != "" {
		fmt.Println("加载环境变量r:", tmp)
		if r, e := strconv.Atoi(tmp); e == nil {
			*roomIdP = r
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
			r := reqf.New()
			showIpOnBoot, _ := t.K_v.LoadV("启动时显示ip").(bool)
			stopPath, _ := t.K_v.LoadV("stop路径").(string)
			waitStop, _ := t.K_v.LoadV("停止其他服务超时").(float64)
			if waitStop <= 0 {
				waitStop = 100
			}
			for ip := range sys.GetIpByCidr() {
				if ip.IsLinkLocalUnicast() {
					continue
				}
				// 停止同配置其他服务
				if stopPath != "" {
					var rval = reqf.Rval{Method: http.MethodOptions, Timeout: 500}
					if ip.To4() != nil {
						rval.Url = fmt.Sprintf("http://%s:%s%s", ip.String(), port, stopPath)
					} else {
						rval.Url = fmt.Sprintf("http://[%s]:%s%s", ip.String(), port, stopPath)
					}
					if err := r.Reqf(rval); err == nil {
						rval.Method = http.MethodGet
						for i := int(waitStop); i > 0; i-- {
							fmt.Printf("\r停止服务: %s %3ds", rval.Url, i)
							_ = r.Reqf(rval)
							if conn, err := net.Dial("tcp", serUrl.Host); err != nil {
								stopPath = ""
								break
							} else {
								_ = conn.Close()
							}
							time.Sleep(time.Second)
						}
						if stopPath == "" {
							fmt.Printf("\n停止服务: 成功\n")
						} else {
							fmt.Printf("\n停止服务: 超时\n")
						}
					}
				}

				// 启动时显示ip
				if !*stop && showIpOnBoot {
					if ip.To4() != nil {
						fmt.Printf("当前地址 http://%s:%s\n", ip.String(), port)
					} else {
						fmt.Printf("当前地址 http://[%s]:%s\n", ip.String(), port)
					}
				}
			}

			if *stop {
				os.Exit(0)
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

		// 登陆
		if v, ok := t.K_v.LoadV(`扫码登录`).(bool); ok && v {
			if scanPath, ok := t.K_v.LoadV("扫码登录路径").(string); ok && scanPath != "" {
				_ = file.Open("qr.png").Delete()
				t.SerF.Store(scanPath, func(w http.ResponseWriter, r *http.Request) {
					if DefaultHttpFunc(t, w, r, http.MethodGet) {
						return
					}
					if q := file.Open("qr.png"); q.IsExist() {
						_ = q.CopyToIoWriter(w, pio.CopyConfig{})
					} else if !t.IsLogin() {
						t.Danmu_Main_mq.Push_tag(`login`, nil)
						_ = q.CopyToIoWriter(w, pio.CopyConfig{})
					} else {
						w.WriteHeader(http.StatusNotFound)
					}
				})
			}
		}

		// debug模式
		if debugP, ok := t.K_v.LoadV(`debug路径`).(string); ok && debugP != "" {
			t.SerF.Store(debugP, func(w http.ResponseWriter, r *http.Request) {
				if DefaultHttpFunc(t, w, r, http.MethodGet, http.MethodPost) {
					return
				}
				if name, found := strings.CutPrefix(r.URL.Path, debugP); found && name != "" {
					switch name {
					case "cmdline":
						pprof.Cmdline(w, r)
					case "profile":
						pprof.Profile(w, r)
					case "trace":
						pprof.Trace(w, r)
					default:
						pprof.Handler(name).ServeHTTP(w, r)
					}
					return
				}
				pprof.Index(w, r)
			})
		}

		if val, ok := t.K_v.LoadV("stop路径").(string); ok && val != "" {
			t.SerF.Store(val, func(w http.ResponseWriter, r *http.Request) {
				if DefaultHttpFunc(t, w, r, http.MethodGet, http.MethodOptions) {
					return
				}
				if r.Method == http.MethodOptions {
					w.Header().Set("Allow", "GET")
					return
				}
				t.Danmu_Main_mq.Push_tag(`interrupt`, nil)
			})
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
			t.SerF.Store(val, func(w http.ResponseWriter, r *http.Request) {
				if DefaultHttpFunc(t, w, r, http.MethodGet) {
					return
				}

				if web.NotModifiedDur(r, w, time.Second*5) {
					return
				}

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

				_, timeOffset := time.Now().Zone()

				ResStruct{0, "ok", map[string]any{
					"pid":       t.PID,
					"version":   t.Version,
					"goVersion": runtime.Version(),
					"timeInfo": map[string]any{
						"timeZone":           timeOffset,
						"biliServerTimeZone": t.SerLocation,
						"startTime":          t.StartT.Format(time.RFC3339),
						"currentTime":        time.Now().Format(time.RFC3339),
					},
					"reqPoolState": map[string]any{
						"pooled":   reqState.Pooled,
						"nopooled": reqState.Nopooled,
						"inuse":    reqState.Inuse,
						"nouse":    reqState.Nouse,
						"sum":      reqState.Sum,
						"qts":      math.Round(reqState.GetPerSec*100) / 100,
					},
					"numGoroutine": runtime.NumGoroutine(),
					"mem": map[string]any{
						"memInUse":      humanize.Bytes(memStats.HeapInuse + memStats.StackInuse),
						"memTotalAlloc": humanize.Bytes(memStats.TotalAlloc),
					},
					"gc": map[string]any{
						"numGC":            memStats.NumGC,
						"lastGC":           time.UnixMicro(int64(memStats.LastGC / 1000)).Format(time.RFC3339),
						"gcCPUFractionPpm": float64(int(memStats.GCCPUFraction*100000000)) / 100,
						"gcAvgS":           float64(int(gcAvgS*100)) / 100,
					},
					"common": streams,
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

	t.Log = plog.New()
	{
		v, _ := t.K_v.LoadV("日志文件输出").(string)
		t.Log.LFile(v)

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
					if e := psql.BeginTx(db, pctx.GenTOCtx(time.Second*5)).SimpleDo(create).Run(); !psql.HasErrTx(e, nil, psql.ErrExec) {
						panic("保存日志至db打开连接错误 " + e.Error())
					}
				}
				switch dbname {
				case "postgres":
					t.Log = t.Log.LDB(psql.NewTxPool(db), psql.PlaceHolderB, insert)
				case "mysql":
					t.Log = t.Log.LDB(psql.NewTxPool(db), psql.PlaceHolderB, insert)
				case "sqlite":
					db.SetMaxOpenConns(1)
					t.Log = t.Log.LDB(psql.NewTxPool(db), psql.PlaceHolderA, insert)
				default:
				}
			}
		}

		levelM := map[plog.Level]string{
			plog.T: "T: ",
			plog.I: "I: ",
			plog.W: "W: ",
			plog.E: "E: ",
		}
		if array, ok := t.K_v.Load(`日志显示`); ok {
			maps.DeleteFunc(levelM, func(k plog.Level, v string) bool {
				return !slices.Contains(array.([]any), any(v))
			})
		}
		t.Log = t.Log.Level(levelM)
	}

	return t
}

func (t *Common) loadConf(customConf string) error {
	var data map[string]any

	// 64k
	f := file.New("config/config_K_v.json", 0, true)
	if !f.IsExist() {
		return os.ErrNotExist
	}
	if e := f.ReadToBuf(&t.buf, 16, 1<<16); e != nil {
		if !errors.Is(e, io.EOF) {
			return e
		} else {
			_ = json.Unmarshal(t.buf, &data)
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
			if req.ResStatusCode()&200 != 200 {
				return fmt.Errorf("无法获取自定义配置文件 %d", req.ResStatusCode())
			} else {
				var tmp map[string]any
				_ = req.ResponUnmarshal(json.Unmarshal, &tmp)
				for k, v := range tmp {
					data[k] = v
				}
			}
		} else {
			//从文件读取
			if bb, err := file.New(customConf, 0, true).ReadAll(100, 1<<16); err != nil {
				if errors.Is(err, io.EOF) {
					var tmp map[string]any
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

	// 从环境变量获取
	if v, ok := t.K_v.LoadV("从环境变量覆盖").([]any); ok && len(v) > 0 {
		for i := 0; i < len(v); i++ {
			if vm, ok := v[i].(map[string]any); ok {
				if err := dealEnv(&t.K_v, vm); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

var (
	ErrDealEnvUnknowType          = errors.New("ErrDealEnvUnknowType")
	ErrDealEnvEnvValueTypeNoMatch = errors.New("ErrDealEnvEnvValueTypeNoMatch")
	ErrDealEnvKeyNoArray          = errors.New("ErrDealEnvKeyNoArray")
	ErrDealEnvKeyNoMap            = errors.New("ErrDealEnvKeyNoMap")
	ErrDealEnvKeyArrayNoUInt      = errors.New("ErrDealEnvKeyArrayNoUInt")
)

// vm:{"key":"","type":"","env":""}
func dealEnv(K_v *syncmap.Map, vm map[string]any) error {
	_key, ok := vm[`key`].(string)
	if !ok || strings.TrimSpace(_key) == `` {
		return nil
	}
	var val any
	{
		_env, ok := vm[`env`].(string)
		if !ok || strings.TrimSpace(_env) == `` {
			return nil
		}
		_val := os.Getenv(_env)
		_type, ok := vm[`type`].(string)
		if !ok || strings.TrimSpace(_type) == `` {
			_type = "string"
		}
		switch _type {
		case `string`:
			val = _val
		case `float64`:
			if v, err := strconv.ParseFloat(_val, 64); err != nil {
				return ErrDealEnvEnvValueTypeNoMatch
			} else {
				val = v
			}
		case `bool`:
			switch strings.ToLower(_val) {
			case `true`:
				val = true
			case `false`:
				val = false
			default:
				return ErrDealEnvEnvValueTypeNoMatch
			}
		default:
			return ErrDealEnvUnknowType
		}
	}

	_keys := strings.Split(_key, ".")
	if len(_keys) == 1 {
		K_v.Store(_keys[0], val)
	} else {
		var key = K_v.LoadV(_keys[0])
		for i := 1; i < len(_keys)-1; i++ {
			if strings.Contains(_keys[i], "[") {
				if tmp, ok := key.([]any); !ok {
					return ErrDealEnvKeyNoArray
				} else if n, err := strconv.ParseInt(_keys[i][1:len(_keys[i])-1], 0, 64); err != nil {
					return ErrDealEnvKeyArrayNoUInt
				} else if int(n) > len(tmp)-1 {
					return nil
				} else {
					key = tmp[int(n)]
				}
			} else {
				if tmp, ok := key.(map[string]any); !ok {
					return ErrDealEnvKeyNoMap
				} else {
					key = tmp[_keys[i]]
				}
			}
		}
		if strings.Contains(_keys[len(_keys)-1], "[") {
			if tmp, ok := key.([]any); !ok {
				return ErrDealEnvKeyNoArray
			} else if n, err := strconv.ParseInt(_keys[len(_keys)-1][1:len(_keys[len(_keys)-1])-1], 0, 64); err != nil {
				return ErrDealEnvKeyArrayNoUInt
			} else if int(n) > len(tmp)-1 {
				return nil
			} else {
				tmp[n] = val
			}
		} else {
			if tmp, ok := key.(map[string]any); !ok {
				return ErrDealEnvKeyNoMap
			} else {
				tmp[_keys[len(_keys)-1]] = val
			}
		}
	}
	return nil
}

var C = new(Common).Init()

// StreamRec
// fmp4
// https://datatracker.ietf.org/doc/html/draft-pantos-http-live-streaming
var StreamO = new(sync.Map)
var Commons = new(syncmap.Map)
var CommonsLoadOrInit = syncmap.NewLoadOrInitFunc[Common](Commons).SetInit(func() *Common {
	return C.Copy()
})

// 消息队列
type Danmu_Main_mq_item struct {
	Class string
	Data  any
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
	if strings.Contains(r.RequestURI, "..") {
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
