// Test various ways to do HTTP method+path routing in Go

// Each router handles the 11 URLs below:
//
// GET  /										# home
// GET  /contact						# contact
// GET  /api/widgets				# apiGetWidgets
// POST /api/widgets        # apiCreateWidget
// POST /api/widgets/:slug                     	# apiUpdateWidget
// POST /api/widgets/:slug/parts               	# apiCreateWidgetPart
// POST /api/widgets/:slug/parts/:id/update    	# apiUpdateWidgetPart
// POST /api/widgets/:slug/parts/:id/delete    	# apiDeleteWidgetPart
// GET  /:slug							# widget
// GET  /:slug/admin        # widgetAdmin
// POST /:slug/image				# widgetImage

package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 9990

func main() {
	router := Serve // In route.go file

	fmt.Printf("listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
