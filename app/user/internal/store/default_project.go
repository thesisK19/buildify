package store

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/user/internal/constant"
	"github.com/thesisK19/buildify/app/user/internal/model"
	"gopkg.in/mgo.v2/bson"
)

func (r *repository) GetDefaultProjectByType(ctx context.Context, projectType int) (*model.DefaultProject, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetDefaultProjectByType")

	filter := bson.M{"type": projectType}
	var defaultProject model.DefaultProject

	if err := r.db.Collection(constant.DEFAULT_PROJECT_COLL).FindOne(ctx, filter).Decode(&defaultProject); err != nil {
		logger.WithError(err).Error("failed to FindOne")
		return nil, err
	}

	return &defaultProject, nil
}
