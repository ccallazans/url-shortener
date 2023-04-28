package sqliteImpl

import (
	"context"
	"database/sql"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/repository"
)

type sqliteUrlRepository struct {
	db *sql.DB
}

func NewSqliteUrlRepository(db *sql.DB) repository.UrlRepositoryInterface {
	return &sqliteUrlRepository{
		db: db,
	}
}

func (r *sqliteUrlRepository) Save(ctx context.Context, url *models.Url) error {
	return nil
}

func (r *sqliteUrlRepository) FindById(ctx context.Context, id int) (*models.Url, error) {
	return nil, nil
}

func (r *sqliteUrlRepository) Update(ctx context.Context, url *models.Url) error {
	return nil
}

func (r *sqliteUrlRepository) DeleteById(ctx context.Context, id int) error {
	return nil
}
