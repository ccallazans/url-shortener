package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"url-shortener/cmd/api/handlers"

	"github.com/golang-jwt/jwt"
)

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			handlers.ErrorJSON(w, http.StatusBadRequest, errors.New("no authorization"))
			return
		}

		var mySigningKey = []byte(os.Getenv("JWT_KEY"))

		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {

			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("error parsing jwt")
			}

			return mySigningKey, nil
		})

		if err != nil {
			handlers.ErrorJSON(w, http.StatusUnauthorized, errors.New("token expired"))
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			w.WriteHeader(http.StatusOK)
			next.ServeHTTP(w, r)
			return
		}

		handlers.ErrorJSON(w, http.StatusUnauthorized, errors.New("error authentication"))
	})
}
