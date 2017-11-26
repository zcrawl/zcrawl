package crawlers

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

func (r *Router) getCrawlers(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) createCrawler(w http.ResponseWriter, req *http.Request) {
	rawCrawler, err := ioutil.ReadAll(req.Body)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	c := models.Crawler{}
	err = json.Unmarshal(rawCrawler, &c)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = c.Save()
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	crawlerJSON, _ := json.Marshal(&c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(crawlerJSON)
}

func (r *Router) getCrawler(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	c := models.Crawler{}
	err := c.Get(id)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	crawlerJSON, _ := json.Marshal(&c)
	w.Header().Set("Content-Type", "application/json")
	w.Write(crawlerJSON)
}

func (r *Router) updateCrawler(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) deleteCrawler(w http.ResponseWriter, req *http.Request) {
}

// New initializes a new router.
func New() *Router {
	r := &Router{chi.NewRouter()}
	r.Get("/", r.getCrawlers)
	r.Post("/", r.createCrawler)

	r.Route("/{id}", func(subrouter chi.Router) {
		subrouter.Get("/", r.getCrawler)
		subrouter.Put("/", r.updateCrawler)
		subrouter.Delete("/", r.deleteCrawler)
	})
	return r
}
