package domain

const(
	USER = "USER"
)

type Role struct {
	ID   uint   `gorm:"primary_key"`
	Role string `gorm:"column:role"`
}
