package adapter

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	genCodeApi "github.com/thesisK19/buildify/app/gen-code/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GenCodeClient interface {
	HelloWorld(ctx context.Context) (*genCodeApi.HelloWorldResponse, error)
}

type ImplGenCodeClient struct {
	host   string
	client genCodeApi.GenCodeServiceClient
}

func (c *ImplGenCodeClient) connect() error {
	if c.client == nil {
		connectorChannel, err := grpc.DialContext(context.Background(), c.host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}

		c.client = genCodeApi.NewGenCodeServiceClient(connectorChannel)
	}
	return nil
}

// TODO: re-check
func NewGenCodeClient(connectorAddr string) (*ImplGenCodeClient, error) {
	serviceGRPC := ImplGenCodeClient{
		host: connectorAddr,
	}
	connectorChannel, err := grpc.DialContext(context.Background(), connectorAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &serviceGRPC, err
	}

	serviceGRPC.client = genCodeApi.NewGenCodeServiceClient(connectorChannel)

	return &serviceGRPC, nil
}

func (c *ImplGenCodeClient) HelloWorld(ctx context.Context) (*genCodeApi.HelloWorldResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "HelloWorld")

	if err := c.connect(); err != nil {
		logger.WithError(err).Error("Failed to connect client")
		return nil, err
	}

	req := genCodeApi.EmptyRequest{}

	resp, err := c.client.HelloWorld(ctx, &req)
	if err != nil {
		logger.WithError(err).Error("Failed to call HelloWorld")
		return nil, err
	}

	return resp, nil
}
