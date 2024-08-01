package handler

import (
	"errors"
	"sportix-cli/internal/entity"
	repository "sportix-cli/internal/repository/user"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	Register(name, email, password, role string) error
	Login(email, password string) (*entity.User, error)
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
		return user, err
	}

	return user, nil
}
