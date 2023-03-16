package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User Model
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty" validate:"required,alpha"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,alpha"`
}
