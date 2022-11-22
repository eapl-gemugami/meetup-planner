package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	return db, err
}
