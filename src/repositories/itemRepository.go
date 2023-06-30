package repositories

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepository struct {
	client *mongo.Client
}

func (r *ItemRepository) FindItem(ctx context.Context) (any, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	coll := r.client.Database("item_db").Collection("items")

	cursor, err := coll.Find(ctx, bson.D{}, nil)
	if err != nil {
		return nil, fmt.Errorf("find items failed: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

	}

	return nil, nil
}
