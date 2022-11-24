package views

import (
	"fmt"
	"log"
	"time"
	"errors"
	"net/url"
	"net/http"
	"html/template"

	"gorm.io/gorm"
	"github.com/go-chi/chi"

	"github.com/eapl-gemugami/meetup-planner/db"
	"github.com/eapl-gemugami/meetup-planner/models"
)

type TemplateData struct {
	Event models.Event
	TimeOptions []string
	Timezone string
	TimezoneEscaped string
}

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

func EventGetDataRange(w http.ResponseWriter, r *http.Request) {
	conn, err := db.GetDBConnection()
	if err != nil {
		fmt.Fprintf(w, "Failed to connect to the DB. Blame your dev!")
		return
	}

	public_code := chi.URLParam(r, "public_code")
	timezone, _ := url.QueryUnescape(chi.URLParam(r, "timezone"))

	// TODO: Check the timezone is valid
	targetLoc, errTz := time.LoadLocation(timezone)
	if errTz != nil {
		fmt.Fprintf(w, "Invalid timezone. Blame your dev!")
		return
	}

	var event models.Event
	err = conn.First(&event, "public_code = ?", public_code).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprintf(w, "Invalid timezone. Blame your dev!")
		return
	}

	timeStart := time.Unix(event.TimeStart, 0)
	timeEnd := time.Unix(event.TimeEnd, 0)

	// TODO: Check that the difference is less than a week
	// https://gosamples.dev/difference-between-dates/
	difference := timeEnd.Sub(timeStart)
	diff_weeks := int64(difference.Hours() / 24 / 7)
	if diff_weeks >= 1 {
		fmt.Fprintf(w, "Data range is too big - Only a week is allowed")
		return
	}

	if timeEnd.Before(timeStart) {
		fmt.Fprintf(w, "End date happens in the past! WTF!")
		return
	}

	currentDate := time.UnixMilli(timeStart.UnixMilli())
	incrementMins := event.TimeInterval

	var timeOptions []string

	currentIdx := 0
	currentTimeIsBeforetimeEnd := true
	for currentTimeIsBeforetimeEnd {
		// https://pkg.go.dev/fmt#hdr-Printing
		//fmt.Fprintf(w, "%v - %s\n", currentIdx, currentDate.UTC())
		/*
		fmt.Fprintf(w, "%v - %s\n",
			currentIdx,
			currentDate.In(targetLoc).Format("02/Jan/2006 - Monday @ 15:04 / 3:04pm"),
		)
		*/

		timeOptions = append(timeOptions, currentDate.In(targetLoc).Format("02/Jan/2006 - Monday @ 15:04 / 3:04pm"))

		currentDate = currentDate.Add(time.Minute * time.Duration(incrementMins))
		currentTimeIsBeforetimeEnd = currentDate.Before(timeEnd)
		currentIdx += 1
	}

	timeOptions = append(timeOptions, currentDate.In(targetLoc).Format("02/Jan/2006 - Monday @ 15:04 / 3:04pm"))

	templateData := TemplateData {
		Event: event,
		TimeOptions: timeOptions,
		Timezone: timezone,
		TimezoneEscaped: chi.URLParam(r, "timezone"),
	}

	tmpl_files := []string{
		"templates/base.tmpl.html",
		"templates/time_poll.tmpl.html",
	}

	ts, err := template.ParseFS(content, tmpl_files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error parsing Templates", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", &templateData)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error parsing Templates", 500)
	}
}

func EventPostDataRange(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	//fmt.Fprintf(w, "r.PostFrom = %v\n", r.PostForm)

	for key, value := range r.PostForm {
		fmt.Fprintf(w, "%v - %v\n", key, value)
	}
}