package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server is the framework instance.
type Server struct {
	grpcServer    *grpcServer
	gatewayServer *gatewayServer
	config        *Config
}

// New creates a server intstance.
func New(opts ...Option) (*Server, error) {
	c := createConfig(opts)

	log.Println("Create grpc server")
	grpcServer := newGrpcServer(c.Grpc, c.ServiceServers)
	reflection.Register(grpcServer.server)
	// if err != nil {
	// 	return nil, fmt.Errorf("Faild to create grpc server. %w", err)
	// }

	conn, err := grpc.Dial(c.Grpc.Addr.String(), grpc.WithInsecure(),
		//nolint:gomnd
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*50)),
		grpc.WithChainUnaryInterceptor(
			otelgrpc.UnaryClientInterceptor(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("fail to dial gRPC server. %w", err)
	}

	log.Println("Create gateway server")
	gatewayServer, err := newGatewayServer(c.Gateway, conn, c.ServiceServers)
	if err != nil {
		return nil, fmt.Errorf("fail to create gateway server. %w", err)
	}

	return &Server{
		grpcServer:    grpcServer,
		gatewayServer: gatewayServer,
		config:        c,
	}, nil
}

// Serve starts gRPC and Gateway servers.
func (s *Server) Serve() error {
	stop := make(chan os.Signal, 1)
	errch := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.gatewayServer.Serve(); err != nil {
			log.Println("Error starting http server, ", err)
			errch <- err
		}
	}()

	go func() {
		if err := s.grpcServer.Serve(); err != nil {
			log.Println("Error starting gRPC server, ", err)
			errch <- err
		}
	}()

	// shutdown
	for {
		select {
		case <-stop:
			log.Println("Shutting down server")
			//nolint:gomnd
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			for _, ss := range s.config.ServiceServers {
				ss.Close(ctx)
			}

			s.gatewayServer.Shutdown(ctx)
			s.grpcServer.Shutdown(ctx)
			return nil
		case err := <-errch:
			return err
		}
	}
}
