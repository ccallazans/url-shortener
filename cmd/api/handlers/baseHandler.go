package handlers

import (
	"database/sql"

	"github.com/ccallazans/url-shortener/cmd/api/repositories"
	"github.com/ccallazans/url-shortener/models"
)

type BaseHandler struct {
	urlRepo  models.UrlRepository
	userRepo models.UserRepository
}

func NewBaseHandler(db *sql.DB) *BaseHandler {

	urlRepo := repositories.NewUrlRepo(db)
	userRepo := repositories.NewUserRepo(db)

	return &BaseHandler{
		urlRepo:  &urlRepo,
		userRepo: &userRepo,
	}
}
