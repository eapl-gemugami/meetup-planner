package views

import (
	"fmt"
	"net/http"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "contact\n")
}
