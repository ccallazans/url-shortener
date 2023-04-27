package main

import (
	"ccallazans/internal/db"
	"ccallazans/internal/models"
	"ccallazans/internal/repository/impl/sqliteImpl"
	"log"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Panic("Error connecting to database", err)
	}

	db.AutoMigrate(&models.Url{})

	rp := sqliteImpl.NewSqliteRepository(db)
	rp.SaveUrl(models.Url{Url: "asdasd", Hash: "asdasd"})
}