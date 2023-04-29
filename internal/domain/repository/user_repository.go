package repository

import (
	"context"

	"github.com/ccallazans/url-shortener/internal/domain/models"
)

type UserRepositoryInterface interface {
	Save(ctx context.Context, user *models.UserEntity) error
	FindAll(ctx context.Context) ([]*models.UserEntity, error)
	FindById(ctx context.Context, id int) (*models.UserEntity, error)
	FindByUsername(ctx context.Context, username string) (*models.UserEntity, error)
	Update(ctx context.Context, user *models.UserEntity) error
	DeleteById(ctx context.Context, id int) error
}