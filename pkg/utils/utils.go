package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func BsonObjectID(objectId string) primitive.ObjectID {
	result, _ := primitive.ObjectIDFromHex(objectId)
	return result
}
