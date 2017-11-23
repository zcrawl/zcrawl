package jobs

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router wraps chi.Router.
type Router struct {
	chi.Router
}

func (r *Router) getJobs(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) createJob(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) getJob(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) updateJob(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) deleteJob(w http.ResponseWriter, req *http.Request) {
}

// New initializes a new router.
func New() *Router {
	r := &Router{chi.NewRouter()}
	r.Get("/", r.getJobs)
	r.Post("/", r.createJob)

	r.Route("/{id}", func(subrouter chi.Router) {
		subrouter.Get("/", r.getJob)
		subrouter.Put("/", r.updateJob)
		subrouter.Delete("/", r.deleteJob)
	})
	return r
}
