package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
    "time"
    "github.com/DemiAina/Anime-Wallpaper-API/db"
)

func stringManpluationName(file string) string {
     extensions := []string{".jpg", ".png"}
	for _, ext := range extensions {
		file = strings.TrimSuffix(file, ext)
	}
	return file
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to retrieve image", http.StatusBadRequest)
		fmt.Printf("Failed to retrieve image: %v\n", err)
		return
	}
	defer file.Close()

	newFilePath := "./dist/images_uploaded/" + header.Filename

	err = os.MkdirAll(filepath.Dir(newFilePath), 0755)
	if err != nil {
		http.Error(w, "Failed to create directories", http.StatusInternalServerError)
		fmt.Printf("Failed to create directories: %v\n", err)
		return
	}

	newFile, err := os.Create(newFilePath)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		fmt.Printf("Failed to save file: %v\n", err)
		return
	}

	name := stringManpluationName(header.Filename)
	fmt.Print(name)


    hash := db.GenerateHash(header.Filename)
    exsits := db.HashExists(hash)
    if exsits{
        log.Print("File has been uploaded before")
        return
    }
    
    // This seems very specific however this was in the documentation
    //	Jan 2 15:04:05 2006 MST
    //https://pkg.go.dev/time#Time.Format
    // Its pretty intresting this design decsion
    currentTime := time.Now()
	uploadedAt := currentTime.Format("2006-01-02 15:04:05")
    err = db.AddImage(header.Filename,newFilePath,hash,uploadedAt)
    if err != nil{
        log.Println("Error inserting image to the database:",err)
    }
    
    fmt.Printf("\n %s is hash \n" , hash)

	fmt.Printf("Uploaded file name: %s\n", header.Filename)
	fmt.Fprintln(w, "Image uploaded and saved successfully")
	fmt.Println("Image uploaded and saved successfully")
}

func main() {

    err := db.InitConnection()
    if err != nil {
        fmt.Println("Error initializing database connection:", err)
        return
    }
    defer db.CloseConnection()

	port := ":8000"
	fs := http.FileServer(http.Dir("dist"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self' http://localhost:8000;")
		if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript")
		}
		fs.ServeHTTP(w, r)
	})

	http.HandleFunc("/upload", uploadHandler)

	fmt.Printf("Serving at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))




}
