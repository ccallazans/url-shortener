package repository

import (
	"context"
	"myapi/internal/app/domain"
)

type IShortener interface {
	Save(ctx context.Context, shortener *domain.Shortener) error
	FindAll(ctx context.Context) ([]*domain.Shortener, error)
	FindByHash(ctx context.Context, hash string) (*domain.Shortener, error)
	DeleteByHash(ctx context.Context, hash string) error
}