package views

import (
	"fmt"
	"log"
	"time"
	"errors"
	"strconv"
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

	timeOptions :=
		GetTimeOptions(
			event.TimeStart, event.TimeEnd,
			event.TimeInterval, *targetLoc,
		)

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
	conn, err := db.GetDBConnection()
	if err != nil {
		panic("Failed to connect database")
	}

	public_code := chi.URLParam(r, "public_code")

	var event models.Event
	err = conn.First(&event, "public_code = ?", public_code).Error

	timeOptions :=
		GetTimeOptions(
			event.TimeStart, event.TimeEnd,
			event.TimeInterval, *time.UTC,
		)

	for idx, _ := range timeOptions {
		currentValue := r.FormValue(strconv.Itoa(idx))
		fmt.Printf("%v - %v\n", idx, currentValue)
	}

	// Get the settings of the Event, and get each of the votes

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	//fmt.Fprintf(w, "r.PostFrom = %v\n", r.PostForm)

	for key, value := range r.PostForm {
		fmt.Fprintf(w, "%v - %v\n", key, value)
	}
}