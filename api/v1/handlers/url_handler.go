package handlers

import (
	"context"
	"net/http"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	urlService service.UrlServiceInterface
}

func NewUrlHandler(urlService service.UrlServiceInterface) *UrlHandler {
	return &UrlHandler{
		urlService: urlService,
	}
}

func (h *UrlHandler) CreateUrl(c *gin.Context) {
	var urlRequest models.UrlRequest

	if err := c.BindJSON(&urlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.urlService.Save(context.TODO(), urlRequest.ToUrl()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"url": urlRequest})
}
