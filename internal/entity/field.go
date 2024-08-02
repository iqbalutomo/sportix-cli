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

type FormAddsField struct {
	Name         string
	Price        string
	CategoryID   int
	LocationID   int
	Address      string
	Bathroom     string
	Cafeteria    string
	VehiclePark  string
	PrayerRoom   string
	ChangingRoom string
	CCTV         string
	CreatedByID  string
}
