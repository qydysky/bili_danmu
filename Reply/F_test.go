package reply

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	c "github.com/qydysky/bili_danmu/CV"
	psql "github.com/qydysky/part/sql"
)

func TestSaveDanmuToDB(t *testing.T) {
	c.C.K_v.Store(`保存弹幕至db`, map[string]any{
		"dbname": "sqlite",
		"url":    "danmu.sqlite3",
		"create": "create table danmu (created text, createdunix text, msg text, color text, auth text, uid text, roomid text)",
		"insert": "insert into danmu  values ({Date},{Unix},{Msg},{Color},{Auth},{Uid},{Roomid})",
	})
	saveDanmuToDB.init(c.C)
	saveDanmuToDB.danmu(Danmu_item{
		msg:    "可能走位配合了他的压枪",
		color:  "#54eed8",
		auth:   "畏未",
		uid:    "96767379",
		roomid: 92613,
	})
	saveDanmuToDB.db.Close()

	if db, e := sql.Open("sqlite", "danmu.sqlite3"); e != nil {
		t.Fatal(e)
	} else {
		tx := psql.BeginTx[any](db, context.Background())
		tx.Do(psql.SqlFunc[any]{Query: "select msg as Msg from danmu"})
		tx.AfterQF(func(_ *any, rows *sql.Rows, e *error) {
			type row struct {
				Msg string
			}

			v, err := psql.DealRows(rows, func() row { return row{} })
			if err != nil {
				*e = err
				return
			}
			if len(v) != 1 || v[0].Msg != "可能走位配合了他的压枪" {
				*e = errors.New("no msg")
				return
			}
		})
		if _, e := tx.Fin(); e != nil {
			t.Fatal(e)
		}
		db.Close()
	}
}

func Test_getRecInfo(t *testing.T) {
	pathInfo, err := getRecInfo("testdata/live/2023_06_27-02_17_48-7734200-【预告】27日15点 JDG vs NIP！-原画-YI2")
	if err != nil {
		t.Fatal(err)
	}
	if pathInfo.Name != "【预告】27日15点 JDG vs NIP！" {
		t.Fatal()
	}
	if pathInfo.Path != "testdata/live/2023_06_27-02_17_48-7734200-【预告】27日15点 JDG vs NIP！-原画-YI2" {
		t.Fatal()
	}
	if pathInfo.Qn != "原画" {
		t.Fatal()
	}
	if pathInfo.StartT != "2023-06-25 15:29:33" {
		t.Fatal()
	}
	if pathInfo.Roomid != 7734200 {
		t.Fatal()
	}
	if pathInfo.UpUid != 50329118 {
		t.Fatal()
	}
}
