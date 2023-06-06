package handlers

import (
	"context"

	"myapi/internal/app/application/usecase"

	"myapi/internal/app/domain"
	"myapi/internal/app/shared"

	"net/http"

	"github.com/gin-gonic/gin"
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
	user, exists := c.MustGet("user").(*domain.UserAuth)
	if !exists {
		user.UUID = uuid.Nil
	}

	type ShortenerRequest struct {
		Url string `json:"url"  validate:"required"`
	}

	var shortenerRequest ShortenerRequest

	err := validateRequest(c, &shortenerRequest)
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	shortener, err := h.shortenerUsecase.Save(context.TODO(), domain.Shortener{Url: shortenerRequest.Url, User: user.UUID})
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusCreated, shortener.ToResponse())
}

func (h *ShortenerHandler) Redirect(c *gin.Context) {

	hash := c.Param("hash")

	url, err := h.shortenerUsecase.FindByHash(context.TODO(), hash)
	if err != nil {
		response := shared.HandleResponseError(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Url)
}
