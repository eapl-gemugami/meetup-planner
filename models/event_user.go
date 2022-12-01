package models

import (
	"gorm.io/gorm"
)

type EventUser struct {
	gorm.Model

	EventID int
	Event   Event

	Name string `gorm:"unique"`
}
