package handler

import (
	"fmt"
	"sportix-cli/internal/entity"
	"sportix-cli/internal/repository"
	"sportix-cli/internal/session"
)

type ReservationHandler interface {
	AddReservation(field entity.Field, ahID int) error
}

type reservationHandler struct {
	repo repository.ReservationRepo
}

func NewReservationHandler(repo repository.ReservationRepo) ReservationHandler {
	return &reservationHandler{repo}
}

func (r *reservationHandler) AddReservation(field entity.Field, ahID int) error {

	// Check Availability of Field
	isAvailable, err := r.repo.CheckFieldAvailability(field.FieldID, ahID)
	if err != nil {
		return err
	}

	if !isAvailable {
		return fmt.Errorf("field is not available at that time")
	}

	// Check User's Balance if price is greater than balance
	if field.Price > session.UserSession.Balance {
		return fmt.Errorf("insufficient balance")
	}

	// Get Total Price
	totalPrice := field.Price

	// Create Reservation
	reservation := entity.ReservationForm{
		UserID:            session.UserSession.UserID,
		FieldID:           field.FieldID,
		ReservationStatus: "confirmed",
		Amount:            totalPrice,
		PaymentStatus:     "paid",
	}

	err = r.repo.CreateReservation(reservation)
	if err != nil {
		return err
	}

	return nil

}
