package views

import (
	"fmt"
	"time"
	"errors"
	"net/http"

	"gorm.io/gorm"
	"github.com/go-chi/chi"

	"github.com/eapl-gemugami/meetup-planner/db"
	"github.com/eapl-gemugami/meetup-planner/models"
)

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func GetDataRange(w http.ResponseWriter, r *http.Request) {
	conn, err := db.GetDBConnection()
	if err != nil {
		panic("Failed to connect database")
	}

	public_code := chi.URLParam(r, "public_code")

	var event models.Event
	err = conn.First(&event, "public_code = ?", public_code).Error // find product with code D42

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprintf(w, "Event not found")
		return
	}

	//fmt.Fprintf(w, "Event ID: %s\n\n", event)

	timeStart := time.Unix(event.TimeStart, 0)
	timeEnd := time.Unix(event.TimeEnd, 0)

	// TODO: Check that the difference is less than a week
	// https://gosamples.dev/difference-between-dates/
	difference := timeEnd.Sub(timeStart)
	fmt.Printf("Weeks: %d\n", int64(difference.Hours() / 24 / 7))

	newLocationName := "Australia/Brisbane"
	newLoc, _ := time.LoadLocation(newLocationName)

	fmt.Fprintf(w, "Times in %s\n", newLocationName)

	currentDate := time.UnixMilli(timeStart.UnixMilli())
	if timeEnd.After(timeStart) {
		incrementMins := event.TimeInterval

		currentIdx := 0
		currentTimeIsBeforetimeEnd := true
		for currentTimeIsBeforetimeEnd {
			// https://pkg.go.dev/fmt#hdr-Printing
			//fmt.Fprintf(w, "%v - %s\n", currentIdx, currentDate.UTC())
			fmt.Fprintf(w, "%v - %s\n",
				currentIdx,
				currentDate.In(newLoc).Format("02/Jan/2006 - Monday @ 15:04 / 3:04pm"),
			)

			currentDate = currentDate.Add(time.Minute * time.Duration(incrementMins))
			currentTimeIsBeforetimeEnd = currentDate.Before(timeEnd)
			currentIdx += 1
		}

		// And add the time end
		fmt.Fprintf(w, "%v - %s\n",
			currentIdx,
			timeEnd.In(newLoc).Format("02/Jan/2006 - Monday @ 15:04 / 3:04pm"),
		)
	} else {
		fmt.Println("ERROR - timeEnd MUST be after timeStart")
	}
}

/*
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

		// // https://gosamples.dev/difference-between-dates/
		// difference := date2.Sub(date1)

		// fmt.Printf("Weeks: %d\n", int64(difference.Hours()/24/7))
		// fmt.Printf("Days: %d\n", int64(difference.Hours()/24))
		// fmt.Printf("Hours: %.f\n", difference.Hours())
		// fmt.Printf("Minutes: %.f\n", difference.Minutes())

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
*/