package models

type Role struct {
	ID     uint   `gorm:"primary_key"`
	Role   string `gorm:"column:role"`
	UserId uint   `gorm:"column:user_id"`
}
