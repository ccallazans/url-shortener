package routes

import (
	"url-shortener/cmd/api/handlers"

	"github.com/go-chi/chi/v5"
)

func ServeRouter(teste *handlers.BaseHandler) *chi.Mux {
	router := chi.NewRouter()

	// Public Routes
	router.Group(func(r chi.Router) {
		r.Get("/", teste.GetAllUrlsHandler)
		r.Get("/{hash}", teste.GetUrlHandler)
		r.Post("/create", teste.InsertUrlHandler)
	})

	return router
}
