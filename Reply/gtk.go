package reply

import (
	"container/list"
	"errors"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/gdk"
	p "github.com/qydysky/part"
	F "github.com/qydysky/bili_danmu/F"
)

const appId = "com.github.qydysky.bili_danmu.reply"
var BTList = list.New()
var BIList = list.New()

var (
	Gtk_on bool
	Gtk_Tra bool
	Gtk_img_path string = "face"
	Gtk_danmuChan chan string = make(chan string, 100)
	Gtk_danmuChan_uid chan string = make(chan string, 100)
)

func Gtk_danmu() {
	if Gtk_on {return}

	application, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	if err != nil {return}

	application.Connect("startup", func() {
		log.Println("application startup")	

		builder, err := gtk.BuilderNewFromFile("ui/1.glade")
		if err != nil {return}

		signals := map[string]interface{}{
			"on_main_window_destroy": onMainWindowDestroy,
		}
		builder.ConnectSignals(signals)

		obj, err := builder.GetObject("main_window")
		if err != nil {return}



		win, err := isWindow(obj)
		if err != nil {return}

		var scrolledwindow0 *gtk.ScrolledWindow
		{
			obj, err := builder.GetObject("scrolledwindow0")
			if err != nil {return}
			if tmp,ok := obj.(*gtk.ScrolledWindow); ok {
				scrolledwindow0 = tmp
			}
		}

		var viewport0 *gtk.Viewport
		{
			obj, err := builder.GetObject("viewport0")
			if err != nil {return}
			if tmp,ok := obj.(*gtk.Viewport); ok {
				viewport0 = tmp
			}
		}

		var grid0 *gtk.Grid;
		{
			obj, err := builder.GetObject("grid0")
			if err != nil {return}
			if tmp,ok := obj.(*gtk.Grid); ok {
				grid0 = tmp
			}
		}

		var y func(bool)
		y = func(tra bool){
			Gtk_Tra = true
			
			s:=<-Gtk_danmuChan
			t,_ := gtk.TextViewNew();
			BTList.PushBack(t)
			t.SetMarginStart(5)
			t.SetEditable(false)
			t.SetHExpand(true)
			t.SetVExpand(false)
			t.SetWrapMode(gtk.WRAP_WORD_CHAR)
			
			img,_ :=gtk.ImageNew();
			BIList.PushBack(img)

			var handle glib.SignalHandle
			handle,_ = t.Connect("size-allocate", func(){
				Gtk_Tra = false
				b,e := t.GetBuffer()
				if e != nil {return}
				b.SetText(s)

				img_src := load_face(<-Gtk_danmuChan_uid)
				pixbuf,e := gdk.PixbufNewFromFileAtSize(img_src, 40, 40);
				if e != nil {return}
				img.SetFromPixbuf(pixbuf)
				t.HandlerDisconnect(handle)

				if tra {
					handle,_ = t.Connect("size-allocate", func(){
						tmp := scrolledwindow0.GetVAdjustment()
						tmp.SetValue(tmp.GetUpper())
						t.HandlerDisconnect(handle)
					})
				}
				if len(Gtk_danmuChan) != 0 {y(tra)}
			})

			tmp,_ := t.Container.Widget.Cast()
			loc := BTList.Len();
			grid0.InsertRow(loc);
			grid0.Attach(img, 0, loc - 1, 1, 1)
			grid0.Attach(tmp, 1, loc - 1, 1, 1)
			for tra && BTList.Len() > 50 {
				BTList.Remove(BTList.Front()).(*gtk.TextView).Destroy()
				BIList.Remove(BIList.Front()).(*gtk.Image).Destroy()
				grid0.RemoveRow(0)
			}
			win.ShowAll()
		}

		glib.TimeoutAdd(uint(100), func() bool {
			if !Gtk_Tra && len(Gtk_danmuChan) != 0 {
				tmp := scrolledwindow0.GetVAdjustment()
				h := viewport0.GetViewWindow().WindowGetHeight()
				y(tmp.GetUpper() - tmp.GetValue() < float64(h) * 1.3)
			}
			return true
		})
		win.Show()
		application.AddWindow(win)
		Gtk_on = true
	})

	application.Connect("activate", func() {
		log.Println("application activate")
	})

	application.Connect("shutdown", func() {
		log.Println("application shutdown")	
		Gtk_on = false
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

func load_face(uid string) string {
	defaultuid := "0default"
	if uid != "" && p.Checkfile().IsExist(Gtk_img_path + `/` + uid) {return Gtk_img_path + `/` + uid}
	if src := F.Get_face_src(uid);src != "" {
		req := p.Req()
		if e := req.Reqf(p.Rval{
			Url:src,
			SaveToPath:Gtk_img_path + `/` + uid,
			Timeout:3,
		}); e != nil{log.Println(e);return Gtk_img_path + `/` + defaultuid}
		return Gtk_img_path + `/` + uid
	}
	return Gtk_img_path + `/` + defaultuid
}