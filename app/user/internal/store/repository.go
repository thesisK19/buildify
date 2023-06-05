package store

import (
	"context"

	"github.com/thesisK19/buildify/app/user/config"
	"github.com/thesisK19/buildify/app/user/internal/dto"
	"github.com/thesisK19/buildify/app/user/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	// user
	CreateUser(ctx context.Context, params model.CreateUserParams) (*string, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	UpdateUserByUsername(ctx context.Context, username string, params model.UpdateUserParams) error
	// project
	CreateProject(ctx context.Context, username string, projectName string, projectType int) (*dto.Project, error)
	GetListProjects(ctx context.Context, username string) ([]*dto.Project, error)
	GetProject(ctx context.Context, username string, id string) (*dto.Project, error)
	GetProjectBasicInfo(ctx context.Context, username string, id string) (*dto.Project, error)
	UpdateProject(ctx context.Context, project model.Project) error
	DeleteProject(ctx context.Context, username string, id string) error
	// default_project
	GetDefaultProjectByType(ctx context.Context, projectType int) (*model.DefaultProject, error)
	// repository
	Ping() error
	Close() error
}

type repository struct {
	config      *config.Config
	mongoClient *mongo.Client
	db          *mongo.Database
}

func NewRepository(config *config.Config, mongoClient *mongo.Client) Repository {
	db := mongoClient.Database(config.ServiceDB)

	return &repository{
		config:      config,
		mongoClient: mongoClient,
		db:          db,
	}
}

func (r *repository) Ping() error {
	return r.mongoClient.Ping(context.TODO(), nil)
}

func (r *repository) Close() error {
	if err := r.mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	return nil
}
