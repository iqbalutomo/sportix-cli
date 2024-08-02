package cli

import (
	"sportix-cli/internal/session"
	"sportix-cli/internal/utils"
	"strconv"

	"github.com/rivo/tview"
)

func ShowDeposit(app *tview.Application, handler Handler) tview.Primitive {
	var totalDeposit float64
	var inputValid bool

	form := tview.NewForm().
		AddTextView("Your Balance:", utils.FormatRupiah(session.UserSession.Balance), 40, 2, true, false).
		AddInputField("Total Deposit", "", 40, nil, func(balanceStr string) {
			balance, err := strconv.ParseFloat(balanceStr, 64)
			if err != nil {
				showAlertModal(app, handler, "Invalid input for total deposit.")
				inputValid = false
				return
			}
			totalDeposit = balance
			inputValid = true
		}).
		AddButton("Pay", func() {
			if !inputValid {
				showAlertModal(app, handler, "Please enter a valid total deposit.")
				return
			}
			if totalDeposit <= 0 {
				showAlertModal(app, handler, "Total deposit must be greater than zero.")
				return
			}

			//if err := handler.User.PutBalance(session.UserSession.UserID, session.UserSession.Balance, totalDeposit); err != nil {
			//	showAlertModal(app, handler, "Failed to deposit, try again!")
			//	return
			//}

			session.UserSession.Balance += totalDeposit

			showAlertModal(app, handler, "Deposit Successfully!")
		})

	form.SetBorder(true).SetTitle("Payment Deposit").SetTitleAlign(tview.AlignLeft)

	return form
}

func showAlertModal(app *tview.Application, handler Handler, textErr string) {
	modal := tview.NewModal().
		SetText(textErr).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {

			OwnerDashboardPage(app, handler)
		})

	app.SetRoot(modal, true)
}
