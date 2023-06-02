package domain

import (
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `gorm:"primary_key"`
	Username string    `gorm:"column:username"`
	Password string    `gorm:"column:password"`
	Role     Role      `gorm:"embedded"`
}
