package main

import (
	"github.com/Rayato159/stupid-inventory/pkg/db"
	"github.com/Rayato159/stupid-inventory/src/config"
	"github.com/Rayato159/stupid-inventory/src/server"
)

func main() {
	cfg := config.NewConfig("./.env.http.item")

	dbClient := db.DBConn(cfg)

	server.NewHttpServer(cfg, dbClient).StartItemServer()
}
