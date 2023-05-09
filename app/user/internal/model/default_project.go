package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// DefaultProject Model
type DefaultProject struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	Type           int                `bson:"type,omitempty"`
	Name           string             `bson:"name,omitempty"`
	CompressString string             `bson:"compress_string,omitempty"`
}
