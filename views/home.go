package views

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/create_event", http.StatusSeeOther)
}
