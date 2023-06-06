package domain

import "github.com/google/uuid"

type User struct {
	UUID       uuid.UUID   `gorm:"type:uuid;primaryKey"`
	Username   string      `gorm:"column:username"`
	Password   string      `gorm:"column:password"`
	Role       string      `gorm:"column:role"`
	Shorteners []Shortener `gorm:"foreignKey:User"`
}

type UserResponse struct {
	UUID       uuid.UUID           `json:"uuid"`
	Username   string              `json:"username"`
	Role       string              `json:"role"`
	Shorteners []ShortenerResponse `json:"shorteners"`
}

//

func (u *User) ToResponse() UserResponse {

	var shortnersResponse []ShortenerResponse
	for _, resp := range u.Shorteners {
		shortnersResponse = append(shortnersResponse, resp.ToResponse())
	}

	return UserResponse{
		UUID:       u.UUID,
		Username:   u.Username,
		Role:       u.Role,
		Shorteners: shortnersResponse,
	}
}

func UsersToResponse(users []User) []UserResponse {

	var usersResponse []UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, user.ToResponse())
	}

	return usersResponse
}
