package models

import "time"

type Url struct {
	ID        int       `json:"-" db:"id"`
	Short     string    `json:"short" db:"short"`
	Url       string    `json:"url" db:"url"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UrlResponse struct {
	Short     string    `json:"short" db:"short"`
	Url       string    `json:"url" db:"url"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type User struct {
	ID        int       `json:"id" db:"id"`
	UUID      string    `json:"uuid" db:"uuid"`
	Email     string    `json:"email" validate:"required,email" db:"email"`
	Password  string    `json:"password" validate:"required,min=8,max=30" db:"password"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
