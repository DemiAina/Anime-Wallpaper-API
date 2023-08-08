package main

import (
	"fmt"
	"log"
	"net/http"
    "strings"
)

func main() {
	port := ":8000"
	fs := http.FileServer(http.Dir("dist"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request path ends with ".js"
		if strings.HasSuffix(r.URL.Path, ".js") {
			// Set the MIME type to "application/javascript" for JS files
			w.Header().Set("Content-Type", "application/javascript")
		}
		fs.ServeHTTP(w, r)
	})

	fmt.Printf("Serving at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
