package main

import (
	"log"
	"net/http"
	"url-shortener/cmd/api/database"
	"url-shortener/cmd/api/routes"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := routes.ServeRouter()

	http.ListenAndServe(":5000", r)
}