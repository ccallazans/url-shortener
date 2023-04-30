package middleware

import (
	"errors"

	"github.com/ccallazans/url-shortener/api/v1/middleware/validation"
	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/utils"
	"github.com/gin-gonic/gin"
)

func ValidateUserRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var userRequest models.UserRequest

		err := c.ShouldBindJSON(&userRequest)
		if err != nil {
			info := utils.MatchError(errors.New(utils.BAD_REQUEST))
			c.AbortWithStatusJSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
			return
		}

		validation := validation.UsernameValidation{
			Next: &validation.PasswordValidation{},
		}

		err = validation.Execute(&userRequest)
		if err != nil {
			info := utils.MatchError(err)
			c.AbortWithStatusJSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
			return
		}

		c.Set("request", &userRequest)
		c.Next()
	}
}

func ValidateUrlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var urlRequest models.UrlRequest

		err := c.ShouldBindJSON(&urlRequest)
		if err != nil {
			info := utils.MatchError(errors.New(utils.BAD_REQUEST))
			c.AbortWithStatusJSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
			return
		}

		validation := validation.UrlValidation{}

		err = validation.Execute(&urlRequest)
		if err != nil {
			info := utils.MatchError(err)
			c.AbortWithStatusJSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
			return
		}

		c.Set("request", &urlRequest)
		c.Next()
	}
}
