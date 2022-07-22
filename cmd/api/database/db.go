package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func ConnectDB() (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
