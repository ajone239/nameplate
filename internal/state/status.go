package state

import (
	"encoding/json"
	"os"

	"github.com/ajone239/nameplate/internal/models"
)

type StatusStore interface {
	GetStatus() (*models.Status, error)
	SetStatus(s *models.Status) error
}

func NewStatusStore() StatusStore {
	return &JsonStatusStore{
		path: "status.json",
	}
}

type JsonStatusStore struct {
	path string
}

func (j *JsonStatusStore) load() (*models.Status, error) {
	data, err := os.ReadFile(j.path)
	if err != nil {
		if os.IsNotExist(err) {
			return &models.Status{}, nil
		}
		return nil, err
	}

	var status models.Status
	json.Unmarshal(data, &status)

	return &status, nil
}

func (j *JsonStatusStore) save(s *models.Status) error {
	data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(j.path, data, 0644)
}

func (j *JsonStatusStore) GetStatus() (*models.Status, error) {
	return j.load()
}

func (j *JsonStatusStore) SetStatus(s *models.Status) error {
	return j.save(s)
}
