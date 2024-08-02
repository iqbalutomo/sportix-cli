package cli

import (
	"sportix-cli/internal/session"
	"sportix-cli/internal/styles"
	"sportix-cli/internal/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func UserDashboardPage(app *tview.Application, handler Handler) {
	filter := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Balance: "+utils.FormatRupiah(session.UserSession.Balance), "", 0, func() {})
	filter.SetTitle("Your Wallet").SetBorder(true).SetTitleAlign(tview.AlignCenter)

	nav := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Profile", "", 0, func() {}).
		AddItem("Field List", "", 0, func() {}).
		AddItem("Reservation Field", "", 0, func() {})
	nav.SetTitle("Menu").SetBorder(true).SetTitleAlign(tview.AlignCenter)

	setting := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Logout", "", 0, func() {})
	setting.SetTitle("Setting").SetBorder(true).SetTitleAlign(tview.AlignCenter)

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().
			SetDynamicColors(true).
			SetTextAlign(tview.AlignLeft).
			SetBorder(true).
			SetTitle("User Dashboard: Welcome, "+session.UserSession.Username).
			SetTitleAlign(tview.AlignCenter), 0, 1, true)

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
			content.Clear().SetTitle("").SetBorder(false)
			depositView := ShowDeposit(app, handler)
			content.AddItem(depositView, 0, 1, true)
			content.SetBorder(true)
			content.SetTitle("Deposit Wallet")
		}
	})

	nav.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			content.Clear().SetTitle("").SetBorder(false)
			profileView := tview.NewForm().AddInputField("Username", "", 20, nil, nil).
				AddPasswordField("Password", "", 20, '*', nil)
			profileView.
				SetBorder(true).
				SetTitle("Profile").
				SetTitleAlign(tview.AlignCenter)

			content.AddItem(profileView, 0, 1, true)
		case 1:
			content.Clear().SetTitle("").SetBorder(false)
			fieldsView := ShowFields(app, handler, content)
			content.AddItem(fieldsView, 0, 1, true)
			content.SetBorder(true)
			content.SetTitle("Field List")
		case 2:
			content.Clear().SetTitle("").SetBorder(false)
			reservationView := tview.NewTextView().
				SetText("You selected Reservation Field.\nHere you can reserve a field.").
				SetDynamicColors(true).
				SetTextAlign(tview.AlignLeft).
				SetBorder(true).
				SetTitle("Reservation Field").
				SetTitleAlign(tview.AlignCenter)
			content.AddItem(reservationView, 0, 1, true)
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
