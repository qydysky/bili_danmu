package cv

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"time"

	file "github.com/qydysky/part/file"
	idpool "github.com/qydysky/part/idpool"
	log "github.com/qydysky/part/log"
	mq "github.com/qydysky/part/msgq"
	reqf "github.com/qydysky/part/reqf"
	syncmap "github.com/qydysky/part/sync"
)

type Common struct {
	Uid               int                //client uid
	Live              []string           //直播流链接
	Live_qn           int                //当前直播流质量
	Live_want_qn      int                //期望直播流质量
	Roomid            int                //房间ID
	Cookie            syncmap.Map        //Cookie
	Title             string             //直播标题
	Uname             string             //主播名
	UpUid             int                //主播uid
	Rev               float64            //营收
	Renqi             int                //人气
	Watched           int                //观看人数
	GuardNum          int                //舰长数
	ParentAreaID      int                //父分区
	AreaID            int                //子分区
	Locked            bool               //直播间封禁
	Note              string             //分区排行
	Live_Start_Time   time.Time          //直播开始时间
	Liveing           bool               //是否在直播
	Wearing_FansMedal int                //当前佩戴的粉丝牌
	Token             string             //弹幕钥
	WSURL             []string           //弹幕链接
	LIVE_BUVID        bool               //cookies含LIVE_BUVID
	Stream_url        []string           //直播Web服务
	Proxy             string             //全局代理
	AcceptQn          map[int]string     //允许的直播流质量
	Qn                map[int]string     //全部直播流质量
	K_v               syncmap.Map        //配置文件
	Log               *log.Log_interface //日志
	Danmu_Main_mq     *mq.Msgq           //消息
	ReqPool           *idpool.Idpool     //请求池
}

func (t *Common) init() Common {
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

	t.Danmu_Main_mq = mq.New(200)

	var (
		ckv     = flag.String("ckv", "", "自定义配置KV文件，将会覆盖config_K_v配置")
		roomIdP = flag.Int("r", 0, "roomid")
	)
	flag.Parse()
	t.Roomid = *roomIdP

	if bb, err := file.New("config/config_K_v.json", 0, true).ReadAll(100, 1<<16); err != nil {
		if errors.Is(err, io.EOF) {
			var data map[string]interface{}
			json.Unmarshal(bb, &data)
			for k, v := range data {
				t.K_v.Store(k, v)
			}
		} else {
			panic(err)
		}
	}

	if t.K_v.Len() == 0 {
		panic("未能加载配置")
	}

	if *ckv != "" {
		if bb, err := file.New(*ckv, 0, true).ReadAll(100, 1<<16); err != nil {
			if errors.Is(err, io.EOF) {
				var data map[string]interface{}
				json.Unmarshal(bb, &data)
				for k, v := range data {
					t.K_v.Store(k, v)
				}
			} else {
				panic(err)
			}
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

			// 64k
			if bb, e := file.New("config/config_K_v.json", 0, true).ReadAll(100, 1<<16); e != nil {
				if !errors.Is(e, io.EOF) {
					panic(e)
				} else {
					var data map[string]interface{}
					json.Unmarshal(bb, &data)
					for k, v := range data {
						t.K_v.Store(k, v)
					}
				}
			}
		}
	}()

	if val, exist := t.K_v.Load("http代理地址"); exist {
		t.Proxy = val.(string)
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
		logmap := make(map[string]struct{})
		if array, ok := t.K_v.Load(`日志显示`); ok {
			for _, v := range array.([]interface{}) {
				logmap[v.(string)] = log.On
			}
		}
		t.Log = t.Log.Level(logmap)
	}

	t.ReqPool = idpool.New(func() interface{} {
		return reqf.New()
	})
	return *t
}

var C = new(Common).init()

// 消息队列
type Danmu_Main_mq_item struct {
	Class string
	Data  interface{}
}
