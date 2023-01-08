package reply

import (
	_ "embed"
	"testing"

	F "github.com/qydysky/bili_danmu/F"
)

var buf []byte

func Test_deal(t *testing.T) {
	err := deal(buf,
		[]string{"moof", "mfhd",
			"traf", "tfhd", "tfdt", "trun",
			"traf", "tfhd", "tfdt", "trun",
			"mdat"},
		func(m []*ie) bool {
			moofSN := int(F.Btoi(buf, m[1].i+12, 4))
			keyframeMoof := buf[m[5].i+20] == byte(0x02) || buf[m[9].i+20] == byte(0x02)
			t.Log(moofSN, "frame", keyframeMoof, m[0].i, m[10].n, m[10].e)
			return false
		})
	t.Log("err", err)
}
