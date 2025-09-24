package savedanmutodb

import (
	"database/sql"
	"errors"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	comp "github.com/qydysky/part/component2"
	pctx "github.com/qydysky/part/ctx"
	log "github.com/qydysky/part/log"
	psql "github.com/qydysky/part/sql"
	_ "modernc.org/sqlite"
)

// 保存弹幕至db
func init() {
	comp.RegisterOrPanic[interface {
		Init(config any, fl *log.Log_interface)
		Danmu(Msg string, Color string, Auth any, Uid string, Roomid int64)
		Close() error
	}](`saveDanmuToDB`, &saveDanmuToDB{})
}

type saveDanmuToDB struct {
	state  atomic.Int32
	dbname string
	db     *sql.DB
	insert string
	fl     *log.Log_interface
}

func (t *saveDanmuToDB) Init(config any, fl *log.Log_interface) {
	if t.state.CompareAndSwap(0, 1) {
		if v, ok := config.(map[string]any); ok && len(v) != 0 {
			var (
				dbname, url, create                 string
				dbnameok, urlok, createok, insertok bool
			)

			dbname, dbnameok = v["dbname"].(string)
			url, urlok = v["url"].(string)
			create, createok = v["create"].(string)
			t.insert, insertok = v["insert"].(string)

			if dbname == "" || url == "" || t.insert == "" || !dbnameok || !urlok || !insertok {
				t.state.CompareAndSwap(1, 0)
				return
			}

			t.dbname = dbname

			t.fl = fl.Base_add("保存弹幕至db")
			if db, e := sql.Open(dbname, url); e != nil {
				t.fl.L(`E: `, e)
			} else {
				db.SetConnMaxLifetime(time.Minute * 3)
				db.SetMaxOpenConns(10)
				db.SetMaxIdleConns(10)
				t.db = db
				if createok {
					tx := psql.BeginTx[any](db, pctx.GenTOCtx(time.Second*5))
					tx.Do(psql.SqlFunc[any]{Sql: create, SkipSqlErr: true})
					if _, e := tx.Fin(); e != nil {
						t.fl.L(`E: `, e)
						t.state.CompareAndSwap(1, 0)
						return
					}
				}
				t.fl.L(`I: `, dbname)
				t.state.CompareAndSwap(1, 2)
				return
			}
		}
	}
	t.state.CompareAndSwap(1, 0)
}

func (t *saveDanmuToDB) Danmu(Msg string, Color string, Auth any, Uid string, Roomid int64) {
	if t.state.Load() == 2 {
		if e := t.db.Ping(); e == nil {
			type DanmuI struct {
				Date   string
				Unix   int64
				Msg    string
				Color  string
				Auth   any
				Uid    string
				Roomid int64
			}

			var replaceF psql.ReplaceF
			switch t.dbname {
			case "postgres":
				replaceF = psql.PlaceHolderB
			default:
				replaceF = psql.PlaceHolderA
			}

			tx := psql.BeginTx[any](t.db, pctx.GenTOCtx(time.Second*5))
			tx.DoPlaceHolder(psql.SqlFunc[any]{Sql: t.insert}, &DanmuI{
				Date:   time.Now().Format(time.DateTime),
				Unix:   time.Now().Unix(),
				Msg:    Msg,
				Color:  Color,
				Auth:   Auth,
				Uid:    Uid,
				Roomid: Roomid,
			}, replaceF)
			tx.AfterEF(func(_ *any, result sql.Result, e *error) {
				if v, err := result.RowsAffected(); err != nil {
					*e = err
					return
				} else if v != 1 {
					*e = errors.New("插入数量错误")
					return
				}
			})
			if _, e := tx.Fin(); e != nil {
				t.fl.L(`E: `, e)
			}
		}
	}
}

func (t *saveDanmuToDB) Close() error {
	t.state.Store(0)
	return t.db.Close()
}
