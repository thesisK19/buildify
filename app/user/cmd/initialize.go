package main

import (
	"context"
	"fmt"
	"log"
	"thesis/be/app/user/config"
	"thesis/be/app/user/internal/service"
	"thesis/be/app/user/internal/store"
	"time"

	"github.com/sirupsen/logrus"
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
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func newService(cfg *config.Config, log *logrus.Logger) (*service.Service, error) {
	mongoClient, err := connectDB(cfg.MongoDB, cfg.ServiceDB)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"DB config": cfg.MongoDB,
			"err":       err,
		}).Error("Error connect database")
		return nil, err
	}
	repository := store.NewRepository(log, cfg, mongoClient)

	return service.NewService(cfg, log, repository), nil
}
