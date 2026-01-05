package state_test

import (
	"testing"

	"github.com/ajone239/nameplate/internal/models"
	"github.com/ajone239/nameplate/internal/state"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func TestSqliteStatusStore_InitStore(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()

	state := state.NewSqliteStatusStore(db)

	created, err := state.InitStore()
	if err != nil {
		t.Fatalf("InitStore failed: %v", err)
	}

	if !created {
		t.Fatalf("expected state to be initialized")
	}
}

func TestSqliteStatusStore_DoubleInitStore(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()

	state := state.NewSqliteStatusStore(db)

	created, err := state.InitStore()
	if err != nil {
		t.Fatalf("InitStore failed: %v", err)
	}

	if !created {
		t.Fatalf("expected state to be initialized")
	}

	created, err = state.InitStore()
	if err != nil {
		t.Fatalf("InitStore failed: %v", err)
	}

	if !created {
		t.Fatalf("expected state to be initialized")
	}
}

func TestSqliteStatusStore_GetStatus(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()

	state := state.NewSqliteStatusStore(db)

	_, err := state.InitStore()
	if err != nil {
		t.Fatalf("InitStore failed: %v", err)
	}

	status, err := state.GetStatus()
	if err != nil {
		t.Fatalf("GetStatus failed: %v", err)
	}

	if status.Status != models.Away {
		t.Fatalf("expected status %v, got %v", models.Away, status.Status)
	}
}

func TestSqliteStatusStore_SetStatus(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()

	state := state.NewSqliteStatusStore(db)

	_, err := state.InitStore()
	if err != nil {
		t.Fatalf("InitStore failed: %v", err)
	}

	err = state.SetStatus(&models.Status{
		Status: models.HeadDown,
	})
	if err != nil {
		t.Fatalf("SetStatus failed: %v", err)
	}

	status, err := state.GetStatus()
	if err != nil {
		t.Fatalf("GetStatus failed: %v", err)
	}

	if status.Status != models.HeadDown {
		t.Fatalf("expected status %v, got %v", models.HeadDown, status.Status)
	}
}
