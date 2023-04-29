package handlers

import (
	"context"
	"net/http"

	"github.com/ccallazans/url-shortener/internal/domain/mappers"
	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/service"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user request"})
		return
	}

	token, err := h.userService.Auth(context.TODO(), mappers.NewUserMapper().UserRequestToUser(userRequest))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}
