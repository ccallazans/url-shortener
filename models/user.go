package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	UUID      string    `json:"uuid" db:"uuid"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UserRepository interface {
	CreateUser(userModel User) error
	GetUserByEmail(email string) (*User, error)
	UUIDExists(uuid string) bool
	EmailExists(email string) bool
}
