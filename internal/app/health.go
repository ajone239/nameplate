package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type healthResponse struct {
	Message string `json:"message"`
}

func (app *App) HealthHandler(w http.ResponseWriter, _ *http.Request) {
	app.hits += 1

	message := fmt.Sprintf("Hello from json API: %d", app.hits)
	response := healthResponse{Message: message}

	json.NewEncoder(w).Encode(response)
}
