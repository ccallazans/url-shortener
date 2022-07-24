package repositories

import (
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

func (r *UrlRepo) InsertUrlModel(newUrl models.UrlRequest) error {

	// Insert new url
	query := `INSERT INTO urls (url, hash, created_at) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, newUrl.Url, newUrl.Hash, newUrl.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *UrlRepo) GetAllUrls() ([]*models.Url, error) {

	// query all data
	query := `SELECT id, hash, url, created_at FROM urls`
	rows, err := r.db.Query(query)
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
			&newUrl.Hash,
			&newUrl.Url,
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

func (r *UrlRepo) GetByHash(hash string) (*models.Url, error) {
	// Query specific data
	query := `SELECT id, hash, url, created_at FROM urls WHERE hash = $1`
	row := r.db.QueryRow(query, hash)

	var newUrl models.Url
	err := row.Scan(
		&newUrl.ID,
		&newUrl.Hash,
		&newUrl.Url,
		&newUrl.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &newUrl, nil
}

func (r *UrlRepo) UpdateByHash(newUrl models.UrlRequest) error {
	// Query specific data
	query := `UPDATE urls SET url = $1 WHERE hash = $2`
	_, err := r.db.Exec(query, newUrl.Url, newUrl.Hash)
	if err != nil {
		return err
	}

	return nil
}

func (r *UrlRepo) HashExists(url string) bool {
	// Query specific data
	query := `SELECT id, hash, url, created_at FROM urls WHERE hash = $1`
	row := r.db.QueryRow(query, url)

	var newUrl models.Url
	err := row.Scan(
		&newUrl.ID,
		&newUrl.Hash,
		&newUrl.Url,
		&newUrl.CreatedAt,
	)
	if err != nil {
		return false
	}

	return true
}
