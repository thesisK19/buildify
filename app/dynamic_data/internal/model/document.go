package model

import "gopkg.in/mgo.v2/bson"

type Document struct {
	ID           int32  `bson:"_id,omitempty"` // Custom ID field
	CollectionID int32  `bson:"collection_id"`
	Data         bson.M `bson:"data"`
	Username     string `bson:"username"`
}

// Define a struct to represent the document data in JSON input
type DocumentInput struct {
	Data         map[string]string `json:"data"`
	CollectionID int32             `json:"collectionId"`
}
