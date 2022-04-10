package stream

import (
	"path/filepath"

	c "github.com/qydysky/bili_danmu/CV"

	log "github.com/qydysky/part/log"
	signal "github.com/qydysky/part/signal"
)

type Stream struct {
	Status *signal.Signal //IsLive()是否运行中
	log    *log.Log_interface
	config Stream_Config //配置
}

type Stream_Config struct {
	save_path     string //直播流保存目录
	want_qn       int    //直播流清晰度
	want_type     string //直播流类型
	bufsize       int    //直播hls流缓冲
	banlance_host bool   //直播hls流均衡
}

func (t *Stream) LoadConfig() {
	//读取配置
	if path, ok := c.C.K_v.LoadV("直播流保存位置").(string); ok {
		if path, err := filepath.Abs(path); err == nil {
			t.config.save_path = path + "/"
		}
	}
	if v, ok := c.C.K_v.LoadV(`直播hls流缓冲`).(float64); ok && v > 0 {
		t.config.bufsize = int(v)
	}
	if v, ok := c.C.K_v.LoadV(`直播hls流均衡`).(bool); ok {
		t.config.banlance_host = v
	}
	if v, ok := c.C.K_v.LoadV(`直播流清晰度`).(int); ok {
		t.config.want_qn = v
	}
	if v, ok := c.C.K_v.LoadV(`直播流类型`).(string); ok {
		t.config.want_type = v
	}
}

func (t *Stream) Start() {
	t.log = c.C.Log.Base(`直播流保存`)
}
