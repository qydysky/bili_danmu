## bilibili 直播弹幕机
[![Go Report Card](https://goreportcard.com/badge/github.com/qydysky/bili_danmu)](https://goreportcard.com/report/github.com/qydysky/bili_danmu)
[![codecov](https://codecov.io/gh/qydysky/bili_danmu/branch/master/graph/badge.svg?token=2R7031P47C)](https://codecov.io/gh/qydysky/bili_danmu)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fqydysky%2Fbili_danmu.svg?type=small)](https://app.fossa.com/projects/git%2Bgithub.com%2Fqydysky%2Fbili_danmu?ref=badge_small)
[![少年Pi的直播回放](https://img.shields.io/website?up_message=%E5%9C%A8%E7%BA%BF&down_message=%E7%A6%BB%E7%BA%BF&url=http%3A%2F%2F8.134.76.53%3A8080%2Fweb%2F&label=%E5%B0%91%E5%B9%B4Pi%E7%9A%84%E7%9B%B4%E6%92%AD%E5%9B%9E%E6%94%BE)](https://r.qydysky.org/bili)
[![Static Badge](https://img.shields.io/badge/%E5%B0%91%E5%B9%B4Pi%E7%9A%84%E7%9B%B4%E6%92%AD%E5%9B%9E%E6%94%BEipv6-green)](http://bilirec.qydysky.org:10000/web/)

![commits](https://img.shields.io/github/commit-activity/m/qydysky/bili_danmu)
![GitHub all releases](https://img.shields.io/github/downloads/qydysky/bili_danmu/total)

<a href="https://github.com/qydysky/bili_danmu/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=qydysky/bili_danmu&max=36" />
  <div>see more</div>
</a>

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
以下内容可能过时，点击查看[功能配置](https://github.com/qydysky/bili_danmu/blob/master/demo/config/config_K_v.json)
- [x] 直播流、弹幕回放服务
- [x] 每天自动发送将要过期的银瓜子礼物(默认发送3天内过期的)
- [x] 保持当前已点亮的粉丝牌总是点亮
- [x] 银瓜子自动兑换硬币
- [x] 发进房弹幕(可选有无粉丝牌(可选每日首次发送后不发))
- [x] 每日签到
- [x] 自定义私信
- [x] 自动切换粉丝牌
- [x] 扫码登录(qrcode in webServer and cmd)
- [x] 自定义语音提醒
- [x] GTK弹幕窗
- [x] GTK信息窗
- [x] 营收统计
- [x] 舰长数统计
- [x] 直播流、弹幕保存
- [x] ASS字幕生成
- [x] 节奏提示
- [x] 反射型弹幕机
- [x] 自动型弹幕机
- [x] 相同弹幕合并
- [x] 重复度高弹幕屏蔽
- [x] 弹幕开头字符相同缩减


#### 其他特性

- [x] 使用http,https,socks5代理
- [x] cookie加密
- [x] 弹幕自动重连（30s无响应）
- [x] 直播流开播自动下载、断流再保存、故障转移、移除历史
- [x] 指定弹幕、标题切换重启录制
- [x] 命令行支持房间切换、弹幕发送、表情发送、启停录制、重载弹幕、查看历史记录、查看正在直播的主播、退出房间
- [x] GTK信息窗支持房间切换、弹幕格式化发送、时长统计
- [x] GTK弹幕窗支持自定义人/事件消息停留

### 说明
本项目使用github action自动构建，构建过程详见[yml](https://github.com/qydysky/bili_danmu/blob/master/.github/workflows/go.yml)

#### 获取视频切片
当请求`http://{Web服务地址}{直播Web服务路径}stream/ref={录播文件夹名}&st={起始时间}&dur={片段时长}`时，将返回从录播文件的切片视频(>v0.14.21)

切片将从大于`{起始时间}`的关键帧开始，`{片段时长}`之后的关键帧结束，故大多数情况不能获得精确时间的切片视频

其中`dur`为空时，将返回全部时长。`st`参数可以为空或不传，此时从录播文件起始点开始。

`{起始时间}`、`{片段时长}`格式使用[time.ParseDuration](https://pkg.go.dev/time#ParseDuration)进行转换。例：`1m`为1分钟、`1h2m`为1小时2分。

注意：当配置`直播流回放连接检查`启用时（默认不启用），你需要配置`直播流回放连接检查忽略key`(>v0.14.21)以避免检查，url加上参数`&key={配置的key}`。

例子：

```json
{
  "Web服务地址":"0.0.0.0:20000",
  "直播Web服务路径":"/web/",
  "直播流回放连接检查": 10,
  "直播流回放连接检查忽略key-help": "字符串数组，默认空，空字符串将忽略，当不为空时，将不会定时检查指定key值的请求",
  "直播流回放连接检查忽略key": ["cut"],
}
```

```
curl -v "http://192.168.31.230:20000/web/stream?ref=2024_11_04-01_29_47-47867-250-edd590-JdB&key=cut&dur=1m"
*   Trying 192.168.31.230:20000...
* Connected to 192.168.31.230 (192.168.31.230) port 20000
> GET /web/stream?ref=2024_11_04-01_29_47-47867-250-edd590-JdB&key=cut&dur=1m HTTP/1.1
> Host: 192.168.31.230:20000
> User-Agent: curl/8.9.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Access-Control-Allow-Credentials: true
< Access-Control-Allow-Headers: *
< Access-Control-Allow-Methods: POST, GET, OPTIONS
< Access-Control-Allow-Origin: *
< Connection: keep-alive
< Content-Disposition: inline; filename="2024_11_04-01_29_47-47867-250-edd590-JdB.1731342591.mp4"
< Content-Transfer-Encoding: binary
< Content-Type: flv-application/octet-stream
< Date: Mon, 11 Nov 2024 16:29:51 GMT
< Transfer-Encoding: chunked
< 
Warning: Binary output can mess up your terminal. Use "--output -" to tell curl to output it to your terminal anyway, or consider "--output <FILE>" to save to a file.
* client returned ERROR on write of 1004 bytes
* Failed reading the chunked-encoded stream
* closing connection #0
```

#### Web自定义响应头
配置文件中添加配置项`Web自定义响应头`(>v0.14.19)。默认为空，当不为空时，将在所有响应中添加指定头。
例子：
```
{
  "Web服务地址":"0.0.0.0:11000",
  "直播Web服务路径":"/web/",
  "Web自定义响应头": {
    "Access-Control-Allow-Origin":"*"
  }
}
```

```
curl -i http://{主机名}:11000/web/emots/4e9621239e8a349b1ad198af90458e06.png
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Cache-Control: private
Etag: 2024-09-24T01:53:03+08:00
Last-Modified: Tue, 24 Sep 2024 01:53:03 CST
Date: Sat, 05 Oct 2024 14:02:18 GMT
Content-Type: image/png
Transfer-Encoding: chunked
```


#### ip路径
配置文件中添加配置项`ip路径`(>v0.14.16)。不为空时，将会在指定路径显示出当前主机ip,从而可以使用主机名来获取主机ipv6,以便于[ddns-go](https://github.com/jeessy2/ddns-go)等通过命令获取ipv6。默认空。

例子：
```
{
  "Web服务地址":"0.0.0.0:11000",
  "ip路径":"/ip/"
}
```

```
curl -s http://{主机名}:11000/ip/
192.168.31.230
172.16.0.1
172.17.144.244
172.17.0.1
172.17.0.10
172.17.121.35
2409:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:8ce5
fd55:d4a5:a9e7:0:aab8:e0ff:fe03:8ce5
fe80::aab8:e0ff:fe03:8ce5
fe80::d8dc:35ff:fe21:d71b
fe80::785f:8dff:fe34:9443
fe80::440a:9fff:fe6d:5da6
fe80::60ff:4cff:fee7:c226
fe80::4c24:e3ff:fe65:b955
fe80::8099:ecff:fefa:f36c
fe80::68fe:63ff:fe74:6e35
```

```
curl -s http://{主机名}:11000/ip/ | awk '/240:?/'

2409:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:8ce5
```

#### 保持粉丝牌点亮
配置文件中添加配置项`保持牌子亮着_指定时间`(>v0.14.11)。将会在指定时间启动保持。默认`00:00:00`

保持期间，可能会频繁发送弹幕（间隔5秒/条），此时可能会影响其他使用。

#### 直播回放显示表情
配置文件中添加配置项`弹幕表情`(>v0.14.9)。默认为true，当为true时，将会保存弹幕中的表情png到emots目录下，并在回放时显示表情。

为了能顺利保存，会将某些字符进行转换(<=v0.14.18)，如：[dog?]=>[dog？].png。

但之后(>v0.14.18)存储通过md5作为文件名，如：[dog?]=>md5("[dog?]")+".png"=>0c8427bdb9854d85e9cfe59c45d70583.png

从而避免对原表情产生修改，且避免不同环境对文件名支持的差异。

注意：需要配套更新`demo/html/artPlayer`下的资源

保存的表情可以通过`http://{Web服务地址}{直播Web服务路径}emots/{md5}.png`获取。

#### 直播流停用服务器
配置文件中添加配置项`直播流停用服务器`(>v0.14.3)。默认为空，编写正则字符串，当获取到的服务器链接与字符串匹配时，将会停用。

```json
{
  "直播流不使用mcdn": false,
  "直播流停用服务器": [
      "\\.mcdn\\."
  ],
}
```

#### fmp4相关优化
配置文件中添加配置项`fmp4获取更多服务器`(>v0.14.2)。当直播流类型为fmp4，为`true`时，会在可用切片服务器仅剩1个时，获取更多服务器，但总数超过5个后，将不会再获取。默认为`true`。

配置文件中添加配置项`fmp4跳过解码出错的帧`(>v0.14.2)。当直播流类型为fmp4，为`true`时，会在跳过解码错误的帧(常见于音视频时间戳差过大)，这将导致错误的部分缺失，时间轴发生跳越，但能获得一个完整记录。默认为`false`。

```json
{
  "fmp4获取更多服务器": true,
  "fmp4跳过解码出错的帧-help": "fmp4跳过解码出错的帧、但可能导致关键帧时间上的跳越",
  "fmp4跳过解码出错的帧": false,
}
```

#### 自定义命令行登陆二维码块
配置文件中添加配置项`登陆二维码-白`、`登陆二维码-黑`。

默认：
```json
{
  "登陆二维码-白":"OO",
  "登陆二维码-黑":"  "
}
```

在某些平台，命令行可能被设置为等宽字符，此时默认配置会导致二维码显示比例错误，可以调整到：
```json
{
  "登陆二维码-白":"OOO",
  "登陆二维码-黑":"   "
}
```

或其他字符，或者使用浏览器打开`扫码登录路径`配置的路径进行扫码登陆

#### 指定用户弹幕切片
配置文件中添加配置项`指定弹幕重启录制`。

- `uid`：进入个人空间，可以在地址栏看到你的uid
- `danmu`：正则匹配表达式
- 多个相同`uid`的规则存在时，则会覆盖，取最后一个规则，一个用户最多一条规则
- `uid`或`danmu`为""时，将忽略这个规则

当有`匹配分组`时，第一个分组将作为新标题，例:`t(.*)`，接收到`t暗潮t炒`时，则会将`暗潮t炒`作为标题。

无`匹配分组`时，将以原标题切片，例:`t`，接收到`t暗潮t炒`时，将会切片。

```json
{
  "指定弹幕重启录制":[
      {
          "uid":"29183321",
          "danmu":"t(.*)"
      }
  ]
}
```

#### 指定房间录制区间
配置文件中添加配置项`指定房间录制区间`。

指定roomid的房间在指定时间段内将会开启录制。注意：注意先配置主机时区
 
- `start`时检查是否在直播，是则开始录制，如已在录制则切片。
- `end`时如已在录制则停止录制。
- 存在同时有`start`，`end`的`fromTo`，则开播时，若在`start`与`end`之间，则录制。不存在同时有`start`，`end`的`fromTo`，则一开播就录制。
- 5s内只能触发一个`fromTo`，所以同房间各`start`、`end`之间间隔不要少于5s。

```json
{
  "指定房间录制区间":[
      {
          "roomid":0,
          "fromTo":[
              {
                  "start": "12:01:00",
                  "end": "12:03:00"
              },
              {
                  "start": "12:02:00"
              }
          ]
      }
  ]
}
```

#### 保存日志至DB
配置文件中添加配置项`保存日志至db`。大部分可以参考保存弹幕至db,但有些许不同：

`insert`语句中不能使用`{...}`的占位符，默认占位符1为Prefix,2为Base,3为具体内容。见下面的mysql实例。
```json
{
  "保存日志至db": {
      "dbname": "mysql",
      "url":"root:root@(192.168.31.103:10836)/test",
      "字段help":"$1:Prefix $2:Base $3:Msgs, use ? or $%d not placeholder",
      "create":"create table log (t datetime,Prefix varchar(20), Base varchar(50), Msgs varchar(500))",
      "insert":"insert into log (t,Prefix, Base, Msgs) values (now(),?,?,?)"
  }
}
```

log：

|t|Prefix|Base|Msgs|
|-|------|----|----|
|2023-06-22 04:59:57|T:|Ass|编码: GB18030|
|2023-06-22 04:59:57|T:|功能 更少弹幕|每秒弹幕数: 1|
|2023-06-22 04:59:57|I:|功能 直播Web服务|启动于 http://0.0.0.0:11000/web/|
|2023-06-22 04:59:57|I:|bilidanmu|当前PID: 716582|
|2023-06-22 04:59:57|I:|bilidanmu|version:  20230614221100|
|2023-06-22 04:59:57|I:|bilidanmu|3s内2次ctrl+c退出|
|2023-06-22 04:59:57|I:|bilidanmu|当前ip：http://192.168.31.101|
|2023-06-22 04:59:57|I:|bilidanmu|当前ip：http://172.16.0.1|
|2023-06-22 04:59:57|I:|bilidanmu|当前ip：http://172.17.0.1|
|2023-06-22 04:59:57|I:|bilidanmu|当前ip：http://172.17.0.10|
|2023-06-22 04:59:57|I:|bilidanmu|当前ip：http://172.17.86.21|
|2023-06-22 04:59:58|T:|api IsConnected|已连接|
|2023-06-22 04:59:58|T:|api Get|Get Cookie|
|2023-06-22 04:59:58|I:|api 获取Cookie|已登录|
|2023-06-22 04:59:59|T:|api Get|Get LIVE_BUVID|
|2023-06-22 04:59:59|I:|api LIVE_BUVID|获取到LIVE_BUVID，保存cookie|
|2023-06-22 04:59:59|T:|api Get|Get Uid|
|2023-06-22 04:59:59|T:|命令行操作|回车查看帮助|
|2023-06-22 04:59:59|T:|api 银瓜子=>硬币|现在有银瓜子 396 个|
|2023-06-22 04:59:59|W:|api 银瓜子=>硬币|当前银瓜子数量不足|
|2023-06-22 04:59:59|T:|api 签到|签到|
|2023-06-22 05:00:01|I:|bilidanmu|结束退出|

#### 保存弹幕至DB
配置文件中添加配置项`保存弹幕至db`。参考以下实例：

postgreSql:
```json
{
  "保存弹幕至db": {
      "dbname": "postgres",
      "url":"postgres://postgres:qydysky@192.168.31.103:5432/postgres?sslmode=disable",
      "字段help":"time.Now().Format(time.DateTime), time.Now().Unix(), item.msg, item.color, item.auth, item.uid, item.roomid",
      "create":"create table danmu (created varchar(20), createdunix varchar(20), msg varchar(100), color varchar(20), auth varchar(50), uid varchar(30), roomid varchar(30))",
      "insert":"insert into danmu (created, createdunix, msg, color, auth, uid, roomid) values ({Date},{Unix},{Msg},{Color},{Auth},{Uid},{Roomid})"
  }
}
```

mysql:
```json
{
  "保存弹幕至db": {
      "dbname": "mysql",
      "url":"root:root@(192.168.31.103:10836)/test",
      "字段help":"time.Now().Format(time.DateTime), time.Now().Unix(), item.msg, item.color, item.auth, item.uid, item.roomid",
      "create":"create table danmu (created varchar(20), createdunix varchar(20), msg varchar(100), color varchar(20), auth varchar(50), uid varchar(30), roomid varchar(30))",
      "insert":"insert into danmu (created, createdunix, msg, color, auth, uid, roomid) values ({Date},{Unix},{Msg},{Color},{Auth},{Uid},{Roomid})"
  }
}
```

sqlite3:

注意：不要并发连接数据库，可能会导致失败
```json
{
  "保存弹幕至db": {
      "dbname": "sqlite",
      "url":"danmu.sqlite3",
      "字段help":"time.Now().Format(time.DateTime), time.Now().Unix(), item.msg, item.color, item.auth, item.uid, item.roomid",
      "create":"create table danmu (created text, createdunix text, msg text, color text, auth text, uid text, roomid text)",
      "insert":"insert into danmu  values ({Date},{Unix},{Msg},{Color},{Auth},{Uid},{Roomid})"
  }
}
```

字段`create`可以为空字符串，此时将不会在初次接收到弹幕时执行`create`语句，你必须要预先创建好表。

数据实例：

|created|createdunix|msg|color|auth|uid|roomid|
|-------|-----------|---|-----|----|---|------|
|2023-04-26 03:20:33|1682450433|可能走位配合了他的压枪|#e33fff|畏未|96767379|92613|
|2023-04-26 03:20:45|1682450445|=。=|#54eed8|青江知暖|282800070|92613|
|2023-04-26 03:20:52|1682450452|这不是铁挂吗|#ffffff|瑶少溪|43048863|92613|
|2023-04-26 03:20:55|1682450455|开个箱？|#58c1de|修阡宇|433678545|92613|
|2023-04-26 03:20:59|1682450459|来了来了|#00fffc|我肚子是吃饱了撑大的|830140|92613|


#### 指定房间录制回调
配置文件添加了如下配置
```json
{
    "指定房间录制回调-help":"当指定roomid的房间结束录制后触发对应的命令，命令执行目录为录播目录，占位符({type}:视频类型)，durationS：录制时长超过指定秒数才触发",
    "指定房间录制回调":[
        {
            "roomid":0,
            "durationS":60,
            "after":["cmd","/c","ffmpeg","-i","0.{type}","-y","-c","copy","-movflags","+faststart","1.{type}","1>1.log","2>&1"]
        },
        {
            "roomid":0,
            "durationS":60,
            "after":["ffmpeg","-i","0.{type}","-y","-c","copy","-movflags","+faststart","1.{type}"]
        }
    ]
}
```
上述例子中演示了windows下使用[ffmpeg](https://ffmpeg.org/) ，这将使得保存的流文件`0.mp4 or 0.flv`转为正常的视频`1.mp4 or 1.flv`。

注意：命令运行是异步的，如同步执行多个命令，应使用脚本。

#### 性能检查
当配置了`Web服务地址`及`性能路径`时，运行中的性能信息将可以通过http获取。  
例如有如下配置:  
config_K_v.json
```json
{
  ...
  "Web服务地址":"0.0.0.0:10000",
  "性能路径":"/state/"
  ...
}
```

此时GET http://127.0.0.1:10000/state/
```json
{
  "code": 0,
  "message": "ok",
  "data": {
    "currentTime": "2023-03-11 15:49:06", //当前时间
    "startTime": "2023-03-11 15:48:26",   //启动时间
    "version": "c797128",                 //版本
    "state": {
      "base": {
        "goVersion": "go1.20.1",          //编译使用的golang版本
        "numGoroutine": 53,               //goroutine数量
        "reqPoolState": {
            "inuse": 0,                   //全局请求池-正在使用数量
            "nopooled": 0,
            "nouse": 4,
            "pooled": 4,
            "qts": 1.8,                   //全局请求池-每秒请求数
            "sum": 4                      //全局请求池-总数量
        }
      },
      "common": {
        "92613": {                        //正在录制的房间
            "pid": 9,
            "version": "c797128",
            "live": [                     //流服务器状态>v0.14.5
                {
                    "Host": "xy111x59x162x229xy.mcdn.bilivideo.cn:486", //流服务器HOST
                    "Up": true,                                         //是否启用
                    "Codec": "avc",                                     //编码格式
                    "ReUpTime": "0001-01-01 00:00:00",                  //重新启用时间
                    "Expires": "2078-12-03 11:53:58",
                    "DisableCount": 0                                   //被禁用次数
                },
                {
                    "Host": "d1--cn-gotcha204-4.bilivideo.com",
                    "Up": true,
                    "Codec": "avc",
                    "ReUpTime": "0001-01-01 00:00:00",
                    "Expires": "2078-12-03 11:53:58",
                    "DisableCount": 0
                },
                {
                    "Host": "d1--cn-gotcha209.bilivideo.com",
                    "Up": true,
                    "Codec": "avc",
                    "ReUpTime": "0001-01-01 00:00:00",
                    "Expires": "2078-12-03 11:53:58",
                    "DisableCount": 0
                },
                {
                    "Host": "d1--cn-gotcha208.bilivideo.com",
                    "Up": true,
                    "Codec": "avc",
                    "ReUpTime": "0001-01-01 00:00:00",
                    "Expires": "2078-12-03 11:53:58",
                    "DisableCount": 0
                }
            ],
            "liveQn": 10000,                                //画质
            "roomid": 92613,                                //房间号
            "title": "补作业，没得D了",                     //标题
            "uname": "少年Pi",
            "upUid": 13046,
            "rev": 0,
            "renqi": 1,
            "watched": 36272,
            "onlineNum": 4615,
            "guardNum": 720,
            "parentAreaID": 6,
            "areaID": 235,
            "locked": false,
            "note": "人气榜 100+",
            "liveStartTime": "2024-06-17T15:55:05+08:00",
            "liveing": true
        }
    },
      "gc": {
        "gcAvgS": 6.73,                   //平均gc间隔 单位秒
        "gcCPUFractionPpm": 4.74,         //gc的STC耗时占总CPU时间比值 单位百万分之
        "lastGC": "2023-03-11 15:49:01",  //最后gc时间
        "numGC": 6                        //总gc数量
      },
      "mem": { "memInUse": "7.2 MB" }     //总使用内存
    }
  }
}
```

另外，当配置文件中的`debug路径`不为空时, 访问此路径可以获取调试信息，为空时关闭，需要/结尾，默认为空

- 标准包[net/http/pprof](https://pkg.go.dev/net/http/pprof)将在`debug路径`路径可用，从而可以使用`go tool pprof`工具进行性能调试。

另外，当配置文件中的`debug模式`为`true`时，默认为`false`, 启用额外调试信息

- 当录制fmp4时，每1分钟打印m4s池状态

#### 自定义config_K_v.json
当启动时使用`-ckv 路径`，将从此路径(或http地址)加载config_K_v.json并覆盖默认config_K_v.json中的配置项。
使用此配置，可以在有新配置项时，默认使用新配置项而保持之前其他的配置。

#### 表情发送
当`demo/config/config_K_v.json`中`弹幕_识别表情代码`为`true`时，发送特定的文字将发送表情。

在`demo/config/config_danmu_official.json`中可找到支持的特殊文字

如，命令行输入点赞回车：
```
点赞

I: 2022/09/15 02:23:21 弹幕发送 [发送 official_147 至 92613]
赞
```

danmu.log
```
I: 2022/09/15 02:23:21 弹幕发送 [发送 official_147 至 92613]
I: 2022/09/15 02:23:23 Msg [qydysky丶 : 赞]
```

#### 流保存、弹幕ass、弹幕回放文件
以下内容可能过时，点击查看[配置](https://github.com/qydysky/bili_danmu/blob/master/demo/config/config_K_v.json#L55)

```json
"直播流清晰度-help": "清晰度可选-1:不保存 0:默认 20000:4K 10000:原画 400:蓝光 250:超清 150:高清 80:流畅,无提供所选清晰度时，使用低一档清晰度",
"直播流清晰度": 400,
"直播流类型-help": "flv,fmp4,flvH,fmp4H,带H后缀的为Hevc格式编码",
"直播流类型": "flv",
"直播流保存位置": "./live",
"直播hls流故障转移-help":"true:hls服务器故障时，使用其他",
"直播hls流故障转移": true,
"直播hls流保存为MP4": true,
"仅保存当前直播间流-help": "启用此项，才会保存Ass",
"仅保存当前直播间流": true,
"直播Web服务口":10000,
"直播Web缓冲长度-help":"非负整数，越长直播流延迟越高 内存占用越高",
"直播Web缓冲长度":5,
"直播Web可以发送弹幕":true,
"弹幕回放-help": "仅保存当前直播间流为true时才有效",
"弹幕回放": true,
"ass-help": "只有保存直播流时才考虑生成ass，ass编码默认GB18030(可选utf-8)",
"生成Ass弹幕": true,
"Ass编码": "GB18030",
```

当直播流类型为`fmp4`或`fmp4H`时，默认将下载的切片合并保存为mp4。

当所选类型在当前直播中不可用时，会按以下顺序[`fmp4`,`flv`]尝试。

ass编码GB18030支持中文

- `GB18030`(默认)
- `utf-8`

弹幕回放(仅直播流Web服务)

在保存直播流时，如`弹幕回放`为`true`，则会将会将弹幕同时保存为csv文件。切片或停止录制时，根据csv文件生成[xml文件](https://socialsisteryi.github.io/bilibili-API-collect/docs/danmaku/danmaku_xml.html#xml%E5%BC%B9%E5%B9%95)

在[直播流Web服务](#直播流Web服务)中，点击进入任意一个目录，进行回放，如`弹幕回放`为`true`，则会将会将弹幕回放出来。

添加配置项`弹幕回放_隐藏发送人`(>v0.14.19)，为`true`时，回放弹幕文件类型记录(ass,xml,csv...)不记录发送人，当前直播的即时回放不显示发送人，但不影响`保存弹幕至db`等日志类型记录。由于回放依赖于csv类型文件，故回放历史录播亦不显示发送人。默认为`false`。

当`直播流保存天数`大于1时(默认4)，当t日有1录播时，会尝试删除t-n日及之前的1or2个最早的录播。目录下有.keep文件将忽略(v0.12.7+)。

#### 直播流回放Web服务
启动Web流服务，为下载的直播流提供局域网内的流服务，提供flv、hls/mp4格式流。

在`demo/config/config_K_v.json`中可找到配置项，`0.0.0.0`:开放至局域网 `127.0.0.1`:仅本机。

```
    "Web服务地址":"0.0.0.0:10000",
    "直播Web服务路径":"/web/",
```

下述路径以上面这个配置作为实例。

直接进入直播Web服务路径，`http://127.0.0.1:10000/web/`，将进入由[Vue](https://vuejs.org/)，[element-plus](https://element-plus.org/)创建的简易web目录。

![](_Screenshot/Snipaste_2023-06-28_08-52-30.png)

点击目录项，将进入由[artplayer](https://artplayer.org/)、[mpegts](https://github.com/xqq/mpegts.js)创建的简易回放界面

![](_Screenshot/Snipaste_2023-02-15_22-13-40.png)

**特殊**
- 路径为`http://127.0.0.1:10000/web/stream?ref=now`

  当前正在获取的流，播放此链接时进度将保持当前流进度，在登陆情况下，可以发送弹幕。流格式为fmp4(Hevc)或flv(Hevc)

  使用`?ref={目录}`参数来获取流，当为`?ref=now`时，为当前直播流

- Hevc格式，可能需要硬件解码支持
- 预览[92613少年Pi的直播回放](https://r.qydysky.org/bili)，此预览地址已应用如下配置
  <details>
    
  <summary>config_K_v.json</summary>
    
  ```json
  {
    "ip路径":"/xxxxxxx/",
    "日志文件输出-help": "为空时不保存日志",
    "日志文件输出": "",
    "指定弹幕重启录制":[
        {
            "uid":"29183321",
            "danmu":"1"
        }
    ],
    "弹幕输出到日志": false,
    "直播流清晰度-help": "清晰度可选-1:不保存 0:默认 20000:4K 10000:原画 400:蓝光 250:超清 150:高清 80:流畅,无提供所选清晰度时，使用低一档清晰度",
    "直播流清晰度": 10000,
    "直播流类型-help": "flv,fmp4,flvH,fmp4H,带H后缀的为Hevc格式编码",
    "直播流类型": "fmp4",
    "直播流保存位置": "/bilirec",
    "Web服务地址":"0.0.0.0:10000",
    "直播Web服务路径":"/web/",
    "直播Web可以发送弹幕":false,
    "直播流不使用mcdn":false,
    "fmp4获取更多服务器": true,
    "fmp4跳过解码出错的帧": true,
    "flv使用内置头": true,
    "直播流回放速率": "3 MB",
    "分段时长min": 60,
    "标题修改检测s": 1200,
    "直播流回放连接检查": 100,
    "Web服务连接限制": [
        {
            "cidr":"0.0.0.0/0",
            "max":10
        }
    ],
    "直播流回放连接限制": [
        {
            "help":"内网网段",
            "cidr":"192.168.31.1/24",
            "max":-1
        },
        {
            "help":"公网网段",
            "cidr":"127.0.0.1/32",
            "max":2
        },
        {
            "help":"公网网段",
            "cidr":"::/0",
            "max":2
        }
    ],
    "cookie保护": "下填路径",
    "cookie加密公钥": "public.pem",
    "cookie解密私钥": "private.pem",
    "debug模式-help":"启用额外调试信息",
    "debug模式":true,
    "debug路径-help":"当Web服务地址不为空时, 访问此路径可以获取调试信息，为空时关闭，需要/结尾",
    "debug路径":"/xxxxxxx/",
    "生成pgo":"",
    "扫码登录":true,
    "保存弹幕至db": {
        "dbname": "postgres",
        "url":"postgres://postgres:qydysky@192.168.31.103:5432/postgres?sslmode=disable",
        "字段help":"time.Now().Format(time.DateTime), time.Now().Unix(), item.msg, item.color, item.auth, item.uid, item.roomid",
        "create":"create table danmu (created varchar(20), createdunix varchar(20), msg varchar(100), color varchar(20), auth varchar(50), uid varchar(30), roomid varchar(30))",
        "insert":"insert into danmu (created, createdunix, msg, color, auth, uid, roomid) values ({Date},{Unix},{Msg},{Color},{Auth},{Uid},{Roomid})"
    },
    "保存日志至db":{
        "dbname": "postgres",
        "url":"postgres://postgres:qydysky@192.168.31.103:5432/postgres?sslmode=disable",
        "字段help":"$1:Prefix $2:Base $3:Msgs, use ? or $%d not placeholder",
        "create":"create table log (created timestamp, Prefix varchar(10), Base varchar(30), Msgs varchar(500))",
        "insert":"insert into log (created, Prefix, Base, Msgs) values (now(),$1,$2,$3)"
    },
    "服务器时区":28800,
    "弹幕表情": true
  }
  ```
  </details>

测试可用项目(测试可连续播放10min+)：

- flv-html播放器
  - [xqq/mpegts.js](https://github.com/xqq/mpegts.js)
- hls-html播放器
  - [bytedance/xgplayer](https://github.com/bytedance/xgplayer)
  - [videojs/video.js](https://github.com/videojs/video.js)([demo](https://videojs-http-streaming.netlify.app))
  - [video-dev/hls.js@v1.0.7+](https://hls-js-10780deb-25d8-41d3-b164-bc334c8dd47f.netlify.app/demo/)
- mp4-html播放器
  - [bytedance/xgplayer](https://github.com/bytedance/xgplayer)
  - [videojs/video.js](https://github.com/videojs/video.js)([demo](https://videojs-http-streaming.netlify.app))
- 客户端播放器
  - [mpv](https://mpv.io/)
  - [MXPlayer](https://sites.google.com/site/mxvpen/home)

#### 命令行操作
在准备动作完成(`T: 2021/03/06 16:22:39 命令行操作 [回车查看帮助]`)后，输入回车将显示帮助
```
切换房间->输入' 数字'回车
发送弹幕->输入'字符串'回车
查看直播中主播->输入' liv'回车
重载弹幕->输入' reload'回车
搜索主播->输入' sea关键词'回车
房间信息->输入' room'回车
开始结束录制->输入' rec'回车
退出当前房间->输入' exit'回车
其他输出隔断不影响
```
用例：
- 直播间切换
```
 his
T: 2023/02/01 02:34:00 命令行操作 [指令( his)]

T: 2023/02/01 02:34:00 api 历史直播主播 [获取中]
T: 2023/02/01 02:34:00 api 历史直播主播 [完成]
0       ☁       少年Pi(92613)
                        MC
1               C酱です(47867)
                        阿C今天又在玩什么捏
2               逍遥散人(1017)
                        日常闲聊 麻将上分训练
3               烤鱼子Official(22259479)
                        ！！！！新年快乐！！！！
4       ☁       Asaki大人(6154037)
                        开心鸭鸭杀
5               观察者网(948391)
                        一周军情观察：海的那边是敌人
6               一米八的坤儿(4350043)
                        单排下饭主播
7               数字猫(3839755)
                        石村号上的朋友真是热情好客
8       ☁       环球时报(796281)
                        家乡的风景 家乡的年
9               千鈴鳴_Official(26724557)
                        【B限】原神やるぞおおおお【千鈴鳴】
回复' to(序号)'进入直播间

3 x 9
满脑子都是小c
 to4
T: 2023/02/01 02:34:07 命令行操作 [指令( to4)]

I: 2023/02/01 02:34:07 bilidanmu [停止，等待服务器断开连接]
I: 2023/02/01 02:34:07 bilidanmu [服务器连接关闭]
I: 2023/02/01 02:34:07 直播流保存 [正在等待下载完成...]
T: 2023/02/01 02:34:08 弹幕回放 [停止]
I: 2023/02/01 02:34:08 直播流保存 [结束录制(92613)]
T: 2023/02/01 02:34:09 功能 弹幕合并 [房间更新: 6154037]
T: 2023/02/01 02:34:09 bilidanmu [准备]
T: 2023/02/01 02:34:09 api IsConnected [已连接]
I: 2023/02/01 02:34:10 bilidanmu [连接到房间 6154037]
I: 2023/02/01 02:34:16 api 切换粉丝牌 [无主播粉丝牌]
I: 2023/02/01 02:34:17 bilidanmu [直播中]
T: 2023/02/01 02:34:17 bilidanmu [连接 wss://hw-sh-live-comet-03.chat.bilibili.com/sub]
I: 2023/02/01 02:34:17 bilidanmu [已连接到房间 Asaki大人 ( 6154037 )]
I: 2023/02/01 02:34:17 bilidanmu [开心鸭鸭杀]
T: 2023/02/01 02:34:18 bilidanmu [获取人气]
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
T: 2023/02/01 02:35:38 命令行操作 [指令( room)]

当前直播间(6154037)信息
Asaki大人 开心鸭鸭杀 直播中
已直播时长: 08:12:18
营收: ￥0.00
舰长数: 1287
分区排行: 人气榜 2 人气： 2049772 观看人数： 278848 在线人数： 10000
直播Web服务: http://0.0.0.0:10000
正在录制的房间：
        Asaki大人(6154037) 开心鸭鸭杀
输入` rec` 来启停当前房间录制
```

还支持登录、搜索主播直播间、查看历史记录、查看关注的直播间、保存直播流等功能

#### cookie加密
*使用X25519和chacha20poly1305*(>v0.14.15) 保护cookie.txt

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
使用 -genKey 生成(>v0.14.15)
main(main.exe) -genKey

公钥：
-----BEGIN ECDH PUBLIC KEY-----
tvdVdbI7DTlRcyE44va7zXhi5rewxcm44/Dmp8DMnGY=
-----END ECDH PUBLIC KEY-----

私钥：
-----BEGIN ECDH PRIVATE KEY-----
xrWweTO5upvzDha6WrEBQKkToUYLyMCI7An2btRqop0=
-----END ECDH PRIVATE KEY-----

请复制以上公私钥并另存为文件,可以在cookie加密公钥、cookie解密私钥中使用
```

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

总开关,自定义响应的事件可在`demo/config/config_tts.json`中编辑。{}为传递过来的变量，将会按设定替换。最后未使用的{}会全部删除。下例：`ABC购买 1个月舰长`。
```
...
"0buyguide-help": "大航海 {username}:用户 {op_name}:购买方式 {role_name}:大航海类型 {num}:个数 {unit}:时间单位",
"0buyguide": "{username}{op_name} {num}个{unit}{role_name}",
...
```

特别说明的，下面结构为全局替换。

```
"replace":{
    "?":"问号",
    "？":"问号"
}
```

在`demo/config/config_K_v.json`中可选使用的服务api
```
    "TTS_服务器-help": "baidu:百度翻译合成 youdao:有道TTS",
    "TTS_服务器": "youdao",
    "TTS_服务器_youdaoId": "",
    "TTS_服务器_youdaoKey": "",
```
支持baidu、[有道](https://ai.youdao.com/gw.s#/)、[讯飞](https://www.xfyun.cn/)。

使用有道则需要Id和Key。
```
    "TTS_服务器_youdaoId": "7xxxxxxxxxxxxxxa",
    "TTS_服务器_youdaoKey": "yxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxP",

```

使用讯飞则需要应用Id、Key、Secret。可设置发音人。
```
    "TTS_服务器_xfId": "6xxxxxxb",
    "TTS_服务器_xfKey": "4xxxxxxxx9",
    "TTS_服务器_xfSecret": "YxxxxxxxBh",
    "TTS_服务器_xfVoice-help": "讯飞发音人 xiaoyan:小燕甜美女声 aisjiuxu:许久亲切男声 aisxping:小萍知性女声 aisjinger:小婧亲切女声 aisbabyxu:许小宝可爱童声 random:随机",
    "TTS_服务器_xfVoice": "random",

```

#### 弹幕窗
构建gtk需要gtk3,先行安装[gtk](https://www.gtk.org/)
release 由于某些[原因](https://github.com/golang/go/issues/57328#issuecomment-1353413399)，Linux默认不编译gtk界面 Windows默认不编译
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
1. 预先编译

下载安装[golang](https://go.dev/)/[golangCN](https://golang.google.cn/)

clone/下载本项目。进入`demo`目录(文件夹)，运行：
```
linux: CGO_ENABLED=0 go build .
windows: set CGO_ENABLED=0;go build .
```
再运行生成的`demo.exe`或`demo`

2. 即时编译

下载安装[golang](https://go.dev/)/[golangCN](https://golang.google.cn/)

clone/下载本项目。进入`demo`目录(文件夹)，运行：
```
linux: CGO_ENABLED=0 go run . [-r 房间ID] [-ckv 自定义config_K_v.json] [-genKey]
windows: set CGO_ENABLED=0;go run . [-r 房间ID] [-ckv 自定义config_K_v.json] [-genKey]
```

3. docker部署

经测试可以部署到`debian`镜像上，注意首先得[更新ca](https://stackoverflow.com/questions/64462922/docker-multi-stage-build-go-image-x509-certificate-signed-by-unknown-authorit)
```
apt-get update && apt-get install -y ca-certificates openssl
```

如果你日常使用windows，但在debian运行，那还需先编译linux版本

下载安装[golang](https://go.dev/)/[golangCN](https://golang.google.cn/)

clone/下载本项目。进入`demo`目录(文件夹)，运行：
```
set GOOS=linux
set CGO_ENABLED=0
go build .
```

注意实际使用时，在`config_K_v.json`关闭`tts`等需要关闭界面及音频的功能，配置好录播存放位置，做好直播流服务的端口映射

4. github编译

前往[releases](https://github.com/qydysky/bili_danmu/releases)页下载对应系统版本。解压后进入`demo`目录(文件夹)，运行`main`(`main.exe`)。
```
./main [-r 房间ID] [-ckv 自定义config_K_v.json] [-genKey]
./main.exe [-r 房间ID] [-ckv 自定义config_K_v.json] [-genKey]
```

#### 注意事项
* 其中[]内的内容为可选项
* 当`-r`,`-ckv`为空时，将尝试从环境变量中获取，分别对应键值`r`,`ckv`(>v0.14.21)
* 法1，2，3建议使用最新提交
* 程序受主机时区配置影响，注意正确配置主机时区
* 弹幕及礼物会记录于danmu.log中
* 部分功能(如签到、发送弹幕、获取原画等)**需要在`demo`目录(文件夹)下放置`cookie.txt`才可用** 或 **运行时按提示使用扫码登录成功后才可用(登录信息会保存在`demo/cookie.txt`中)**
* 在golang1.20+，由于某些[原因](https://github.com/golang/go/issues/57328#issuecomment-1353413399)，你可能需要在构建时添加`CGO_ENABLED=0`
* 由于通常是发现功能不正常时，才会检查b站是否更新，又因日常录播并不会使用到全部功能，所以并不能确保全部功能都能正常运行

#### 关于更新版本v0.{y}.{z}
* 当bilibili停止服务时，将会发布v1版本。
* 当发生**不兼容**的改动(删减原来有效、修改有效项的含义)时，将会发布v0.{y+1}.0版。更新前，你需要关注更新说明，以确认变动的配置项的影响。
* 当发生**向下兼容**的改动(修复、新增功能)时，将会发布v0.{y}.{z+1}版。通常只需要覆盖主体`main`（`main.exe`）即可更新。
* 另外`releases`页面中，`Pre-release`表示此版本为尚未经过验证的版本，或存在重大bug，`Latest`表示此版本经过验证，不存在明显的重大bug。

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

#### 结束后的文件播放效果(显于左上)
![](_Screenshot/Screenshot_20200926_173834.png)

[截图地址](//zdir.ntsdtt.bid/ALL/Admin/pack/file/Screenshot_20200926_173834.png)

#### Gtk弹幕窗(Linux Only)

![](_Screenshot/2020-12-12_16-43-09.gif)

[截图地址](//zdir.ntsdtt.bid/ALL/Admin/pack/file/2020-12-12_16-43-09.gif)

![](_Screenshot/Screenshot_20201212_164610.png)

[截图地址](//zdir.ntsdtt.bid//ALL/Admin/pack/file/Screenshot_20201212_164610.png)

更多内容详见注释，如有疑问请发issues，欢迎pr
