package service

import (
	"context"

	"github.com/thesisK19/buildify/app/user/api"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// RegisterWithServer implementing service server interface
func (s *Service) RegisterWithServer(server *grpc.Server) {
	api.RegisterUserServiceServer(server, s)
}

// RegisterWithHandler implementing service server interface
func (s *Service) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	entry := ctxlogrus.Extract(ctx)

	err := api.RegisterUserServiceHandler(ctx, mux, conn)
	if err != nil {
		entry.Error("Error register servers")
	}

	return err
}
