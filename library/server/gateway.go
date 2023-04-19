package server_lib

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"

	"google.golang.org/grpc"
)

// gatewayServer wraps gRPC gateway server setup process.
type gatewayServer struct {
	server *http.Server
	config *gatewayConfig
}

type gatewayConfig struct {
	Addr Listen
}

func createDefaultGatewayConfig() *gatewayConfig {
	//nolint:gomnd
	config := &gatewayConfig{
		Addr: Listen{
			Host: "0.0.0.0",
			Port: 80,
		},
	}

	return config
}

func newGatewayServer(c *gatewayConfig, conn *grpc.ClientConn, servers []ServiceServer) (*gatewayServer, error) {
	// init mux
	gw := runtime.NewServeMux()

	for _, svr := range servers {
		err := svr.RegisterWithHandler(context.Background(), gw, conn)
		if err != nil {
			return nil, fmt.Errorf("failed to register handler. %w", err)
		}
	}

	mux := http.NewServeMux()
	mux.Handle("/", gw)

	svr := &http.Server{
		Addr: c.Addr.String(),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      cors.AllowAll().Handler(mux),
	}

	return &gatewayServer{
		server: svr,
		config: c,
	}, nil
}

// Serve
func (s *gatewayServer) Serve() error {
	log.Println("http server starting at", s.config.Addr.String())

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("Error starting http server, ", err)
		return err
	}

	return nil
}

func (s *gatewayServer) Shutdown(ctx context.Context) {
	err := s.server.Shutdown(ctx)
	if err != nil {
		log.Println("Failed to shutdown grpc-gateway server: ", err)
	}
	log.Println("All http(s) requests finished")
}
