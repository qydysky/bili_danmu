## bilibili 直播弹幕机

### demo 
```
git clone https://github.com/qydysky/bili_danmu.git
cd demo
go run main.go
```
```
输入房间号: 213 
INFO:  2020/09/14 21:02:58 [send hello to wss://tx-sh-live-comet-02.chat.bilibili.com/sub]
ERROR:  2020/09/14 21:02:58 [-> websocket: close 1006 (abnormal closure): unexpected EOF]
INFO:  2020/09/14 21:02:59 [send hello to wss://tx-gz-live-comet-01.chat.bilibili.com/sub]
ERROR:  2020/09/14 21:02:59 [-> websocket: close 1006 (abnormal closure): unexpected EOF]
INFO:  2020/09/14 21:03:00 [send hello to wss://broadcastlv.chat.bilibili.com/sub]
INFO:  2020/09/14 21:03:00 [wss://broadcastlv.chat.bilibili.com/sub hello!]
INFO:  2020/09/14 21:06:56 [祸美人M : 鬼来了]
INFO:  2020/09/14 21:06:56 [泽北SAMAい : wc]
INFO:  2020/09/14 21:06:57 [OurTube丶Now : 偷窥]
INFO:  2020/09/14 21:06:57 [云又云又云 : 什么游戏？]
INFO:  2020/09/14 21:06:57 [GNsetusna : 草]
INFO:  2020/09/14 21:06:58 [amuseustillwedie : 闸总你出来！]
INFO:  2020/09/14 21:06:58 [Kinoko7pro : 游戏名：零~月蚀之假面~]
INFO:  2020/09/14 21:07:00 [mousebat04 : 就是能登麻美子]
INFO:  2020/09/14 21:07:02 [紅魔の月時計 : 毫无感觉]
INFO:  2020/09/14 21:07:03 [0小牙0 : 想到了冲自己冲冲冲]
INFO:  2020/09/14 21:07:03 [gamestarts0 : 拿到武器必遇怪]
INFO:  2020/09/14 21:07:04 [为什么鸽 : 我敲]
INFO:  2020/09/14 21:07:04 [欢迎舰长 <%林嘉驹%> 进入直播间]
INFO:  2020/09/14 21:07:04 [Icarus丶 : 这不就是小圆]
INFO:  2020/09/14 21:07:05 [伶伶miss : 妖孽！哪里逃]
INFO:  2020/09/14 21:07:07 [阿克酱 : 胆子好大啊]
INFO:  2020/09/14 21:07:08 [Fdalmir : 偷窥]
INFO:  2020/09/14 21:07:08 [欣泠丶 : 那个是圆香吧]
INFO:  2020/09/14 21:07:08 [[[粉丝总人数: 599214] [粉丝团人数: 13796]]]
INFO:  2020/09/14 21:07:09 [_墨轩- : 要是我直接哭了]
INFO:  2020/09/14 21:07:09 [百合花业余植物学家 : 要是等流歌来可能就不会全灭了]
INFO:  2020/09/14 21:07:09 [单机 第70名]
INFO:  2020/09/14 21:07:11 [红脸的野郎 : 小姐姐；爱玩啊]
INFO:  2020/09/14 21:07:11 [三月三日三重樱 : 大威天龙！]

ctrl+c退出，弹幕会同时记录到文件danmu.log中
```
更多内容详见注释，如有疑问请发issues