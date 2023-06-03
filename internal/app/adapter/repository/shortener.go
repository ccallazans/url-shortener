package repository

import (
	"context"
	"myapi/internal/app/domain"
	"myapi/internal/app/domain/repository"

	"gorm.io/gorm"
)

type shortenerRepository struct {
	db *gorm.DB
}

func NewShortenerRepository(db *gorm.DB) repository.IShortener {
	return &shortenerRepository{
		db: db,
	}
}

func (r *shortenerRepository) Save(ctx context.Context, shortener *domain.Shortener) error {

	result := r.db.Create(shortener)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *shortenerRepository) FindAll(ctx context.Context) ([]*domain.Shortener, error) {

	var shorteners []*domain.Shortener
	result := r.db.Find(&shorteners)
	if result.Error != nil {
		return nil, result.Error
	}

	return shorteners, nil
}

func (r *shortenerRepository) FindByHash(ctx context.Context, hash string) (*domain.Shortener, error) {

	var shortener *domain.Shortener
	result := r.db.Find(&shortener, "hash = ?", hash).Limit(1)
	if result.Error != nil {
		return nil, result.Error
	}

	return shortener, nil
}

func (r *shortenerRepository) DeleteByHash(ctx context.Context, hash string) error {
	return nil
}
