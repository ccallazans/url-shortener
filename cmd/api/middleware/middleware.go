package middleware

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/time/rate"

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
				return nil, fmt.Errorf("error parsing token")
			}
			return mySigningKey, nil
		})

		if err != nil {
			utils.ErrorJSON(w, http.StatusBadRequest, fmt.Errorf("token has expired"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			type contextKey string
			userContextKey := contextKey("user")

			ctx := context.WithValue(r.Context(), userContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		utils.ErrorJSON(w, http.StatusUnauthorized, fmt.Errorf("token not valid"))
	})
}

func RateLimitIp(next http.Handler) http.Handler {

	type client struct {
		limiter *rate.Limiter
		lastSeen time.Time
	}
		

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			utils.ErrorJSON(w, http.StatusInternalServerError, err)
			return
		}

		mu.Lock()

		if _, found := clients[ip]; !found {
			rt := rate.Every(5* time.Second) // Permite até 5 requisições por segundo
			clients[ip] = &client{limiter: rate.NewLimiter(rt, 1)}
		}

		clients[ip].lastSeen = time.Now()


		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			utils.ErrorJSON(w, http.StatusTooManyRequests, errors.New("too many requests"))
			return
		}

		mu.Unlock()
		next.ServeHTTP(w, r)

	})
}
