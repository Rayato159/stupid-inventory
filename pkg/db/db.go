package db

import (
	"context"
	"log"
	"time"

	"github.com/Rayato159/stupid-inventory/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBConn(cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Db.Url))
	if err != nil {
		log.Fatalf("connect to db -> %s failed: %v", cfg.Db.Url, err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("pinging to db -> %s failed: %v", cfg.Db.Url, err)
	}
	log.Printf("connected to db -> %s", cfg.Db.Url)
	return client
}
