package service

import (
	"context"

	"github.com/thesisK19/buildify/app/user/api"
	"github.com/thesisK19/buildify/app/user/internal/model"
	"google.golang.org/genproto/googleapis/rpc/code"
)

func (s *Service) CreateUser(ctx context.Context, in *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	createParams := model.CreateUserParams{
		Name:     in.Name,
		Username: in.Username,
		Password: in.Password,
	}
	id, err := s.repository.CreateUser(ctx, createParams)
	if err != nil {
		return &api.CreateUserResponse{
			Code:    code.Code_OK,
			Message: err.Error(),
		}, nil
	}

	return &api.CreateUserResponse{
		Code:    code.Code_OK,
		Message: code.Code_OK.String(),
		Id:      *id,
	}, nil
}

func (s *Service) GetUser(ctx context.Context, in *api.GetUserRequest) (*api.GetUserResponse, error) {
	if in.Id == "" && in.Username == "" {
		return &api.GetUserResponse{
			Code:    code.Code_INTERNAL,
			Message: "id or username is required",
		}, nil
	}

	var (
		user *model.User
		err  error
	)

	if in.Id != "" {
		user, err = s.repository.GetUserByID(ctx, in.Id)
	} else {
		user, err = s.repository.GetUserByUsername(ctx, in.Username)
	}

	if err != nil {
		return &api.GetUserResponse{
			Code:    code.Code_INTERNAL,
			Message: err.Error(),
		}, nil
	}

	return &api.GetUserResponse{
		Code:    code.Code_OK,
		Message: code.Code_OK.String(),
		User: &api.User{
			Username: user.Username,
			Name:     user.Name,
		},
	}, nil
}

func (s *Service) UpdateUser(ctx context.Context, in *api.UpdateUserRequest) (*api.UpdateUserResponse, error) {
	if in.Id == "" && in.Username == "" {
		return &api.UpdateUserResponse{
			Code:    code.Code_INVALID_ARGUMENT,
			Message: "id or username is required",
		}, nil
	}

	var (
		err error
	)
	updateParams := model.UpdateUserParams{
		Name:     in.Name,
		Username: in.Username,
		Password: in.Password,
	}

	if in.Id != "" {
		err = s.repository.UpdateUserByID(ctx, in.Id, updateParams)
	} else {
		err = s.repository.UpdateUserByUsername(ctx, in.Username, updateParams)
	}

	if err != nil {
		return &api.UpdateUserResponse{
			Code:    code.Code_INTERNAL,
			Message: err.Error(),
		}, nil
	}

	return &api.UpdateUserResponse{
		Code:    code.Code_OK,
		Message: code.Code_OK.String(),
	}, nil
}
