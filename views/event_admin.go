package views

import (
	"fmt"
	"log"
	"errors"
	"net/http"
	"html/template"

	"github.com/go-chi/chi"
	"gorm.io/gorm"

	"github.com/eapl-gemugami/meetup-planner/db"
	"github.com/eapl-gemugami/meetup-planner/models"
)

type AdminData struct {
	Event models.Event
	Users []models.EventUser
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	conn, err := db.GetDBConnection()
	if err != nil {
		panic("Failed to connect database")
	}

	admin_code := chi.URLParam(r, "admin_code")

	// Get the event
	var event models.Event
	err = conn.First(&event, "admin_code = ?", admin_code).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprintf(w, "Event not found")
		return
	}

	// Get the Users for that event
	var users []models.EventUser
	conn.Where("event_id = ?", uint(event.ID)).Find(&users)

	adminData := AdminData{
		Event: event,
		Users: users,
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

	err = ts.ExecuteTemplate(w, "base", &adminData)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
