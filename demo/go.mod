module github.com/qydysky/bili_danmu/demo

go 1.14

require (
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/christopher-dG/go-obs-websocket v0.0.0-20200720193653-c4fed10356a5 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/gotk3/gotk3 v0.6.0 // indirect
	github.com/klauspost/compress v1.13.0 // indirect
	github.com/mdp/qrterminal/v3 v3.0.0 // indirect
	github.com/miekg/dns v1.1.42 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/qydysky/bili_danmu v0.5.9
	github.com/qydysky/part v0.5.23 // indirect
	github.com/shirou/gopsutil v3.21.5+incompatible // indirect
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e // indirect
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966 // indirect
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210608053332-aa57babbf139 // indirect
)

replace (
	github.com/gotk3/gotk3 v0.5.2 => github.com/qydysky/gotk3 v0.0.0-20210103171910-327affdaaa80
	github.com/qydysky/bili_danmu => ../
//github.com/qydysky/part => ../../part
)
