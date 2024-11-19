package ass

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	p "github.com/qydysky/part"
	comp "github.com/qydysky/part/component2"
	pctx "github.com/qydysky/part/ctx"
	file "github.com/qydysky/part/file"
	encoder "golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type i interface {
	Assf(s string) error
	Ass_f(ctx context.Context, enc, savePath string, st time.Time)
}

func init() {
	if e := comp.Register[i]("ass", &Ass{header: Ass_header}); e != nil {
		panic(e)
	}
}

var (
	Ass_height = 720  //字幕高度
	Ass_width  = 1280 //字幕宽度
	Ass_font   = 50   //字幕字体大小
	Ass_T      = 7    //单条字幕显示时间
	Ass_loc    = 7    //字幕位置 小键盘对应的位置
	accept     = map[string]encoder.Encoding{
		``:        simplifiedchinese.GB18030,
		`GB18030`: simplifiedchinese.GB18030,
		`utf-8`:   nil,
	}
	Ass_header = `[Script Info]
	Title: Default Ass file
	ScriptType: v4.00+
	WrapStyle: 0
	ScaledBorderAndShadow: yes
	PlayResX: ` + strconv.Itoa(Ass_height) + `
	PlayResY: ` + strconv.Itoa(Ass_width) + `
	
	[V4+ Styles]
	Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
	Style: Default,,` + strconv.Itoa(Ass_font) + `,&H40FFFFFF,&H000017FF,&H80000000,&H40000000,0,0,0,0,100,100,0,0,1,4,4,` + strconv.Itoa(Ass_loc) + `,20,20,50,1
	
	[Events]
	Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
	`
)

// Ass 弹幕转字幕
type Ass struct {
	savePath string           //弹幕ass文件名
	startT   time.Time        //开始记录的基准时间
	header   string           //ass开头
	wrap     encoder.Encoding //编码
}

func (t *Ass) Assf(s string) error {
	if t.savePath == "" || s == "" {
		return nil
	}

	st := time.Since(t.startT) + time.Duration(p.Rand().MixRandom(0, 2000))*time.Millisecond
	et := st + time.Duration(Ass_T)*time.Second

	var b string
	// b += "Comment: " + strconv.Itoa(loc) + " "+ Dtos(showedt) + "\n"
	b += `Dialogue: 0,`
	b += dtos(st) + `,` + dtos(et)
	b += `,Default,,0,0,0,,{\fad(200,500)\blur3}` + s + "\n"

	f := file.New(t.savePath+"0.ass", -1, true)
	f.Config.Coder = t.wrap
	if _, e := f.Write([]byte(b), true); e != nil {
		return e
	}
	return nil
}

func (t *Ass) Ass_f(ctx context.Context, enc, savePath string, st time.Time) {
	if v1, ok := accept[enc]; ok {
		t.wrap = v1
	}

	t.savePath = savePath
	f := &file.File{
		Config: file.Config{
			FilePath:  t.savePath + "0.ass",
			AutoClose: true,
			Coder:     t.wrap,
		},
	}
	_, _ = f.Write([]byte(t.header), true)
	t.startT = st

	ctx, done := pctx.WaitCtx(ctx)
	defer done()
	<-ctx.Done()

	t.savePath = ""
}

// 时间转化为0:00:00.00规格字符串
func dtos(t time.Duration) string {
	M := int(math.Floor(t.Minutes())) % 60
	S := int(math.Floor(t.Seconds())) % 60
	Ns := t.Nanoseconds() / int64(time.Millisecond) % 1000 / 10

	return fmt.Sprintf("%d:%02d:%02d.%02d", int(math.Floor(t.Hours())), M, S, Ns)
}
