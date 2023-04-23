package util

import "gopkg.in/mgo.v2/bson"

func ConvertMapToBsonM(m map[string]string) bson.M {
	bsonM := bson.M{}
	for key, value := range m {
		bsonM[key] = value
	}
	return bsonM
}

func ConvertBsonMToMap(bsonM bson.M) map[string]string {
	m := make(map[string]string)
	for key, value := range bsonM {
		m[key] = value.(string)
	}
	return m
}
