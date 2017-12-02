package users

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

func (r *Router) getUsers(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) createUser(w http.ResponseWriter, req *http.Request) {
	rawUser, err := ioutil.ReadAll(req.Body)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	u := models.User{}
	err = json.Unmarshal(rawUser, &u)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = u.Save()
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userJSON, _ := json.Marshal(&u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(userJSON)
}

func (r *Router) getUser(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	u := models.User{}
	err := u.Get(id)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userJSON, _ := json.Marshal(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
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
