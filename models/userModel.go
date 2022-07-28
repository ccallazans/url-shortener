package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	UUID      string    `json:"uuid" db:"uuid"`
	Email     string    `json:"email" validate:"required,email" db:"email"`
	Password  string    `json:"password" validate:"required,min=8,max=30" db:"password"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UserRepository interface {
	CreateUser(newUser User) error

	GetAllUsers() ([]*User, error)
	GetUserByEmail(email string) (*User, error)

	UpdateUserByEmail(oldEmail string, newEmail string) error

	DeleteUserByEmail(email string) error

	ValueExists(value string, column string) bool
}
