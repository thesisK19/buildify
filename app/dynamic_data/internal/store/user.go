package store

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/constant"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/model"
	"gopkg.in/mgo.v2/bson"
)

func (r *repository) CreateUser(ctx context.Context, username string, lastDocId int32, lastCollId int32) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateUser")

	_, err := r.db.Collection(constant.USER_COLL).InsertOne(ctx, model.User{
		Username:   username,
		LastDocId:  lastDocId,
		LastCollId: lastCollId,
	})
	if err != nil {
		logger.WithError(err).Error("failed to InsertOne")
		return err
	}

	return nil
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
