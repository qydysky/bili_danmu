module github.com/qydysky/bili_danmu/demo

go 1.14

require (
	github.com/christopher-dG/go-obs-websocket v0.0.0-20200720193653-c4fed10356a5 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/gotk3/gotk3 v0.5.1 // indirect
	github.com/klauspost/compress v1.11.3 // indirect
	github.com/miekg/dns v1.1.35 // indirect
	github.com/mitchellh/mapstructure v1.4.0 // indirect
	github.com/qydysky/bili_danmu v0.5.2
	github.com/qydysky/part v0.3.3 // indirect
	github.com/qydysky/part/msgq v0.0.0-20201213031129-ca3253dc72ad // indirect
	github.com/shirou/gopsutil v3.20.11+incompatible // indirect
	golang.org/x/crypto v0.0.0-20201203163018-be400aefbc4c // indirect
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb // indirect
	golang.org/x/sys v0.0.0-20201204225414-ed752295db88 // indirect
)

replace (
	github.com/gotk3/gotk3 v0.5.1 => github.com/qydysky/gotk3 v0.5.1-0.20201205180217-ed1a98fbc6dc
	github.com/qydysky/bili_danmu => ../
)
