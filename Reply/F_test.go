package Reply

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	c "github.com/qydysky/bili_danmu/CV"
	replyFunc "github.com/qydysky/bili_danmu/Reply/F"
	videoInfo "github.com/qydysky/bili_danmu/Reply/F/videoInfo"
	pctx "github.com/qydysky/part/ctx"
	psql "github.com/qydysky/part/sql"
)

func TestSaveDanmuToDB(t *testing.T) {
	c.C.K_v.Store(`保存弹幕至db`, map[string]any{
		"dbname": "sqlite",
		"url":    "danmu.sqlite3",
		"create": "create table danmu (created text, createdunix text, msg text, color text, auth text, uid text, roomid text)",
		"insert": "insert into danmu  values ({Date},{Unix},{Msg},{Color},{Auth},{Uid},{Roomid})",
	})
	_ = replyFunc.SaveDanmuToDB.Run(func(sdtd replyFunc.SaveDanmuToDBI) error {
		sdtd.Init(c.C.K_v.LoadV(`保存弹幕至db`), msglog)
		sdtd.Danmu("可能走位配合了他的压枪", "#54eed8", "畏未", "96767379", 92613)
		return sdtd.Close()
	})

	if db, e := sql.Open("sqlite", "danmu.sqlite3"); e != nil {
		t.Fatal(e)
	} else {
		tx := psql.BeginTx[any](db, pctx.GenTOCtx(time.Second*5))
		tx.Do(&psql.SqlFunc[any]{Sql: "select msg as Msg from danmu"})
		tx.AfterQF(func(_ *any, rows *sql.Rows) (e error) {
			type row struct {
				Msg string
			}

			v, err := psql.DealRows[row](rows)
			if err != nil {
				return err
			}
			if len(v) != 1 || v[0].Msg != "可能走位配合了他的压枪" {
				return errors.New("no msg")
			}
			return
		})
		if _, e := tx.Fin(); e != nil {
			t.Fatal(e)
		}
		_ = db.Close()
	}
}

func Test_getRecInfo(t *testing.T) {
	pathInfo, err := videoInfo.Get.Run(context.Background(), "testdata/live/2023_07_10-14_49_10-22259479-10000-644c3e-vfv")
	if err != nil {
		t.Fatal(err)
	}
	if pathInfo.Name != "🐟在太空宿舍里该做什么呢🐟" {
		t.Fatal()
	}
	if pathInfo.Path != "2023_06_28-09_15_48-22259479-🐟在太空宿舍里该做什么呢🐟-原画-aw4" {
		t.Fatal()
	}
	if pathInfo.Qn != "原画" {
		t.Fatal()
	}
	if pathInfo.StartT != "2023-06-28 09:15:48" {
		t.Fatal()
	}
	if pathInfo.StartLiveT != "2023-06-28 00:23:35" {
		t.Fatal()
	}
	if pathInfo.Roomid != 22259479 {
		t.Fatal()
	}
	if pathInfo.UpUid != 592507317 {
		t.Fatal()
	}
	if pathInfo.Uname != "烤鱼子Official" {
		t.Fatal()
	}
}
