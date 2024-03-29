package service

import (
	"context"

	"github.com/thesisK19/buildify/app/gen_code/api"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// RegisterWithServer implementing service server interface
func (s *Service) RegisterWithServer(server *grpc.Server) {
	api.RegisterGenCodeServiceServer(server, s)
}

// RegisterWithHandler implementing service server interface
func (s *Service) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "RegisterWithHandler")

	err := api.RegisterGenCodeServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.WithError(err).Error("failed to register servers")
		return err
	}

	return nil
}
