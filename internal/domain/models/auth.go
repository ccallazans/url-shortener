package models

import "github.com/golang-jwt/jwt/v4"

type UserClaims struct {
	User User `json:"user"`
	jwt.StandardClaims
}
