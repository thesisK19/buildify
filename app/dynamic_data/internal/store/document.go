package store

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/constant"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Function to get the next available document Id for a user
func (r *repository) getNextDocId(ctx context.Context, username string) (int32, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "getNextDocId")

	// Create a filter for the username
	filter := bson.M{"username": username}

	// Update the user document to increment the LastDocId field by 1 and return the updated document
	update := bson.M{"$inc": bson.M{"last_doc_id": 1}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	var user model.User
	err := r.db.Collection(constant.USER_COLL).FindOneAndUpdate(ctx, filter, update, &opt).Decode(&user)
	if err == nil {
		return user.LastDocId, nil
	}

	if err != mongo.ErrNoDocuments {
		logger.WithError(err).Error("failed to FindOneAndUpdate")
		return 0, err
	}

	// If user not found, create a new user and return 1
	err = r.CreateUser(ctx, username, 1, 0)

	if err != nil {
		logger.WithError(err).Error("failed to create user")
		return 0, err
	}
	return 1, nil
}

// Function to create a new document for a user
func (r *repository) CreateDocument(ctx context.Context, doc model.Document) (int32, error) {
	documentId, err := r.getNextDocId(ctx, doc.Username)
	if err != nil {
		return 0, err
	}

	doc.Id = documentId
	// Insert the new document into the document collection
	_, err = r.db.Collection(constant.DOCUMENT_COLL).InsertOne(ctx, &doc)
	if err != nil {
		return 0, err
	}

	return documentId, nil
}

func (r *repository) GetDocument(ctx context.Context, username string, id int32) (*model.Document, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetDocument")

	filter := bson.M{"id": id, "username": username}
	var doc model.Document

	if err := r.db.Collection(constant.DOCUMENT_COLL).FindOne(ctx, filter).Decode(&doc); err != nil {
		logger.WithError(err).Error("failed to FindOne")
		return nil, err
	}
	return &doc, nil
}

func (r *repository) UpdateDocument(ctx context.Context, doc model.Document) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UpdateDocument")

	filter := bson.M{"id": doc.Id, "username": doc.Username}
	updateDoc := bson.M{"$set": doc}

	result, err := r.db.Collection(constant.DOCUMENT_COLL).UpdateOne(ctx, filter, updateDoc)
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

func (r *repository) DeleteDocument(ctx context.Context, username string, id int32) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "DeleteDocument")

	filter := bson.M{"id": id, "username": username}

	// Delete the document
	result, err := r.db.Collection(constant.DOCUMENT_COLL).DeleteOne(ctx, filter)
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

	return nil
}

func (r *repository) GetListDocuments(ctx context.Context, username string, collectionId int32) ([]*model.Document, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetListDocuments")

	var result []*model.Document

	filter := bson.M{"username": username, "collection_id": collectionId}

	// Find all documents
	cursor, err := r.db.Collection(constant.DOCUMENT_COLL).Find(ctx, filter)
	if err != nil {
		logger.WithError(err).Error("failed to Find documents")
		return nil, fmt.Errorf("failed to find documents: %v", err)
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &result); err != nil {
		logger.WithError(err).Error("failed to decode documents")
		return nil, fmt.Errorf("failed to decode documents: %v", err)
	}
	return result, nil
}
