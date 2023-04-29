package middleware

import (
	"net/http"

	"github.com/ccallazans/url-shortener/api/v1/middleware/validation"
	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func ValidateUserRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var userRequest models.UserRequest

		err := c.ShouldBindJSON(&userRequest)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		validation := validation.UsernameValidation{
			Next: &validation.PasswordValidation{},
		}

		err = validation.Execute(&userRequest)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Set("request", &userRequest)
		c.Next()
	}
}
