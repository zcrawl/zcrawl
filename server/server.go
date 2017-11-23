package server

import (
	"net/http"

	"github.com/zcrawl/zcrawl/server/api"
)

// Server is the main server data structure.
type Server struct {
	settings *Settings
	http.Handler
}

// Settings holds the server settings.
type Settings struct {
	ListenAddr string
}

// New initializes a new server with the given settings.
func New(settings *Settings) Server {
	s := Server{
		settings: settings,
	}
	api := api.New()
	s.Handler = api
	return s
}

// Start starts the server.
func (s *Server) Start() error {
	return http.ListenAndServe(s.settings.ListenAddr, s)
}
