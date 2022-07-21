package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ccallazans/url-shortener/cmd/api/routes"
	"github.com/ccallazans/url-shortener/models"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)


type config struct {
	port string
	env  string
}

type application struct {
	config config
	logger *log.Logger
	router *chi.Mux
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create server configuration
	cfg := config{
		port: os.Getenv("PORT"),
		env:  os.Getenv("ENV"),
	}

	// Create log
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Initiate db connection
	db, err := models.OpenDB()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// Create app
	app := application{
		config: cfg,
		logger: logger,
		router: routes.NewRouter(),
	}

	// Create server
	server := &http.Server{
		Addr:         cfg.port,
		Handler:      app.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.port)

	// Start server
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
