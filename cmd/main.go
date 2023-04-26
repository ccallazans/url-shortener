package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	_, err := gorm.Open(sqlite.Open("database.db"))
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}
}

func main() {

}
