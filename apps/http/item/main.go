package main

import (
	"context"
	"os"

	"github.com/Rayato159/stupid-inventory/pkg/db"
	"github.com/Rayato159/stupid-inventory/src/config"
	"github.com/Rayato159/stupid-inventory/src/server"
)

func main() {
	cfg := config.NewConfig(func() string {
		if len(os.Args) > 1 {
			return os.Args[1]
		}
		return "./.env.http.item"
	}())

	dbClient := db.DBConn(cfg)
	defer dbClient.Disconnect(context.Background())

	server.NewHttpServer(cfg, dbClient).StartItemServer()
}
