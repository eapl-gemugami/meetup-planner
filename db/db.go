package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func GetDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	return db, err
}