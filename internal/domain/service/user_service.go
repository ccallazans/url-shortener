package service

import (
	"context"

	"github.com/ccallazans/url-shortener/internal/domain/models"
)

type UserServiceInterface interface {
	Save(ctx context.Context, user *models.User) error
	FindAll(ctx context.Context) ([]*models.User, error)
	FindById(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	DeleteById(ctx context.Context, id int) error
}
