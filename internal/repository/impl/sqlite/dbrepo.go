package dbrepo

import (
	"github.com/ccallazans/url-shortener/internal/repository"
	"gorm.io/gorm"
)

type sqliteRepository struct {
	DB *gorm.DB
}

func NewSqliteRepository(db *gorm.DB) repository.RepositoryInterface {
	return &sqliteRepository{DB: db}
}
