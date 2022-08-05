package models

type UrlRepository interface {
	CreateUrl(newUrl Url) error

	GetAllUrls() ([]*Url, error)
	GetUrlByShort(short string) (*Url, error)

	UpdateUrlByShort(short string, newUrl string) error

	DeleteUrlByShort(short string) error

	ValueExists(value string, column string) bool
}
