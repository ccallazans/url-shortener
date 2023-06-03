package middleware

import (
	"errors"
	"myapi/internal/app/shared"

	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := shared.HandleError(errors.New(shared.AUTHORIZATION_HEADER_EMPTY))
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			response := shared.HandleError(errors.New(shared.AUTHORIZATION_HEADER_FORMAT_ERROR))
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}
		tokenString := authHeaderParts[1]

		// validate the token
		token, err := jwt.ParseWithClaims(tokenString, &shared.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(shared.INVALID_SIGNING_METHOD)
			}
			return []byte(os.Getenv("AUTH_SECRET_KEY")), nil
		})
		if err != nil {
			response := shared.HandleError(err)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}
		claims, ok := token.Claims.(*shared.JWTClaim)
		if !ok || !token.Valid {
			response := shared.HandleError(errors.New(shared.AUTHORIZATION_HEADER_FORMAT_ERROR))
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		// set the user context for the downstream handlers
		user := claims.User

		c.Set("user", &user)
		c.Next()
	}
}
