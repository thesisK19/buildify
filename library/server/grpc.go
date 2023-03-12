package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type grpcConfig struct {
	Addr                               Listen
	PreServerUnaryInterceptors         []grpc.UnaryServerInterceptor
	PreServerStreamInterceptors        []grpc.StreamServerInterceptor
	PostServerUnaryInterceptors        []grpc.UnaryServerInterceptor
	PostServerStreamInterceptors       []grpc.StreamServerInterceptor
	DefaultUnaryInterceptorsLog        grpc.UnaryServerInterceptor
	DefaultStreamInterceptorsLog       grpc.StreamServerInterceptor
	DefaultUnaryInterceptorsValidator  grpc.UnaryServerInterceptor
	DefaultStreamInterceptorsValidator grpc.StreamServerInterceptor
	ServerOption                       []grpc.ServerOption
	MaxConcurrentStreams               uint32
}

func (c *grpcConfig) ServerOptions() []grpc.ServerOption {
	enforcement := keepalive.EnforcementPolicy{
		MinTime:             SERVER_OPTION_ENFORCEMENT_MIN_TIME * time.Second,
		PermitWithoutStream: true,
	}
	return append(
		[]grpc.ServerOption{
			grpc.MaxConcurrentStreams(c.MaxConcurrentStreams),
			grpc.KeepaliveEnforcementPolicy(enforcement), // here
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle:     SERVER_OPTION_MAX_CONNECTION_IDLE * time.Second,
				MaxConnectionAge:      SERVER_OPTION_MAX_CONNECTION_AGE * time.Second,
				MaxConnectionAgeGrace: SERVER_OPTION_MAX_CONNECTION_AGE_GRACE * time.Second,
			}),
		},
		c.ServerOption...,
	)
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
	// lgr := createLog(zapcore.InfoLevel)
	// grpc_prometheus.EnableHandlingTimeHistogram()
	config := &grpcConfig{
		Addr: Listen{
			Host: "0.0.0.0",
			Port: 10443,
		},
		// PreServerUnaryInterceptors: []grpc.UnaryServerInterceptor{
		// 	otelgrpc.UnaryServerInterceptor(),
		// 	grpc_prometheus.UnaryServerInterceptor,
		// 	grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		// 	grpc_ctx.UnaryServerInterceptor(),
		// },
		// PreServerStreamInterceptors: []grpc.StreamServerInterceptor{
		// 	otelgrpc.StreamServerInterceptor(),
		// 	grpc_prometheus.StreamServerInterceptor,
		// 	grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		// 	grpc_ctx.StreamServerInterceptor(),
		// },
		// DefaultUnaryInterceptorsLog:        grpc_logr.UnaryServerInterceptor(lgr, grpc_logr.WithDecider(DefaultShouldLog)),
		// DefaultStreamInterceptorsLog:       grpc_logr.StreamServerInterceptor(lgr, grpc_logr.WithDecider(DefaultShouldLog)),
		// DefaultUnaryInterceptorsValidator:  grpc_validator.UnaryServerInterceptor(),
		// DefaultStreamInterceptorsValidator: grpc_validator.StreamServerInterceptor(),
		// MaxConcurrentStreams:               1000,
	}

	return config
}
