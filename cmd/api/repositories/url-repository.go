package repositories

import (
	"database/sql"
	"errors"
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

func (r *UrlRepo) AddUrl(newUrl models.UrlShort) error {

	// Verify if hash already exists
	query := `SELECT original_url, hash 
			  FROM urls 
			  WHERE hash = $1`
	rows, err := r.db.Query(query, newUrl.Hash)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return errors.New("error generating hash, try again")
	}

	// Insert new url
	query = `INSERT INTO urls (original_url, hash, created_at) 
			 VALUES ($1, $2, $3);`
	_, err = r.db.Exec(query, newUrl.OriginalUrl, newUrl.Hash, newUrl.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *UrlRepo) GetAllUrls() ([]*models.UrlShort, error) {

	// query all data
	query := `SELECT original_url, hash, created_at
			  FROM urls`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []*models.UrlShort

	// read all rows and save into 'urls'
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

		urls = append(urls, &addUrl)
	}

	return urls, nil
}

func (r *UrlRepo) GetUrl(hash string) (*string, error) {

	// query specific data
	query := `SELECT original_url
			  FROM urls
			  WHERE hash = $1`
	row := r.db.QueryRow(query, hash)

	var original_url string

	err := row.Scan(&original_url)
	if err == sql.ErrNoRows {
		return nil, err
	}

	return &original_url, nil
}
