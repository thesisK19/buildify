package service

import (
	"context"
	"fmt"

	"github.com/thesisK19/buildify/app/gen-code/api"
)

func (s *Service) HelloWorld(ctx context.Context, request *api.HelloWorldRequest) (*api.HelloWorldResponse, error) {
	fmt.Println("helu")
	// url, err := util.UploadFile()
	// if err == nil {
	// 	return &api.HelloWorldResponse{
	// 		Code:    "OK",
	// 		Message: *url,
	// 	}, nil
	// }

	return &api.HelloWorldResponse{
		Code:    "OK",
		Message: "GENCODE: zuizui",
	}, nil
}
