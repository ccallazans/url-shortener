package handlers

import "url-shortener/models"

type BaseHandler struct {
	urlRepo  models.UrlShortRepository
	userRepo models.UserRepository
}

func NewBaseHandler(urlRepo models.UrlShortRepository, userRepo models.UserRepository) *BaseHandler {
	return &BaseHandler{
		urlRepo:  urlRepo,
		userRepo: userRepo,
	}
}
