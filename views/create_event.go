package views

import (
	"fmt"
	"time"
	"net/http"
  "math/rand"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

	"github.com/eapl-gemugami/meetup-planner/models"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("Failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&models.Event{})

	// Create random name
	rand.Seed(time.Now().UnixNano())
	eventName := randSeq(10)

  // Create event
  db.Create(&models.Event{Name: eventName, TimeStart: time.Now().Unix()})

	fmt.Fprintf(w, "Create event: %v\n", eventName)
}
