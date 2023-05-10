package dto

type Collection struct {
	Id        int32
	Name      string
	ProjectId string
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
	ProjectId string
	DataKeys  []string
	DataTypes []int32
	Documents []Document
}
