module github.com/qydysky/bili_danmu

go 1.14

require (
	github.com/christopher-dG/go-obs-websocket v0.0.0-20200720193653-c4fed10356a5
	github.com/gorilla/websocket v1.4.2
	github.com/klauspost/compress v1.11.0 // indirect
	github.com/qydysky/part v0.0.4
	github.com/shirou/gopsutil v2.20.8+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200917061948-648f2a039071 // indirect
)

//replace github.com/qydysky/part => ../part
