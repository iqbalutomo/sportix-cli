package repository

import (
	"database/sql"
	"sportix-cli/internal/entity"
)

type CategoryRepo interface {
	GetAllCategory() ([]entity.Category, error)
}

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepo {
	return &categoryRepo{db}
}

func (c *categoryRepo) GetAllCategory() ([]entity.Category, error) {
	var categories []entity.Category

	query := `SELECT category_id, name FROM categories;`
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category entity.Category
		err = rows.Scan(&category.CategoryID, &category.Name)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
