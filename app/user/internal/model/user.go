package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User Model
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,alpha"`
	Username string             `json:"username,omitempty" bson:"username,omitempty" validate:"required,alpha"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,alpha"`
}

type CreateUserParams struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty" validate:"required,alpha"`
	Username string `json:"username,omitempty" bson:"username,omitempty" validate:"required,alpha"`
	Password string `json:"password,omitempty" bson:"password,omitempty" validate:"required,alpha"`
}

type UpdateUserParams struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty" validate:"alpha"`
	Username string `json:"username,omitempty" bson:"username,omitempty" validate:"alpha"`
	Password string `json:"password,omitempty" bson:"password,omitempty" validate:"alpha"`
}
