package views

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"gorm.io/gorm"

	"github.com/eapl-gemugami/meetup-planner/db"
	"github.com/eapl-gemugami/meetup-planner/models"
)

func EventGetAskTimezone(w http.ResponseWriter, r *http.Request) {
	conn, err := db.GetDBConnection()
	if err != nil {
		panic("Failed to connect database")
	}

	public_code := chi.URLParam(r, "public_code")

	var event models.Event
	err = conn.First(&event, "public_code = ?", public_code).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//fmt.Fprintf(w, "Event not found")
		panic("Event not found")
		return
	}

	tmpl_files := []string{
		"templates/base.tmpl.html",
		"templates/ask_timezone.tmpl.html",
	}

	ts, err := template.ParseFS(content, tmpl_files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", &event)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func EventPostAskTimezone(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}

	http.Redirect(w, r,
		r.URL.String()+"/"+url.QueryEscape(r.FormValue("tz")),
		http.StatusSeeOther,
	)
}
