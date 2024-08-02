package handler

import (
	"errors"
	"sportix-cli/internal/entity"
	"sportix-cli/internal/repository"
)

type FacilityHandler interface {
	GetFacilityById(facility_id int) (*entity.Facility, error)
	EditFacilityById(updatedFacility *entity.Facility) error
}

type facilityHandler struct {
	repo repository.FacilityRepo
}

func NewFacilityHandler(repo repository.FacilityRepo) FacilityHandler {
	return &facilityHandler{repo: repo}
}

func (fa *facilityHandler) GetFacilityById(facility_id int) (*entity.Facility, error) {
	facility, err := fa.repo.FindFacilityById(int(facility_id))
	if err != nil {
		return nil, errors.New("failed to fetch facility By Id")
	}
	return facility, nil
}

func (fa *facilityHandler) EditFacilityById(updatedFacility *entity.Facility) error {
	return fa.repo.EditFacility(updatedFacility)
}
