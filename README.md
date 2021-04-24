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
|Json/|json的golang struct表述|
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
- [github.com/skratchdot/open-golang/open](https://github.com/skratchdot/open-golang) under [MIT](https://raw.githubusercontent.com/skratchdot/open-golang/master/LICENSE)
- [7z](https://www.7-zip.org/) under [LICENSE](https://www.7-zip.org/license.txt)
- [github.com/mdp/qrterminal/v3](github.com/mdp/qrterminal/v3) under [MIT](https://github.com/mdp/qrterminal/blob/master/LICENSE)
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
以下内容可能过时，点击查看[功能配置](https://github.com/qydysky/bili_danmu/blob/maintenance/demo/config/config_K_v.json)
- [x] 直播流服务
- [x] 每天自动发送将要过期的银瓜子礼物(默认发送3天内过期的)
- [x] 保持当前已点亮的粉丝牌总是点亮
- [x] 银瓜子自动兑换硬币
- [x] 发进房弹幕(可选有无粉丝牌(可选每日首次发送后不发))
- [x] 每日签到
- [x] 每日有粉丝牌获取小心心
- [x] 自定义私信
- [x] 自动切换粉丝牌
- [x] 扫码登录(qrcode in webServer and cmd)
- [x] 自定义语音提醒
- [x] GTK弹幕窗
- [x] GTK信息窗
- [x] 营收统计
- [x] 舰长数统计
- [x] 直播流保存(默认hls，支持flv)
- [x] ASS字幕生成
- [x] OBS调用
- [x] 节奏提示
- [x] 反射型弹幕机
- [x] 自动型弹幕机
- [x] 相同弹幕合并
- [x] 重复度高弹幕屏蔽
- [x] 弹幕开头字符相同缩减


#### 其他特性

- [x] cookie加密
- [x] 弹幕自动重连（30s）
- [x] 直播流开播自动下载
- [x] 直播流断流再保存
- [x] 命令行支持房间切换、弹幕发送、启停录制
- [x] GTK信息窗支持房间切换、弹幕格式化发送、时长统计
- [x] GTK弹幕窗支持自定义人/事件消息停留

### 说明
本项目使用github action自动构建，构建过程详见[yml](https://github.com/qydysky/bili_danmu/blob/master/.github/workflows/go.yml)

#### 直播流Web服务
启动Web流服务，为下载的直播流提供局域网内的流服务。

在`demo/config/config_K_v.json`中可找到配置项，0:随机可用端口 >0:固定可用端口 <0:禁用服务。

```
    "直播保存位置Web服务":0,
```

开启之后，启动会显示服务地址，在局域网内打开网址可以取得所有直播流的串流地址

支持跨域，注意：在https网站默认无法加载非本机http服务

- dtmp结尾：当前正在获取的流，播放此链接时进度将保持当前流进度
- flv/m3u8结尾：保存完毕的直播流，播放此链接时将从头开始播放
- ass结尾：保存完毕的直播流字幕，有些播放器会在串流时获取此文件
- m4s结尾：hls切片

**特殊的：路径为`/now`(例：当服务地址为下方的38259口时，此对应的路径为`http://192.168.31.245:38259/now`)，会重定向到当前正在获取的流，播放此链接时进度将保持当前流进度**

服务地址也可通过命令行` room`查看。

```
I: 2021/04/13 20:07:45 命令行操作 [直播Web服务: http://192.168.31.245:38259]
```

测试可用项目：

- [bilibili/flv.js](https://github.com/bilibili/flv.js)
- [bytedance/xgplayer](https://github.com/bytedance/xgplayer)
- [video-dev/hls.js](https://github.com/video-dev/hls.js)
- [mpv](https://mpv.io/)


#### 命令行操作
在准备动作完成(`T: 2021/03/06 16:22:39 命令行操作 [回车查看帮助]`)后，输入回车将显示帮助
```
I: 2021/04/01 11:36:46 命令行操作 [切换房间->输入数字回车]
I: 2021/04/01 11:36:46 命令行操作 [发送弹幕->输入' 字符串'回车]
I: 2021/04/01 11:36:46 命令行操作 [房间信息->输入' room'回车]
I: 2021/04/01 11:36:46 命令行操作 [开始结束录制->输入' rec'回车]
I: 2021/04/01 11:36:46 命令行操作 [查看直播中主播->输入' live'回车]
I: 2021/04/01 11:36:46 命令行操作 [其他输出隔断不影响]

```
用例：
- 直播间切换
```
 live
T: 2021/03/06 16:18:48 api 正在直播主播 [获取中]
T: 2021/03/06 16:18:48 api 正在直播主播 [完成]
I: 2021/03/06 16:18:48 命令行操作 [0 奈さま B站第一坑老公]
I: 2021/03/06 16:18:48 命令行操作 [1 古守血遊official 【B】Apex on]
I: 2021/03/06 16:18:48 命令行操作 [2 一米八的坤儿 午饭直播]
I: 2021/03/06 16:18:48 命令行操作 [3 哔哩哔哩英雄联盟赛事 【直播】 LNG vs LGD]
I: 2021/03/06 16:18:48 命令行操作 [回复' live(序号)'进入直播间]
 live1
I: 2021/03/06 16:19:12 命令行操作 [进入房间 8725120]
```
```
120
I: 2021/03/06 16:21:35 命令行操作 [进入 120]
```
- 发送弹幕
```
 1
I: 2021/03/06 16:21:17 弹幕发送 [发送 1 至 7734200]
```
- 查看房间信息
```
 room
I: 2021/04/13 20:04:56 命令行操作 [当前直播间信息]
I: 2021/04/13 20:04:56 命令行操作 [哔哩哔哩英雄联盟赛事 【直播】EDG vs RNG 直播中]
I: 2021/04/13 20:04:56 命令行操作 [已直播时长: 244:02:58]
I: 2021/04/13 20:04:56 命令行操作 [营收: ￥3193.10]
I: 2021/04/13 20:04:56 命令行操作 [舰长数: 16]
I: 2021/04/13 20:04:56 命令行操作 [分区排行: 50+ 人气： 41802746]
I: 2021/04/13 20:04:56 命令行操作 [直播Web服务: http://192.168.31.245:38259]
```
#### cookie加密
保护cookie.txt

在`demo/config/config_K_v.json`中可找到配置项
```
"cookie加密公钥":"public.pem",
"cookie解密私钥":"private.pem"
```
- 当配置了公钥路径后，cookie将被加密(若公钥无效，将会导致cookie无法储存)。若未配置私钥路径，则每次启动都会要求输入私钥路径。(若私钥无效，将会导致cookie被清除)
- 当未配置公钥路径(空字符串)，cookie将明文储存。
- 默认使用了`demo/`下的(public.pem)(private.pem)进行加密，使用时注意自行生成公私钥并按照上述说明使用

注意，每次更换设置(设置或未设置公钥)，cookie会失效。

附：创建公(public.pem)私(private.pem)钥
```
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem
```
#### 小心心
在登录后，可以自动获取小心心，获取小心心需要加密

加密方式：
- 浏览器(默认)

当`小心心nodjs加密服务地址`为空时启用，需要支持webassembly的浏览器(通常可以在bili直播间获得小心心的浏览器均可)  
golang通过websocket与浏览器js进行通讯，在浏览器js调用bilibili的webassembly组件，对信息进行加密。最后返回加密字符串，并由golang进行获取请求。因此需要保持浏览器的相关标签页不被关闭。

- NodeJs

支持使用nodeJs服务来进行加密，在`config/config_K_v.json`配置。当`小心心nodjs加密服务地址`不为空(如Nodejs服务在本地`5200`端口启动：`http://127.0.0.1:5200/enc`)时，将使用此服务来进行加密。注意：加密失败将导致小心心获取退出。  
nodejs小心心加密项目地址[lkeme/bilibili-pcheartbeat](https://github.com/lkeme/bilibili-pcheartbeat)。请自行配置启动。

- golang?暂无

至于为什么没有直接的golang实现，是因为查找资料一番后发现golang执行wasm是使用虚拟机。出于效率及平台普遍性的考量，故没使用，等相关项目更加完善在添加。

相关项目

- [mathetake/gasm](https://github.com/mathetake/gasm)
- [wasmerio/wasmer-go](https://github.com/wasmerio/wasmer-go)

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

### 运行 
#### 方法
1. 前往[releases](https://github.com/qydysky/bili_danmu/releases)页下载对应系统版本。解压后进入`demo`目录(文件夹)，运行`demo.run`(`demo.exe`)。
```
./demo.run [-r 房间ID]
```

2. clone本项目。进入`demo`目录(文件夹)，运行：
```
go run [-tags "gtk"] main.go [-r 房间ID]
```

#### 注意事项
* 其中[]内的内容为可选项
* 法2的golang需1.15并建议使用最新提交
* 弹幕及礼物会记录于danmu.log中
* 部分功能(如获取小心心、签到、发送弹幕、获取原画等)**需要在`demo`目录(文件夹)下放置`cookie.txt`才可用** 或 **运行时按提示使用扫码登录成功后才可用(登录信息会保存在`demo/cookie.txt`中)**

### 效果展示
以下内容可能过时，以实际运行为准

#### 命令窗口(以下为截取)
```
//启动
qydysky@DESKTOP-5CV1EFA:~/程序/git/go/src/github.com/qydysky/bili_danmu/demo$ go run -tags "gtk" main.go -r 21320551
I: 2021/02/18 20:33:09 api 小心心加密 [如需加密，会自动打开 http://127.0.0.1:33673]
I: 2021/02/18 20:33:09 api 小心心加密 [启动]
PID:14544
房间号: 21320551
T: 2021/02/18 20:33:09 api 新建 [ok]
T: 2021/02/18 20:33:09 api 获取房号 [获取房号]
T: 2021/02/18 20:33:10 api LIVE_BUVID [获取LIVE_BUVID]
I: 2021/02/18 20:33:10 api LIVE_BUVID [存在]
T: 2021/02/18 20:33:11 api 获取Token [ok]
I: 2021/02/18 20:33:13 api 获取直播流 [轮播中]
T: 2021/02/18 20:33:13 api 银瓜子=>硬币 [银瓜子=>硬币]
I: 2021/02/18 20:33:15 api 银瓜子=>硬币 [现在有银瓜子 540 个]
W: 2021/02/18 20:33:15 api 银瓜子=>硬币 [当前银瓜子数量不足]
T: 2021/02/18 20:33:15 api 签到 [签到]
I: 2021/02/18 20:33:19 api 获取客户版本 [api version 2.6.25]
I: 2021/02/18 20:33:21 api 获取热门榜 [热门榜: 虚拟主播 50+]
I: 2021/02/18 20:33:23 bilidanmu Demo [连接到房间 21320551]
I: 2021/02/18 20:33:23 bilidanmu Demo [连接 wss://tx-bj-live-comet-02.chat.bilibili.com/sub]
T: 2021/02/18 20:33:23 api 小心心 [获取小心心]
I: 2021/02/18 20:33:23 bilidanmu Demo [已连接到房间 乙女音Official ( 21320551 )]
I: 2021/02/18 20:33:23 bilidanmu Demo [【b限】学《巴啦啦小魔仙》]
I: 2021/02/18 20:33:24 bilidanmu Demo [获取人气]
T: 2021/02/18 20:33:25 api 礼物列表 [获取礼物列表]
I: 2021/02/18 20:33:27 api 获取舰长数 [舰长数获取成功 471]
I: 2021/02/18 20:33:27 弹幕发送 [发送   至 21320551]
I: 2021/02/18 20:33:27 功能 [营收 ￥0.00]
 
I: 2021/02/18 20:33:29 api 礼物列表 [成功]
I: 2021/02/18 20:33:29 api 小心心 [今天小心心已满！]
```
```
//普通弹幕
老鸡捉小鹰
你快扒拉他
你这好像是补刀
吓人
```
```
//大航海
>>> 欢迎舰长 Mana_单推... 进入直播间
```
```
//礼物
====
超级角击 投喂 1 个 摩天轮
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
SC:  凪穗 

有了OTO能量，我们才能够坚强，prprpr
OTOエネルギーがあれば、私たちは強くなれる。prprprpr
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
```
//其他会出现在命令行的信息
//热门榜
I: 2021/02/18 14:59:00 Msg 房 [热门榜 虚拟 4]
//人气
I: 2021/02/18 14:58:51 Reply 人气 [当前人气 450869]
//营收
I: 2021/02/18 14:58:24 功能 [营收 ￥247.80]
//语音
I: 2021/02/18 14:59:00 TTS [0superchat SC:  三得笠·阿克曼茶   8888888  天才oto天才 ]
```
ctrl+c退出，会同时追加记录到文件danmu.log中（文件记录完整信息,不会减少附加功能作用的弹幕）

#### danmu.log
基本同命令行显示，不同下列：
```
//弹幕 上述合并、忽略都不会起作用
I: 2021/02/18 07:24:55 Msg [从天上掉下来的骚年 : 秀才]
//礼物 超过设定限额的将会在命令行中显示，级别为I
T: 2021/02/18 07:30:30 Msg 礼 [正道的光博航同志 投喂 1 个 上上签 ￥1.0]
I: 2021/02/18 14:52:31 Msg 礼 [三千千千千千千 投喂 1 个 爱之魔力 ￥28.0]
//sc
I: 2021/02/18 14:40:54 Msg 礼 [SC:  加拉入我心 ￥ 30 关注了乙女音，我才能够得到快乐 乙女音符に注目してこそ、私は幸せになれるのです。]
I: 2021/02/18 21:48:49 Msg 房 [欢迎舰长 Mana_单推... 进入直播间]
```

#### 流保存以及弹幕ass
```
结束后会保存为
房间号_时间.mkv
房间号_时间.ass
```

#### 结束后的文件播放效果(显于左上)
![](_Screenshot/Screenshot_20200926_173834.png)

[截图地址](//zdir.ntsdtt.bid/ALL/Admin/pack/file/Screenshot_20200926_173834.png)

#### Gtk弹幕窗(Linux Only)

![](_Screenshot/2020-12-12_16-43-09.gif)

[截图地址](//zdir.ntsdtt.bid/ALL/Admin/pack/file/2020-12-12_16-43-09.gif)

![](_Screenshot/Screenshot_20201212_164610.png)

[截图地址](//zdir.ntsdtt.bid//ALL/Admin/pack/file/Screenshot_20201212_164610.png)

更多内容详见注释，如有疑问请发issues，欢迎pr
