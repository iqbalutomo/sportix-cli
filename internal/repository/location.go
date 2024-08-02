package repository

import (
	"database/sql"
	"sportix-cli/internal/entity"
)

type LocationRepo interface {
	GetAllLocation() ([]entity.Location, error)
}

type locationRepo struct {
	db *sql.DB
}

func NewLocationRepo(db *sql.DB) LocationRepo {
	return &locationRepo{db}
}

func (l *locationRepo) GetAllLocation() ([]entity.Location, error) {
	var locations []entity.Location

	query := `SELECT location_id, name FROM locations;`
	rows, err := l.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var location entity.Location
		err = rows.Scan(&location.LocationID, &location.Name)
		if err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	return locations, nil
}
