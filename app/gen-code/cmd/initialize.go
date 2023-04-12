package main

import (
	"context"
	"time"

	"github.com/thesisK19/buildify/app/gen-code/config"
	"github.com/thesisK19/buildify/app/gen-code/internal/service"
	"github.com/thesisK19/buildify/app/gen-code/internal/store"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB(connectionString string, serviceDB string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		logger.WithError(err).Error("Failed to connect mongoDB. ", "connectionStr=", connectionString)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logger.WithError(err).Error("Failed to ping mongoDB. ", "connectionStr=", connectionString)
		return nil, err
	}

	logger.Info("Connected to MongoDB!!")
	return client, nil
}

func newService(cfg *config.Config) (*service.Service, error) {
	mongoClient, err := connectDB(cfg.MongoDB, cfg.ServiceDB)
	if err != nil {
		logger.WithError(err).Error("Failed to connectDB")
		return nil, err
	}
	repository := store.NewRepository(cfg, mongoClient)

	return service.NewService(cfg, repository), nil
}
