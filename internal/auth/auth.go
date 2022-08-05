package auth

import (
	"os"
	"time"

	"github.com/ccallazans/url-shortener/models"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	jwt.RegisteredClaims
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}

func GenerateJWT(user *models.User) (*string, error) {
	claims := &JWTClaim{
		UUID:  user.UUID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"localhost"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "localhost",
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   user.UUID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
