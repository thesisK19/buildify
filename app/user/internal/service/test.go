package service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/user/api"
	"github.com/thesisK19/buildify/library/errors"
)

func (s *Service) Test(ctx context.Context, in *api.TestRequest) (*api.TestResponse, error) {
	// logger := ctxlogrus.Extract(ctx).WithField("func", "Test")

	resp, err := s.adapters.genCode.HelloWorld(ctx)
	if err != nil {
		return nil, errors.ToNotFoundError(err)
	}
	return &api.TestResponse{
		Message: resp.Message + "\nAnd this is user new version kk",
	}, nil
}

func (s *Service) HealthCheck(ctx context.Context, in *api.HealthCheckRequest) (*api.HealthCheckResponse, error) {
	logger := ctxlogrus.Extract(ctx)
	logger.Info("Check Health")

	return &api.HealthCheckResponse{}, nil
}
