package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var hits uint = 0

func main() {
	api := http.NewServeMux()
	api.HandleFunc("/api/data", dataHandler)
	http.Handle("/api/", api)

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

	fmt.Println("Listening on: localhost:8080")

	http.ListenAndServe(":8080", nil)
}

type ApiResponse struct {
	Message string `json:"message"`
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	hits += 1

	message := fmt.Sprintf("Hello from json API: %d", hits)
	response := ApiResponse{Message: message}

	json.NewEncoder(w).Encode(response)
}
