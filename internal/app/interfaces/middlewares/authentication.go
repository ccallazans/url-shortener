package middlewares

import (
	"errors"
	"myapi/internal/app/interfaces/auth"
	"myapi/utils"
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
			info := utils.MatchError(errors.New(utils.AUTHORIZATION_HEADER_EMPTY))
			c.JSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
			return
		}
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			info := utils.MatchError(errors.New(utils.AUTHORIZATION_HEADER_FORMAT_ERROR))
			c.AbortWithStatusJSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
			return
		}
		tokenString := authHeaderParts[1]

		// validate the token
		token, err := jwt.ParseWithClaims(tokenString, &auth.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(utils.INVALID_SIGNING_METHOD)
			}
			return []byte(os.Getenv("AUTH_SECRET_KEY")), nil
		})
		if err != nil {
			info := utils.MatchError(errors.New(utils.AUTHORIZATION_HEADER_FORMAT_ERROR))
			c.AbortWithStatusJSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
			return
		}
		claims, ok := token.Claims.(*auth.JWTClaim)
		if !ok || !token.Valid {
			info := utils.MatchError(errors.New(utils.AUTHORIZATION_HEADER_FORMAT_ERROR))
			c.AbortWithStatusJSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
			return
		}

		// set the user context for the downstream handlers
		user := claims.User

		c.Set("user", &user)
		c.Next()
	}
}
