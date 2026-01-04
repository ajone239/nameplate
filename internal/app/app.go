package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ajone239/nameplate/internal/state"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type App struct {
	hits        int
	staticDir   string
	statusStore state.StatusStore
}

func NewApp(initHit int, staticDirPath string) *App {
	db, err := sql.Open("sqlite3", "file:test.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)

	statusStore := state.NewSqliteStatusStore(db)
	statusStore.InitStore()

	return &App{
		hits:        initHit,
		staticDir:   staticDirPath,
		statusStore: statusStore,
	}
}

func (a *App) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Build and register the API layer
	api := a.apiRoutes()
	mux.Handle("/api/", http.StripPrefix("/api", api))

	// Build file server
	fs := http.FileServer(http.Dir(a.staticDir))
	mux.HandleFunc("/", a.spaHandler(fs))

	return mux
}

func (a *App) apiRoutes() *http.ServeMux {
	api := http.NewServeMux()
	api.HandleFunc("GET /data", a.HealthHandler)
	api.HandleFunc("GET /status", a.GetStatusHandler)
	api.HandleFunc("POST /status", a.PostStatusHandler)

	return api
}

func (a *App) spaHandler(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// the route should have caught this already so not found
		if strings.HasPrefix(path, "/api") {
			http.NotFound(w, r)
			return
		}

		// route through spa if theres no file to match
		_, err := os.Stat(filepath.Join(a.staticDir, path))
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(a.staticDir, "index.html"))
			return
		}

		// give a static file otherwise
		fs.ServeHTTP(w, r)
	}
}
