package service

import (
	"context"

	"github.com/thesisK19/buildify/app/dynamic_data/api"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// RegisterWithServer implementing service server interface
func (s *Service) RegisterWithServer(server *grpc.Server) {
	api.RegisterDynamicDataServiceServer(server, s)
}

// RegisterWithHandler implementing service server interface
func (s *Service) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "RegisterWithHandler")

	err := api.RegisterDynamicDataServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.WithError(err).Error("Failed to register servers")
		return err
	}

	return err
}
