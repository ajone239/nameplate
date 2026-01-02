package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	app := &App{hits: 0, incs: 3}

	api := http.NewServeMux()
	api.HandleFunc("/data", app.dataHandler)

	http.Handle("/api/", http.StripPrefix("/api", api))

	staticDir := "./web/build"
	fs := http.FileServer(http.Dir(staticDir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if strings.HasPrefix(path, "/api") {
			http.NotFound(w, r)
			return
		}

		_, err := os.Stat(filepath.Join(staticDir, path))
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
			return
		}

		fs.ServeHTTP(w, r)
	})

	fmt.Println("Listening on: localhost:80")

	http.ListenAndServe(":80", nil)
}

type App struct {
	hits int
	incs int
}

type ApiResponse struct {
	Message string `json:"message"`
}

func (app *App) dataHandler(w http.ResponseWriter, r *http.Request) {
	app.hits += app.incs

	message := fmt.Sprintf("Hello from json API: %d", app.hits)
	response := ApiResponse{Message: message}

	json.NewEncoder(w).Encode(response)
}
