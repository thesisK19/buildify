package store

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/user/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

const userCol = "users"

func (r *repository) CreateUser(ctx context.Context, params model.CreateUserParams) (*string, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateUser")

	result, err := r.db.Collection(userCol).InsertOne(ctx, params)
	if err != nil {
		logger.WithError(err).Error("Failed to InsertOne")
		return nil, err
	}

	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		logger.WithError(err).Error("Failed to convert objectID")
		return nil, err
	}

	stringId := objID.Hex()

	return &stringId, nil
}

func (r *repository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetUserByID")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.WithError(err).Error("Failed to convert objectID")
		return nil, err
	}

	filter := bson.M{"_id": objID}
	var user model.User

	if err := r.db.Collection(userCol).FindOne(ctx, filter).Decode(&user); err != nil {
		logger.WithError(err).Error("Failed to FindOne")
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetUserByUsername")

	filter := bson.M{"username": username}
	var user model.User

	if err := r.db.Collection(userCol).FindOne(ctx, filter).Decode(&user); err != nil {
		logger.WithError(err).Error("Failed to FindOne")
		return nil, err
	}
	return &user, nil
}

func (r *repository) UpdateUserByID(ctx context.Context, id string, params model.UpdateUserParams) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UpdateUserByID")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	updateDoc := bson.M{"$set": params}

	result, err := r.db.Collection(userCol).UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		logger.WithError(err).Error("Failed to UpdateOne")
		return err
	}

	// Check the number of documents matched and modified
	if result.MatchedCount == 0 {
		return fmt.Errorf("no documents matched the filter")
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("no documents were modified")
	}

	return nil
}

func (r *repository) UpdateUserByUsername(ctx context.Context, username string, params model.UpdateUserParams) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UpdateUserByUsername")

	filter := bson.M{"username": username}
	updateDoc := bson.M{"$set": params}

	result, err := r.db.Collection(userCol).UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		logger.WithError(err).Error("Failed to UpdateOne")
		return err
	}

	// Check the number of documents matched and modified
	if result.MatchedCount == 0 {
		return fmt.Errorf("no documents matched the filter")
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("no documents were modified")
	}

	return nil
}
