package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ccallazans/url-shortener/cmd/api/handlers"
	"github.com/ccallazans/url-shortener/cmd/api/routes"
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

	appHandler := handlers.NewBaseHandler(db)
	router := routes.ServeRouter(appHandler)

	fmt.Println("Starting Server")
	err = http.ListenAndServe(":5000", router)
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
