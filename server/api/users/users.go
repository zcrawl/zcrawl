package users

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router wraps chi.Router.
type Router struct {
	chi.Router
}

func (r *Router) getUsers(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) createUser(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) getUser(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) updateUser(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) deleteUser(w http.ResponseWriter, req *http.Request) {
}

// New initializes a new router.
func New() *Router {
	r := &Router{chi.NewRouter()}
	r.Get("/", r.getUsers)
	r.Post("/", r.createUser)

	r.Route("/{id}", func(subrouter chi.Router) {
		subrouter.Get("/", r.getUser)
		subrouter.Put("/", r.updateUser)
		subrouter.Delete("/", r.deleteUser)
	})
	return r
}
