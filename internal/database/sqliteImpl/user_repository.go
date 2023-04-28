package sqliteImpl

import (
	"context"
	"database/sql"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/repository"
)

type sqliteUserRepository struct {
	db *sql.DB
}

func NewSqliteUserRepository(db *sql.DB) repository.UserRepositoryInterface {
	return &sqliteUserRepository{
		db: db,
	}
}

func (r *sqliteUserRepository) Save(ctx context.Context, user *models.User) error {

	query := "INSERT INTO users (username, password) VALUES (?, ?)"

	statement, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteUserRepository) FindById(ctx context.Context, id int) (*models.User, error) {

	query := "SELECT * FROM users WHERE id = ?"

	statement, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	user := models.User{}
	err = statement.QueryRowContext(ctx, id).Scan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *sqliteUserRepository) Update(ctx context.Context, user *models.User) error {
	return nil
}

func (r *sqliteUserRepository) DeleteById(ctx context.Context, id int) error {
	return nil
}
