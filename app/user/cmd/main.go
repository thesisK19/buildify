package main

import (
	"log"

	"github.com/thesisK19/buildify/app/user/config"
	"github.com/thesisK19/buildify/library/server"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
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

	// run server
	err = runServer()
	if err != nil {
		return err
	}

	return nil
}

func runServer() error {
	service, err := newService(cfg, logger)
	if err != nil {
		logger.Error("Error init servers", zap.Error(err))
		return err
	}

	s, err := server.New(
		server.WithGrpcAddrListen(cfg.Server.GRPC),
		server.WithGatewayAddrListen(cfg.Server.HTTP),
		server.WithServiceServer(service),
	)
	if err != nil {
		logger.Error("Error new servers", zap.Error(err))
		return err
	}

	if err := s.Serve(); err != nil {
		logger.Error("Error start servers", zap.Error(err))
		return err
	}
	return nil
}
