package cli

import (
	"fmt"
	"sportix-cli/constants"
	"sportix-cli/internal/entity"
	"sportix-cli/internal/session"
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
			SetAlign(tview.AlignCenter)).SetSelectedFunc(func(row, column int) {
			errorModal := tview.NewModal().
				SetText("Are you sure to book this field time?").
				AddButtons([]string{"YES", "NO"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "YES" {
						if err := handler.Reservation.AddReservation(field, int(data.AvailableHourID.AvailableHourID)); err != nil {
							modal := tview.NewModal().
								SetText("Failed reservation").
								AddButtons([]string{""}).
								SetDoneFunc(func(buttonIndex int, buttonLabel string) {
									if session.UserSession.Role == "user" {
										UserDashboardPage(app, handler)
									} else {
										OwnerDashboardPage(app, handler)
									}
								})

							app.SetRoot(modal, true)
						}

						modal := tview.NewModal().
							SetText("Reservation Successfully!").
							AddButtons([]string{""}).
							SetDoneFunc(func(buttonIndex int, buttonLabel string) {
								if session.UserSession.Role == "user" {
									UserDashboardPage(app, handler)
								} else {
									OwnerDashboardPage(app, handler)
								}
							})

						app.SetRoot(modal, true)
					}
					if session.UserSession.Role == "user" {
						UserDashboardPage(app, handler)
					} else {
						OwnerDashboardPage(app, handler)
					}
				})
			app.SetRoot(errorModal, true).EnableMouse(true).Run()
		})
	}
	hoursTable.SetBorder(true).SetTitle("Available Hours").SetTitleAlign(tview.AlignCenter)

	detailAndFacilitiesFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	detailAndFacilitiesFlex.AddItem(detailText, 0, 2, true)
	detailAndFacilitiesFlex.AddItem(facilitiesText, 0, 1, false)

	flex.AddItem(detailAndFacilitiesFlex, 0, 2, true)
	flex.AddItem(hoursTable, 0, 1, false)

	return flex
}

func UpdateFieldForm(app *tview.Application, handler Handler, content *tview.Flex) tview.Primitive {
	form := tview.NewForm()

	idInput := tview.NewInputField()
	idInput.SetLabel("Field ID:").
		SetFieldWidth(20).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				id := idInput.GetText()
				fieldID, err := strconv.Atoi(id)
				if err != nil {
					return
				}

				field, err := handler.Field.GetFieldById(fieldID)
				if err != nil {
					showModal(app, handler, "Error", fmt.Sprintln("Your field not found."))
					return
				}
				if err == nil {
					field.FieldID = fieldID
				}
				content.Clear().SetTitle("").SetBorder(false)
				form.AddInputField("New Name", field.Name, 30, nil, func(text string) {
					field.Name = text
				}).
					AddInputField("New Address", field.Address, 100, nil, func(text string) {
						field.Address = text
					}).
					AddInputField("New Price", fmt.Sprintf("%.2f", field.Price), 10, nil, func(text string) {
						price, err := strconv.ParseFloat(text, 64)
						if err == nil {
							field.Price = price
						}
					}).
					AddInputField("New Bathroom Count", strconv.Itoa(field.Facility.Bathroom), 10, nil, func(text string) {
						bathroom, err := strconv.Atoi(text)
						if err == nil {
							field.Facility.Bathroom = bathroom
						}
					}).
					AddDropDown("Has Cafeteria (yes/no)", constants.YesNoOptions, 0, func(option string, index int) {
						field.Facility.Cafeteria = utils.IsYes(option)
					}).
					AddInputField("New Vehicle Park Area", strconv.Itoa(field.Facility.VehiclePark), 10, nil, func(text string) {
						vehiclePark, err := strconv.Atoi(text)
						if err == nil {
							field.Facility.VehiclePark = vehiclePark
						}
					}).
					AddDropDown("Has Prayer Room (yes/no)", constants.YesNoOptions, 0, func(option string, index int) {
						field.Facility.PrayerRoom = utils.IsYes(option)
					}).
					AddInputField("New Changing Room Count", strconv.Itoa(field.Facility.ChangingRoom), 10, nil, func(text string) {
						changingRoom, err := strconv.Atoi(text)
						if err == nil {
							field.Facility.ChangingRoom = changingRoom
						}
					}).
					AddDropDown("Has CCTV (yes/no)", constants.YesNoOptions, 0, func(option string, index int) {
						field.Facility.CCTV = utils.IsYes(option)
					}).
					AddButton("Update Field", func() {
						err := handler.Field.EditField(field)
						if err != nil {
							showModal(app, handler, "Error", fmt.Sprintf("Error updating field: %v", err))
							return
						}
						showModal(app, handler, "Success", "Field updated successfully!")
					}).
					AddButton("Cancel", func() {
						OwnerDashboardPage(app, handler)
					})

				content.AddItem(form, 0, 1, true)
			}
		})

	form.SetBorder(true).SetTitle("Update Field").SetTitleAlign(tview.AlignCenter)

	content.AddItem(idInput, 0, 1, true)
	return form
}

func showModal(app *tview.Application, handler Handler, title, message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if session.UserSession.Role == constants.User {
				UserDashboardPage(app, handler)
			} else {
				OwnerDashboardPage(app, handler)
			}
		})
	modal.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignCenter)
	app.SetRoot(modal, true).SetFocus(modal)
}
