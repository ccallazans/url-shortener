package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ccallazans/url-shortener/cmd/api/utils"
	"github.com/golang-jwt/jwt/v4"
)

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			utils.ErrorJSON(w, http.StatusBadRequest, fmt.Errorf("token not found"))
			return
		}

		var mySigningKey = []byte(os.Getenv("JWT_KEY"))

		token, err := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			utils.ErrorJSON(w, http.StatusBadRequest, fmt.Errorf("token has expired"))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "email", claims["email"].(string))
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		utils.ErrorJSON(w, http.StatusUnauthorized, fmt.Errorf("sai daqui"))
	})
}
