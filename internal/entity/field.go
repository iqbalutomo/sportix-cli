package entity

import "time"

type Field struct {
	FieldID   int
	Name      string
	Price     float64
	Category  Category
	Location  Location
	Address   string
	Facility  Facility
	CreatedBy User
	CreatedAt time.Time
	UpdatedAt time.Time
}
