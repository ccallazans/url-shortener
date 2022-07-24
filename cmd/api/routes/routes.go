package routes

import (
	"url-shortener/cmd/api/handlers"

	"github.com/go-chi/chi/v5"
)

func ServeRouter(hand *handlers.BaseHandler) *chi.Mux {
	router := chi.NewRouter()

	// Public Routes
	router.Group(func(r chi.Router) {
		r.NotFound(hand.NotFoundHandler)
		r.Get("/", hand.GetAllUrlsHandler)
		r.Get("/{hash}", hand.GetUrlHandler)
	})

	// Protected Routes
	router.Group(func(r chi.Router) {
		r.Post("/create", hand.InsertUrlHandler)
		r.Post("/edit", hand.UpdateHashHandler)
	})


	return router
}
