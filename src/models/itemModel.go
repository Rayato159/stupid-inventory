package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ObjectId    primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Title       string      `bson:"title" json:"title"`
	Description string      `bson:"description" json:"description"`
	Damage      float64     `bson:"damage" json:"damage"`
}
