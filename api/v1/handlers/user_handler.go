package handlers

import (
	"context"
	"net/http"

	"github.com/ccallazans/url-shortener/internal/domain/mappers"
	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	userRequest, exists := c.MustGet("request").(*models.UserRequest)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "message": "failed to get user request"})
		return
	}

	err := h.userService.Save(context.TODO(), mappers.NewUserMapper().UserRequestToUser(userRequest))
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "conflict", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mappers.NewUserMapper().UserRequestToUserResponse(userRequest))
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userService.FindById(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusFound, gin.H{"error": "not found", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mappers.NewUserMapper().UserToUserResponse(user))
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {

	users, err := h.userService.FindAll(context.TODO())
	if err != nil {
		c.JSON(http.StatusFound, gin.H{"error": "not found", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mappers.NewUserMapper().UsersToUserResponses(users))
}
