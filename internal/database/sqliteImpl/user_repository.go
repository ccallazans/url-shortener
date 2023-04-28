package sqliteImpl

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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
		return fmt.Errorf("deu erro1: ", err)
	}

	_, err = statement.ExecContext(ctx, user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("deu erro2: ", err)
	}

	return nil
}

func (r *sqliteUserRepository) FindById(ctx context.Context, id int) (*models.User, error) {

	query := "SELECT * FROM users WHERE id = ?"

	statement, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("deu erro1: ", err)
	}

	rows, err := statement.QueryContext(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("deu erro2: ", err)
	}

	ourUser := models.User{}
	for rows.Next() {
		fmt.Println("rows -> ", rows)
		err = rows.Scan(&ourUser.ID, &ourUser.Username, &ourUser.Password)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &ourUser, nil
}

func (r *sqliteUserRepository) Update(ctx context.Context, user *models.User) error {
	return nil
}

func (r *sqliteUserRepository) DeleteById(ctx context.Context, id int) error {
	return nil
}
