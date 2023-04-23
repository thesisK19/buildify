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

// Function to get the next available document ID for a user
func (r *repository) getNextDocID(ctx context.Context, username string) (int32, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "getNextDocID")

	// Create a filter for the username
	filter := bson.M{"username": username}

	// Update the user document to increment the LastDocID field by 1 and return the updated document
	update := bson.M{"$inc": bson.M{"last_doc_id": 1}}
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
	err = r.CreateUser(ctx, username, 1, 0)
	if err != nil {
		logger.WithError(err).Error("Failed to create user")
		return 0, err
	}
	return 1, nil
}

// Function to create a new document for a user
func (r *repository) CreateDocument(ctx context.Context, username string, collectionID int32, data map[string]string) (int32, error) {
	documentID, err := r.getNextDocID(ctx, username)
	if err != nil {
		return 0, err
	}

	dataConvert := make(map[string]interface{})
	for key, value := range data {
		dataConvert[key] = value
	}

	document := model.Document{
		ID:           documentID,
		CollectionID: collectionID,
		Data:         dataConvert,
		Username:     username,
	}

	// Insert the new document into the document collection
	_, err = r.db.Collection(constant.DOCUMENT_COLL).InsertOne(ctx, &document)
	if err != nil {
		return 0, err
	}

	// TODO: return response
	return 0, nil
}
