package sqliteImpl

import (
	"context"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/repository"
	"gorm.io/gorm"
)

type sqliteUrlRepository struct {
	db *gorm.DB
}

func NewSqliteUrlRepository(db *gorm.DB) repository.UrlRepositoryInterface {
	return &sqliteUrlRepository{
		db: db,
	}
}

func (r *sqliteUrlRepository) Save(ctx context.Context, url *models.Url) error {

	result := r.db.Create(&url)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *sqliteUrlRepository) FindAll(ctx context.Context) ([]*models.Url, error) {

	var urls []*models.Url
	result := r.db.Find(&urls)
	if result.Error != nil {
		return nil, result.Error
	}

	return urls, nil
}

func (r *sqliteUrlRepository) FindById(ctx context.Context, id int) (*models.Url, error) {

	var url models.Url
	result := r.db.First(&url, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &url, nil
}

func (r *sqliteUrlRepository) FindByHash(ctx context.Context, hash string) (*models.Url, error) {

	var url models.Url
	result := r.db.First(&url, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}

	return &url, nil
}

func (r *sqliteUrlRepository) Update(ctx context.Context, url *models.Url) error {
	return nil
}

func (r *sqliteUrlRepository) DeleteById(ctx context.Context, id int) error {
	return nil
}
