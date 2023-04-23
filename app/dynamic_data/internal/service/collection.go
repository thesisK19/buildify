package service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/dynamic_data/api"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/model"
	server_lib "github.com/thesisK19/buildify/library/server"
)

func (s *Service) CreateCollection(ctx context.Context, in *api.CreateCollectionRequest) (*api.CreateCollectionResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateCollection")

	username := server_lib.GetUsernameFromContext(ctx)

	newId, err := s.repository.CreateCollection(ctx, model.Collection{
		Name:      in.Name,
		DataKeys:  in.DataKeys,
		DataTypes: in.DataTypes,
		Username:  username,
	})
	if err != nil {
		logger.WithError(err).Error("failed to repo.CreateUser")
		return nil, err
	}

	return &api.CreateCollectionResponse{
		Id: newId,
	}, nil
}

func (s *Service) GetListCollections(ctx context.Context, in *api.EmptyRequest) (*api.GetListCollectionsResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetListCollections")

	username := server_lib.GetUsernameFromContext(ctx)

	listCollections, err := s.repository.GetListCollections(ctx, username)
	if err != nil {
		logger.WithError(err).Error("failed to repo.GetListCollections")
		return nil, err
	}

	var resp api.GetListCollectionsResponse
	for _, coll := range listCollections.Collections {
		resp.Collections = append(resp.Collections, &api.Collection{
			Id:        coll.Id,
			Name:      coll.Name,
			DataKeys:  coll.DataKeys,
			DataTypes: coll.DataTypes,
		})
	}
	for _, doc := range listCollections.Documents {
		resp.Documents = append(resp.Documents, &api.Document{
			Id:           doc.Id,
			Data:         doc.Data,
			CollectionId: doc.CollectionId,
		})
	}

	return &resp, nil
}

func (s *Service) GetCollection(ctx context.Context, in *api.GetCollectionRequest) (*api.GetCollectionResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetCollection")

	username := server_lib.GetUsernameFromContext(ctx)

	coll, err := s.repository.GetCollection(ctx, username, in.Id)
	if err != nil {
		logger.WithError(err).Error("failed to repo.GetCollection")
		return nil, err
	}
	var docs []*api.Document

	for _, doc := range coll.Documents {
		docs = append(docs, &api.Document{
			Id:           doc.Id,
			Data:         doc.Data,
			CollectionId: doc.CollectionId,
		})
	}

	return &api.GetCollectionResponse{
		Id:        coll.Id,
		Name:      coll.Name,
		DataKeys:  coll.DataKeys,
		DataTypes: coll.DataTypes,
		Documents: docs,
	}, nil
}

func (s *Service) UpdateCollection(ctx context.Context, in *api.UpdateCollectionRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UpdateCollection")

	username := server_lib.GetUsernameFromContext(ctx)

	err := s.repository.UpdateCollection(ctx, model.Collection{
		Id:        in.Id,
		Name:      in.Name,
		DataKeys:  in.DataKeys,
		DataTypes: in.DataTypes,
		Username:  username,
	})
	if err != nil {
		logger.WithError(err).Error("failed to repo.UpdateCollection")
		return nil, err
	}

	return &api.EmptyResponse{}, nil
}

func (s *Service) DeleteCollection(ctx context.Context, in *api.DeleteCollectionRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "DeleteCollection")

	username := server_lib.GetUsernameFromContext(ctx)

	err := s.repository.DeleteCollection(ctx, username, in.Id)
	if err != nil {
		logger.WithError(err).Error("failed to repo.DeleteCollection")
		return nil, err
	}

	return &api.EmptyResponse{}, nil
}
