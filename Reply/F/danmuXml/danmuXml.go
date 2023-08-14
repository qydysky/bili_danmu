package danmuXml

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dustin/go-humanize"
	comp "github.com/qydysky/part/component"
	file "github.com/qydysky/part/file"
)

type Sign struct {
	// path: csv所在目录，末尾无
	toXml func(ctx context.Context, path *string) error
}

func init() {
	sign := Sign{
		toXml: toXml,
	}
	if e := comp.Put[string](comp.Sign[Sign](`toXml`), sign.toXml); e != nil {
		panic(e)
	}
}

type danmu struct {
	XMLName    xml.Name `xml:"i"`
	Chatserver string   `xml:"chatserver"`
	Chatid     int      `xml:"chatid"`
	Mission    int      `xml:"mission"`
	Maxlimit   int      `xml:"maxlimit"`
	State      int      `xml:"state"`
	RealName   int      `xml:"real_name"`
	Source     string   `xml:"source"`
	D          []danmuD `xml:"d"`
}

type danmuD struct {
	P    string `xml:"p,attr"`
	Data string `xml:",chardata"`
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

func toXml(ctx context.Context, path *string) error {
	d := danmu{
		Chatserver: "chat.bilibili.com",
		Chatid:     0,
		Mission:    0,
		Maxlimit:   1500,
		State:      0,
		RealName:   0,
		Source:     "e-r",
	}

	csvf := file.New((*path)+"0.csv", 0, false)
	var data = Data{}
	for i := 0; true; i += 1 {
		if line, e := csvf.ReadUntil('\n', humanize.KByte, humanize.MByte); len(line) != 0 {
			lined := bytes.SplitN(line, []byte{','}, 3)
			if len(lined) == 3 {
				if e := json.Unmarshal(lined[2], &data); e == nil {
					if data.Style.Color == "" {
						data.Style.Color = "#FFFFFF"
					}
					if color, e := strconv.ParseInt(data.Style.Color[1:], 16, 0); e == nil {
						d.D = append(d.D, danmuD{
							P:    fmt.Sprintf("%s,1,25,%d,0,0,0,%d", lined[0], color, i),
							Data: data.Text,
						})
					}
				}
			}
		} else if e != nil {
			break
		}
	}
	csvf.Close()

	output, err := xml.MarshalIndent(d, "", "    ")
	if err != nil {
		return err
	}

	f := file.New((*path)+"0.xml", 0, false)
	defer f.Close()

	if _, err := f.Write([]byte(xml.Header), true); err != nil {
		return err
	}

	if _, err := f.Write(output, true); err != nil {
		return err
	}

	return nil
}
