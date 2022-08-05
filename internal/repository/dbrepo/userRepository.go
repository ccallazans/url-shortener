package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/ccallazans/url-shortener/internal/models"
)

func (r *postgresDBRepo) CreateUser(newUser models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Insert new url
	query := `INSERT INTO users (uuid, email, password, updated_at, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.ExecContext(ctx, query, newUser.UUID, newUser.Email, newUser.Password, newUser.UpdatedAt, newUser.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDBRepo) GetAllUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// query all data
	query := `SELECT uuid, email, updated_at, created_at FROM users`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allUsers []*models.User

	// Read all rows and save into 'urls'
	for rows.Next() {
		var oneUser models.User
		err = rows.Scan(
			&oneUser.UUID,
			&oneUser.Email,
			&oneUser.UpdatedAt,
			&oneUser.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Append data
		allUsers = append(allUsers, &oneUser)
	}

	return allUsers, nil
}

func (r *postgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := `SELECT uuid, email, password, updated_at, created_at FROM users WHERE email = $1`
	row := r.DB.QueryRowContext(ctx, query, email)

	var oneUser models.User
	err := row.Scan(
		&oneUser.UUID,
		&oneUser.Email,
		&oneUser.Password,
		&oneUser.UpdatedAt,
		&oneUser.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &oneUser, nil
}

func (r *postgresDBRepo) UpdateUserByEmail(oldEmail string, newEmail string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := `UPDATE users SET email = $1 WHERE email = $2`
	_, err := r.DB.ExecContext(ctx, query, newEmail, oldEmail)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDBRepo) DeleteUserByEmail(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := `DELETE FROM users WHERE email = $1`
	_, err := r.DB.ExecContext(ctx, query, email)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDBRepo) UserValueExists(value string, column string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := fmt.Sprintf(`SELECT %s FROM users WHERE %s = $1`, column, column)
	row := r.DB.QueryRowContext(ctx, query, value)

	var user models.User
	err := row.Scan(
		&user.Email,
	)
	if err != nil {
		return false
	}

	return true
}
