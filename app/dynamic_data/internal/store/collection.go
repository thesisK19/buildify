package store

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/constant"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/dto"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/model"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Function to get the next available collection Id for a user
func (r *repository) getNextCollId(ctx context.Context, username string) (int32, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "getNextCollId")
	// Create a filter for the username
	filter := bson.M{"username": username}

	// Update the user document to increment the LastCollId field by 1 and return the updated document
	update := bson.M{"$inc": bson.M{"last_coll_id": 1}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	var user model.User
	err := r.db.Collection(constant.USER_COLL).FindOneAndUpdate(ctx, filter, update, &opt).Decode(&user)
	if err == nil {
		return user.LastCollId, nil
	}

	if err != mongo.ErrNoDocuments {
		logger.WithError(err).Error("failed to FindOneAndUpdate")
		return 0, err
	}

	// If user not found, create a new user and return 1
	err = r.CreateUser(ctx, username, 0, 1)
	if err != nil {
		logger.WithError(err).Error("failed to create user")
		return 0, err
	}
	return 1, nil
}

// Function to create a new collection for a user
func (r *repository) CreateCollection(ctx context.Context, coll model.Collection) (int32, error) {
	collectionId, err := r.getNextCollId(ctx, coll.Username)
	if err != nil {
		return 0, err
	}

	coll.Id = collectionId

	// Insert the new collection into the collection
	_, err = r.db.Collection(constant.COLLECTION_COLL).InsertOne(ctx, &coll)
	if err != nil {
		return 0, err
	}

	return collectionId, nil
}

func (r *repository) GetListCollections(ctx context.Context, username string, projectId primitive.ObjectID) (*dto.ListCollections, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetListCollections")

	var (
		result        dto.ListCollections
		collectionIds []int32
	)

	filter := bson.M{"username": username, "project_id": projectId}

	// Find all collections
	cursor, err := r.db.Collection(constant.COLLECTION_COLL).Find(ctx, filter)
	if err != nil {
		logger.WithError(err).Error("failed to Find collections")
		return nil, fmt.Errorf("failed to find collections: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var coll model.Collection
		if err := cursor.Decode(&coll); err != nil {
			logger.WithError(err).Error("failed to decode collection")
			return nil, fmt.Errorf("failed to decode collection: %v", err)
		}
		result.Collections = append(result.Collections, dto.Collection{
			Id:        coll.Id,
			Name:      coll.Name,
			DataKeys:  coll.DataKeys,
			DataTypes: coll.DataTypes,
		})
		collectionIds = append(collectionIds, coll.Id)
	}

	if len(collectionIds) == 0 {
		return &dto.ListCollections{}, nil
	}

	cursor, err = r.db.Collection(constant.DOCUMENT_COLL).Find(ctx, bson.M{"collection_id": bson.M{"$in": collectionIds}})
	if err != nil {
		return nil, fmt.Errorf("failed to find documents: %v", err)
	}

	for cursor.Next(ctx) {
		var doc model.Document
		if err := cursor.Decode(&doc); err != nil {
			return nil, fmt.Errorf("failed to decode document: %v", err)
		}
		result.Documents = append(result.Documents, dto.Document{
			Id:           doc.Id,
			Data:         util.ConvertBsonMToMap(doc.Data),
			CollectionId: doc.CollectionId,
		})
	}

	return &result, nil
}
func (r *repository) GetCollection(ctx context.Context, username string, id int32) (*dto.GetCollection, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetCollection")

	var result dto.GetCollection

	filter := bson.M{"id": id, "username": username}
	var coll model.Collection

	if err := r.db.Collection(constant.COLLECTION_COLL).FindOne(ctx, filter).Decode(&coll); err != nil {
		logger.WithError(err).Error("failed to FindOne")
		return nil, err
	}

	result.Id = coll.Id
	result.Name = coll.Name
	result.ProjectId = coll.ProjectId.Hex()
	result.DataKeys = coll.DataKeys
	result.DataTypes = coll.DataTypes

	cursor, err := r.db.Collection(constant.DOCUMENT_COLL).Find(ctx, bson.M{"collection_id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to find documents: %v", err)
	}

	for cursor.Next(ctx) {
		var doc model.Document
		if err := cursor.Decode(&doc); err != nil {
			return nil, fmt.Errorf("failed to decode document: %v", err)
		}
		result.Documents = append(result.Documents, dto.Document{
			Id:           doc.Id,
			Data:         util.ConvertBsonMToMap(doc.Data),
			CollectionId: doc.CollectionId,
		})
	}

	return &result, nil
}

func (r *repository) UpdateCollection(ctx context.Context, coll model.Collection) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UpdateCollection")

	filter := bson.M{"id": coll.Id, "username": coll.Username}
	updateDoc := bson.M{"$set": coll}

	result, err := r.db.Collection(constant.COLLECTION_COLL).UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		logger.WithError(err).Error("failed to UpdateOne")
		return err
	}

	// Check the number of documents matched and modified
	if result.MatchedCount == 0 {
		err = fmt.Errorf("no documents matched the filter")
		logger.WithError(err).Error("failed to UpdateOne")
		return err
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("no documents were modified")
		logger.WithError(err).Error("failed to UpdateOne")
		return err
	}

	return nil
}

func (r *repository) DeleteCollection(ctx context.Context, username string, id int32) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "DeleteCollection")

	filter := bson.M{"id": id, "username": username}

	// Delete the document
	result, err := r.db.Collection(constant.COLLECTION_COLL).DeleteOne(ctx, filter)
	if err != nil {
		logger.WithError(err).Error("failed to DeleteOne")
		return err
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		err = fmt.Errorf("document not found")
		logger.WithError(err).Error("failed to DeleteOne")
		return err
	}

	filter2 := bson.M{"collection_id": id, "username": username}
	// Delete all documents matching the filter
	result2, err := r.db.Collection(constant.DOCUMENT_COLL).DeleteMany(ctx, filter2)
	if err != nil {
		logger.WithError(err).Error("failed to DeleteMany")
		return err
	}

	// Check if any documents were deleted
	if result2.DeletedCount == 0 {
		logger.Warn("no documents found to delete")
		return err
	}

	return nil
}
