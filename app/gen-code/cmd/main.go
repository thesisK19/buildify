package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"thesis/be/app/gen-code/config"
	"thesis/be/library/server"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var (
	cfg    *config.Config
	logger *logrus.Logger
)

func run(_ []string) error {
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

	// app
	app := cli.NewApp()
	app.Name = "service"
	// app.Usage = "tekit tool"
	// app.Version = Version
	app.Commands = []*cli.Command{
		{
			Name:   "server",
			Usage:  "start grpc/http server",
			Action: serverAction,
		},
		{
			Name:   "config-dump",
			Usage:  "dump config out",
			Action: configDumpAction,
		},
	}

	if app.Run(os.Args) != nil {
		panic(err)
	}
	return nil
}

func configDumpAction(cliCtx *cli.Context) error {
	b, err := json.MarshalIndent(cfg, "", "\t")

	fmt.Println(string(b))
	return err
}

func serverAction(cliCtx *cli.Context) error {
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
		// grpc_recovery.UnaryServerInterceptor(utils.LogTraceStack()), // This is for recover from panic
		// interceptors.SchemaLogInterceptor,
		// interceptors.LogSlowAPI(logger, cfg.SlowResponseCheckingMap, cfg.SlowResponseExceedTime, static),
		// utils.CatchContextErrorInterceptor,
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
