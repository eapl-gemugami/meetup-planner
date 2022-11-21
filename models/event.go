package models

import (
	"time"
	"gorm.io/gorm"
)

// https://golangbyexample.com/exported-unexported-fields-struct-go/
// https://gorm.io/docs/models.html

type Event struct {
	gorm.Model
	Name string
	TimeStart int64 // Unix Timestamp
	TimeEnd int64 // Unix Timestamp
}
