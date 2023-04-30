package models

type UserEntity struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Role     Role   `gorm:"foreignKey:UserId"`
	Urls     []Url  `gorm:"foreignKey:UserID;references:ID"`
}

//

type User struct {
	ID       uint
	Username string
	Password string
	Role     Role
	Urls     []Url
}

//

type UserResponse struct {
	Username string `json:"username"`
	Urls     []Url  `json:"urls"`
}

//

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
