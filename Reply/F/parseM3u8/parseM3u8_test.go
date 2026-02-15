package parsem3u8

import (
	"iter"
	"testing"

	_ "embed"

	comp "github.com/qydysky/part/component2"
)

//go:embed test.m3u8
var testM3u8 []byte

func TestMain(t *testing.T) {
	var ParseM3u8 = comp.GetV3[interface {
		Parse(respon []byte, lastNo int) (m4sLink iter.Seq[interface {
			IsHeader() bool
			M4sLink() string
		}], redirectUrl string, err error)
		IsErrRedirect(e error) bool
	}](`parseM3u8`).Inter()
	_, url, e := ParseM3u8.Parse(testM3u8, 0)
	if !ParseM3u8.IsErrRedirect(e) {
		t.Fatal(e, url)
	}
	if url != "https://d1--cn-gotcha209.bilivideo.com/live-bvc/192693/live_194484313_8775758_bluray/index.m3u8?expires=1737129387&len=0&oi=x&pt=web&qn=10000&trid=x&sigparams=cdn,expires,len,oi,pt,qn,trid&cdn=cn-gotcha209&sign=x&site=x&free_type=0&mid=x&sche=ban&bvchls=1&trace=64&isp=cm&rg=South&pv=Guangdong&deploy_env=prod&hot_cdn=909773&origin_bitrate=1272698&source=puv3_master&flvsk=x&suffix=bluray&pp=rtmp&qp=bqor_250&sl=10&p2p_type=1&sk=x&score=42&info_source=cache&vd=bc&src=puv3&order=1" {
		t.Fatal()
	}
}
