package common

import "go.mongodb.org/mongo-driver/bson"

func ConvertToBsonM(v interface{}) (bson.M, error) {
	data, err := bson.Marshal(v)
	var bsonDoc bson.M
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(data, &bsonDoc)
	if err != nil {
		return nil, err
	}

	return bsonDoc, nil
}
