package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ajone239/nameplate/internal/models"
)

type statusResponse struct {
	Status  string `json:"status"`
	TimeSet string `json:"time"`
}

type statusRequest struct {
	Status string `json:"status"`
}

func (app *App) GetStatusHandler(w http.ResponseWriter, _ *http.Request) {
	status, err := app.statusStore.GetStatus()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	statusResponse := &statusResponse{
		Status:  status.Status.String(),
		TimeSet: status.TimeSet.UTC().String(),
	}

	json.NewEncoder(w).Encode(statusResponse)
}

func (app *App) PostStatusHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	var input statusRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	status := models.Status{
		Status:  models.FromString(input.Status),
		TimeSet: time.Now(),
	}

	err = app.statusStore.SetStatus(&status)
	if err != nil {
		log.Println(err)
		http.Error(w, "Status Save Failed", http.StatusInternalServerError)
		return
	}

	statusResponse := &statusResponse{
		Status:  status.Status.String(),
		TimeSet: status.TimeSet.UTC().String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(statusResponse)
}
