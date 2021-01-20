//+build gtk

package reply

import (
	"container/list"
	"errors"
	"strconv"
	"time"
	"strings"
	"log"
	"fmt"
	"sync"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/gdk"
	p "github.com/qydysky/part"
	F "github.com/qydysky/bili_danmu/F"
	c "github.com/qydysky/bili_danmu/CV"
	s "github.com/qydysky/part/buf"
)

type danmu_item struct {
	text *gtk.TextView
	img *gtk.Image
	handle glib.SignalHandle
}
type gtk_item_source struct {
	text string
	img string
	time time.Time
}

var gtkGetList = make(chan string,100)

var keep_key = map[string]int{
	"face/0default":0,
	"face/0room":0,
	"face/0buyguide":9,
	"face/0gift":8,
	"face/0jiezou":8,
	"face/0level1":5,
	"face/0level2":3,
	"face/0level3":1,
	"face/0superchat":13,
	"face/0tianxuan":5,
}
var (
	Gtk_on bool
	Gtk_img_path string = "face"
	Gtk_danmu_pool_index uint
	Gtk_danmu_pool = make(map[uint]Danmu_mq_t)
	win *gtk.Window
	imgbuf = struct{
		b map[string](*gdk.Pixbuf)
		sync.Mutex
	}{
		b:make(map[string](*gdk.Pixbuf)),
	}

	danmu_win_running bool//弹幕窗体是否正在运行
	contrl_win_running bool//控制窗体是否正在运行
	in_smooth_roll bool
	grid0 *gtk.Grid
	grid1 *gtk.Grid
	keep_list = list.New()
)

func init(){
	if!IsOn("Gtk") {return}
	go Gtk_danmu()
	//使用带tag的消息队列在功能间传递消息
	Danmu_mq.Pull_tag(map[string]func(interface{})(bool){
		`danmu`:func(data interface{})(bool){//弹幕
			if int(Gtk_danmu_pool_index) - len(Gtk_danmu_pool) > 1e5 {Gtk_danmu_pool_index = 0}
			Gtk_danmu_pool_index += 1
			Gtk_danmu_pool[Gtk_danmu_pool_index] = data.(Danmu_mq_t)
			return false
		},
	})
	//
	go func(){//copy map
		for {
			time.Sleep(time.Duration(60)*time.Second)
			{
				tmp := make(map[string](*gdk.Pixbuf))
				for k,v := range imgbuf.b {tmp[k] = v}
				imgbuf.Lock()
				imgbuf.b = tmp
				imgbuf.Unlock()
			}
		}
	}()
}

