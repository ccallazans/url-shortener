package routes

import (
	"url-shortener/cmd/api/handlers"

	"github.com/go-chi/chi/v5"
)

func ServeRouter() *chi.Mux {
	router := chi.NewRouter()

	// Public Routes
	router.Group(func(r chi.Router) {
		r.Get("/", handlers.Teste)
	})

	return router
}
