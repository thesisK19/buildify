package store

import (
	"context"
	"thesis/be/app/gen-code/config"
	"thesis/be/app/gen-code/internal/model"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	// user
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*model.User, error)

	// repository
	Ping() error
	Close() error
}

type repository struct {
	log         *logrus.Logger
	config      *config.Config
	mongoClient *mongo.Client
	db          *mongo.Database
}

func NewRepository(log *logrus.Logger, config *config.Config, mongoClient *mongo.Client) Repository {
	db := mongoClient.Database(config.ServiceDB)

	return &repository{
		log:         log,
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
