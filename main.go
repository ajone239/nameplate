package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
	response := ApiResponse{Message: "Hello from json API"}
	json.NewEncoder(w).Encode(response)
}
