module github.com/qydysky/bili_danmu/demo

go 1.14

require (
	github.com/christopher-dG/go-obs-websocket v0.0.0-20200720193653-c4fed10356a5 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/gotk3/gotk3 v0.5.2 // indirect
	github.com/klauspost/compress v1.11.4 // indirect
	github.com/miekg/dns v1.1.35 // indirect
	github.com/mitchellh/mapstructure v1.4.0 // indirect
	github.com/qydysky/bili_danmu v0.5.7
	github.com/qydysky/part v0.3.5-0.20210105160037-508c706d691b // indirect
	github.com/shirou/gopsutil v3.20.12+incompatible // indirect
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/sys v0.0.0-20201231184435-2d18734c6014 // indirect
)

replace (
	github.com/gotk3/gotk3 v0.5.2 => github.com/qydysky/gotk3 v0.0.0-20210103171910-327affdaaa80
	github.com/qydysky/bili_danmu => ../
//github.com/qydysky/part => ../../part
)
