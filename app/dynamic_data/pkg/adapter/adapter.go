package adapter

import (
	"context"
	"time"

	"github.com/thesisK19/buildify/app/dynamic_data/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ADAPTER_CONTEXT_TIMEOUT_DEFAULT = 3000 * time.Millisecond

type DynamicDataAdapter interface {
	// collection
	CreateCollection(ctx context.Context, in *api.CreateCollectionRequest) (*api.CreateCollectionResponse, error)
	GetListCollections(ctx context.Context, in *api.GetListCollectionsRequest) (*api.GetListCollectionsResponse, error)
	GetCollection(ctx context.Context, in *api.GetCollectionRequest) (*api.GetCollectionResponse, error)
	UpdateCollection(ctx context.Context, in *api.UpdateCollectionRequest) (*api.EmptyResponse, error)
	DeleteCollection(ctx context.Context, in *api.DeleteCollectionRequest) (*api.EmptyResponse, error)
	// document
	CreateDocument(ctx context.Context, in *api.CreateDocumentRequest) (*api.CreateDocumentResponse, error)
	GetDocument(ctx context.Context, in *api.GetDocumentRequest) (*api.GetDocumentResponse, error)
	GetListDocument(ctx context.Context, in *api.GetListDocumentRequest) (*api.GetListDocumentResponse, error)
	UpdateDocument(ctx context.Context, in *api.UpdateDocumentRequest) (*api.EmptyResponse, error)
	DeleteDocument(ctx context.Context, in *api.DeleteDocumentRequest) (*api.EmptyResponse, error)
	// test
}

type dynamicDataAdapter struct {
	grpcAddr string
	client   api.DynamicDataServiceClient
}

func NewDynamicDataAdapter(dynamicDataGRPCAddr string) (DynamicDataAdapter, error) {
	a := dynamicDataAdapter{
		grpcAddr: dynamicDataGRPCAddr,
		client:   nil, // lazy init
	}
	return &a, a.connect()
}

func (a *dynamicDataAdapter) connect() error {
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

		a.client = api.NewDynamicDataServiceClient(conn)
	}
	return nil
}

func (a *dynamicDataAdapter) CreateCollection(ctx context.Context, in *api.CreateCollectionRequest) (*api.CreateCollectionResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.CreateCollection(ctxWithTimeout, in)
}

func (a *dynamicDataAdapter) GetListCollections(ctx context.Context, in *api.GetListCollectionsRequest) (*api.GetListCollectionsResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.GetListCollections(ctxWithTimeout, in)
}

func (a *dynamicDataAdapter) GetCollection(ctx context.Context, in *api.GetCollectionRequest) (*api.GetCollectionResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.GetCollection(ctxWithTimeout, in)
}

func (a *dynamicDataAdapter) UpdateCollection(ctx context.Context, in *api.UpdateCollectionRequest) (*api.EmptyResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.UpdateCollection(ctxWithTimeout, in)
}

func (a *dynamicDataAdapter) DeleteCollection(ctx context.Context, in *api.DeleteCollectionRequest) (*api.EmptyResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.DeleteCollection(ctxWithTimeout, in)
}

func (a *dynamicDataAdapter) CreateDocument(ctx context.Context, in *api.CreateDocumentRequest) (*api.CreateDocumentResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.CreateDocument(ctxWithTimeout, in)
}

func (a *dynamicDataAdapter) GetDocument(ctx context.Context, in *api.GetDocumentRequest) (*api.GetDocumentResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.GetDocument(ctxWithTimeout, in)
}

func (a *dynamicDataAdapter) GetListDocument(ctx context.Context, in *api.GetListDocumentRequest) (*api.GetListDocumentResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.GetListDocument(ctxWithTimeout, in)
}

func (a *dynamicDataAdapter) UpdateDocument(ctx context.Context, in *api.UpdateDocumentRequest) (*api.EmptyResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.UpdateDocument(ctxWithTimeout, in)
}

func (a *dynamicDataAdapter) DeleteDocument(ctx context.Context, in *api.DeleteDocumentRequest) (*api.EmptyResponse, error) {
	if err := a.connect(); err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, ADAPTER_CONTEXT_TIMEOUT_DEFAULT)
	defer cancel()

	return a.client.DeleteDocument(ctxWithTimeout, in)
}
