package models

type UserRepository interface {
	CreateUser(newUser User) error

	GetAllUsers() ([]*User, error)
	GetUserByEmail(email string) (*User, error)

	UpdateUserByEmail(oldEmail string, newEmail string) error

	DeleteUserByEmail(email string) error

	ValueExists(value string, column string) bool
}
