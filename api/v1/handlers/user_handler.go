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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": utils.INTERNAL_SERVER_ERROR})
		return
	}

	err := h.userService.Save(context.TODO(), mappers.NewUserMapper().UserRequestToUser(userRequest))
	if err != nil {
		info := utils.MatchError(err)
		c.JSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
		return
	}

	c.JSON(http.StatusCreated, mappers.NewUserMapper().UserRequestToUserResponse(userRequest))
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userService.FindById(context.TODO(), id)
	if err != nil {
		info := utils.MatchError(err)
		c.JSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
		return
	}

	c.JSON(http.StatusCreated, mappers.NewUserMapper().UserToUserResponse(user))
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {

	users, err := h.userService.FindAll(context.TODO())
	if err != nil {
		info := utils.MatchError(err)
		c.JSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
		return
	}

	c.JSON(http.StatusCreated, mappers.NewUserMapper().UsersToUserResponses(users))
}
