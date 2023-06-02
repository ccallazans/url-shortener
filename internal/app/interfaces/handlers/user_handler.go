package handlers

import (
	"context"
	"errors"
	"myapi/internal/app/application/usecase"
	"myapi/internal/app/domain"
	"myapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

type UserRequest struct {
	Username string `json:username`
	Password string `json:password`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		info := utils.MatchError(errors.New(utils.BAD_REQUEST))
		c.AbortWithStatusJSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
		return
	}

	err = h.userUsecase.Save(context.TODO(), &domain.User{Username: userRequest.Username, Password: userRequest.Password})
	if err != nil {
		info := utils.MatchError(err)
		c.JSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
		return
	}

	c.JSON(http.StatusCreated, userRequest)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userUsecase.FindById(context.TODO(), id)
	if err != nil {
		info := utils.MatchError(err)
		c.JSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {

	users, err := h.userUsecase.FindAll(context.TODO())
	if err != nil {
		info := utils.MatchError(err)
		c.JSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
		return
	}

	c.JSON(http.StatusCreated, users)
}
