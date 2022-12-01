package views

import (
	"errors"
	"fmt"
	"log"
	"time"

	"html/template"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"gorm.io/gorm"

	"github.com/eapl-gemugami/meetup-planner/db"
	"github.com/eapl-gemugami/meetup-planner/models"
)

type ResultsData struct {
	Event     models.Event
	OptionSum []OptionSum
}

type OptionSum struct {
	Text string
	Sum  int
}

func EventGetResults(w http.ResponseWriter, r *http.Request) {
	conn, err := db.GetDBConnection()
	if err != nil {
		panic("Failed to connect database")
	}

	public_code := chi.URLParam(r, "public_code")
	timezone, _ := url.QueryUnescape(chi.URLParam(r, "timezone"))

	// TODO: Check the timezone is valid
	targetLoc, errTz := time.LoadLocation(timezone)
	if errTz != nil {
		fmt.Fprintf(w, "Invalid timezone. Blame your dev!")
		return
	}

	// Get the event
	var event models.Event
	err = conn.First(&event, "public_code = ?", public_code).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprintf(w, "Event not found")
		return
	}

	timeOptions :=
		GetTimeOptions(
			event.TimeStart, event.TimeEnd,
			event.TimeInterval, *targetLoc,
		)

	// Get all the Votes for that event
	var votes []models.EventVote
	conn.Where("event_id = ?", uint(event.ID)).Find(&votes)

	optionSum := map[int]int{}

	for _, vote := range votes {
		if _, exists := optionSum[vote.TimeOption]; !exists {
			optionSum[vote.TimeOption] = 0
		}

		//fmt.Printf("Option: %v - Availability: %d\n", vote.TimeOption, vote.TimeAvailability)
		optionSum[vote.TimeOption] += vote.TimeAvailability
	}

	var optionsSum []OptionSum

	for idx, optionText := range timeOptions {
		sum := 0
		if _, exists := optionSum[idx]; exists {
			sum = optionSum[idx]
		}

		optionsSum = append(optionsSum, OptionSum{
			Text: optionText,
			Sum:  sum,
		})
	}

	fmt.Println(optionsSum)

	resultsData := ResultsData{
		Event:     event,
		OptionSum: optionsSum,
	}

	tmpl_files := []string{
		"templates/base.tmpl.html",
		"templates/results.tmpl.html",
	}

	ts, err := template.ParseFS(content, tmpl_files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", &resultsData)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
