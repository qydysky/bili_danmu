package reply

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var (
	centralLayout       *widgets.QGridLayout
	QtDanmuChan chan string
	QtOn bool
	Qt_MaxMun int = 30//danmu max limit
	Qt_LineHeight float64 = 90//percent
	Qt_BlockMargin float64 = 7
)

func Qtdanmu() {
	widgets.NewQApplication(len(os.Args), os.Args)

	//主窗口
	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetWindowTitle("danmu")
	mainWindow.SetSizePolicy2(widgets.QSizePolicy__Maximum, widgets.QSizePolicy__Maximum)
	mainWindow.SetContentsMargins(0, 0, 0, 0)
	mainWindow.SetWindowOpacity(1)
	mainWindow.SetAttribute(core.Qt__WA_TranslucentBackground, true)
	{
		Qp := gui.NewQPalette()
		Qp.SetColor2(gui.QPalette__Background, gui.NewQColor3(0, 0, 0, 0));
		mainWindow.SetPalette(Qp)
	}

	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetContentsMargins(0, 0, 0, 0)
	{
		Qp := gui.NewQPalette()
		Qp.SetColor2(gui.QPalette__Background, gui.NewQColor3(0, 0, 0, 0));
		centralWidget.SetPalette(Qp)
	}

	
	vbox := widgets.NewQGridLayout(centralWidget)
	t := new(centralWidget, vbox)

	centralWidget.SetLayout(vbox)
	mainWindow.SetCentralWidget(centralWidget)
	mainWindow.ShowNormalDefault()

	go func(){
		QtDanmuChan = make(chan string, 10)
		QtOn = true
		// var list []string
		for QtOn {
			select{
			case i :=<-QtDanmuChan:
				text(i, t)
			}
		}
	}()
	widgets.QApplication_Exec()
	QtOn = false
}

func new(pare *widgets.QWidget, layouts *widgets.QGridLayout) (t *widgets.QTextEdit) {
	t = widgets.NewQTextEdit(pare)
	{
		Qp := gui.NewQPalette()
		Qp.SetColor2(gui.QPalette__Base, gui.NewQColor3(0, 0, 0, 140));
		t.SetPalette(Qp)
	}
	t.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	t.SetWordWrapMode(gui.QTextOption__WrapAnywhere)
	// t.SetBackgroundVisible(true)
	// t.SetMaximumBlockCount(100)
	t.SetContentsMargins(0, 0, 0, 0)
	// t.SetCenterOnScroll(false)
	t.SetReadOnly(true)
	{
		t.SetTextBackgroundColor(gui.NewQColor3(0, 0, 0, 0))

		f := gui.NewQFont()
		f.SetPixelSize(18)
		f.SetWeight(63)
		t.SetCurrentFont(f)
	}
	{
		tc := t.TextCursor()
		b := tc.BlockFormat()
		b.SetLineHeight(Qt_LineHeight, int(gui.QTextBlockFormat__ProportionalHeight))
		b.SetBottomMargin(Qt_BlockMargin)
		tc.SetBlockFormat(b)
		t.SetTextCursor(tc)
	}
	layouts.AddWidget2(t, layouts.RowCount(), 0, 0)
	return
}

func text(s string, pare *widgets.QTextEdit) {
	c := pare.TextCursor()
	c.BeginEditBlock()
	c.InsertText(s)
	c.EndEditBlock()
	c.InsertBlock()
	if pare.Document().BlockCount() > Qt_MaxMun {
		c.MovePosition(gui.QTextCursor__Start, gui.QTextCursor__MoveAnchor, 1)
		// c.BeginEditBlock()
		c.MovePosition(gui.QTextCursor__NextBlock, gui.QTextCursor__KeepAnchor, 1)
		// c.Select(gui.QTextCursor__BlockUnderCursor)
		c.RemoveSelectedText()
		// c.EndEditBlock()
		c.MovePosition(gui.QTextCursor__End, gui.QTextCursor__MoveAnchor, 1)
	}
	// t := pare.ToPlainText()

	pare.EnsureCursorVisible()
	// pare.SetPlainText(s + "\n" + t)
}
