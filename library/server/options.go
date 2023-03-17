package server

import (
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Option configures a gRPC and a gateway server.
type Option func(*Config)

func createConfig(opts []Option) *Config {
	c := createDefaultConfig()
	for _, f := range opts {
		f(c)
	}
	return c
}

// WithGatewayAddr
func WithGatewayAddr(host string, port int) Option {
	return func(c *Config) {
		c.Gateway.Addr = Listen{
			Host: host,
			Port: port,
		}
	}
}

// WithGrpcAddrListen
func WithGatewayAddrListen(l Listen) Option {
	return func(c *Config) {
		c.Gateway.Addr = l
	}
}

// WithGatewayServerMiddlewares returns an Option that sets middleware(s) for http.Server to a gateway server.
func WithGatewayServerMiddlewares(middlewares ...HTTPServerMiddleware) Option {
	return func(c *Config) {
		c.Gateway.ServerMiddlewares = append(c.Gateway.ServerMiddlewares, middlewares...)
	}
}

// WithPassedHeader returns an Option that sets configurations about passed headers for a gateway server.
func WithPassedHeader(decider PassedHeaderDeciderFunc) Option {
	return WithGatewayServerMiddlewares(createPassingHeaderMiddleware(decider))
}

///-------------------------- GRPC options below--------------------------

// WithGrpcAddr
func WithGrpcAddr(host string, port int) Option {
	return func(c *Config) {
		c.Grpc.Addr = Listen{
			Host: host,
			Port: port,
		}
	}
}

// WithGrpcAddrListen
func WithGrpcAddrListen(l Listen) Option {
	return func(c *Config) {
		c.Grpc.Addr = l
	}
}

// WithGrpcServerUnaryInterceptors returns an Option that sets unary interceptor(s) for a gRPC server.
func WithGrpcServerUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(c *Config) {
		c.Grpc.PostServerUnaryInterceptors = append(c.Grpc.PostServerUnaryInterceptors, interceptors...)
	}
}

// WithGrpcServerStreamInterceptors returns an Option that sets stream interceptor(s) for a gRPC server.
func WithGrpcServerStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(c *Config) {
		c.Grpc.PostServerStreamInterceptors = append(c.Grpc.PostServerStreamInterceptors, interceptors...)
	}
}

// WithDefaultLogger returns an Option that sets default grpclogger.LoggerV2 object.
func WithDefaultLogger() Option {
	return func(c *Config) {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))
	}
}

// WithServiceServer
func WithServiceServer(srv ...ServiceServer) Option {
	return func(c *Config) {
		c.ServiceServers = append(c.ServiceServers, srv...)
	}
}
