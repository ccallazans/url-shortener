package routes

import (
	"github.com/ccallazans/url-shortener/cmd/api/handlers"
	"github.com/ccallazans/url-shortener/cmd/api/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func ServeRouter(hand *handlers.BaseHandler) *chi.Mux {
	router := chi.NewRouter()
	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.RateLimitIp)

	// Public Routes
	router.Group(func(r chi.Router) {
		r.Get("/", hand.GetAllUrlsHandler)
		r.Get("/{short}", hand.GetUrlByShortHandler)

		r.Post("/register", hand.CreateUserHandler)
		r.Post("/login", hand.AuthUserHandler)
	})

	// Protected Routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.IsAuthorized)
		r.Post("/create", hand.CreateUrlHandler)
		r.Post("/edit", hand.UpdateUrlByShortHandler)
	})

	return router
}
