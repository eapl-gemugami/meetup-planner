package views

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func GetDataRange(w http.ResponseWriter, r *http.Request) {
	// Define a starting and end date, and get a list of
	// hours in between

	loc, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		fmt.Printf("Error with the timezone")
	} else {
		t := time.Date(2022, 10, 10, 23, 0, 0, 0, loc)
		fmt.Fprintf(w, "Go launched at %s\n", t.Local())
	}

	event_id := chi.URLParam(r, "event_id")
	fmt.Fprintf(w, "Event ID %s\n", event_id)
}
