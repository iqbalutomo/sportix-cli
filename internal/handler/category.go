package handler

import (
	"sportix-cli/internal/entity"
	"sportix-cli/internal/repository"
)

type CategoryHandler interface {
	GetAllCategory() ([]entity.Category, error)
}

type categoryHandler struct {
	repo repository.CategoryRepo
}

func NewCategoryHandler(repo repository.CategoryRepo) CategoryHandler {
	return &categoryHandler{repo}
}

func (c *categoryHandler) GetAllCategory() ([]entity.Category, error) {
	categories, err := c.repo.GetAllCategory()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
