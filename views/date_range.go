package views

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func GetDataRange(w http.ResponseWriter, r *http.Request) {
	event_id := chi.URLParam(r, "event_id")
	fmt.Fprintf(w, "Event ID: %s\n\n", event_id)

	// Define a starting and end date and get a list of hours in between
	loc, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		fmt.Printf("Timezone invalid")
	} else {
		// Example - Friday Nov 18th, 2022 - 9pm - México City
		timeStart := time.Date(2022, 11, 18, 21, 0, 0, 0, loc)

		// Saturday Nov 19th, 2022 - 9pm - México City
		timeEnd := time.Date(2022, 11, 19, 21, 0, 0, 0, loc)

		/*
			// https://gosamples.dev/difference-between-dates/
			difference := date2.Sub(date1)

			fmt.Printf("Weeks: %d\n", int64(difference.Hours()/24/7))
			fmt.Printf("Days: %d\n", int64(difference.Hours()/24))
			fmt.Printf("Hours: %.f\n", difference.Hours())
			fmt.Printf("Minutes: %.f\n", difference.Minutes())
		*/

		currentDate := time.UnixMilli(timeStart.UnixMilli())
		if timeEnd.After(timeStart) {
			incrementMins := 60

			currentIdx := 0
			currentTimeIsBeforetimeEnd := true
			for currentTimeIsBeforetimeEnd {
				// https://pkg.go.dev/fmt#hdr-Printing
				fmt.Fprintf(w, "%v - %s\n", currentIdx, currentDate.Local())

				currentDate = currentDate.Add(time.Minute * time.Duration(incrementMins))
				currentTimeIsBeforetimeEnd = currentDate.Before(timeEnd)
				currentIdx += 1
			}
		} else {
			fmt.Println("ERROR - timeEnd 2 MUST be after timeStart")
		}
	}
}
