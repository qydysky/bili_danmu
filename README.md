## bilibili 直播弹幕机

### 当前支持功能列表
```
Msg.go
case 后有函数调用的为支持，无调用的为待完善，注释掉的调用为未启用

case "COMBO_SEND":
case "INTERACT_WORD":
case "ACTIVITY_BANNER_UPDATE_V2":
case "NOTICE_MSG":
case "ROOM_BANNER":
case "ONLINERANK":
case "WELCOME":
case "HOUR_RANK_AWARDS":
case "ROOM_RANK":
case "ROOM_SHIELD":
case "USER_TOAST_MSG":
case "GUARD_BUY"://大航海购买
case "WELCOME_GUARD"://welcome_guard(s)//大航海进入
case "ROOM_SILENT_OFF", "ROOM_SILENT_ON":roomsilent(s);//禁言
case "SEND_GIFT":send_gift(s)//礼物
case "ROOM_BLOCK_MSG":room_block_msg(s)//封禁
case "PREPARING":preparing(s)//下播
case "LIVE":live(s)//开播
case "SUPER_CHAT_MESSAGE", "SUPER_CHAT_MESSAGE_JPN":super_chat_message(s)//打赏
case "PANEL":panel(s)//排行榜
case "ENTRY_EFFECT":entry_effect(s)//进入特效
case "ROOM_REAL_TIME_MESSAGE_UPDATE":roominfo(s)//粉丝数
case "DANMU_MSG":danmu(s)//弹幕

```

### demo 
```
git clone https://github.com/qydysky/bili_danmu.git
cd demo
go run main.go
```
```
输入房间号: 13946381
INFO: 2020/09/15 06:40:21 [bili_danmu.go>测试] [连接到房间 13946381]
INFO: 2020/09/15 06:40:21 [bili_danmu.go>测试] [连接 wss://tx-sh-live-comet-01.chat.bilibili.com/sub]
INFO: 2020/09/15 06:40:21 [bili_danmu.go>测试] [已连接到房间 13946381]
INFO: 2020/09/15 06:40:22 [bili_danmu.go>测试] [开始心跳]
弹幕
INFO: 2020/09/15 06:40:29 [Msg.go>] [pek0pek0 : 外掛]
弹幕机
INFO: 2020/09/15 14:45:54 [弹幕发送] [发送 在 至 12345]
INFO: 2020/09/15 14:45:55 [弹幕发送] [成功]
INFO: 2020/09/15 14:45:55 [Msg.go>] [12345 : 弹幕机在么]
INFO: 2020/09/15 14:45:56 [Msg.go>] [12345 : 在]
礼物
INFO: 2020/09/15 09:00:26 [Msg.go> 礼] [不能一命通关的M桑 投喂 5 x 冰阔落 ( 5000 x 金瓜子 )]
入场提示
INFO: 2020/09/15 09:00:41 [Msg.go> 房] [欢迎舰长 <%不同选择%> 进入直播间]
排行
INFO: 2020/09/15 09:01:00 [Msg.go> 房] [排行 手游 第4名]
粉丝更新
INFO: 2020/09/15 09:01:00 [Msg.go> 粉] [粉丝总人数: 395189 粉丝团人数: 2391]
。。。

^CINFO: 2020/09/15 09:18:28 [ws.go>关闭] [关闭!]
INFO: 2020/09/15 09:18:28 [bili_danmu.go>测试] [停止，等待服务器断开连接]
INFO: 2020/09/15 09:18:28 [ws.go>处理] [捕获到中断]
INFO: 2020/09/15 09:18:28 [ws.go>心跳] [停止！]
ERROR: 2020/09/15 09:18:28 [ws.go>处理] [服务器意外关闭连接]
INFO: 2020/09/15 09:18:29 [bili_danmu.go>测试] [结束退出]

ctrl+c退出，日志会同时追加记录到文件danmu.log中
```
更多内容详见注释，如有疑问请发issues，欢迎pr