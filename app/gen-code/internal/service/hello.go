package service

import (
	"context"

	"github.com/thesisK19/buildify/app/gen-code/api"
)

func (s *Service) HelloWorld(ctx context.Context, request *api.HelloWorldRequest) (*api.HelloWorldResponse, error) {

	return &api.HelloWorldResponse{
		Message: "hellu <3",
	}, nil
}
