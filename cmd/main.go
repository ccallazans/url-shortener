package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ccallazans/url-shortener/internal/database/sqliteImpl"
	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := DBConnection()
	if err != nil {
		log.Panic(err)
	}

	userRepository := sqliteImpl.NewSqliteUserRepository(db)
	// urlRepository := sqliteImpl.NewSqliteUrlRepository(db)

	err = userRepository.Save(context.Background(), &models.User{Username: "ciro", Password: "1234"})
	if err != nil {
		log.Panic(err)
	}

	myuser, err := userRepository.FindById(context.Background(), 12)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(myuser)

	fmt.Println(err)
}

func DBConnection() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sql, err := utils.ReadSqlFile("internal/database/schema/schema.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(sql)
	if err != nil {
		return nil, err
	}

	return db, nil

}
