package store

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/constant"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Function to get the next available collection ID for a user
func (r *repository) getNextCollID(ctx context.Context, username string) (int32, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "getNextCollID")
	// Create a filter for the username
	filter := bson.M{"username": username}

	// Update the user document to increment the LastCollID field by 1 and return the updated document
	update := bson.M{"$inc": bson.M{"last_coll_id": 1}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	var user model.User
	err := r.db.Collection(constant.USER_COLL).FindOneAndUpdate(ctx, filter, update, &opt).Decode(&user)
	if err == nil {
		return user.LastCollID, nil
	}

	if err != mongo.ErrNoDocuments {
		logger.WithError(err).Error("Failed to FindOneAndUpdate")
		return 0, err
	}

	// If user not found, create a new user and return 1
	err = r.CreateUser(ctx, username, 0, 1)
	if err != nil {
		logger.WithError(err).Error("Failed to create user")
		return 0, err
	}
	return 1, nil
}

// Function to create a new collection for a user
func (r *repository) CreateCollection(ctx context.Context, username string, name string, semanticKey string, keys []string, types []int32) (int32, error) {
	collectionID, err := r.getNextCollID(ctx, username)
	if err != nil {
		return 0, err
	}

	collection := model.Collection{
		ID:          collectionID,
		Name:        name,
		SemanticKey: semanticKey,
		Keys:        keys,
		Types:       types,
		DocumentIDs: []int32{},
		Username:    username,
	}

	// Insert the new collection into the collection
	_, err = r.db.Collection(constant.COLLECTION_COLL).InsertOne(ctx, &collection)
	if err != nil {
		return 0, err
	}

	// TODO: return response
	return 0, nil
}
