package server

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type grpcConfig struct {
	Addr         Listen
}

// grpcServer wraps grpc.Server setup process.
type grpcServer struct {
	server *grpc.Server
	config *grpcConfig
}

func newGrpcServer(c *grpcConfig, servers []ServiceServer) *grpcServer {
	s := grpc.NewServer()
	for _, svr := range servers {
		svr.RegisterWithServer(s)
	}
	return &grpcServer{
		server: s,
		config: c,
	}
}

// Serve implements Server.Server
func (s *grpcServer) Serve() error {
	l, err := s.config.Addr.CreateListener()
	if err != nil {
		return fmt.Errorf("failed to create listener %w", err)
	}

	log.Println("gRPC server is starting ", l.Addr())

	err = s.server.Serve(l)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to serve gRPC server %w", err)
	}
	log.Println("gRPC server ready")

	return nil
}

// Shutdown
func (s *grpcServer) Shutdown(ctx context.Context) {
	s.server.GracefulStop()
}

func createDefaultGrpcConfig() *grpcConfig {
	config := &grpcConfig{
		Addr: Listen{
			Host: "0.0.0.0",
			Port: 443,
		},
	}

	return config
}