func Gtk_danmu() {
	if Gtk_on {return}
	gtk.Init(nil)

	var (
		win2 *gtk.Window
		scrolledwindow0 *gtk.ScrolledWindow
		viewport0 *gtk.Viewport
		viewport1 *gtk.Viewport
		w2_textView0 *gtk.TextView
		w2_textView1 *gtk.TextView
		w2_textView2 *gtk.TextView
		w2_textView3 *gtk.TextView
		w2_textView4 *gtk.TextView
		renqi_old = 1
		w2_Entry0 *gtk.Entry
		w2_Entry0_editting = make(chan bool,10)
	)


	application, err := gtk.ApplicationNew(
	"com.github.qydysky.bili_danmu.reply"+p.Sys().GetTime(),//时间戳允许多开
	glib.APPLICATION_FLAGS_NONE)

	if err != nil {log.Println(err);return}

	application.Connect("startup", func() {

		builder, err := gtk.BuilderNewFromFile("ui/1.glade")
		if err != nil {log.Println(err);return}
		builder2, err := gtk.BuilderNewFromFile("ui/2.glade")
		if err != nil {log.Println(err);return}

		{
			signals := map[string]interface{}{
				"on_main_window_destroy": onMainWindowDestroy,
			}
			builder.ConnectSignals(signals)
			builder2.ConnectSignals(signals)
		}

		{
			obj, err := builder.GetObject("main_window")
			if err != nil {log.Println(err);return}
			win, err = isWindow(obj)
			if err != nil {log.Println(err);return}
			danmu_win_running = true
			win.Connect("delete-event", func() {
				log.Println(`弹幕窗已关闭`)
				danmu_win_running = false//关闭后置空
			})
			application.AddWindow(win)
		}
		{
			obj, err := builder2.GetObject("main_window")
			if err != nil {log.Println(err);return}
			win2, err = isWindow(obj)
			if err != nil {log.Println(err);return}
			contrl_win_running = true
			win2.Connect("delete-event", func() {
				log.Println(`弹幕信息窗已关闭`)
				contrl_win_running = false//关闭后置空
			})
			application.AddWindow(win2)
		}
		{//营收
			obj, err := builder2.GetObject("t0")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.TextView); ok {
				w2_textView0 = tmp
			}else{log.Println("cant find #t0 in .glade");return}
		}
		{//直播时长
			obj, err := builder2.GetObject("t1")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.TextView); ok {
				w2_textView1 = tmp
			}else{log.Println("cant find #t1 in .glade");return}
		}
		{//人气值
			obj, err := builder2.GetObject("t2")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.TextView); ok {
				w2_textView2 = tmp
			}else{log.Println("cant find #t2 in .glade");return}
		}
		{//舰长数
			obj, err := builder2.GetObject("t3")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.TextView); ok {
				w2_textView3 = tmp
			}else{log.Println("cant find #t3 in .glade");return}
		}
		{//排名
			obj, err := builder2.GetObject("t4")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.TextView); ok {
				w2_textView4 = tmp
			}else{log.Println("cant find #t4 in .glade");return}
		}
		{//发送弹幕
			var danmu_send_form string
			{//发送弹幕格式
				obj, err := builder2.GetObject("send_danmu_form")
				if err != nil {log.Println(err);return}
				if tmp,ok := obj.(*gtk.Entry); ok {
					tmp.Connect("focus-out-event", func() {
						if t,e := tmp.GetText();e == nil {//可设置为空
							danmu_send_form = t
							log.Println("弹幕格式已设置为",danmu_send_form)
						}
					})
				}else{log.Println("cant find #send_danmu in .glade");return}
			}
			obj, err := builder2.GetObject("send_danmu")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.Entry); ok {
				tmp.Connect("key-release-event", func(entry *gtk.Entry, event *gdk.Event) {
					eventKey := gdk.EventKeyNewFromEvent(event)
					if eventKey.KeyVal() == gdk.KEY_Return {
						if t,e := entry.GetText();e == nil && t != ``{
							danmu_want_send := t
							if danmu_send_form != `` {danmu_want_send = strings.ReplaceAll(danmu_send_form, "{D}", t)}
							if len([]rune(danmu_want_send)) > 20 {
								log.Println(`弹幕长度大于20,不做格式处理`)
								danmu_want_send = t
							} 
							Msg_senddanmu(danmu_want_send)
							entry.SetText(``)
						}
					}
				})
			}else{log.Println("cant find #send_danmu in .glade");return}
		}
		{//房间id
			obj, err := builder2.GetObject("want_room_id")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.Entry); ok {
				w2_Entry0 = tmp
				tmp.Connect("focus-out-event", func() {
					glib.TimeoutAdd(uint(3000), func()bool{//3s后才解除，避免刚想切换又变回去
						for len(w2_Entry0_editting) != 0 {<-w2_Entry0_editting}
						w2_Entry0_editting <- false
						return false
					})
				})
			}else{log.Println("cant find #want_room_id in .glade");return}
		}
		{//房间id click
			obj, err := builder2.GetObject("want_click")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.Button); ok {
				tmp.Connect("clicked", func() {
					if t,e := w2_Entry0.GetText();e != nil {
						show("读取错误",load_face("0room"))
					} else if t != `` {
						if i,e := strconv.Atoi(t);e != nil {
							show(`输入错误`,load_face("0room"))
						} else {
							c.Roomid =  i
							c.Danmu_Main_mq.Push_tag(`change_room`,nil)
						}
					} else {
						show(`房间号输入为空`,load_face("0room"))
					}
				})
			}else{log.Println("cant find #want_click in .glade");return}
		}
		{
			obj, err := builder.GetObject("scrolledwindow0")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.ScrolledWindow); ok {
				scrolledwindow0 = tmp
			}else{log.Println("cant find #scrolledwindow0 in .glade");return}
		}

		{
			obj, err := builder.GetObject("viewport0")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.Viewport); ok {
				viewport0 = tmp
			}else{log.Println("cant find #viewport0 in .glade");return}
		}
		{
			obj, err := builder.GetObject("viewport1")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.Viewport); ok {
				viewport1 = tmp
			}else{log.Println("cant find #viewport1 in .glade");return}
		}
		{
			obj, err := builder.GetObject("grid0")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.Grid); ok {
				grid0 = tmp
			}else{log.Println("cant find #grid0 in .glade");return}
		}
		{
			obj, err := builder.GetObject("grid1")
			if err != nil {log.Println(err);return}
			if tmp,ok := obj.(*gtk.Grid); ok {
				grid1 = tmp
			}else{log.Println("cant find #grid1 in .glade");return}
		}
		imgbuf.Lock()
		imgbuf.b["face/0default"],_ = gdk.PixbufNewFromFileAtSize("face/0default", 40, 40);
		imgbuf.Unlock()

		{
			if pro_style,e := gtk.CssProviderNew();e == nil{
				if e = pro_style.LoadFromPath(`ui/1.css`);e == nil{
					if scr := win.GetScreen();scr != nil {
						gtk.AddProviderForScreen(scr,pro_style,gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
					}
				}else{log.Println(e)}
			}else{log.Println(e)}
		}

		//先展示弹幕信息窗
		win.ShowAll()
		win2.ShowAll()

		Gtk_on = true
	})

	application.Connect("activate", func() {

		go func(){
			for danmu_win_running {
				time.Sleep(time.Second)
				glib.TimeoutAdd(uint(10),func()(value bool){
					el := keep_list.Front()
					value = el != nil && time.Now().After(el.Value.(gtk_item_source).time)
					if value {
						if i,e := grid1.GetChildAt(0,0); e != nil{i.(*gtk.Widget).Destroy()}
						if i,e := grid1.GetChildAt(1,0); e != nil{i.(*gtk.Widget).Destroy()}
						grid1.RemoveRow(0)
						show(el.Value.(gtk_item_source).text,el.Value.(gtk_item_source).img, 0)
						keep_list.Remove(el)
						el = el.Next()
					}
					return
				})
				glib.TimeoutAdd(uint(1000 / (len(Gtk_danmu_pool) + 1)),func()(bool){
					for id,item := range Gtk_danmu_pool {
						delete(Gtk_danmu_pool,id)
						show(item.msg,load_face(item.uid))
						return true
					}
					return false
				})
			}
		}()
		var old_cu float64
		{//平滑滚动效果
			tmp := scrolledwindow0.GetVAdjustment()
			glib.TimeoutAdd(uint(30),func()(true_value bool){
				true_value = true
				if !in_smooth_roll {return}
				h := viewport0.GetViewWindow().WindowGetHeight()
				// g1h := viewport1.GetViewWindow().WindowGetHeight()
				max := tmp.GetUpper() - float64(h)
				cu := tmp.GetValue()
				
				//用户在回看
				if old_cu != 0 &&//非初始
				max - cu > 100 &&//当前位置低于100
				old_cu != cu {//上一次滚动有移动
					return
				}

				step := (max - cu) / 30
				if step > 20 || max > 5 * float64(h){//太长或太快
					if i,e := grid0.GetChildAt(0,0); e != nil{i.(*gtk.Widget).Destroy()}
					if i,e := grid0.GetChildAt(1,0); e != nil{i.(*gtk.Widget).Destroy()}
					grid0.RemoveRow(0)
				} else if step > 0.5 {
					if step > 5{step = 5}
					tmp.SetValue(cu + step)
				} else {
					in_smooth_roll = false
					tmp.SetValue(max)
					loc := int(grid0.Container.GetChildren().Length())/2
					if v,ok := K_v[`gtk_保留弹幕数量`].(int);ok {
						loc -= v
					} else {
						loc -= 25
					}
					for loc > 0 {
						if i,e := grid0.GetChildAt(0,0); e != nil{i.(*gtk.Widget).Destroy()}
						if i,e := grid0.GetChildAt(1,0); e != nil{i.(*gtk.Widget).Destroy()}
						grid0.RemoveRow(0)
						loc -= 1
					}
				}
				old_cu = tmp.GetValue()
				return
			})
		}
		glib.TimeoutAdd(uint(3000), func()(o bool){
			o = contrl_win_running
			//y("sssss",load_face(""))
			{//加载特定信息驻留时长
				buf := s.New()
				buf.Load("config/config_gtk_keep_key.json")
				for k,_ := range keep_key {delete(keep_key,k)}
				for k,v := range buf.B {
					keep_key[k] = int(v.(float64))
				}
			}
			{//营收
				if IsOn("ShowRev") {
					b,e := w2_textView0.GetBuffer()
					if e != nil {log.Println(e);return}
					b.SetText(fmt.Sprintf("￥%.2f",c.Rev))					
				}
			}
			{//舰长
				b,e := w2_textView3.GetBuffer()
				if e != nil {log.Println(e);return}
				b.SetText(fmt.Sprintf("%d",c.GuardNum))
			}
			{//分区排行
				b,e := w2_textView4.GetBuffer()
				if e != nil {log.Println(e);return}
				b.SetText(c.Note)
			}
			{//时长
				b,e := w2_textView1.GetBuffer()
				if e != nil {log.Println(e);return}
				if c.Liveing {
					d := time.Since(c.Live_Start_Time).Round(time.Second)
					h := d / time.Hour
					d -= h * time.Hour
					m := d / time.Minute
					d -= m * time.Minute
					s := d / time.Second
					b.SetText(fmt.Sprintf("%02d:%02d:%02d", h, m, s))
				} else {
					b.SetText("00:00:00")
				}
			}
			{//人气
				b,e := w2_textView2.GetBuffer()
				if e != nil {log.Println(e);return}
				if c.Liveing {
					if c.Renqi != renqi_old {
						var Renqi string = strconv.Itoa(c.Renqi)
						L:=len([]rune(Renqi))

						var tmp string
						if renqi_old != 1 {
							if c.Renqi > renqi_old {tmp += `+`}
							tmp += fmt.Sprintf("%.1f",100*float64(c.Renqi - renqi_old)/float64(renqi_old)) + `% | `
						}
						if c.Renqi != 0 {renqi_old = c.Renqi}

						for k,v := range []rune(Renqi) {
							tmp += string(v)
							if (L - k)%3 == 1 && L - k != 1{
								tmp += `,`
							}
						}
						b.SetText(tmp)
					}
				} else {
					b.SetText(`0`)
				}
			}
			{//房间id
				for len(w2_Entry0_editting) > 1 {<-w2_Entry0_editting}
				select{
				case tmp:=<-w2_Entry0_editting:
					if !tmp {
						w2_Entry0.SetText(strconv.Itoa(c.Roomid))
					}
				default:
				}
			}
			select {
			case uid:=<-gtkGetList:
				go func(){
					if p.Checkfile().IsExist(Gtk_img_path + `/` + uid) {return}
					src := F.Get_face_src(uid)
					if src == "" {return}
					req := p.Req()
					if e := req.Reqf(p.Rval{
						Url:src,
						SaveToPath:Gtk_img_path + `/` + uid,
						Timeout:3,
					}); e != nil{log.Println(e);}
				}()
			default:
			}
			return
		})
	})

	application.Connect("shutdown", func() {
		log.Println("application shutdown")	
		Gtk_on = false
		c.Danmu_Main_mq.Push_tag(`gtk_close`,nil)
	})

	application.Run(nil)
}

