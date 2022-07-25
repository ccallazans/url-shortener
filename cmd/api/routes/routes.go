package routes

import (
	"url-shortener/cmd/api/handlers"
	mid "url-shortener/cmd/api/middleware"

	"github.com/go-chi/chi/v5"
)

func ServeRouter(hand *handlers.BaseHandler) *chi.Mux {
	router := chi.NewRouter()

	// Public Routes
	router.Group(func(r chi.Router) {
		r.Get("/", hand.GetAllHandler)
		r.Get("/{hash}", hand.GetByHashHandler)
		r.Post("/register", hand.CreateUserHandler)
		r.Post("/login", hand.AuthUserHandler)
	})

	// Protected Routes
	router.Group(func(r chi.Router) {
		r.Use(mid.IsAuthorized)
		r.Post("/create", hand.InsertUrlHandler)
		r.Post("/edit", hand.UpdateByHashHandler)
	})

	return router
}
