package main

import (
	"context"

	"github.com/Rayato159/stupid-inventory/pkg/db"
	"github.com/Rayato159/stupid-inventory/src/config"
	"github.com/Rayato159/stupid-inventory/src/server"
)

func main() {
	cfg := config.NewConfig("./.env.grpc.item")

	dbClient := db.DBConn(cfg)
	defer dbClient.Disconnect(context.Background())

	server.NewGrpcServer(cfg, dbClient).StartItemServer()
}
