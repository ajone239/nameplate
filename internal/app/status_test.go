package app_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ajone239/nameplate/internal/app"
	"github.com/ajone239/nameplate/internal/models"
)

type fakeStatusStore struct {
	status *models.Status
	err    error
}

func (f *fakeStatusStore) InitStore() (bool, error) {
	return true, nil
}

func (f *fakeStatusStore) GetStatus() (*models.Status, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.status, nil
}

func (f *fakeStatusStore) SetStatus(s *models.Status) error {
	if f.err != nil {
		return f.err
	}
	f.status = s
	return nil
}

func newTestApp(store *fakeStatusStore) *app.App {
	return &app.App{
		StatusStore: store,
	}
}

func TestGetStatusHandler_OK(t *testing.T) {
	store := &fakeStatusStore{
		status: &models.Status{Status: models.Away},
	}

	a := newTestApp(store)

	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	rec := httptest.NewRecorder()

	a.GetStatusHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("invalid json response")
	}

	if resp["status"] != models.Away.String() {
		t.Fatalf("unexpected status: %v", resp["status"])
	}
}

func TestGetStatusHandler_Error(t *testing.T) {
	store := &fakeStatusStore{
		err: errors.New("db error"),
	}

	a := newTestApp(store)

	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	rec := httptest.NewRecorder()

	a.GetStatusHandler(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rec.Code)
	}
}

func TestPostStatusHandler_OK(t *testing.T) {
	store := &fakeStatusStore{}
	a := newTestApp(store)

	body := []byte(`{"status":"Headdown"}`)
	req := httptest.NewRequest(http.MethodPost, "/status", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	a.PostStatusHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("invalid json response")
	}

	if resp["status"] != "Headdown" {
		t.Fatalf("unexpected status: %v", resp["status"])
	}
}

func TestPostStatusHandler_BadContentType(t *testing.T) {
	store := &fakeStatusStore{}
	a := newTestApp(store)

	req := httptest.NewRequest(http.MethodPost, "/status", nil)
	rec := httptest.NewRecorder()

	a.PostStatusHandler(rec, req)

	if rec.Code != http.StatusUnsupportedMediaType {
		t.Fatalf("expected 415, got %d", rec.Code)
	}
}

func TestPostStatusHandler_InvalidJSON(t *testing.T) {
	store := &fakeStatusStore{}
	a := newTestApp(store)

	req := httptest.NewRequest(
		http.MethodPost,
		"/status",
		bytes.NewBufferString("{invalid json"),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	a.PostStatusHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

func TestPostStatusHandler_StoreError(t *testing.T) {
	store := &fakeStatusStore{
		err: errors.New("write failed"),
	}
	a := newTestApp(store)

	body := []byte(`{"status":"Away"}`)
	req := httptest.NewRequest(http.MethodPost, "/status", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	a.PostStatusHandler(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rec.Code)
	}
}
