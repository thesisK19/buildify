package service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/user/api"
	"github.com/thesisK19/buildify/app/user/internal/model"
	server_lib "github.com/thesisK19/buildify/library/server"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) CreateProject(ctx context.Context, in *api.CreateProjectRequest) (*api.CreateProjectResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateProject")

	username := server_lib.GetUsernameFromContext(ctx)

	newProject, err := s.repository.CreateProject(ctx, username, in.Name, int(in.GetType().Number()))
	if err != nil {
		logger.WithError(err).Error("failed to repo.CreateProject")
		return nil, err
	}

	return &api.CreateProjectResponse{
		Id:             newProject.Id,
		Name:           newProject.Name,
		CompressString: newProject.CompressString,
		CreatedAt:      newProject.CreatedAt,
		UpdatedAt:      newProject.UpdatedAt,
	}, nil
}

func (s *Service) GetListProjects(ctx context.Context, in *api.EmptyRequest) (*api.GetListProjectsResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetListProjects")

	username := server_lib.GetUsernameFromContext(ctx)
	var resp api.GetListProjectsResponse

	listProjects, err := s.repository.GetListProjects(ctx, username)
	if err != nil {
		logger.WithError(err).Error("failed to repo.GetListProjects")
		return nil, err
	}

	for _, project := range listProjects {
		resp.Projects = append(resp.Projects, &api.Project{
			Id:             project.Id,
			Name:           project.Name,
			CompressString: project.CompressString,
			CreatedAt:      project.CreatedAt,
			UpdatedAt:      project.UpdatedAt,
		})
	}

	return &resp, nil
}

func (s *Service) GetProject(ctx context.Context, in *api.GetProjectRequest) (*api.GetProjectResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetProject")

	username := server_lib.GetUsernameFromContext(ctx)

	project, err := s.repository.GetProject(ctx, username, in.Id)
	if err != nil {
		logger.WithError(err).Error("failed to repo.GetProject")
		return nil, err
	}

	return &api.GetProjectResponse{
		Id:             project.Id,
		Name:           project.Name,
		CompressString: project.CompressString,
		CreatedAt:      project.CreatedAt,
		UpdatedAt:      project.UpdatedAt,
	}, nil
}

func (s *Service) UpdateProject(ctx context.Context, in *api.UpdateProjectRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UpdateProject")

	username := server_lib.GetUsernameFromContext(ctx)

	objectId, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		logger.WithError(err).Error("failed to convert ObjectIDFromHex")
		return nil, err
	}

	err = s.repository.UpdateProject(ctx, model.Project{
		Id:             objectId,
		Name:           in.Name,
		Username:       username,
		CompressString: in.CompressString,
	})
	if err != nil {
		logger.WithError(err).Error("failed to repo.UpdateProject")
		return nil, err
	}

	return &api.EmptyResponse{}, nil
}

func (s *Service) DeleteProject(ctx context.Context, in *api.DeleteProjectRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "DeleteProject")

	username := server_lib.GetUsernameFromContext(ctx)

	err := s.repository.DeleteProject(ctx, username, in.Id)
	if err != nil {
		logger.WithError(err).Error("failed to repo.DeleteProject")
		return nil, err
	}

	return &api.EmptyResponse{}, nil
}
