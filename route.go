package main

import (
	"net/http"

	"github.com/eapl-gemugami/meetup-planner/views"
	"github.com/go-chi/chi"
)

var Serve http.Handler

func init() {
	r := chi.NewRouter()

	r.Get("/", views.Home)

	// Debug Routes
	r.Get("/debug/time", views.Time)

	// Admin
	r.Get("/admin/migrate", views.Migrate)

	// Event routes
	r.Get("/create_event", views.CreateEventGet)
	r.Post("/create_event", views.CreateEventPost)

	// e = Event shortcut
	r.Get("/e/{public_code}", views.EventGetAskTimezone)
	r.Post("/e/{public_code}", views.EventPostAskTimezone)
	r.Get("/e/{public_code}/{timezone}", views.EventGetDataRange)
	r.Post("/e/{public_code}/{timezone}", views.EventPostDataRange)

	// r = Results (of survey)
	r.Get("/r/{public_code}", views.EventGetResultsAskTimezone)
	r.Post("/r/{public_code}", views.EventPostResultsAskTimezone)
	r.Get("/r/{public_code}/{timezone}", views.EventGetResults)

	// a = Admin shortcut
	r.Get("/a/{admin_code}", views.GetAdmin)

	// Serve statics
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/s/*", http.StripPrefix("/s/", fs))

	Serve = r
}
