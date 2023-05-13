package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Document struct {
	Id           int32  `bson:"id,omitempty"`
	Data         bson.M `bson:"data,omitempty"`
	CollectionId int32  `bson:"collection_id,omitempty"`
	Username     string `bson:"username,omitempty"`
}
