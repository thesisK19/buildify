package service

import (
	"context"

	genCodeApi "github.com/thesisK19/buildify/app/gen_code/api"
	"github.com/thesisK19/buildify/app/user/api"
	errors_lib "github.com/thesisK19/buildify/library/errors"
)

func (s *Service) Test(ctx context.Context, in *api.EmptyRequest) (*api.TestResponse, error) {
	resp, err := s.adapters.genCode.HelloWorld(ctx, &genCodeApi.EmptyRequest{})
	if err != nil {
		return nil, errors_lib.ToNotFoundError(err)
	}
	return &api.TestResponse{
		Message: resp.Message + " Success call gencode yeah!",
	}, nil
}

func (s *Service) HealthCheck(ctx context.Context, in *api.EmptyRequest) (*api.EmptyResponse, error) {
	// logger := ctxlogrus.Extract(ctx)
	// logger.Info("Check Health")

	return &api.EmptyResponse{}, nil
}
