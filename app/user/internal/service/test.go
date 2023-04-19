package service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/user/api"
	errors_lib "github.com/thesisK19/buildify/library/errors"
)

func (s *Service) Test(ctx context.Context, in *api.EmptyRequest) (*api.TestResponse, error) {
	// logger := ctxlogrus.Extract(ctx).WithField("func", "Test")

	resp, err := s.adapters.genCode.HelloWorld(ctx)
	if err != nil {
		return nil, errors_lib.ToNotFoundError(err)
	}
	return &api.TestResponse{
		Message: resp.Message + " Success call gencode yeah!",
	}, nil
}

func (s *Service) HealthCheck(ctx context.Context, in *api.EmptyRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx)
	logger.Info("Check Health")

	return &api.EmptyResponse{}, nil
}
