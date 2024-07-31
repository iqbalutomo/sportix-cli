package styles

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	BackgroundColor          = tcell.NewHexColor(0x282a36)
	TitleColor               = tcell.NewHexColor(0xf8f8f2)
	BorderColor              = tcell.NewHexColor(0x44475a)
	PrimaryColor             = tcell.NewHexColor(0xff79c6)
	SecondaryColor           = tcell.NewHexColor(0x8be9fd)
	CurrentLineColor         = tcell.NewHexColor(0x44475a)
	SelectionBackgroundColor = tcell.NewHexColor(0x6272a4)
	SelectionForegroundColor = tcell.NewHexColor(0xf8f8f2)
	CommentColor             = tcell.NewHexColor(0x6272a4)
	CyanColor                = tcell.NewHexColor(0x8be9fd)
	GreenColor               = tcell.NewHexColor(0x50fa7b)
	OrangeColor              = tcell.NewHexColor(0xffb86c)
	PinkColor                = tcell.NewHexColor(0xff79c6)
	PurpleColor              = tcell.NewHexColor(0xbd93f9)
	RedColor                 = tcell.NewHexColor(0xff5555)
	YellowColor              = tcell.NewHexColor(0xf1fa8c)
)

func ApplyTheme(p tview.Primitive) {
	switch component := p.(type) {
	case *tview.List:
		component.SetBackgroundColor(BackgroundColor)
		component.SetBorderColor(BorderColor)
		component.SetTitleColor(TitleColor)
	case *tview.Table:
		component.SetBackgroundColor(BackgroundColor)
		component.SetBorderColor(BorderColor)
		component.SetTitleColor(TitleColor)
	case *tview.TextView:
		component.SetBackgroundColor(BackgroundColor)
		component.SetBorderColor(BorderColor)
		component.SetTitleColor(TitleColor)
	}
}
