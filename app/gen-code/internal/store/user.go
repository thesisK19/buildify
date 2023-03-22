package store

import (
	"buildify/app/gen-code/internal/model"
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

const usersCollection = "users"

func (r *repository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	usersCollection := r.db.Collection(usersCollection)

	result, err := usersCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "InsertOne")
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.Wrap(err, "ConvertInsertID")
	}

	user.ID = objectID

	return user, nil
}
func (r *repository) GetUserByID(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	usersCollection := r.db.Collection(usersCollection)

	var user model.User
	if err := usersCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, errors.Wrap(err, "FindOne")
	}
	return &user, nil
}
