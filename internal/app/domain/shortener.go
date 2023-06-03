package domain

import "github.com/google/uuid"

type Shortener struct {
	ID   uint      `gorm:"primaryKey"`
	Url  string    `gorm:"column:url"`
	Hash string    `gorm:"column:hash"`
	User uuid.UUID `gorm:"foreignKey:UUID"`
}
