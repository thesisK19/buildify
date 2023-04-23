package service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/dynamic_data/api"
)

func (s *Service) CreateDocument(ctx context.Context, in *api.CreateDocumentRequest) (*api.CreateDocumentResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "SignUp")

	_, err := s.repository.CreateDocument(ctx, "ynek", in.CollectionId, in.Data)
	if err != nil {
		logger.WithError(err).Error("Failed to repo.CreateUser")
		return nil, err
	}

	return &api.CreateDocumentResponse{}, nil
}

func (s *Service) CreateCollection(ctx context.Context, in *api.CreateCollectionRequest) (*api.CreateCollectionResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "SignUp")

	_, err := s.repository.CreateCollection(ctx, "ynek", in.Name, in.SemanticKey, in.Keys, in.Types)
	if err != nil {
		logger.WithError(err).Error("Failed to repo.CreateUser")
		return nil, err
	}

	return &api.CreateCollectionResponse{}, nil
}
