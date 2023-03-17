package service

import (
	"context"
	"fmt"
	"thesis/be/app/gen-code/api"
)

func (s *Service) HelloWorld(ctx context.Context, request *api.HelloWorldRequest) (*api.HelloWorldResponse, error) {
	fmt.Println("helu")

	return &api.HelloWorldResponse{
		Code:    "OK",
		Message: "zuizui",
	}, nil
}
