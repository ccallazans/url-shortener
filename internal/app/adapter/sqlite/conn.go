package sqliteimpl

import (
	"myapi/internal/app/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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