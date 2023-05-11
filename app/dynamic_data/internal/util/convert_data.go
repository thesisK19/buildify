package util

import (
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/mgo.v2/bson"
)

func ConvertMapToBsonM(m map[string]interface{}) bson.M {
	bsonM := bson.M{}
	for k, v := range m {
		bsonM[k] = v
	}
	return bsonM
}

func ConvertBsonMToMap(bsonM bson.M) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range bsonM {
		m[k] = v
	}
	return m
}

func ConvertStructpbValueMapToBsonM(structpbValueMap map[string]*structpb.Value) bson.M {
	bsonM := bson.M{}
	for k, v := range structpbValueMap {
		bsonM[k] = v.AsInterface()
	}
	return bsonM
}

func ConvertBsonMToStructpbValueMap(bsonM bson.M) (map[string]*structpb.Value, error) {
	m := make(map[string]*structpb.Value)
	for k, v := range bsonM {
		value, err := structpb.NewValue(v)
		if err != nil {
			return nil, err
		}
		m[k] = value
	}
	return m, nil
}

func ConvertStructpbValueMapToStringMap(structpbValueMap map[string]*structpb.Value) map[string]interface{} {
	interfaceMap := make(map[string]interface{})
	for k, v := range structpbValueMap {
		interfaceMap[k] = v.AsInterface()
	}
	return interfaceMap
}

func ConvertStringMapToStructpbValueMap(interfaceMap map[string]interface{}) (map[string]*structpb.Value, error) {
	structpbValueMap := make(map[string]*structpb.Value)
	for k, v := range interfaceMap {
		value, err := structpb.NewValue(v)
		if err != nil {
			return nil, err
		}
		structpbValueMap[k] = value
	}
	return structpbValueMap, nil
}

func ConvertStructpbValueMapToInt32Map(structpbValueMap map[int32]*structpb.Value) map[int32]interface{} {
	interfaceMap := make(map[int32]interface{})
	for k, v := range structpbValueMap {
		interfaceMap[k] = v.AsInterface()
	}
	return interfaceMap
}

func ConvertInt32MapToStructpbValueMap(interfaceMap map[int32]interface{}) (map[int32]*structpb.Value, error) {
	structpbValueMap := make(map[int32]*structpb.Value)
	for k, v := range interfaceMap {
		value, err := structpb.NewValue(v)
		if err != nil {
			return nil, err
		}
		structpbValueMap[k] = value
	}
	return structpbValueMap, nil
}


// func ConvertPbStructToBsonM(pbStruct *structpb.Struct) bson.M {
// 	mapData := pbStruct.AsMap()
// 	bsonM := bson.M{}
// 	for key, value := range mapData {
// 		bsonM[key] = value
// 	}
// 	return bsonM
// }

// func ConvertBsonMToPbStruct(bsonM bson.M) (*structpb.Struct, error) {
// 	mapData := ConvertBsonMToMap(bsonM)
// 	return structpb.NewStruct(mapData)
// }