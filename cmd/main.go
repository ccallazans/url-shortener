package main

import (
	"log"
	"os"

	"github.com/ccallazans/url-shortener/api/v1/router"
	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	server := router.Config(db)

	err = server.Run()
	if err != nil {
		log.Panic(err)
	}
}

func DBConnection() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_STRING")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.UserEntity{}, &models.Role{}, &models.Url{})
	if err != nil {
		return nil, err
	}

	// sql, err := utils.ReadSqlFile(os.Getenv("DB_SCHEMA"))
	// if err != nil {
	// 	return nil, err
	// }

	// db = db.Raw(sql)
	// if err != nil {
	// 	return nil, err
	// }

	return db, nil

}
