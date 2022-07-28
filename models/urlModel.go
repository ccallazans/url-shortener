package models

import "time"

type Url struct {
	ID        int       `json:"-" db:"id"`
	Short     string    `json:"short" db:"short"`
	Url       string    `json:"url" db:"url"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UrlResponse struct {
	Short     string    `json:"short" db:"short"`
	Url       string    `json:"url" db:"url"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UrlRepository interface {
	CreateUrl(newUrl Url) error

	GetAllUrls() ([]*Url, error)
	GetUrlByShort(short string) (*Url, error)

	UpdateUrlByShort(short string, newUrl string) error

	DeleteUrlByShort(short string) error

	ValueExists(value string, column string) bool
}
