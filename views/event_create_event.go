package views

import (
	"fmt"
	"log"
	"time"
	"strconv"
	"strings"
	"net/http"
	"math/rand"
	"html/template"

	"github.com/eapl-gemugami/meetup-planner/db"
	"github.com/eapl-gemugami/meetup-planner/models"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
// From Microsoft Product Activation
//var letters = []rune("2346789BCDFGHJKMPQRTVWXY")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateEventGet(w http.ResponseWriter, r *http.Request) {
	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"templates/base.tmpl.html",
		"templates/create_event.tmpl.html",
	}

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	ts, err := template.ParseFS(content, files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func CreateEventPost(w http.ResponseWriter, r *http.Request) {
	conn, err := db.GetDBConnection()
	if err != nil {
		panic("Failed to connect database")
	}

	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		// Redirect to Error
		fmt.Printf("ParseForm() err: %v", err)
		return
	}

	//fmt.Printf("r.PostFrom = %v\n", r.PostForm)

	loc, err := time.LoadLocation(r.FormValue("tz"))
	if err != nil {
		panic("Invalid timezone")
	}

	// Try to parse the date
	start_date_split := strings.Split(r.FormValue("start_date"), "-")
	if len(start_date_split) != 3 {
		panic("Incorrect date format")
	}

	start_year, _ := strconv.Atoi(start_date_split[0])
	start_month, _ := strconv.Atoi(start_date_split[1])
	start_day, _ := strconv.Atoi(start_date_split[2])
	start_hour, _ := strconv.Atoi(r.FormValue("start_time"))

	timeStart := time.Date(start_year, time.Month(start_month), start_day, start_hour, 0, 0, 0, loc)
	//fmt.Printf("Start time: %v\n", timeStart)

	end_date_split := strings.Split(r.FormValue("end_date"), "-")
	if len(end_date_split) != 3 {
		panic("Incorrect date format")
	}

	end_year, _ := strconv.Atoi(end_date_split[0])
	end_month, _ := strconv.Atoi(end_date_split[1])
	end_day, _ := strconv.Atoi(end_date_split[2])
	end_hour, _ := strconv.Atoi(r.FormValue("end_time"))

	timeEnd := time.Date(end_year, time.Month(end_month), end_day, end_hour, 0, 0, 0, loc)
	//fmt.Printf("End time: %v\n", timeEnd)

	time_interval, _ := strconv.Atoi(r.FormValue("interval"))

	// Create random codes
	rand.Seed(time.Now().UnixNano())
	publicCode := randSeq(15)
	adminCode := randSeq(35)

	// Create event
	conn.Create(&models.Event{
		Name: strings.TrimSpace(r.FormValue("event_name")),

		PublicCode: publicCode,
		AdminCode:  adminCode,

		TimeStart:    timeStart.Unix(),
		TimeEnd:      timeEnd.Unix(),
		TimeInterval: time_interval,
	})

	// TODO: Add a field of num of options

	//fmt.Printf("Created event: %v, %v\n", publicCode, adminCode)

	// https://pkg.go.dev/net/http#pkg-constants
	http.Redirect(w, r, "/a/" + adminCode, http.StatusSeeOther)
}
