package main

import (
	"log"

	"github.com/thesisK19/buildify/app/gen-code/config"
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
		// server.WithGrpcServerUnaryInterceptors(
		// grpc_zap.UnaryServerInterceptor(logger),                     // This is for assign logger to server -> Have to stand before every interceptor need logging
		// interceptors.InjectTraceIDLogger,                            // This is for inject trace id to logger -> Have to stand after logger inject and before another processes
		// grpc_recovery.UnaryServerInterceptor(util.LogTraceStack()), // This is for recover from panic
		// interceptors.SchemaLogInterceptor,
		// interceptors.LogSlowAPI(logger, cfg.SlowResponseCheckingMap, cfg.SlowResponseExceedTime, static),
		// util.CatchContextErrorInterceptor,
		// ),
		server.WithGatewayAddrListen(cfg.Server.HTTP),
		server.WithServiceServer(service),
		// server.WithErrorLogger(cfg.ListErrorNotLog),
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
