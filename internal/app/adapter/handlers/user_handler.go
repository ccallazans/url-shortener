package handlers

import (
	"context"
	"errors"
	"myapi/internal/app/application/usecase"
	"myapi/internal/app/domain"
	"myapi/internal/app/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	type UserRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var userRequest UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	err = validator.New().Struct(userRequest)
	if err != nil {
		response := shared.HandleResponseError(errors.New(shared.BAD_REQUEST))
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	err = h.userUsecase.Save(context.TODO(), domain.User{Username: userRequest.Username, Password: userRequest.Password})
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusCreated, userRequest)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	username := c.Param("username")

	user, err := h.userUsecase.FindByUsername(context.TODO(), username)
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusCreated, user.ToResponse())
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {

	users, err := h.userUsecase.FindAll(context.TODO())
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusCreated, domain.UsersToResponse(users))
}

func (h *UserHandler) AuthUser(c *gin.Context) {

	type UserRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var userRequest UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	err = validator.New().Struct(userRequest)
	if err != nil {
		response := shared.HandleResponseError(errors.New(shared.BAD_REQUEST))
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	token, err := h.userUsecase.Auth(context.TODO(), domain.User{Username: userRequest.Username, Password: userRequest.Password})
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}
