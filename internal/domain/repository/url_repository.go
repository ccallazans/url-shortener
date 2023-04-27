package repository

import (
	"context"

	"github.com/ccallazans/url-shortener/internal/domain/models"
)

type UrlRepository interface {
	Save(ctx context.Context, url *models.Url) error
	FindById(ctx context.Context, id int) (*models.Url, error)
	Update(ctx context.Context, url *models.Url) error
	DeleteById(ctx context.Context, id int) error
}
