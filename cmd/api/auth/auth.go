package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwt(uuid string) (*string, error) {
	var mySigningKey = []byte(os.Getenv("JWT_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["uuid"] = uuid
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return nil, err
	}
	return &tokenString, nil

}
