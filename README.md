## bilibili 直播弹幕机
golang go version go1.15 linux/amd64

---
### 目录释义
|目录|说明|
|-|-|
|./|项目根|
|CV/|全局变常量|
|F/|项目小工具(ws消息生成、api、整数字节转换)|
|Replay/|接收的数据处理区|
|Send/|发送数据区|
|_Screenshot/|截图保存目录|
|_msg_sample/|ws接收数据示例|
|_source/|bilijs文件示例|
|demo/|运行目录|
|.gitignore|项目忽略文件|
|7za.exe|githubAction的windows打包程序|
|LICENSE|许可|
|VERSION|项目版本|
|bili_danmu.go|主运行文件|
|go.mod|goMod文件|
---

---
### LICENSE
使用了下述的项目，十分感谢
- [golang](https://golang.org/) under [BSD](https://golang.org/LICENSE)
- [github.com/gotk3/gotk3](https://github.com/gotk3/gotk3) under [ISC](https://raw.githubusercontent.com/gotk3/gotk3/master/LICENSE)
- [github.com/qydysky/part](https://github.com/qydysky/part) under [MIT](https://raw.githubusercontent.com/qydysky/part/master/LICENSE)
- [github.com/christopher-dG/go-obs-websocket](https://github.com/christopher-dG/go-obs-websocket) under [MIT](https://raw.githubusercontent.com/christopher-dG/go-obs-websocket/master/LICENSE)
- [github.com/gorilla/websocket](https://github.com/gorilla/websocket) under [BSD 2-Clause](https://raw.githubusercontent.com/gorilla/websocket/master/LICENSE)
- [github.com/skip2/go-qrcode](https://github.com/skip2/go-qrcode) under [MIT](https://github.com/skip2/go-qrcode/blob/master/LICENSE)
- [github.com/gofrs/uuid](https://github.com/gofrs/uuid) under [MIT](https://github.com/gofrs/uuid/blob/master/LICENSE)
- [7z](https://www.7-zip.org/) under [LICENSE](https://www.7-zip.org/license.txt)
---

### 当前支持显示/功能

#### 当前支持显示
以下内容可能过时，点击查看[当前支持显示](https://github.com/qydysky/bili_danmu/blob/master/Reply/Msg.go#L13)
- [x] 人气
- [x] 天选之人开始
- [x] 天选之人获奖
- [x] 直播间关注提示
- [x] 大航海购买
- [x] 节奏风暴
- [x] 大航海进入
- [x] 弹幕
- [x] 房间信息分区改变
- [x] 禁言
- [x] 礼物
- [x] 封禁
- [x] 下播
- [x] 开播
- [x] SC
- [x] 排行榜

#### 当前支持功能
以下内容可能过时，点击查看[当前支持功能](https://github.com/qydysky/bili_danmu/blob/master/Reply/F.go#L16)
- [x] 自定义私信
- [x] 自动切换粉丝牌
- [x] 扫码登录
- [x] 自定义语音提醒
- [x] GTK弹幕窗
- [x] GTK信息窗
- [x] 营收统计
- [x] 舰长数统计
- [x] 直播流保存
- [x] ASS字幕生成
- [x] OBS调用
- [x] 节奏提示
- [x] 反射型弹幕机
- [x] 自动型弹幕机
- [x] 相同弹幕合并
- [x] 重复度高弹幕屏蔽
- [x] 弹幕开头字符相同缩减


#### 其他特性

- [x] 弹幕自动重连（30s）
- [x] 直播流开播自动下载
- [x] 直播流断流再保存
- [x] GTK信息窗支持房间切换、弹幕格式化发送、时长统计
- [x] GTK弹幕窗支持自定义人/事件消息停留

### 说明
本项目使用github action自动构建，构建过程详见[yml](https://github.com/qydysky/bili_danmu/blob/master/.github/workflows/go.yml)

#### 私信
在登录后，可以使用私信

私信配置在`demo/config/config_K_v.json`有说明

#### 语音
调用tts默认使用ffplay,安装[ffmpeg](http://ffmpeg.org/download.html)

或使用其他程序：可在`demo/config/config_K_v.json`中编辑调用的程序及附加选项
```
config_K_v.json
默认
    "TTS_使用程序路径":"ffplay",
    "TTS_使用程序参数":"-autoexit -nodisp"

使用mpv
    "TTS_使用程序路径":"mpv",
    "TTS_使用程序参数":"--no-video"

使用potplayer(例程序位置D:\potplayer\PotPlayerMini64.exe)
    "TTS_使用程序路径":"D:\\potplayer\\PotPlayerMini64.exe",
    "TTS_使用程序参数":"/current /autoplay"
```
release默认编译tts

总开关,自定义响应的事件可在`demo/config/config_tts.json`中编辑
```
{D}:为tts内容
key为demo/face下的文件名
{
    "0multi": "观众：{D}",
    "29183321":"{D}"
}
```
#### 弹幕窗
构建gtk需要gtk3,先行安装[gtk](https://www.gtk.org/)
release Linux默认编译gtk界面 Windows默认不编译
```
编译命令
cd demo
go build -v -tags `gtk` -o demo.exe -i main.go
```
#### 弹幕处理/响应
默认开启了

- 反射弹幕机

启动时加载，当弹幕内容与`demo/config_auto_reply.json`中所设键名相同时，在登录的情况下，会自动发送对应值的弹幕

- 相同合并

当短时间存在大量完全相同的弹幕时，他们将合并显示。

- 更少弹幕

过滤掉自身重复度及最近弹幕重复度高的弹幕

- 更短弹幕

当与上条弹幕具有相同开头的开头时，重复的部分会用...替代

仅对显示效果进行处理，而不处理输出到日志。更多设置见`demo/config/config_F.json`

### demo 
前往[releases](https://github.com/qydysky/bili_danmu/releases)页下载对应系统版本。解压后进入`demo`目录(文件夹)，运行`demo.run`(`demo.exe`)。
```
./demo.run -q 清晰度 -r 房间ID
```

> 清晰度可取[数值](https://github.com/qydysky/bili_danmu/blob/cf52498a88e885fb66dbc94fb8652cb6fa35fb26/CV/Var.go#L37)
> 弹幕及礼物会记录于danmu.log中

**部分功能需要在`demo`目录(文件夹)下放置`cookie.txt`才可用**

### 效果展示
以下内容可能过时，以实际运行为准
- 命令窗口(以下为截取)
```
$ go run main.go 
输入房间号: 213
INFO: 2020/09/16 16:48:11 [bili_danmu.go 测试] [连接到房间 213]
INFO: 2020/09/16 16:48:11 [bili_danmu.go 测试] [连接 wss://tx-sh-live-comet-01.chat.bilibili.com/sub]
INFO: 2020/09/16 16:48:11 [bili_danmu.go 测试] [已连接到房间 213]
INFO: 2020/09/16 16:48:11 [bili_danmu.go 测试] [开始心跳]
```
```
//大航海进入
>>> 欢迎 舰长 茶摊儿在森林喝碗山海 进入直播间
```
```
//普通弹幕
老鸡捉小鹰
你快扒拉他
你这好像是补刀
吓人
```
```
//礼物
====
孤单猫与淋雨猪 投喂 1314 x 辣条 ( 131400 x 金瓜子 )
====
```
```
//同字符串合并
7 x 原神公测B服冲冲冲
```
```
//同字符忽略
原神公测B站冲冲冲
...B服冲冲冲
```
```
//SC
====
SC:  吹舞火 ￥ 30
我旁边的一万是幻觉吗？
私の隣の一万は幻ですか？
====
```
```
//gtk的弹幕格式化发送
2020/11/20 15:39:57 弹幕格式已设置为 [{D}]
INFO: 2020/11/20 15:40:05 [弹幕发送] [发送 [就是这样] 至 394988]
[就是这样]
INFO: 2020/11/20 15:40:15 [弹幕发送] [发送 [你知道么] 至 394988]
[你知道么]
2020/11/20 15:42:38 弹幕长度大于20,不做格式处理
INFO: 2020/11/20 15:42:38 [弹幕发送] [发送 11111111111111111111 至 394988]
11111111111111111111
```
ctrl+c退出，会同时追加记录到文件danmu.log中（文件记录完整信息,不会减少附加功能作用的弹幕）
- 流保存以及弹幕ass
```
结束后会保存为
房间号_时间.mkv
房间号_时间.ass
```
结束后的文件播放效果(显于左上)
![](_Screenshot/Screenshot_20200926_173834.png)

[截图地址](//zdir.ntsdtt.bid/ALL/Admin/pack/file/Screenshot_20200926_173834.png)

- Gtk弹幕窗(Linux Only)

![](_Screenshot/2020-12-12_16-43-09.gif)

[截图地址](//zdir.ntsdtt.bid/ALL/Admin/pack/file/2020-12-12_16-43-09.gif)

![](_Screenshot/Screenshot_20201212_164610.png)

[截图地址](//zdir.ntsdtt.bid//ALL/Admin/pack/file/Screenshot_20201212_164610.png)

更多内容详见注释，如有疑问请发issues，欢迎pr
