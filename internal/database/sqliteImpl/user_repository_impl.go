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

func (r *sqliteUserRepository) Save(ctx context.Context, user *models.UserEntity) error {

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

func (r *sqliteUserRepository) FindAll(ctx context.Context) ([]*models.UserEntity, error) {
	query := "SELECT * FROM users"

	statement, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := statement.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()


	var users []*models.UserEntity
	for rows.Next() {
		var user models.UserEntity

		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *sqliteUserRepository) FindById(ctx context.Context, id int) (*models.UserEntity, error) {

	query := "SELECT * FROM users WHERE id = ?"

	statement, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	user := models.UserEntity{}
	err = statement.QueryRowContext(ctx, id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *sqliteUserRepository) FindByUsername(ctx context.Context, username string) (*models.UserEntity, error) {
	
	query := "SELECT * FROM users WHERE username = ?"

	statement, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	user := models.UserEntity{}
	err = statement.QueryRowContext(ctx, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *sqliteUserRepository) Update(ctx context.Context, user *models.UserEntity) error {
	return nil
}

func (r *sqliteUserRepository) DeleteById(ctx context.Context, id int) error {
	return nil
}
