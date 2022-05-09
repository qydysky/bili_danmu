module github.com/qydysky/bili_danmu/demo

go 1.14

require (
	github.com/qydysky/bili_danmu v0.5.9
	github.com/qydysky/part v0.8.0 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
)

replace (
	github.com/gotk3/gotk3 v0.5.2 => github.com/qydysky/gotk3 v0.0.0-20210103171910-327affdaaa80
	github.com/qydysky/bili_danmu => ../
//github.com/qydysky/part => ../../part
)
