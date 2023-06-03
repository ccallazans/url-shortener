package main

import (
	"log"
	"myapi/internal/app/adapter/handlers"
	sqliteimpl "myapi/internal/app/adapter/sqlite"

	"github.com/joho/godotenv"
)

const (
	ERROR_LOAD_ENV_VARIABLES = "Error loading .env file"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(ERROR_LOAD_ENV_VARIABLES)
	}
}

func main() {

	db, err := sqliteimpl.DBConnection()
	if err != nil {
		log.Panic(err)
	}

	server := handlers.RouterConfig(db)
	server.Run()
}
