package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/thesisK19/buildify/app/file-management/config"
	"github.com/thesisK19/buildify/app/file-management/internal/store"
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

func newService(cfg *config.Config, log *logrus.Logger) (*Service, error) {
	mongoClient, err := connectDB(cfg.MongoDB, cfg.ServiceDB)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"DB config": cfg.MongoDB,
			"err":       err,
		}).Error("Error connect database")
		return nil, err
	}
	repository := store.NewRepository(log, cfg, mongoClient)

	s := NewService(cfg, log, repository, mux.NewRouter())
	// s.UseMiddleware(handler.JSONContentTypeMiddleware)

	// Add routes
	s.setRouter()

	return s, nil
}
