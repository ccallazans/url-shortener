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
	UUID       uuid.UUID   `json:"uuid"`
	Username   string      `json:"username"`
	Role       string      `json:"role"`
	Shorteners []ShortenerResponse `json:"shorteners"`
}

func (u *User) toResponse() UserResponse {

	var shortnerResponses []ShortenerResponse
	for _, resp := range u.Shorteners {
		shortnerResponses = append(shortnerResponses, resp.toResponse())
	}

	return UserResponse{
		UUID: u.UUID,
		Username: u.Username,
		Role: u.Role,
		Shorteners: shortnerResponses,
	}
}