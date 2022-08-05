package repository

import "github.com/ccallazans/url-shortener/internal/models"

type DBRepo interface {
	// Url Repo
	CreateUrl(newUrl models.Url) error
	GetAllUrls() ([]*models.Url, error)
	GetUrlByShort(short string) (*models.Url, error)
	UpdateUrlByShort(short string, newUrl string) error
	DeleteUrlByShort(short string) error
	UrlValueExists(value string, column string) bool

	// User Repo
	CreateUser(newUser models.User) error
	GetAllUsers() ([]*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUserByEmail(oldEmail string, newEmail string) error
	DeleteUserByEmail(email string) error
	UserValueExists(value string, column string) bool
}
