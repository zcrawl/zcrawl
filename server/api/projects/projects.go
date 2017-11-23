package projects

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router wraps chi.Router.
type Router struct {
	chi.Router
}

func (r *Router) getProjects(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) createProject(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) getProject(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) updateProject(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) deleteProject(w http.ResponseWriter, req *http.Request) {
}

// New initializes a new router.
func New() *Router {
	r := &Router{chi.NewRouter()}
	r.Get("/", r.getProjects)
	r.Post("/", r.createProject)

	r.Route("/{id}", func(subrouter chi.Router) {
		subrouter.Get("/", r.getProject)
		subrouter.Put("/", r.updateProject)
		subrouter.Delete("/", r.deleteProject)
	})
	return r
}
