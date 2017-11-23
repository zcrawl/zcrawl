package main

import (
	"github.com/zcrawl/zcrawl/server"
)

const (
	defaultAddr = ":9999"
)

func main() {
	settings := &server.Settings{
		ListenAddr: defaultAddr,
	}
	s := server.New(settings)
	err := s.Start()
	if err != nil {
		panic(err)
	}
}
