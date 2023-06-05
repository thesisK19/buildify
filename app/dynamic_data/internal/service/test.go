package service

import (
	"context"

	"github.com/thesisK19/buildify/app/dynamic_data/api"
)

func (s *Service) HealthCheck(ctx context.Context, in *api.EmptyRequest) (*api.EmptyResponse, error) {
	// logger := ctxlogrus.Extract(ctx)
	// logger.Info("Check Health")

	return &api.EmptyResponse{}, nil
}
