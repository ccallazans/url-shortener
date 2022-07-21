package routes

import (
	"net/http"

	"github.com/ccallazans/url-shortener/cmd/api/handlers"
	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	// Public Routes
	router.Group(func(r chi.Router) {
		r.MethodFunc(http.MethodGet, "/status", handlers.GetStatus)
	})

	return router
}
