package domain

import "github.com/google/uuid"

type User struct {
	UUID       uuid.UUID   `gorm:"type:uuid;primaryKey"`
	Username   string      `gorm:"column:username"`
	Password   string      `gorm:"column:password"`
	Role       string      `gorm:"column:role"`
	Shorteners []Shortener `gorm:"foreignKey:User"`
}
