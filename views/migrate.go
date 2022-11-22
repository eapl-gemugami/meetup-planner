package views

import (
	"fmt"
	"net/http"

	"github.com/eapl-gemugami/meetup-planner/db"
	"github.com/eapl-gemugami/meetup-planner/models"
)

func Migrate(w http.ResponseWriter, r *http.Request) {
	conn, err := db.GetDBConnection()

	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	conn.AutoMigrate(&models.Event{})

	fmt.Fprintf(w, "Sucessfully migrated")
}
