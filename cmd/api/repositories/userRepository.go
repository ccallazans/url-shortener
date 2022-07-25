package repositories

import (
	"context"
	"time"
	"url-shortener/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(userModel models.User) error {

	// Insert new url
	query := `INSERT INTO auth (uuid, email, password, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, userModel.UUID, userModel.Email, userModel.Password, userModel.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := `SELECT uuid, email, password FROM auth WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)

	var user models.User
	err := row.Scan(
		&user.UUID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) EmailExists(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := `SELECT email FROM auth WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)

	var verify models.User
	err := row.Scan(
		&verify.Email,
	)
	if err != nil {
		return false
	}

	return true
}

func (r *UserRepo) UUIDExists(uuid string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := `SELECT uuid FROM auth WHERE uuid = $1`
	row := r.db.QueryRowContext(ctx, query, uuid)

	var verify models.User
	err := row.Scan(
		&verify.UUID,
	)
	if err != nil {
		return false
	}

	return true
}
