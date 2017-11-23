package workers

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router wraps chi.Router.
type Router struct {
	chi.Router
}

func (r *Router) getWorkers(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) createWorker(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) getWorker(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) updateWorker(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) deleteWorker(w http.ResponseWriter, req *http.Request) {
}

// New initializes a new router.
func New() *Router {
	r := &Router{chi.NewRouter()}
	r.Get("/", r.getWorkers)
	r.Post("/", r.createWorker)

	r.Route("/{id}", func(subrouter chi.Router) {
		subrouter.Get("/", r.getWorker)
		subrouter.Put("/", r.updateWorker)
		subrouter.Delete("/", r.deleteWorker)
	})
	return r
}
