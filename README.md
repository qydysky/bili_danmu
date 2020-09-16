## bilibili 直播弹幕机

### 当前支持功能列表
```
Msg.go
case 后有函数调用的为支持，无调用的为待完善，注释掉的调用为未启用

var Msg_map = map[string]func(replayF, string) {
	"ANCHOR_LOT_START":nil,//天选之人开始
	"ANCHOR_LOT_CHECKSTATUS":nil,
	"ANCHOR_LOT_END":nil,//天选之人结束
	"ANCHOR_LOT_AWARD":nil,//天选之人获奖
	"COMBO_SEND":nil,
	"INTERACT_WORD":nil,
	"ACTIVITY_BANNER_UPDATE_V2":nil,
	"NOTICE_MSG":nil,
	"ROOM_BANNER":nil,
	"ONLINERANK":nil,
	"WELCOME":nil,
	"HOUR_RANK_AWARDS":nil,
	"ROOM_RANK":nil,
	"ROOM_SHIELD":nil,
	"USER_TOAST_MSG":nil,
	"GUARD_BUY":nil,//大航海购买
	"WELCOME_GUARD":nil,//replayF.welcome_guard,//大航海进入
	"DANMU_MSG":replayF.danmu,//弹幕
	"ROOM_CHANGE":replayF.room_change,//房间信息分区改变
	"ROOM_SILENT_OFF":replayF.roomsilent,//禁言结束
	"ROOM_SILENT_ON":replayF.roomsilent,//禁言开始
	"SEND_GIFT":replayF.send_gift,//礼物
	"ROOM_BLOCK_MSG":replayF.room_block_msg,//封禁
	"PREPARING":replayF.preparing,//下播
	"LIVE":replayF.live,//开播
	"SUPER_CHAT_MESSAGE":nil,//replayF.super_chat_message,//打赏
	"SUPER_CHAT_MESSAGE_JPN":replayF.super_chat_message,//打赏
	"PANEL":replayF.panel,//排行榜
	"ENTRY_EFFECT":replayF.entry_effect,//进入特效
	"ROOM_REAL_TIME_MESSAGE_UPDATE":replayF.roominfo,//粉丝数
}

其他功能
var AllF = map[string]bool{
	"Autoban":true,//自动封禁(仅提示，未完成)
	"Danmuji":true,//反射型弹幕机
	"Danmuji_auto":false,//自动型弹幕机
	"Autoskip":true,//刷屏缩减
}

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