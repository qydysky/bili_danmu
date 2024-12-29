package cv

import (
	_ "embed"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	syncmap "github.com/qydysky/part/sync"
	_ "modernc.org/sqlite"
)

func TestDealEnv(t *testing.T) {
	os.Setenv("tes", "2")
	os.Setenv("tes1", "true")
	os.Setenv("tes2", "true")
	var m syncmap.Map
	m.Store("k", float64(1))
	m.Store("b", false)
	m.Store("s", "s")
	if e := dealEnv(&m, map[string]any{`key`: `k`, `type`: `float64`, `env`: `tes`}); e != nil {
		t.Fatal(e)
	}
	if v, ok := m.LoadV("k").(float64); !ok || v != 2 {
		t.Fatal(v)
	}
	if e := dealEnv(&m, map[string]any{`key`: `b`, `type`: `bool`, `env`: `tes1`}); e != nil {
		t.Fatal(e)
	}
	if v, ok := m.LoadV("b").(bool); !ok || !v {
		t.Fatal(v)
	}
	if e := dealEnv(&m, map[string]any{`key`: `s`, `env`: `tes2`}); e != nil {
		t.Fatal(e)
	}
	if v, ok := m.LoadV("s").(string); !ok || v != "true" {
		t.Fatal(v)
	}
}

func TestDealEnv2(t *testing.T) {
	os.Setenv("tes", "2")
	var m syncmap.Map
	m.Store("k", map[string]any{"d": float64(1)})
	if e := dealEnv(&m, map[string]any{`key`: `k.d`, `type`: `float64`, `env`: `tes`}); e != nil {
		t.Fatal(e)
	}
	if v, ok := m.LoadV("k").(map[string]any); !ok {
		t.Fatal(v)
	} else if v[`d`].(float64) != 2 {
		t.Fatal()
	}
}

func TestDealEnv3(t *testing.T) {
	os.Setenv("tes", "2")
	var m syncmap.Map
	m.Store("k", []any{float64(1)})
	if e := dealEnv(&m, map[string]any{`key`: `k.[0]`, `type`: `float64`, `env`: `tes`}); e != nil {
		t.Fatal(e)
	}
	if v, ok := m.LoadV("k").([]any); !ok {
		t.Fatal(v)
	} else if v[0].(float64) != 2 {
		t.Fatal()
	}
}

func TestDealEnv4(t *testing.T) {
	os.Setenv("tes", "2")
	var m syncmap.Map
	m.Store("k", []any{map[string]any{"d": float64(1)}})
	if e := dealEnv(&m, map[string]any{`key`: `k.[0].d`, `type`: `float64`, `env`: `tes`}); e != nil {
		t.Fatal(e)
	}
	if v, ok := m.LoadV("k").([]any); !ok {
		t.Fatal(v)
	} else if q, ok := v[0].(map[string]any); !ok {
		t.Fatal(q)
	} else if q["d"].(float64) != 2 {
		t.Fatal(q["d"])
	}
}
