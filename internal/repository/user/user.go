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
	FindUserByEmail(email string) (*entity.User, error)
	ValidateUser(email, password string) (*entity.User, error)
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
