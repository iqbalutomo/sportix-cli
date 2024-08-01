package repository

import (
	"database/sql"
	"fmt"
	"log"
	"sportix-cli/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(user *entity.User) error
	CreateWallet(user *entity.User) error
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
	query := `INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?);`
	_, err := u.db.Exec(query, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}

func (u *userRepo) CreateWallet(user *entity.User) error {
	query := `INSERT INTO wallets (user_id) VALUES (?);`
	_, err := u.db.Exec(query, user.UserID)
	if err != nil {
		log.Fatal(err.Error())
		return err
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

func (auth *userRepo) ValidateUser(email, password string) (*entity.User, error) {
	user, err := auth.FindUserByEmail(email)
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
