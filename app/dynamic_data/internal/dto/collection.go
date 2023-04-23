package dto

type Collection struct {
	Id        int32
	Name      string
	DataKeys  []string
	DataTypes []int32
}

type ListCollections struct {
	Collections []Collection
	Documents   []Document
}

type GetCollection struct {
	Id        int32
	Name      string
	DataKeys  []string
	DataTypes []int32
	Documents []Document
}
