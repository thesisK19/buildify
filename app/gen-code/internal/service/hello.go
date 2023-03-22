package service

import (
	"buildify/app/gen-code/api"
	"context"
	"fmt"
)

func (s *Service) HelloWorld(ctx context.Context, request *api.HelloWorldRequest) (*api.HelloWorldResponse, error) {
	fmt.Println("helu")

	return &api.HelloWorldResponse{
		Code:    "OK",
		Message: "zuizui",
	}, nil
}
