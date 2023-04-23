package model

type Collection struct {
	ID          int32    `bson:"_id,omitempty"` // Custom ID field
	Name        string   `bson:"name"`
	SemanticKey string   `bson:"semantic_key"`
	Keys        []string `bson:"keys"`
	Types       []int32  `bson:"types"`
	DocumentIDs []int32  `bson:"document_ids,omitempty"`
	Username    string   `bson:"username"`
}
