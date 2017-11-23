package crawlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router wraps chi.Router.
type Router struct {
	chi.Router
}

func (r *Router) getCrawlers(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) createCrawler(w http.ResponseWriter, req *http.Request) {
}

func (r *Router) getCrawler(w http.ResponseWriter, req *http.Request) {
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
