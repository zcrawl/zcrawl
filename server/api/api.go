package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// API wraps the API handlers.
type API struct {
	chi.Router
}

func (a *API) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}

func (a *API) loadRoutes() {
	a.Get("/ping", a.ping)
}

// New is used to initialize a new router.
func New() http.Handler {
	r := API{chi.NewRouter()}
	r.loadRoutes()
	return r
}
