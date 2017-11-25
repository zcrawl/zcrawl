package main

import (
	"github.com/zcrawl/zcrawl/server"
)

const (
	defaultAddr      = ":9999"
	defaultMongoAddr = "localhost:27017"
)

func main() {
	settings := &server.Settings{
		ListenAddr: defaultAddr,
		MongoAddr:  defaultMongoAddr,
	}
	s := server.New(settings)
	err := s.Start()
	if err != nil {
		panic(err)
	}
}
