package ass

import (
	"bytes"
	"encoding/json"
	"fmt"
	"iter"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/dustin/go-humanize"
	comp "github.com/qydysky/part/component2"
	file "github.com/qydysky/part/file"
)

var (
	playResX = 1280 //字幕宽度
	playResY = 720  //字幕高度
)

type i interface {
	ToAss(savePath string, filename ...string)
	Init(cfg any)
}

func init() {
	if e := comp.Register[i]("ass", &Ass{
		fontsize: 40,
		showSec:  10,
		area:     1.0,
		alpha:    0,
		header: `[Script Info]
Title: Default Ass file
ScriptType: v4.00+
WrapStyle: 0
ScaledBorderAndShadow: yes
PlayResX: ` + strconv.Itoa(playResX) + `
PlayResY: ` + strconv.Itoa(playResY) + `

[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: Default,,40,&H40FFFFFF,&H000017FF,&H80000000,&H40000000,0,0,0,0,100,100,0,0,1,1,0,7,20,20,50,1

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`,
	}); e != nil {
		panic(e)
	}
}

// Ass 弹幕转字幕
type Ass struct {
	showSec, fontsize int
	area              float64
	alpha             int
	header            string //ass开头
}

func (t *Ass) Init(cfg any) {
	if c, ok := cfg.(map[string]any); !ok {
		return
	} else {
		fontname := c[`fontname`].(string)
		if tmp, ok := c[`fontsize`].(float64); ok && tmp > 0 {
			t.fontsize = int(tmp)
		}
		if tmp, ok := c[`showSec`].(float64); ok && tmp > 0 {
			t.showSec = int(tmp)
		}
		if tmp, ok := c[`area`].(float64); ok && tmp >= 0 {
			t.area = tmp
		}
		if tmp, ok := c[`alpha`].(float64); ok && tmp >= 0 {
			t.alpha = int(tmp * 255)
		}
		t.header = `[Script Info]
Title: Default Ass file
ScriptType: v4.00+
WrapStyle: 0
ScaledBorderAndShadow: yes
PlayResX: ` + strconv.Itoa(playResX) + `
PlayResY: ` + strconv.Itoa(playResY) + `

[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: Default,` + fontname + `,` + strconv.Itoa(t.fontsize) + `,&H40FFFFFF,&H000017FF,&H80000000,&H40000000,0,0,0,0,100,100,0,0,1,1,0,7,20,20,50,1

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`
	}
}

func (t *Ass) ToAss(savePath string, filename ...string) {
	f := file.New(savePath+append(filename, "0.ass")[0], 0, false)
	defer f.Close()
	if f.IsExist() {
		_ = f.Delete()
	}

	lsSize := int(float64(playResY) * t.area / float64(t.fontsize))
	var lsd = make([]float64, lsSize)
	var lso = make([]float64, lsSize)

	var write bool
	for line := range loadCsv(savePath, strings.Split(append(filename, "0.ass")[0], `.`)[0]+".csv") {
		if !write {
			_, _ = f.Write([]byte(t.header))
			write = true
		}

		danmul := utf8.RuneCountInString(line.Text)
		danmuSec := (float64(t.showSec*t.fontsize*danmul) / float64(t.fontsize*danmul+playResX))

		c := -1

		for i := 0; i < lsSize; i++ {
			if lsd[i] > line.Time+float64(t.showSec)-danmuSec {
				continue
			}
			if lso[i] > line.Time {
				continue
			}
			{
				lsd[i] = line.Time + float64(t.showSec) //line.Time + (float64(showSec*fontsize*danmul) / float64(fontsize*danmul+playResX))
				lso[i] = line.Time + +danmuSec
				c = i
				break
			}
		}

		if c == -1 {
			continue
		}

		_, _ = f.Write([]byte(
			`Dialogue: 0,` +
				stos(line.Time) + `,` + stos(line.Time+float64(t.showSec)) +
				`,Default,,0,0,0,,{` +
				`\c&H` + line.Style.Color[5:7] + line.Style.Color[3:5] + line.Style.Color[1:3] + `&` +
				`\alpha&H` + fmt.Sprintf("%02x", t.alpha) + `&` +
				`\move(` + strconv.Itoa(playResX) + `,` + strconv.Itoa(c*t.fontsize) + `,-` + strconv.Itoa(t.fontsize*danmul) + `,` + strconv.Itoa(c*t.fontsize) + `)` +
				`}` + line.Text + "\n"))
	}
}

func loadCsv(savePath string, filename ...string) iter.Seq[Data] {
	return func(yield func(Data) bool) {
		csvf := file.New(savePath+append(filename, "0.csv")[0], 0, false)
		defer csvf.Close()

		if !csvf.IsExist() {
			return
		}

		var data = Data{}
		for i := 0; true; i += 1 {
			if line, e := csvf.ReadUntil([]byte{'\n'}, humanize.KByte, humanize.MByte); len(line) != 0 {
				lined := bytes.SplitN(line, []byte{','}, 3)
				if len(lined) == 3 {
					if t, e := strconv.ParseFloat(string(lined[0]), 64); e == nil {
						if e := json.Unmarshal(lined[2], &data); e == nil {
							data.Time = t
							if data.Style.Color == "" {
								data.Style.Color = "#FFFFFF"
							}
							if !yield(data) {
								return
							}
						}
					}
				}
			} else if e != nil {
				break
			}
		}
	}
}

type DataStyle struct {
	Color  string `json:"color"`
	Border bool   `json:"border"`
	Mode   int    `json:"mode"`
}

type Data struct {
	Text  string    `json:"text"`
	Style DataStyle `json:"style"`
	Time  float64   `json:"time"`
}

func stos(sec float64) string {
	H := int(math.Floor(sec)) / 3600
	M := int(math.Floor(sec)) % 3600 / 60
	S := int(math.Floor(sec)) % 60
	Ns := int(sec*1000) % 1000 / 10

	return fmt.Sprintf("%d:%02d:%02d.%02d", H, M, S, Ns)
}
