package state

import (
	"database/sql"
	"time"

	"github.com/ajone239/nameplate/internal/models"
)

type HighScoreStore interface {
	InitStore() (bool, error)
	AddScore(s *models.HighScore) (int64, error)
	GetScore(id int) (*models.HighScore, error)
	GetAllScores() ([]models.HighScore, error)
	GetAllGameScores(gameName string) ([]models.HighScore, error)
}

type SqliteHighScoreStore struct {
	db *sql.DB
}

func NewSqliteHighScoreStore(db *sql.DB) *SqliteHighScoreStore {
	return &SqliteHighScoreStore{
		db: db,
	}
}

func (h *SqliteHighScoreStore) InitStore() (bool, error) {
	sqlCreate := `CREATE TABLE IF NOT EXISTS highscores(
		id INTEGER PRIMARY KEY,
		player_name string not null,
		game_name string not null,
		score integer,
		date string not null
    );`

	_, err := h.db.Exec(sqlCreate)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (h *SqliteHighScoreStore) AddScore(s *models.HighScore) (int64, error) {
	sql := `INSERT INTO highscores (player_name, game_name, score, date)
			VALUES (?, ?, ?, ?);`

	result, err := h.db.Exec(
		sql,
		s.PlayerName,
		s.GameName,
		s.Score,
		s.Date.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()

}

func (h *SqliteHighScoreStore) GetScore(id int) (*models.HighScore, error) {
	sql := `SELECT id, player_name, game_name, score, date from highscores where id = ?;`
	row := h.db.QueryRow(sql, id)

	return highScoreFromRow(row)
}

func highScoreFromRow(row *sql.Row) (*models.HighScore, error) {
	var ts string
	hs := &models.HighScore{}

	err := row.Scan(&hs.Id, &hs.PlayerName, &hs.GameName, &hs.Score, &ts)
	if err != nil {
		return nil, err
	}

	hs.Date, err = time.Parse(time.RFC3339, ts)
	if err != nil {
		return nil, err
	}

	return hs, nil
}

func (h *SqliteHighScoreStore) GetAllScores() ([]models.HighScore, error) {
	sql := `SELECT id, player_name, game_name, score, date
			from highscores
			order by score desc;`

	rows, err := h.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var highscores []models.HighScore

	for rows.Next() {
		var ts string
		hs := models.HighScore{}

		err := rows.Scan(&hs.Id, &hs.PlayerName, &hs.GameName, &hs.Score, &ts)
		if err != nil {
			return nil, err
		}

		hs.Date, err = time.Parse(time.RFC3339, ts)
		if err != nil {
			return nil, err
		}

		highscores = append(highscores, hs)
	}

	return highscores, nil
}

func (h *SqliteHighScoreStore) GetAllGameScores(gameName string) ([]models.HighScore, error) {
	sql := `SELECT id, player_name, game_name, score, date
			from highscores
			where game_name = ?
			order by score desc;`

	rows, err := h.db.Query(sql, gameName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var highscores []models.HighScore

	for rows.Next() {
		var ts string
		hs := models.HighScore{}

		err := rows.Scan(&hs.Id, &hs.PlayerName, &hs.GameName, &hs.Score, &ts)
		if err != nil {
			return nil, err
		}

		hs.Date, err = time.Parse(time.RFC3339, ts)
		if err != nil {
			return nil, err
		}

		highscores = append(highscores, hs)
	}

	return highscores, nil
}
