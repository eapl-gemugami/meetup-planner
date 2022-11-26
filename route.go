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

	// r = Results (of voting)
	//r.Get("/r/{public_code}", views.EventGetResultsAskTimezone)
	r.Get("/r/{public_code}/{timezone}", views.EventGetResults)

	// a = Admin shortcut
	r.Get("/a/{admin_code}", views.GetAdmin)

	// Serve statics
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/s/*", http.StripPrefix("/s/", fs))

	Serve = r
}

/*
r.Get("/api/widgets", apiGetWidgets)
r.Post("/api/widgets", apiCreateWidget)
r.Post("/api/widgets/{slug}", apiUpdateWidget)
r.Post("/api/widgets/{slug}/parts", apiCreateWidgetPart)
r.Post("/api/widgets/{slug}/parts/{id:[0-9]+}/update", apiUpdateWidgetPart)
r.Post("/api/widgets/{slug}/parts/{id:[0-9]+}/delete", apiDeleteWidgetPart)
r.Get("/{slug}", widgetGet)
r.Get("/{slug}/admin", widgetAdmin)
r.Post("/{slug}/image", widgetImage)

func apiGetWidgets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "apiGetWidgets\n")
}

func apiCreateWidget(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "apiCreateWidget\n")
}

func apiUpdateWidget(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	fmt.Fprintf(w, "apiUpdateWidget %s\n", slug)
}

func apiCreateWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	fmt.Fprintf(w, "apiCreateWidgetPart %s\n", slug)
}

func apiUpdateWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
}

func apiDeleteWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	fmt.Fprintf(w, "apiDeleteWidgetPart %s %d\n", slug, id)
}

func widgetGet(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		fmt.Fprintf(w, "widget %s\n", slug)
}

func widgetAdmin(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	fmt.Fprintf(w, "widgetAdmin %s\n", slug)
}

func widgetImage(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	fmt.Fprintf(w, "widgetImage %s\n", slug)
}
*/
