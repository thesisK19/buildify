package server_lib

import (
	"context"
	"fmt"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

type grpcConfig struct {
	Addr Listen

	PreServerUnaryInterceptors  []grpc.UnaryServerInterceptor
	PostServerUnaryInterceptors []grpc.UnaryServerInterceptor
	MaxConcurrentStreams        uint32

	ServerOption []grpc.ServerOption
}

// grpcServer wraps grpc.Server setup process.
type grpcServer struct {
	server *grpc.Server
	config *grpcConfig
}

func newGrpcServer(c *grpcConfig, servers []ServiceServer) *grpcServer {
	s := grpc.NewServer(c.ServerOptions()...)
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
		return fmt.Errorf("Failed to create listener %w", err)
	}

	log.Println("gRPC server is starting ", l.Addr())

	err = s.server.Serve(l)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Failed to serve gRPC server %w", err)
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
		PreServerUnaryInterceptors: []grpc.UnaryServerInterceptor{
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		},
		MaxConcurrentStreams: 100,
	}

	return config
}

func (c *grpcConfig) buildServerUnaryInterceptors() []grpc.UnaryServerInterceptor {
	unaryInterceptors := c.PreServerUnaryInterceptors
	unaryInterceptors = append(unaryInterceptors, c.PostServerUnaryInterceptors...)
	return unaryInterceptors
}

func (c *grpcConfig) ServerOptions() []grpc.ServerOption {
	return append(
		[]grpc.ServerOption{
			grpc_middleware.WithUnaryServerChain(c.buildServerUnaryInterceptors()...),
			grpc.MaxConcurrentStreams(c.MaxConcurrentStreams),
		},
		c.ServerOption...,
	)
}
