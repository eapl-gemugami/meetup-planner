package views

import (
	"fmt"
	"log"
	"errors"
	"net/http"
	"html/template"

	"gorm.io/gorm"
	"github.com/go-chi/chi"

	"github.com/eapl-gemugami/meetup-planner/db"
	"github.com/eapl-gemugami/meetup-planner/models"
)

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	conn, err := db.GetDBConnection()
	if err != nil {
		panic("Failed to connect database")
	}

	admin_code := chi.URLParam(r, "admin_code")

	var event models.Event
	err = conn.First(&event, "admin_code = ?", admin_code).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprintf(w, "Event not found")
		return
	}

	tmpl_files := []string{
		"templates/base.tmpl.html",
		"templates/admin.tmpl.html",
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
