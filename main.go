package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web/build"))
	http.Handle("/", fs)

	api := http.NewServeMux()
	api.HandleFunc("/api/data", dataHandler)
	http.Handle("/api/", api)

	http.ListenAndServe(":8080", nil)
}

type ApiResponse struct {
	Message string `json:"message"`
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	response := ApiResponse{Message: "Hello from json API"}
	json.NewEncoder(w).Encode(response)
}
