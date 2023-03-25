package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thesisK19/buildify/app/user/api"
	"github.com/thesisK19/buildify/app/user/internal/model"
)

func (s *Service) CreateUser(ctx context.Context, in *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	newUser := model.User{
		Username: in.Username,
		Password: in.Password,
	}
	_, err := s.repository.CreateUser(ctx, &newUser)
	if err != nil {
		return &api.CreateUserResponse{
			Code:    "err",
			Message: err.Error(),
		}, nil
	}

	byteArr, _ := json.Marshal(newUser)

	return &api.CreateUserResponse{
		Code:    "OK",
		Message: string(byteArr),
	}, nil
}
func (s *Service) GetUser(ctx context.Context, in *api.GetUserRequest) (*api.GetUserResponse, error) {
	fmt.Print("helllllllllll")
	return &api.GetUserResponse{
		Code:    "3",
		Message: "à haha",
	}, nil
}
