package models

import (
	"gorm.io/gorm"
)

// https://golangbyexample.com/exported-unexported-fields-struct-go/
// https://gorm.io/docs/models.html

type Event struct {
	gorm.Model
	Name string

	PublicCode string `gorm:"unique"`
	AdminCode string `gorm:"unique"`

	TimeStart int64 // Unix Timestamp
	TimeEnd int64 // Unix Timestamp
	TimeInterval int
}
