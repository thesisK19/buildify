package adapter

import (
	"context"
	"time"

	"github.com/thesisK19/buildify/app/gen_code/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ADAPTER_CONTEXT_TIMEOUT_DEFAULT = 3000 * time.Millisecond

type GenCodeAdapter interface {
	// gen code
	GenReactSourceCode(ctx context.Context, request *api.GenReactSourceCodeRequest) (*api.GenReactSourceCodeResponse, error)
	// test
	HelloWorld(ctx context.Context, request *api.EmptyRequest) (*api.HelloWorldResponse, error)
}

type genCodeAdapter struct {
	grpcAddr string
	client   api.GenCodeServiceClient
}

func NewGenCodeAdapter(genCodeGRPCAddr string) (GenCodeAdapter, error) {
	a := genCodeAdapter{
		grpcAddr: genCodeGRPCAddr,
		client:   nil, // lazy init
	}
	return &a, a.connect()
}

func (a *genCodeAdapter) connect() error {
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

		a.client = api.NewGenCodeServiceClient(conn)
	}
	return nil
}

func (a *genCodeAdapter) GenReactSourceCode(ctx context.Context, request *api.GenReactSourceCodeRequest) (*api.GenReactSourceCodeResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.GenReactSourceCode(ctxWithTimeout, request)
}

func (a *genCodeAdapter) HelloWorld(ctx context.Context, request *api.EmptyRequest) (*api.HelloWorldResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.HelloWorld(ctxWithTimeout, request)
}
