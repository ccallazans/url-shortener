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

	query := "INSERT INTO urls (url, hash, user_id) VALUES (?, ?, ?)"

	statement, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, url.Url, url.Hash, url.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteUrlRepository) FindById(ctx context.Context, id int) (*models.Url, error) {

	query := "SELECT * FROM urls WHERE id = ?"

	statement, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	url := models.Url{}
	err = statement.QueryRowContext(ctx, id).Scan(&url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *sqliteUrlRepository) Update(ctx context.Context, url *models.Url) error {
	return nil
}

func (r *sqliteUrlRepository) DeleteById(ctx context.Context, id int) error {
	return nil
}
