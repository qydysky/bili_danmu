package json

import (
	"encoding/json"
	"testing"

	cmp "github.com/qydysky/part/component2"
)

// 11004 ns/op               8 B/op          1 allocs/op
func Benchmark_2(b *testing.B) {
	type g struct {
		A123 int
		Asf  string
	}
	a := g{}

	data := []byte(`{"A123":123,"asf":"1"}`)

	json := cmp.GetV3[interface {
		Unmarshal(data []byte, v any) error
	}](`json`).Inter()

	for b.Loop() {
		json.Unmarshal(data, &a)
	}
}

// 6388 ns/op             224 B/op          4 allocs/op
func Benchmark_1(b *testing.B) {
	type g struct {
		A123 int
		Asf  string
	}
	a := g{}
	data := []byte(`{"A123":123,"asf":"1"}`)
	for b.Loop() {
		json.Unmarshal(data, &a)
	}
}

func Test(t *testing.T) {
	type g struct {
		A123 int
		Asf  string
	}
	a := g{}

	data := []byte(`{"A123":123,"asf":"1"}`)

	json := cmp.GetV3[interface {
		Unmarshal(data []byte, v any) error
	}](`json`).Inter()

	if e := json.Unmarshal(data, &a); e != nil || a.A123 != 123 || a.Asf != "1" {
		t.Fatal(e)
	}
}

func Test2(t *testing.T) {
	type g struct {
		A123 int
		Asf  string
	}
	a := g{}

	data := []byte(`{"A123":123,"asf:"1"}`)
	data1 := []byte(`{"A123":123,"asf":"1"}`)

	json := cmp.GetV3[interface {
		Unmarshal(data []byte, v any) error
	}](`json`).Inter()

	if e := json.Unmarshal(data, &a); e == nil {
		t.Fatal(e)
	}
	if e := json.Unmarshal(data1, &a); e != nil || a.A123 != 123 || a.Asf != "1" {
		t.Fatal(e)
	}
}
