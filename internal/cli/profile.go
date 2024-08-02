package cli

import (
	"sportix-cli/internal/session"
	"sportix-cli/internal/utils"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ShowFieldOwners(app *tview.Application, handler Handler, content *tview.Flex) tview.Primitive {
	table := tview.NewTable().
		SetBorders(true).
		SetFixed(1, 1).SetSelectable(true, false)

	headers := []string{"ID", "Name", "Price", "Category", "Location", "Address"}
	for col, header := range headers {
		table.SetCell(0, col, tview.NewTableCell(header).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter).
			SetSelectable(false))
	}

	fields, err := handler.Field.GetFieldOwners(session.UserSession.UserID)
	if err != nil {
		modal := tview.NewModal().
			SetText(err.Error()).
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "OK" {
					UserDashboardPage(app, handler)
				}
			})

		app.SetRoot(modal, true)
	}

	for row, item := range fields {
		fieldIDStr := strconv.FormatUint(uint64(item.FieldID), 10)

		table.SetCell(row+1, 0, tview.NewTableCell(fieldIDStr).
			SetAlign(tview.AlignCenter))
		table.SetCell(row+1, 1, tview.NewTableCell(item.Name).
			SetAlign(tview.AlignCenter))
		table.SetCell(row+1, 2, tview.NewTableCell(utils.FormatRupiah(item.Price)).
			SetAlign(tview.AlignCenter))
		table.SetCell(row+1, 3, tview.NewTableCell(item.Category.Name).
			SetAlign(tview.AlignCenter))
		table.SetCell(row+1, 4, tview.NewTableCell(item.Location.Name).
			SetAlign(tview.AlignCenter))
		table.SetCell(row+1, 5, tview.NewTableCell(item.Address).
			SetAlign(tview.AlignCenter))
	}

	table.SetSelectedFunc(func(row, column int) {
		content.Clear().SetTitle("").SetBorder(false)
		fieldDetail := ShowFieldDetail(app, row-1, fields[row-1], handler)
		content.AddItem(fieldDetail, 0, 1, true)
		content.SetBorder(true)
		content.SetTitle("Field Detail")
	})

	return table
}
