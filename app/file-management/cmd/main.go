package main

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/thesisK19/buildify/app/file-management/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

var (
	cfg    *config.Config
	logger *logrus.Logger
)

func run() error {
	var err error

	// load config
	cfg, err = config.Load()
	if err != nil {
		return err
	}

	// init logging
	logger, err = cfg.Log.Build()
	if err != nil {
		return err
	}

	// start server
	server, err := newService(cfg, logger)
	if err != nil {
		return err
	}
	server.Serve()

	return nil
}
