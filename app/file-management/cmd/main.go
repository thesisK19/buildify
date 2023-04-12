package main

import (
	"log"

	"github.com/thesisK19/buildify/app/file-management/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

var (
	cfg *config.Config
)

func run() error {
	var err error

	// load config
	cfg, err = config.Load()
	if err != nil {
		return err
	}

	// start server
	server, err := newService(cfg)
	if err != nil {
		return err
	}
	server.Serve()

	return nil
}
