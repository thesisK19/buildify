package main

import (
	"log"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"github.com/thesisK19/buildify/app/gen_code/config"
	server_lib "github.com/thesisK19/buildify/library/server"
	"google.golang.org/grpc"
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
		log.Fatal("failed to load config", err)
		return err
	}

	// init logging
	logger, err = cfg.Log.Build()
	if err != nil {
		log.Fatal("failed to build logger", err)
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
	service, err := newService(cfg)
	if err != nil {
		logger.WithError(err).Error("failed to init servers")
		return err
	}

	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	logrusEntry := logrus.NewEntry(logger)

	s, err := server_lib.New(
		server_lib.WithGrpcAddrListen(cfg.Server.GRPC),
		server_lib.WithGatewayAddrListen(cfg.Server.HTTP),
		server_lib.WithGrpcServerUnaryInterceptors(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			grpc.UnaryServerInterceptor(server_lib.AuthInterceptor),
		),
		server_lib.WithServiceServer(service),
	)
	if err != nil {
		logger.WithError(err).Error("failed to get new servers")
		return err
	}

	if err := s.Serve(); err != nil {
		logger.WithError(err).Error("failed to start servers")
		return err
	}
	return nil
}
