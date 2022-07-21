package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func createDsn() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DB_DRIVER"), fmt.Sprintf(
		"host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}

func OpenDB() (*sql.DB, error) {

	driver, dsn := createDsn()

	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error database ping")
		return nil, err
	}

	return db, nil
}
