package entity

import "time"

type Reservation struct {
	ReservationID   int
	User            User
	Field           Field
	Status          string
	ReservationDate time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ReservationForm struct {
	UserID            uint
	FieldID           int
	ReservationStatus string
	ReservationDate   time.Time
	Amount            float64
	PaymentStatus     string
}
