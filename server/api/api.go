package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/matiasinsaurralde/zcrawl-platform/server/api/crawlers"
	"github.com/matiasinsaurralde/zcrawl-platform/server/api/jobs"
	"github.com/matiasinsaurralde/zcrawl-platform/server/api/projects"
	"github.com/matiasinsaurralde/zcrawl-platform/server/api/users"
	"github.com/matiasinsaurralde/zcrawl-platform/server/api/workers"
	"github.com/sirupsen/logrus"
)

// API wraps the main router.
type API struct {
	chi.Router
}

func (api *API) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}

// loadRoutes mounts all subrouters.
func (api *API) loadRoutes() {
	crawlersRouter := crawlers.New()
	api.Mount("/crawlers", crawlersRouter)

	jobsRouter := jobs.New()
	api.Mount("/jobs", jobsRouter)

	projectsRouter := projects.New()
	api.Mount("/projects", projectsRouter)

	usersRouter := users.New()
	api.Mount("/users", usersRouter)

	workersRouter := workers.New()
	api.Mount("/workers", workersRouter)

	api.Get("/ping", api.ping)
}

// New is used to initialize a new router.
func New() http.Handler {
	// Setup logging:
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}

	// Initialize the main router and attach the base middleware:
	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware.RequestID)
	mainRouter.Use(NewStructuredLogger(logger))
	mainRouter.Use(middleware.Recoverer)

	r := API{mainRouter}
	r.loadRoutes()
	return r
}
