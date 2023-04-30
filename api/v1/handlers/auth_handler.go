package handlers

import (
	"context"
	"net/http"

	"github.com/ccallazans/url-shortener/internal/domain/mappers"
	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/service"
	"github.com/ccallazans/url-shortener/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService service.UserServiceInterface
}

func NewAuthHandler(userService service.UserServiceInterface) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) AuthUser(c *gin.Context) {
	userRequest, exists := c.MustGet("request").(*models.UserRequest)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": utils.INTERNAL_SERVER_ERROR})
		return
	}

	token, err := h.userService.Auth(context.TODO(), mappers.NewUserMapper().UserRequestToUser(userRequest))
	if err != nil {
		info := utils.MatchError(err)
		c.JSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}