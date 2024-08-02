package handler

import (
	"errors"
	"sportix-cli/internal/entity"
	"sportix-cli/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	Register(name, email, password, role string) error
	Login(email, password string) (*entity.User, error)
	GetBalanceByEmail(email string) (float64, error)
	PutBalance(userID uint, balance, deposit float64) error
}

type userHandler struct {
	repo repository.UserRepo
}

func NewUserHandler(repo repository.UserRepo) UserHandler {
	return &userHandler{repo}
}

func (u *userHandler) Register(name, email, password, role string) error {
	user, err := u.repo.FindUserByEmail(email)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &entity.User{
		Username: name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := u.repo.CreateUser(newUser); err != nil {
		return errors.New("failed to register")
	}

	return nil
}

func (auth *userHandler) Login(email, password string) (*entity.User, error) {
	user, err := auth.repo.ValidateUser(email, password)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (u *userHandler) GetBalanceByEmail(email string) (float64, error) {
	balance, err := u.repo.FindBalanceByEmail(email)
	if err != nil {
		return 0, nil
	}

	return balance, nil
}

func (u *userHandler) PutBalance(userID uint, balance, deposit float64) error {
	totalBalance := balance + deposit
	if err := u.repo.UpdateBalance(userID, totalBalance); err != nil {
		return err
	}

	return nil
}
