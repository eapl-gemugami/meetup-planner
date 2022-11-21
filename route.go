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
	r.Get("/contact", views.Contact)
	r.Get("/time", views.Time)
	r.Get("/e/{event_id}", views.GetDataRange) // Event
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
	*/

	// Serve statics
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/s/*", http.StripPrefix("/s/", fs))

	Serve = r
}

/*
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
