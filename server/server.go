package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/matiasinsaurralde/zcrawl-platform/server/api"
)

// Server is the main server data structure.
type Server struct {
	settings *Settings
	*chi.Mux
}

// Settings holds the server settings.
type Settings struct {
	ListenAddr string
}

// New initializes a new server with the given settings.
func New(settings *Settings) Server {
	s := Server{
		settings: settings,
		Mux:      chi.NewRouter(),
	}
	s.mountRoutes()
	return s
}

func (s *Server) mountRoutes() {
	api := api.New()
	s.Mux.Mount("/", api)
}

// Start starts the server.
func (s *Server) Start() error {
	return http.ListenAndServe(s.settings.ListenAddr, s.Mux)
}
