package routes

import (
	"url-shortener/cmd/api/handlers"

	"github.com/go-chi/chi/v5"
)

func ServeRouter(hand *handlers.BaseHandler) *chi.Mux {
	router := chi.NewRouter()

	// Public Routes
	router.Group(func(r chi.Router) {
		r.Get("/", hand.GetAllHandler)
		r.Get("/{hash}", hand.GetByHashHandler)
	})

	// Protected Routes
	router.Group(func(r chi.Router) {
		r.Post("/create", hand.InsertUrlHandler)
		r.Post("/editurl", hand.UpdateByHashHandler)
	})


	return router
}
