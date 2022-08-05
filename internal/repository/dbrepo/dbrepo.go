package dbrepo

import (
	"database/sql"

	"github.com/ccallazans/url-shortener/internal/repository"
)

type postgresDBRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(db *sql.DB) repository.DBRepo {
	return &postgresDBRepo{
		DB: db,
	}
}

