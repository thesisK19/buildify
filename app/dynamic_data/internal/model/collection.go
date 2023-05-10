package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Collection struct {
	Id        int32              `bson:"id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	DataKeys  []string           `bson:"data_keys,omitempty"`
	DataTypes []int32            `bson:"data_types,omitempty"`
	Username  string             `bson:"username,omitempty"`
	ProjectId primitive.ObjectID `bson:"project_id,omitempty"`
}

// DocumentIds []int32  `bson:"document_ids,omitempty"`
