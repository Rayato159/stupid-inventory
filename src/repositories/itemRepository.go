package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Rayato159/stupid-inventory/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepository struct {
	Client *mongo.Client
}

func (r *ItemRepository) FindItems(ctx context.Context, items *[]models.Item) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	coll := r.Client.Database("item_db").Collection("items")

	cursor, err := coll.Find(ctx, bson.D{}, nil)
	if err != nil {
		return fmt.Errorf("find items failed: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item models.Item
		if err := cursor.Decode(&item); err != nil {
			return fmt.Errorf("decode item failed: %v", err)
		}
		*items = append(*items, item)
	}
	return nil
}

func (r *ItemRepository) FindOneItem(ctx context.Context, itemId string) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	itemObjectId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, err
	}

	item := new(models.Item)
	if err := r.Client.Database("item_db").Collection("items").FindOne(ctx, bson.M{"_id": itemObjectId}, nil).Decode(item); err != nil {
		return nil, fmt.Errorf("find one item id: %v failed: %v", itemObjectId, err)
	}
	return item, nil
}
