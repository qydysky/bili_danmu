## bilibili 直播弹幕机

### 当前支持显示/功能
```
显示
case 后有函数调用的为支持，为nil的为待完善，注释掉的调用为未启用

var Msg_map = map[string]func(replayF, string) {
	"ROOM_SKIN_MSG":nil,
	"GUARD_ACHIEVEMENT_ROOM":nil,
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
	"WIN_ACTIVITY":nil,
	"GUARD_BUY":replayF.guard_buy,//大航海购买
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
	"ROOM_REAL_TIME_MESSAGE_UPDATE":nil,//replayF.roominfo,//粉丝数
}

其他功能
//功能开关
var AllF = map[string]bool{
	"Autoban":false,//自动封禁(仅提示，未完成)
	"Danmuji":true,//反射型弹幕机，回应弹幕
	"Danmuji_auto":false,//自动型弹幕机，定时输出
	"Autoskip":true,//刷屏缩减，相同合并
	"Lessdanmu":true,//弹幕缩减，显示差异大的
}

```

### demo 
```
git clone https://github.com/qydysky/bili_danmu.git
cd demo
go run main.go
```
```
$ go run main.go 
输入房间号: 213
INFO: 2020/09/16 16:48:11 [bili_danmu.go 测试] [连接到房间 213]
INFO: 2020/09/16 16:48:11 [bili_danmu.go 测试] [连接 wss://tx-sh-live-comet-01.chat.bilibili.com/sub]
INFO: 2020/09/16 16:48:11 [bili_danmu.go 测试] [已连接到房间 213]
INFO: 2020/09/16 16:48:11 [bili_danmu.go 测试] [开始心跳]
欢迎舰长 <%初鑫不变%> 进入直播间
C语言get DAZE！ + 10
老鸡捉小鹰
你快扒拉他
你这好像是补刀
吓人
====
孤单猫与淋雨猪 投喂 1314 x 辣条 ( 131400 x 金瓜子 )
====

ctrl+c退出，日志会同时追加记录到文件danmu.log中（文件记录完整信息）
```
更多内容详见注释，如有疑问请发issues，欢迎pr