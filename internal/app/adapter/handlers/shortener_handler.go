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
	"github.com/google/uuid"
)

type ShortenerHandler struct {
	shortenerUsecase usecase.ShortenerUsecase
}

func NewShortenerHandler(shortenerUsecase usecase.ShortenerUsecase) *ShortenerHandler {
	return &ShortenerHandler{
		shortenerUsecase: shortenerUsecase,
	}
}

func (h *ShortenerHandler) CreateShortener(c *gin.Context) {
	user, exists := c.MustGet("user").(*shared.UserAuth)
	if !exists {
		user.UUID = uuid.Nil
	}

	type ShortenerRequest struct {
		Url string `json:"url"  validate:"required"`
	}

	var shortenerRequest ShortenerRequest

	err := c.ShouldBindJSON(&shortenerRequest)
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	err = validator.New().Struct(shortenerRequest)
	if err != nil {
		response := shared.HandleResponseError(errors.New(shared.BAD_REQUEST))
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	shortener, err := h.shortenerUsecase.Save(context.TODO(), &domain.Shortener{Url: shortenerRequest.Url, User: user.UUID})
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusCreated, shortener)
}

func (h *ShortenerHandler) Redirect(c *gin.Context) {

	hash := c.Param("hash")

	url, err := h.shortenerUsecase.FindByHash(context.TODO(), hash)
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, hash)
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Url)
}
