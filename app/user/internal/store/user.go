package store

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/user/internal/constant"
	"github.com/thesisK19/buildify/app/user/internal/model"
	errors_lib "github.com/thesisK19/buildify/library/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func (r *repository) CreateUser(ctx context.Context, params model.CreateUserParams) (*string, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateUser")

	result, err := r.db.Collection(constant.USER_COLL).InsertOne(ctx, params)
	if err != nil {
		logger.WithError(err).Error("failed to InsertOne")
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors_lib.ToInvalidArgumentError(fmt.Errorf("username already exists"))
		}
		return nil, err
	}

	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		logger.WithError(err).Error("failed to convert objectID")
		return nil, err
	}

	stringId := objID.Hex()

	return &stringId, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetUserByUsername")

	filter := bson.M{"username": username}
	var user model.User

	if err := r.db.Collection(constant.USER_COLL).FindOne(ctx, filter).Decode(&user); err != nil {
		logger.WithError(err).Error("failed to FindOne")
		return nil, err
	}
	return &user, nil
}

func (r *repository) UpdateUserByUsername(ctx context.Context, username string, params model.UpdateUserParams) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UpdateUserByUsername")

	filter := bson.M{"username": username}
	updateDoc := bson.M{"$set": params}

	result, err := r.db.Collection(constant.USER_COLL).UpdateOne(ctx, filter, updateDoc)
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
