package state

import (
	"database/sql"
	"encoding/json"
	"errors"
	"os"

	"github.com/ajone239/nameplate/internal/models"
)

type StatusStore interface {
	InitStore() (bool, error)
	GetStatus() (*models.Status, error)
	SetStatus(s *models.Status) error
}

type SqliteStatusStore struct {
	db *sql.DB
}

func NewSqliteStatusStore(db *sql.DB) *SqliteStatusStore {
	return &SqliteStatusStore{
		db: db,
	}
}

func (s *SqliteStatusStore) InitStore() (bool, error) {
	sqlCreate := `CREATE TABLE IF NOT EXISTS status(
		id INTEGER PRIMARY KEY CHECK (id = 0),
        status_num INTEGER NOT NULL
    );`

	_, err := s.db.Exec(sqlCreate)
	if err != nil {
		return false, err
	}

	sqlInsert := `INSERT OR IGNORE INTO status (id, status_num)
			VALUES (0, ?);`

	_, err = s.db.Exec(sqlInsert, models.Away)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *SqliteStatusStore) GetStatus() (*models.Status, error) {
	sql := `SELECT status_num from status where id = 0;`
	row := s.db.QueryRow(sql)

	var status_val int

	err := row.Scan(&status_val)
	if err != nil {
		return nil, err
	}

	status := &models.Status{
		Status: models.StatusState(status_val),
	}

	return status, nil
}

func (s *SqliteStatusStore) SetStatus(status *models.Status) error {
	sql := `UPDATE status
			SET status_num = ?
			where id = 0;`

	result, err := s.db.Exec(sql, status.Status)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("Somehow messed up the database")
	}

	return nil
}

type JsonStatusStore struct {
	path string
}

func (j *JsonStatusStore) InitStore() (bool, error) {
	return true, nil
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
