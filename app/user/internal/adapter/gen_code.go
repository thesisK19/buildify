package adapter

import (
	"context"
	"fmt"

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

func (c *ImplGenCodeClient) getConnection() error {
	if c.client == nil {
		connectorChannel, err := grpc.DialContext(context.Background(), c.host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}

		c.client = genCodeApi.NewGenCodeServiceClient(connectorChannel)
	}
	return nil
}

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
	if err := c.getConnection(); err != nil {
		return nil, fmt.Errorf("error to connect to gen code service %v", err.Error())
	}
	var err error
	var resp *genCodeApi.HelloWorldResponse
	req := genCodeApi.HelloWorldRequest{}
	resp, err = c.client.HelloWorld(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("GenCode adapter: HelloWorld failed %w", err)
	}

	if resp.Code != "OK" {
		return nil, fmt.Errorf("GenCode adapter: HelloWorld failed %v", resp.Message)
	}

	return resp, nil
}
