package models

import "time"

type Url struct {
	ID        int       `json:"-" db:"id"`
	Hash      string    `json:"hash" db:"hash"`
	Url       string    `json:"url" db:"url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UrlRequest struct {
	Hash      string    `json:"hash" db:"hash"`
	Url       string    `json:"url" db:"url"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}

type UrlShortRepository interface {
	InsertUrlModel(newUrl UrlRequest) error

	GetAllUrls() ([]*Url, error)
	GetByHash(hash string) (*Url, error)

	UpdateByHash(newUrl UrlRequest) error

	HashExists(url string) bool
}
