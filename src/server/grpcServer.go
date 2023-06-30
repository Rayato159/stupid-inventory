package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Rayato159/stupid-inventory/src/config"
	"github.com/Rayato159/stupid-inventory/src/controllers/grpcControllers"
	pbItem "github.com/Rayato159/stupid-inventory/src/proto/item"
	"github.com/Rayato159/stupid-inventory/src/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type grpcServer struct {
	cfg      *config.Config
	dbClient *mongo.Client
}

func NewGrpcServer(cfg *config.Config, dbClient *mongo.Client) *grpcServer {
	return &grpcServer{
		cfg:      cfg,
		dbClient: dbClient,
	}
}

func (s *grpcServer) StartItemServer() {
	opts := make([]grpc.ServerOption, 0)
	gs := grpc.NewServer(opts...)

	log.Printf("%s server is starting on %s", s.cfg.App.AppName, s.cfg.App.Url)
	lis, err := net.Listen("tcp", s.cfg.App.Url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pbItem.RegisterItemServiceServer(
		gs,
		&grpcControllers.ItemGrpcController{
			ItemRepository: &repositories.ItemRepository{
				Client: s.dbClient,
			},
		},
	)
	gs.Serve(lis)
}
