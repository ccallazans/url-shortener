package handlers

import (
	"database/sql"

	"github.com/ccallazans/url-shortener/internal/repository"
	"github.com/ccallazans/url-shortener/internal/repository/dbrepo"
)

//Repo is the repository
var MyRepo *DBRepo

type DBRepo struct {
	DB repository.DBRepo
}

// NewHandlers creates the handlers
func NewHandlers(repo *DBRepo) {
	MyRepo = repo
}

// NewPostgresqlHandlers creates db repo for postgres
func NewPostgresqlHandlers(db *sql.DB) *DBRepo {
	return &DBRepo{
		DB:  dbrepo.NewPostgresRepo(db),
	}
}
