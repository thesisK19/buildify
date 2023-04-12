package service

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/user/api"
	"github.com/thesisK19/buildify/app/user/internal/model"
	"github.com/thesisK19/buildify/library/errors"
)

func (s *Service) CreateUser(ctx context.Context, in *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateUser")

	createParams := model.CreateUserParams{
		Name:     in.Name,
		Username: in.Username,
		Password: in.Password,
	}
	id, err := s.repository.CreateUser(ctx, createParams)
	if err != nil {
		logger.WithError(err).Error("Failed to repo.CreateUser")
		return nil, err
	}

	return &api.CreateUserResponse{
		Id: *id,
	}, nil
}

func (s *Service) GetUser(ctx context.Context, in *api.GetUserRequest) (*api.GetUserResponse, error) {
	if in.Id == "" && in.Username == "" {
		return nil, errors.ToInvalidArgumentError(fmt.Errorf("id or username is required"))
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
		return nil, err
	}

	return &api.GetUserResponse{
		User: &api.User{
			Username: user.Username,
			Name:     user.Name,
		},
	}, nil
}

func (s *Service) UpdateUser(ctx context.Context, in *api.UpdateUserRequest) (*api.UpdateUserResponse, error) {
	if in.Id == "" && in.Username == "" {
		return nil, errors.ToInvalidArgumentError(fmt.Errorf("id or username is required"))
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
		return nil, err
	}

	return &api.UpdateUserResponse{}, nil
}
