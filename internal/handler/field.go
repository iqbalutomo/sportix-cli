package handler

import (
	"fmt"
	"sportix-cli/internal/entity"
	"sportix-cli/internal/repository"
)

type FieldHandler interface {
	GetFields() ([]entity.Field, error)
	GetFieldAvailableHours(fieldID uint) ([]entity.FieldAvailableHour, error)
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

func (f *fieldHandler) GetFieldAvailableHours(fieldID uint) ([]entity.FieldAvailableHour, error) {
	fields, err := f.repo.FindAllHoursByFieldID(fieldID)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return fields, nil
}

func (f *fieldHandler) EditField(updatedField *entity.Field) error {
	return f.repo.EditField(updatedField)
}
