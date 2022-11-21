package views

import (
	"fmt"
	"net/http"
	"time"
)

func Time(w http.ResponseWriter, r *http.Request) {
	// About Monotonic clock, m=+/-
	// https://stackoverflow.com/a/52809985/13173382
	fmt.Fprintf(w, "Current time: %s\n\n", time.Now())
}
