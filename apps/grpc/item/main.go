package main

import (
	"github.com/Rayato159/stupid-inventory/pkg/db"
	"github.com/Rayato159/stupid-inventory/src/config"
	"github.com/Rayato159/stupid-inventory/src/server"
)

func main() {
	cfg := config.NewConfig("./.env.grpc.item")

	dbClient := db.DBConn(cfg)

	server.NewGrpcServer(cfg, dbClient).StartItemServer()
}
