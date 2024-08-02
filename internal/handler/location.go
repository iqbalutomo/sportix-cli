package handler

import (
	"sportix-cli/internal/entity"
	"sportix-cli/internal/repository"
)

type LocationHandler interface {
	GetAllLocation() ([]entity.Location, error)
}

type locationHandler struct {
	repo repository.LocationRepo
}

func NewLocationHandler(repo repository.LocationRepo) LocationHandler {
	return &locationHandler{repo}
}

func (l *locationHandler) GetAllLocation() ([]entity.Location, error) {
	locations, err := l.repo.GetAllLocation()
	if err != nil {
		return nil, err
	}

	return locations, nil
}
