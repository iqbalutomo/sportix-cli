package cli

import (
	"fmt"
	"sportix-cli/constants"
	"sportix-cli/internal/entity"
	"sportix-cli/internal/utils"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ShowFields(app *tview.Application, handler Handler, content *tview.Flex) tview.Primitive {
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

	fields, err := handler.Field.GetFields()
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

func ShowFieldDetail(app *tview.Application, selectedRow int, field entity.Field, handler Handler) tview.Primitive {
	flex := tview.NewFlex().SetDirection(tview.FlexColumn)

	detailText := tview.NewTextView().
		SetText(fmt.Sprintf("Name: %s\n\nCategory: %s\n\nPrice: %s", field.Name, field.Category.Name, utils.FormatRupiah(field.Price))).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)

	detailText.SetBorder(true).
		SetTitle("Informations").
		SetTitleAlign(tview.AlignCenter)

	facilitiesText := tview.NewTextView().
		SetText(fmt.Sprintf("Bathroom: %d\n\nCafeteria: %s\n\nVehicle Park: %d\n\nPrayer Room: %s\n\nChanging Room: %d\n\nCCTV: %s", field.Facility.Bathroom, utils.BoolToYesNo(field.Facility.Cafeteria), field.Facility.VehiclePark, utils.BoolToYesNo(field.Facility.PrayerRoom), field.Facility.ChangingRoom, utils.BoolToYesNo(field.Facility.CCTV))).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)
	facilitiesText.SetBorder(true).
		SetTitle("Facilities").
		SetTitleAlign(tview.AlignCenter)

	hoursTable := tview.NewTable().
		SetBorders(true).
		SetBordersColor(tcell.ColorWhite).
		SetFixed(1, 1).SetSelectable(true, false)

	headers := []string{"Start Time", "End Time", "Status"}
	for col, header := range headers {
		hoursTable.SetCell(0, col, tview.NewTableCell(header).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter).
			SetSelectable(false))
	}

	fieldAvailableHours, err := handler.Field.GetFieldAvailableHours(uint(field.FieldID))
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

	for row, data := range fieldAvailableHours {
		hoursTable.SetCell(row+1, 0, tview.NewTableCell(data.AvailableHourID.StartTime).
			SetAlign(tview.AlignCenter))
		hoursTable.SetCell(row+1, 1, tview.NewTableCell(data.AvailableHourID.EndTime).
			SetAlign(tview.AlignCenter))
		hoursTable.SetCell(row+1, 2, tview.NewTableCell(data.Status).
			SetAlign(tview.AlignCenter))
	}
	hoursTable.SetBorder(true).SetTitle("Available Hours").SetTitleAlign(tview.AlignCenter)

	detailAndFacilitiesFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	detailAndFacilitiesFlex.AddItem(detailText, 0, 2, true)
	detailAndFacilitiesFlex.AddItem(facilitiesText, 0, 1, false)

	flex.AddItem(detailAndFacilitiesFlex, 0, 2, true)
	flex.AddItem(hoursTable, 0, 1, false)

	return flex
}

func AddField(app *tview.Application, handler Handler, content *tview.Flex) tview.Primitive {

	addFieldForm := &entity.FormAddsField{}

	categories, _ := handler.Category.GetAllCategory()
	categoriesOptions, _ := utils.ConvertStructSliceToStringSlice(categories, "Name")

	locations, _ := handler.Location.GetAllLocation()

	locationOptions, _ := utils.ConvertStructSliceToStringSlice(locations, "Name")

	// Create form
	form := tview.NewForm().
		AddInputField("Name:", "", 40, nil, func(text string) {
			addFieldForm.Name = text
		}).
		AddInputField("Price", "", 40, nil, func(text string) {
			addFieldForm.Price = text
		}).
		AddDropDown("Category", categoriesOptions, 0, func(option string, index int) {
			addFieldForm.CategoryID = index + 1
		}).
		AddDropDown("Location", locationOptions, 0, func(option string, index int) {
			addFieldForm.LocationID = index + 1
		}).
		AddInputField("Address", "", 40, nil, func(text string) {
			addFieldForm.Address = text
		}).
		AddInputField("Bathroom", "", 40, nil, func(text string) {
			addFieldForm.Bathroom = text
		}).
		AddDropDown("Cafeteria", constants.YesNoOptions, 0, func(option string, index int) {
			addFieldForm.Cafeteria = option
		}).
		AddInputField("Vehicle Park", "", 40, nil, func(text string) {
			addFieldForm.VehiclePark = text
		}).
		AddDropDown("Prayer Room", constants.YesNoOptions, 0, func(option string, index int) {
			addFieldForm.PrayerRoom = option
		}).
		AddInputField("Changing Room", "", 40, nil, func(text string) {
			addFieldForm.ChangingRoom = text
		}).
		AddDropDown("CCTV", constants.YesNoOptions, 0, func(option string, index int) {
			addFieldForm.CCTV = option
		}).
		AddButton("Add Field", func() {
			// Check if it AddInputField is empty
			if addFieldForm.Name == "" || addFieldForm.Price == "" || addFieldForm.Address == "" || addFieldForm.Bathroom == "" || addFieldForm.VehiclePark == "" || addFieldForm.ChangingRoom == "" {
				errorModal := tview.NewModal().
					SetText("All field cannot be empty").
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						OwnerDashboardPage(app, handler)
					})
				app.SetRoot(errorModal, true).EnableMouse(true).Run()
				return
			}

			err := handler.Field.AddField(addFieldForm)

			if err != nil {
				errorModal := tview.NewModal().
					SetText(err.Error()).
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						OwnerDashboardPage(app, handler)
					})
				app.SetRoot(errorModal, true).EnableMouse(true).Run()
				return
			}

			successModal := tview.NewModal().
				SetText("Add a New Field Successfully").
				AddButtons([]string{"OK"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					OwnerDashboardPage(app, handler)
				})

			app.SetRoot(successModal, true).EnableMouse(true).Run()
		})

	return form

}
