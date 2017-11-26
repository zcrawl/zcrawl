package projects

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/zcrawl/zcrawl/server/api/helpers"
	"github.com/zcrawl/zcrawl/server/api/models"
)

// Router wraps chi.Router.
type Router struct {
	chi.Router
}

func (r *Router) getProjects(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) createProject(w http.ResponseWriter, req *http.Request) {

	rawProject, err := ioutil.ReadAll(req.Body)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	p := models.Project{}
	err = json.Unmarshal(rawProject, &p)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = p.Save()
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	projectJSON, _ := json.Marshal(&p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(projectJSON)
}

func (r *Router) getProject(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	p := models.Project{}
	err := p.Get(id)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	projectJSON, _ := json.Marshal(&p)
	w.Header().Set("Content-Type", "application/json")
	w.Write(projectJSON)
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
