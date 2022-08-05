package main

import (
	"net/http"

	"github.com/ccallazans/url-shortener/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(RateLimitIp)

	// Public Routes
	router.Group(func(r chi.Router) {
		r.Get("/", handlers.MyRepo.GetAllUrlsHandler)
		r.Get("/{short}", handlers.MyRepo.GetUrlByShortHandler)

		r.Post("/register", handlers.MyRepo.CreateUserHandler)
		r.Post("/login", handlers.MyRepo.AuthUserHandler)
	})

	// Protected Routes
	router.Group(func(r chi.Router) {
		r.Use(IsAuthorized)
		r.Post("/create", handlers.MyRepo.CreateUrlHandler)
		r.Post("/edit", handlers.MyRepo.UpdateUrlByShortHandler)
	})

	return router
}
