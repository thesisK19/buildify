package store

import (
	"context"

	"github.com/thesisK19/buildify/app/dynamic_data/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateDocument(ctx context.Context, username string, collectionID int32, data map[string]string) (int32, error)
	CreateCollection(ctx context.Context, username string, name string, semanticKey string, keys []string, types []int32) (int32, error)
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
