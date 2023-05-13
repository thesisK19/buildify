package service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/dynamic_data/api"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/model"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/util"
	errors_lib "github.com/thesisK19/buildify/library/errors"
	server_lib "github.com/thesisK19/buildify/library/server"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *Service) CreateCollection(ctx context.Context, in *api.CreateCollectionRequest) (*api.CreateCollectionResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateCollection")

	username := server_lib.GetUsernameFromContext(ctx)

	projectObjectId, err := primitive.ObjectIDFromHex(in.ProjectId)
	if err != nil {
		logger.WithError(err).Error("failed to convert ObjectIDFromHex")
		return nil, errors_lib.ToInvalidArgumentError(err)
	}

	newId, err := s.repository.CreateCollection(ctx, model.Collection{
		Name:      in.Name,
		ProjectId: projectObjectId,
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

func (s *Service) GetListCollections(ctx context.Context, in *api.GetListCollectionsRequest) (*api.GetListCollectionsResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetListCollections")

	username := server_lib.GetUsernameFromContext(ctx)

	var (
		projectObjectId = primitive.NilObjectID
		err             error
	)

	projectObjectId, err = primitive.ObjectIDFromHex(in.ProjectId)
	if err != nil {
		logger.WithError(err).Error("failed to convert ObjectIDFromHex")
		return nil, errors_lib.ToInvalidArgumentError(err)
	}

	listCollections, err := s.repository.GetListCollections(ctx, username, projectObjectId)
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
		structpbValueMap, err := util.ConvertStringMapToStructpbValueMap(doc.Data)
		if err != nil {
			logger.WithError(err).Error("failed to convert map to structpb")
			return nil, err
		}
		resp.Documents = append(resp.Documents, &api.Document{
			Id:           doc.Id,
			Data:         structpbValueMap,
			CollectionId: doc.CollectionId,
		})
	}

	return &resp, nil
}

func (s *Service) GetCollectionMapping(ctx context.Context, in *api.GetCollectionMappingRequest) (*api.GetCollectionMappingResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetCollectionMapping")

	username := server_lib.GetUsernameFromContext(ctx)

	var (
		projectObjectId = primitive.NilObjectID
		err             error
	)

	projectObjectId, err = primitive.ObjectIDFromHex(in.ProjectId)
	if err != nil {
		logger.WithError(err).Error("failed to convert ObjectIDFromHex")
		return nil, errors_lib.ToInvalidArgumentError(err)
	}

	collectionMap, err := s.repository.GetCollectionMapping(ctx, username, projectObjectId)
	if err != nil {
		logger.WithError(err).Error("failed to repo.GetListCollections")
		return nil, err
	}

	dynamicDataTypes := make(map[int32]string)
	dynamicDataTypes[0] = "unknown"
	dynamicDataTypes[1] = "string"
	dynamicDataTypes[2] = "number"

	result := make(map[int32]*api.CollectionInfo)
	for id, coll := range *collectionMap {
		// convert dataTypes
		stringDataTypes := []string{}
		for _, dataType := range coll.DataTypes {
			stringDataTypes = append(stringDataTypes, dynamicDataTypes[dataType])
		}

		result[id] = &api.CollectionInfo{
			Name:      coll.Name,
			DataKeys:  coll.DataKeys,
			DataTypes: stringDataTypes,
			Documents: make(map[int32]*structpb.Value),
		}
		// convert doc
		structpbValueMap, err := util.ConvertInt32MapToStructpbValueMap(coll.Documents)
		if err != nil {
			logger.WithError(err).Error("failed to convert map to structpb")
			return nil, err
		}
		result[id].Documents = structpbValueMap
	}

	return &api.GetCollectionMappingResponse{
		Data: result,
	}, nil
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
		structpbValueMap, err := util.ConvertStringMapToStructpbValueMap(doc.Data)
		if err != nil {
			logger.WithError(err).Error("failed to convert map to structpb")
			return nil, err
		}
		docs = append(docs, &api.Document{
			Id:           doc.Id,
			Data:         structpbValueMap,
			CollectionId: doc.CollectionId,
		})
	}

	return &api.GetCollectionResponse{
		Id:        coll.Id,
		Name:      coll.Name,
		ProjectId: coll.ProjectId,
		DataKeys:  coll.DataKeys,
		DataTypes: coll.DataTypes,
		Documents: docs,
	}, nil
}

func (s *Service) UpdateCollection(ctx context.Context, in *api.UpdateCollectionRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UpdateCollection")
	username := server_lib.GetUsernameFromContext(ctx)

	projectObjectId, err := primitive.ObjectIDFromHex(in.ProjectId)
	if err != nil {
		logger.WithError(err).Error("failed to convert ObjectIDFromHex")
		return nil, errors_lib.ToInvalidArgumentError(err)
	}

	err = s.repository.UpdateCollection(ctx, model.Collection{
		Id:        in.Id,
		Name:      in.Name,
		ProjectId: projectObjectId,
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
