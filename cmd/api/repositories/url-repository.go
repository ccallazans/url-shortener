package repositories

import (
	"database/sql"
	"fmt"
	"url-shortener/models"

	"github.com/jmoiron/sqlx"
)

type UrlRepo struct {
	db *sqlx.DB
}

func NewUrlRepo(db *sqlx.DB) *UrlRepo {
	return &UrlRepo{
		db: db,
	}
}

func (r *UrlRepo) AddUrl(newUrl models.UrlShortRequest) error {

	// Insert new url
	query := `INSERT INTO urls (original_url, hash, created_at) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, newUrl.OriginalUrl, newUrl.Hash, newUrl.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *UrlRepo) GetAllUrls() ([]*models.UrlShort, error) {

	// query all data
	query := `SELECT original_url, hash, created_at FROM urls`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []*models.UrlShort

	// Read all rows and save into 'urls'
	for rows.Next() {
		var addUrl models.UrlShort
		err = rows.Scan(
			&addUrl.OriginalUrl,
			&addUrl.Hash,
			&addUrl.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Append data
		urls = append(urls, &addUrl)
	}

	return urls, nil
}

func (r *UrlRepo) GetUrl(hash string) (*string, error) {

	// Query specific data
	query := `SELECT original_url FROM urls WHERE hash = $1`
	row, err := r.db.Query(query, hash)
	if err != nil {
		return nil, err
	}

	var url string

	// Append data
	err = row.Scan(&url)
	if err == sql.ErrNoRows {
		return nil, err
	}

	return &url, nil
}

func (r *UrlRepo) UpdateHash(newUrl models.UrlShortRequest) error {
	// Query specific data
	query := `UPDATE urls SET hash = $1 WHERE original_url = $2`
	_, err := r.db.Exec(query, newUrl.Hash, newUrl.OriginalUrl)
	if err != nil {
		return err
	}

	return nil
}

func (r *UrlRepo) VerifyExists(column string, value string) error {

	// Query data
	query := fmt.Sprintf(`SELECT * FROM urls WHERE %s = $1`, column)
	rows, err := r.db.Query(query, value)
	if err != nil {
		return err
	}

	// If exists
	if rows.Next() {
		return fmt.Errorf("%s already exists", column)
	}

	return nil
}
