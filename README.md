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
INFO: 2020/09/15 06:40:29 [Msg.go>] [pek0pek0 : 外掛]
INFO: 2020/09/15 06:40:30 [Msg.go>] [NealxS : 明显是挂了]
INFO: 2020/09/15 06:40:33 [Msg.go>] [懒得起昵称丶 : 真大哥]
INFO: 2020/09/15 06:40:36 [Msg.go>] [恩里克-普奇-神父 : 这场战役我们失去了天义佬]
INFO: 2020/09/15 06:40:36 [Msg.go>] [ntwww 投喂 1 x 冰阔落 ( 1000 x 金瓜子 )]
INFO: 2020/09/15 06:40:38 [Msg.go>] [方舟之下幽兰呆鹅 : 外挂]
INFO: 2020/09/15 06:40:38 [Msg.go>] [一般通りのまこちゅう : 科技大佬]
^CINFO: 2020/09/15 06:46:14 [ws.go>心跳] [fin]
INFO: 2020/09/15 06:46:14 [ws.go>关闭] [*ws.Close]
INFO: 2020/09/15 06:46:14 [ws.go>关闭] [ok]

ctrl+c退出，日志会同时追加记录到文件danmu.log中
```
更多内容详见注释，如有疑问请发issues，欢迎pr