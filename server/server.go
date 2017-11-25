package server

import (
	"net/http"

	"github.com/zcrawl/zcrawl/server/api"
	"github.com/zcrawl/zcrawl/server/api/models"
)

// Server is the main server data structure.
type Server struct {
	settings *Settings
	http.Handler
}

// Settings holds the server settings.
type Settings struct {
	ListenAddr string
	MongoAddr  string
}

// New initializes a new server with the given settings.
func New(settings *Settings) Server {
	s := Server{
		settings: settings,
	}
	// Initialize the API handlers:
	api := api.New()
	s.Handler = api
	return s
}

// Start starts the server.
func (s *Server) Start() error {
	// Setup the Mongo connection:
	err := models.DialMongo(s.settings.MongoAddr)
	if err != nil {
		return err
	}
	return http.ListenAndServe(s.settings.ListenAddr, s)
}
