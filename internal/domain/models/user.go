package models

type UserEntity struct {
	ID       uint
	Username string
	Password string
}

//

type User struct {
	ID       uint
	Username string
	Password string
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