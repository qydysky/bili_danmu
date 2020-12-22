module github.com/qydysky/bili_danmu/demo

go 1.14

require (
	github.com/christopher-dG/go-obs-websocket v0.0.0-20200720193653-c4fed10356a5 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/gotk3/gotk3 v0.5.1 // indirect
	github.com/klauspost/compress v1.11.3 // indirect
	github.com/miekg/dns v1.1.35 // indirect
	github.com/mitchellh/mapstructure v1.4.0 // indirect
	github.com/qydysky/bili_danmu v0.5.4
	github.com/qydysky/part v0.3.5-0.20201222075205-70243aca6682 // indirect
	github.com/shirou/gopsutil v3.20.11+incompatible // indirect
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e // indirect
	golang.org/x/crypto v0.0.0-20201217014255-9d1352758620 // indirect
	golang.org/x/net v0.0.0-20201216054612-986b41b23924 // indirect
	golang.org/x/sys v0.0.0-20201218084310-7d0127a74742 // indirect
)

replace (
	github.com/gotk3/gotk3 v0.5.1 => github.com/qydysky/gotk3 v0.5.1-0.20201205180217-ed1a98fbc6dc
	github.com/qydysky/bili_danmu => ../
)
