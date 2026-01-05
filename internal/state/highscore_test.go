package state_test

import (
	"testing"
	"time"

	"github.com/ajone239/nameplate/internal/models"
	"github.com/ajone239/nameplate/internal/state"
)

func seedScores(t *testing.T, store *state.SqliteHighScoreStore) {
	t.Helper()

	scores := []models.HighScore{
		{PlayerName: "Alice", GameName: "Tetris", Score: 1000, Date: time.Now()},
		{PlayerName: "Bob", GameName: "Tetris", Score: 3000, Date: time.Now()},
		{PlayerName: "Carol", GameName: "Pacman", Score: 2000, Date: time.Now()},
	}

	for _, s := range scores {
		_, err := store.AddScore(&s)
		if err != nil {
			t.Fatalf("seed AddScore failed: %v", err)
		}
	}
}

func TestSqliteHighScoreStore_InitStore(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()

	store := state.NewSqliteHighScoreStore(db)

	ok, err := store.InitStore()
	if err != nil {
		t.Fatalf("InitStore failed: %v", err)
	}

	if !ok {
		t.Fatalf("expected InitStore to return true")
	}
}

func TestSqliteHighScoreStore_AddAndGetScore(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()

	store := state.NewSqliteHighScoreStore(db)

	_, err := store.InitStore()
	if err != nil {
		t.Fatalf("InitStore failed: %v", err)
	}

	now := time.Now().UTC()

	score := &models.HighScore{
		PlayerName: "Alice",
		GameName:   "Tetris",
		Score:      4200,
		Date:       now,
	}

	id, err := store.AddScore(score)
	if err != nil {
		t.Fatalf("AddScore failed: %v", err)
	}

	if id == 0 {
		t.Fatalf("expected non-zero id")
	}

	got, err := store.GetScore(int(id))
	if err != nil {
		t.Fatalf("GetScore failed: %v", err)
	}

	if got.PlayerName != score.PlayerName {
		t.Fatalf("player mismatch: got %q", got.PlayerName)
	}

	if got.GameName != score.GameName {
		t.Fatalf("game mismatch: got %q", got.GameName)
	}

	if got.Score != score.Score {
		t.Fatalf("score mismatch: got %d", got.Score)
	}
}

func TestSqliteHighScoreStore_GetScore_NotFound(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()

	store := state.NewSqliteHighScoreStore(db)

	_, err := store.InitStore()
	if err != nil {
		t.Fatalf("InitStore failed: %v", err)
	}

	_, err = store.GetScore(999)
	if err == nil {
		t.Fatalf("expected error for missing row")
	}
}

func TestSqliteHighScoreStore_GetAllScores(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()

	store := state.NewSqliteHighScoreStore(db)

	_, err := store.InitStore()
	if err != nil {
		t.Fatalf("InitStore failed: %v", err)
	}

	seedScores(t, store)

	scores, err := store.GetAllScores()
	if err != nil {
		t.Fatalf("GetAllScores failed: %v", err)
	}

	if len(scores) != 3 {
		t.Fatalf("expected 3 scores, got %d", len(scores))
	}

	// DESC order check
	if scores[0].Score < scores[1].Score {
		t.Fatalf("scores not ordered descending")
	}
}

func TestSqliteHighScoreStore_GetAllGameScores(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()

	store := state.NewSqliteHighScoreStore(db)

	_, err := store.InitStore()
	if err != nil {
		t.Fatalf("InitStore failed: %v", err)
	}

	seedScores(t, store)

	scores, err := store.GetAllGameScores("Tetris")
	if err != nil {
		t.Fatalf("GetAllGameScores failed: %v", err)
	}

	if len(scores) != 2 {
		t.Fatalf("expected 2 scores, got %d", len(scores))
	}

	for _, s := range scores {
		if s.GameName != "Tetris" {
			t.Fatalf("unexpected game: %q", s.GameName)
		}
	}

	// DESC order check
	if scores[0].Score < scores[1].Score {
		t.Fatalf("scores not ordered descending")
	}
}
