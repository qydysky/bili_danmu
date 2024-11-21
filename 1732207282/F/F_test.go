package F

import (
	"net/url"
	"testing"
)

func Test2(t *testing.T) {
	rawURL := "http://127.0.0.1:10841/1?workspace=/codefile/qydysky.code-workspace#12"
	u, _ := url.Parse(rawURL)
	if u.Host != ParseHost(rawURL) {
		t.Fatal()
	}
	if u.Query().Get("workspace") != ParseQuery(rawURL, "workspace=") {
		t.Log(u.Query().Get("workspace"))
		t.Log(ParseQuery(rawURL, "workspace="))
		t.Fatal()
	}
}

func Test3(t *testing.T) {
	rawURL := "http://127.0.0.1:10841/1?workspace=/codefile/qydysky.code-workspace#12"

	u, _ := url.Parse(rawURL)
	u1, _ := url.Parse("./2")

	if u.ResolveReference(u1).String() != ResolveReferenceLast(rawURL, "2") {
		t.Fatal()
	}
}
