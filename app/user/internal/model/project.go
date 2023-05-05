package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Project Model
type Project struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name,omitempty"`
	Username       string             `bson:"username,omitempty"`
	CompressString string             `bson:"compress_string,omitempty"`
}
