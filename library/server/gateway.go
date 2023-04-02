package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"

	"google.golang.org/grpc"
)

// gatewayServer wraps gRPC gateway server setup process.
type gatewayServer struct {
	mux    *http.ServeMux
	config *gatewayConfig
}

type gatewayConfig struct {
	Addr              Listen
	ServerMiddlewares []HTTPServerMiddleware
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
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

	return &gatewayServer{
		mux:    mux,
		config: c,
	}, nil
}

// Serve
func (s *gatewayServer) Serve() error {
	log.Println("http server starting at", s.config.Addr.String())
	handler := cors.Default().Handler(s.mux)

	if err := http.ListenAndServe(s.config.Addr.String(), handler); err != nil && err != http.ErrServerClosed {
		log.Println("Error starting http server, ", err)
		return err
	}

	return nil
}

func (s *gatewayServer) Shutdown(ctx context.Context) {
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	// err := s.mux.Shutdown(ctx)
	log.Println("All http(s) requests finished")
	// if err != nil {
	// 	log.Println("failed to shutdown grpc-gateway server: ", err)
	// }
}
