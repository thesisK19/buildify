package service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	dynamicDataApi "github.com/thesisK19/buildify/app/dynamic_data/api"
	"github.com/thesisK19/buildify/app/user/api"
	"github.com/thesisK19/buildify/app/user/internal/model"
	server_lib "github.com/thesisK19/buildify/library/server"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *Service) CreateProject(ctx context.Context, in *api.CreateProjectRequest) (*api.CreateProjectResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateProject")

	username := server_lib.GetUsernameFromContext(ctx)

	newProject, err := s.repository.CreateProject(ctx, username, in.Name, int(in.GetType().Number()))
	if err != nil {
		logger.WithError(err).Error("failed to repo.CreateProject")
		return nil, err
	}

	err = s.createExampleDatabase(ctx, newProject.Id)
	if err != nil {
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

func (s *Service) createExampleDatabase(ctx context.Context, projectId string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "createExampleDatabase")
	collResp, err := s.adapters.dynamicData.CreateCollection(ctx, &dynamicDataApi.CreateCollectionRequest{
		ProjectId: projectId,
		Name:      "Product",
		DataKeys:  []string{"name", "description", "price", "image"},
		DataTypes: []int32{1, 1, 1, 1},
	})
	if err != nil {
		logger.WithError(err).Error("failed to createExampleDatabase")
		return err
	}
	collId := collResp.GetId()
	_, err = s.adapters.dynamicData.CreateDocument(ctx, &dynamicDataApi.CreateDocumentRequest{
		CollectionId: collId,
		Data: map[string]*structpb.Value{
			"name":        createStructValue("iPhone 12 64GB"),
			"description": createStructValue("Chip Apple A14 Bionic RAM: 4 GB Dung lượng: 64 GB Camera sau: 2 camera 12 MP Camera trước: 12 MP Pin 2815 mAh, Sạc 20 W"),
			"price":       createStructValue("15.490.000 đ"),
			"image":       createStructValue("https://storage.googleapis.com/dynamic-data-bucket/example/iphone-12-tim.jpg"),
		},
	})
	if err != nil {
		logger.WithError(err).Error("failed to createExampleDatabase")
		return err
	}

	_, err = s.adapters.dynamicData.CreateDocument(ctx, &dynamicDataApi.CreateDocumentRequest{
		CollectionId: collId,
		Data: map[string]*structpb.Value{
			"name":        createStructValue("iPhone 14"),
			"description": createStructValue("Chip Apple A15 Bionic RAM: 6 GB Dung lượng: 128 GB Camera sau: 2 camera 12 MP Camera trước: 12 MP Pin 3279 mAh, Sạc 20 W"),
			"price":       createStructValue("19.390.000 đ"),
			"image":       createStructValue("https://storage.googleapis.com/dynamic-data-bucket/example/iPhone-14-do.jpg"),
		},
	})
	if err != nil {
		logger.WithError(err).Error("failed to createExampleDatabase")
		return err
	}
	_, err = s.adapters.dynamicData.CreateDocument(ctx, &dynamicDataApi.CreateDocumentRequest{
		CollectionId: collId,
		Data: map[string]*structpb.Value{
			"name":        createStructValue("iPhone 14 Pro"),
			"description": createStructValue("Chip Apple A16 Bionic RAM: 6 GB Dung lượng: 128 GB Camera sau: Chính 48 MP & Phụ 12 MP, 12 MP Camera trước: 12 MP Pin 3200 mAh, Sạc 20 W"),
			"price":       createStructValue("24.990.000 đ"),
			"image":       createStructValue("https://storage.googleapis.com/dynamic-data-bucket/example/iphone-14-pro-vang.jpg"),
		},
	})
	if err != nil {
		logger.WithError(err).Error("failed to createExampleDatabase")
		return err
	}
	return nil
}

func createStructValue(input string) *structpb.Value {
	// Create a *structpb.Value object for the string value
	value := &structpb.Value{
		Kind: &structpb.Value_StringValue{
			StringValue: input,
		},
	}

	return value
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

func (s *Service) InternalGetProjectBasicInfo(ctx context.Context, in *api.InternalGetProjectBasicInfoRequest) (*api.InternalGetProjectBasicInfoResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "InternalGetProjectBasicInfo")

	username := server_lib.GetUsernameFromContext(ctx)

	project, err := s.repository.GetProjectBasicInfo(ctx, username, in.Id)
	if err != nil {
		logger.WithError(err).Error("failed to repo.GetProjectBasicInfo")
		return nil, err
	}

	return &api.InternalGetProjectBasicInfoResponse{
		Id:   project.Id,
		Name: project.Name,
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
