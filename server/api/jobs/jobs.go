package jobs

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

func (r *Router) getJobs(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) createJob(w http.ResponseWriter, req *http.Request) {
	rawJob, err := ioutil.ReadAll(req.Body)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	j := models.Job{}
	err = json.Unmarshal(rawJob, &j)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = j.Save()
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	jobJSON, _ := json.Marshal(&j)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jobJSON)
}

func (r *Router) getJob(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	j := models.Job{}
	err := j.Get(id)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	jobJSON, _ := json.Marshal(&j)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jobJSON)
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
