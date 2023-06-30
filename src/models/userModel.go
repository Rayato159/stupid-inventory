package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ObjectId primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username string             `bson:"username" json:"username"`
	Items    []Item             `bson:"items,omitempty" json:"items"`
}

// Many to many -> users *-* items
type UserItem struct {
	UserId string `bson:"user_id"`
	ItemId string `bson:"item_id"`
}