func isWindow(obj glib.IObject) (*gtk.Window, error) {
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}

func load_face(uid string) (loc string) {
	loc = Gtk_img_path + `/` + "0default"
	if uid == "" {return}
	if _,ok := keep_key[Gtk_img_path + `/` + uid];ok{
		loc = Gtk_img_path + `/` + uid
		return
	}
	if p.Checkfile().IsExist(Gtk_img_path + `/` + uid) && p.Rand().MixRandom(1,100) > 1 {
		loc = Gtk_img_path + `/` + uid
		return
	}
	if v,ok := K_v[`gtk_头像获取等待最大数量`].(int);ok && len(gtkGetList) > v {return}
	select{
		case gtkGetList <- uid:
		default:
	}
	return
}

func show(s,img_src string,to_grid ...int){
	sec := 0

	var item danmu_item

	item.text,_ = gtk.TextViewNew();
	{
		item.text.SetMarginStart(5)
		item.text.SetEditable(false)
		item.text.SetHExpand(true)
		item.text.SetWrapMode(gtk.WRAP_WORD_CHAR)
		if tsec,ok := keep_key[img_src];ok && tsec != 0 {
			sec = tsec
			if sty,e := item.text.GetStyleContext();e == nil{
				sty.AddClass("highlight")
			}
		}
		item.handle,_ = item.text.Connect("size-allocate", func(){
			b,e := item.text.GetBuffer()
			if e != nil {log.Println(e);return}
			b.SetText(s)
			in_smooth_roll = true
		})
	}

	item.img,_ = gtk.ImageNew();
	{
		var (
			pixbuf *gdk.Pixbuf
			e error
		)
		if v,ok := imgbuf.b[img_src];ok{
			pixbuf,e = gdk.PixbufCopy(v)
		} else {
			pixbuf,e = gdk.PixbufNewFromFileAtSize(img_src, 40, 40);
			if e == nil {
				imgbuf.Lock()
				if v,ok := K_v[`gtk_内存头像数量`].(int);ok && len(imgbuf.b) > v {
					for k,_ := range imgbuf.b {delete(imgbuf.b,k);break}
				}
				imgbuf.b[img_src],e = gdk.PixbufCopy(pixbuf)
				imgbuf.Unlock()
			}
		}
		if e == nil {item.img.SetFromPixbuf(pixbuf)}
	}
	{
		if len(to_grid) != 0 && to_grid[0] == 0 {//突出显示结束后，显示在普通弹幕区
			loc := int(grid0.Container.GetChildren().Length())/2;
			grid0.InsertRow(loc);
			grid0.Attach(item.img, 0, loc, 1, 1)
			grid0.Attach(item.text, 1, loc, 1, 1)
			grid0.ShowAll()
			return
		}
		/*
			front
			|
			back index:0
		*/
		var InsertIndex int = keep_list.Len()
		if sec > InsertIndex / 5 {//5不是指最大值，而是当list太大时，sec小的将直接跳过
			var cu_To = time.Now().Add(time.Second * time.Duration(sec))
			var hasInsert bool
			for el := keep_list.Front(); el != nil; el = el.Next(){
				if cu_To.After(el.Value.(gtk_item_source).time) {InsertIndex -= 1;continue}
				keep_list.InsertBefore(gtk_item_source{
					text:s,
					img:img_src,
					time:cu_To,
				},el)
				hasInsert = true
				break
			}
			if !hasInsert {
				keep_list.PushBack(gtk_item_source{
					text:s,
					img:img_src,
					time:cu_To,
				})
			}
			loc := int(grid1.Container.GetChildren().Length())/2;
			grid1.InsertRow(loc - InsertIndex);
			grid1.Attach(item.img, 0, loc - InsertIndex, 1, 1)
			grid1.Attach(item.text, 1, loc - InsertIndex, 1, 1)
			grid1.ShowAll()
		} else {
			loc := int(grid0.Container.GetChildren().Length())/2;
			grid0.InsertRow(loc);
			grid0.Attach(item.img, 0, loc, 1, 1)
			grid0.Attach(item.text, 1, loc, 1, 1)
			grid0.ShowAll()
		}
	}

}
