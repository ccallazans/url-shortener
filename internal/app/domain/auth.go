package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserAuth struct {
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
}

type JWTClaim struct {
	User UserAuth `json:"user"`
	jwt.RegisteredClaims
}
