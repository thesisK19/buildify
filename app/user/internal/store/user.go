package store

import (
	"context"
	"fmt"

	"github.com/thesisK19/buildify/app/user/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

const userCol = "users"

func (r *repository) CreateUser(ctx context.Context, params model.CreateUserParams) (*string, error) {
	result, err := r.db.Collection(userCol).InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}

	stringId := objID.Hex()

	return &stringId, nil
}

func (r *repository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	var user model.User

	if err := r.db.Collection(userCol).FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {

	filter := bson.M{"username": username}
	var user model.User

	if err := r.db.Collection(userCol).FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) UpdateUserByID(ctx context.Context, id string, params model.UpdateUserParams) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	updateDoc := bson.M{"$set": params}

	result, err := r.db.Collection(userCol).UpdateOne(ctx, filter, updateDoc)
	if err != nil {
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
	filter := bson.M{"username": username}
	updateDoc := bson.M{"$set": params}

	result, err := r.db.Collection(userCol).UpdateOne(ctx, filter, updateDoc)
	if err != nil {
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
