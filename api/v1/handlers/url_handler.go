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

type UrlHandler struct {
	urlService service.UrlServiceInterface
}

func NewUrlHandler(urlService service.UrlServiceInterface) *UrlHandler {
	return &UrlHandler{
		urlService: urlService,
	}
}

func (h *UrlHandler) CreateUrl(c *gin.Context) {
	user, exists := c.MustGet("user").(*models.User)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": utils.INTERNAL_SERVER_ERROR})
		return
	}

	urlRequest, exists := c.MustGet("request").(*models.UrlRequest)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": utils.INTERNAL_SERVER_ERROR})
		return
	}

	urlRequest.UserID = user.ID

	err := h.urlService.Save(context.TODO(), mappers.NewUrlMapper().UrlRequestToUrl(urlRequest))
	if err != nil {
		info := utils.MatchError(err)
		c.JSON(info.Status, gin.H{"error": info.ErrorType, "message": info.Message})
		return
	}
}
