package server_lib

import "google.golang.org/grpc"

// Option configures a gRPC and a gateway server.
type Option func(*Config)

func createConfig(opts []Option) *Config {
	c := createDefaultConfig()
	for _, f := range opts {
		f(c)
	}
	return c
}

// WithGrpcAddrListen
func WithGatewayAddrListen(l Listen) Option {
	return func(c *Config) {
		c.Gateway.Addr = l
	}
}

// WithGrpcAddrListen
func WithGrpcAddrListen(l Listen) Option {
	return func(c *Config) {
		c.Grpc.Addr = l
	}
}

// WithServiceServer
func WithServiceServer(srv ...ServiceServer) Option {
	return func(c *Config) {
		c.ServiceServers = append(c.ServiceServers, srv...)
	}
}

// WithGrpcServerUnaryInterceptors returns an Option that sets unary interceptor(s) for a gRPC server.
func WithGrpcServerUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(c *Config) {
		c.Grpc.PostServerUnaryInterceptors = append(c.Grpc.PostServerUnaryInterceptors, interceptors...)
	}
}
