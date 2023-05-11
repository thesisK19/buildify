package store

import (
	"context"

	"github.com/thesisK19/buildify/app/dynamic_data/config"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/dto"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	// document
	CreateDocument(ctx context.Context, doc model.Document) (int32, error)
	GetDocument(ctx context.Context, username string, id int32) (*model.Document, error)
	GetListDocuments(ctx context.Context, username string, collectionId int32) ([]*model.Document, error)
	UpdateDocument(ctx context.Context, doc model.Document) error
	DeleteDocument(ctx context.Context, username string, id int32) error
	// collection
	CreateCollection(ctx context.Context, coll model.Collection) (int32, error)
	GetCollection(ctx context.Context, username string, id int32) (*dto.GetCollection, error)
	GetListCollections(ctx context.Context, username string, projectId primitive.ObjectID) (*dto.ListCollections, error)
	GetCollectionMapping(ctx context.Context, username string, projectId primitive.ObjectID) (*dto.CollectionMap, error)
	UpdateCollection(ctx context.Context, coll model.Collection) error
	DeleteCollection(ctx context.Context, username string, id int32) error
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
