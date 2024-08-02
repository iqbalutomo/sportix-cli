package handler

import (
	"fmt"
	"sportix-cli/internal/entity"
	"sportix-cli/internal/repository"
	"sportix-cli/internal/session"
	"sportix-cli/internal/utils"
)

type FieldHandler interface {
	GetFields() ([]entity.Field, error)
	GetFieldOwners(userID uint) ([]entity.Field, error)
	GetFieldAvailableHours(fieldID uint) ([]entity.FieldAvailableHour, error)
	AddField(field *entity.FormAddsField) error
}

type fieldHandler struct {
	repo repository.FieldRepo
}

func NewFieldHandler(repo repository.FieldRepo) FieldHandler {
	return &fieldHandler{repo}
}

func (f *fieldHandler) GetFields() ([]entity.Field, error) {
	fields, err := f.repo.FindAllFields()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return fields, nil
}

func (f *fieldHandler) GetFieldOwners(userID uint) ([]entity.Field, error) {
	fields, err := f.repo.FindAllFieldsByOwner(userID)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return fields, nil
}

func (f *fieldHandler) GetFieldAvailableHours(fieldID uint) ([]entity.FieldAvailableHour, error) {
	fields, err := f.repo.FindAllHoursByFieldID(fieldID)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return fields, nil
}

func (f *fieldHandler) AddField(fieldForm *entity.FormAddsField) error {

	price, err := utils.ParseDataType[float64](fieldForm.Price)
	if err != nil {
		return fmt.Errorf("invalid price input, please input a number")
	}
	if err := utils.CheckNonNegative(price, "Price"); err != nil {
		return err
	}

	bathrooms, err := utils.ParseDataType[int](fieldForm.Bathroom)
	if err != nil {
		return fmt.Errorf("invalid Bathroom input, please input a number")
	}
	if err := utils.CheckNonNegative(bathrooms, "Bathroom"); err != nil {
		return fmt.Errorf("invalid Bathroom input, please input a positive number")
	}

	cafeteria := utils.IsYes(fieldForm.Cafeteria)

	vehiclePark, err := utils.ParseDataType[int](fieldForm.VehiclePark)
	if err != nil {
		return fmt.Errorf("invalid Vehicle Park input, please input a number")
	}
	if err := utils.CheckNonNegative(vehiclePark, "VehiclePark"); err != nil {
		return fmt.Errorf("invalid Vehicle Park input, please input a positive number")
	}

	prayerRoom := utils.IsYes(fieldForm.PrayerRoom)

	changingRoom, err := utils.ParseDataType[int](fieldForm.ChangingRoom)
	if err != nil {
		return fmt.Errorf("failed to convert Changing Room value: %v", err)
	}
	if err := utils.CheckNonNegative(changingRoom, "ChangingRoom"); err != nil {
		return fmt.Errorf("invalid Changing Room input, please input a positive number")
	}

	cctv := utils.IsYes(fieldForm.CCTV)

	field := &entity.Field{
		Name:     fieldForm.Name,
		Price:    price,
		Category: entity.Category{CategoryID: fieldForm.CategoryID},
		Location: entity.Location{LocationID: fieldForm.LocationID},
		Address:  fieldForm.Address,
		Facility: entity.Facility{
			Bathroom:     bathrooms,
			Cafeteria:    cafeteria,
			VehiclePark:  vehiclePark,
			PrayerRoom:   prayerRoom,
			ChangingRoom: changingRoom,
			CCTV:         cctv,
		},
		CreatedBy: entity.User{
			UserID: session.UserSession.UserID,
		},
	}

	err = f.repo.AddNewField(field)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
