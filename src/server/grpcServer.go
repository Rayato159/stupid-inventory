package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Rayato159/stupid-inventory/src/config"
	pbItem "github.com/Rayato159/stupid-inventory/src/proto/item"
	"go.mongodb.org/mongo-driver/mongo"
)

type grpcServer struct {
	cfg *config.Config
	db  *mongo.Client
	pbItem.UnimplementedItemServiceServer
}

func NewGrpcServer(cfg *config.Config, db *mongo.Client) *grpcServer {
	return &grpcServer{
		cfg: cfg,
		db:  db,
	}
}

func (s *grpcServer) StartItemServer() {
	var opts []grpc.ServerOption
	gs := grpc.NewServer(opts...)
	log.Printf("%s server is starting on %s", s.cfg.App.AppName, s.cfg.App.Url)
	lis, err := net.Listen("tcp", s.cfg.App.Url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// pbItem.RegisterItemServiceServer(
	// 	grpcServer,
	// )
	gs.Serve(lis)
}
