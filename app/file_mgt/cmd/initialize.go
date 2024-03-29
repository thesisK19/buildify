package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/mux"
	"github.com/thesisK19/buildify/app/file_mgt/config"
	"github.com/thesisK19/buildify/app/file_mgt/internal/handler"
	"github.com/thesisK19/buildify/app/file_mgt/internal/store"
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
		log.Println("failed to connect mongoDB. ", "connectionStr=", connectionString, "err=", err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("failed to ping mongoDB. ", "connectionStr=", connectionString, "err=", err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func newService(cfg *config.Config) (*Service, error) {
	mongoClient, err := connectDB(cfg.MongoDB, cfg.ServiceDB)
	if err != nil {
		log.Println("failed to connectDB", "err=", err)
		return nil, err
	}
	repository := store.NewRepository(cfg, mongoClient)

	s := NewService(cfg, repository, mux.NewRouter())
	// s.UseMiddleware(handler.JSONContentTypeMiddleware)
	s.UseMiddleware(handler.Recovery)

	// Add routes
	s.setRouter()

	return s, nil
}
