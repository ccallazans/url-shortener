package models

import "time"

type UrlShort struct {
	ID          int       `json:"-" db:"id"`
	Hash        string    `json:"hash" db:"hash"`
	OriginalUrl string    `json:"original_url" db:"original_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type UrlShortRequest struct {
	Hash        string    `json:"hash" db:"hash"`
	OriginalUrl string    `json:"original_url" db:"original_url"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
}


type UrlShortRepository interface {
	AddUrl(newUrl UrlShortRequest) error

	GetAllUrls() ([]*UrlShort, error)
	GetUrl(hash string) (*string, error)

	UpdateHash(newUrl UrlShortRequest) error

	VerifyExists(column string, value string) error
}
