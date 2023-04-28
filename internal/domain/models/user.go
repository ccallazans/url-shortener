package models

type User struct {
	ID       uint
	Username string
	Password string
	Urls     []Url
}

func (u *User) ToUserResponse() *UserResponse {
	return &UserResponse{
		Username: u.Username,
		Urls: u.Urls,
	}
}

//

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserRequest) ToUser() *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
	}
}

//

type UserResponse struct {
	Username string `json:"username"`
	Urls     []Url  `json:"urls"`
}