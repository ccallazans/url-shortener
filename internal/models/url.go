package models

type Url struct {
	ID     uint   `gorm:"primaryKey"`
	Url    string `gorm:"not null`
	Hash   string `gorm:"not null`
	UserID uint   `gorm:"not null`
}
