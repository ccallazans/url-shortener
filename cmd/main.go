package main

import (
	"log"
	"myapi/internal/app/adapter/handlers"
	"myapi/internal/app/domain"

	"github.com/joho/godotenv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	db, err := DBConnection()
	if err != nil {
		log.Panic(err)
	}

	server := handlers.RouterConfig(db)
	server.Run()
}

func DBConnection() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.User{}, &domain.Shortener{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
