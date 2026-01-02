package app

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type App struct {
	hits      int
	staticDir string
}

func NewApp(initHit int, staticDirPath string) *App {
	return &App{
		hits:      initHit,
		staticDir: staticDirPath,
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
	api.HandleFunc("/data", a.HealthHandler)

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
