package cli

import (
	"sportix-cli/internal/session"
	"sportix-cli/internal/styles"
	"sportix-cli/internal/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func OwnerDashboardPage(app *tview.Application, handler Handler) {
	wallet := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Balance: "+utils.FormatRupiah(session.UserSession.Balance), "", 0, func() {})
	wallet.SetTitle("Your Wallet").SetBorder(true).SetTitleAlign(tview.AlignCenter)

	nav := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Profile", "", 0, func() {}).
		AddItem("Field List", "", 0, func() {}).
		AddItem("Add Field", "", 0, func() {})
	nav.SetTitle("Menu").SetBorder(true).SetTitleAlign(tview.AlignCenter)

	setting := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Logout", "", 0, func() {})
	setting.SetTitle("Setting").SetBorder(true).SetTitleAlign(tview.AlignCenter)

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("Select an item from the sidebar to see the content here").
			SetDynamicColors(true).
			SetTextAlign(tview.AlignLeft).
			SetBorder(true).
			SetTitle("Owner Dashboard: Welcome, "+session.UserSession.Username).
			SetTitleAlign(tview.AlignCenter), 0, 1, true)

	styles.ApplyTheme(wallet)
	styles.ApplyTheme(nav)
	styles.ApplyTheme(setting)
	styles.ApplyTheme(content)

	focusOrder := []tview.Primitive{wallet, nav, setting, content}
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

	wallet.SetInputCapture(setFocus)
	nav.SetInputCapture(setFocus)
	setting.SetInputCapture(setFocus)
	content.SetInputCapture(setFocus)

	wallet.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
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
			title := tview.NewForm().
				AddTextView("Name:", session.UserSession.Username, 40, 2, true, false).
				AddTextView("Email:", session.UserSession.Email, 40, 2, true, false)
			content.AddItem(title, 0, 1, true)

			// fieldsView := ShowFieldOwners(app, handler, content)
			// content.AddItem(fieldsView, 0, 1, true)

			content.SetBorder(true)
			content.SetTitle("Profile")

		case 1:
			content.Clear().SetTitle("").SetBorder(false)
			fieldsView := ShowFields(app, handler, content)
			content.AddItem(fieldsView, 0, 1, true)
			content.SetBorder(true)
			content.SetTitle("Field List")
		case 2:
			content.Clear().SetTitle("").SetBorder(false)
			addFieldView := AddField(app, handler, content)
			content.AddItem(addFieldView, 0, 1, true)
			content.SetBorder(true)
			content.SetTitle("Add Field")
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
		AddItem(wallet, 0, 0, 1, 1, 0, 0, false).
		AddItem(nav, 1, 0, 1, 1, 0, 0, true).
		AddItem(setting, 2, 0, 1, 1, 0, 0, true).
		AddItem(content, 0, 1, 3, 1, 0, 0, false)

	if err := app.SetRoot(mainLayout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
