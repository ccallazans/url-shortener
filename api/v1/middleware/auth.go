package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
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
		token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
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
		claims, ok := token.Claims.(*models.UserClaims)
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

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, exists := c.MustGet("user").(*models.User)
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": utils.INTERNAL_SERVER_ERROR})
			return
		}

		// Get user role from JWT or session
		userRole := user.Role.Role

		// Check if user has required role
		for _, role := range roles {
			if userRole == role {

				c.Next() // Authorized, proceed to next middleware
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden", "message": "You are not authorized to access this resource."})
	}
}
