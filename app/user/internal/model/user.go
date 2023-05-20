package model

import (
	"github.com/golang-jwt/jwt"
)

// User Model
type User struct {
	FullName  string `bson:"full_name,omitempty" validate:"required,alpha"`
	Email     string `bson:"email,omitempty" validate:"required,alpha"`
	Username  string `bson:"username,omitempty" validate:"required,alpha"`
	Password  string `bson:"password,omitempty" validate:"required"`
	AvatarUrl string `bson:"avatar_url,omitempty"`
}

type CreateUserParams struct {
	FullName  string `bson:"full_name,omitempty" validate:"required,alpha"`
	Email     string `bson:"email,omitempty" validate:"required,alpha"`
	Username  string `bson:"username,omitempty" validate:"required,alpha"`
	Password  string `bson:"password,omitempty" validate:"required"`
	AvatarUrl string `bson:"avatar_url,omitempty"`
}

type UpdateUserParams struct {
	FullName  string `bson:"full_name,omitempty" validate:"alpha"`
	Email     string `bson:"email,omitempty" validate:"alpha"`
	Password  string `bson:"password,omitempty" validate:""`
	AvatarUrl string `bson:"avatar_url,omitempty"`
}

// Claims struct for JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
