package repository

import (
	"database/sql"
	"fmt"
	"sportix-cli/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(user *entity.User) error
	//CreateWallet(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
	ValidateUser(email, password string) (*entity.User, error)
	FindBalanceByEmail(email string) (float64, error)
	UpdateBalance(userID uint, deposit float64) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db}
}

func (u *userRepo) CreateUser(user *entity.User) error {
	// Begin a transaction
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?);`
	result, err := u.db.Exec(query, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error creating user table: %v", err)
	}

	// Get the last inserted user ID
	userID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error getting last inserted user ID: %v", err)
	}

	user.UserID = uint(userID)

	// Create a wallet for the user
	query = `INSERT INTO wallets (user_id) VALUES (?);`
	_, err = u.db.Exec(query, user.UserID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error creating wallet table: %v", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func (u *userRepo) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	query := `SELECT user_id, username, password, email, role FROM users WHERE email = ?;`
	rows := u.db.QueryRow(query, email)

	if err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (u *userRepo) ValidateUser(email, password string) (*entity.User, error) {
	user, err := u.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	return user, nil
}

func (u *userRepo) FindBalanceByEmail(email string) (float64, error) {
	var wallet float64

	query := `SELECT w.balance
			FROM users u
			JOIN wallets w ON u.user_id = w.user_id
			WHERE email = ?;`
	rows := u.db.QueryRow(query, email)

	if err := rows.Scan(&wallet); err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}

		return 0, err
	}

	return wallet, nil
}

func (u *userRepo) UpdateBalance(userID uint, totalBalance float64) error {
	query := `UPDATE wallets SET balance = ? WHERE user_id = ?;`

	_, err := u.db.Exec(query, totalBalance, userID)
	if err != nil {
		return err
	}

	return nil
}
