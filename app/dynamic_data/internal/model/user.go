package model

// User Model
type User struct {
	Username   string `json:"username,omitempty" bson:"username,omitempty" validate:"required,alpha"`
	LastDocId  int32  `json:"last_doc_id,omitempty" bson:"last_doc_id,omitempty" validate:"required"`
	LastCollId int32  `json:"last_coll_id,omitempty" bson:"last_coll_id,omitempty" validate:"required"`
}
