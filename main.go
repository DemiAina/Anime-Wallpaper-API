package main

import (
	"fmt"
	"log"
	"net/http"
	"embed"
    "github.com/DemiAina/Anime-Wallpaper-API/server/render"
)

// go:embed dist/
var embeddedFiles embed.FS

func main() {
var port = ":8000"
http.Handle("/", http.FileServer(http.Dir("dist")))
http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		err := render.RenderTemplate(embeddedFiles, w, r, "about.html")
		if err != nil {
			fmt.Println("ERROR: cannot render", err)
		}
	})

	fmt.Printf("Serving at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

