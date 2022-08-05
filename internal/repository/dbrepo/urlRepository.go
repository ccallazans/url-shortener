package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/ccallazans/url-shortener/internal/models"
)

func (r *postgresDBRepo) CreateUrl(newUrl models.Url) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Insert new url
	query := `INSERT INTO urls (short, url, created_by, updated_at, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.ExecContext(ctx, query, newUrl.Short, newUrl.Url, newUrl.CreatedBy, newUrl.UpdatedAt, newUrl.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDBRepo) GetAllUrls() ([]*models.Url, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// query all data
	query := `SELECT id, short, url, created_by, updated_at, created_at FROM urls`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allUrl []*models.Url

	// Read all rows and save into 'urls'
	for rows.Next() {
		var newUrl models.Url
		err = rows.Scan(
			&newUrl.ID,
			&newUrl.Short,
			&newUrl.Url,
			&newUrl.CreatedBy,
			&newUrl.UpdatedAt,
			&newUrl.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Append data
		allUrl = append(allUrl, &newUrl)
	}

	return allUrl, nil
}

func (r *postgresDBRepo) GetUrlByShort(short string) (*models.Url, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := `SELECT id, short, url, created_by, updated_at, created_at FROM urls WHERE short = $1`
	row := r.DB.QueryRowContext(ctx, query, short)

	var newUrl models.Url
	err := row.Scan(
		&newUrl.ID,
		&newUrl.Short,
		&newUrl.Url,
		&newUrl.CreatedBy,
		&newUrl.UpdatedAt,
		&newUrl.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &newUrl, nil
}

func (r *postgresDBRepo) UpdateUrlByShort(short string, newUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := `UPDATE urls SET url = $1 WHERE short = $2`
	_, err := r.DB.ExecContext(ctx, query, newUrl, short)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDBRepo) DeleteUrlByShort(short string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := `DELETE FROM urls WHERE short = $1`
	_, err := r.DB.ExecContext(ctx, query, short)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDBRepo) UrlValueExists(value string, column string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query specific data
	query := fmt.Sprintf(`SELECT %s FROM urls WHERE %s = $1`, column, column)
	row := r.DB.QueryRowContext(ctx, query, value)

	var newHash models.Url
	err := row.Scan(
		&newHash.Short,
	)
	if err != nil {
		return false
	}

	return true
}
