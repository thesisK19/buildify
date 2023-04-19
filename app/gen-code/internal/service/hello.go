package service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/gen-code/api"
)

func (s *Service) HelloWorld(ctx context.Context, request *api.EmptyRequest) (*api.HelloWorldResponse, error) {

	return &api.HelloWorldResponse{
		Message: "hellu from gencode <3",
	}, nil
}

func (s *Service) HealthCheck(ctx context.Context, in *api.EmptyRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx)
	logger.Info("Check Health")

	return &api.EmptyResponse{}, nil
}
