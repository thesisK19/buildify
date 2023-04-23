package service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/dynamic_data/api"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/model"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/util"
	server_lib "github.com/thesisK19/buildify/library/server"
)

func (s *Service) CreateDocument(ctx context.Context, in *api.CreateDocumentRequest) (*api.CreateDocumentResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateDocument")

	username := server_lib.GetUsernameFromContext(ctx)

	newId, err := s.repository.CreateDocument(ctx, model.Document{
		Data:         util.ConvertMapToBsonM(in.Data),
		CollectionId: in.CollectionId,
		Username:     username,
	})
	if err != nil {
		logger.WithError(err).Error("failed to repo.CreateUser")
		return nil, err
	}

	return &api.CreateDocumentResponse{
		Id: newId,
	}, nil
}

func (s *Service) GetDocument(ctx context.Context, in *api.GetDocumentRequest) (*api.GetDocumentResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetDocument")

	username := server_lib.GetUsernameFromContext(ctx)

	doc, err := s.repository.GetDocument(ctx, username, in.Id)
	if err != nil {
		logger.WithError(err).Error("failed to repo.GetDocument")
		return nil, err
	}

	return &api.GetDocumentResponse{
		Id:           doc.Id,
		Data:         util.ConvertBsonMToMap(doc.Data),
		CollectionId: doc.CollectionId,
	}, nil
}

func (s *Service) GetListDocument(ctx context.Context, in *api.GetListDocumentRequest) (*api.GetListDocumentResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetListDocument")

	username := server_lib.GetUsernameFromContext(ctx)

	docs, err := s.repository.GetListDocuments(ctx, username, in.CollectionId)
	if err != nil {
		logger.WithError(err).Error("failed to repo.GetListDocuments")
		return nil, err
	}

	var resp api.GetListDocumentResponse
	for _, doc := range docs {
		resp.Documents = append(resp.Documents, &api.Document{
			Id:           doc.Id,
			Data:         util.ConvertBsonMToMap(doc.Data),
			CollectionId: doc.CollectionId,
		})
	}
	return &resp, nil
}

func (s *Service) UpdateDocument(ctx context.Context, in *api.UpdateDocumentRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UpdateDocument")

	username := server_lib.GetUsernameFromContext(ctx)

	err := s.repository.UpdateDocument(ctx, model.Document{
		Id:           in.Id,
		Data:         util.ConvertMapToBsonM(in.Data),
		CollectionId: in.CollectionId,
		Username:     username,
	})
	if err != nil {
		logger.WithError(err).Error("failed to repo.UpdateDocument")
		return nil, err
	}

	return &api.EmptyResponse{}, nil
}

func (s *Service) DeleteDocument(ctx context.Context, in *api.DeleteDocumentRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "DeleteDocument")

	username := server_lib.GetUsernameFromContext(ctx)

	err := s.repository.DeleteDocument(ctx, username, in.Id)
	if err != nil {
		logger.WithError(err).Error("failed to repo.DeleteDocument")
		return nil, err
	}

	return &api.EmptyResponse{}, nil
}
