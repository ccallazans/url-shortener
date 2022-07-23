package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortener/cmd/api/handlers"
	"url-shortener/cmd/api/repositories"
	"url-shortener/cmd/api/routes"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	urlRepo := repositories.NewUrlRepo(db)
	h := handlers.NewBaseHandler(urlRepo)
	r := routes.ServeRouter(h)

	fmt.Println("Starting Server")
	err = http.ListenAndServe(":5000", r)
	fmt.Println(err)
}

func ConnectDB() (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", os.Getenv("DB_DSN"))
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
