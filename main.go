package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var hits uint = 0

func main() {
	fs := http.FileServer(http.Dir("./web/build"))
	http.Handle("/", fs)

	api := http.NewServeMux()
	api.HandleFunc("/api/data", dataHandler)
	http.Handle("/api/", api)

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
