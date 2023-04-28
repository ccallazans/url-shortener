package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ccallazans/url-shortener/api/v1/router"
	"github.com/ccallazans/url-shortener/internal/utils"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	db, err := DBConnection()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	server := router.Config(db)

	err = server.Run()
	if err != nil {
		log.Panic(err)
	}
}

func DBConnection() (*sql.DB, error) {

	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_STRING"))
	if err != nil {
		return nil, err
	}

	sql, err := utils.ReadSqlFile(os.Getenv("DB_SCHEMA"))
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(sql)
	if err != nil {
		return nil, err
	}

	return db, nil

}
