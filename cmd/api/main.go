package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ccallazans/url-shortener/internal/handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	psql := handlers.NewPostgresqlHandlers(db)
	handlers.NewHandlers(psql)
	routes := NewRouter()

	srv := &http.Server{
		Addr:              ":5000",
		Handler:           routes,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	fmt.Println("Starting Server on port 5000")
	srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}

func ConnectDB() (*sql.DB, error) {

	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
