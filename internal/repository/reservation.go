package repository

import (
	"database/sql"
	"fmt"
	"sportix-cli/internal/entity"
	"sportix-cli/internal/session"
	"time"
)

type ReservationRepo interface {
	CheckFieldAvailability(fieldID, ahID int) (bool, error)
	CreateReservation(reservation entity.ReservationForm) error
}

type reservationRepo struct {
	db *sql.DB
}

func NewReservationRepo(db *sql.DB) ReservationRepo {
	return &reservationRepo{db}
}

func (r *reservationRepo) CreateReservation(reservation entity.ReservationForm) error {

	// Begin a transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Insert into reservations table
	reservationQuery := `INSERT INTO reservations (user_id, field_id, status, reservation_date) VALUES (?, ?, ?, ?);`
	reservationDate := time.Now().Format("2006-01-02")
	result, err := r.db.Exec(reservationQuery, reservation.UserID, reservation.FieldID, reservation.ReservationStatus, reservationDate)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting into reservations table: %v", err)
	}

	// Get the last inserted reservation ID
	reservationID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error getting last inserted reservation ID: %v", err)
	}

	// Insert into payments table
	paymentQuery := `INSERT INTO payments (reservation_id, amount, status) VALUES (?, ?, ?);`
	_, err = r.db.Exec(paymentQuery, reservationID, reservation.Amount, reservation.PaymentStatus)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting into payments table: %v", err)
	}

	// Update Wallets table
	updateBalance := session.UserSession.Balance - reservation.Amount
	session.UserSession.Balance = updateBalance
	walletQuery := `UPDATE wallets SET balance = ? WHERE user_id = ?;`
	_, err = r.db.Exec(walletQuery, updateBalance, reservation.UserID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating wallets table: %v", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func (r *reservationRepo) CheckFieldAvailability(fieldID, ahID int) (bool, error) {
	var availableHours entity.FieldAvailableHour

	queryAvailableHours := `SELECT fa.status
			FROM field_available_hours fa
			JOIN available_hours a ON fa.available_hour_id = a.available_hour_id
			WHERE fa.field_id = ? AND fa.available_hour_id = ?;`

	rows, err := r.db.Query(queryAvailableHours, fieldID, ahID)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		err = rows.Scan(&availableHours.Status)

		if availableHours.Status != "available" {
			return false, nil
		}
	}

	return true, nil
}
