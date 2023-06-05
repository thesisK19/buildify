package adapter

import (
	"context"
	"time"

	"github.com/thesisK19/buildify/app/user/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ADAPTER_CONTEXT_TIMEOUT_DEFAULT = 3000 * time.Millisecond

type UserAdapter interface {
	// user
	GetUser(ctx context.Context, in *api.EmptyRequest) (*api.GetUserResponse, error)
	UpdateUser(ctx context.Context, in *api.UpdateUserRequest) (*api.EmptyResponse, error)
	// project
	CreateProject(ctx context.Context, in *api.CreateProjectRequest) (*api.CreateProjectResponse, error)
	GetListProjects(ctx context.Context, in *api.EmptyRequest) (*api.GetListProjectsResponse, error)
	GetProject(ctx context.Context, in *api.GetProjectRequest) (*api.GetProjectResponse, error)
	InternalGetProjectBasicInfo(ctx context.Context, in *api.InternalGetProjectBasicInfoRequest) (*api.InternalGetProjectBasicInfoResponse, error)
	UpdateProject(ctx context.Context, in *api.UpdateProjectRequest) (*api.EmptyResponse, error)
	DeleteProject(ctx context.Context, in *api.DeleteProjectRequest) (*api.EmptyResponse, error)
	// test
	Test(ctx context.Context, in *api.EmptyRequest) (*api.TestResponse, error)
}

type userAdapter struct {
	grpcAddr string
	client   api.UserServiceClient
}

func NewUserAdapter(userGRPCAddr string) (UserAdapter, error) {
	a := userAdapter{
		grpcAddr: userGRPCAddr,
		client:   nil, // lazy init
	}
	return &a, a.connect()
}

func (a *userAdapter) connect() error {
	if a.client == nil {
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
		defer cancel()

		conn, err := grpc.DialContext(
			ctxWithTimeout,
			a.grpcAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return err
		}

		a.client = api.NewUserServiceClient(conn)
	}
	return nil
}

func (a *userAdapter) GetUser(ctx context.Context, in *api.EmptyRequest) (*api.GetUserResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.GetUser(ctxWithTimeout, in)
}

func (a *userAdapter) UpdateUser(ctx context.Context, in *api.UpdateUserRequest) (*api.EmptyResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.UpdateUser(ctxWithTimeout, in)
}

func (a *userAdapter) CreateProject(ctx context.Context, in *api.CreateProjectRequest) (*api.CreateProjectResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.CreateProject(ctxWithTimeout, in)
}

func (a *userAdapter) GetListProjects(ctx context.Context, in *api.EmptyRequest) (*api.GetListProjectsResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.GetListProjects(ctxWithTimeout, in)
}

func (a *userAdapter) GetProject(ctx context.Context, in *api.GetProjectRequest) (*api.GetProjectResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.GetProject(ctxWithTimeout, in)
}

func (a *userAdapter) InternalGetProjectBasicInfo(ctx context.Context, in *api.InternalGetProjectBasicInfoRequest) (*api.InternalGetProjectBasicInfoResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.InternalGetProjectBasicInfo(ctxWithTimeout, in)
}

func (a *userAdapter) UpdateProject(ctx context.Context, in *api.UpdateProjectRequest) (*api.EmptyResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.UpdateProject(ctxWithTimeout, in)
}

func (a *userAdapter) DeleteProject(ctx context.Context, in *api.DeleteProjectRequest) (*api.EmptyResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.DeleteProject(ctxWithTimeout, in)
}

func (a *userAdapter) Test(ctx context.Context, in *api.EmptyRequest) (*api.TestResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.Test(ctxWithTimeout, in)
}
