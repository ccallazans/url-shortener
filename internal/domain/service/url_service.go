package service

import (
	"context"

	"github.com/ccallazans/url-shortener/internal/domain/models"
)

type UrlServiceInterface interface {
	Save(ctx context.Context, url *models.Url) error
	FindAll(ctx context.Context) ([]*models.Url, error)
	FindById(ctx context.Context, id string) (*models.Url, error)
	FindByHash(ctx context.Context, hash string) (*models.Url, error)
	Update(ctx context.Context, url *models.Url) error
	DeleteById(ctx context.Context, id int) error
}
