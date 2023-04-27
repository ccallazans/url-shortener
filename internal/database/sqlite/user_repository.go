package databse

import (
	"context"
	"database/sql"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/repository"
)

type sqliteUserRepository struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) repository.UserRepository {
	return &sqliteUserRepository{
		db: db,
	}
}

func (r *sqliteUserRepository) Save(ctx context.Context, user *models.User) error {
	return nil
}

func (r *sqliteUserRepository) FindById(ctx context.Context, id int) (*models.User, error) {
	return nil, nil
}

func (r *sqliteUserRepository) Update(ctx context.Context, user *models.User) error {
	return nil
}

func (r *sqliteUserRepository) DeleteById(ctx context.Context, id int) error {
	return nil
}
