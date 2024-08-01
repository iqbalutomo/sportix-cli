package cli

import (
	"sportix-cli/internal/session"
	"sportix-cli/internal/styles"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func UserDashboardPage(app *tview.Application, handler Handler) {
	filter := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Search", "", 0, func() {}).
		AddItem("Filter", "", 0, func() {})
	filter.SetTitle("Filter / Search").SetBorder(true).SetTitleAlign(tview.AlignCenter)

	nav := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Profile", "", 0, func() {}).
		AddItem("Show Fields", "", 0, func() {}).
		AddItem("Reservation Field", "", 0, func() {})
	nav.SetTitle("Welcome, " + session.UserSession.Username).SetBorder(true).SetTitleAlign(tview.AlignCenter)

	setting := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Logout", "", 0, func() {})
	setting.SetTitle("Setting").SetBorder(true).SetTitleAlign(tview.AlignCenter)

	content := tview.NewTextView().
		SetText("Select an item from the sidebar to see the content here")
	content.SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetBorder(true).
		SetTitle("Content").
		SetTitleAlign(tview.AlignCenter)

	styles.ApplyTheme(filter)
	styles.ApplyTheme(nav)
	styles.ApplyTheme(setting)
	styles.ApplyTheme(content)

	focusOrder := []tview.Primitive{filter, nav, setting, content}
	setFocus := func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			for i, p := range focusOrder {
				if app.GetFocus() == p {
					app.SetFocus(focusOrder[(i+1)%len(focusOrder)])
					return nil
				}
			}
		case tcell.KeyBacktab:
			for i, p := range focusOrder {
				if app.GetFocus() == p {
					app.SetFocus(focusOrder[(i+len(focusOrder)-1)%len(focusOrder)])
					return nil
				}
			}
		}
		return event
	}

	filter.SetInputCapture(setFocus)
	nav.SetInputCapture(setFocus)
	setting.SetInputCapture(setFocus)
	content.SetInputCapture(setFocus)

	filter.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			content.SetText("search content")
			content.SetTitle("Search")
		case 1:
			content.SetText("filter content")
			content.SetTitle("Filter")
		}
	})

	nav.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			content.SetText("profile content")
			content.SetTitle("Profile")
		case 1:
			content.SetText("field list content")
			content.SetTitle("Field List")
		case 2:
			content.SetText("reservation content")
			content.SetTitle("Reservation Field")
		}
	})

	setting.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			showLogoutModal(app, handler)
		}
	})

	mainLayout := tview.NewGrid().
		SetBorders(false).
		SetColumns(20, -1).
		SetRows(3, 0, -1).
		AddItem(filter, 0, 0, 1, 1, 0, 0, false).
		AddItem(nav, 1, 0, 1, 1, 0, 0, true).
		AddItem(setting, 2, 0, 1, 1, 0, 0, true).
		AddItem(content, 0, 1, 3, 1, 0, 0, false)

	if err := app.SetRoot(mainLayout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func showLogoutModal(app *tview.Application, handler Handler) {
	modal := tview.NewModal().
		SetText("Are you sure you want to logout?").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				session.UserSession = nil
				AuthModal(app, handler)
			}
			UserDashboardPage(app, handler)
		})

	app.SetRoot(modal, true)
}
