package handlers

import (
	"database/sql"

	"github.com/ccallazans/url-shortener/internal/repository"
	"github.com/ccallazans/url-shortener/internal/repository/dbrepo"
)

type DBRepo struct {
	DB repository.DBRepo
}

// NewPostgresqlHandlers creates db repo for postgres
func NewHandlers(db *sql.DB) *DBRepo {
	return &DBRepo{
		DB:  dbrepo.NewPostgresRepo(db),
	}
}
