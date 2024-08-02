package cli

import (
	"sportix-cli/internal/session"

	"github.com/rivo/tview"
)

func OwnerDashboardPage(app *tview.Application, handler Handler) {
	nav := tview.NewList().
		ShowSecondaryText(false).
		AddItem("1. Show Fields", "", 0, func() {}).
		AddItem("2. Add Field", "", 0, func() {}).
		AddItem("3. Edit Field", "", 0, func() {}).
		AddItem("4. Revenue Report", "", 0, func() {})
	nav.SetTitle("Welcome, " + session.UserSession.Username).SetBorder(true).SetTitleAlign(tview.AlignCenter)

	mainLayout := tview.NewGrid().
		SetBorders(false).
		SetColumns(20, -1).
		SetRows(3, 0, -1).
		// AddItem(filter, 0, 0, 1, 1, 0, 0, false).
		AddItem(nav, 1, 0, 1, 1, 0, 0, true)
		// AddItem(setting, 2, 0, 1, 1, 0, 0, true).
		// AddItem(content, 0, 1, 3, 1, 0, 0, false)

	if err := app.SetRoot(mainLayout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
