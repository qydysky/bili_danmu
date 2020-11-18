module github.com/qydysky/bili_danmu/demo

go 1.14

require (
	github.com/gotk3/gotk3 v0.5.0 // indirect
	github.com/klauspost/compress v1.11.2 // indirect
	github.com/miekg/dns v1.1.35 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/qydysky/bili_danmu v0.4.0
	github.com/qydysky/part v0.2.2-0.20201117231944-b989fc77f39b // indirect
	github.com/shirou/gopsutil v3.20.10+incompatible // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/sys v0.0.0-20201110211018-35f3e6cf4a65 // indirect
)

replace (
	github.com/gotk3/gotk3 v0.5.0 => github.com/qydysky/gotk3 v0.5.1-0.20201114200959-3165c4dc990f
	github.com/qydysky/bili_danmu => ../
)
