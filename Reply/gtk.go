package reply

import (
	"container/list"
	"errors"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appId = "com.github.qydysky.bili_danmu.reply"
var BList = list.New()

var (
	Gtk_on bool
	Gtk_danmuChan chan string = make(chan string, 10)
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

		var grid0 *gtk.Grid;
		{
			obj, err := builder.GetObject("grid0")
			if err != nil {return}
			if tmp,ok := obj.(*gtk.Grid); ok {
				grid0 = tmp
			}
		}

		var y func()
		y = func(){
			s:=<-Gtk_danmuChan
			t,_ := gtk.TextViewNew();
			t.SetEditable(false)
			t.SetHExpand(true)
			t.SetVExpand(false)
			t.SetWrapMode(gtk.WRAP_WORD_CHAR)
			var handle glib.SignalHandle
			handle,_ = t.Connect("size-allocate", func(){
				b,e := t.GetBuffer()
				if e != nil {return}
				b.SetText(s)
				t.HandlerDisconnect(handle)

				tmp := scrolledwindow0.GetVAdjustment()
				tmp.SetValue(tmp.GetUpper())
				if len(Gtk_danmuChan) != 0 {y()}
			})
			
			tmp,_ := t.Container.Widget.Cast()
			loc := int(grid0.Container.GetChildren().Length());
			grid0.InsertRow(loc);
			grid0.Attach(tmp, 0, loc, 1, 1)
			if loc > 50 {
				l,_ := grid0.GetChildAt(0, 0)
				l.ToWidget().Destroy()
				grid0.RemoveRow(0)
			}
			win.ShowAll()
		}

		glib.TimeoutAdd(uint(100), func() bool {
			if len(Gtk_danmuChan) != 0 {y()}
			
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

	os.Exit(application.Run(nil))
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
