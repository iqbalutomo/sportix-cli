package entity

import "time"

type User struct {
	UserID    uint
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CurrentUser struct {
	UserID    uint
	Username  string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
