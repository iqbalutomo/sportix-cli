package entity

type User struct {
	UserID    uint
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt string
	UpdatedAt string
}

type CurrentUser struct {
	UserID    uint
	Username  string
	Email     string
	Role      string
	CreatedAt string
	UpdatedAt string
}
